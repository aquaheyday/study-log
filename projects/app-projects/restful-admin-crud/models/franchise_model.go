package models

// FranchiseUserLeads represents the franchise_user_leads model
type FranchiseUserLeads struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	BirthDate      string `json:"birth_date"`
	Gender         string `json:"gender"`
	PhoneNumber    string `json:"phone_number"`
	BranchLocation string `json:"branch_location"`
	Inquiry        string `json:"inquiry"`
	CreatedAt      string `json:"created_at"`
}

// FranchiseUserLeadsPagination represents the paginated response structure
type FranchiseUserLeadsPagination struct {
	CurrentPage int                  `json:"current_page"`
	Limit       int                  `json:"limit"`
	TotalItems  int                  `json:"total_items"`
	TotalPages  int                  `json:"total_pages"`
	Inquiries   []FranchiseUserLeads `json:"inquiries"`
}

// FranchiseUserLeadsInsertResponse 가맹점 등록 응답 모델
type FranchiseUserLeadsInsertResponse struct {
	InsertedID int64 `json:"inserted_id"`
}

// DeleteFranchiseUserLeadsRequest 가맹점 문의 삭제 요청 모델
type DeleteFranchiseUserLeadsRequest struct {
	IDs []int `json:"ids"`
}

// DeleteFranchiseUserLeadsResponse 가맹점 문의 삭제 응답 모델
type DeleteFranchiseUserLeadsResponse struct {
	RowsAffected int64 `json:"rows_affected"`
}
