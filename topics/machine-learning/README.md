# 🤖 Machine Learning

이 공간은 **머신러닝의 핵심 이론, 알고리즘 구현, 실습 예제, 프로젝트**를 정리한 학습용 자료 저장소입니다.  

> 📌 머신러닝의 원리부터 실전 적용까지 체계적으로 학습하는 데 목적이 있습니다.

---

## 📋 머신러닝 이론 및 기본기 정리

### 📌 머신러닝 기초 개념
| 주제 | 파일명 | 설명 |
|---|---|---|
| 머신러닝이란? | [ml-intro.md](./notes/ml-intro.md) | 머신러닝의 정의, 역사, 분류 |
| 지도/비지도/강화 학습 | [learning-types.md](./notes/learning-types.md) | 학습 방식별 차이점과 활용 예 |
| 머신러닝 vs 딥러닝 | [ml-vs-dl.md](./notes/ml-vs-dl.md) | 개념 차이와 연계성 |
| 머신러닝 파이프라인 | [pipeline.md](./notes/pipeline.md) | 데이터 준비 → 모델링 → 평가 과정 |

---

### 🧮 수학 기초 & 통계
| 주제 | 파일명 | 설명 |
|---|---|---|
| 선형대수 | [linear-algebra.md](./notes/linear-algebra.md) | 벡터, 행렬, 내적/외적 개념 |
| 미분 & 경사하강법 | [calculus.md](./notes/calculus.md) | 비용 함수, 경사하강법 적용 |
| 확률 및 통계 | [statistics.md](./notes/statistics.md) | 확률 분포, 평균/분산, 조건부 확률 |
| 정규분포와 이상치 | [normal-distribution.md](./notes/normal-distribution.md) | 데이터 분포 이해 및 이상치 처리 |

---

### 🧹 데이터 전처리
| 주제 | 파일명 | 설명 |
|---|---|---|
| 결측치와 이상치 처리 | [missing-outlier.md](./notes/missing-outlier.md) | NaN, 이상치 탐지 및 처리 방법 |
| 특성 스케일링 | [scaling.md](./notes/scaling.md) | 표준화, 정규화, 로그 변환 등 |
| 범주형 인코딩 | [encoding.md](./notes/encoding.md) | 원-핫 인코딩, 레이블 인코딩 |
| 특성 선택 및 추출 | [feature-selection.md](./notes/feature-selection.md) | 불필요한 변수 제거, PCA 소개 |

---

### 📊 모델 평가 지표
| 주제 | 파일명 | 설명 |
|---|---|---|
| 회귀 모델 평가 | [regression-metrics.md](./notes/regression-metrics.md) | MSE, RMSE, MAE |
| 분류 모델 평가 | [classification-metrics.md](./notes/classification-metrics.md) | 정확도, 정밀도, 재현율, F1 |
| ROC 곡선 & AUC | [roc-auc.md](./notes/roc-auc.md) | 이진 분류 성능 시각화 |
| 교차검증 | [cross-validation.md](./notes/cross-validation.md) | k-fold, Stratified K-Fold |

---

## 🔍 주요 알고리즘 이론

| 알고리즘 | 파일명 | 설명 |
|---|---|---|
| 선형 회귀 | [linear-regression.md](./notes/linear-regression.md) | 수식 유도, 과적합 방지 기법 포함 |
| 로지스틱 회귀 | [logistic-regression.md](./notes/logistic-regression.md) | 시그모이드 함수와 확률 기반 분류 |
| 결정 트리 | [decision-tree.md](./notes/decision-tree.md) | 분할 기준, 정보이득, 가지치기 |
| KNN | [knn.md](./notes/knn.md) | 거리 기반 예측, K 선택법 |
| SVM | [svm.md](./notes/svm.md) | 초평면 개념, 커널 트릭 소개 |
| K-Means | [kmeans.md](./notes/kmeans.md) | 군집화 원리 및 시각화 |
| 나이브 베이즈 | [naive-bayes.md](./notes/naive-bayes.md) | 조건부 확률 기반 분류 |
| 랜덤 포레스트 | [random-forest.md](./notes/random-forest.md) | 앙상블 기법, 부트스트래핑 |

---

## 📚 참고 자료

- [Scikit-learn 공식 문서](https://scikit-learn.org/stable/)
- [Hands-On ML with Scikit-Learn, Keras, and TensorFlow (O'Reilly)](https://www.oreilly.com/library/view/hands-on-machine-learning/)
- [Kaggle](https://www.kaggle.com/) - 데이터 분석 및 대회 플랫폼
- [머신러닝 공부 로드맵 (Github)](https://github.com/rasbt/mlxtend) - 구현 중심 라이브러리
