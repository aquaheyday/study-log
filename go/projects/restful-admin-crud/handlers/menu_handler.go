package handlers

import (
	"my-production-app/models"
	"my-production-app/services"
	"my-production-app/utils"
	"net/http"
)

// MenuHandler 사용자 권한과 그룹을 기반으로 메뉴 목록 조회
// @Summary 메뉴 목록 조회
// @Description 사용자 권한 및 그룹 필터에 따라 메뉴 목록을 검색합니다.
// @Tags 메뉴
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT Token" default(Bearer <token>)
// @Param group query string true "메뉴 항목 필터링을 위한 그룹 식별 코드"
// @Success 200 {array} models.MenuItem "메뉴 항목 목록"
// @Failure 400 {object} models.ErrorResponse400 "잘못된 요청 형식"
// @Failure 500 {object} models.ErrorResponse500 "서버 오류"
// @Router /v1/menu [get]
func MenuHandler(w http.ResponseWriter, r *http.Request) {
	// 컨텍스트에서 이메일 정보 가져오기
	email, ok := r.Context().Value("email").(string)
	if !ok {
		utils.SendErrorResponse(w, http.StatusUnauthorized, utils.ErrEmailClaimNotFound.Error())
		return
	}

	// URL 파라미터에서 `group` 값 가져오기
	group := r.URL.Query().Get("group")
	if group == "" {
		utils.SendErrorResponse(w, http.StatusBadRequest, utils.ErrMissingParameters.Error())
		return
	}

	// 서비스 계층에서 메뉴 데이터 가져오기
	menus, err := services.GetUserMenu(email, group)
	if err != nil {
		switch err {
		case utils.ErrMissingParameters:
			utils.SendErrorResponse(w, http.StatusBadRequest, err.Error())
		case utils.ErrDatabaseError:
			utils.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		default:
			utils.SendErrorResponse(w, http.StatusInternalServerError, utils.ErrDatabaseError.Error())
		}
		return
	}

	// JSON 응답 반환
	response := models.SuccessResponse{
		Message: "Menu list retrieved successfully",
		Data:    menus,
	}

	utils.SendJSONResponse(w, http.StatusOK, response)
}
