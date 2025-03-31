# ⚠️ Implicit Flow (비추천)

**Implicit Flow**는 브라우저 기반 앱(SPA 등)에서 사용하던 OAuth2 인증 방식입니다.  
Access Token을 클라이언트 측에서 **직접 발급받는 방식**으로, **서버를 거치지 않기 때문에 보안에 취약**합니다.

> 현재는 보안상의 이유로 OAuth 공식 문서에서도 사용을 지양하며, **PKCE Flow**를 권장합니다.

---

## 1️⃣ 인증 흐름

1. 클라이언트 → Authorization 요청 (response_type=token)
2. 사용자 로그인 & 권한 승인
3. Access Token이 리디렉션 URL에 **해시(#)** 포함되어 전달됨
4. 클라이언트는 이 토큰을 파싱해 API 요청에 사용

---

## 2️⃣ 예시

#### 1. 요청
```
GET /authorize?  
 response_type=token  
 &client_id=CLIENT_ID  
 &redirect_uri=https://yourapp.com/callback  
 &scope=read_profile email  
 &state=xyz123
```

#### 2. 응답
```
https://yourapp.com/callback#access_token=ABC123&token_type=Bearer&expires_in=3600
```
---

## 3️⃣ 보안상 단점

- Access Token이 **URL 해시값으로 노출됨**
- **Refresh Token을 받을 수 없음**
- XSS에 취약 (브라우저에 저장 시 위험)
- **Token 저장을 안전하게 관리할 방법이 부족**

---

## 🎯 정리

| 항목             | 내용 |
|------------------|------|
| 발급 대상         | 브라우저 기반 클라이언트 |
| 발급 토큰         | Access Token (Refresh X) |
| 보안성            | 낮음 (현재는 사용 권장 ❌) |
| 대체 플로우        | Authorization Code Flow + PKCE |
