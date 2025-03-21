# 📊 Python 딥러닝 기초 (Deep Learning Basics)

딥러닝(Deep Learning)은 **인공신경망(Artificial Neural Network, ANN)** 을 기반으로 한 머신러닝 기법으로,  
대량의 데이터를 학습하여 **이미지 인식, 음성 처리, 자연어 처리(NLP) 등 다양한 분야에서 뛰어난 성능**을 발휘합니다.

---

## 1️⃣ 딥러닝이란?

- **인공신경망(ANN, Artificial Neural Network)** 을 기반으로 한 학습 방법
- 데이터에서 **고차원적인 패턴을 자동으로 학습** 가능
- 대표적인 활용 사례:
  - **이미지 처리** → CNN (Convolutional Neural Network)
  - **자연어 처리(NLP)** → RNN (Recurrent Neural Network), Transformer
  - **강화학습** → AlphaGo, 자율주행

---

## 2️⃣ 머신러닝 vs 딥러닝 차이점

| 구분 | 머신러닝 (ML) | 딥러닝 (DL) |
|------|-------------|-------------|
| 특징 | 특징을 사람이 직접 정의 (Feature Engineering) | 신경망이 자동으로 특징 추출 |
| 모델 | 선형 회귀, 의사결정트리, SVM 등 | CNN, RNN, Transformer 등 |
| 성능 | 데이터가 적어도 잘 작동 | 데이터가 많을수록 성능 향상 |
| 계산량 | 비교적 적음 | 높은 연산량 필요 (GPU 필수) |

✔ 딥러닝은 **대량의 데이터와 강력한 컴퓨팅 자원(GPU)** 이 필요하지만,  
복잡한 패턴을 더 잘 학습할 수 있음.

---

## 3️⃣ 인공신경망(ANN) 기본 개념

### 1) 뉴런(Neuron)
- **입력 → 가중치 적용 → 활성화 함수 → 출력** 과정을 거치는 단위
- 여러 뉴런이 연결되어 **레이어(Layer)를 형성**함.

---

### 2) 신경망 구조
- **입력층(Input Layer)** → 데이터를 받아들이는 층
- **은닉층(Hidden Layer)** → 특징을 학습하는 층 (1개 이상 존재)
- **출력층(Output Layer)** → 최종 결과를 출력하는 층

✔ **신경망이 깊어질수록(레이어 증가) 더 복잡한 패턴을 학습 가능**  

---

## 4️⃣ 딥러닝 활성화 함수 (Activation Function)

| 활성화 함수 | 수식 | 특징 |
|------------|------|------|
| **시그모이드(Sigmoid)** | 𝑓(𝑥) = 1 / (1 + e⁻ˣ) | 값이 0~1 사이, 이진 분류에 사용 |
| **렐루(ReLU)** | 𝑓(𝑥) = max(0, 𝑥) | 음수를 0으로 변경, 학습 속도 빠름 |
| **소프트맥스(Softmax)** | 𝑓(𝑥) = e^𝑥 / ∑e^𝑥 | 다중 클래스 분류에 사용 |

✔ **ReLU** 가 가장 널리 사용됨 (시그모이드는 기울기 소실 문제 있음)  

---

## 5️⃣ 딥러닝 학습 과정

1. **순전파(Forward Propagation)** → 입력 데이터가 신경망을 통과하여 예측값 생성  
2. **손실 계산(Loss Calculation)** → 예측값과 실제값의 차이를 손실 함수로 계산  
3. **역전파(Backpropagation)** → 손실을 줄이기 위해 가중치 업데이트  
4. **최적화(Optimization)** → 경사 하강법(Gradient Descent)으로 모델 학습  

✔ 이 과정을 반복하면서 모델의 정확도가 점점 향상됨.  

---

## 6️⃣ 딥러닝 라이브러리 (TensorFlow & Keras)

#### TensorFlow/Keras 설치
```sh
pip install tensorflow keras
```

