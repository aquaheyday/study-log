package models

// AdminUsers represents the admin_users model
type AdminUsers struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	IsActive string `json:"is_active"`
}

// LoginRequest 로그인 요청 데이터 구조체
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse 로그인 응답 데이터 구조체
type LoginResponse struct {
	Token         string `json:"token"`
	ApprovedCount int    `json:"approved_count"`
	PendingCount  int    `json:"pending_count"`
	RejectedCount int    `json:"rejected_count"`
}

type UserInsertRequest struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	IsActive string `json:"is_active"`
}
