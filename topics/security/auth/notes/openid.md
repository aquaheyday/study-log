# 🔐 OpenID Connect (OIDC)

**OpenID Connect**는 OAuth2 위에 **인증(Authentication)** 기능을 추가한 **표준 프로토콜**입니다.  
OAuth2는 기본적으로 **권한 위임(Authorization)** 만 처리할 수 있지만, OpenID Connect는 **사용자 정보 확인 및 로그인 처리**까지 확장해줍니다.

> OIDC = OAuth2 + ID 토큰 + 사용자 프로필 조회

---

## 1️⃣ 왜 OIDC가 필요한가?

| 항목                 | OAuth2만 사용했을 때         | OpenID Connect 도입 시               |
|----------------------|------------------------------|--------------------------------------|
| 로그인 처리          | 자체 구현 필요               | 표준 로그인 프로토콜 제공            |
| 사용자 정보 확인     | 별도 API 필요                | `id_token`으로 기본 정보 포함         |
| 인증 보장 수준       | 명확하지 않음                | 표준 명세에 따라 보장 가능            |
| 소셜 로그인 통합     | 구현 방식 제각각              | Google, Apple 등 OIDC 기반 통일 가능  |

---

## 2️⃣ 주요 개념

| 용어              | 설명 |
|-------------------|------|
| ID Token          | 사용자의 인증 정보를 담은 JWT (로그인 증명) |
| UserInfo Endpoint | 인증된 사용자의 상세 정보(JSON) 제공 |
| Scope             | `openid` 필수 + 추가 정보 권한 (`profile`, `email` 등) |
| OP (Provider)     | OpenID Provider, 예: Google, Apple |
| RP (Client)       | Relying Party, 즉 로그인 기능 사용하는 앱/서비스 |

---

## 3️⃣ 인증 흐름 (Authorization Code + OIDC)

OIDC는 일반적으로 **Authorization Code Flow**와 함께 사용됩니다.

1. 클라이언트 → `/authorize` 요청  
   - scope에 `openid` 포함  
2. 사용자 로그인 & 동의  
3. 리디렉션 → 인가 코드(code) 전달  
4. 클라이언트 → `/token` 요청  
   - 인가 코드로 `access_token` + `id_token` 획득  
5. (선택) `/userinfo` 엔드포인트에서 사용자 정보 조회  

---

## 4️⃣ 요청/응답 예시

### 1) 인가 요청

```http
GET /authorize?
  response_type=code
  &client_id=CLIENT_ID
  &redirect_uri=https://yourapp.com/callback
  &scope=openid profile email
  &state=abc123
```

---

### 2) 토큰 요청
```http
POST /token
Content-Type: application/x-www-form-urlencoded

grant_type=authorization_code
&code=AUTH_CODE
&redirect_uri=https://yourapp.com/callback
&client_id=CLIENT_ID
&client_secret=CLIENT_SECRET
```

---

### 3) 응답 예시
```json
{
  "access_token": "ACCESS_TOKEN",
  "id_token": "ID_TOKEN",
  "token_type": "Bearer",
  "expires_in": 3600
}
```

---

## 5️⃣ ID Token 구조
id_token은 JWT 형식이며, 다음과 같은 정보가 포함됩니다:

- sub: 사용자 고유 식별자
- name, email, picture: 사용자 프로필 정보 (scope에 따라)
- iss: 발급자 URL
- aud: 클라이언트 ID
- exp: 만료 시간

---

## 6️⃣ OIDC Scope 예시
| Scope | 설명 |
|---|---|
| openid | ✅ 필수. OIDC 사용을 나타냄 |
| profile	| 이름, 프로필 사진 등 |
| email	| 이메일 주소 |
| address	| 사용자 주소 |
| phone	| 전화번호 |

---

## 🎯 정리
✔ OpenID Connect는 OAuth2를 확장한 인증 프로토콜  
✔ id_token을 통해 로그인 여부, 사용자 정보 확인 가능  
✔ 소셜 로그인(Google, Apple 등) 대부분 OIDC 기반  
✔ OAuth2와 같은 플로우를 사용하되, 목적은 “인증”  
✔ scope에 openid가 포함되어야 OIDC 활성화됨
