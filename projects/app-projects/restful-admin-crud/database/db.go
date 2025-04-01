package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

// InitializeDB initializes the MySQL connection
func InitializeDB() {
	// .env 파일 로드
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables.")
	}

	// DSN(Data Source Name) 생성
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		getEnv("DB_USERNAME", "root"),  // 기본값: root
		getEnv("DB_PASSWORD", ""),      // 기본값: 비밀번호 없음
		getEnv("DB_HOST", "127.0.0.1"), // 기본값: localhost
		getEnv("DB_PORT", "3306"),      // 기본값: 3306
		getEnv("DB_NAME", "testdb"),    // 기본값: testdb
	)

	// MySQL 연결
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 연결 확인
	err = DB.Ping()
	if err != nil {
		log.Printf("Database is not reachable: %v", err)
		DB = nil // 연결 실패 시 DB를 nil로 설정
	}

	fmt.Println("Connected to MySQL successfully!")
}

// getEnv는 환경 변수를 읽어오거나 기본값을 반환합니다.
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
