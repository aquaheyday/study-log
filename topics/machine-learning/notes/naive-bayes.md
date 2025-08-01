# 📦 나이브 베이즈 (Naive Bayes) — 조건부 확률 기반 분류

나이브 베이즈는 **베이즈 정리**와 **독립성 가정**을 이용해 빠르고 간단하게 분류를 수행하는 확률 모델입니다.

---

## 🔍 모델 정의 & 수식

### 📌 베이즈 정리  
\[
P(y \mid \mathbf{x})
= \frac{P(y)\,P(\mathbf{x}\mid y)}{P(\mathbf{x})}
\]
- \(y\): 클래스  
- \(\mathbf{x}=(x_1, \dots, x_p)\): 특성 벡터  
- **사후확률** \(P(y\mid \mathbf{x})\) 계산 후 가장 큰 클래스로 예측

### 📌 나이브 가정 (조건부 독립성)  
\[
P(\mathbf{x}\mid y)
= \prod_{j=1}^p P(x_j \mid y)
\]
- 특성 간 독립이라 가정 → 계산 단순화  
- 실제로는 약한 성능 저하 있지만, 높은 효율성 장점

### 📌 최종 예측  
\[
\hat y = \arg\max_{y\in\mathcal{Y}} 
\; P(y)\,\prod_{j=1}^p P(x_j \mid y)
\]

---

## 🛠️ 주요 분포별 NB 종류

| 모델                    | 특성 유형             | 우도 \(P(x_j\mid y)\) 예시                          |
|-------------------------|-----------------------|-----------------------------------------------------|
| **Gaussian NB**         | 연속형 실수값         | \(\displaystyle \mathcal{N}(x_j; \mu_{y,j}, \sigma_{y,j}^2)\) |
| **Multinomial NB**      | 이산형 빈도/카운트값  | \(\displaystyle \frac{(\sum_k x_{k|y})!}{\prod_j x_{j|y}!}\prod_j \theta_{y,j}^{x_j}\) |
| **Bernoulli NB**        | 이진형(0/1) 특성      | \(\theta_{y,j}^{x_j}(1-\theta_{y,j})^{1-x_j}\)       |

- \(\theta_{y,j}\): 클래스 \(y\)에서 특성 \(j\)가 1일 확률  
- 빈도 기반 텍스트 분류: Multinomial NB 자주 사용  

---

## ✅ 장점 / ❌ 단점

| 장점                                                    | 단점                                                     |
|---------------------------------------------------------|----------------------------------------------------------|
| 매우 빠르고 메모리·계산 효율성 높음                     | 특성 독립성 가정이 현실에 맞지 않을 수 있어 성능 제한    |
| 소량 데이터에도 안정적                                   | 연속 변수의 경우 정규성 가정 필요 (Gaussian NB)          |
| 출력이 확률이므로 **의사결정 임계값** 조절 용이           | 다수 클래스가 있을 때 사전확률 편향 영향                  |
| 과적합 위험이 낮아 **고차원 데이터**에도 강건            | 특성 간 상호작용 포착 불가                               |

---

## 💡 실무 팁 요약

1. **텍스트 분류**: 단어 빈도·TF-IDF → Multinomial NB  
2. **스케일링 불필요**: 확률 계산에 거리·크기 개념 없음  
3. **라플라스 스무딩**: 희소 데이터·0확률 방지를 위해 \(\alpha>0\) 적용  
4. **이진 특성**: Bernoulli NB로 0/1 피처 다루기 효율적  
5. **하이퍼파라미터 튜닝**: 스무딩 계수 \(\alpha\), 분포 가정(정규 vs 대체) 검토  
