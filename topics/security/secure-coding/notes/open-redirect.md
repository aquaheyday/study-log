# 🛡️ Secure Coding - Open Redirect 대응

**Open Redirect**는 웹 애플리케이션에서 **리다이렉션 URL을 외부 입력으로 받아 처리할 때**,  
공격자가 이를 악용해 **사용자를 악성 사이트로 우회시키는 공격**입니다.

> 사용자는 정상 도메인으로부터 이동한다고 믿고 악성 사이트로 접속할 수 있습니다.

---

## 1️⃣ Open Redirect의 위험성

| 위험 요소              | 설명 |
|------------------------|------|
| 피싱 유도              | 정상 도메인으로 가장한 악성 링크 유도 |
| 세션 탈취 시도         | 로그인 후 리디렉션되는 페이지에 악성 URL 삽입 |
| 브랜드 신뢰도 하락     | 도메인을 통해 악성 행위가 이뤄지는 것으로 인식 |
| 인증 우회 가능성       | OAuth 등 리디렉션 기반 로직에서 우회 발생 가능 |

---

## 2️⃣ 취약한 코드 예시

#### Node.js Express 예시
```js
app.get("/redirect", (req, res) => {
  const target = req.query.url;
  res.redirect(target);
});
```

> `/redirect?url=https://evil.com` → 사용자가 공격자 사이트로 이동됨

---

## 3️⃣ Open Redirect 대응 전략

| 대응 방법               | 설명 |
|------------------------|------|
| **화이트리스트 사용**    | 리디렉션 가능한 URL을 제한된 목록으로 관리 |
| **상대 경로만 허용**     | 외부 도메인 리디렉션 금지, 내부 경로(`/home`, `/dashboard`)만 허용 |
| **정규표현식 검증**      | URL 포맷 검사 및 도메인 필터링 적용 |
| **인코딩된 파라미터 제한** | Base64 등으로 우회된 외부 URL도 검사 필요 |
| **사용자 알림 제공**     | 리디렉션 전 경고 메시지 또는 확인창 제공 가능 |

---

## 4️⃣ 안전한 코드 예시

#### 1. 상대 경로만 허용 (Node.js)

```js
app.get("/redirect", (req, res) => {
  const target = req.query.url;

  if (!target || target.startsWith("http")) {
    return res.status(400).send("잘못된 리디렉션 요청입니다.");
  }

  res.redirect(target);
});
```

#### 2. 화이트리스트 기반 리디렉션

```js
const allowedUrls = ["/home", "/dashboard", "/profile"];

app.get("/redirect", (req, res) => {
  const target = req.query.url;
  if (!allowedUrls.includes(target)) {
    return res.redirect("/home");
  }
  res.redirect(target);
});
```

---

## 5️⃣ Open Redirect 보안 체크리스트

- [ ] 리디렉션 URL을 사용자 입력 그대로 사용하고 있지 않은가?
- [ ] 외부 URL이 아닌, 내부 경로만 허용하고 있는가?
- [ ] URL 입력값에 대해 정규표현식 또는 화이트리스트 검증이 적용되고 있는가?
- [ ] OAuth 또는 로그인 후 리디렉션에서 외부 경로로 이동이 불가능한가?
- [ ] 공격자가 Base64, URL 인코딩 등으로 우회 시도를 하지 못하는가?

---

## 6️⃣ 테스트 입력 예시

| 입력값 (url 파라미터)        | 기대 결과     | 실제 위험                        |
|------------------------------|----------------|-----------------------------------|
| `/dashboard`                 | 정상 이동      | 안전                             |
| `https://evil.com`           | 차단되어야 함  | 피싱, 세션 탈취 가능             |
| `//evil.com`                 | 차단되어야 함  | 브라우저가 외부 도메인으로 인식 |
| `%2f%2fevil.com`             | 차단되어야 함  | URL 인코딩 우회 가능             |

---

## 🎯 정리

✔ Open Redirect는 **정상 사이트를 이용한 악성 리다이렉션 공격**  
✔ **화이트리스트**, **상대경로 제한**, **도메인 검증** 등을 통해 방어  
✔ 입력값은 반드시 검증하고, 사용자 리디렉션은 **신중히 처리**  
✔ 단순한 리디렉션 로직도 **브랜드, 신뢰도, 인증 보안**에 직접적 영향을 미친다
