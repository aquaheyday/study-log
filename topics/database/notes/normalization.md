# 📄 데이터 정규화 (Database Normalization)

---

## 1️⃣ 정규화란?

**정규화(Normalization)** 는 데이터베이스 설계에서 **중복을 최소화하고, 데이터의 일관성과 무결성을 유지하기 위해 테이블을 구조화하는 과정**입니다.

- 중복 제거 (Redundancy ↓)  
- 이상 현상(Anomaly) 방지: 삽입, 수정, 삭제 이상  
- 데이터 무결성(Integrity) 향상

---

## 2️⃣ 정규화의 목적

| 목적 | 설명 |
|------|------|
| **중복 제거** | 동일한 정보 반복 저장 방지 |
| **이상 현상 방지** | 삽입, 삭제, 갱신 시 오류 방지 |
| **데이터 무결성 유지** | 논리적으로 정합성 있는 구조 |
| **관리 용이성 증가** | 유지보수 및 구조 변경이 쉬움 |

---

## 3️⃣ 정규형(Normal Form)의 단계

| 단계 | 조건 | 설명 |
|------|------|------|
| **1NF (제1정규형)** | 원자값(Atomic Value)만 저장 | 컬럼에 반복/다중 값 저장 금지 |
| **2NF (제2정규형)** | 1NF 만족 + 부분 함수 종속 제거 | 기본키의 일부가 아닌 전체에 종속 |
| **3NF (제3정규형)** | 2NF 만족 + 이행적 종속 제거 | 기본키를 거치지 않는 종속 제거 |
| **BCNF (보이스-코드 정규형)** | 결정자 → 후보키 조건 만족 | 모든 결정자가 후보키여야 함 |
| **4NF (제4정규형)** | 다치 종속 제거 | 하나의 속성이 다른 속성에 다치 종속이면 분리 |
| **5NF (제5정규형)** | 조인 종속 제거 | 재구성 가능한 조인 분해 가능 |

---

## 4️⃣ 이상 현상(Anomalies) 예시

| 유형 | 설명 | 예시 |
|------|------|------|
| **삽입 이상** | 일부 데이터만으로는 삽입 불가 | 과목 없이 교수 추가 불가 |
| **삭제 이상** | 데이터 삭제 시 의도치 않은 정보 손실 | 마지막 학생 삭제 → 과목 정보도 사라짐 |
| **갱신 이상** | 중복된 데이터 갱신 누락 가능 | 교수 전화번호 변경 누락 |

---

## 5️⃣ 반정규화(Denormalization)

성능 향상을 위해 **일부러 정규화를 완화**하는 작업  

| 특징 | 설명 |
|------|------|
| 조회 속도 향상 | 조인 감소로 빠른 응답 가능 |
| 쓰기 성능 감소 | 중복된 데이터 갱신 부담 증가 |
| 사용처 | DW, 통계 시스템 등 읽기 위주 환경에서 사용 |

---

## 6️⃣ 정규화 vs 반정규화 비교

| 항목 | 정규화 | 반정규화 |
|------|--------|----------|
| 중복 | 최소화 | 허용 |
| 무결성 | 높음 | 낮음 (관리 필요) |
| 성능 | 쓰기/변경에 유리 | 읽기/조회에 유리 |
| 사용 목적 | 일관성 중시 시스템 | 성능 중시 시스템 |

---

## 🎯 정리

✔ 정규화는 **중복 제거 & 무결성 유지**를 위한 테이블 구조화 작업  
✔ 정규형은 1NF → 5NF로 갈수록 엄격하고 구조화됨  
✔ 이상 현상 방지를 위해 필수  
✔ 읽기 성능이 중요한 경우에는 **반정규화** 고려 가능

