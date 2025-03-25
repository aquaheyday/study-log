# 🕳️ Vulnerabilities - 보안 취약점 정리

이 디렉토리는 웹 및 시스템 개발 과정에서 발생할 수 있는 주요 보안 취약점들을 정리합니다.  
취약점의 개념, 발생 원인, 공격 방식, 대응 전략 등을 포함하며, OWASP Top 10 기준을 기반으로 구성되어 있습니다.

---

### 🧨 OWASP Top 10 (2023)
| 주제 | 파일명 | 설명 |
|------|--------|------|
| A01 - Broken Access Control | [broken-access-control.md](./broken-access-control.md) | 권한 없는 리소스에 접근 가능 |
| A02 - Cryptographic Failures | [crypto-failures.md](./crypto-failures.md) | 암호화 관련 설정 미흡, 민감정보 노출 |
| A03 - Injection | [injection.md](./injection.md) | SQL, Command, LDAP 등 인젝션 공격 |
| A04 - Insecure Design | [insecure-design.md](./insecure-design.md) | 보안 고려 없이 설계된 시스템 구조 |
| A05 - Security Misconfiguration | [misconfiguration.md](./misconfiguration.md) | 잘못된 서버 설정, 디버그 모드 노출 등 |
| A06 - Vulnerable & Outdated Components | [outdated-components.md](./outdated-components.md) | 보안 패치되지 않은 라이브러리 사용 |
| A07 - Identification and Authentication Failures | [auth-failures.md](./auth-failures.md) | 인증 절차 우회, 세션 처리 취약 |
| A08 - Software and Data Integrity Failures | [integrity-failures.md](./integrity-failures.md) | 무결성 검증 없는 배포/업데이트 |
| A09 - Security Logging and Monitoring Failures | [logging-failures.md](./logging-failures.md) | 이상 징후를 기록/감지하지 못함 |
| A10 - Server-Side Request Forgery (SSRF) | [ssrf.md](./ssrf.md) | 서버가 악성 URL 요청을 보내도록 유도 |

---

### 💥 기타 주요 취약점
| 주제 | 파일명 | 설명 |
|------|--------|------|
| Cross-Site Scripting (XSS) | [xss.md](./xss.md) | 스크립트를 삽입해 사용자 브라우저에서 실행 |
| Cross-Site Request Forgery (CSRF) | [csrf.md](./csrf.md) | 사용자의 의지와 무관한 요청 전송 |
| Clickjacking | [clickjacking.md](./clickjacking.md) | 보이지 않는 프레임을 통한 클릭 유도 |
| Directory Traversal | [directory-traversal.md](./directory-traversal.md) | ../ 등을 이용한 파일 시스템 탈출 |
| Open Redirect | [open-redirect.md](./open-redirect.md) | 악성 사이트로 리디렉션 유도 |
