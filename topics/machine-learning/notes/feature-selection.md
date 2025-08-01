# 🔍 특성 선택 및 추출 (Feature Selection & Extraction)

머신러닝 모델의 성능과 해석력을 높이기 위해 **중요하지 않거나 중복된 피처**를 걸러내고, **새로운 요약된 피처**를 만들어내는 과정입니다.

---

## 🗑️ 특성 선택 (Feature Selection)

### 📌 정의  
원본 피처 중에서 모델 성능에 기여도가 낮거나 불필요한 피처를 제거하여 학습 효율과 일반화 성능을 높이는 기법.

### 🔍 주요 방법

| 방법                         | 설명                                                              | 장단점                             |
|------------------------------|-------------------------------------------------------------------|------------------------------------|
| **필터 방식 (Filter)**       | 통계적 지표(분산, 상관계수, 카이제곱 등)로 사전 선택               | 간단·빠름<br>모델 비의존적         |
| **랩퍼 방식 (Wrapper)**      | 피처 서브셋을 만들어 모델 성능을 평가하며 선택 (RFE, Forward/Backward) | 최적화 가능<br>계산 비용 큼        |
| **임베디드 방식 (Embedded)** | 모델 학습 중 가중치·규제(Regularization)로 자동 선택 (L1, Tree 기반) | 계산 효율적<br>모델 의존적         |
| **차원 축소 기반**           | PCA, LDA 등으로 요약된 피처를 생성 후 선택                         | 노이즈 제거 효과<br>해석 어려움     |

### 💡 실무 팁
- **필터**로 1차 제거 후, **랩퍼/임베디드**로 세밀하게 다듬기  
- 상관관계 |ρ| > 0.8 인 피처 중 하나 제거하여 다중공선성 완화  
- `SelectKBest`, `SelectFromModel` 등 scikit-learn 도구 활용  

---

## 📐 특성 추출 (Feature Extraction)

### 📌 정의  
기존 피처들을 조합·변환하여 새로운 저차원 피처를 만들어내는 기법.  
특히, **고차원 데이터**를 요약하여 정보 손실을 최소화합니다.

### 🔑 PCA (Principal Component Analysis)

#### 📌 정의  
공분산 행렬의 고유벡터 분석을 통해 데이터 분산이 최대가 되는 방향(주성분)으로 축을 재설정하는 차원 축소 기법.

#### 🛠️ 처리 방법

| 단계                       | 설명                                                                                  |
|----------------------------|---------------------------------------------------------------------------------------|
| 1. 표준화(Standardization) | 평균 0, 분산 1이 되도록 스케일 조정                                                   |
| 2. 공분산 행렬 계산        | 피처 간 공분산 행렬 Σ 계산                                                             |
| 3. 고유값·고유벡터 분해     | Σ의 고유값(λ)과 고유벡터(v)를 구함                                                     |
| 4. 주성분 선택            | λ가 큰 순서대로 k개 고유벡터(v₁…vₖ)를 선택                                             |
| 5. 변환                     | 원본 데이터를 [v₁…vₖ] 행렬에 투영하여 k차원 저차원 데이터 생성                         |

#### ✅ 장점 / ❌ 단점

| 장점                                                    | 단점                                         |
|---------------------------------------------------------|----------------------------------------------|
| 분산이 큰 정보를 최대한 보존 → 노이즈·중복 제거 효과    | 피처의 물리적 의미 해석 어려움               |
| 고차원 데이터를 저차원으로 압축 → 계산 속도·저장공간 절약 | 비선형 관계 포착 불가                        |
| 잡음(noise)·이상치(outlier)에 상대적으로 견고           | 표준화를 반드시 수행해야 함                  |

#### 💡 실무 팁
- 전체 분산의 **95% 이상**을 설명하는 주성분 개수 선정  
- 스케일 차이가 큰 피처는 반드시 **표준화** 후 적용  
- 시각화를 위해 **2~3차원** 주성분만 사용  
- 비선형 구조가 강하면 Kernel PCA, t-SNE 등 고려  

---

## 🧠 실무 팁 요약

- **필터 → 래퍼/임베디드** 순으로 단계별 피처 선택  
- **상관관계**, **분산 임계치** 등 간단한 지표로 1차 제거  
- PCA로 차원 축소 시 **표준화** 필수, 분산 설명 비율 확인  
- 비선형 특성이 많으면 PCA 외 **커널 기법** 검토  

