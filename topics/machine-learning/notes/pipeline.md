# 🔁 머신러닝 파이프라인 정리

머신러닝 프로젝트는 단순히 알고리즘을 적용하는 것이 아니라,  
**데이터 수집부터 모델 배포까지의 전체 과정**을 체계적으로 관리하는 것이 중요합니다.  
이를 **머신러닝 파이프라인(Machine Learning Pipeline)** 이라고 합니다.

---

## 🛠 전체 파이프라인 흐름

```plaintext
1. 데이터 수집
2. 데이터 탐색 및 전처리
3. 특성 선택 및 엔지니어링
4. 데이터 분할 (학습/검증/테스트)
5. 모델 선택 및 학습
6. 하이퍼파라미터 튜닝
7. 모델 평가
8. 배포 및 모니터링
```

---

## 1️⃣ 데이터 수집 (Data Collection)

- CSV, 데이터베이스, API, 웹 크롤링 등 다양한 방식으로 수집
- **정량/정성, 정형/비정형 데이터 구분**
- 주의: 결측치, 중복 데이터, 이상치 포함 여부 확인 필요

---

## 2️⃣ 데이터 탐색 및 전처리 (EDA & Preprocessing)

- EDA(탐색적 데이터 분석)를 통해 데이터 분포, 상관관계 파악
- 결측값 처리: 평균/중앙값 대체, 삭제 등  
- 이상치 제거 또는 수정
- 범주형 변수 인코딩 (Label, One-hot)
- 수치형 변수 스케일링 (표준화, 정규화 등)

---

## 3️⃣ 특성 선택 및 엔지니어링 (Feature Engineering)

- 중요하지 않은 변수 제거 (예: 상관계수, 트리 기반 중요도)
- 파생 변수 생성 (예: 날짜 → 연, 월, 요일)
- 차원 축소 기법 적용 (PCA 등)
- 도메인 지식을 활용한 특성 강화

---

## 4️⃣ 데이터 분할 (Train/Validation/Test Split)

- 일반적인 비율: 70%(train) / 15%(validation) / 15%(test)
- 교차 검증 (k-fold cross validation)도 자주 활용
- 랜덤 시드 고정으로 재현 가능성 확보

---

## 5️⃣ 모델 선택 및 학습 (Model Training)

- 문제 유형에 따라 알고리즘 선택
  - 분류: 로지스틱 회귀, SVM, 랜덤 포레스트 등
  - 회귀: 선형 회귀, XGBoost 등
- Scikit-learn, LightGBM, TensorFlow 등 사용

---

## 6️⃣ 하이퍼파라미터 튜닝 (Hyperparameter Tuning)

- 그리드 서치 (GridSearchCV)
- 랜덤 서치 (RandomizedSearchCV)
- 베이지안 최적화, Optuna 등도 가능
- **과적합 방지용 정규화, 조기 종료 등 고려**

---

## 7️⃣ 모델 평가 (Model Evaluation)

- 회귀: MSE, RMSE, R² 등
- 분류: 정확도, 정밀도, 재현율, F1, AUC 등
- Confusion Matrix 및 ROC Curve 시각화
- 학습/검증 오차 비교를 통한 과적합 점검

---

## 8️⃣ 모델 배포 및 모니터링 (Deployment & Monitoring)

- Flask, FastAPI, Docker, AWS Lambda 등을 이용한 서빙
- 모델 버전 관리 (MLflow, DVC 등)
- 예측 결과 모니터링 및 재학습 자동화 고려 (MLOps)

---

## 🎯 파이프라인의 이점

- 전체 과정 자동화로 **재현성** 확보
- 반복 작업 최소화 및 협업 효율 향상
- 지속적인 성능 추적 및 유지보수 가능
