# 🗂️ 토큰 저장 위치 정리 (LocalStorage vs Cookie vs SessionStorage)

- **JWT 등의 인증 토큰을 클라이언트에 저장할 때, 어디에 저장할 것인가?**
- 보안성, 사용 편의성, 공격 위험성 등을 고려해 적절한 저장 위치를 선택해야 함

---

## 1️⃣ 저장 방식 개요

| 저장 방식        | 설명 |
|------------------|------|
| `LocalStorage`     | 브라우저에 영구 저장 (탭을 종료해도 유지) |
| `SessionStorage`   | 탭(세션) 단위로 저장 (탭을 닫으면 삭제됨) |
| `Cookie`           | 브라우저에 저장되며, 유효기간까지 유지 가능브라우저가 자동으로 요청에 포함 (`Set-Cookie`로 설정) |

---

## 2️⃣ 주요 비교

| 항목             | LocalStorage             | SessionStorage           | Cookie                            |
|------------------|---------------------------|---------------------------|------------------------------------|
| 저장 위치         | 클라이언트 브라우저       | 클라이언트 브라우저       | 클라이언트 브라우저               |
| 유지 시간         | 브라우저를 꺼도 유지         | 탭을 닫으면 삭제             | 설정한 만료 시간까지 유지          |
| 요청 시 자동 전송 | ❌ 수동으로 헤더에 포함해야 함 | ❌ 수동으로 헤더에 포함해야 함 | ✅ 자동 전송 (같은 도메인 요청 시) |
| JS 접근 가능 여부 | ✅ 가능 (`window.localStorage`) | ✅ 가능 (`window.sessionStorage`) | ❌ HttpOnly 설정 시 접근 불가       |
| XSS 취약성        | ❗ 높음                   | ❗ 높음                   | ✅ HttpOnly 설정 시 안전            |
| CSRF 취약성       | ❌ 낮음 (자동 전송 X)      | ❌ 낮음                   | ❗ 자동 전송 → CSRF 취약 가능성 있음 |
| 보안 설정 옵션     | 없음                      | 없음                      | Secure, HttpOnly, SameSite 가능    |

---

## 3️⃣ 저장 방식 선택 기준

| 상황                                | 추천 저장 방식        |
|-------------------------------------|------------------------|
| SPA(싱글페이지 앱), 모바일 앱         | LocalStorage / SessionStorage |
| 민감한 보안이 중요한 시스템            | **Cookie (HttpOnly + SameSite)** |
| 자동 인증 유지 필요 (ex: 새로고침 유지) | LocalStorage / Cookie |
| CSRF 방지 필요                        | Header 인증 방식 + 토큰 저장 |
| 짧은 세션 유지(탭 단위)               | SessionStorage        |

---

## 4️⃣ 보안 팁

- LocalStorage/SessionStorage는 **XSS에 매우 취약** → 반드시 입력값 필터링, CSP, Content-Type 설정 등 필요함  
- 쿠키 사용 시에는:
  - `HttpOnly` → JS 접근 차단
  - `Secure` → HTTPS에서만 전송
  - `SameSite=Strict|Lax` → CSRF 방지

#### `Set-Cookie` 예시
```
Set-Cookie: access_token=abc123; HttpOnly; Secure; SameSite=Strict
```

---

## 🎯 정리

✔ `LocalStorage` → 유지 길고 간편하지만 XSS 취약  
✔ `Cookie` → 자동 전송 + 보안 옵션 풍부하지만 CSRF 주의  
✔ `SessionStorage` → 탭 단위 보안은 괜찮지만 휘발성  
✔ 상황에 맞게 선택하고, 반드시 **보안 설정**과 함께 사용  

