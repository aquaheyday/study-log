package utils

import (
	"encoding/json"
	"my-production-app/models"
	"net/http"
)

// JSON 응답을 전송하는 함수
func SendJSONResponse(w http.ResponseWriter, status int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

// 에러 응답을 전송하는 함수
func SendErrorResponse(w http.ResponseWriter, status int, message string) {
	SendJSONResponse(w, status, models.ErrorResponse{
		Message: message,
	})
}
