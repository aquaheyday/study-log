package models

// MenuItem 메뉴 데이터 구조체
type MenuItem struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	Level    int    `json:"level"`
	Sort     int    `json:"sort"`
	ParentID string `json:"parent_id"`
	Role     string `json:"role"`
}
