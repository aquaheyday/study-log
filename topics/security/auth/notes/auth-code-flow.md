# 🔐 Authorization Code Flow

**Authorization Code Flow**는 OAuth2 인증 방식 중 가장 **보안성이 높고 일반적으로 권장되는 플로우**입니다.  
사용자가 로그인하면 **인가 코드(Authorization Code)** 를 받아서, 서버 측에서 이 코드를 사용해 **Access Token**과 **Refresh Token**을 안전하게 발급받습니다.

---

## 1️⃣ 언제 사용하나요?

| 사용 대상              | 이유 |
|------------------------|------|
| 서버 사이드 웹 앱       | 클라이언트 비밀(Client Secret) 안전하게 보관 가능 |
| 보안이 중요한 서비스   | 토큰 발급 과정이 사용자 노출 없이 진행됨 |
| 로그인 + 장기 인증 필요 | Refresh Token까지 안전하게 활용 가능 |

---

## 2️⃣ 인증 흐름 요약

1. **사용자 로그인 요청**  
   사용자가 애플리케이션에서 로그인 버튼 클릭  
   클라이언트는 Authorization Server로 리디렉션

2. **권한 승인**  
   사용자 로그인 & 권한 범위(Scope) 승인  
   서버는 인가 코드(Authorization Code)를 클라이언트에게 리디렉션 URL로 전달

3. **인가 코드 → 토큰 교환**  
   클라이언트는 받은 인가 코드를 Authorization Server에 전달  
   클라이언트 ID/Secret 포함  
   서버는 Access Token (+ Refresh Token)을 응답으로 발급

4. **API 요청 시 Access Token 사용**  
   클라이언트는 이 토큰으로 Resource Server(API 서버)에 요청

---

## 3️⃣ 흐름 다이어그램

```
사용자 ──▶ 클라이언트 ──▶ Authorization Server  
   ▲                                │  
   │         ◀── 인가 코드 ────────┘  
   │  
   └──────── Access Token 요청 ───▶ (인가 코드 사용)  
             ◀── Access Token 발급 ─┘  
```

---

## 4️⃣ 실제 요청 예시

### 1) 인가 코드 요청

```
GET /authorize?response_type=code  
  &client_id=CLIENT_ID  
  &redirect_uri=https://yourapp.com/callback  
  &scope=read_profile email  
  &state=xyz123
```

---

### 2) 토큰 요청 (서버 측)
```
POST /token  
Content-Type: application/x-www-form-urlencoded  

grant_type=authorization_code  
&code=AUTH_CODE_HERE  
&redirect_uri=https://yourapp.com/callback  
&client_id=CLIENT_ID  
&client_secret=CLIENT_SECRET
```

---

### 3) 응답 예시
```
{
  "access_token": "ACCESS_TOKEN_HERE",
  "token_type": "Bearer",
  "expires_in": 3600,
  "refresh_token": "REFRESH_TOKEN_HERE",
  "scope": "read_profile email"
}
```

---

## 5️⃣ 보안 팁

| 전략                         | 설명 |
|------------------------------|------|
| ✅ HTTPS 필수                 | 민감 정보 탈취 방지 |
| ✅ redirect_uri 고정 & 검증   | 악성 사이트 리디렉션 방지 |
| ✅ state 파라미터 사용       | CSRF 공격 방지 |
| ✅ Access Token은 짧게, Refresh Token은 안전하게 | 수명 관리 및 쿠키/DB 활용 |

---

## 🎯 정리

✔ **Authorization Code Flow**는 OAuth2의 기본이자 가장 안전한 방식  
✔ 클라이언트는 인가 코드만 받고, **토큰은 서버에서만 처리**  
✔ 사용자 정보는 절대 클라이언트에 노출되지 않음  
✔ 로그인, API 인증, 장기 세션 모두 대응 가능
