# 🪙 JWT (JSON Web Token) 개요

**JWT**는 인증 정보를 JSON 형태로 담아 **서명한 토큰**으로, **무상태(stateless)** 인증에 널리 사용되는 방식입니다.


---

## 1️⃣ JWT의 구조

`JWT`는 **3개의 파트**로 구성된 문자열: `Header`.`Payload`.`Signature`


| 파트     | 설명 |
|----------|------|
| Header   | 토큰 타입 & 서명 알고리즘 (예: HS256) |
| Payload  | 사용자 정보(Claims), exp 등 포함 |
| Signature | 비밀 키를 이용한 서명 (위조 방지용) |

#### 예시
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.
eyJ1c2VySWQiOiIxMjMiLCJyb2xlIjoiYWRtaW4ifQ.
zKqZcHqzv0bLDK3YOuZHa9VcD1VmM7n8YI1iy4gGn2w
```


---

## 2️⃣ JWT의 특징

| 항목           | 설명 |
|----------------|------|
| ✅ 자체 정보 포함   | 토큰 안에 사용자 정보 직접 포함 (ex: userId, role) |
| ✅ 무상태 인증     | 서버가 상태(Session)를 저장할 필요 없음 |
| ✅ 서명 기반 검증   | 서버는 서명을 확인해 토큰의 무결성을 보장 |
| ❗ 노출 시 위험     | 탈취되면 위조는 어렵지만, 사용은 가능 → 보안 중요 |

---

## 3️⃣ JWT 사용 흐름

1. 사용자가 로그인 → 서버가 JWT 발급
2. 클라이언트는 JWT를 저장 (`localStorage`, `cookie` 등)
3. 이후 요청 시 JWT를 **Authorization 헤더에 포함**
```
Authorization: Bearer <JWT>
```
4. 서버는 JWT를 검증하고 → 요청 처리

---

## 4️⃣ JWT 주요 클레임 (Payload 내용)

| 필드    | 설명 |
|---------|------|
| `sub`   | 사용자 ID (Subject) |
| `exp`   | 만료 시간 (timestamp) |
| `iat`   | 발급 시간 |
| `iss`   | 발급자 (Issuer) |
| `aud`   | 대상자 (Audience) |
| 커스텀 필드 | userId, role 등 자유롭게 추가 가능 |

---

## 5️⃣ JWT의 장단점

| 장점                        | 단점 |
|-----------------------------|------|
| ✅ 서버 확장에 유리 (무상태)    | ❌ 노출 시 탈취 위험 (HTTPS 필수) |
| ✅ 빠른 인증 처리             | ❌ 토큰 자체가 커질 수 있음 |
| ✅ 서명으로 위변조 방지 가능     | ❌ 토큰 만료 전에는 강제 무효화 어려움 |

---

## 6️⃣ JWT 보안

- 반드시 **HTTPS 사용**
- **짧은 만료 시간 + Refresh Token 분리**
- 필요 시 토큰 블랙리스트 관리
- 민감한 정보는 **Payload에 넣지 말 것** (암호화 안 됨)

---

## ✅ 요약 키워드

✔ JWT → 서명된 인증 토큰  
✔ 사용자 정보 + 만료 시간 포함  
✔ 서버는 토큰만 보고 인증 상태 판별 (Stateless)  
✔ 탈취 시 재사용 가능 → HTTPS + 짧은 수명 + Refresh 전략이 보안에 중요
