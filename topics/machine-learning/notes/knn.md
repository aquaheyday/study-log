# 📐 K-최근접 이웃 (K-Nearest Neighbors, KNN) — 거리 기반 예측 & K 선택법

KNN은 학습 단계에서 모델 파라미터 학습 없이, **거리 계산**만으로 가장 가까운 이웃 \(K\)명의 레이블(또는 값을) 참조해 예측하는 **비모수(non-parametric)** 방식입니다.

---

## 🗺️ 거리 기반 예측 (Distance-Based Prediction)

### 📌 정의  
입력 샘플 \(\mathbf{x}_*\)과 학습 샘플 \(\{\mathbf{x}_i,y_i\}\) 간 **거리**를 계산하여, 가장 가까운 \(K\)개의 이웃을 찾아 예측.

### 🔢 주요 거리 지표

| 거리 유형           | 수식                                               | 특징                           |
|---------------------|----------------------------------------------------|--------------------------------|
| **유클리드 거리**   | \(\displaystyle d(\mathbf{x},\mathbf{x}') = \sqrt{\sum_j (x_j - x'_j)^2}\) | 직관적, 연속 변수에 적합        |
| **맨해튼 거리**     | \(\displaystyle d(\mathbf{x},\mathbf{x}') = \sum_j |x_j - x'_j|\)          | 고차원에서 과민 반응 ↓           |
| **민코프스키 거리** | \(\displaystyle \bigl(\sum_j |x_j - x'_j|^p\bigr)^{1/p}\)              | \(p\)-값 변경으로 유연성 제공   |
| **코사인 유사도**   | \(\displaystyle 1 - \frac{\mathbf{x}\cdot\mathbf{x}'}{\|\mathbf{x}\|\|\mathbf{x}'\|}\) | 방향성 기준, 텍스트·문서 비교에 유리 |

### 🛠️ 예측 방식

| 예측 유형    | 방법                                                                             |
|--------------|----------------------------------------------------------------------------------|
| **분류**     | 다수결 투표(Majority Vote)  
:  \(\hat y = \mathrm{mode}\{y_{(1)},...,y_{(K)}\}\)                               |
| **회귀**     | 이웃 값의 평균 또는 가중평균  
:  \(\hat y = \frac{1}{K}\sum_{i=1}^K y_{(i)}\)  
  또는 거리 역수 가중  
  \(\hat y = \frac{\sum_{i=1}^K w_i\,y_{(i)}}{\sum_{i=1}^K w_i},\; w_i = 1/d_i\) |

> 💡 **Tip**: 가중치 사용 시 가까운 이웃에 더 큰 영향

---

## 🎯 K 값 선택법 (Choosing \(K\))

| 방법                          | 설명                                                                                |
|-------------------------------|-------------------------------------------------------------------------------------|
| **Rule of Thumb**             | \(K \approx \sqrt{N}\) (전체 샘플 수 \(N\)) 기준                                     |
| **교차검증 (Cross-Validation)** | 여러 \(K\) 값에 대해 CV 수행 → 평균 오차 또는 정확도 확인 → 최적 \(K\) 선택          |
| **오류 곡선(Elbow Method)**   | \(K\) 증가에 따른 검증 오차 그래프 작성 → 오차 감소 폭이 완만해지는 지점 포착        |
| **도메인 지식 활용**          | 문제 특성·클래스 분포 고려 → 적절한 이웃 수 조정                                    |
| **연속 vs 불연속 예측**       | 회귀보다 분류에서 \(K\) 민감도↑ → 일반적으로 분류는 \(K\) 크게, 회귀는 작게 시작 권장 |

### 🔍 시각화 예시

```python
from sklearn.model_selection import cross_val_score
import matplotlib.pyplot as plt

ks = range(1, 21)
scores = [cross_val_score(KNeighborsClassifier(n_neighbors=k), X, y, cv=5).mean() for k in ks]

plt.plot(ks, scores)
plt.xlabel('Number of Neighbors K')
plt.ylabel('CV Accuracy')
plt.title('Elbow Curve for K Selection')
plt.show()
```

---

## 💡 실무 팁 요약
1. 스케일링 필수: 거리 계산 민감 → `StandardScaler` 또는 `MinMaxScaler` 적용  
2. 이상치 제거: 이상치가 이웃 검색에 큰 영향 → 사전 처리 필수  
3. 가중치 적용: `weights='distance'` 옵션으로 거리 역수 가중 투표  
4. 메모리·속도 고려: 큰 데이터셋은 `BallTree`/`KDTree` 구조나 근사 이웃 알고리즘(Annoy, FAISS) 사용  
5. 실험: 다양한 거리 지표·K 조합을 교차검증으로 비교  
