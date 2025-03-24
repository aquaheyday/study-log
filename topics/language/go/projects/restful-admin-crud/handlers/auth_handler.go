package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"my-production-app/database"
	"my-production-app/models"
	"my-production-app/repositories"
	"my-production-app/services"
	"my-production-app/utils"
	"net/http"
	"strconv"
)

// LoginHandler 로그인 요청을 처리하는 핸들러
// @Summary 사용자 로그인
// @Description 사용자를 인증하고 권한 정보와 JWT 토큰을 반환합니다.
// @Tags 사용자
// @Accept json
// @Produce json
// @Param request body models.LoginRequest true "사용자 로그인 요청"
// @Success 200 {object} models.LoginResponse "로그인 성공"
// @Failure 400 {object} models.ErrorResponse400 "잘못된 요청 형식"
// @Failure 500 {object} models.ErrorResponse500 "서버 오류"
// @Router /v1/login [post]
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// 요청 데이터 디코딩
	var loginReq models.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, utils.ErrInvalidJSON.Error())
		return
	}

	// 이메일이 비어있는지 확인
	if loginReq.Email == "" {
		utils.SendErrorResponse(w, http.StatusBadRequest, utils.ErrEmptyEmail.Error())
		return
	}

	// 비밀번호가 비어있는지 확인
	if loginReq.Password == "" {
		utils.SendErrorResponse(w, http.StatusBadRequest, utils.ErrEmptyPassword.Error())
		return
	}

	// 서비스 계층에서 로그인 처리
	data, err := services.AuthenticateUser(loginReq.Email, loginReq.Password)
	if err != nil {
		if errors.Is(err, utils.ErrEmailNotFound) {
			utils.SendErrorResponse(w, http.StatusUnauthorized, utils.ErrEmailNotFound.Error())
		} else if errors.Is(err, utils.ErrPasswordMismatch) {
			utils.SendErrorResponse(w, http.StatusUnauthorized, utils.ErrPasswordMismatch.Error())
		} else if errors.Is(err, utils.ErrTokenGeneration) {
			utils.SendErrorResponse(w, http.StatusUnauthorized, utils.ErrTokenGeneration.Error())
		} else if errors.Is(err, utils.ErrInsufficientPermissions) {
			utils.SendErrorResponse(w, http.StatusUnauthorized, utils.ErrInsufficientPermissions.Error())
		} else {
			utils.SendErrorResponse(w, http.StatusInternalServerError, utils.ErrInvalidRequest.Error())
		}
		return
	}

	response := models.SuccessResponse{
		Message: "login successfully",
		Data:    data,
	}

	// JSON 응답 반환
	utils.SendJSONResponse(w, http.StatusOK, response)
}

// UserInsert handles admin user creation
// @Summary 사용자 계정 생성
// @Description 새로운 사용자 계정을 생성하고 JWT 토큰을 반환합니다
// @Tags 사용자
// @Accept json
// @Produce json
// @Param request body models.UserInsertRequest true "사용자 계정 생성 요청"
// @Success 201 {object} models.SuccessResponse{data=models.TokenResponse}
// @Failure 400 {object} models.ErrorResponse400 "잘못된 요청 형식"
// @Failure 500 {object} models.ErrorResponse500 "서버 오류"
// @Router /v1/admin/user [post]
func UserInsert(w http.ResponseWriter, r *http.Request) {
	var admin models.UserInsertRequest

	// 요청 본문에서 데이터 읽기
	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, utils.ErrRequestBodyEmpty.Error())
		return
	}

	// 서비스 호출
	token, err := services.CreateAdminUserService(admin)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, utils.ErrAdminUserCreationFailed.Error())
		return
	}

	// JSON 응답 반환
	response := models.SuccessResponse{
		Message: "Admin user created successfully",
		Data:    models.TokenResponse{Token: token},
	}

	utils.SendJSONResponse(w, http.StatusCreated, response)
}

// GetPermissions
// @Summary 사용자 권한 목록 조회
// @Description 모든 사용자 권한을 검색합니다
// @Tags 사용자
// @Accept json
// @Produce json
// @Success 200 {object} models.SuccessResponse{data=[]models.Permission}
// @Failure 400 {object} models.ErrorResponse400 "잘못된 요청 형식"
// @Failure 500 {object} models.ErrorResponse500 "서버 오류"
// @Router /v1/admin/permissions [get]
func GetPermissions(w http.ResponseWriter, r *http.Request) {
	permissions, err := repositories.GetPermissionsFromDB(database.DB)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, utils.ErrAdminUserCreationFailed.Error())
		return
	}

	// JSON 응답 반환
	response := models.SuccessResponse{
		Message: "Permissions retrieved successfully",
		Data:    permissions,
	}

	utils.SendJSONResponse(w, http.StatusOK, response)
}

// InsertPermissions
// @Summary 사용자 권한 요청
// @Description 특정 계정에 대한 사용자 권한을 요청합니다
// @Tags 사용자
// @Accept json
// @Produce json
// @Param email query string true "사용자 계정"
// @Param permission query int true "권한 ID"
// @Success 201 {object} models.SuccessResponse "새로운 요소 생성"
// @Failure 400 {object} models.ErrorResponse400 "잘못된 요청 형식"
// @Failure 500 {object} models.ErrorResponse500 "서버 오류"
// @Router /v1/admin/permissions [post]
func InsertPermissions(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	permissionIDStr := r.URL.Query().Get("permission")

	if email == "" || permissionIDStr == "" {
		utils.SendErrorResponse(w, http.StatusBadRequest, utils.ErrMissingEmailOrPermission.Error())
		return
	}

	// string -> int 변환
	permissionID, err := strconv.Atoi(permissionIDStr)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, utils.ErrInvalidPermissionID.Error())
		return
	}

	// DB에 권한 요청 추가
	err = repositories.InsertPermission(database.DB, email, permissionID)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSONResponse(w, http.StatusCreated, models.SuccessResponse{
		Message: "계정 권한 요청이 완료되었습니다.",
	})
}

// GetPermissionGroups
// @Summary 사용자 권한 그룹 목록 조회
// @Description 모든 사용자 권한 그룹을 검색합니다
// @Tags 사용자
// @Accept json
// @Produce json
// @Success 200 {object} models.SuccessResponse{data=[]models.PermissionGroups}
// @Failure 400 {object} models.ErrorResponse400 "잘못된 요청 형식"
// @Failure 500 {object} models.ErrorResponse500 "서버 오류"
// @Router /v1/admin/permission-groups [get]
func GetPermissionGroups(w http.ResponseWriter, r *http.Request) {
	permissionGroups, err := repositories.GetPermissionGroupsFromDB(database.DB)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, models.SuccessResponse{
		Message: "권한 그룹 목록 조회 성공",
		Data:    permissionGroups,
	})
}
