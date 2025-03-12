package routes

import (
	"my-production-app/config"
	"my-production-app/handlers"
	"my-production-app/middleware"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// SetupRoutes configures the API routes
func SetupRoutes() http.Handler {
	r := mux.NewRouter()

	// API v1 그룹
	v1 := r.PathPrefix("/v1").Subrouter()
	// 메뉴 관련
	v1.HandleFunc("/menu", middleware.AuthMiddleware(handlers.MenuHandler)).Methods("GET")
	// Swagger 경로 추가
	v1.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// 사용자 관련
	v1.HandleFunc("/admin/user", handlers.UserInsert).Methods("POST")
	v1.HandleFunc("/admin/permissions", handlers.GetPermissions).Methods("GET")
	v1.HandleFunc("/admin/permissions", handlers.InsertPermissions).Methods("POST")
	v1.HandleFunc("/admin/permission-groups", handlers.GetPermissionGroups).Methods("GET")

	// 권한 승인 관련
	v1.HandleFunc("/admin/permission-approvals-requests", handlers.GetPermissionApprovalsRequests).Methods("GET")
	v1.HandleFunc("/admin/permission-approvals", middleware.AuthMiddleware(handlers.SetPermissionApprovalsRequest)).Methods("PATCH")

	// 로그인 관련
	v1.Handle("/login", middleware.JSONValidationMiddleware(http.HandlerFunc(handlers.LoginHandler))).Methods("POST")

	// 가맹점 문의 접수
	v1.HandleFunc("/franchise-user-leads", middleware.AuthMiddleware(handlers.GetFranchiseUserLeadsHandler)).Methods("GET")
	v1.HandleFunc("/franchise-user-leads", handlers.InsertFranchiseUserLeads).Methods("POST")
	v1.HandleFunc("/franchise-user-leads", middleware.AuthMiddleware(handlers.DeleteFranchiseUserLeads)).Methods("DELETE")

	// CORS 미들웨어 추가
	corsHandler := config.LoadCORSConfig()
	// CORS를 포함한 라우터 반환
	return corsHandler.Handler(r)
}
