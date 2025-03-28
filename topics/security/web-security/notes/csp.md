# 🛡️ CSP (Content Security Policy)

웹 애플리케이션이 신뢰되지 않은 콘텐츠를 로드하거나 실행하지 않도록 제어하는 보안 정책입니다.  
XSS, 데이터 인젝션 등의 공격을 방지하기 위해 사용되며, 서버가 브라우저에게 “어떤 출처의 리소스를 허용할 것인가”를 선언합니다.

---

## 1️⃣ CSP란?

- Content Security Policy (CSP) 는 웹 브라우저에게 허용된 콘텐츠 출처를 명시해, 악성 콘텐츠 실행을 방지하는 보안 기능입니다.
- XSS, Clickjacking, 데이터 인젝션 등을 방지하는 데 효과적입니다.

---

## 2️⃣ 동작 방식

CSP는 서버가 HTTP 응답 헤더 또는 HTML <meta> 태그를 통해 선언합니다.

#### 1. HTTP 헤더 방식

```
Content-Security-Policy: default-src 'self'; script-src 'self' https://cdn.example.com;
```

#### 2. META 태그 방식

```
<meta http-equiv="Content-Security-Policy" content="default-src 'self'; script-src 'self'">
```

---

## 3️⃣ 주요 디렉티브

| 디렉티브 | 설명 |
|----------|------|
| `default-src` | 모든 리소스의 기본 출처 |
| `script-src` | JavaScript 파일 허용 출처 |
| `style-src` | CSS 파일 허용 출처 |
| `img-src` | 이미지 허용 출처 |
| `font-src` | 웹 폰트 허용 출처 |
| `connect-src` | AJAX, WebSocket 등 네트워크 요청 |
| `media-src` | 오디오, 비디오 등 미디어 리소스 |
| `frame-src` | iframe, frame 등 포함 리소스 |
| `object-src` | 플러그인 객체 (e.g. Flash 등) |
| `base-uri` | <base> 태그 허용 출처 |
| `form-action` | <form> 제출 허용 도메인 |

---

## 4️⃣ 자주 쓰이는 정책 예시

#### 기본 정책 예시
```
Content-Security-Policy:
  default-src 'self';
  script-src 'self' https://cdn.example.com;
  style-src 'self' 'unsafe-inline';
  img-src *;
  connect-src 'none';
```

#### 의미

| 항목 | 의미 |
|------|------|
| `self` | 현재 페이지의 출처만 허용 |
| `unsafe-inline` | 인라인 <style> 허용 (권장하지 않음) |
| `*` | 모든 출처 허용 (가능한 피해야 함) |
| `none` | 아무 것도 허용하지 않음 |

---

## 5️⃣ CSP 우회 방지 및 강화 방법

| 방법 | 설명 |
|------|------|
| `'strict-dynamic'` | 동적 script 삽입 방지 (nonce/hash 기반 정책과 함께 사용) |
| `'nonce-xxxx'` | 무작위로 생성된 nonce 값이 있는 <script>만 허용 |
| `'hash-xxxx'` | 특정 스크립트의 해시가 일치할 때만 허용 |
| `'report-uri'` | 위반 사항을 리포트 받을 URL 설정 |
| `'upgrade-insecure-requests'` | HTTP 요청을 자동으로 HTTPS로 업그레이드 |

### 1) 'nonce-xxxx' 사용 예시
랜덤 난수 값(nonce)을 부여한 스크립트만 실행 허용

#### 1. HTML
```html
<script nonce="abc123">console.log('이 스크립트는 허용됨');</script>
```

#### 2. CSP 헤더
```http
Content-Security-Policy: script-src 'self' 'nonce-abc123';
```

✅ nonce 값이 일치하는 <script>만 실행됨  
❌ inline script 중 nonce 없는 건 모두 차단

---

### 2) 'hash-xxxx' 사용 예시
스크립트 내용의 해시(SHA-256 등)를 미리 계산해 허용

