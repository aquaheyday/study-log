# 🔐 로지스틱 회귀 (Logistic Regression)

이진 분류(Binary Classification) 문제에서 **확률** 기반으로 클래스를 예측하는 선형 모델입니다.  

---

## 🎯 모델 정의 & 확률 해석

### 📌 선형 결합  
먼저 입력 특성 \(\mathbf{x} = (x_1, \dots, x_p)\)와 계수 \(\boldsymbol\beta = (\beta_0, \dots, \beta_p)\)를 이용해  
\[
z = \beta_0 + \sum_{j=1}^p \beta_j x_j = \mathbf{x}^\top \boldsymbol\beta
\]  
를 계산합니다.

### 🔄 시그모이드 함수 (Sigmoid)  
이 \(z\)를 **시그모이드** 함수를 통해 \([0,1]\) 범위의 확률로 변환:  
\[
\sigma(z) = \frac{1}{1 + e^{-z}}
\]  
- \(\sigma(z)\)는 \(P(y=1 \mid \mathbf{x})\)  
- \(1 - \sigma(z)\)는 \(P(y=0 \mid \mathbf{x})\)

---

## 🛠️ 학습 (Parameter Estimation)

### 🔍 최대우도법 (Maximum Likelihood)  
- **로그우도 (Log-Likelihood)** 최댓값을 찾음:  
  \[
  \ell(\boldsymbol\beta)
  = \sum_{i=1}^n \Bigl[\,y_i\log\sigma(z_i) + (1-y_i)\log(1 - \sigma(z_i))\Bigr]
  \]
- 이를 **음의 로그우도(Negative Log-Loss)**를 최소화하는 문제로 변환하여 경사하강법(Gradient Descent) 등으로 최적화.

### 🔄 손실 함수 (Log Loss)  
\[
\text{LogLoss} = -\frac{1}{n}\sum_{i=1}^n \Bigl[y_i\log(\hat p_i) + (1-y_i)\log(1-\hat p_i)\Bigr]
\]

---

## 🚦 분류 임계값 (Decision Threshold)

- 일반적으로 \(\hat p = \sigma(z) \ge 0.5\) 이면 Positive(\(y=1\)), 아니면 Negative(\(y=0\)).  
- 임계값을 조정하여 **Precision–Recall** 또는 **TPR–FPR** 트레이드오프 최적화 가능.

---

## 🚫 과적합 방지 (Regularization)

| 기법         | 목적                       | 손실 함수 예시                                 |
|--------------|----------------------------|-----------------------------------------------|
| **L2 (Ridge)**  | 계수 크기 제약 → 분산 ↓    | LogLoss + \(\alpha \sum_j \beta_j^2\)          |
| **L1 (Lasso)**  | 희소성 유도 → 변수 선택    | LogLoss + \(\alpha \sum_j |\beta_j|\)          |
| **ElasticNet**  | L1+L2 복합                | 두 패널티 결합                                |

> 💡 **Tip**: \(\alpha\)는 교차검증으로 최적화

---

## 🌐 다중 클래스 확장 (Multiclass)

- **One-vs-Rest (OvR)**: 각 클래스 vs 나머지 이진 모델 학습  
- **Softmax 회귀 (Multinomial Logistic)**: 출력층에 Softmax 함수 사용  
  \[
  P(y=k\mid \mathbf{x}) = \frac{\exp(\mathbf{x}^\top \boldsymbol\beta_k)}{\sum_{c}\exp(\mathbf{x}^\top \boldsymbol\beta_c)}
  \]

---

## 💡 실무 팁 요약

- **특성 스케일링**: 계수 수렴 속도·안정성 위해 표준화  
- **이상치**: 로지스틱은 완벽 선형 분리 시 과적합 가능 → 정규화 필수  
- **임계값 조정**: 불균형 데이터에서 ROC/PR 커브 참고  
- **피처 엔지니어링**: 다항 특성, 교호작용 추가 시 비선형 관계 포착  
- **해석력**: \(\exp(\beta_j)\)는 **오즈비(Odds Ratio)**로 해석  

