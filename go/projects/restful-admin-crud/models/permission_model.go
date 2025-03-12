package models

// Permission 모델
type Permission struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Sort        int    `json:"sort"`
	Role        string `json:"role"`
	Description string `json:"description"`
}

type PermissionGroups struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Sort        int    `json:"sort"`
	Description string `json:"description"`
	Permissions []int  `json:"permissions"`
}

// PermissionApprovalsRequests 권한 요청 모델
type PermissionApprovalsRequests struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Role        string `json:"role"`
	Status      string `json:"status"`
	RequestedAt string `json:"requested_at"`
}

// PermissionApprovalsRequestsPagination 페이징을 포함한 응답 모델
type PermissionApprovalsRequestsPagination struct {
	CurrentPage int                           `json:"current_page"`
	Limit       int                           `json:"limit"`
	TotalItems  int                           `json:"total_items"`
	TotalPages  int                           `json:"total_pages"`
	Requests    []PermissionApprovalsRequests `json:"requests"`
}

// PermissionApprovalRequest 권한 승인 요청 모델
type PermissionApprovalRequest struct {
	RequestID string `json:"request_id"`
	Status    string `json:"status" enums:"approved,rejected" example:"approved" default:"approved" binding:"required"`
}

// PermissionApproval 권한 승인 모델
type PermissionApproval struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	PermissionID int    `json:"permission_id"`
	Status       string `json:"status"`
	AdminEmail   string `json:"admin_email"`
	ApprovedAt   string `json:"approved_at,omitempty"`
	RejectedAt   string `json:"rejected_at,omitempty"`
}
