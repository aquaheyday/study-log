# ✍️ 입력 검증과 필터링

사용자의 입력은 항상 신뢰할 수 없기 때문에, 웹 애플리케이션에서는 **입력값을 철저히 검증하고, 필요시 필터링 또는 정규화**해야 합니다.

---

## 1️⃣ 입력 검증(Input Validation)이란?

사용자가 입력한 값이 **기대하는 형식, 범위, 패턴에 맞는지 검사**하는 것

- 숫자인가? 이메일 형식인가? 길이 제한은?
- 비정상 입력을 차단하여 **애플리케이션 로직을 보호**
- 프론트엔드에서는 **UX 개선용**으로, 백엔드에서는 **보안 목적**으로 필수

---

## 2️⃣ 필터링(Sanitization)이란?

입력값에서 **불필요하거나 위험한 요소를 제거/변환**하는 작업

- `<script>`, `onerror=`, `javascript:` 등 제거
- HTML 인코딩, 정규식 치환 등을 통해 무해한 값으로 변환
- 주로 **출력 전에 적용**하여 XSS 등 공격 방지

---

## 3️⃣ 클라이언트 vs 서버 입력 검증

| 구분 | 목적 | 설명 |
|------|------|------|
| 클라이언트 (브라우저) | UX 개선 | 실시간 오류 표시, 잘못된 입력 차단 |
| 서버 (백엔드) | 보안 유지 | 악의적인 입력 차단, 데이터베이스 보호 |

✔ 클라이언트 입력 검사는 우회 가능하기 때문에 **클라이언트 검증만으로는 절대 충분하지 않음**  

---

## 4️⃣ 브라우저 입력 검증 방법

### 1) HTML5 기본 속성

`required`, `minlength`, `maxlength`, `pattern`, `type` 등

```html
<input type="email" required pattern=".+@example\.com">
```

---

### 2) JavaScript 유효성 검사

길이, 정규식 검사, 특수문자 필터링 등

```js
const email = document.getElementById("email").value;
if (!/^[^@]+@[^@]+\.[^@]+$/.test(email)) {
  alert("이메일 형식이 잘못되었습니다.");
}
```

---

## 5️⃣ 필터링(Sanitization) 예시

### 1) HTML 이스케이프
```js
function escapeHTML(str) {
  return str
    .replace(/&/g, "&amp;")
    .replace(/</g, "&lt;")
    .replace(/>/g, "&gt;")
    .replace(/"/g, "&quot;")
    .replace(/'/g, "&#039;");
}
```

#### 결과
`<script>alert(1)</script>` → `&lt;script&gt;alert(1)&lt;/script&gt;`

---

### 2) DOM 조작 시 안전한 API 사용

- `element.textContent = userInput` → 안전 ✅  
- `element.innerHTML = userInput` → 위험 ❌

✔ `innerHTML`은 문자열을 HTML로 파싱하여 실행하기 때문에, 악성 코드가 브라우저에서 실행될 수 있음 

---

## 6️⃣ 입력 검증/필터링 실패 시 보안 위험

| 공격 기법 | 설명 |
|-----------|------|
| XSS | 입력된 스크립트가 실행됨 (출력 필터링 실패) |
| SQL Injection | DB 쿼리 조작 (백엔드 필터링 실패) |
| Open Redirect | URL 입력값 미검증 시 악성 사이트 유도 |
| Command Injection | 쉘 명령어 인젝션 등 심각한 위협 |

---

## 7️⃣ 안전한 입력 처리 전략

| 단계 | 설명 |
|------|------|
| ✅ 입력 검증 | 허용할 값만 통과 (화이트리스트 기반) |
| ✅ 길이 제한 | 너무 긴 값은 차단 |
| ✅ 타입 확인 | 숫자, 문자열, 이메일 등 명확하게 검사 |
| ✅ 출력 이스케이프 | HTML, JavaScript, URL 등 컨텍스트에 맞는 인코딩 |
| ✅ DOM API 안전 사용 | innerHTML 대신 textContent, setAttribute 사용

---

## 🎯 정리

✔ 브라우저는 UX 개선을 위한 **입력 검증**을 제공하지만, 보안을 위해선 서버 검증이 필수  
✔ **입력값은 신뢰하지 말고**, 항상 검증 + 필터링 적용  
✔ 클라이언트에서는 `HTML 속성 + JS`, 서버에서는 정규식 + 인코딩으로 다중 방어  
✔ 모든 출력 지점에서는 **컨텍스트에 맞는 이스케이프 처리**가 핵심
