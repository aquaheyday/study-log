# 📊 Python 머신러닝 개요 (Machine Learning Overview)

머신러닝(Machine Learning, ML)은 **데이터를 학습하여 패턴을 찾아 예측하는 기술**입니다.  
Python에서는 **Scikit-Learn**, **TensorFlow**, **PyTorch** 등의 라이브러리를 사용하여 머신러닝을 구현할 수 있습니다.

---

## 1️⃣ 머신러닝이란?

- **명시적인 프로그래밍 없이** 데이터에서 패턴을 학습하는 기술
- **지도학습(Supervised Learning)**, **비지도학습(Unsupervised Learning)**, **강화학습(Reinforcement Learning)** 으로 구분
- **데이터 기반 의사결정, 추천 시스템, 이미지 인식, 자연어 처리 등 다양한 분야에 활용**

---

## 2️⃣ 머신러닝의 종류

### 1) 지도학습 (Supervised Learning)
- **입력(X)과 정답(Y)이 주어진 데이터**를 학습하여 새로운 데이터 예측
- **주요 알고리즘**:
  - **회귀(Regression)** → 연속된 숫자 예측 (ex. 집값 예측)
  - **분류(Classification)** → 카테고리 분류 (ex. 이메일 스팸 필터링)
  
| 문제 유형 | 알고리즘 예시 | 활용 사례 |
|---------|--------------|----------|
| 회귀 | 선형 회귀(Linear Regression), 랜덤 포레스트 회귀 | 주가 예측, 날씨 예측 |
| 분류 | 로지스틱 회귀(Logistic Regression), 의사결정트리(Decision Tree), 서포트 벡터 머신(SVM), 랜덤 포레스트 | 스팸 분류, 암 진단 |

---

### 2) 비지도학습 (Unsupervised Learning)
- **정답(Y) 없이 데이터에서 패턴을 스스로 찾음**
- **주요 알고리즘**:
  - **군집화(Clustering)** → 유사한 데이터 그룹화 (ex. 고객 세분화)
  - **차원 축소(Dimensionality Reduction)** → 데이터 요약 및 시각화 (ex. PCA)

| 문제 유형 | 알고리즘 예시 | 활용 사례 |
|---------|--------------|----------|
| 군집화 | K-평균(K-Means), 계층적 군집(Hierarchical Clustering) | 고객 세분화, 유전자 분석 |
| 차원 축소 | PCA(주성분 분석), t-SNE | 데이터 시각화, 특징 선택 |

---

### 3) 강화학습 (Reinforcement Learning)
- **보상을 기반으로 최적의 행동을 학습**
- **알고리즘 예시**: Q-Learning, Deep Q-Network(DQN)
- **활용 사례**: 게임 AI, 자율주행, 로봇 제어

---

## 3️⃣ 머신러닝 프로세스

1. **데이터 수집** → CSV, 데이터베이스, API 등에서 데이터 가져오기  
2. **데이터 전처리** → 결측값 처리, 이상치 제거, 데이터 정규화  
3. **특징 엔지니어링** → 의미 있는 데이터 변환 및 선택  
4. **모델 선택** → 적절한 알고리즘 선택 (선형 회귀, 의사결정 트리 등)  
5. **모델 학습(Training)** → 데이터로 학습 진행  
6. **모델 평가(Evaluation)** → 정확도, RMSE, F1-score 등 성능 평가  
7. **모델 배포(Deployment)** → 웹 서비스, 앱, API로 활용  

---

## 4️⃣ Scikit-Learn을 활용한 머신러닝 구현

### 1) Scikit-Learn 설치
```sh
pip install scikit-learn
```

---

### 2) 데이터 불러오기
```python
from sklearn.datasets import load_iris

data = load_iris()
X, y = data.data, data.target  # 특징(X), 정답(y)
print(X.shape, y.shape)  # (150, 4), (150,)
```

---

### 3) 데이터 분할 (훈련/테스트 데이터)
```python
from sklearn.model_selection import train_test_split

X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=42)
print(X_train.shape, X_test.shape)  # (120, 4), (30, 4)
```
✔ `test_size=0.2` → **전체 데이터의 20%를 테스트용으로 분할**  
✔ `random_state=42` → **재현 가능성을 위한 랜덤 시드 설정**  

---

### 4) 머신러닝 모델 학습 (로지스틱 회귀)
```python
from sklearn.linear_model import LogisticRegression

model = LogisticRegression()
model.fit(X_train, y_train)  # 모델 학습
```

---

### 5) 예측 및 평가
```python
from sklearn.metrics import accuracy_score

y_pred = model.predict(X_test)
accuracy = accuracy_score(y_test, y_pred)
print(f"모델 정확도: {accuracy:.2f}")
```
✔ `accuracy_score()` → **모델의 정확도(%) 측정**  

---

## 5️⃣ 주요 머신러닝 알고리즘

### 1) 회귀 (Regression)
| 알고리즘 | 설명 | 활용 사례 |
|---------|------|----------|
| 선형 회귀 (Linear Regression) | 직선 관계 예측 | 집값 예측 |
| 랜덤 포레스트 회귀 (Random Forest) | 여러 결정 트리 평균값 예측 | 날씨 예측 |

---

### 2) 분류 (Classification)
| 알고리즘 | 설명 | 활용 사례 |
|---------|------|----------|
| 로지스틱 회귀 (Logistic Regression) | 이진 분류 | 스팸 필터링 |
| 서포트 벡터 머신 (SVM) | 초평면으로 분류 | 얼굴 인식 |
| 랜덤 포레스트 (Random Forest) | 여러 트리 조합 | 질병 예측 |

---

### 3) 군집화 (Clustering)
| 알고리즘 | 설명 | 활용 사례 |
|---------|------|----------|
| K-평균 (K-Means) | 유사한 그룹 자동 분류 | 고객 분류 |
| 계층적 군집화 | 트리 형태로 데이터 그룹화 | 유전자 분석 |

---

## 6️⃣ 머신러닝 모델 평가 지표

| 평가 지표 | 설명 |
|---------|------|
| 정확도 (Accuracy) | 전체 데이터 중 올바르게 예측한 비율 |
| 정밀도 (Precision) | 모델이 양성으로 예측한 것 중 실제로 양성인 비율 |
| 재현율 (Recall) | 실제 양성 중 모델이 올바르게 예측한 비율 |
| F1-score | 정밀도와 재현율의 조화 평균 |
| RMSE | 회귀 모델의 평균 오차 크기 |

```python
from sklearn.metrics import classification_report

print(classification_report(y_test, y_pred))
```
✔ `classification_report()` → **정확도, 정밀도, 재현율 출력**  

---

## 🎯 정리

✔ **머신러닝** → 데이터에서 패턴을 학습하여 예측하는 기술  
✔ **지도학습** → 정답이 있는 데이터 학습 (회귀, 분류)  
✔ **비지도학습** → 정답 없이 패턴을 탐색 (군집화, 차원 축소)  
✔ **Scikit-Learn** → 머신러닝 모델 구현에 가장 널리 사용됨  
✔ **모델 평가** → 정확도, 정밀도, RMSE 등 다양한 지표 활용  
