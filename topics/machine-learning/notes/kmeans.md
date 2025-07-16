# 🔍 K-Means 군집화 (Clustering) — 원리 & 시각화

비지도 학습에서 **데이터를 \(K\)개의 군집(cluster)** 으로 나누는 대표적 알고리즘입니다. 각 군집은 내부의 샘플 간 거리가 작고, 군집 간 거리는 크게 유지합니다.

---

## ⚙️ 알고리즘 원리

1. **초기화 (Initialization)**  
   - 군집 수 \(K\) 지정  
   - 중심점(centroid) \( \{\mu_1, \dots, \mu_K\}\) 을 무작위 또는 K-Means++ 방식으로 선택  

2. **할당 단계 (Assignment Step)**  
   \[
   \text{for each } x_i:\quad
   c_i = \arg\min_{j\in\{1..K\}} \; \|x_i - \mu_j\|^2
   \]  
   – 각 샘플을 가장 가까운 중심점에 할당  

3. **업데이트 단계 (Update Step)**  
   \[
   \mu_j = \frac{1}{|C_j|}\sum_{x_i \in C_j} x_i
   \]  
   – 각 군집 \(C_j\) 의 할당된 샘플들의 평균으로 중심점 이동  

4. **수렴 판단 (Convergence Check)**  
   - 중심점 이동 거리가 기준 이하이거나  
   - 최대 반복 횟수 도달 시 종료  

### 🎯 목적 함수 (Within-Cluster Sum of Squares)  
\[
J = \sum_{j=1}^K \sum_{x_i\in C_j} \|x_i - \mu_j\|^2
\]  
– 각 군집 내 거리 제곱합을 최소화

---

## 📊 시각화 기법

### 1️⃣ 엘보우 방법 (Elbow Method)  
- \(K\)를 다양하게 설정하고 \(J(K)\) 값을 계산  
- 그래프에서 “팔꿈치(Elbow)” 지점이 최적 \(K\) 후보  

```python
# 의사 코드 예시
errors = []
for k in range(1,11):
    model = KMeans(n_clusters=k, random_state=0).fit(X)
    errors.append(model.inertia_)  # J(K)
plt.plot(range(1,11), errors, marker='o')
plt.xlabel('Number of Clusters K')
plt.ylabel('Inertia (Sum of Squared Distances)')
plt.show()
```

### 2️⃣ 군집 할당 결과 시각화  
- **2D/3D 산점도**에 각 샘플을 군집 색상별로 표시  
- 중심점(\(\mu_j\)) 마커로 오버레이  

```python
# 의사 코드 예시
labels = model.labels_
centroids = model.cluster_centers_
plt.scatter(X[:,0], X[:,1], c=labels, cmap='tab10')
plt.scatter(centroids[:,0], centroids[:,1], s=200, marker='X', edgecolor='k')
plt.show()
```

### 3️⃣ 실루엣 분석 (Silhouette Analysis)  
- 각 샘플 \(i\) 의 실루엣 계수  
  \[
  s(i) = \frac{b(i) - a(i)}{\max\{a(i),\,b(i)\}}
  \]  
  - \(a(i)\): 같은 군집 내 평균 거리  
  - \(b(i)\): 가장 가까운 다른 군집까지의 평균 거리  
- **전체 평균 실루엣**이 클수록 군집화 품질 우수  

```python
from sklearn.metrics import silhouette_score
score = silhouette_score(X, labels)
print(f"Silhouette Score: {score:.3f}")
```

---

## 💡 실무 팁 요약

1. **스케일링**: 거리 기반 알고리즘이므로 `StandardScaler` 등으로 피처 표준화  
2. **K-Means++ 초기화**: 초기 중심점 선택을 개선해 수렴 속도·안정성 ↑  
3. **반복 횟수**: `max_iter`를 충분히 크게 설정, `tol`로 수렴 기준 조정  
4. **잡음 & 이상치**: 이상치에 민감 → 사전 제거하거나 `MiniBatchKMeans` 고려  
5. **시각화**: 엘보우 + 실루엣 분석으로 최적 \(K\) 결정, 산점도로 군집 구조 직관적 확인  
