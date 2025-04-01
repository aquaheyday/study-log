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

// GetFranchiseUserLeadsHandler retrieves paginated franchise user leads
// @Summary 가맹점 문의 목록 조회
// @Description 페이징을 지원하여 가맹점 문의 목록을 가져옵니다
// @Tags 가맹점 문의
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT Token" default(Bearer <token>)
// @Param page query int false "페이지 번호 (기본값: 1)"
// @Param limit query int false "페이지당 항목 수 (기본값: 10)"
// @Success 200 {object} models.SuccessResponse{data=models.FranchiseUserLeadsPagination}
// @Failure 400 {object} models.ErrorResponse400 "잘못된 요청 형식"
// @Failure 500 {object} models.ErrorResponse500 "서버 오류"
// @Router /v1/franchise-user-leads [get]
func GetFranchiseUserLeadsHandler(w http.ResponseWriter, r *http.Request) {
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

	// DB에서 가맹점 문의 목록 조회
	inquiries, totalItems, totalPages, err := repositories.GetFranchiseUserLeadsFromDB(database.DB, page, limit)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// 공통 Pagination 구조체 활용
	response := models.FranchiseUserLeadsPagination{
		CurrentPage: page,
		Limit:       limit,
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		Inquiries:   inquiries,
	}

	utils.SendJSONResponse(w, http.StatusOK, models.SuccessResponse{
		Message: "가맹점 문의 조회 성공",
		Data:    response,
	})
}

// InsertFranchiseUserLeads handles inserting a new franchise user lead
// @Summary 가맹점 문의 삽입
// @Description 새로운 가맹점 문의를 추가합니다
// @Tags 가맹점 문의
// @Accept json
// @Produce json
// @Param request body models.FranchiseUserLeads true "가맹점 문의 요청 데이터"
// @Success 201 {object} models.FranchiseUserLeadsInsertResponse
// @Failure 400 {object} models.ErrorResponse400 "잘못된 요청 형식"
// @Failure 500 {object} models.ErrorResponse500 "서버 오류"
// @Router /v1/franchise-user-leads [post]
func InsertFranchiseUserLeads(w http.ResponseWriter, r *http.Request) {
	var franchiseUserLeads models.FranchiseUserLeads

	// 요청 본문에서 데이터 읽기
	err := json.NewDecoder(r.Body).Decode(&franchiseUserLeads)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "잘못된 요청 형식입니다: "+err.Error())
		return
	}

	// DB에 데이터 등록
	insertedID, err := repositories.InsertFranchiseUserLead(database.DB, franchiseUserLeads)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// 응답 반환
	utils.SendJSONResponse(w, http.StatusCreated, models.SuccessResponse{
		Message: "가맹점 문의가 성공적으로 추가되었습니다.",
		Data:    models.FranchiseUserLeadsInsertResponse{InsertedID: insertedID},
	})
}

// DeleteFranchiseUserLeads handles deleting franchise user leads
// @Summary 가맹점 문의 삭제
// @Description 가맹점 문의 데이터를 삭제 상태로 업데이트합니다
// @Tags 가맹점 문의
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT Token" default(Bearer <token>)
// @Param request body models.DeleteFranchiseUserLeadsRequest true "삭제할 가맹점 문의 ID 목록"
// @Success 200 {object} models.SuccessResponse{data=models.DeleteFranchiseUserLeadsResponse}
// @Failure 400 {object} models.ErrorResponse400 "잘못된 요청 형식"
// @Failure 500 {object} models.ErrorResponse500 "서버 오류"
// @Router /v1/franchise-user-leads [delete]
func DeleteFranchiseUserLeads(w http.ResponseWriter, r *http.Request) {
	// 요청 본문에서 데이터 읽기
	var requestPayload models.DeleteFranchiseUserLeadsRequest

	// 요청 본문 디코딩
	err := json.NewDecoder(r.Body).Decode(&requestPayload)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "잘못된 요청 형식입니다: "+err.Error())
		return
	}

	// ID 리스트가 비어 있는지 확인
	if len(requestPayload.IDs) == 0 {
		utils.SendErrorResponse(w, http.StatusBadRequest, "삭제할 ID가 제공되지 않았습니다.")
		return
	}

	// 컨텍스트에서 이메일 정보 가져오기
	email, ok := r.Context().Value("email").(string)
	if !ok {
		utils.SendErrorResponse(w, http.StatusUnauthorized, utils.ErrEmailClaimNotFound.Error())
		return
	}

	// DB에서 데이터 삭제 처리
	rowsAffected, err := repositories.DeleteFranchiseUserLeads(database.DB, email, requestPayload.IDs)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected == 0 {
		utils.SendErrorResponse(w, http.StatusNotFound, "제공된 ID와 일치하는 레코드가 없습니다.")
		return
	}

	// 응답 반환
	utils.SendJSONResponse(w, http.StatusCreated, models.SuccessResponse{
		Message: "가맹점 문의가 성공적으로 삭제되었습니다.",
		Data:    models.DeleteFranchiseUserLeadsResponse{RowsAffected: rowsAffected},
	})
}
