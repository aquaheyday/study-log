package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gocql/gocql" // 오타 수정: github.comcom -> github.com
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
)

// Message 구조체 (서버<->클라이언트 통신용)
type Message struct {
	Event   string `json:"event"`
	RoomID  string `json:"roomId"`
	Content string `json:"content"`
}

var clients = make(map[*websocket.Conn]string) // key: 클라이언트, value: 방 ID
var clientsMutex = sync.Mutex{}
var serverID = gocql.TimeUUID().String() // 각 서버 인스턴스에 고유 ID 부여

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}
var rdb *redis.Client
var kafkaWriter *kafka.Writer
var cassandraSession *gocql.Session

const kafkaTopic = "chat_messages_storage"

func main() {
	initRedis()
	initKafka()
	initCassandra()

	go subscribeToRooms()

	http.HandleFunc("/ws", handleConnections)
	log.Printf("[INFO] Server %s started on :8080", serverID)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer ws.Close()
	log.Printf("[CONN] Client connected to server %s", serverID)

	defer removeClient(ws)

	for {
		_, msgBytes, err := ws.ReadMessage()
		if err != nil {
			log.Printf("[CONN] Read error from client: %v", err)
			break
		}

		var msg Message
		if err := json.Unmarshal(msgBytes, &msg); err != nil {
			log.Println("[ERROR] Invalid message format:", err)
			continue
		}

		log.Printf("[RECV] Server %s received message for room '%s': %s", serverID, msg.RoomID, msg.Content)

		switch msg.Event {
		case "join":
			clientsMutex.Lock()
			clients[ws] = msg.RoomID
			clientsMutex.Unlock()
			log.Printf("[JOIN] Client %p joined room '%s' on server %s", ws, msg.RoomID, serverID)

		case "message":
			redisChannel := "room:" + msg.RoomID
			fullMessage, _ := json.Marshal(msg)

			err = rdb.Publish(context.Background(), redisChannel, fullMessage).Err()
			if err != nil {
				log.Println("[ERROR] Redis publish error:", err)
			} else {
				log.Printf("[PUB] Server %s published message to Redis channel '%s'", serverID, redisChannel)
			}

			err = kafkaWriter.WriteMessages(context.Background(), kafka.Message{Value: fullMessage})
			if err != nil {
				log.Printf("[ERROR] Kafka write error: %v", err)
			}
		}
	}
}

func subscribeToRooms() {
	pubsub := rdb.PSubscribe(context.Background(), "room:*")
	defer pubsub.Close()
	log.Printf("[SUB] Server %s subscribed to Redis 'room:*' channels", serverID)
	ch := pubsub.Channel()
	for redisMsg := range ch {
		log.Printf("[SUB] Server %s received message from Redis channel '%s'", serverID, redisMsg.Channel)
		clientsMutex.Lock()
		roomID := strings.TrimPrefix(redisMsg.Channel, "room:")
		sentCount := 0
		for client, clientRoomID := range clients {
			if clientRoomID == roomID {
				err := client.WriteMessage(websocket.TextMessage, []byte(redisMsg.Payload))
				if err != nil {
					log.Printf("[ERROR] Write error to client in room '%s': %v", roomID, err)
					client.Close()
					delete(clients, client)
				} else {
					sentCount++
				}
			}
		}
		clientsMutex.Unlock()
		log.Printf("[SEND] Server %s sent message to %d clients in room '%s'", serverID, sentCount, roomID)
	}
}

func removeClient(ws *websocket.Conn) {
	clientsMutex.Lock()
	delete(clients, ws)
	clientsMutex.Unlock()
	log.Printf("[CONN] Client disconnected from server %s", serverID)
}

func initRedis() {
	redisAddr := os.Getenv("REDIS_ADDR")
	rdb = redis.NewClient(&redis.Options{Addr: redisAddr})
	log.Println("Connected to Redis")
}

func initKafka() {
	kafkaBroker := os.Getenv("KAFKA_BROKERS")
	conn, err := kafka.Dial("tcp", kafkaBroker)
	if err != nil {
		panic(fmt.Sprintf("failed to dial kafka broker: %v", err))
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		panic(fmt.Sprintf("failed to get kafka controller: %v", err))
	}

	controllerConn, err := kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(fmt.Sprintf("failed to dial kafka controller: %v", err))
	}
	defer controllerConn.Close()

	topicConfig := kafka.TopicConfig{
		Topic:             kafkaTopic,
		NumPartitions:     1,
		ReplicationFactor: 1,
	}

	err = controllerConn.CreateTopics(topicConfig)
	if err != nil {
		if e, ok := err.(kafka.Error); ok && e == kafka.TopicAlreadyExists {
			log.Printf("Kafka topic '%s' already exists, skipping creation.", kafkaTopic)
		} else {
			panic(fmt.Sprintf("failed to create kafka topic: %v", err))
		}
	} else {
		log.Printf("Kafka topic '%s' created successfully.", kafkaTopic)
	}

	kafkaWriter = &kafka.Writer{
		Addr:     kafka.TCP(kafkaBroker),
		Topic:    kafkaTopic,
		Balancer: &kafka.LeastBytes{},
	}
	log.Println("Connected to Kafka")
}

func initCassandra() {
	cassandraHost := os.Getenv("CASSANDRA_HOSTS")
	keyspace := os.Getenv("CASSANDRA_KEYSPACE")
	for i := 0; i < 10; i++ {
		cluster := gocql.NewCluster(cassandraHost)
		cluster.Keyspace = keyspace
		cluster.Consistency = gocql.Quorum
		var err error
		cassandraSession, err = cluster.CreateSession()
		if err == nil {
			log.Println("Connected to Cassandra")
			return
		}
		log.Printf("Could not connect to Cassandra, retrying in 5s... (%v)\n", err)
		time.Sleep(5 * time.Second)
	}
	log.Fatal("Could not connect to Cassandra after multiple retries")
}
