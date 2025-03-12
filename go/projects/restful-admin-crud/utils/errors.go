package utils

import "errors"

// 공통 오류 메시지
var (
	// 인증 관련
	ErrInvalidCredentials    = errors.New("잘못된 로그인 정보입니다.")
	ErrTokenGeneration       = errors.New("토큰 생성에 실패했습니다.")
	ErrTokenValidationFailed = errors.New("토큰 검증에 실패했습니다.")
	ErrDatabaseError         = errors.New("데이터베이스 오류가 발생했습니다.")

	// 이메일 및 비밀번호 유효성 검사
	ErrInvalidEmail     = errors.New("이메일 형식이 올바르지 않습니다.")
	ErrEmptyEmail       = errors.New("이메일을 입력해주세요.")
	ErrEmailNotFound    = errors.New("등록되지 않은 이메일입니다.")
	ErrInvalidPassword  = errors.New("비밀번호 형식이 올바르지 않습니다.")
	ErrEmptyPassword    = errors.New("비밀번호를 입력해주세요.")
	ErrPasswordMismatch = errors.New("비밀번호가 일치하지 않습니다.")

	// 클레임 관련 오류
	ErrEmailClaimNotFound = errors.New("토큰에서 이메일 정보를 찾을 수 없습니다.")
	ErrInvalidEmailClaim  = errors.New("토큰의 이메일 클레임이 올바르지 않습니다.")

	// 권한 관련
	ErrUnauthorizedAccess      = errors.New("로그인이 필요합니다.")
	ErrInsufficientPermissions = errors.New("권한이 부족합니다.")
	ErrPermissionDenied        = errors.New("이 기능을 사용할 수 없습니다.")

	// 요청 관련
	ErrInvalidJSON       = errors.New("잘못된 JSON 형식입니다.")
	ErrRequestBodyEmpty  = errors.New("요청 본문이 비어 있습니다.")
	ErrInvalidRequest    = errors.New("잘못된 요청입니다.")
	ErrResourceNotFound  = errors.New("요청하신 정보를 찾을 수 없습니다.")
	ErrMissingParameters = errors.New("필수 요청 파라미터가 누락되었습니다.")
	ErrInvalidParameters = errors.New("잘못된 요청 파라미터입니다.")

	// 계정 관련
	ErrAdminUserCreationFailed = errors.New("관리자 계정 생성에 실패했습니다.")

	// 권한 요청 관련
	ErrMissingEmailOrPermission = errors.New("이메일과 권한 ID는 필수 입력 항목입니다.")
	ErrInvalidPermissionID      = errors.New("권한 ID는 숫자로 입력해야 합니다.")
)
