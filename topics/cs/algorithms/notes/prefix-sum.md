# 📊 누적합 (Prefix Sum)

**누적합(Prefix Sum)** 은 배열이나 수열에서 **구간 합을 빠르게 계산**하기 위해 사용하는 전처리 기법입니다.  
한 번 계산해두면 **임의 구간의 합을 O(1)에 구할 수 있어** 매우 효율적입니다.

---

## 1️⃣ 문제 정의

- 배열에서 **i번째부터 j번째까지의 합**을 자주 구해야 할 때
- 반복적으로 `sum(arr[i:j+1])` 하면 느리므로, **미리 누적합 배열을 만들어 O(1)로 합 계산**

---

## 2️⃣ 핵심 아이디어

- **누적합 배열 prefix[i] = 0 ~ i번째까지의 합**

### 공식

- `prefix[0] = arr[0]`  
- `prefix[i] = prefix[i - 1] + arr[i]`

### 구간 합 계산

- `sum(i ~ j) = prefix[j] - prefix[i - 1]`  
- 단, i = 0이면 → `sum = prefix[j]`

---

## 3️⃣ 예시

### 입력 배열
```
arr = [3, 2, 4, 1, 5]
```

### 누적합 배열
```
prefix = [3, 5, 9, 10, 15]
```

### 예: 1번 ~ 3번 인덱스의 합
```
sum(1 ~ 3) = prefix[3] - prefix[0] = 10 - 3 = 7
```

---

## 4️⃣ 파이썬 코드 예제

```python
# 배열로부터 누적합 배열 만들기
def build_prefix_sum(arr):
    prefix = [0] * len(arr)
    prefix[0] = arr[0]
    for i in range(1, len(arr)):
        prefix[i] = prefix[i - 1] + arr[i]
    return prefix

# 구간 합 계산
def range_sum(prefix, i, j):
    if i == 0:
        return prefix[j]
    return prefix[j] - prefix[i - 1]

# 예시
arr = [3, 2, 4, 1, 5]
prefix = build_prefix_sum(arr)

print(range_sum(prefix, 1, 3))  # 출력: 7 (2+4+1)
```

---

## 5️⃣ 시간 복잡도

- 누적합 배열 생성: **O(N)**
- 구간 합 계산: **O(1)**

---

## 6️⃣ 활용 예시

- 구간 합을 빠르게 구하는 문제
- 누적 차이 계산 (이동 거리, 누적 비용 등)
- 2차원 배열 누적합 (누적합 + 사각형 범위 계산)
- 부분합 조건 만족하는 구간 찾기 (투 포인터와 함께 사용)

---

## 🎯 정리 요약

✔ 구간 합 계산을 O(1)로 줄여줌  
✔ 전처리로 prefix 배열 한 번 만들면 이후 빠름  
✔ i ~ j의 합 = `prefix[j] - prefix[i - 1]`  
✔ 2차원으로 확장 가능 (2D Prefix Sum)  
