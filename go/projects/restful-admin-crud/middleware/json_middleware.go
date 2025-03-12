package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"my-production-app/utils"
	"net/http"
)

// JSONValidationMiddleware - Content-Type 및 JSON 유효성 검사 미들웨어
func JSONValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1️⃣ Content-Type 검사
		if r.Header.Get("Content-Type") != "application/json" {
			utils.SendErrorResponse(w, http.StatusUnsupportedMediaType, "Content-Type must be application/json")
			return
		}

		// 2️⃣ 요청 바디 읽기 (바디 복사본 생성)
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Failed to read request body: %v", err)
			utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid request body")
			return
		}

		// 3️⃣ JSON 형식 검사
		var temp map[string]interface{}
		err = json.Unmarshal(bodyBytes, &temp)
		if err != nil {
			log.Printf("Failed to decode JSON: %v", err)
			utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid JSON format")
			return
		}

		// ✅ 4️⃣ JSON이 올바르면 `r.Body`를 복원하여 핸들러에서 다시 사용할 수 있도록 설정
		r.Body = io.NopCloser(bytes.NewReader(bodyBytes))

		// 5️⃣ 다음 핸들러 실행
		next.ServeHTTP(w, r)
	})
}
