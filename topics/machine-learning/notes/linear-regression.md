# 📐 선형 회귀 (Linear Regression)

회귀 분석의 기초 모델로, 한 개 이상의 독립 변수(Feature)와 종속 변수(Target) 간의 **선형 관계**를 가정합니다.

---

## 🔍 정의  
- **단순 선형 회귀(Simple Linear Regression)**  
  하나의 독립 변수 \(x\)와 종속 변수 \(y\) 간의 관계를  
  \[
    y = \beta_0 + \beta_1 x + \varepsilon
  \]
  로 모델링  
- **다중 선형 회귀(Multiple Linear Regression)**  
  여러 개의 독립 변수 \(\mathbf{x} = (x_1, x_2, \dots, x_p)\)와  
  \[
    y = \beta_0 + \beta_1 x_1 + \dots + \beta_p x_p + \varepsilon
  \]
  로 표현  

여기서 \(\beta_0\)는 절편(intercept), \(\beta_j\)는 기울기(회귀 계수, coefficient), \(\varepsilon\)은 오차항(error term)입니다.

---

## 🔧 모형 추정 (Parameter Estimation)

### 1. 최소제곱법(Ordinary Least Squares, OLS)  
- **목표**: 실제값 \(y_i\)와 예측값 \(\hat y_i\)의 오차 제곱합(SSE)을 최소화  
  \[
    \min_{\boldsymbol\beta} \sum_{i=1}^n (y_i - \hat y_i)^2
    \quad\Longrightarrow\quad
    \hat{\boldsymbol\beta} = (\mathbf{X}^\top \mathbf{X})^{-1} \mathbf{X}^\top \mathbf{y}
  \]

### 2. 경사하강법(Gradient Descent)  
- 대규모·고차원 데이터에서 OLS 역행렬 계산이 부담스러울 때 사용  
- 반복적으로 \(\beta_j\)를 업데이트하며 SSE 감소  

---

## ✅ 가정 (OLS 가정조건)

1. **선형성(Linearity)**  
   \(\mathbb{E}[y \mid \mathbf{x}]\)가 독립 변수의 선형 조합
2. **독립성(Independence)**  
   오차항 \(\varepsilon_i\)가 서로 독립
3. **등분산성(Homoscedasticity)**  
   모든 \(\mathbf{x}\)에서 오차 분산이 일정 \(\mathrm{Var}(\varepsilon_i)=\sigma^2\)
4. **정규성(Normality)**  
   오차항이 정규분포 \(\varepsilon_i \sim N(0,\sigma^2)\)  
   (주로 신뢰구간·가설검정 시 필요)
5. **다중공선성 없음(No Multicollinearity)**  
   독립 변수들 사이에 강한 선형 상관관계 배제

---

## 📊 평가 지표

| 지표                            | 정의                                                     |
|---------------------------------|----------------------------------------------------------|
| **R² (결정계수)**               | \(\,1 - \frac{\mathrm{SSE}}{\mathrm{SST}}\)  
| **조정 R² (Adjusted R²)**       | 변수 개수 보정된 R²: \(1 - \frac{(1 - R^2)(n-1)}{n-p-1}\) |
| **MSE / RMSE / MAE**            | 평균 제곱 오차 / 제곱근 오차 / 평균 절대 오차             |
| **F-통계량 (F-test)**           | 전체 회귀모형 유의성 검정                                 |
| **회귀 계수 t-검정 (t-test)**    | 각 \(\beta_j\)의 유의성 검정                             |

---

## 💡 실무 팁

- **변수 선택**: 불필요한 변수 제거(필터·임베디드·래퍼 방식)로 다중공선성 완화  
- **표준화/정규화**: 계수 해석·수렴 속도 향상을 위해 독립 변수 스케일 맞추기  
- **잔차 분석**: 잔차 플롯으로 선형성·등분산성·정규성 검토  
- **정규화 회귀**: Ridge/Lasso/ElasticNet으로 과적합 방지  
- **다항 회귀**: 비선형 관계 시 특징 확장(다항식) 후 선형 회귀 적용  

