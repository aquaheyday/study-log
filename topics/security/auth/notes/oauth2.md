# 🔐 OAuth2 개요

**OAuth2**는 **제3자 애플리케이션이 사용자 자원에 접근할 수 있도록 권한을 위임하는 인증 프레임워크**입니다.  
비밀번호를 공유하지 않고 안전하게 리소스에 접근할 수 있도록 합니다.

> 사용자 → 인증 서버에게 권한 허락 → 애플리케이션은 **토큰**으로 접근

---

## 1️⃣ 왜 OAuth2가 필요한가?

| 문제                           | OAuth2 도입 전                             | OAuth2 도입 후                             |
|--------------------------------|--------------------------------------------|---------------------------------------------|
| 비밀번호 공유                   | 서비스에 비밀번호를 직접 입력해야 함       | 토큰 기반 인증으로 비밀번호 노출 없음       |
| 권한 제어 어려움                | 서비스가 모든 권한을 가져감                | 범위(Scope)로 접근 권한을 세분화 가능       |
| 접근 철회 어려움                | 비밀번호 변경 외에 접근 차단 어려움        | 토큰 만료/철회로 접근 통제 가능             |

---

## 2️⃣ 주요 개념 정리

| 용어             | 설명 |
|------------------|------|
| **Resource Owner** | 자원의 실제 소유자 (예: 사용자) |
| **Client**         | 자원에 접근하려는 제3자 애플리케이션 |
| **Authorization Server** | 사용자 인증 및 토큰 발급 담당 |
| **Resource Server** | 보호된 자원이 있는 서버 (API 서버 등) |
| **Access Token**   | 자원에 접근할 수 있는 권한 증명 토큰 |
| **Refresh Token**  | 새로운 Access Token을 발급받기 위한 토큰 |

---

## 3️⃣ 인증 플로우 예시 (Authorization Code Grant)

1. **사용자 로그인 & 권한 승인**  
   → 사용자가 브라우저에서 OAuth2 서버에 로그인하고, 권한 범위(scope)를 승인

2. **Authorization Code 발급**  
   → 승인 후 리디렉션 URL로 **인가 코드(authorization code)** 전달

3. **Access Token 교환**  
   → 클라이언트는 인가 코드를 서버에 전달하고 **Access Token**과 **Refresh Token**을 발급받음

4. **API 요청 시 Access Token 사용**  
   → 클라이언트는 이 토큰으로 Resource Server에 안전하게 요청

---

## 4️⃣ 주요 Grant Type 비교

### 🔑 Grant Type 이란?
OAuth2에서 다양한 상황에 맞게 토큰을 발급받기 위해 **여러 가지 인증 흐름(grant type)** 을 제공합니다.

| Grant Type               | 사용 시나리오                     | 특징 |
|--------------------------|-----------------------------------|------|
| Authorization Code       | 웹 앱 (서버 사이드)              | 보안성 가장 높음, Refresh Token 사용 가능 |
| Implicit                 | SPA, 모바일 (현재는 비추천)       | 액세스 토큰만 발급, 보안 취약 가능성 ↑ |
| Resource Owner Password  | 신뢰 높은 앱 (테스트용)           | 사용자 비밀번호 직접 입력 → 비추 |
| Client Credentials       | 서버 간 통신                      | 사용자 개입 없이 클라이언트 자체 인증 |

---

## 5️⃣ 토큰 구조 예시 (JWT 기반)

```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "token_type": "Bearer",
  "expires_in": 3600,
  "refresh_token": "def502009ea3b1e..."
}
```

---

## 6️⃣ OAuth2 보안 전략

| 전략                              | 설명 |
|-----------------------------------|------|
| ✅ HTTPS 사용                      | 토큰 탈취 방지 필수 |
| ✅ Redirect URI 화이트리스트 등록  | 리디렉션 URL 위조 방지 |
| ✅ Scope 최소화                    | 필요한 권한만 요청 |
| ✅ 토큰 수명 짧게 설정             | 탈취 시 피해 최소화 |
| ✅ Refresh Token 보호              | HttpOnly + Secure 쿠키 등 활용 |

---

## 7️⃣ 관련 엔드포인트 예시

- `GET /authorize` → 인가 코드 발급
- `POST /token` → 액세스 토큰 + 리프레시 토큰 발급
- `POST /refresh` → 새 액세스 토큰 재발급 (리프레시 토큰 사용)
- `POST /revoke` → 토큰 무효화 처리

---

## 🎯 정리

✔ OAuth2는 **비밀번호 없이 권한 위임**을 가능하게 함  
✔ **액세스 토큰**으로 API 접근, **리프레시 토큰**으로 재발급  
✔ Grant Type에 따라 적절한 플로우 선택  
✔ 보안은 **HTTPS, Scope 제한, 토큰 관리**가 핵심
