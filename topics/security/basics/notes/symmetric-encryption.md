# 🔐 대칭키 암호화 (Symmetric Key Encryption)

**대칭키 암호화**는 **하나의 동일한 키**를 사용하여 데이터를 암호화(Encrypt)하고 복호화(Decrypt)하는 암호화 방식

- 암호화 키 = 복호화 키
- 송신자와 수신자가 **같은 키를 사전에 공유**해야 함

---

## 1️⃣ 특징

- **속도가 빠름** → 대량의 데이터 암호화에 적합
- **키 관리가 어려움** → 키를 안전하게 공유하는 것이 관건
- **양방향 암호화** 방식

---

## 2️⃣ 주요 알고리즘

- **AES (Advanced Encryption Standard)**: 현재 가장 널리 사용됨
- **DES (Data Encryption Standard)**: 과거에 사용, 현재는 보안 취약
- **3DES (Triple DES)**: DES를 3번 적용, 보안성 강화
- **Blowfish**, **RC4**, **ChaCha20** 등

---

## 3️⃣ 암호화 과정

1. **송신자**는 공유된 키를 이용해 평문(Plaintext)을 암호문(Ciphertext)으로 변환
2. 암호문을 **수신자에게 전송**
3. **수신자**는 같은 키를 사용하여 암호문을 복호화하여 원래의 평문 복원

---

## 4️⃣ 예시

- 압축 파일(.zip) 비밀번호 설정 (.zip 파일에 비밀번호를 걸면, 내부적으로는 대칭키 암호화 방식이 사용됨)
- 기업 내부 네트워크에서 빠른 데이터 암호화 (기업 내에서 직원끼리 파일을 공유하거나 서버와 클라이언트 간에 데이터를 주고받을 때 대칭키 암호화 사)
- VPN에서 사용하는 일부 암호화 트래픽 (VPN은 사용자와 VPN 서버 간의 통신 내용을 암호화해서 보호)

---

## 5️⃣ 보안상의 단점

- **키 분배 문제**: 안전한 키 전달이 어려움
- **키 개수 증가**: 사용자 수가 많아질수록 키 관리가 복잡해짐
  - 사용자 N명이 있을 때 필요한 키 수 = N(N - 1) / 2
- **기밀성 위협**: 키가 노출되면 암호화 전체가 무력화됨

---

## 6️⃣ 대칭키 vs 비대칭키 비교 (요약)

| 구분             | 대칭키 암호화               | 비대칭키 암호화               |
|------------------|-----------------------------|-------------------------------|
| 키 종류          | 하나의 키                   | 공개키 + 개인키               |
| 속도             | 빠름                        | 느림                          |
| 키 관리          | 어렵다 (키 공유 필요)       | 상대적으로 쉽다 (공개키 사용) |
| 주요 용도        | 데이터 암호화               | 키 교환, 인증, 전자서명 등    |

