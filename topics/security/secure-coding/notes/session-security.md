# 🛡️ Secure Coding - 세션 관리 보안 (Session Management Security)

**세션(Session)** 은 사용자의 로그인 상태를 유지하기 위해 서버가 생성하는 일시적인 인증 정보입니다.  
보안에 취약한 세션 관리 방식은 세션 탈취(Session Hijacking), 세션 고정(Session Fixation) 등의 공격으로 이어질 수 있습니다.

---

## 1️⃣ 세션 보안의 핵심 원칙

| 항목                  | 설명 |
|-----------------------|------|
| 고유하고 예측 불가능한 세션 ID | UUID, 랜덤 토큰 등 사용, 단순 숫자나 반복 패턴 금지 |
| 세션 고정 방지        | 로그인 시 항상 새 세션 ID 발급 |
| HTTPS 적용            | 세션 ID가 노출되지 않도록 전송 시 암호화 |
| 세션 쿠키 보안 속성   | `HttpOnly`, `Secure`, `SameSite` 속성 필수 |
| 세션 만료 설정        | 일정 시간 비활동 시 자동 로그아웃 |
| 로그아웃 처리 철저    | 세션 삭제, 쿠키 만료, 서버 무효화까지 처리 |

---

## 2️⃣ 쿠키 속성 보안 설정

| 속성        | 설명 |
|-------------|------|
| `HttpOnly`  | 자바스크립트 접근 차단 (XSS 방지) |
| `Secure`    | HTTPS 연결에서만 전송 가능 |
| `SameSite`  | 교차 사이트 요청 차단 (`Strict` 또는 `Lax`) |

```http
Set-Cookie: sessionId=abc123; HttpOnly; Secure; SameSite=Strict
```

---

## 3️⃣ 세션 고정(Session Fixation) 방지

**문제:** 로그인 전 세션 ID가 유지되어, 공격자가 미리 세션을 할당할 수 있음  
**해결:** 로그인 시 새로운 세션 ID를 강제로 발급

#### Express + express-session 예시
```js
req.session.regenerate((err) => {
  if (err) { return next(err); }
  req.session.userId = user.id;
});
```

---

## 4️⃣ 세션 타임아웃 설정

| 유형           | 설명 |
|----------------|------|
| 절대 만료 시간 | 로그인 후 최대 지속 시간 (예: 2시간) |
| 비활동 만료 시간 | 일정 시간 활동 없을 경우 종료 (예: 15분) |

#### express-session 설정 예시
```js
app.use(session({
  secret: "secureSecret",
  resave: false,
  saveUninitialized: false,
  cookie: {
    httpOnly: true,
    secure: true,
    maxAge: 1000 * 60 * 15 // 15분
  }
}));
```

---

## 5️⃣ 세션 관리 보안 체크리스트

- [ ] 세션 ID는 충분히 랜덤하고 추측 불가능한 값인가?
- [ ] 로그인 시 세션을 재생성(regenerate)하는가?
- [ ] 세션 쿠키에 `HttpOnly`, `Secure`, `SameSite` 설정이 있는가?
- [ ] 세션 만료 시간이 적절히 설정되어 있는가?
- [ ] 로그아웃 시 세션이 완전히 종료되는가?
- [ ] HTTPS가 항상 강제되고 있는가?

---

## 🎯 정리

✔ 세션은 인증 이후 상태를 유지하는 핵심 메커니즘  
✔ 세션 ID는 예측 불가능하고, 로그인 시 반드시 갱신  
✔ 쿠키 보안 속성(`HttpOnly`, `Secure`, `SameSite`)은 필수  
✔ 타임아웃과 로그아웃 처리를 통해 세션 수명 관리  
✔ 세션 관리 미흡 = 인증 보안 실패와 직결됨
