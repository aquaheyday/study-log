# 📈 알고리즘 - 최장 증가 부분 수열(Longest Increasing Subsequence, LIS)

**최장 증가 부분 수열(LIS)** 이란, 주어진 수열에서 **오름차순으로 정렬된 가장 긴 부분 수열**을 찾는 문제입니다.

---

## 1️⃣ 문제 설명

주어진 수열에서 **순서를 유지하며**, 숫자가 **점점 증가하는 부분 수열 중 가장 긴 것**의 길이를 구하세요.

---

## 2️⃣ 예시

### 입력:
```
[10, 20, 10, 30, 20, 50]
```

### 가능한 증가 수열:
- [10, 20, 30, 50]
- [10, 20, 50]

### LIS 길이: **4**

---

## 3️⃣ 접근 방식

### 방법 1: 동적 계획법 (DP) `O(N²)`

- `dp[i]`: i번째 원소까지 고려했을 때 LIS의 길이
- 점화식:
  ```text
  dp[i] = max(dp[j] + 1)  (j < i and arr[j] < arr[i])
  ```
- 초기값: 모든 `dp[i] = 1`

---

### 방법 2: 이분 탐색 + 그리디 `O(N log N)`

- 배열을 따로 만들어 **LIS를 흉내만 냄**
- 각 수에 대해:
  - **해당 값이 LIS 배열 끝보다 크면** → 맨 뒤에 추가
  - **작거나 같으면** → 적절한 위치에 교체 (`lower_bound` 사용)

---

## 4️⃣ 코드 예시

### 방법 1: DP `O(N²)`

```python
def lis(arr):
    n = len(arr)
    dp = [1] * n

    for i in range(n):
        for j in range(i):
            if arr[j] < arr[i]:
                dp[i] = max(dp[i], dp[j] + 1)
    
    return max(dp)

print(lis([10, 20, 10, 30, 20, 50]))  # 출력: 4
```

---

### 방법 2: 이분 탐색 + 그리디 `O(N log N)`

```python
import bisect

def lis(arr):
    result = []

    for num in arr:
        idx = bisect.bisect_left(result, num)
        if idx == len(result):
            result.append(num)
        else:
            result[idx] = num
    
    return len(result)

print(lis([10, 20, 10, 30, 20, 50]))  # 출력: 4
```

> ❗ `result` 배열은 실제 LIS는 아니지만, **길이는 정확함**

---

## 5️⃣ 시간 복잡도 비교

| 방식            | 시간 복잡도 | 설명 |
|-----------------|--------------|------|
| DP (중첩 반복문) | `O(N²)`      | 작은 입력에서 충분 |
| 이분 탐색 방식   | `O(N log N)` | 대규모 데이터 처리 가능 (최적화된 풀이)

---

## 6️⃣ 응용 문제 예시

- 실제 LIS 수열 출력하기 (역추적 필요)
- 감소 수열 LIS (Longest Decreasing Subsequence)
- 바이토닉 수열, 전깃줄 문제 (백준 2565)
- 2차원 LIS: 좌표 정렬 + LIS
- LCS(최장 공통 부분 수열)과의 차이 비교

---

## 🎯 정리 요약

✔ LIS는 수열에서 **오름차순이 되도록 선택한 가장 긴 부분 수열**  
✔ DP로 쉽게 구현 가능 (`O(N²)`)  
✔ `bisect`를 활용한 이분 탐색으로 `O(N log N)` 최적화 가능  
✔ 실전에서는 대부분 길이만 구하는 문제 → **이분 탐색 버전이 더 유리**

