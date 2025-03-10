# Python 기타 유용한 팁 정리

Python을 사용할 때 유용한 기능과 팁을 정리합니다.

---

## 1. 리스트 컴프리헨션 (List Comprehension)

리스트를 효율적으로 생성하는 방법입니다.

```python
# 기존 방식
numbers = []
for i in range(10):
    numbers.append(i * 2)

# 리스트 컴프리헨션 사용
numbers = [i * 2 for i in range(10)]
print(numbers)  # [0, 2, 4, 6, 8, 10, 12, 14, 16, 18]
```

### 1.1 조건문 포함
```python
evens = [i for i in range(10) if i % 2 == 0]
print(evens)  # [0, 2, 4, 6, 8]
```

---

## 2. 딕셔너리 컴프리헨션 (Dictionary Comprehension)

딕셔너리를 간결하게 생성하는 방법입니다.

```python
squares = {x: x**2 for x in range(5)}
print(squares)  # {0: 0, 1: 1, 2: 4, 3: 9, 4: 16}
```

---

## 3. `zip()`을 활용한 여러 리스트 병합

여러 개의 리스트를 동시에 순회할 때 사용합니다.

```python
names = ["Alice", "Bob", "Charlie"]
ages = [25, 30, 35]

for name, age in zip(names, ages):
    print(f"{name} is {age} years old.")
```

출력:
```
Alice is 25 years old.
Bob is 30 years old.
Charlie is 35 years old.
```

---

## 4. `enumerate()`를 활용한 인덱스 포함 반복

리스트 반복 시 인덱스와 값을 함께 가져올 수 있습니다.

```python
fruits = ["apple", "banana", "cherry"]

for index, fruit in enumerate(fruits, start=1):
    print(f"{index}: {fruit}")
```

출력:
```
1: apple
2: banana
3: cherry
```

---

## 5. `map()`과 `filter()` 사용하기

### 5.1 `map()`을 사용한 리스트 변환
```python
numbers = [1, 2, 3, 4, 5]
squared = list(map(lambda x: x**2, numbers))
print(squared)  # [1, 4, 9, 16, 25]
```

### 5.2 `filter()`를 사용한 조건 필터링
```python
evens = list(filter(lambda x: x % 2 == 0, numbers))
print(evens)  # [2, 4]
```

---

## 6. `set()`을 활용한 중복 제거

리스트에서 중복을 제거할 때 유용합니다.

```python
numbers = [1, 2, 2, 3, 4, 4, 5]
unique_numbers = list(set(numbers))
print(unique_numbers)  # [1, 2, 3, 4, 5]
```

---

## 7. `collections.Counter()`를 활용한 빈도수 계산

리스트에서 요소별 개수를 쉽게 계산할 수 있습니다.

```python
from collections import Counter

words = ["apple", "banana", "apple", "cherry", "banana", "banana"]
count = Counter(words)
print(count)  # Counter({'banana': 3, 'apple': 2, 'cherry': 1})
```

---

## 8. `itertools`를 활용한 순열과 조합

### 8.1 순열 (Permutations)
```python
from itertools import permutations

data = ['A', 'B', 'C']
perms = list(permutations(data, 2))
print(perms)  # [('A', 'B'), ('A', 'C'), ('B', 'A'), ('B', 'C'), ('C', 'A'), ('C', 'B')]
```

### 8.2 조합 (Combinations)
```python
from itertools import combinations

combs = list(combinations(data, 2))
print(combs)  # [('A', 'B'), ('A', 'C'), ('B', 'C')]
```

---

## 9. `defaultdict`를 사용한 편리한 딕셔너리

존재하지 않는 키를 자동으로 기본값으로 설정할 수 있습니다.

```python
from collections import defaultdict

word_count = defaultdict(int)

words = ["apple", "banana", "apple"]
for word in words:
    word_count[word] += 1

print(word_count)  # defaultdict(<class 'int'>, {'apple': 2, 'banana': 1})
```

---

## 10. 리스트의 요소 개수 제한하기 (`heapq` 사용)

최대값이나 최소값을 빠르게 찾을 때 유용합니다.

```python
import heapq

numbers = [5, 1, 8, 3, 2, 7]
print(heapq.nlargest(3, numbers))  # [8, 7, 5]
print(heapq.nsmallest(3, numbers))  # [1, 2, 3]
```

---

## 11. `functools.lru_cache()`를 활용한 함수 결과 캐싱

자주 호출되는 함수를 캐싱하여 성능을 향상시킬 수 있습니다.

```python
from functools import lru_cache

@lru_cache(maxsize=100)
def fibonacci(n):
    if n < 2:
        return n
    return fibonacci(n-1) + fibonacci(n-2)

print(fibonacci(10))  # 55
```

---

## 12. `dataclasses`를 활용한 간결한 클래스 정의

데이터를 다루는 클래스를 더 간결하게 정의할 수 있습니다.

```python
from dataclasses import dataclass

@dataclass
class Person:
    name: str
    age: int
    city: str

p = Person("Alice", 25, "Seoul")
print(p)  # Person(name='Alice', age=25, city='Seoul')
```

---

## 13. `with` 문을 활용한 파일 처리

파일을 열고 닫는 과정을 자동으로 처리할 수 있습니다.

```python
with open("example.txt", "w") as file:
    file.write("Hello, Python!")
```

---

## 14. `try-except`에서 `else`와 `finally` 활용

```python
try:
    result = 10 / 2
except ZeroDivisionError:
    print("0으로 나눌 수 없습니다.")
else:
    print("연산 성공:", result)
finally:
    print("이 문장은 항상 실행됩니다.")
```

출력:
```
연산 성공: 5.0
이 문장은 항상 실행됩니다.
```

---

## 15. `argparse`를 활용한 명령줄 인자 처리

Python 스크립트에서 명령줄 인자를 쉽게 받을 수 있습니다.

```python
import argparse

parser = argparse.ArgumentParser(description="테스트 프로그램")
parser.add_argument("--name", type=str, help="사용자 이름")

args = parser.parse_args()
print(f"Hello, {args.name}!")
```

실행 예시:
```
python script.py --name Alice
# 출력: Hello, Alice!
```

---

## 16. 정리

| 기능 | 설명 |
|------|------|
| 리스트 컴프리헨션 | 리스트를 간결하게 생성 |
| `zip()` | 여러 리스트를 병합하여 순회 |
| `enumerate()` | 리스트 순회 시 인덱스 포함 |
| `map()`, `filter()` | 데이터 변환 및 필터링 |
| `set()` | 리스트 중복 제거 |
| `Counter()` | 요소 개수 계산 |
| `itertools` | 순열, 조합 생성 |
| `defaultdict` | 딕셔너리 기본값 설정 |
| `heapq` | 최댓값, 최솟값 찾기 |
| `lru_cache()` | 함수 결과 캐싱 |
| `dataclasses` | 간결한 데이터 클래스 |
| `with open()` | 파일 자동 관리 |
| `try-except-else-finally` | 예외 처리 |
| `argparse` | 명령줄 인자 처리 |
