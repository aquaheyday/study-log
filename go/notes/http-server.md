# Go 언어: HTTP 서버 만들기

Go 언어에서는 `net/http` 패키지를 사용하여 간단하게 HTTP 서버를 구축할 수 있습니다.

---

## 1. 기본 HTTP 서버 만들기

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

### 주요 패턴
- `http.HandleFunc("/", handler)`: 특정 경로에 대한 핸들러 등록
- `http.ListenAndServe(":8080", nil)`: 8080 포트에서 서버 실행
- `http.ResponseWriter`를 사용하여 응답 작성

---

## 2. 라우팅 (Routing)
Go의 `http.ServeMux`를 사용하여 URL별로 핸들러를 등록할 수 있습니다.

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

### 주요 패턴
- `http.NewServeMux()`: HTTP 요청을 라우팅할 멀티플렉서 생성
- `mux.HandleFunc("/route", handler)`: 경로별 핸들러 등록

---

## 3. 요청(Request) 처리
요청 정보를 가져오고, 쿼리 매개변수를 읽는 방법을 살펴보겠습니다.

```go
func handler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" {
        name = "Guest"
    }
    fmt.Fprintf(w, "Hello, %s!", name)
}
```

### 주요 패턴
- `r.URL.Query().Get("key")`: 쿼리 매개변수 읽기
- 기본값 처리 (`if name == "" {}`)

---

## 4. JSON 응답 처리
Go에서 JSON을 반환하는 방법입니다.

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

### 주요 패턴
- `w.Header().Set("Content-Type", "application/json")`: JSON 응답 설정
- `json.NewEncoder(w).Encode(response)`: JSON 변환 및 응답

---

## 5. HTTP 요청 처리 (POST 요청 데이터 받기)
POST 요청에서 데이터를 처리하는 방법입니다.

```go
package main

import (
    "fmt"
    "net/http"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }
    r.ParseForm()
    name := r.FormValue("name")
    fmt.Fprintf(w, "Received name: %s", name)
}

func main() {
    http.HandleFunc("/submit", postHandler)
    http.ListenAndServe(":8080", nil)
}
```

### 주요 패턴
- `r.Method != http.MethodPost`: 요청 메서드 확인
- `r.ParseForm()`: 폼 데이터 파싱
- `r.FormValue("key")`: POST 데이터 읽기

---

## 6. 미들웨어 활용
미들웨어를 사용하여 로깅이나 인증을 추가할 수 있습니다.

```go
package main

import (
    "fmt"
    "net/http"
    "time"
)

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        fmt.Printf("%s %s %s\n", r.Method, r.URL.Path, time.Since(start))
    })
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, Middleware!")
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", handler)

    wrappedMux := loggingMiddleware(mux)
    http.ListenAndServe(":8080", wrappedMux)
}
```

### 주요 패턴
- `http.Handler`를 감싸는 미들웨어 패턴
- 요청 시간 로깅 추가

---

## 7. 정적 파일 제공
정적 파일(HTML, CSS, JS)을 서빙하는 방법입니다.

```go
package main

import "net/http"

func main() {
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.ListenAndServe(":8080", nil)
}
```

### 주요 패턴
- `http.FileServer(http.Dir("./static"))`: 특정 디렉터리 서빙
- `http.StripPrefix("/static/", fs)`: 경로 조정

---

## 8. HTTPS 서버 만들기
TLS(SSL) 인증서를 사용하여 HTTPS 서버를 실행하는 방법입니다.

```go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, HTTPS!")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
}
```

### 주요 패턴
- `http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)`: HTTPS 서버 실행
- `cert.pem`, `key.pem` 파일 필요

---

## 9. HTTP 서버 베스트 프랙티스
**핸들러에서 `http.ResponseWriter` 설정을 먼저 수행**
**JSON 응답 시 `Content-Type`을 `application/json`으로 설정**
**미들웨어를 사용하여 로깅 및 인증 추가**
**보안 강화를 위해 HTTPS 사용**
**요청 크기 제한 및 타임아웃 설정 적용**
