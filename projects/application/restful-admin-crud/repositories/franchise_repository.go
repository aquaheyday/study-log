package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"my-production-app/models"
	"strings"
)

// GetFranchiseUserLeadsFromDB retrieves paginated franchise user leads from the database
func GetFranchiseUserLeadsFromDB(db *sql.DB, page, limit int) ([]models.FranchiseUserLeads, int, int, error) {
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
		    id,
			COALESCE(name, '') AS name, 
			COALESCE(birth_date, '') AS birth_date, 
			COALESCE(gender, '') AS gender, 
			COALESCE(phone_number, '') AS phone_number, 
			COALESCE(branch_location, '') AS branch_location, 
			COALESCE(inquiry, '') AS inquiry, 
			created_at 
		FROM 
			franchise_user_leads 
		WHERE status = 'live'
		LIMIT ? OFFSET ?
	`
	rows, err := db.Query(query, limit, offset)
	if err != nil {
		return nil, 0, 0, errors.New("데이터 조회 실패:" + ": " + err.Error())
	}
	defer rows.Close()

	// 전체 항목 수 계산
	var total int
	err = db.QueryRow("SELECT COUNT(*) FROM franchise_user_leads WHERE status = 'live'").Scan(&total)
	if err != nil {
		return nil, 0, 0, errors.New("데이터 조회 실패:" + ": " + err.Error())
	}

	totalPages := (total + limit - 1) / limit // 전체 페이지 수 계산

	var inquiries []models.FranchiseUserLeads
	for rows.Next() {
		var inquiry models.FranchiseUserLeads
		err := rows.Scan(
			&inquiry.Id,
			&inquiry.Name,
			&inquiry.BirthDate,
			&inquiry.Gender,
			&inquiry.PhoneNumber,
			&inquiry.BranchLocation,
			&inquiry.Inquiry,
			&inquiry.CreatedAt,
		)
		if err != nil {
			return nil, 0, 0, errors.New("데이터 변환 실패: " + err.Error())
		}
		inquiries = append(inquiries, inquiry)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, 0, errors.New("데이터 처리 중 오류 발생: " + err.Error())
	}

	return inquiries, total, totalPages, nil
}

// InsertFranchiseUserLead 가맹점 문의 데이터 등록
func InsertFranchiseUserLead(db *sql.DB, lead models.FranchiseUserLeads) (int64, error) {
	query := `
		INSERT INTO franchise_user_leads (name, birth_date, gender, phone_number, branch_location, inquiry)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	result, err := db.Exec(
		query,
		lead.Name,
		lead.BirthDate,
		lead.Gender,
		lead.PhoneNumber,
		lead.BranchLocation,
		lead.Inquiry,
	)
	if err != nil {
		return 0, errors.New("데이터 등록 실패:" + ": " + err.Error())
	}

	// 등록된 행의 ID 반환
	insertedID, err := result.LastInsertId()
	if err != nil {
		return 0, errors.New("등록된 ID를 가져오는 데 실패했습니다: " + err.Error())
	}

	return insertedID, nil
}

// DeleteFranchiseUserLeads 가맹점 문의 삭제 (상태 업데이트)
func DeleteFranchiseUserLeads(db *sql.DB, adminEmail string, ids []int) (int64, error) {
	if len(ids) == 0 {
		return 0, errors.New("삭제할 ID가 없습니다.")
	}

	// 동적 IN 절을 위한 쿼리 구성
	placeholders := strings.Repeat(",?", len(ids)-1)
	query := fmt.Sprintf(
		"UPDATE franchise_user_leads SET status = 'delete', delete_at = CURRENT_TIMESTAMP, admin_email = ? WHERE id IN (?%s)",
		placeholders,
	)

	args := []interface{}{adminEmail}
	for _, id := range ids {
		args = append(args, id)
	}

	result, err := db.Exec(query, args...)
	if err != nil {
		return 0, errors.New("데이터 삭제 실패:" + ": " + err.Error())
	}

	// 삭제된 행 수 반환
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, errors.New("삭제된 행 수를 가져오는 데 실패했습니다: " + err.Error())
	}

	return rowsAffected, nil
}
