package services

import (
	"my-production-app/models"
	"my-production-app/repositories"
	"my-production-app/utils"
)

// CreateAdminUserService 새로운 관리자 사용자 생성 및 JWT 반환
func CreateAdminUserService(admin models.UserInsertRequest) (string, error) {
	// 비밀번호 해싱
	hashedPassword, err := utils.HashPassword(admin.Password)
	if err != nil {
		return "", err
	}
	admin.Password = hashedPassword

	// DB에 사용자 추가
	adminID, err := repositories.CreateAdminUser(admin)
	if err != nil {
		return "", err
	}

	// 생성된 사용자의 이메일 조회
	email, err := repositories.GetAdminEmailByID(adminID)
	if err != nil {
		return "", err
	}

	// JWT 생성
	token, err := utils.GenerateToken(email)
	if err != nil {
		return "", err
	}

	return token, nil
}
