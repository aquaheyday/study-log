package repositories

import (
	"database/sql"
	"errors"
	"my-production-app/models"
	"strconv"
	"strings"
)

// GetPermissionsFromDB retrieves permission data from the database
func GetPermissionsFromDB(db *sql.DB) ([]models.Permission, error) {
	query := `
		SELECT 
			id,
			COALESCE(name, '') AS name,
			type,
			sort,
			role,
			COALESCE(description, '') AS description
		FROM 
			permissions 
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, errors.New("failed to retrieve permissions: " + err.Error())
	}
	defer rows.Close()

	var permissions []models.Permission

	for rows.Next() {
		var permission models.Permission
		err := rows.Scan(&permission.ID, &permission.Name, &permission.Type, &permission.Sort, &permission.Role, &permission.Description)
		if err != nil {
			return nil, errors.New("failed to parse permission data: " + err.Error())
		}
		permissions = append(permissions, permission)
	}

	// rows 처리 중 에러 확인
	if err := rows.Err(); err != nil {
		return nil, errors.New("error while processing rows: " + err.Error())
	}

	return permissions, nil
}

// InsertPermission inserts a new permission request for an admin user
func InsertPermission(db *sql.DB, email string, permission int) error {
	query := "INSERT INTO admin_user_permission_requests (email, permission_id) VALUES (?, ?)"
	_, err := db.Exec(query, email, permission)
	if err != nil {
		return errors.New("failed to insert permission request: " + err.Error())
	}
	return nil
}

// GetPermissionGroupsFromDB 데이터베이스에서 권한 그룹을 검색
func GetPermissionGroupsFromDB(db *sql.DB) ([]models.PermissionGroups, error) {
	query := `
		SELECT 
		    group_permissions.id,
			COALESCE(name, '') AS name,
			sort,
			COALESCE(description, '') AS description,
			GROUP_CONCAT(group_permissions.permission_id) as permissions
		FROM 
			permission_groups
		JOIN group_permissions ON group_permissions.group_id = permission_groups.id
		GROUP BY permission_groups.id
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, errors.New("권한 그룹 데이터 조회 실패: " + err.Error())
	}
	defer rows.Close()

	var permissionGroups []models.PermissionGroups

	for rows.Next() {
		var permission models.PermissionGroups
		var permissionsString string

		err := rows.Scan(&permission.ID, &permission.Name, &permission.Sort, &permission.Description, &permissionsString)
		if err != nil {
			return nil, errors.New("권한 그룹 데이터 변환 실패: " + err.Error())
		}

		// 쉼표(,) 기준으로 권한 ID를 나누고 int 배열로 변환
		permissionsArr := strings.Split(permissionsString, ",")
		for _, p := range permissionsArr {
			if permissionID, err := strconv.Atoi(p); err == nil {
				permission.Permissions = append(permission.Permissions, permissionID)
			}
		}
		permissionGroups = append(permissionGroups, permission)
	}

	// rows 처리 중 에러 확인
	if err := rows.Err(); err != nil {
		return nil, errors.New("권한 그룹 데이터 처리 중 오류 발생: " + err.Error())
	}

	return permissionGroups, nil
}

