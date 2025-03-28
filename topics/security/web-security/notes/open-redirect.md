# 🚨 Open Redirect (오픈 리디렉션)

Open Redirect는 웹 애플리케이션에서 리디렉션 대상 URL을 검증 없이 받아서, **외부 악성 사이트로 리디렉션이 가능해지는 취약점**입니다.  
공격자는 이를 이용해 피싱, 악성 사이트 유도 등에 사용합니다.

---

## 1️⃣ 정의

사용자가 특정 사이트를 클릭했지만, **의도하지 않은 외부 사이트**로 리디렉션되는 취약점

- `redirect=외부주소` 와 같은 파라미터를 이용
- 브라우저 주소창에 **정상적인 도메인이 보이기 때문에** 사용자가 속기 쉬움

---

## 2️⃣ 공격 시나리오

1. 공격자가 다음과 같은 링크를 생성
```
https://trust-site.com/login?redirect=https://evil.com
```
2. 사용자는 `trust-site.com` 도메인이므로 안심하고 클릭
3. 서버는 리다이렉트 처리 → `https://evil.com` 으로 이동
4. 사용자는 피싱 페이지 또는 악성 코드 유포 페이지에 접속됨

---

## 3️⃣ 주요 발생 위치

- 로그인 후 이동 (ex: `redirect=...`)
- 로그아웃, 회원가입 등 이동 페이지
- 공유 링크 또는 이메일 인증 링크

---

## 4️⃣ 공격 활용 예

- 피싱 공격: 사용자가 **정상 사이트로 로그인하는 줄 착각**
- 악성코드 배포: 외부 악성 URL로 우회 유도
- URL 스푸핑: `trust-site.com` 링크처럼 보이지만 실제는 악성 사이트

---

## 5️⃣ 실제 URL 예시
```
https://safe-site.com?next=http://evil.com
https://example.com/redirect?url=https://malicious.site
```
사용자는 클릭만으로 악성 사이트로 리디렉션 됨

---

## 6️⃣ 방어 방법

| 방법 | 설명 |
|------|------|
| ✅ 리디렉션 대상 **화이트리스트**화 | 미리 정의된 URL만 리디렉션 허용 |
| ✅ 상대 경로만 허용 | `/dashboard`, `/home` 처럼 내부 경로만 |
| ✅ 외부 주소 입력 시 경고 또는 중단 | 외부 URL로의 리디렉션은 직접 처리 안 함 |
| ✅ 서버 측 검증 로직 추가 | 도메인 체크, `startsWith('/')` 등 조건 검사

---

### 예시: 서버 측 방어 로직 (Node.js Express)
```js
app.get('/redirect', (req, res) => {
  const url = req.query.url;
  if (url && url.startsWith('/')) {
    res.redirect(url);
  } else {
    res.status(400).send('Invalid redirect target');
  }
});
```
---

## 7️⃣ 취약점 탐지 방법

- 툴 없이 직접 테스트 가능:
    1. `?redirect=http://example.com` 시도
    2. 외부 사이트로 실제 이동되는지 확인
- 자동화 도구 사용 (Burp Suite, ZAP 등)

---

## 🎯 요약

✔ Open Redirect는 **리디렉션 대상 검증 없이 외부 이동이 가능**한 취약점  
✔ 피싱, 악성 링크 유도 등으로 악용 가능  
✔ 반드시 **리디렉션 대상 검증**, **화이트리스트 적용**, **상대 경로만 허용** 등을 통해 방어해야 함  
✔ 단독 공격보다는 **다른 공격과 조합 시 위험도 증가**
