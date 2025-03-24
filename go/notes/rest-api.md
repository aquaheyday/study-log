# 🌐 Go 언어 REST API 개발

Go는 `net/http`, `encoding/json` 등 표준 라이브러리만으로도 RESTful API 서버를 쉽게 만들 수 있습니다.  
이 문서는 JSON 요청/응답, CRUD, 라우팅 구조 등 REST API 구축에 필요한 실전 예제를 다룹니다.

---

## 1️⃣ 기본 API 서버 구조

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "OK")
    })

    fmt.Println("API 서버 실행 중: http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
```

✔ `/health` 엔드포인트로 상태 확인 가능  
✔ 기본 포트 8080, 핸들러 함수 등록  

---

## 2️⃣ JSON 응답 처리

```go
package main

import (
    "encoding/json"
    "net/http"
)

type Message struct {
    Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    res := Message{Message: "Hello, API!"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(res)
}
```

✔ 구조체 → JSON으로 자동 인코딩  
✔ `w.Header().Set("Content-Type", "application/json")` 필수  

---

## 3️⃣ JSON 요청 파싱 (POST)

```go
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func createUser(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    json.NewEncoder(w).Encode(user)
}
```

✔ JSON 요청 바디를 구조체로 디코딩  
✔ 유효성 검사는 별도 처리 필요  

---

## 4️⃣ 간단한 CRUD 예제 (메모리 기반)

```go
var users = map[int]User{}
var idCounter = 1

func getUsers(w http.ResponseWriter, r *http.Request) {
    userList := []User{}
    for _, user := range users {
        userList = append(userList, user)
    }
    json.NewEncoder(w).Encode(userList)
}

func addUser(w http.ResponseWriter, r *http.Request) {
    var user User
    json.NewDecoder(r.Body).Decode(&user)

    user.ID = idCounter
    users[idCounter] = user
    idCounter++

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}
```

```go
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}
```

✔ 메모리 기반 저장소 사용  
✔ POST로 추가, GET으로 조회  

---

## 5️⃣ 라우팅 구조 분리

```go
func main() {
    http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            getUsers(w, r)
        case http.MethodPost:
            addUser(w, r)
        default:
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        }
    })

    fmt.Println("서버 실행: http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
```

✔ 한 경로에서 HTTP 메서드에 따라 처리 분기  
✔ `GET /users`, `POST /users` 형태 구현 가능  

---

## 6️⃣ 외부 라우터 사용 (gorilla/mux)

```go
import (
    "github.com/gorilla/mux"
)

func getUserByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    fmt.Fprintf(w, "ID: %s", id)
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/users/{id}", getUserByID).Methods("GET")
    http.ListenAndServe(":8080", r)
}
```

✔ `{id}` 같은 경로 파라미터 지원  
✔ `.Methods("GET")`으로 메서드 제한 가능  

---

## 7️⃣ 상태 코드 및 에러 응답

```go
func getUser(w http.ResponseWriter, r *http.Request) {
    id := 123 // 예시
    user, ok := users[id]
    if !ok {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(user)
}
```

✔ 에러 응답은 `http.Error()`로 처리  
✔ 상태 코드를 명확하게 지정할 것  

---

## 8️⃣ REST API 설계 기본 규칙

| HTTP 메서드 | 용도 | 예시 |
|-------------|------|------|
| GET    | 데이터 조회   | `/users`, `/users/1` |
| POST   | 새 데이터 추가 | `/users` |
| PUT    | 데이터 전체 수정 | `/users/1` |
| PATCH  | 데이터 일부 수정 | `/users/1` |
| DELETE | 데이터 삭제 | `/users/1` |

---

## 🎯 정리

✔ 항상 `Content-Type: application/json` 명시  
✔ URL은 명사 중심, 소문자 사용 (`/users`)  
✔ 상태 코드: `200`, `201`, `400`, `404`, `500` 명확히 사용  
✔ 인증/인가는 미들웨어로 분리하는 게 깔끔  
✔ 구조 분리: `router.go`, `handler.go`, `model.go`, `main.go` 등 추천  
✔ Go의 표준 패키지로도 REST API는 충분히 구현 가능하지만, 규모가 커지면 `gin`, `echo`, `fiber` 같은 프레임워크 사용도 고려해야함  
