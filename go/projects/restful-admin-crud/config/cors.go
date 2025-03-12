package config

import (
	"os"
	"strings"

	"github.com/rs/cors"
)

// CORS 설정
func LoadCORSConfig() *cors.Cors {
	allowedOrigins := os.Getenv("CORS_ALLOWED_ORIGINS")
	if allowedOrigins == "" {
		allowedOrigins = "https://ppeum.com,http://admin-hub.ppeum.com,https://admin-hub.ppeum.com,http://localhost:*"
	}

	return cors.New(cors.Options{
		AllowedOrigins:   strings.Split(allowedOrigins, ","),
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
	})
}
