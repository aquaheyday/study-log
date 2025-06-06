# 🔐 Secure Coding - 패스워드 보안 (Password Security)

**패스워드 보안**은 인증 보안의 가장 기본이자 핵심입니다.  
사용자의 비밀번호는 해킹 시 가장 먼저 노리는 정보이며, 안전하게 저장하고 검증하지 않으면 심각한 보안 사고로 이어질 수 있습니다.

---

## 1️⃣ 패스워드 보안 원칙

| 항목                  | 설명 |
|-----------------------|------|
| 평문 저장 금지        | 절대 평문으로 저장하지 않음 |
| 해시+Salt 적용        | bcrypt, argon2 등 사용 / 사용자마다 다른 salt 적용 |
| 해시 알고리즘 선택    | 느리고 강력한 해시 사용 (bcrypt, argon2id 권장) |
| 비밀번호 최소 요건    | 길이, 문자 조합 정책 (예: 8자 이상, 대소문자 포함 등) |
| 동일 비밀번호 사용 방지 | 이전 비밀번호 재사용 제한, 공통 패턴 차단 |
| 로그인 시도 제한      | brute-force 방어 (5회 실패 시 차단 등) |
| 주기적 변경 권장      | 장기 사용 시 만료 및 재설정 요구 가능 |

---

## 2️⃣ 안전한 해시 저장 방식

### 1) 권장 해시 알고리즘

| 알고리즘  | 특징                        |
|-----------|-----------------------------|
| **bcrypt**   | 가장 널리 사용, 검증된 알고리즘 |
| **argon2id** | 메모리 기반, 최신 권장 표준 |
| scrypt    | 메모리/CPU 병렬 저항 있음    |

> ❌ SHA-1, MD5, SHA-256 단독 사용은 금지 (빠르고 취약함)

---

### 2) 해시 예시 (bcrypt)

```js
const bcrypt = require('bcrypt');

// 해시 생성
const hash = await bcrypt.hash(password, 12);

// 검증
const match = await bcrypt.compare(inputPassword, hash);
```

---

## 3️⃣ 비밀번호 정책 예시

- 최소 8자 이상
- 대소문자, 숫자, 특수문자 포함
- 연속 문자/숫자 제한 (ex. `123456`, `aaaaaa`)
- 사전에 포함된 쉬운 비밀번호 금지 (`password`, `qwerty`, `1234` 등)

```regex
^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?#&])[A-Za-z\d@$!%*?#&]{8,}$
```

---

## 4️⃣ 체크리스트

- [ ] 비밀번호는 절대 평문으로 저장하지 않는다
- [ ] bcrypt/argon2 등 강력한 해시 알고리즘을 사용한다
- [ ] 사용자별 salt를 자동으로 적용한다
- [ ] 비밀번호 입력은 최소 길이 및 복잡도 정책을 적용한다
- [ ] 로그인 시도 횟수를 제한한다
- [ ] 탈취 우려 시 즉시 변경하도록 한다 (알림 포함)
- [ ] 비밀번호 변경 시 이전 비밀번호와의 유사성 체크

---

## 🎯 정리

✔ 비밀번호는 **단순히 암호화가 아니라, 해시 + salt 기반으로 안전하게 저장**해야 함  
✔ bcrypt, argon2id 같이 **검증된 해시 알고리즘** 사용  
✔ **입력 정책, 중복/공통 패턴 차단, 시도 제한, 변경 정책**을 통해 보안 강화  
