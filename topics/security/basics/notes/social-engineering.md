# 🎭 사회공학 공격 (Social Engineering Attack)

**사회공학 공격**은 기술적인 해킹보다, **사람의 심리와 행동을 조작**해서 정보를 탈취하거나 시스템에 접근하는 공격 방식

> 기술보다 **사람을 노리는 공격**

---

## 1️⃣ 특징

- 피해자의 **심리(신뢰, 공포, 호기심)** 를 악용
- 기술적 보안이 아무리 강해도 **사람을 속이면 뚫림**
- 이메일, 전화, 대면, 메시지 등 **모든 커뮤니케이션 수단**이 대상
- 보통 **다른 공격의 전 단계**로 활용됨 (ex: 피싱 → 계정 탈취 → 내부 침투)

---

## 2️⃣ 주요 공격 기법

| 공격 기법          | 설명 |
|--------------------|------|
| **피싱 (Phishing)**         | 가짜 이메일/사이트로 로그인 유도 → 계정 탈취  
| **스미싱 (Smishing)**        | 문자 메시지(SMS)로 악성 링크 전송  
| **비싱 (Vishing)**           | 전화로 사칭(은행, 회사 등) → 개인정보 요구  
| **사칭 (Impersonation)**     | 관리자, 동료, 택배기사 등으로 위장해 접근  
| **프리텍스팅 (Pretexting)** | 거짓 명분(조사 중, 업데이트 요청 등)으로 정보 수집  
| **퀴드 프로 쿼 (Quid Pro Quo)** | 어떤 보상을 미끼로 정보 요구  
| **태일게이팅 (Tailgating)**   | 출입증 없는 사람이 뒤따라 들어가는 물리적 침입

---

## 3️⃣ 실제 사례

- 회사 직원에게 "IT팀입니다. 비밀번호 재설정해드릴게요" → 계정 탈취
- "택배 주소 확인해 주세요" 문자 클릭 → 악성코드 감염
- 가짜 은행 전화로 OTP 요청 → 계좌 이체

> 대부분 **정교하게 설계된 '정상처럼 보이는 상황'** 을 만들어냄

---

## 4️⃣ 방어 방법

| 방법                          | 설명 |
|-------------------------------|------|
| ✅ **보안 교육/훈련**             | 직원들에게 주기적으로 사회공학 기법 교육  
| ✅ **출처 불명 링크/파일 경계**     | 메일, 문자, 웹사이트 등 신뢰되지 않으면 클릭 금지  
| ✅ **개인정보 최소 공유**          | 전화나 메시지로 받는 개인정보 요청은 의심  
| ✅ **이중 인증 (2FA)**            | 비밀번호 탈취 당해도 접근 막기  
| ✅ **사칭 탐지 훈련 (모의 공격)**  | 실제 피싱 메일을 활용한 실습 훈련

---

## 5️⃣ 사회공학 vs 기술 해킹

| 구분           | 사회공학 공격                 | 기술적 해킹                     |
|----------------|-------------------------------|----------------------------------|
| 공격 대상       | 사람                           | 시스템/네트워크                  |
| 접근 방식       | 속임수, 신뢰 조작               | 취약점 분석, 자동화 도구 활용     |
| 방어 방식       | 인식 교육, 절차 강화             | 보안 솔루션, 패치, 접근 제어 등   |
| 성공 확률       | **매우 높음 (사람은 속기 쉬움)** | 시스템 보안 강화되면 낮아짐       |

---

## 🎯 정리

✔ 사람을 속여 정보를 얻는 심리 조작 공격  
✔ 피싱, 스미싱, 비싱, 사칭 등 다양한 형태  
✔ 기술보다 **사람이 취약점**이라는 점을 악용  
✔ **보안 인식 + 절차적 대응**이 가장 중요한 방어  

