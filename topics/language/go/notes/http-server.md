# 🌐 Go 언어 기본 웹 서버

Go 언어에서는 `net/http` 패키지를 사용하여 간단하고 강력한 HTTP 서버를 구현할 수 있습니다.  
이 문서는 웹 서버 작성, 요청 처리, 정적 파일 서빙, JSON 응답, 미들웨어, 보안(HTTPS), 운영 설정까지 포괄합니다.

---

## 1️⃣ 기본 HTTP 서버 만들기

```go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Starting server on :8080")
    http.ListenAndServe(":8080", nil)
}
```

✔ `http.HandleFunc("/", handler)` : 특정 경로에 핸들러 등록  
✔ `http.ListenAndServe(":8080", nil)` : 기본 서버 실행  

---

## 2️⃣ 라우팅 (Routing)

```go
package main

import (
    "fmt"
    "net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the Home Page!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is the About Page.")
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", homeHandler)
    mux.HandleFunc("/about", aboutHandler)

    fmt.Println("Server running on :8080")
    http.ListenAndServe(":8080", mux)
}
```

✔ `http.NewServeMux()` : 라우팅 멀티플렉서 생성  
✔ `mux.HandleFunc("/path", handler)` : 경로별 처리  

---

## 3️⃣ 요청 정보 다루기

```go
func handler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" {
        name = "Guest"
    }
    fmt.Fprintf(w, "Hello, %s!", name)
}
```

✔ `r.URL.Query().Get("key")` : 쿼리 매개변수 읽기  
✔ 기본값 처리도 손쉽게 가능

---

## 4️⃣ POST 요청 및 폼 데이터 처리

```go
func postHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }
    r.ParseForm()
    name := r.FormValue("name")
    fmt.Fprintf(w, "Received name: %s", name)
}
```

✔ `r.Method`로 HTTP 메서드 확인  
✔ `r.ParseForm()` 후 `FormValue()`로 값 추출  

---

## 5️⃣ JSON 응답 처리

```go
package main

import (
    "encoding/json"
    "net/http"
)

type Response struct {
    Message string `json:"message"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
    response := Response{Message: "Hello, JSON!"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func main() {
    http.HandleFunc("/json", jsonHandler)
    http.ListenAndServe(":8080", nil)
}
```

✔ `Content-Type` 설정은 필수  
✔ `json.NewEncoder(w).Encode()`로 직렬화  

---

## 6️⃣ 정적 파일 제공

```go
func main() {
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    http.ListenAndServe(":8080", nil)
}
```

✔ `http.FileServer`로 파일 서버 생성  
✔ `http.StripPrefix`로 URL 경로 조정  

---

## 7️⃣ 미들웨어 패턴

```go
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("%s %s\n", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello with Middleware")
    })

    wrapped := loggingMiddleware(mux)
    http.ListenAndServe(":8080", wrapped)
}
```

✔ `http.Handler`를 감싸서 기능 추가  
✔ 로깅, 인증, 요청 제한 등 구현 가능  

---

## 8️⃣ HTTPS 서버 실행

```go
func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, HTTPS!")
    })
    http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
}
```

✔ 인증서 파일(`cert.pem`, `key.pem`) 필요  
✔ HTTPS는 기본 포트 443 사용  

---

## 9️⃣ 고급 서버 설정 (커스텀 http.Server)

```go
srv := &http.Server{
    Addr:         ":8080",
    ReadTimeout:  5 * time.Second,
    WriteTimeout: 10 * time.Second,
    IdleTimeout:  30 * time.Second,
    Handler:      nil,
}

srv.ListenAndServe()
```

✔ 시간 설정으로 연결 안정성 확보  
✔ Graceful shutdown 시에도 유용  

---

## 🎯 정리

| 항목 | 권장 사항 |
|------|-----------|
| 경로별 핸들러 | `ServeMux` 또는 라우팅 프레임워크 사용 |
| 응답 작성 | 항상 `Content-Type` 명시 |
| JSON 응답 | `encoding/json` 사용, 구조체 명세 지정 |
| 보안 | HTTPS 적용 (`ListenAndServeTLS`) |
| 미들웨어 | 로깅/인증/권한 체크 시 활용 |
| 운영 안정성 | 커스텀 서버 + 타임아웃 설정 적용 |
| 정적 파일 | `http.StripPrefix`로 경로 매칭 조정 |
