# 📊 ROC 곡선 & AUC (Receiver Operating Characteristic Curve & Area Under the Curve)

분류 모델의 임계값(threshold) 변화에 따른 성능 trade-off를 시각화하고, 전체 성능을 하나의 수치로 요약하는 기법입니다.

---

## 📈 ROC 곡선 (ROC Curve)

### 📌 정의  
- **ROC (Receiver Operating Characteristic) Curve**는 다양한 분류 임계값에서의  
  - **TPR (True Positive Rate, 민감도, Recall)**  
  - **FPR (False Positive Rate, 1 – 특이도)**  
  의 변화를 2차원 평면에 그린 곡선입니다.

\[
\begin{aligned}
\mathrm{TPR} &= \frac{TP}{TP + FN} \\
\mathrm{FPR} &= \frac{FP}{FP + TN}
\end{aligned}
\]

### 🛠️ 작성 방법

| 단계 | 설명 |
|------|------|
| 1. 예측 확률 계산 | 각 샘플에 대해 Positive일 확률 \(\hat p_i\) 산출 |
| 2. 임계값 순회 | threshold를 0부터 1까지 변화시키며 분류 |
| 3. TPR/FPR 계산  | 각 threshold에서 TP, FP, TN, FN 집계 후 TPR, FPR 계산 |
| 4. 곡선 그리기   | x축: FPR, y축: TPR로 점 연결 |

### ✅ 장점 / ❌ 단점

| 장점                                                                  | 단점                                                   |
|-----------------------------------------------------------------------|--------------------------------------------------------|
| 임계값 독립적 → 모델의 전체 분류 능력 평가                              | 불균형 클래스일 때 “거의 모든 예측을 Negative로” 해도 TPR/FPR 왜곡 |
| 시각적으로 다양한 모델 간 비교 및 threshold 결정 지점 선택 가능        | 실제 운영 환경에서는 특정 임계값 성능이 더 중요할 수 있음 |

---

## 📐 AUC (Area Under Curve)

### 📌 정의  
- ROC 곡선 아래 면적(Area)을 구한 값.  
- 0.5 (랜덤) ~ 1.0 (완벽 분류) 사이.

\[
\mathrm{AUC} = \int_{0}^{1} \mathrm{TPR}(f) \, d(\mathrm{FPR}(f))
\]

### ✅ 장점 / ❌ 단점

| 장점                                                            | 단점                                                   |
|-----------------------------------------------------------------|--------------------------------------------------------|
| 단일 수치로 모델 전체 성능 요약                                 | 불균형 데이터에서는 AUC가 높아도 실서비스 성능과 괴리 발생 가능 |
| 모델 간 비교·선택이 쉬움                                       | “어느 지점의 성능” 정보를 제공하지 않음               |

---

## 💡 실무 팁

- **임계값 선택**: ROC 곡선에서 원하는 TPR/FPR trade-off 지점을 찾아 최적 threshold 결정  
- **불균형 데이터**: Precision–Recall Curve와 PR-AUC 병행 검토  
- **시각화**: 여러 모델의 ROC를 한 그래프에 겹쳐 비교하면 성능 차이 파악이 용이  
- **해석 주의**: AUC가 높아도 특정 FPR 구간에서의 TPR 성능이 떨어질 수 있으므로, 비즈니스 목표에 맞춰 부분 영역(AUC@FPR≤0.1 등) 분석 고려  

