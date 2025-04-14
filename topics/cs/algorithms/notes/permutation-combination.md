# 🔁 알고리즘 - 순열 & 조합 생성

**순열(Permutation)** 과 **조합(Combination)** 은 **N개의 원소 중 일부 또는 전부를 선택**하는 기본적인 탐색 기법입니다.

---

## 1️⃣ 문제 정의

### 순열 (Permutation)
- 서로 다른 **N개 중에서 R개를 골라** **순서를 고려하여 나열**하는 경우의 수  
- 예: A, B, C → A-B, B-A 등 (순서 다르면 다른 경우)
- 수식: `n! / (n - r)!`

### 조합 (Combination)
- 서로 다른 **N개 중에서 R개를 골라**, **순서를 고려하지 않고 선택**하는 경우의 수  
- 예: A, B, C → A-B와 B-A는 같은 경우 (순서 무시)
- 수식: `n! / (r! × (n - r)!)`

---

## 2️⃣ 예시 (N = 3, R = 2)

### 순열  

입력: `['A', 'B', 'C']`, R = 2  

#### 가능한 결과:

```
A B  
A C  
B A  
B C  
C A  
C B  
```

총 **6가지** (3P2, 3P2 = 3 × 2 = 6)

---

### 조합  

입력: `['A', 'B', 'C']`, R = 2  

#### 가능한 결과:

```
A B  
A C  
B C  
```

총 **3가지** (3C2, 3C2 = 3! / (2! × (3-2)!) = 3)

---

## 3️⃣ 해결 방법: `itertools` 사용 or 직접 구현

### 순열 (itertools)

```python
from itertools import permutations

data = ['A', 'B', 'C']
result = list(permutations(data, 2))  # 2개 뽑는 순열
print(result)
# [('A', 'B'), ('A', 'C'), ('B', 'A'), ('B', 'C'), ('C', 'A'), ('C', 'B')]
```

---

### 조합 (itertools)

```python
from itertools import combinations

data = ['A', 'B', 'C']
result = list(combinations(data, 2))  # 2개 뽑는 조합
print(result)
# [('A', 'B'), ('A', 'C'), ('B', 'C')]
```

---

## 4️⃣ 직접 구현 (백트래킹 방식)

### 순열 - 직접 구현

```python
def generate_permutations(arr, r):
    result = []
    visited = [False] * len(arr)

    def backtrack(path):
        if len(path) == r:
            result.append(path[:])
            return
        for i in range(len(arr)):
            if not visited[i]:
                visited[i] = True
                path.append(arr[i])
                backtrack(path)
                path.pop()
                visited[i] = False

    backtrack([])
    return result

print(generate_permutations(['A', 'B', 'C'], 2))
```

---

### 🛠 조합 - 직접 구현

```python
def generate_combinations(arr, r):
    result = []

    def backtrack(start, path):
        if len(path) == r:
            result.append(path[:])
            return
        for i in range(start, len(arr)):
            path.append(arr[i])
            backtrack(i + 1, path)
            path.pop()

    backtrack(0, [])
    return result

print(generate_combinations(['A', 'B', 'C'], 2))
```

---

## 5️⃣ 시간 복잡도

| 종류 | 시간 복잡도 |
|------|--------------|
| 순열 | O(N!)       |
| 조합 | O(N! / (R! * (N-R)!)) |

- 순열은 경우의 수가 많기 때문에 빠르게 가지치기(backtracking)하는 것이 중요
- itertools는 내부적으로 최적화 되어 있어 일반적인 경우 사용 추천

---

## 6️⃣ 주요 포인트 요약

| 항목       | 설명 |
|------------|------|
| 순열       | 순서 O / 중복 X |
| 조합       | 순서 X / 중복 X |
| itertools | 코드 짧고 간단 (추천) |
| 직접 구현 | 알고리즘 원리 이해에 유용 |
| 백트래킹   | 재귀 기반으로 경로 구성하며 탐색 |

---

## 🎯 정리 요약

✔ **순열**은 순서를 고려해 나열, **조합**은 순서 무시하고 선택  
✔ `itertools.permutations`, `combinations` 사용하면 매우 간단  
✔ **직접 구현** 시 백트래킹으로 모든 경우의 수 탐색 가능  
✔ **탐색 문제**에서 많이 사용되므로, 원리 + 라이브러리 숙지 필수
