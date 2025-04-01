# 🔐 Secure Coding - 안전한 인증 구현

**인증(Authentication)** 은 사용자가 누구인지 확인하는 절차이며, 시스템 보안의 핵심 중 하나입니다.  
안전하지 않은 인증 로직은 계정 탈취, 세션 하이재킹, 권한 상승 등 심각한 보안 문제로 이어질 수 있습니다.

---

## 1️⃣ 안전한 인증 구현 원칙

| 원칙                    | 설명 |
|-------------------------|------|
| 비밀번호는 해시로 저장  | 평문 저장 금지, bcrypt/argon2 등 강력한 해시 사용 |
| 로그인 시도 제한        | brute-force 공격 방지 (예: 5회 실패 후 잠금) |
| 인증 실패 응답 통일     | 아이디/비밀번호 구분 없이 동일 메시지 (`로그인 실패`) |
| 이중 인증(MFA) 권장     | SMS, OTP, FIDO 등 추가 인증 수단 |
| 세션 고정 방지          | 로그인 시 세션 ID 재발급 |
| 토큰 검증 철저          | JWT 서명 검증, 만료 체크, 토큰 탈취 대응 |
| HTTPS 필수              | 인증 정보 전송 시 암호화 필수 (중간자 공격 방지) |

---

## 2️⃣ 비밀번호 보안 처리

#### 1. 해시 알고리즘 추천
  - ✅ `bcrypt`, `argon2`, `scrypt`  
  - ❌ SHA-1, MD5 (충돌 위험, 빠름 → brute-force에 취약)

#### 2. Salting 적용
  - 사용자마다 랜덤 salt 추가로 rainbow table 공격 방지

#### bcrypt 예시 (Node.js)
```js
const bcrypt = require('bcrypt');
const hash = await bcrypt.hash(password, 12);
const isMatch = await bcrypt.compare(inputPassword, hash);
```

---

## 3️⃣ 로그인 보안 처리

### 1) 응답 메시지

```text
❌ "비밀번호가 틀렸습니다."
❌ "존재하지 않는 아이디입니다."

✅ "아이디 또는 비밀번호가 일치하지 않습니다."
```

---

### 2) 로그인 시도 제한 (Rate Limiting)

```js
// Express + rate-limit 미들웨어 예시
const rateLimit = require("express-rate-limit");
app.use("/login", rateLimit({
  windowMs: 10 * 60 * 1000, // 10분
  max: 5, // 5회 시도 제한
  message: "로그인 시도 횟수를 초과했습니다."
}));
```

---

## 4️⃣ 토큰 기반 인증 (JWT)

- **서명 검증 필수**
- **만료 시간 설정 (`exp`)**
- **토큰 저장은 Secure/HttpOnly 쿠키 또는 메모리 권장**
- **탈취 대비 리프레시 토큰 관리 정책 필요**

#### JWT 검증 예시 (Node.js)
```js
const decoded = jwt.verify(token, process.env.JWT_SECRET);
```

---

## 5️⃣ 인증 관련 체크리스트

- [ ] 비밀번호 평문 저장 금지, 반드시 해시+salt 적용
- [ ] 로그인 실패 응답은 모호하게 처리
- [ ] 로그인 시도 제한 or CAPTCHA 적용
- [ ] 인증 시 HTTPS 사용 강제
- [ ] MFA(2FA) 도입 고려
- [ ] 세션 고정(Session Fixation) 방지
- [ ] JWT 서명 및 만료 검증 필수
- [ ] 인증 관련 로그 기록 (단, 민감정보 제외)

---

## 🎯 정리

✔ 인증은 시스템의 **입구**, 반드시 철저하게 구현  
✔ 비밀번호는 **해시 + salt**, 평문 절대 금지  
✔ 로그인/토큰 처리 시 **정보 노출 방지 + 반복 공격 방어**  
✔ HTTPS, MFA, 세션 고정 방지 등 **다중 보호 레이어 적용**
