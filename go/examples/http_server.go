package main

import (
	"fmt"
	"net/http"
)

// 기본 핸들러 함수
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "안녕하세요! Go HTTP 서버입니다.")
}

// 경로 매개변수 예제 (URL에서 값 받기)
func helloHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("name") // URL 쿼리에서 name 값을 가져오기
	if query == "" {
		query = "손님"
	}
	fmt.Fprintf(w, "안녕하세요, %s!", query)
}

// JSON 응답 예제
func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // JSON 응답 설정
	jsonResponse := `{"message": "Go HTTP 서버에서 JSON 응답을 보냅니다."}`
	fmt.Fprintln(w, jsonResponse)
}

func main() {
	// 라우트 설정
	http.HandleFunc("/", homeHandler)       // 루트 경로
	http.HandleFunc("/hello", helloHandler) // 쿼리 매개변수 사용
	http.HandleFunc("/json", jsonHandler)   // JSON 응답

	// 서버 실행
	port := ":8080"
	fmt.Println("서버가 http://localhost" + port + " 에서 실행됩니다.")
	err := http.ListenAndServe(port, nil) // 서버 실행
	if err != nil {
		fmt.Println("서버 실행 오류:", err)
	}
}
