# ✖️ Karatsuba 곱셈 (Karatsuba Multiplication)

**Karatsuba 곱셈**은 두 수를 곱할 때, 기존의 `O(n²)` 시간복잡도보다 빠른 **`O(n^log₂3) ~ O(n^1.585)`** 알고리즘입니다.  
**분할 정복** 기법을 사용하며, 큰 수의 곱셈을 빠르게 처리할 수 있습니다.

---

## 1️⃣ 문제 정의

- 두 정수 A, B를 곱해야 함
- 단순 곱셈은 `O(n²)` 시간이 걸림 (자리 수 n개일 경우)
- **Karatsuba 알고리즘은 이를 더 빠르게 처리**

---

## 2️⃣ 핵심 아이디어 (Divide & Conquer)

1. 수 A와 B를 **두 부분으로 나눔**  
   → A = A₁×10ᵈ + A₀, B = B₁×10ᵈ + B₀

2. 전체 곱셈을 다음 항으로 표현:

```
A × B = (A₁×10ᵈ + A₀)(B₁×10ᵈ + B₀)
      = A₁B₁×10²ᵈ + (A₁B₀ + A₀B₁)×10ᵈ + A₀B₀
```

3. Karatsuba는 중간항을 다음처럼 줄임:

```
(A₁ + A₀)(B₁ + B₀) - A₁B₁ - A₀B₀ = A₁B₀ + A₀B₁
```

→ 즉, 곱셈 4번 대신 곱셈을 **3번만 수행**하게 됨

---

## 3️⃣ 파이썬 코드 예제

```python
def karatsuba(x, y):
    # 기본 곱셈 처리
    if x < 10 or y < 10:
        return x * y

    # 자리 수 계산
    n = max(len(str(x)), len(str(y)))
    m = n // 2

    # 자릿수 분할
    high1, low1 = divmod(x, 10**m)
    high2, low2 = divmod(y, 10**m)

    # 분할된 부분 곱셈
    z0 = karatsuba(low1, low2)
    z1 = karatsuba((low1 + high1), (low2 + high2))
    z2 = karatsuba(high1, high2)

    return (z2 * 10**(2*m)) + ((z1 - z2 - z0) * 10**m) + z0
```

---

## 4️⃣ 예제 입력

```python
a = 1234
b = 5678
print(karatsuba(a, b))  # 출력: 7006652
```

---

## 5️⃣ 시간 복잡도

- 일반 곱셈: `O(n²)`
- Karatsuba: **`O(n^log₂3) ~ O(n^1.585)`**
- n이 충분히 클 때 성능 차이 발생

---

## 6️⃣ 사용 예시

- 큰 수 연산 (정수 범위를 초과하는 곱셈)
- 임베디드, 보안, 암호화 분야의 빠른 곱셈 연산
- 다항식 곱셈 최적화

---

## 🎯 정리 요약

✔ 자리 수를 반으로 나눠서 곱셈 수를 줄임  
✔ 4번 → 3번 곱셈으로 최적화  
✔ 분할 정복 기반 구조  
✔ 빠른 큰 수 연산에 유리  
