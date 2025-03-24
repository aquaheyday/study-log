# 🔀 Go 언어 라우팅과 핸들러

Go에서는 기본 패키지 `net/http`를 이용해 **간단한 라우팅과 핸들러 등록**이 가능합니다.  
또한, `http.ServeMux` 또는 외부 라우팅 프레임워크를 사용하여 더 유연한 라우팅도 구현할 수 있습니다.

---

## 1️⃣ 기본 핸들러 등록 (`http.HandleFunc`)

```go
package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, Go!")
}

func main() {
    http.HandleFunc("/", helloHandler)
    http.ListenAndServe(":8080", nil)
}
```

✔ `http.HandleFunc`는 URL 경로에 함수를 매핑  
✔ `ResponseWriter`로 응답 작성, `*Request`로 요청 정보 접근  

---

## 2️⃣ ServeMux로 라우팅 세분화

```go
func home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "홈 페이지입니다")
}

func about(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "소개 페이지입니다")
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
    mux.HandleFunc("/about", about)
    http.ListenAndServe(":8080", mux)
}
```

✔ `http.NewServeMux()`로 라우터 객체 생성  
✔ `mux.HandleFunc`으로 라우팅 구성  
✔ 같은 prefix가 있는 라우트는 먼저 등록된 핸들러가 우선됨  

---

## 3️⃣ http.Handler 인터페이스 구현

핸들러는 단순 함수뿐 아니라 **인터페이스 기반 구조체**로도 구현 가능

```go
type HelloHandler struct{}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello from struct handler")
}

func main() {
    http.Handle("/hello", HelloHandler{})
    http.ListenAndServe(":8080", nil)
}
```

✔ `ServeHTTP(w http.ResponseWriter, r *http.Request)` 메서드 구현 시 핸들러로 사용 가능  
✔ 미들웨어나 공통 로직을 구조화할 때 유용  

---

## 4️⃣ 라우팅 규칙과 우선순위

```go
mux := http.NewServeMux()
mux.HandleFunc("/", rootHandler) // fallback (가장 마지막에 호출됨)  
mux.HandleFunc("/admin", adminHandler) //prefix 매칭 
mux.HandleFunc("/admin/settings", settingsHandler) // 정확 매칭 
```

✔ 경로는 **접두사 기반(prefix match)** 으로 처리됨  
✔ `/admin/`으로 등록하면 `/admin/settings`까지 매칭 가능

---

## 5️⃣ URL 파라미터 처리 (기본 방식은 없음)

`net/http` 기본 라우터는 **동적 URL 파라미터 처리 기능이 없음**  
예를 들어 `/user/123` 같은 경로는 직접 파싱해야 함

```go
func userHandler(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path // 예: /user/123
    id := strings.TrimPrefix(path, "/user/")
    fmt.Fprintf(w, "User ID: %s", id)
}
```

✔ 간단한 경우에는 직접 문자열 파싱  
✔ 복잡한 라우팅은 외부 패키지 사용 추천

---

## 6️⃣ 외부 라우팅 라이브러리 예시 (gorilla/mux)

```go
import (
    "github.com/gorilla/mux"
)

func getUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    fmt.Fprintf(w, "User ID: %s", id)
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/user/{id}", getUser)
    http.ListenAndServe(":8080", r)
}
```

✔ `gorilla/mux`는 동적 라우팅, 메서드 매칭, 서브라우터 등을 지원  
✔ `{id}` 같은 URL 파라미터 매핑 가능  

---

## 7️⃣ HTTP 메서드 매핑

```go
func methodHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        fmt.Fprintln(w, "GET 요청입니다")
    case http.MethodPost:
        fmt.Fprintln(w, "POST 요청입니다")
    default:
        http.Error(w, "허용되지 않은 메서드", http.StatusMethodNotAllowed)
    }
}
```

✔ `r.Method`로 HTTP 메서드 확인  
✔ RESTful API 구현 시 필수 패턴  

---

## 🎯 정리

✔ 기본 `net/http`만으로도 간단한 라우팅 처리 가능  
✔ 복잡한 라우팅, 파라미터, 미들웨어 체인은 `gorilla/mux`, `chi`, `gin` 등의 라우터 사용 권장  
✔ 라우팅은 **서버 구조의 중심**이므로, 잘 설계하면 유지보수와 확장성이 매우 좋아짐
