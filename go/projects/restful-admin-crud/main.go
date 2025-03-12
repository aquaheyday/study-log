package main

import (
	"log"
	"my-production-app/database"
	"my-production-app/routes"
	"net/http"

	_ "my-production-app/docs" // Swagger 문서 경로
)

// @title API-HUB
// @version 1.0
// @description API 제공 HUB 입니다.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	// 데이터베이스 초기화
	database.InitializeDB()

	// 라우터 설정
	r := routes.SetupRoutes()

	// 서버 실행
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
