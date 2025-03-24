package services

import (
	"my-production-app/models"
	"my-production-app/repositories"
	"my-production-app/utils"
)

// GetUserMenu 사용자 이메일과 그룹에 맞는 메뉴 데이터 반환
func GetUserMenu(email, group string) ([]models.MenuItem, error) {
	if email == "" || group == "" {
		return nil, utils.ErrMissingParameters
	}

	menus, err := repositories.FetchUserMenu(email, group)
	if err != nil {
		return nil, utils.ErrDatabaseError
	}

	return menus, nil
}
