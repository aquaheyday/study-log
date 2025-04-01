# 🛡️ Secure Coding - 안전한 코드 작성

이 디렉토리는 보안 취약점을 방지하기 위한 안전한 코드 작성 원칙과 사례를 정리합니다.  
OWASP Top 10을 포함한 주요 공격 벡터에 대한 대응 전략, 입력 검증, 인증 처리, 보안 헤더 설정 등 실무에서 꼭 알아야 할 내용을 포함합니다.

---

### 🧱 보안 코딩 기본 원칙
| 주제 | 파일명 | 설명 |
|------|--------|------|
| 보안 코딩이란? | [what-is-secure-coding.md](./notes/what-is-secure-coding.md) | 보안 취약점을 방지하는 개발 방식 개요 |
| 입력 검증과 정제 | [input-validation.md](./notes/input-validation.md) | 사용자 입력 유효성 검사, 화이트리스트 방식 |
| 출력 인코딩 | [output-encoding.md](./notes/output-encoding.md) | XSS 방지, HTML/URL/JavaScript 인코딩 전략 |
| 에러 처리 보안 | [error-handling.md](./notes/error-handling.md) | 민감 정보 노출 방지, 에러 메시지 처리 가이드 |

---

### 🔐 인증/세션 관련 안전 처리
| 주제 | 파일명 | 설명 |
|------|--------|------|
| 안전한 인증 구현 | [secure-auth.md](./notes/secure-auth.md) | 비밀번호 저장, 토큰 보관, 인증 흐름 점검 |
| 세션 관리 보안 | [session-security.md](./notes/session-security.md) | 세션 고정 공격 방지, 세션 만료 처리 |
| 패스워드 보안 | [password-security.md](./notes/password-security.md) | Salt, Hash, bcrypt 등 안전한 비밀번호 처리 |

---

### ☠️ 주요 취약점과 대응
| 주제 | 파일명 | 설명 |
|------|--------|------|
| XSS 대응 | [xss.md](./notes/xss.md) | DOM 기반/Reflected/Persistent XSS 방지 |
| SQL Injection 대응 | [sql-injection.md](./notes/sql-injection.md) | Prepared Statement, ORM 사용 전략 |
| Command Injection 대응 | [command-injection.md](./notes/command-injection.md) | 쉘 명령어 실행 시 입력 검증 필요성 |
| Open Redirect 대응 | [open-redirect.md](./notes/open-redirect.md) | 외부 URL 리디렉션 취약점과 방어 방법 |

---

### 🛡️ 보안 설정 및 기타
| 주제 | 파일명 | 설명 |
|------|--------|------|
| 보안 헤더 설정 | [security-headers.md](./notes/security-headers.md) | CSP, X-Frame-Options, HSTS 등 |
| 의존성 관리 & 취약점 점검 | [dependency-security.md](./notes/dependency-security.md) | npm, pip 등 패키지 취약점 대응 전략 |
| 로깅과 민감 정보 | [logging.md](./notes/logging.md) | 로그에 민감 정보 노출 방지하기 |
