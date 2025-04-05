# 🧮 Data Structures - 순열(Permutation) / 조합(Combination) 구조

**순열과 조합**은 경우의 수를 다루는 핵심 개념으로, 알고리즘 문제, 확률 계산, 백트래킹 등에 매우 자주 사용됩니다.

> Python에서는 `itertools` 모듈을 활용하여 쉽게 구현할 수 있습니다.

---

## 1️⃣ 순열 (Permutation)

### 1) 개념
- **순서를 고려한** 나열 방법
- n개 중에서 r개를 **골라 순서를 고려해 나열**
- 공식: `nPr = n! / (n - r)!`

---

### 2) Python 예시

```python
from itertools import permutations

data = [1, 2, 3]

# 모든 순열 (3P3)
for p in permutations(data, 3):
    print(p)

# 일부 순열 (3P2)
for p in permutations(data, 2):
    print(p)
```

#### 예시 출력
```text
(1, 2, 3)
(1, 3, 2)
(2, 1, 3)
...
```

---

## 2️⃣ 조합 (Combination)

### 1) 개념
- **순서를 고려하지 않고** 뽑는 방법
- n개 중 r개를 **순서 없이 고르는 경우**
- 공식: `nCr = n! / (r! * (n - r)!)`

---

### 2) Python 예시

```python
from itertools import combinations

data = ['A', 'B', 'C']

# 모든 조합 (3C2)
for c in combinations(data, 2):
    print(c)
```

#### 예시 출력
```text
('A', 'B')
('A', 'C')
('B', 'C')
```

---

## 3️⃣ 중복 순열 / 중복 조합

### 1) 중복 순열 (Product)

```python
from itertools import product

data = [1, 2]

# 중복을 허용한 순열 (2자리)
for p in product(data, repeat=2):
    print(p)
```

---

### 2) 중복 조합 (Combinations with Replacement)

```python
from itertools import combinations_with_replacement

data = ['A', 'B']

# 중복을 허용한 조합 (2개 선택)
for c in combinations_with_replacement(data, 2):
    print(c)
```

---

## 4️⃣ 요약 비교표

| 구분        | 중복 허용 | 순서 고려 | 함수                            |
|-------------|-----------|------------|----------------------------------|
| 순열        | X         | O          | `itertools.permutations()`      |
| 조합        | X         | X          | `itertools.combinations()`      |
| 중복 순열   | O         | O          | `itertools.product(repeat=r)`   |
| 중복 조합   | O         | X          | `itertools.combinations_with_replacement()` |

---

## 5️⃣ 활용 예시

- 백트래킹(DFS) 기반 조합/순열 탐색  
- 로또 번호 생성기  
- 암호 조합, 게임 아이템 조합  
- 브루트포스 탐색 문제  
- 확률 및 경우의 수 문제

---

## 🎯 정리 요약

✔ **순열**: 순서 O, 중복 X → `permutations()`  
✔ **조합**: 순서 X, 중복 X → `combinations()`  
✔ **중복 순열**: 순서 O, 중복 O → `product()`  
✔ **중복 조합**: 순서 X, 중복 O → `combinations_with_replacement()`  
✔ Python의 `itertools`로 간편하게 구현 가능

