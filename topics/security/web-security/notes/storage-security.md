# 🔐 Browser - 저장소 보안

## 1️⃣ 주요 브라우저 저장소 종류

| 저장소 | 설명 | 만료 시점 | 도메인 접근 제한 | 보안 설정 가능 |
|--------|------|-----------|------------------|----------------|
| `LocalStorage` | 도메인 기반 영구 저장소 | 수동 삭제 전까지 유지 | Same-Origin | ❌ |
| `SessionStorage` | 탭 단위 세션 저장소 | 탭 종료 시 | Same-Origin | ❌ |
| `Cookie` | 서버와 함께 전송되는 작은 데이터 | 설정된 `expires` 또는 `max-age` | 경로/도메인 제한 가능 | ✅ (`Secure`, `HttpOnly`, `SameSite`) |
| `IndexedDB` | 비동기 DB 기반 대용량 저장소 | 수동 삭제 전까지 유지 | Same-Origin | ❌ |

---

## 2️⃣ 보안 이슈 및 주의점

### 1) `LocalStorage` / `SessionStorage`

| 위험 요소 | 설명 |
|-----------|------|
| XSS 공격에 취약 | 자바스크립트로 접근 가능 → 스크립트 삽입 시 탈취 가능 |
| 민감 정보 저장 금지 | 토큰, 패스워드 등 **절대 저장 금지** |
| 브라우저 확장 프로그램 | 악성 확장 프로그램이 저장소 접근 가능 |

✔ `LocalStorage`는 암호화되지 않음 + 브라우저 디버거로 쉽게 볼 수 있음

---

### 2) `Cookie`

| 위험 요소 | 설명 |
|-----------|------|
| 탈취 위험 | XSS로 접근 가능 (`HttpOnly` 미설정 시) |
| 세션 하이재킹 | 쿠키 위조 또는 탈취로 인증 우회 가능 |
| CSRF 공격 | 인증 쿠키가 자동 전송 → 타 사이트 요청에 악용 가능 |

---

## 3️⃣ 안전한 저장소 사용 전략

### 1) `LocalStorage` / `SessionStorage`

- 저장은 **비민감한 데이터만**
- 가능하면 **토큰은 메모리(변수)** 에 저장
- XSS 방어 필수 (입력 검증, CSP, innerHTML 지양 등)
- 저장 시 암호화 고려 (단, JS 암호화는 완전한 보안 아님)

---

### 2) `Cookie`

| 설정 | 설명 |
|------|------|
| `Secure` | HTTPS에서만 전송 가능 |
| `HttpOnly` | JS에서 쿠키 접근 불가 (XSS 방어) |
| `SameSite` | 외부 요청에 대한 쿠키 전송 제한 (`Strict`, `Lax`, `None`) |
| `Path`, `Domain` 제한 | 접근 범위 명확히 지정 |

```http
Set-Cookie: session_id=abc123; Secure; HttpOnly; SameSite=Lax; Path=/; Domain=example.com
```

---

## 4️⃣ 권장 저장 방식 (JWT 등)

| 항목 | 권장 위치 |
|------|-----------|
| Access Token | 메모리 (JS 변수) 저장 → 재시작 시 로그인 필요 |
| Refresh Token | `HttpOnly` 쿠키 (자동 전송, JS 접근 불가) |
| 일반 설정 값 | `LocalStorage` 또는 `IndexedDB` |

---

## 5️⃣ 요약 비교표

| 항목 | LocalStorage | SessionStorage | Cookie |
|------|--------------|----------------|--------|
| 수명 | 영구 | 탭 종료까지 | 지정된 만료 |
| 접근 범위 | Same-Origin | Same-Origin (탭) | 경로/도메인 제한 가능 |
| JS 접근 가능 여부 | 가능 | 가능 | 설정에 따라 차단 가능 |
| 서버 전송 여부 | ❌ | ❌ | ✅ 자동 포함 |
| 보안 옵션 | ❌ | ❌ | ✅ (`Secure`, `HttpOnly`, `SameSite`) |

---

## 🎯 정리

✔ 브라우저 저장소는 **XSS 방어가 전제 조건**  
✔ `LocalStorage`는 편리하지만 보안 취약 → **토큰 저장에 비추천**  
✔ 쿠키는 서버 전송과 보안 옵션 활용 가능 → **세션/인증에 적합**  
✔ 중요한 데이터는 JS 저장소가 아니라 **HttpOnly 쿠키**에 저장 + **암호화 + CSP + 보안 설정 조합**으로 안전 확보
