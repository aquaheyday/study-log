package handlers

import (
	"encoding/json"
	"my-production-app/database"
	"my-production-app/models"
	"my-production-app/repositories"
	"my-production-app/utils"
	"net/http"
	"strconv"
)

// GetPermissionApprovalsRequests handles paginated permission request retrieval
// @Summary 권한 승인 요청 목록 조회
// @Description 페이징을 지원하여 권한 요청 목록을 검색합니다
// @Tags 권한 승인 요청
// @Accept json
// @Produce json
// @Param page query int false "페이지 번호 (기본값: 1)"
// @Param limit query int false "페이지당 항목 수 (기본값: 10)"
// @Success 200 {object} models.SuccessResponse{data=models.PermissionApprovalsRequestsPagination}
// @Failure 400 {object} models.ErrorResponse400 "잘못된 요청 형식"
// @Failure 500 {object} models.ErrorResponse500 "서버 오류"
// @Router /v1/admin/permission-approvals-requests [get]
func GetPermissionApprovalsRequests(w http.ResponseWriter, r *http.Request) {
	// GET 요청에서 page와 limit 값 받기
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	// 기본값 설정
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	// DB에서 권한 요청 목록 조회
	permissions, totalItems, totalPages, err := repositories.GetPermissionApprovalsRequestsFromDB(database.DB, page, limit)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// 페이징 정보 포함 응답
	response := models.PermissionApprovalsRequestsPagination{
		CurrentPage: page,
		Limit:       limit,
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		Requests:    permissions,
	}

	utils.SendJSONResponse(w, http.StatusOK, models.SuccessResponse{
		Message: "권한 요청 목록 조회 성공",
		Data:    response,
	})
}

// SetPermissionApprovalsRequest 권한 요청 승인 또는 거절
// @Summary 권한 요청 승인 또는 거절
// @Description 관리자가 권한 요청을 승인 또는 거절합니다
// @Tags 권한 승인 요청
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT Token" default(Bearer <token>)
// @Param request body models.PermissionApprovalRequest true "권한 승인 또는 거절 요청"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse400 "잘못된 요청 형식"
// @Failure 500 {object} models.ErrorResponse500 "서버 오류"
// @Router /v1/admin/permission-approvals [patch]
func SetPermissionApprovalsRequest(w http.ResponseWriter, r *http.Request) {
	// 컨텍스트에서 이메일 정보 가져오기
	email, ok := r.Context().Value("email").(string)
	if !ok {
		utils.SendErrorResponse(w, http.StatusUnauthorized, utils.ErrEmailClaimNotFound.Error())
		return
	}

	// 요청 메서드 확인
	if r.Method != http.MethodPatch {
		utils.SendErrorResponse(w, http.StatusMethodNotAllowed, "허용되지 않은 메서드입니다.")
		return
	}

	// 요청 데이터 파싱
	var approvalRequest models.PermissionApprovalRequest
	err := json.NewDecoder(r.Body).Decode(&approvalRequest)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "잘못된 요청 형식입니다: "+err.Error())
		return
	}

	// 요청 ID 및 상태 검증
	if approvalRequest.RequestID == "0" || (approvalRequest.Status != "approved" && approvalRequest.Status != "rejected") {
		utils.SendErrorResponse(w, http.StatusBadRequest, "유효하지 않은 요청입니다.")
		return
	}

	// 권한 요청 존재 여부 확인
	exists, err := repositories.CheckPermissionRequestExists(database.DB, approvalRequest.RequestID)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	if !exists {
		utils.SendErrorResponse(w, http.StatusNotFound, "해당 권한 요청을 찾을 수 없습니다.")
		return
	}

	// 권한 요청 상세 정보 조회
	userEmail, permissionID, err := repositories.GetPermissionRequestDetails(database.DB, approvalRequest.RequestID)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// 승인 또는 거절 처리
	if approvalRequest.Status == "approved" {
		err = repositories.ApprovePermissionRequest(database.DB, approvalRequest.RequestID, email, userEmail, permissionID)
	} else {
		err = repositories.RejectPermissionRequest(database.DB, approvalRequest.RequestID, email, userEmail, permissionID)
	}
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// 성공 응답 반환
	utils.SendJSONResponse(w, http.StatusOK, models.SuccessResponse{
		Message: "권한 요청이 성공적으로 처리되었습니다.",
		Data:    approvalRequest,
	})
}
