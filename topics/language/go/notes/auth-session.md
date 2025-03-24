# 🔐 Go 언어 인증과 세션 관리

Go는 `net/http`, `cookie`, 외부 라이브러리 등을 통해 세션 기반 또는 토큰 기반 인증 구현이 가능합니다.

---

## 1️⃣ 기본 로그인 흐름

```go
func loginHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    username := r.FormValue("username")
    password := r.FormValue("password")

    if username == "admin" && password == "pass123" {
        http.SetCookie(w, &http.Cookie{
            Name:  "session",
            Value: "some-session-id",
            Path:  "/",
        })
        fmt.Fprintln(w, "Login successful")
    } else {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
    }
}
```

✔ 로그인 성공 시 쿠키에 세션 ID 저장  
✔ 이후 요청에서 쿠키로 사용자 식별 가능  

---

## 2️⃣ 세션 확인 미들웨어

```go
func withAuth(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        cookie, err := r.Cookie("session")
        if err != nil || cookie.Value != "some-session-id" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        next(w, r)
    }
}
```

✔ 쿠키 값으로 인증 상태 확인  
✔ 공통 인증 검증 로직을 미들웨어로 분리  

---

## 3️⃣ 외부 세션 라이브러리 (e.g. gorilla/sessions)

```bash
go get github.com/gorilla/sessions
```

```go
var store = sessions.NewCookieStore([]byte("secret-key"))

func loginHandler(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "auth-session")
    session.Values["authenticated"] = true
    session.Save(r, w)
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "auth-session")
    if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    }
    fmt.Fprintln(w, "Welcome!")
}
```

✔ `gorilla/sessions`는 세션을 쿠키 또는 서버에 저장 가능  
✔ `session.Values`에 필요한 값 저장 후 `Save()` 호출  

---

## 4️⃣ 세션 종료 (로그아웃)

```go
func logoutHandler(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "auth-session")
    session.Options.MaxAge = -1 // 즉시 만료
    session.Save(r, w)
    fmt.Fprintln(w, "Logged out")
}
```

✔ `MaxAge = -1`로 설정 시 쿠키 삭제  

---

## 5️⃣ JWT 기반 인증 (토큰 방식)

```bash
go get github.com/golang-jwt/jwt/v5
```

### 토큰 생성

```go
token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "username": "admin",
    "exp":      time.Now().Add(time.Hour).Unix(),
})

tokenString, err := token.SignedString([]byte("my-secret-key"))
```

### 토큰 검증

```go
parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    return []byte("my-secret-key"), nil
})

if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
    fmt.Println("Username:", claims["username"])
} else {
    fmt.Println("Invalid token")
}
```

✔ JWT는 클라이언트가 토큰을 **Authorization 헤더**로 보내는 방식  
✔ 장점: **무상태(stateless)**, 서버에 세션 저장 불필요  
✔ 단점: 토큰 탈취 시 위험 → HTTPS + 만료시간 필수

---

## 6️⃣ Authorization 헤더에서 토큰 추출

```go
func extractToken(r *http.Request) string {
    authHeader := r.Header.Get("Authorization")
    if strings.HasPrefix(authHeader, "Bearer ") {
        return strings.TrimPrefix(authHeader, "Bearer ")
    }
    return ""
}
```

✔ JWT는 보통 `Authorization: Bearer <token>` 형식으로 전달됨  

---

## 7️⃣ 세션 vs JWT 비교

| 항목 | 세션 기반 | JWT 기반 |
|------|-----------|----------|
| 저장 위치 | 서버(메모리/DB) 또는 쿠키 | 클라이언트 (토큰 자체에 정보 포함) |
| 서버 확장성 | 낮음 (세션 공유 필요) | 높음 (무상태) |
| 보안 위험 | 세션 탈취 | 토큰 탈취 |
| 상태 유지 | 상태ful | 무상태 (stateless) |
| 복호화 필요 | ❌ | 필요 (서명 검증) |

---

## 8️⃣ 실전에서의 구조 분리 추천

```
project/
├── main.go
├── auth/
│   ├── session.go
│   └── jwt.go
├── handler/
│   ├── login.go
│   └── protected.go
└── middleware/
    └── auth.go
```

---

## 🎯 정리

✔ 모든 인증 관련 요청은 **HTTPS** 필수  
✔ 쿠키엔 `HttpOnly`, `Secure`, `SameSite` 속성 설정 추천  
✔ JWT에는 너무 민감한 정보 담지 말 것  
✔ 토큰/세션의 만료시간 반드시 설정  
✔ 인증 미들웨어는 라우팅 전 처리 (router-level middleware)
