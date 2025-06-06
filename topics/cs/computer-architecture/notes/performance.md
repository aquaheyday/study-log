# 📊  Computer Architecture - CPI & MIPS

**CPI (Cycles Per Instruction)** 와 **MIPS (Million Instructions Per Second)** 는  
**컴퓨터의 성능을 정량적으로 평가**할 때 사용하는 대표적인 지표입니다.

> 얼마나 빨리 명령어를 실행하느냐? 를 **CPI와 MIPS로 수치화**할 수 있습니다.

---

## 1️⃣ CPI (Cycles Per Instruction)

- **한 명령어를 실행하는 데 걸리는 평균 클럭 수**  
- 클럭이 빠르더라도 CPI가 높으면 성능이 낮아질 수 있음

#### 공식

```
CPI = 총 클럭 사이클 수 / 총 명령어 수
```

#### 예시
- 총 1억 클럭 사이클이 소모되고, 5천만 개 명령어 실행 → CPI = 2.0  
- 한 명령어당 평균 2사이클이 걸린다는 의미

---

## 2️⃣ MIPS (Million Instructions Per Second)

- **초당 몇 백만 개의 명령어를 실행할 수 있는지**를 나타내는 단위  
- 명령어 수가 많다고 항상 빠른 건 아님 (명령어 복잡도 고려해야 함)

#### 공식

```
MIPS = (클럭 주파수 in MHz) / CPI

또는

MIPS = (명령어 수 / 실행 시간) ÷ 10⁶
```

#### 예시
- 클럭 속도 2GHz, CPI = 2 → MIPS = (2000 / 2) = **1000 MIPS**

---

## 3️⃣ CPI vs MIPS 비교

| 항목       | CPI                                | MIPS                                |
|------------|-------------------------------------|--------------------------------------|
| 의미       | 명령어 1개당 평균 사이클 수           | 초당 실행 가능한 명령어 수 (단위: 백만 개) |
| 수치 의미  | 낮을수록 성능이 좋음                  | 높을수록 성능이 좋음                  |
| 고려 요소  | 명령어 종류, 파이프라인, 구조 등       | 클럭 속도와 CPI                      |
| 단점       | 복잡한 명령어 차이 반영 어려움         | 명령어 복잡도 무시됨                 |

---

## 4️⃣ 실제 성능 측정에서 주의할 점

- MIPS는 **다른 아키텍처 간 비교에 부적절** → 같은 "백만 개" 명령어라도 구조나 복잡도가 다르기 때문  
- CPI는 **명령어별로 다를 수 있어서 평균값만으론 부족** → **CPI × 명령어 수**가 총 사이클 수를 결정

---

## 5️⃣ 함께 쓰이는 공식 정리

| 공식 | 의미 |
|------|------|
| `CPU 시간 = (명령어 수 × CPI) / 클럭 주파수` | 프로그램 실행 시간 |
| `MIPS = 클럭(MHz) / CPI`                    | 초당 명령어 수     |

---

## 🎯 정리 요약

✔ **CPI**: 명령어 1개 실행에 필요한 평균 클럭 수 → **낮을수록 좋음**  
✔ **MIPS**: 초당 실행 가능한 명령어 수 (단위: 백만 개) → **높을수록 좋음**  
✔ **MIPS는 단순 비교용**, **CPI는 구조적 분석용**  
✔ 진짜 성능 평가는 **CPI + 명령어 수 + 클럭 속도**를 함께 고려해야 정확!
