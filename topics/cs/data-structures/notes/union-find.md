# 🗃️ Data Structures - 유니온 파인드 (Disjoint Set / Union-Find)

**유니온 파인드(Disjoint Set)** 는 **서로소 집합**을 관리하기 위한 자료구조입니다.  
주로 **그래프에서 사이클 판별**, **네트워크 연결 여부 확인**, **MST 알고리즘(크루스칼)** 등에서 사용됩니다.

> 핵심 연산은 두 가지: `Find`(집합 찾기), `Union`(집합 합치기)

#### 서로소(Disjoint)의 의미
- 두 집합 A, B가 있을 때, **A ∩ B = ∅ (공집합)** 이면 → A와 B는 서로소라고 함.
- 즉, 공통된 원소가 하나도 없음

---

## 1️⃣ 유니온 파인드의 주요 특징

- 원소들을 **여러 개의 집합**으로 나눠 관리  
- 두 원소가 **같은 집합에 속해 있는지 확인** 가능  
- **Find 연산**: 특정 원소가 속한 집합의 대표 찾기  
- **Union 연산**: 두 집합을 하나로 합침  
- **트리 기반 구조 + 최적화 기법**으로 빠른 연산 가능

---

## 2️⃣ 주요 연산 및 시간 복잡도

| 연산       | 설명                               | 시간 복잡도 |
|------------|------------------------------------|--------------|
| `Find(x)`  | 원소 x가 속한 집합의 대표(루트) 찾기 | O(α(N))      |
| `Union(x, y)` | x와 y가 속한 두 집합을 합치기     | O(α(N))      |

> α(N): 매우 느리게 증가하는 **아커만 함수의 역함수**  
> 사실상 **상수 시간**이라고 보면 됨

---

## 3️⃣ 파이썬 구현 예시

```python
class DisjointSet:
    def __init__(self, n):
        self.parent = [i for i in range(n)]

    def find(self, x):
        # 경로 압축 (Path Compression)
        if self.parent[x] != x:
            self.parent[x] = self.find(self.parent[x])
        return self.parent[x]

    def union(self, x, y):
        x_root = self.find(x)
        y_root = self.find(y)
        if x_root != y_root:
            self.parent[y_root] = x_root
```

---

## 4️⃣ 최적화 기법

| 기법                  | 설명                                              | 효과         |
|-----------------------|---------------------------------------------------|--------------|
| 경로 압축 (Path Compression) | `find()` 호출 시 경로상의 노드들을 루트로 직접 연결 | 트리 높이 감소 |
| 랭크 기반 유니온 (Union by Rank / Size) | 더 낮은 트리를 높은 트리에 붙임                   | 트리 균형 유지 |

> 두 기법을 같이 쓰면 **거의 상수 시간**으로 동작

---

## 5️⃣ 유니온 파인드의 장점

✅ **효율적인 집합 관리 (`O(1)` 수준)**  
✅ **서로소 판별 및 합집합 연산에 최적**  
✅ 그래프에서 **사이클 탐지**, **MST(크루스칼)** 등에 활용  

---

## 6️⃣ 유니온 파인드의 단점

⚠️ **정렬된 탐색**이나 **전체 탐색**에는 적합하지 않음  
⚠️ 내부 구조가 **트리 기반**이라 직관적으로 보기 어려울 수 있음  
⚠️ **초기 노드 수 지정 필요**

---

## 7️⃣ 활용 사례

- **사이클 판별 (그래프 탐색)**  
- **네트워크 연결 여부 확인**  
- **크루스칼 알고리즘 (최소 신장 트리)**  
- **집단 분리/합치기 문제**  
- **동적 집합 추적 문제**

---

## 🎯 정리 요약

✔ 유니온 파인드(Disjoint Set)는 **서로소 집합 자료구조**  
✔ `Find`와 `Union` 연산이 핵심 (`O(α(N))` 거의 상수)  
✔ 경로 압축 + 랭크 최적화로 **고속 집합 처리 가능**  
✔ 그래프 문제, 네트워크, 그룹 분리/합치 등에 자주 쓰임