#### 1. HTML
```html
<script>console.log("Hello")</script>
```

#### 2. CSP 헤더
```http
Content-Security-Policy: script-src 'self' 'sha256-oI+3YcM0Yje8AqHskxXHUTUqRMkhdnfr+OBu2nKtoB0=';
```
✅ console.log("Hello") 스크립트는 이 해시와 일치하면 허용  
❌ 내용이 조금이라도 바뀌면 실행 불가

#### 3. 해시 계산은 개발 도구나 openssl 등으로 가능

```bash
echo -n 'console.log("Hello")' | openssl dgst -sha256 -binary | openssl base64
```

---

### 3) 'strict-dynamic' 사용 예시
기본적으로 모든 동적 로딩 <script> 차단, 오직 nonce 또는 hash 스크립트에서 불러오는 것만 허용

#### 1. CSP 헤더
```http
Content-Security-Policy: script-src 'nonce-abc123' 'strict-dynamic';
```

✅ nonce-abc123가 있는 스크립트만 실행됨  
❌ 그 외 외부 CDN, eval, inline 등은 전부 차단됨 (even with trusted src)  

✔ 'strict-dynamic' 사용 시 'self'나 도메인 명시는 무시됨 → 무조건 nonce 또는 hash 기반 신뢰만 허용  

---

### 4) report-uri 또는 report-to 사용 예시
CSP 위반이 발생했을 때, 리포트를 받을 서버 설정

#### 1. CSP 헤더
```http
Content-Security-Policy: default-src 'self'; report-uri /csp-violation-report
```

### 2. 위반 시 브라우저가 보내는 JSON
```json
{
  "csp-report": {
    "document-uri": "https://example.com",
    "violated-directive": "script-src",
    "blocked-uri": "https://evil.com/malware.js"
  }
}
```

✅ 위반 내역을 서버에서 수집해 알림/로그로 활용 가능

---

### 5) upgrade-insecure-requests 사용 예시
HTTP로 된 리소스 요청을 자동으로 HTTPS로 업그레이드

#### 1. CSP 헤더
```http
Content-Security-Policy: upgrade-insecure-requests;
```

#### 2. 예시
`<img src="http://example.com/image.jpg">` → 자동으로 `https://example.com/image.jpg`로 바뀜

✅ 혼합 콘텐츠(Mixed Content) 이슈 방지  
✅ 별도로 URL 바꾸지 않아도 자동 보안 향상  

---

## 6️⃣ CSP 설정 시 주의 사항

- `'unsafe-inline'` 및 `'unsafe-eval'` 은 가능한 피하기
- CDN 사용 시 명확한 출처 도메인 명시
- 동적 스크립트 삽입은 `nonce` 또는 `hash`로 제어
- 리포트용으로 `Content-Security-Policy-Report-Only` 헤더 먼저 활용해 테스트 가능

---

## 7️⃣ CSP + XSS 방어 시너지

| 보안 계층 | 설명 |
|-----------|------|
| 입력 검증 | 유저 입력 필터링 |
| 출력 이스케이프 | HTML, JS, URL 컨텍스트에 맞게 이스케이프 |
| CSP 설정 | 인라인 스크립트 및 외부 리소스 제어 |
| HttpOnly 쿠키 | 세션 쿠키 접근 차단 (JS에서 접근 불가) |
| 프레임워크 보안 기능 | React, Vue, Angular 등 기본 방어 제공 |

---

## 🎯 정리

✔ CSP는 콘텐츠 로딩을 엄격히 제어해 XSS 등 클라이언트 사이드 공격을 방어  
✔ `default-src`, `script-src`, `style-src` 등 컨텍스트별로 출처를 세분화  
✔ `'self'`, `'none'`, `'unsafe-inline'` 등의 키워드 조합에 유의  
✔ CSP는 방어 수단 중 하나일 뿐, 입력 검증 및 출력 이스케이프와 함께 사용해야 안전
