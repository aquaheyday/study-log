# 📉 결측치와 이상치 처리 (Missing Values & Outliers in Machine Learning)

데이터 전처리 단계에서 **결측치와 이상치**는 모델 성능과 해석력에 큰 영향을 주는 핵심 요소입니다.

---

## 🕳️ 결측치 (Missing Values)

### 📌 정의  
관측값이 누락된 데이터로, 수집 오류, 사용자 미입력, 시스템 장애 등으로 발생.

### 🔍 탐지 방법
- isnull(), info(), describe() 등
- 시각화: heatmap, missingno 라이브러리 등

### 🛠️ 처리 방법

| 방법 | 설명 | 장단점 |
|------|------|--------|
| **삭제 (Drop)** | 결측치 있는 행/열 제거 | 간단하지만 정보 손실 |
| **대체 (Imputation)** | 평균, 중앙값, 최빈값 등으로 채움 | 편향 가능성 있음 |
| **모델 기반 대체** | KNN, 회귀 등으로 예측 | 정확하지만 복잡 |
| **특정값 대입** | -9999, 'Unknown' 등으로 채움 | 단순하지만 주의 필요 |
| **모델 자동 처리** | XGBoost, LightGBM 등은 결측치 자동 처리 | 구조적으로 유리 |

### 💡 실무 팁
- 수치형 → 평균/중앙값 / 범주형 → 최빈값 사용  
- 중요한 열은 삭제보단 대체가 바람직  
- 결측치 패턴도 예측에 사용 가능 (결측 여부를 Feature로 활용)

---

## 🚨 이상치 (Outliers)

### 📌 정의  
전체 데이터 패턴에서 **멀리 벗어난 비정상적 값**  
예: 키 데이터에서 250cm, 월급 데이터에서 10억 등

### 🔍 탐지 방법

#### 1️⃣ Z-score (표준점수)

    Z = (x - 평균) / 표준편차

- |Z| ≥ 3 이면 이상치로 판단 (정규분포 기준)

#### 2️⃣ IQR (사분위 범위)

    IQR = Q3 - Q1  
    이상치 범위:  
    x < Q1 - 1.5 * IQR  
    x > Q3 + 1.5 * IQR

- 비정규분포에서도 사용 가능

#### 3️⃣ 시각화
- 박스플롯(Boxplot)
- 히스토그램
- 산점도(Scatter)
- QQ-plot

### 🛠️ 처리 방법

| 방법 | 설명 |
|------|------|
| 제거 | 명백한 오류나 이상값은 삭제 |
| 대체 | 평균 또는 경계값으로 대체 |
| 변환 | 로그, 루트 등으로 스케일 축소 |
| 모델 기반 처리 | RobustScaler, IsolationForest 등 사용 |

---

## 🤖 머신러닝에서의 활용

| 항목 | 설명 |
|------|------|
| 스케일링 | 정규분포 가정하에 표준화(Standardization) 사용 |
| 이상치 제거 | 회귀, KNN, SVM 등에서 민감하게 반응 |
| 정규성 검정 | Shapiro-Wilk, Anderson-Darling 등 사용 가능 |
| 로그 변환 | 데이터 왜도(Skewness) 줄여 정규성 확보 |

---

## 🧠 실무 팁 요약

- 결측치는 **삭제보다 대체**, 이상치는 **완전 제거보다 변환 또는 모델 기반 탐지**
- 정규분포 가정 시 Z-score / 비정규분포 시 IQR 사용
- 이상치와 결측치는 종종 **데이터의 중요한 패턴을 반영**하므로 무작정 제거 X
