# 📐 선형 회귀 (Linear Regression)

회귀 분석의 기초 모델로, 독립 변수와 종속 변수 간의 **선형 관계**를 가정하고, 최적의 회귀 계수를 찾는 과정입니다.

---

## 🔧 모형 수식 유도 (Derivation of OLS)

### 1. 모델 정의  
\[
y_i = \beta_0 + \beta_1 x_{i1} + \dots + \beta_p x_{ip} + \varepsilon_i
\quad\Longleftrightarrow\quad
\mathbf{y} = \mathbf{X}\boldsymbol\beta + \boldsymbol\varepsilon
\]

- \(\mathbf{X}\)는 \(n\times (p+1)\) 설계 행렬(첫 열은 1)  
- \(\boldsymbol\beta=[\beta_0,\dots,\beta_p]^\top\)  
- \(\boldsymbol\varepsilon\)는 오차 벡터  

### 2. 손실 함수 (SSE)  
\[
\mathrm{SSE}(\boldsymbol\beta)
= \sum_{i=1}^n (y_i - \hat y_i)^2
= (\mathbf{y} - \mathbf{X}\boldsymbol\beta)^\top(\mathbf{y} - \mathbf{X}\boldsymbol\beta)
\]

### 3. 미분하여 최적화  
\[
\frac{\partial \mathrm{SSE}}{\partial \boldsymbol\beta}
= -2\mathbf{X}^\top(\mathbf{y} - \mathbf{X}\boldsymbol\beta)
\;\stackrel{!}{=}\; 0
\quad\Longrightarrow\quad
\mathbf{X}^\top\mathbf{X}\,\hat{\boldsymbol\beta} = \mathbf{X}^\top\mathbf{y}
\]

- 이를 **정규방정식 (Normal Equation)** 이라 부르며,  
\[
\boxed{\hat{\boldsymbol\beta} = (\mathbf{X}^\top\mathbf{X})^{-1}\mathbf{X}^\top\mathbf{y}}
\]

---

## 🚫 과적합(Overfitting) 방지 기법

### 1️⃣ 정규화(Regularization)

| 기법         | 손실 함수                                       | 특징                               |
|--------------|-------------------------------------------------|------------------------------------|
| **Ridge (L2)**  | \(\mathrm{SSE} + \alpha \sum_j \beta_j^2\)    | 계수 크기 제약 → 분산 ↓             |
| **Lasso (L1)**  | \(\mathrm{SSE} + \alpha \sum_j |\beta_j|\)    | 희소성 유도 → 불필요 변수 제거      |
| **ElasticNet**  | L1 + L2 복합                                    | Lasso·Ridge 장점 결합              |

> 💡 **Tip**: \(\alpha\) 값은 교차검증으로 최적화

---

### 2️⃣ 교차검증 (Cross-Validation)

- **K-Fold CV**, **Stratified K-Fold** 등을 사용해  
  여러 분할에서 모델 평가 → 일반화 성능 검증  
- **하이퍼파라미터 튜닝**(예: \(\alpha\), degree) 시 과적합 방지

---

### 3️⃣ 특성 선택 & 차원 축소

- **필터·랩퍼·임베디드** 방식으로 불필요 변수 제거  
- **PCA** 등으로 차원 축소 후 회귀 적용 → 노이즈·다중공선성 ↓

---

### 4️⃣ 모델 복잡도 제어

- **다항 회귀(Polynomial Regression)**: 차수 \(d\) 과다 설정 시 과적합 → 교차검증으로 최적 \(d\) 선택  
- **스플라인 회귀(Spline Regression)**, **커널 회귀** 등 부드러운 함수 사용  

---

### 5️⃣ 조기 종료 & 앙상블

- **Gradient Descent** 기반 시 **Early Stopping** 활용  
- **Bagging**(Random Forest)·**Boosting** 대신 선형 모델 사용 시, 트리 계열 앙상블과 비교하여 과적합 정도 확인

---

## 🧠 실무 팁 요약

- **수식 유도**: 정규방정식으로 Closed-form 해 구하고, 대규모 데이터는 경사하강법 활용  
- **과적합 방지**:  
  1. **정규화**(\(\alpha\) 교차검증)  
  2. **교차검증**으로 평가·튜닝  
  3. **특성 선택/차원 축소**로 입력 변수 간소화  
  4. **모델 복잡도 제어**(다항 차수, 스플라인)  
  5. **Early Stopping**, **앙상블 기법** 검토  

