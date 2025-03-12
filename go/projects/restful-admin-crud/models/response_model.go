package models

// SuccessResponse 일반적인 성공 응답 모델
type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // `omitempty`: 값이 없을 경우 JSON에서 제외
}

// TokenResponse 로그인/회원가입 시 반환할 응답 모델
type TokenResponse struct {
	Token string `json:"token"`
}

// ErrorResponse 표준 오류 응답 모델
type ErrorResponse struct {
	Code    int    `json:"code"`    // HTTP 상태 코드 (예: 400, 500)
	Message string `json:"message"` // 오류 메시지
}

// ErrorResponse400 스웨거 표출 400
type ErrorResponse400 struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message"`
}

// ErrorResponse500 스웨거 표출 500
type ErrorResponse500 struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message"`
}
