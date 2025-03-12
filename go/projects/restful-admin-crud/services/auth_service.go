package services

import (
	"golang.org/x/crypto/bcrypt"
	"my-production-app/models"
	"my-production-app/repositories"
	"my-production-app/utils"
)

// AuthenticateUser 사용자를 인증하고 JWT 토큰을 반환하는 서비스
func AuthenticateUser(email, password string) (*models.LoginResponse, error) {
	// 데이터베이스에서 사용자 정보 조회
	dbPassword, err := repositories.GetUserPasswordByEmail(email)
	if err != nil {
		return nil, utils.ErrEmailNotFound
	}

	// 비밀번호 비교
	err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
	if err != nil {
		return nil, utils.ErrPasswordMismatch
	}

	// JWT 토큰 생성
	token, err := utils.GenerateToken(email)
	if err != nil {
		return nil, utils.ErrTokenGeneration
	}

	// 권한 개수 조회
	approvedCount, pendingCount, rejectedCount, err := repositories.GetUserPermissionCounts(email)
	if err != nil {
		return nil, utils.ErrInsufficientPermissions
	}

	// 응답 데이터 반환
	return &models.LoginResponse{
		Token:         token,
		ApprovedCount: approvedCount,
		PendingCount:  pendingCount,
		RejectedCount: rejectedCount,
	}, nil
}
