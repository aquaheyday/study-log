package middleware

import (
	"context"
	"my-production-app/utils"
	"net/http"
	"strings"
)

// AuthMiddleware는 인증된 사용자만 핸들러를 호출하도록 한다.
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Authorization 헤더 확인
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			utils.SendErrorResponse(w, http.StatusUnauthorized, utils.ErrUnauthorizedAccess.Error())
			return
		}

		// "Bearer " 접두어 제거
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// 토큰 검증
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			utils.SendErrorResponse(w, http.StatusUnauthorized, utils.ErrTokenValidationFailed.Error())
			return
		}

		// 이메일 클레임 추출
		email, ok := claims["email"].(string)
		if !ok {
			utils.SendErrorResponse(w, http.StatusUnauthorized, utils.ErrEmailClaimNotFound.Error())
			return
		}

		// 요청 컨텍스트에 이메일 저장
		ctx := r.Context()
		ctx = context.WithValue(ctx, "email", email)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
