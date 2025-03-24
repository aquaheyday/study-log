package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtSecret = []byte("your_secret_key")

// GenerateToken 사용자 이메일을 기반으로 JWT 토큰 생성
func GenerateToken(email string) (string, error) {
	// JWT 클레임 설정
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // 토큰 만료 시간 (24시간 후)
		"iat":   time.Now().Unix(),                     // 토큰 발행 시간
	}

	// JWT 토큰 생성
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 토큰 서명 (시크릿 키 사용)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", errors.New("failed to sign token")
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	// 토큰 파싱 및 검증
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 알고리즘 검증
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	// 검증 실패 시 에러 반환
	if err != nil {
		return nil, err
	}

	// 토큰 클레임 추출
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
