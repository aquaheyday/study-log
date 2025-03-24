# Python 자료구조 기본 정리

Python에서 자주 사용되는 기본 자료구조를 정리합니다.

---

## 1. 리스트 (List)
리스트는 변경 가능한(mutable) 순서가 있는 데이터 구조입니다.

### 1.1 리스트 선언과 사용
```python
# 리스트 선언
numbers = [1, 2, 3, 4, 5]
mixed = [1, "Hello", 3.14, True]

# 리스트 요소 접근
print(numbers[0])   # 1
print(numbers[-1])  # 5 (마지막 요소)

# 리스트 슬라이싱
print(numbers[1:4])  # [2, 3, 4]
```

### 1.2 리스트 메서드
```python
numbers.append(6)   # 요소 추가
numbers.insert(2, 99)  # 특정 위치에 요소 삽입
numbers.remove(3)   # 특정 값 제거
numbers.pop()       # 마지막 요소 제거
numbers.sort()      # 정렬
numbers.reverse()   # 역순 정렬
print(numbers)
```

---

## 2. 튜플 (Tuple)
튜플은 리스트와 비슷하지만 **변경할 수 없는(immutable)** 자료구조입니다.

### 2.1 튜플 선언과 사용
```python
# 튜플 선언
t = (1, 2, 3, 4, 5)

# 요소 접근
print(t[0])    # 1
print(t[1:3])  # (2, 3)

# 요소 변경 불가능
# t[0] = 10  # 오류 발생!
```

### 2.2 튜플 언패킹
```python
a, b, c = (1, 2, 3)
print(a, b, c)  # 1 2 3
```

---

## 3. 딕셔너리 (Dictionary)
딕셔너리는 키(key)와 값(value)의 쌍으로 이루어진 자료구조입니다.

### 3.1 딕셔너리 선언과 사용
```python
person = {"name": "Alice", "age": 25, "city": "Seoul"}

# 값 접근
print(person["name"])  # Alice
print(person.get("age"))  # 25
```

### 3.2 딕셔너리 메서드
```python
person["job"] = "Developer"  # 키-값 추가
del person["age"]            # 키 삭제
print(person.keys())         # 키 목록
print(person.values())       # 값 목록
print(person.items())        # (키, 값) 쌍 목록
```

---

## 4. 집합 (Set)
집합은 중복을 허용하지 않는 **순서가 없는(unordered)** 자료구조입니다.

### 4.1 집합 선언과 사용
```python
# 집합 선언
s = {1, 2, 3, 4, 5}

# 요소 추가 및 삭제
s.add(6)   # 추가
s.remove(3)  # 제거
print(s)
```

### 4.2 집합 연산
```python
a = {1, 2, 3, 4}
b = {3, 4, 5, 6}

print(a | b)  # 합집합 {1, 2, 3, 4, 5, 6}
print(a & b)  # 교집합 {3, 4}
print(a - b)  # 차집합 {1, 2}
```

---

## 5. 스택 (Stack)
스택은 **후입선출(LIFO, Last In First Out)** 구조입니다.

### 5.1 스택 구현 (리스트 사용)
```python
stack = []
stack.append(1)
stack.append(2)
stack.append(3)
print(stack.pop())  # 3
print(stack.pop())  # 2
print(stack)  # [1]
```

---

## 6. 큐 (Queue)
큐는 **선입선출(FIFO, First In First Out)** 구조입니다.

### 6.1 큐 구현 (collections.deque 사용)
```python
from collections import deque

queue = deque()
queue.append(1)
queue.append(2)
queue.append(3)
print(queue.popleft())  # 1
print(queue.popleft())  # 2
print(queue)  # deque([3])
```

---

## 7. 우선순위 큐 (Priority Queue)
우선순위 큐는 **값의 우선순위에 따라 먼저 나가는** 자료구조입니다.

### 7.1 우선순위 큐 구현 (heapq 사용)
```python
import heapq

pq = []
heapq.heappush(pq, (2, "Task2"))
heapq.heappush(pq, (1, "Task1"))  # 우선순위가 높은 값 (1)이 먼저 나옴
heapq.heappush(pq, (3, "Task3"))

print(heapq.heappop(pq))  # (1, "Task1")
print(heapq.heappop(pq))  # (2, "Task2")
```
