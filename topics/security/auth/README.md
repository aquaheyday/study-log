# 🔐 Authentication & Authorization

이 디렉토리는 인증(Authentication)과 인가(Authorization)에 대한 개념과 실무 기술을 정리합니다.  
사용자 식별, 권한 부여, 세션 관리, OAuth2, JWT 등 다양한 인증/인가 방식의 원리와 구현 사례를 포함합니다.

---

### 🧾 인증 & 인가 기본 개념
| 주제 | 파일명 | 설명 |
|------|--------|------|
| 인증 vs 인가 | [auth-vs-authz.md](./notes/auth-vs-authz.md) | 사용자 신원 확인과 권한 부여의 차이점 |
| 인증 흐름 | [auth-flow.md](./notes/auth-flow.md) | 로그인 요청부터 인증 성공까지의 단계 |
| 상태 기반 vs 무상태 인증 | [state-vs-stateless.md](./notes/state-vs-stateless.md) | 세션 vs 토큰 방식 비교 |

---

### 🍪 세션 & 쿠키 기반 인증
| 주제 | 파일명 | 설명 |
|------|--------|------|
| 세션 기반 인증 | [session-auth.md](./notes/session-auth.md) | 서버에 사용자 상태를 저장하는 인증 방식 |
| 쿠키 개념과 보안 설정 | [cookie.md](./notes/cookie.md) | Secure, HttpOnly, SameSite 속성 설명 |

---

### 🔐 토큰 기반 인증
| 주제 | 파일명 | 설명 |
|------|--------|------|
| JWT 개요 | [jwt.md](./notes/jwt.md) | JSON Web Token의 구조와 활용법 |
| JWT vs 세션 | [jwt-vs-session.md](./notes/jwt-vs-session.md) | 토큰 기반과 상태 기반 인증 비교 |
| 토큰 저장 위치 | [token-storage.md](./notes/token-storage.md) | 로컬스토리지 vs 쿠키 vs 세션스토리지 보안 비교 |
| 리프레시 토큰 전략 | [refresh-token.md](./notes/refresh-token.md) | 액세스 토큰 갱신 및 보안 고려사항 |

---

### 🌐 OAuth 2.0 & OpenID
| 주제 | 파일명 | 설명 |
|------|--------|------|
| OAuth2 개요 | [oauth2.md](./notes/oauth2.md) | 제3자 인증을 위한 프로토콜 개요 |
| Authorization Code Flow | [auth-code-flow.md](./notes/auth-code-flow.md) | 서버 사이드 앱에서 주로 사용되는 흐름 |
| Implicit Flow |	[implicit.md](./notes/implicit.md) | 주로 SPA에서 사용되던 간편하지만 보안에 취약한 방식 (현재는 비추천) |
| PKCE Flow | [pkce.md](./notes/pkce.md) | 모바일, SPA 환경에서 권장되는 보안 강화 흐름 |
| OpenID Connect | [openid.md](./notes/openid.md) | OAuth 위에 인증을 추가한 표준 프로토콜 |
