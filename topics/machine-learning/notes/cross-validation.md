# 🔄 교차검증 (Cross-Validation)

모델의 일반화 성능을 안정적으로 평가하기 위해 데이터를 여러 번 나누어 학습·검증하는 기법입니다.

---

## 🎯 목적  
- **과적합 방지**: 학습·검증 데이터 분리 불균형으로 인한 오버핏팅 검증  
- **평가 안정성**: 하나의 분할(random split)에 의존하지 않고, 여러 번 평균 성능 확인  
- **하이퍼파라미터 튜닝**: 최적 파라미터 탐색 시 과적합 리스크 낮춤

---

## 🔍 주요 방법

| 방식                          | 설명                                                                 | 장단점                             |
|-------------------------------|----------------------------------------------------------------------|------------------------------------|
| **홀드아웃 (Hold-Out)**       | 데이터셋을 학습/검증(또는 테스트)으로 한 번만 분할                    | 간단·빠름<br>평가 불안정            |
| **K-폴드 (K-Fold)**           | 데이터를 K등분 → K번 순환하며 각 분할을 검증셋으로 사용               | 안정적<br>계산 비용 proportional to K |
| **스트라티파이드 K-폴드**     | 레이블 비율(클래스 분포) 유지하며 K-폴드 분할                         | 불균형 분류에 유리                |
| **LOOCV (Leave-One-Out CV)**  | N개 샘플 중 1개를 검증, 나머지 N–1개로 학습 → N번 반복                | 최대 활용<br>계산 비용 매우 큼     |
| **타임시리즈 CV**             | 시계열 데이터 순서 유지하며 과거→미래 순으로 분할                      | 순서 반영<br>미래 정보 유출 방지    |

---

## 🛠️ 구현 예시 (scikit-learn)

```python
from sklearn.model_selection import KFold, StratifiedKFold, TimeSeriesSplit

# K-Fold
kf = KFold(n_splits=5, shuffle=True, random_state=42)

# Stratified K-Fold (분류)
skf = StratifiedKFold(n_splits=5, shuffle=True, random_state=42)

# Time Series Split
tscv = TimeSeriesSplit(n_splits=5)