✔ TensorFlow는 Google이 개발한 **딥러닝 프레임워크**  
✔ Keras는 TensorFlow 기반의 **사용하기 쉬운 API** 제공  

---

## 7️⃣ 첫 번째 딥러닝 모델 만들기 (Keras)

### 1) 데이터 준비 (MNIST 숫자 손글씨 데이터)
```python
import tensorflow as tf
from tensorflow import keras
import numpy as np

# MNIST 데이터 불러오기
mnist = keras.datasets.mnist
(X_train, y_train), (X_test, y_test) = mnist.load_data()

# 데이터 정규화 (0~1 사이 값으로 변환)
X_train, X_test = X_train / 255.0, X_test / 255.0
```

✔ **MNIST 데이터셋** → 28x28 크기의 숫자 이미지 (0~9)  
✔ **픽셀 값을 0 ~ 1 사이로 정규화** → 학습 속도 향상  

---

### 2) 모델 생성 (신경망 설계)
```python
model = keras.Sequential([
    keras.layers.Flatten(input_shape=(28, 28)),  # 입력층 (Flatten: 1D 변환)
    keras.layers.Dense(128, activation="relu"),  # 은닉층 (128개 뉴런)
    keras.layers.Dense(10, activation="softmax") # 출력층 (10개 클래스 분류)
])
```

✔ `Flatten()` → 2D 이미지를 1D 벡터로 변환  
✔ `Dense(128, activation="relu")` → 은닉층 (128개 뉴런)  
✔ `Dense(10, activation="softmax")` → 출력층 (10개 클래스)  

---

### 3) 모델 컴파일
```python
model.compile(optimizer="adam",
              loss="sparse_categorical_crossentropy",
              metrics=["accuracy"])
```
✔ `adam` → 최적화 알고리즘 (경사 하강법 변형)  
✔ `sparse_categorical_crossentropy` → 다중 분류 손실 함수  
✔ `accuracy` → 정확도를 평가  

---

### 4) 모델 학습 (Training)
```python
model.fit(X_train, y_train, epochs=5)
```
✔ `epochs=5` → 5번 반복 학습  

---

### 5) 모델 평가 (테스트 데이터 성능 확인)
```python
test_loss, test_acc = model.evaluate(X_test, y_test)
print(f"테스트 정확도: {test_acc:.4f}")
```

---

## 8️⃣ 딥러닝 성능 개선 방법

### 1) 데이터 증강 (Augmentation)
```python
from tensorflow.keras.preprocessing.image import ImageDataGenerator

datagen = ImageDataGenerator(rotation_range=10, zoom_range=0.1)
datagen.fit(X_train)
```
✔ 데이터 다양성을 높여 과적합 방지  

---

### 2) 드롭아웃 (Dropout)
```python
keras.layers.Dropout(0.3)  # 30% 뉴런 비활성화
```
✔ 과적합 방지 (훈련 시 랜덤으로 일부 뉴런 제거)  

---

### 3) 하이퍼파라미터 튜닝
- **배치 크기** → `batch_size=32, 64, 128` 테스트
- **학습률(Learning Rate)** → `0.001`, `0.0001` 조정
- **은닉층 개수/뉴런 개수** 변경하여 성능 비교  

---

## 🎯 정리

✔ **딥러닝(Deep Learning)** → 데이터에서 패턴을 학습하는 강력한 AI 기술  
✔ **신경망 구조** → 입력층, 은닉층, 출력층으로 구성  
✔ **활성화 함수** → ReLU, Sigmoid, Softmax  
✔ **딥러닝 학습 과정** → 순전파 → 손실 계산 → 역전파 → 최적화  
✔ **TensorFlow/Keras 사용법** → 딥러닝 모델 구현 가능  
✔ **성능 개선 방법** → 데이터 증강, 드롭아웃, 하이퍼파라미터 튜닝  
