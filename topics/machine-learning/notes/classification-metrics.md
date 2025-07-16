# 🧩 분류 모델 평가 지표 (Classification Model Evaluation)

분류 모델의 성능을 평가하기 위해 여러 가지 지표를 사용합니다. 목적과 데이터 특성에 맞는 지표를 선택하세요.

---

## 🗂️ 혼동 행렬 (Confusion Matrix)

### 📌 정의  
실제 레이블과 예측 레이블의 교차 집계표. 이진 분류 기준:

|               | 예측 Positive | 예측 Negative |
|---------------|---------------|---------------|
| 실제 Positive | True Positive (TP)  | False Negative (FN) |
| 실제 Negative | False Positive (FP) | True Negative (TN)  |

### 💡 실무 팁
- 불균형 클래스 문제 있을 때 **절대적 수치** 확인 용이  
- 다음 지표들의 기반이 되는 핵심 테이블

---

## 🎯 정확도 (Accuracy)

### 📌 정의  
전체 샘플 중 맞게 예측한 비율.  
\[
\mathrm{Accuracy} = \frac{TP + TN}{TP + TN + FP + FN}
\]

### ✅ 장점 / ❌ 단점

| 장점                           | 단점                                                       |
|--------------------------------|------------------------------------------------------------|
| 직관적 이해·계산이 쉬움        | 클래스 불균형 시 높은 기만성 (예: 다수 클래스 비율만 반영)  |

### 💡 실무 팁
- **균형 데이터**에 활용  
- 불균형 시에는 다른 지표와 함께 봐야 함

---

## 🔍 정밀도 (Precision) & 재현율 (Recall)

### 📌 정의  
- **Precision (정밀도)**: 예측 Positive 중 실제 Positive 비율  
  \(\mathrm{Precision} = \frac{TP}{TP + FP}\)  
- **Recall (재현율)**: 실제 Positive 중 예측 Positive 비율  
  \(\mathrm{Recall} = \frac{TP}{TP + FN}\)

### ✅ 장점 / ❌ 단점

| 지표        | 장점                                         | 단점                                            |
|-------------|----------------------------------------------|-------------------------------------------------|
| Precision   | FP를 줄여야 할 때(잘못된 Positive 최소화) 좋음 | FN은 고려하지 않아 탐지율 낮출 수 있음         |
| Recall      | FN을 줄여야 할 때(놓치는 Positive 최소화) 좋음 | FP는 고려하지 않아 오탐(False Alarm) 증가시킬 수 있음 |

### 💡 실무 팁
- **스팸 필터**: Recall↑ (놓치는 스팸 최소화)  
- **의료 진단**: Recall 우선, Precision도 함께 고려  

---

## ⚖️ F1-Score

### 📌 정의  
Precision과 Recall의 조화평균(Harmonic Mean).  
\[
\mathrm{F1} = 2 \times \frac{\mathrm{Precision} \times \mathrm{Recall}}{\mathrm{Precision} + \mathrm{Recall}}
\]

### ✅ 장점 / ❌ 단점

| 장점                                       | 단점                                   |
|--------------------------------------------|----------------------------------------|
| Precision·Recall 균형 이룬 단일 지표        | 실제 비즈니스 비용 반영 어렵다         |

### 💡 실무 팁
- Precision과 Recall 중 한쪽으로 편향되면 F1도 낮아짐  
- 두 지표 균형이 중요한 경우에 주로 사용

---

## 📈 ROC-AUC (Receiver Operating Characteristic – Area Under Curve)

### 📌 정의  
- **ROC Curve**: 다양한 임계값에서의 TPR(Recall) vs FPR(=FP/(FP+TN))  
- **AUC**: ROC 곡선 아래 면적, 0.5~1 사이

### ✅ 장점 / ❌ 단점

| 장점                                        | 단점                                              |
|---------------------------------------------|---------------------------------------------------|
| 임계값(threshold)에 독립적                   | 클래스 불균형 심할 때 과대평가 가능                |
| 모델의 분류 경계 성능 전체적으로 평가        | 실제 예측 점수(probability) 필요                  |

### 💡 실무 팁
- 여러 모델 비교 시 유용  
- **Precision-Recall AUC**도 함께 확인하면 불균형 문제 완화 가능

---

## 🧠 로지스틱 손실 (Log Loss)

### 📌 정의  
예측 확률과 실제 레이블 간 로그 손실 평균.  
\[
\mathrm{LogLoss} = -\frac{1}{n} \sum_{i=1}^{n} \Bigl[y_i\log(\hat p_i) + (1-y_i)\log(1-\hat p_i)\Bigr]
\]

### ✅ 장점 / ❌ 단점

| 장점                                 | 단점                            |
|--------------------------------------|---------------------------------|
| 확률 예측 품질 평가 가능             | 해석이 직관적이지 않음          |
| 잘못된 확률 예측에 엄격한 페널티 부여 | 로그 계산으로 수치 불안정 가능  |

### 💡 실무 팁
- **확률 기반 모델** 성능 모니터링에 적합  
- 예측 확률 캘리브레이션 필요 여부 판단  

---

## 🧠 실무 팁 요약

- **불균형 데이터**: Precision, Recall, F1, PR-AUC  
- **균형 데이터**: Accuracy, ROC-AUC  
- **확률 예측 품질**: Log Loss, Calibration Curve  
- 여러 지표를 **종합**해 모델의 강·약점을 파악하고, 비즈니스 목적에 맞춰 최적 지표 선택  
