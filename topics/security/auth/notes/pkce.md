# 🔐 Authorization Code Flow with PKCE

**PKCE(Proof Key for Code Exchange)** 는 서버 없는 클라이언트 앱(SPA, 모바일 앱 등)에서도 **보안성 높은 인증**을 가능하게 만든 OAuth2 확장 기능입니다.  
기존 Authorization Code Flow를 **공개키 기반으로 강화**한 방식입니다.

> 주로 **공공 클라이언트(public client)** 즉, 클라이언트 시크릿을 보관할 수 없는 환경(SPA, 모바일)에서 사용됩니다.

---

## 1️⃣ PKCE란?

- 클라이언트가 `code_verifier`를 생성하고, 이를 해싱한 `code_challenge`를 Authorization 요청에 포함
- 토큰 요청 시 `code_verifier`를 함께 전송하여 일치 여부 검증

---

## 2️⃣ 인증 흐름

1. 클라이언트 → Authorization 요청 (code_challenge 포함)
2. 사용자 로그인 & 권한 승인
3. 인가 코드 전달 (Authorization Code)
4. 클라이언트 → 토큰 요청 (code_verifier 포함)
5. 서버에서 검증 후 Access Token 발급

---

## 3️⃣ 예시

### 1) 클라이언트가 먼저 `code_verifier`, `code_challenge` 생성
```
code_verifier:  
abcxyz1234567890abcdefghijk

code_challenge (SHA256 해시 후 Base64):  
oGzURFnFYk-Sz5meB1...
```

---

### 2) 인가 요청

```
GET /authorize?  
 response_type=code  
 &client_id=CLIENT_ID  
 &redirect_uri=https://yourapp.com/callback  
 &code_challenge=oGzURFnFYk-Sz5meB1...  
 &code_challenge_method=S256  
 &scope=read_profile email  
 &state=xyz123
```

---

### 3) 토큰 요청

```
POST /token  
 grant_type=authorization_code  
 &code=AUTH_CODE  
 &redirect_uri=https://yourapp.com/callback  
 &client_id=CLIENT_ID  
 &code_verifier=abcxyz1234567890abcdefghijk
```

---

## 4️⃣ PKCE 보안 효과

| 항목                         | 설명 |
|------------------------------|------|
| code_challenge 검증           | 인가 코드 탈취 방지 (중간 공격 방어) |
| 서버 시크릿 없이 동작         | SPA, 모바일에서 안전하게 사용 가능 |
| Refresh Token 사용 가능       | 장기 세션 유지 가능 |

---

## 🎯 정리

| 항목             | 내용 |
|------------------|------|
| 발급 대상         | SPA, 모바일 앱 등 (서버 없는 앱) |
| 발급 토큰         | Access Token + Refresh Token |
| 보안성            | 높음 (현재 표준 권장 방식) |
| 특징              | code_verifier → code_challenge 검증 |
