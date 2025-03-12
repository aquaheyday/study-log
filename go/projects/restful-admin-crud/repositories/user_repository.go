package repositories

import (
	"database/sql"
	"log"
	"my-production-app/database"
	"my-production-app/models"
)

// CreateAdminUser 새로운 관리자 계정 생성
func CreateAdminUser(admin models.UserInsertRequest) (int64, error) {
	query := "INSERT INTO admin_users (username, email, password) VALUES (?, ?, ?)"
	result, err := database.DB.Exec(query, admin.Username, admin.Email, admin.Password)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// GetAdminEmailByID 관리자 ID로 이메일 조회
func GetAdminEmailByID(adminID int64) (string, error) {
	var email string
	query := "SELECT email FROM admin_users WHERE id = ?"
	err := database.DB.QueryRow(query, adminID).Scan(&email)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}

	return email, nil
}

// GetUserPasswordByEmail 이메일로 사용자 비밀번호 조회
func GetUserPasswordByEmail(email string) (string, error) {
	var password string
	query := "SELECT password FROM admin_users WHERE email = ?"
	err := database.DB.QueryRow(query, email).Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("User not found: %s", email)
			return "", err
		}
		log.Printf("Database query error: %v", err)
		return "", err
	}
	return password, nil
}

// GetUserPermissionCounts 사용자의 승인/대기/거부된 권한 개수 조회
func GetUserPermissionCounts(email string) (int, int, int, error) {
	var approvedCount, pendingCount, rejectedCount int
	query := `
		SELECT 
			COALESCE(SUM(CASE WHEN status = 'approved' THEN 1 ELSE 0 END), 0),
			COALESCE(SUM(CASE WHEN status = 'pending' THEN 1 ELSE 0 END), 0),
			COALESCE(SUM(CASE WHEN status = 'rejected' THEN 1 ELSE 0 END), 0)
		FROM admin_user_permission_requests WHERE email = ?`
	err := database.DB.QueryRow(query, email).Scan(&approvedCount, &pendingCount, &rejectedCount)
	if err != nil {
		log.Printf("Database query error: %v", err)
		return 0, 0, 0, err
	}
	return approvedCount, pendingCount, rejectedCount, nil
}
