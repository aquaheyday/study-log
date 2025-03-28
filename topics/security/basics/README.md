# 🔐 Security Basics - 보안 기초 개념 정리

보안의 기본이 되는 이론과 필수 개념들을 정리합니다.  
암호화, 해시, 인증, 인가, 공격 기법 이해 등 모든 보안 학습의 기초입니다.

---

### 🧠 보안의 3요소 (CIA Triad)
| 주제 | 파일명 | 설명 |
|------|--------|------|
| 기밀성 (Confidentiality) | [confidentiality.md](./notes/confidentiality.md) | 정보 접근 권한이 없는 사용자로부터 보호 |
| 무결성 (Integrity) | [integrity.md](./notes/integrity.md) | 데이터가 인가되지 않은 방식으로 변경되지 않음 |
| 가용성 (Availability) | [availability.md](./notes/availability.md) | 합법적인 사용자가 정보에 접근 가능해야 함 |

---

### 🔐 암호화와 해싱
| 주제 | 파일명 | 설명 |
|------|--------|------|
| 대칭키 암호화 | [symmetric-encryption.md](./notes/symmetric-encryption.md) | 하나의 키로 암호화/복호화 (AES 등) |
| 비대칭키 암호화 | [asymmetric-encryption.md](./notes/asymmetric-encryption.md) | 공개키/개인키 쌍 (RSA, ECC 등) |
| 해시 함수 | [hash-function.md](./notes/hash-function.md) | 단방향 함수, 무결성 검증 (SHA, MD5 등) |
| 디지털 서명 | [digital-signature.md](./notes/digital-signature.md) | 서명 생성/검증으로 위변조 방지 |

---

### ☠️ 기본적인 공격 기법
| 주제 | 파일명 | 설명 |
|------|--------|------|
| 브루트포스 공격 | [brute-force.md](./notes/brute-force.md) | 무작위 대입 방식의 인증 우회 시도 |
| 사전 공격 | [dictionary-attack.md](./notes/ictionary-attack.md) | 일반적인 패스워드 조합을 사용하는 공격 |
| 리플레이 공격 | [replay-attack.md](./notes/replay-attack.md) | 이전의 유효 요청을 재전송하는 공격 방식 |
| 사회공학 공격 | [social-engineering.md](./notes/social-engineering.md) | 사람의 심리를 이용한 보안 우회 기법 |