// GetPermissionRequestsFromDB 데이터베이스에서 페이지별로 구분된 권한 요청 목록을 조회
func GetPermissionApprovalsRequestsFromDB(db *sql.DB, page, limit int) ([]models.PermissionApprovalsRequests, int, int, error) {
	// 기본값 설정
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	// SQL 쿼리를 위한 offset 계산
	offset := (page - 1) * limit

	query := `
		SELECT 
		    aupr.id,
		    aupr.email,
		    per.name,
		    per.type,
		    per.role,
		    aupr.status,
		    aupr.requested_at
		FROM 
			admin_user_permission_requests aupr
		JOIN permissions per ON per.id = aupr.permission_id
		ORDER BY 
		    aupr.id DESC
		LIMIT ? OFFSET ?
	`

	rows, err := db.Query(query, limit, offset)
	if err != nil {
		return nil, 0, 0, errors.New("권한 요청 데이터 조회 실패:" + ": " + err.Error())
	}
	defer rows.Close()

	// 전체 항목 수 계산
	var total int
	err = db.QueryRow("SELECT COUNT(*) FROM admin_user_permission_requests").Scan(&total)
	if err != nil {
		return nil, 0, 0, errors.New("권한 요청 데이터 카운터 조회 실패:" + ": " + err.Error())
	}

	totalPages := (total + limit - 1) / limit // 전체 페이지 수 계산

	var permissions []models.PermissionApprovalsRequests

	// rows 데이터를 반복하며 PermissionRequests 구조체에 매핑
	for rows.Next() {
		var permission models.PermissionApprovalsRequests
		err := rows.Scan(&permission.ID, &permission.Email, &permission.Name, &permission.Type, &permission.Role, &permission.Status, &permission.RequestedAt)
		if err != nil {
			return nil, 0, 0, errors.New("데이터 변환 실패: " + err.Error())
		}
		permissions = append(permissions, permission)
	}

	// rows 처리 중 에러 확인
	if err := rows.Err(); err != nil {
		return nil, 0, 0, errors.New("데이터 처리 중 오류 발생: " + err.Error())
	}

	return permissions, total, totalPages, nil
}

// CheckPermissionRequestExists 요청이 존재하는지 확인
func CheckPermissionRequestExists(db *sql.DB, requestID string) (bool, error) {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM admin_user_permission_requests WHERE id = ?)", requestID).Scan(&exists)
	if err != nil {
		return false, errors.New("데이터 조회 실패:" + ": " + err.Error())
	}
	return exists, nil
}

// GetPermissionRequestDetails 권한 요청 정보 조회
func GetPermissionRequestDetails(db *sql.DB, requestID string) (string, int, error) {
	var email string
	var permissionID int
	err := db.QueryRow("SELECT email, permission_id FROM admin_user_permission_requests WHERE id = ?", requestID).
		Scan(&email, &permissionID)
	if err != nil {
		return "", 0, errors.New("데이터 조회 실패:" + ": " + err.Error())
	}
	return email, permissionID, nil
}

// ApprovePermissionRequest 권한 요청 승인
func ApprovePermissionRequest(db *sql.DB, requestID, adminEmail, userEmail string, permissionID int) error {
	// 승인 상태 업데이트
	updateQuery := `
		UPDATE admin_user_permission_requests
		SET status = 'approved', admin_email = ?, approved_at = NOW()
		WHERE id = ?
	`
	_, err := db.Exec(updateQuery, adminEmail, requestID)
	if err != nil {
		return errors.New("권한 요청 승인 중 오류 발생: " + err.Error())
	}

	// 승인된 권한을 admin_user_permissions 테이블에 추가
	insertQuery := "INSERT INTO admin_user_permissions (email, permission_id) VALUES (?, ?)"
	_, err = db.Exec(insertQuery, userEmail, permissionID)
	if err != nil {
		return errors.New("관리자 권한 추가 중 오류 발생: " + err.Error())
	}

	return nil
}

// RejectPermissionRequest 권한 요청 거절
func RejectPermissionRequest(db *sql.DB, requestID, adminEmail, userEmail string, permissionID int) error {
	// 거절 상태 업데이트
	updateQuery := `
		UPDATE admin_user_permission_requests
		SET status = 'rejected', admin_email = ?, rejected_at = NOW()
		WHERE id = ?
	`
	_, err := db.Exec(updateQuery, adminEmail, requestID)
	if err != nil {
		return errors.New("권한 요청 거절 중 오류 발생: " + err.Error())
	}

	// 거절된 경우 admin_user_permissions 테이블에서 삭제
	deleteQuery := "DELETE FROM admin_user_permissions WHERE email = ? AND permission_id = ?"
	_, err = db.Exec(deleteQuery, userEmail, permissionID)
	if err != nil {
		return errors.New("관리자 권한 삭제 중 오류 발생: " + err.Error())
	}

	return nil
}
