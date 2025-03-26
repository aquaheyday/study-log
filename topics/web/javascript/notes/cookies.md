# 🍪 JavaScript - 쿠키(Cookie)

**쿠키(Cookie)** 는 **작은 데이터를 브라우저에 저장**하고, **서버와 자동으로 주고받을 수 있는 클라이언트-서버 간 데이터 저장 방식**입니다.

---

## 1️⃣ 쿠키란?

- 브라우저에 저장되는 **작은 텍스트 데이터(4KB 이하)**
- **도메인, 경로, 만료일, 보안 옵션 등**과 함께 저장됨
- **HTTP 요청 시 자동 전송**되어 **세션 유지, 사용자 식별** 등에 사용됨

---

## 2️⃣ 쿠키 생성 및 조회

- 쿠키는 **문자열 형태**로 저장되며, 여러 개가 `;`로 구분되어 있음  
- `document.cookie`로는 **읽기/쓰기만 가능**, 삭제는 만료일 조작으로 처리함

```js
// 쿠키 설정
document.cookie = "username=Alice";

// 여러 옵션을 함께 지정
document.cookie = "token=abc123; max-age=3600; path=/";
```

```js
// 쿠키 조회
console.log(document.cookie); 
// 결과: "username=Alice; token=abc123"
```

---

## 3️⃣ 쿠키 옵션

| 옵션          | 설명                                                                 |
|---------------|----------------------------------------------------------------------|
| `expires`     | 쿠키 만료일 (UTC 문자열) 예: `expires=Wed, 01 Jan 2025 00:00:00 GMT` |
| `max-age`     | 쿠키 유효 시간 (초)                                                  |
| `path`        | 쿠키를 적용할 경로 (`/`, `/user` 등)                                |
| `domain`      | 쿠키를 적용할 도메인                                                |
| `secure`      | HTTPS에서만 쿠키 전송                                                |
| `SameSite`    | 크로스사이트 요청 시 쿠키 전송 방식 (`Strict`, `Lax`, `None`)        |

---

## 4️⃣ 쿠키 삭제

```js
// 삭제 = 동일한 이름으로, 과거 만료일 설정
document.cookie = "username=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/";
```

---

## 5️⃣ 사용 예시

### 1) 로그인 세션 유지

- 서버가 로그인 시 `Set-Cookie` 헤더로 쿠키를 내려보냄
- 브라우저는 이후 요청에 자동으로 쿠키 포함

---

### 2) 다크모드 설정 저장

```js
document.cookie = "theme=dark; max-age=86400; path=/";
```

---

## 6️⃣ 쿠키와 보안

| 항목              | 설명                                                            |
|-------------------|-----------------------------------------------------------------|
| `HttpOnly`        | JavaScript에서 접근 불가 (XSS 공격 방지) - 서버에서만 설정 가능 |
| `Secure`          | HTTPS 연결일 때만 전송 |
| `SameSite`        | CSRF 방지 목적 - 크로스사이트 쿠키 전송 제한 |

✔ `HttpOnly`, `Secure`, `SameSite`는 서버에서 `Set-Cookie`로 설정해야 함  
> 예: `Set-Cookie: token=abc; HttpOnly; Secure; SameSite=Strict`

---

## ✅ 쿠키 vs localStorage vs sessionStorage

| 항목              | `Cookies`              | `localStorage`         | `sessionStorage`       |
|-------------------|------------------------|------------------------|-------------------------|
| 용량              | 약 4KB 이하            | 5~10MB                 | 5~10MB                  |
| 만료 설정         | O (`expires`, `max-age`) | X (직접 삭제 필요)     | 자동 (탭 닫힐 때 삭제)  |
| 서버 자동 전송     | ✅ 있음                | ❌ 없음                 | ❌ 없음                 |
| 사용 용도          | 인증, 세션             | 설정 저장, 상태 유지    | 임시 상태 저장          |

---

## 🎯 정리

✔ 쿠키는 서버와 자동으로 데이터를 주고받을 수 있는 저장 수단  
✔ 주로 **인증/세션 유지**에 사용됨  
✔ `document.cookie`로 생성, 조회 가능  
✔ **작고 보안 이슈가 있을 수 있으므로** 중요한 정보는 HttpOnly/Secure 옵션으로 보호
