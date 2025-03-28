# 🍪 쿠키 보안 속성 (Cookie Security Attributes)

쿠키는 웹에서 상태를 유지하거나 세션을 관리하는 데 사용됩니다.  
하지만 설정을 잘못하면 **보안 취약점**으로 이어질 수 있습니다.  
안전한 쿠키 사용을 위해 다양한 속성을 반드시 설정해야 합니다.

---

## 1️⃣ `Secure` (HTTPS 연결일 때만 쿠키가 전송되도록 제한)
`HTTP` 요청에서는 쿠키 전송 ❌ (중간자 공격 방지)

```
Set-Cookie: sessionId=abc123; Secure
```

✅ `HTTPS` 요청 → sessionId=abc123 쿠키를 함께 전송  
❌ `HTTP` 요청 → sessionId=abc123 쿠키는 전송 안함  

---

## 2️⃣ `HttpOnly`

**JavaScript에서 쿠키 접근을 차단**하여 XSS로부터 보호
```
Set-Cookie: sessionId=abc123; HttpOnly
```
✔ `document.cookie`를 통해 쿠키를 읽을 수 없음  
✔ 세션 쿠키, 인증 쿠키에는 반드시 설정해야 함  

---

## 3️⃣ `SameSite` **(CSRF 방어 핵심 속성)**

쿠키가 **다른 사이트에서 요청될 때 전송되는지** 제어  

```
Set-Cookie: sessionId=abc123; SameSite=Strict
```
| 값 | 설명 |
|-----|------|
| `Strict` | 타 사이트에서 온 요청에는 전송 안함 (가장 안전) |
| `Lax` | 기본값, 안전한 방식(GET, 링크 클릭 등)만 허용 |
| `None` | 어떤 경우든 전송 → 반드시 `Secure`와 함께 써야 함 (`Secure` 없으면 브라우저가 무시)

---

## 4️⃣ `Expires` / `Max-Age`

쿠키의 **만료 시점** 설정
- `Max-Age`는 초 단위
- `Expires`는 정확한 일시 지정
- 만료 시간 지정 안 하면 브라우저가 닫힐 때까지 유지
```
Set-Cookie: token=abc; Max-Age=3600
Set-Cookie: token=abc; Expires=Wed, 10 Apr 2025 12:00:00 GMT
```

✔ 세션 쿠키를 만들고 싶다면 해당 속성을 생략  

---

## 5️⃣ `Domain` & `Path`

쿠키가 **어떤 범위에서 유효한지** 결정

### 1) Domain
- 서브도메인 전체(`*.example.com`)에서 사용 가능
- 생략 시 현재 도메인에서만 유효함
```
Set-Cookie: id=abc; Domain=example.com
```

---

### 2) Path
- 지정된 경로 이하에서만 쿠키가 전송됨
- 기본값은 현재 경로
```
Set-Cookie: id=abc; Path=/admin
```

---

## 6️⃣ `Secure` + `HttpOnly` + `SameSite` 조합 예시
CSRF + XSS + HTTPS 보안까지 모두 강화됨
```
Set-Cookie: session=abc123; Secure; HttpOnly; SameSite=Strict
```
✔ 세션, 로그인 쿠키에는 **이 조합이 거의 필수**

---

## 7️⃣ 보안 쿠키 설정 체크리스트

| 속성 | 적용 이유 |
|------|-----------|
| `Secure` | HTTPS 환경에서만 전송되도록 제한 |
| `HttpOnly` | JS를 통한 쿠키 접근 차단 (XSS 방어) |
| `SameSite` | 외부 요청에서의 쿠키 전송 제어 (CSRF 방어) |
| `Expires` / `Max-Age` | 세션 관리 및 자동 만료 |
| `Domain` / `Path` | 쿠키 전송 범위 제어

---

## 8️⃣ 주의사항

- `SameSite=None`은 **반드시 Secure와 함께** 사용해야 브라우저가 인정
- `HttpOnly`는 JS에서 읽을 수 없지만, **쿠키 자체는 여전히 서버로 전송됨**
- 클라이언트 측 쿠키 저장 시에는 민감 정보 절대 저장하지 말 것 (ex: 비밀번호, 사용자 정보)

---

## 🎯 정리

✔ 쿠키는 보안 속성을 제대로 설정하지 않으면 **XSS, CSRF, 세션 탈취**에 노출될 수 있음  
✔ `Secure + HttpOnly + SameSite` 조합은 보안 쿠키의 기본  
✔ 만료, 도메인, 경로 범위도 상황에 맞게 설정  
✔ 모든 인증/세션 쿠키에는 반드시 보안 속성을 적용해야 안전
