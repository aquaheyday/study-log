# 🛡️ Secure Coding - XSS 대응 (Cross-Site Scripting Prevention)

**XSS (교차 사이트 스크립팅)** 는 공격자가 악성 스크립트를 웹 페이지에 삽입해,  
다른 사용자의 브라우저에서 실행되도록 만드는 공격입니다.

---

## 1️⃣ XSS의 위험성

| 유형       | 설명 |
|------------|------|
| 저장형     | 악성 스크립트가 DB 등에 저장되어 다른 사용자에게 전달됨 |
| 반사형     | URL/파라미터로 전달된 스크립트가 즉시 응답에 포함되어 실행됨 |
| DOM 기반   | 클라이언트 측 JavaScript에서 DOM을 통해 삽입/실행됨 |

#### 공격 예시
```html
<script>alert('XSS')</script>
```

---

## 2️⃣ XSS 대응 전략

| 대응 방법              | 설명 |
|------------------------|------|
| 출력 인코딩            | HTML/속성/JS에 맞게 특수문자 이스케이프 (`<` → `&lt;`) |
| 입력값 검증            | 허용된 값만 통과 (정규식, 길이 제한 등) |
| 템플릿 엔진 사용       | React, Handlebars 등 자동 인코딩 기능 활용 |
| JavaScript 직접 삽입 금지 | `innerHTML`, `document.write` 등 사용 지양 |
| 콘텐츠 보안 정책(CSP) 적용 | 스크립트 출처 제한, inline 스크립트 차단 |

---

## 3️⃣ 출력 인코딩 예시

| 위치            | 대응 방식 |
|-----------------|-----------|
| HTML 본문       | `&lt;`, `&gt;`, `&amp;` 등으로 인코딩 |
| HTML 속성       | `"`, `'` 등 포함 문자 인코딩 |
| JavaScript 내 삽입 | 문자열 리터럴로 escape |
| URL 파라미터    | `encodeURIComponent()` 사용 |

#### HTML 이스케이프 함수 (Node.js)
```js
function escapeHtml(str) {
  return str
    .replace(/&/g, "&amp;")
    .replace(/</g, "&lt;")
    .replace(/>/g, "&gt;")
    .replace(/"/g, "&quot;")
    .replace(/'/g, "&#039;");
}
```

---

## 4️⃣ 템플릿 엔진 안전 예시

#### 1) React
  ```jsx
  <div>{userInput}</div> // 자동 인코딩됨
  ```

#### 2) Handlebars / Mustache
  ```html
  {{userInput}}  <!-- HTML 이스케이프 자동 적용 -->
  ```

#### 3) Thymeleaf (Spring)
  ```html
  <span th:text="${userInput}"></span>
  ```

> ❗주의: `dangerouslySetInnerHTML`, `{{{ }}}`, `th:utext` 등 **비인코딩 출력은 위험**

---

## 5️⃣ CSP (Content Security Policy) 예시

```http
Content-Security-Policy: default-src 'self'; script-src 'self'
```

- **inline script 차단**, 외부 스크립트 출처 제한 등 가능
- 브라우저가 악성 스크립트 로딩/실행 자체를 막음

---

## 6️⃣ XSS 보안 체크리스트

- [ ] 사용자 입력 출력 시 HTML/속성/JS 위치에 맞게 인코딩하고 있는가?
- [ ] 템플릿 엔진 자동 인코딩 기능을 사용하고 있는가?
- [ ] `innerHTML`, `eval`, `document.write()` 등을 지양하고 있는가?
- [ ] 외부 스크립트 삽입을 CSP로 제한하고 있는가?
- [ ] 개발 환경에서 XSS 필터나 웹 방화벽이 적용되어 있는가?

---

## 🎯 정리

✔ XSS는 사용자 입력이 **출력될 때 스크립트로 실행되는 것**  
✔ 입력 검증과 **출력 인코딩**을 함께 적용해야 실질적인 대응 가능  
✔ 템플릿 엔진 자동 인코딩 활용, 위험한 DOM API는 사용 금지  
✔ CSP로 **브라우저 레벨에서 방어**하는 것도 효과적  
