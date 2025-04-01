# 🛡️ Secure Coding - 보안 헤더 설정 (Security Headers)

**보안 헤더(Security Headers)** 는 웹 서버가 응답 시 브라우저에 전달하는 보안 관련 정책입니다.  
적절한 설정을 통해 XSS, Clickjacking, 데이터 유출, MIME 타입 스니핑 등 다양한 공격을 예방할 수 있습니다.

---

## 1️⃣ 주요 보안 헤더 목록

| 헤더 이름                    | 설명 |
|-----------------------------|------|
| `Content-Security-Policy`   | 스크립트, 이미지, 스타일 등의 허용 출처를 정의하여 XSS 방지 |
| `X-Frame-Options`           | 페이지를 iframe으로 삽입할 수 있는지 제어 (Clickjacking 방지) |
| `X-Content-Type-Options`    | MIME 타입을 브라우저가 추측하지 못하게 함 |
| `Strict-Transport-Security` | HTTPS 강제 적용, 중간자 공격 방지 |
| `Referrer-Policy`           | 브라우저가 참조자 정보(Referrer)를 어디까지 보낼지 제어 |
| `Permissions-Policy` (구: Feature-Policy) | 브라우저 API 사용 권한 제어 (예: 카메라, 마이크 등) |
| `Cross-Origin-Opener-Policy` | 사이트 간 리소스 격리 정책 |
| `Cross-Origin-Resource-Policy` | 다른 출처에서 리소스 요청 허용 여부 제어 |

---

## 2️⃣ 헤더 설정 예시

### 1) Apache (`.htaccess` 또는 `httpd.conf`)

```apache
Header always set Content-Security-Policy "default-src 'self'; script-src 'self'; object-src 'none'"
Header always set X-Frame-Options "DENY"
Header always set X-Content-Type-Options "nosniff"
Header always set Strict-Transport-Security "max-age=31536000; includeSubDomains; preload"
Header always set Referrer-Policy "no-referrer-when-downgrade"
Header always set Permissions-Policy "geolocation=(), camera=()"
```

---

### 2) Nginx (`nginx.conf`)

```nginx
add_header Content-Security-Policy "default-src 'self'";
add_header X-Frame-Options "DENY";
add_header X-Content-Type-Options "nosniff";
add_header Strict-Transport-Security "max-age=31536000; includeSubDomains; preload";
add_header Referrer-Policy "no-referrer-when-downgrade";
add_header Permissions-Policy "geolocation=(), camera=()";
```

---

## 3️⃣ 체크리스트

- [ ] 모든 응답에 보안 헤더가 포함되어 있는가?
- [ ] CSP를 통해 외부 스크립트, 스타일 등을 제한하고 있는가?
- [ ] iframe 삽입을 제한하여 Clickjacking을 방지하고 있는가?
- [ ] MIME 타입 스니핑을 차단하고 있는가?
- [ ] HTTPS 강제(HSTS)가 적용되어 있는가?
- [ ] 브라우저 API 권한(카메라, 마이크 등)이 최소화되어 있는가?

---

## 4️⃣ 테스트 도구

| 도구 | 설명 |
|------|------|
| [securityheaders.com](https://securityheaders.com) | 보안 헤더 적용 여부 점검 |
| [Mozilla Observatory](https://observatory.mozilla.org/) | CSP, SSL, 헤더 등 보안 점수 확인 |
| [CSP Evaluator](https://csp-evaluator.withgoogle.com/) | Content-Security-Policy 구성 검증 |

---

## 🎯 정리

✔ 보안 헤더는 **서버 설정만으로 브라우저 수준의 보안 방어를 가능하게 함**  
✔ XSS, Clickjacking, 데이터 유출 등 다양한 공격을 예방할 수 있음  
✔ 기본적으로는 최소한 아래 4가지를 꼭 적용하자:

```text
Content-Security-Policy
X-Frame-Options
X-Content-Type-Options
Strict-Transport-Security
```

✔ 서버 유형에 따라 `.htaccess`, `nginx.conf`, `app.js` 등에서 설정 가능
