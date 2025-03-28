# 🌐 Web Security - 웹 보안 정리

이 디렉토리는 웹 애플리케이션에서 발생할 수 있는 다양한 보안 위협과 그에 대한 대응 방안을 정리합니다.  
브라우저 기반 위협(XSS, CSRF), 보안 정책(CORS, CSP), HTTPS, 쿠키 보안 설정 등 웹 개발자와 보안 담당자에게 필수적인 내용을 포함합니다.

---

### ☠️ 주요 공격 기법
| 주제 | 파일명 | 설명 |
|------|--------|------|
| XSS (Cross-Site Scripting) | [xss.md](./notes/xss.md) | 악성 스크립트를 삽입해 사용자 브라우저에서 실행 |
| CSRF (Cross-Site Request Forgery) | [csrf.md](./notes/csrf.md) | 사용자의 인증 정보를 이용해 의도치 않은 요청 전송 |
| Clickjacking | [clickjacking.md](./notes/clickjacking.md) | 투명 iframe 등을 이용한 클릭 유도 공격 |
| Open Redirect | [open-redirect.md](./notes/open-redirect.md) | 악의적인 외부 사이트로 강제 리디렉션 유도 |

---

### 🔐 브라우저 보안 정책
| 주제 | 파일명 | 설명 |
|------|--------|------|
| CORS (Cross-Origin Resource Sharing) | [cors.md](./notes/cors.md) | 교차 출처 요청의 허용/제한 원리 |
| CSP (Content Security Policy) | [csp.md](./notes/csp.md) | XSS 방지를 위한 리소스 로딩 제한 정책 |
| SOP (Same-Origin Policy) | [same-origin.md](./notes/same-origin.md) | 도메인/포트/프로토콜 기준 리소스 접근 제한 |
| 브라우저 저장소 보안 | [storage-security.md](./notes/storage-security.md) | localStorage, sessionStorage, 쿠키 보안 이슈 |
| sandbox & iframe 보안 | [sandbox.md](./notes/sandbox.md) | iframe 분리, 속성 설정으로 보안 격리 |
| 보안 헤더 설정 가이드 | [security-headers.md](./snotes/ecurity-headers.md) | X-Frame-Options, HSTS, Referrer-Policy 등 |

---

### 🔒 HTTPS & 인증 보안
| 주제 | 파일명 | 설명 |
|------|--------|------|
| HTTPS & SSL/TLS | [https.md](./notes/https.md) | 암호화된 통신 채널, 인증서 기반 보안 |
| HSTS (Strict Transport Security) | [hsts.md](./notes/hsts.md) | HTTPS 강제 적용 정책 |
| 쿠키 보안 속성 | [cookie-security.md](./notes/cookie-security.md) | HttpOnly, Secure, SameSite 옵션 설명 |

---

### 🛡️ 기타 보안 주제
| 주제 | 파일명 | 설명 |
|------|--------|------|
| 브라우저 보안 메커니즘 | [browser-security.md](./notes/browser-security.md) | SOP, sandbox, iframe 보안 등 |
| 입력 검증과 필터링 | [input-sanitization.md](./notes/input-sanitization.md) | 사용자 입력 유효성 검사 및 정규화 |
| Referrer Policy | [referrer-policy.md](./notes/referrer-policy.md) | 요청 시 참조 헤더 포함 여부 제어 |
