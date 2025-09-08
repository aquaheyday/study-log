package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gocql/gocql"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
)

var clients = make(map[*websocket.Conn]bool)
var clientsMutex = sync.Mutex{}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var rdb *redis.Client
var kafkaWriter *kafka.Writer
var cassandraSession *gocql.Session

const redisChannel = "chat_messages"
const kafkaTopic = "chat_messages_storage"

func main() {
	initRedis()
	initKafka()
	initCassandra()

	go handleRedisMessages()

	http.HandleFunc("/ws", handleConnections)
	log.Println("Chat server started on :8080")
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

	clientsMutex.Lock()
	clients[ws] = true
	clientsMutex.Unlock()
	log.Println("Client connected")

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			clientsMutex.Lock()
			delete(clients, ws)
			clientsMutex.Unlock()
			break
		}

		// Redis publish
		err = rdb.Publish(context.Background(), redisChannel, msg).Err()
		if err != nil {
			log.Println("Redis publish error:", err)
		}

		// Kafka write
		err = kafkaWriter.WriteMessages(context.Background(), kafka.Message{Value: msg})
		if err != nil {
			log.Println("Kafka write error:", err)
		}

		// Cassandra insert
		if cassandraSession != nil {
			err := cassandraSession.Query(`INSERT INTO chat.messages (id, room_id, sender, content, created_at)
				VALUES (?, ?, ?, ?, ?)`,
				gocql.TimeUUID(), "default", "anonymous", string(msg), time.Now(),
			).Exec()
			if err != nil {
				log.Println("Cassandra insert error:", err)
			}
		}
	}
}

func handleRedisMessages() {
	pubsub := rdb.Subscribe(context.Background(), redisChannel)
	defer pubsub.Close()

	ch := pubsub.Channel()
	for msg := range ch {
		clientsMutex.Lock()
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
			if err != nil {
				log.Println("Write error:", err)
				client.Close()
				delete(clients, client)
			}
		}
		clientsMutex.Unlock()
	}
}

// --- 초기화 함수들 ---
func initRedis() {
	redisAddr := os.Getenv("REDIS_ADDR")
	rdb = redis.NewClient(&redis.Options{Addr: redisAddr})
	log.Println("Connected to Redis")
}

func initKafka() {
	kafkaBrokers := os.Getenv("KAFKA_BROKERS")
	kafkaWriter = &kafka.Writer{
		Addr:     kafka.TCP(kafkaBrokers),
		Topic:    kafkaTopic,
		Balancer: &kafka.LeastBytes{},
	}
	log.Println("Connected to Kafka")
}

func initCassandra() {
	host := os.Getenv("CASSANDRA_HOSTS")
	keyspace := os.Getenv("CASSANDRA_KEYSPACE")
	user := os.Getenv("CASSANDRA_USER")
	pass := os.Getenv("CASSANDRA_PASSWORD")

	for i := 0; i < 10; i++ {
		cluster := gocql.NewCluster(host)
		cluster.Keyspace = keyspace
		cluster.Consistency = gocql.Quorum
		cluster.Authenticator = gocql.PasswordAuthenticator{
			Username: user,
			Password: pass,
		}

		session, err := cluster.CreateSession()
		if err == nil {
			cassandraSession = session
			log.Println("Connected to Cassandra")
			return
		}
		log.Println("Cassandra connect retry:", err)
		time.Sleep(5 * time.Second)
	}
	log.Fatal("Could not connect to Cassandra")
}
