# 🧠 Computer Architecture - 제어 장치 (Control Unit, CU) 란?

**제어 장치 (Control Unit, CU)** 는 CPU 내부에서 **명령어를 해석하고 각 구성 요소를 제어하는 두뇌 역할**을 하는 장치입니다.

> ALU가 계산기라면, CU는 계산기를 조작하는 사람!

---

## 1️⃣ 제어 장치의 정의

- CPU의 핵심 구성 요소 중 하나  
- 메모리에서 명령어를 가져와 **해석하고** 그에 따라 **ALU, 레지스터, 메모리, I/O 장치 등**을 제어  
- **명령어 실행 순서 결정 + 제어 신호 생성**

---

## 2️⃣ 제어 장치의 주요 역할

| 역할                         | 설명 |
|------------------------------|------|
| 명령어 해석                  | 메모리에서 가져온 명령어를 분석 (Op-code 해석) |
| 제어 신호 생성               | ALU, 레지스터, 메모리 등에 제어 신호 전송 |
| 명령어 순서 제어             | 프로그램 카운터(PC)를 조작하여 명령 순서 제어 |
| 데이터 흐름 관리             | 데이터가 어디서 어디로 이동할지 결정 |
| 플래그 판별 및 분기 처리     | 조건 분기, 반복문 등 제어 흐름 판단 |

---

## 3️⃣ 제어 장치의 동작 과정 (명령어 사이클 연계)

1. **Fetch**: 메모리에서 명령어 가져오기 (PC 사용)
2. **Decode**: 명령어 해석 (CU가 담당)
3. **Execute**: 해석 결과에 따라 ALU 등 구성요소 제어
4. **Write Back**: 결과 저장 지시

> 제어 장치는 위 사이클에서 **전체 흐름을 조율**하는 역할

---

## 4️⃣ 제어 장치의 구성 요소 (일반적 개념)

| 구성 요소         | 설명 |
|------------------|------|
| IR (Instruction Register) | 현재 명령어 저장 |
| Decoder          | 명령어 해석기 |
| 제어 신호 생성기  | AND/OR 게이트 등으로 구성 |
| 상태 플래그 확인기 | 조건 분기 판단용 |

---

## 5️⃣ 하드와이어드 vs 마이크로프로그램

제어 신호를 생성하는 방법에는 두 가지 방식이 비교

| 구분             | 하드와이어드 제어 | 마이크로프로그램 제어 |
|------------------|------------------|------------------------|
| 방식             | 논리 회로 기반    | 마이크로명령어 테이블 기반 |
| 속도             | 빠름              | 느림                   |
| 유연성           | 낮음              | 높음                   |
| 사용 예시        | RISC 구조         | CISC 구조              |

> 대부분의 현대 CPU는 두 가지 방식을 조합해서 사용

---

## 6️⃣ 제어 장치 vs ALU 비교

| 항목          | 제어 장치 (CU) | 산술논리장치 (ALU) |
|---------------|----------------|---------------------|
| 기능          | 명령어 해석 & 제어 | 실제 연산 수행 |
| 데이터 처리   | 하지 않음         | 덧셈, 뺄셈 등 연산 |
| 흐름 제어     | 수행             | 제어받음 |

---

## 🎯 정리 요약

✔ **제어 장치(CU)** 는 CPU의 컨트롤 센터, 명령어 해석 및 구성 요소 제어 담당  
✔ ALU, 레지스터, 메모리, I/O를 제어하여 **명령어 실행을 지휘**  
✔ 명령어 사이클의 **Decode 단계**와 **제어 흐름 전체**를 담당  
✔ 하드와이어드 / 마이크로프로그램 방식으로 구현 가능

