# 📌 Python 조건문과 반복문

Python에서 조건문과 반복문을 활용하여 프로그램의 흐름을 제어하는 방법을 정리합니다.

---

## 1️⃣ 조건문 (Conditional Statements)

조건문은 특정 조건을 만족할 때 코드 블록을 실행하는 데 사용됩니다.

### 1) `if` 문 (기본 조건문)
```python
x = 10

if x > 5:
    print("x는 5보다 크다")
```

✔ Python에서는 중괄호 `{}` 대신 들여쓰기(Indentation)로 블록을 구분함  
✔ 조건이 참(`True`)이면 해당 블록이 실행됨  

---

### 2) `if-else` 문 (조건이 참/거짓일 때 실행)
```python
x = 3

if x > 5:
    print("x는 5보다 크다")
else:
    print("x는 5 이하이다")
```

✔ `else` 블록은 `if` 조건이 거짓(`False`)일 때 실행됨  

---

### 3) `if-elif-else` 문 (여러 조건 검사)
```python
score = 85

if score >= 90:
    print("A 학점")
elif score >= 80:
    print("B 학점")
elif score >= 70:
    print("C 학점")
else:
    print("F 학점")
```

✔ `elif`는 추가적인 조건을 검사할 때 사용  
✔ 조건이 위에서부터 순차적으로 평가됨 → 첫 번째로 참(`True`)인 조건이 실행됨  

---

### 4) 조건문에서 논리 연산자 사용 (`and`, `or`, `not`)
```python
age = 20
is_student = True

if age >= 18 and is_student:
    print("성인 학생입니다.")
```

✔ `and` → 두 조건이 모두 참(`True`)이어야 실행  
✔ `or` → 둘 중 하나라도 참이면 실행  
✔ `not` → 논리값을 반전 (`True` ↔ `False`)  

---

## 2️⃣ 반복문 (Loops)

Python에서 반복문은 특정 코드 블록을 여러 번 실행하는 데 사용됩니다.

### 1) `for` 문 (리스트, 문자열, 범위 반복)
```python
fruits = ["사과", "바나나", "체리"]

for fruit in fruits:
    print(fruit)
```

✔ 리스트, 튜플, 문자열 등의 **반복 가능한 객체(iterable)** 를 순회  

---

### 2) `for` 문과 `range()` 함수 (숫자 범위 반복)
```python
for i in range(5):  # 0부터 4까지 반복
    print(i)
```

---

#### 3) `range(start, stop, step)` 형태로 사용 가능  
```python
for i in range(1, 10, 2):  # 1부터 9까지 2씩 증가
    print(i)
```

---

### 4) `while` 문 (조건이 참인 동안 반복)
```python
x = 0

while x < 5:
    print(x)
    x += 1
```

✔ `while` 문은 조건이 참(`True`)인 동안 계속 반복  
✔ 반복을 중단하려면 조건을 변경해야 함 (`x += 1` 등의 업데이트 필수)  

---

## 3️⃣ 반복문 제어 (Loop Control)

### 1) `break` 문 (반복문 종료)
```python
for i in range(10):
    if i == 5:
        break  # i가 5일 때 반복문 종료
    print(i)
```

✔ `break`를 만나면 반복문이 즉시 종료됨  

---

### 2) `continue` 문 (현재 반복 건너뛰기)
```python
for i in range(5):
    if i == 2:
        continue  # i가 2일 때 현재 반복을 건너뜀
    print(i)
```

✔ `continue`를 만나면 현재 반복만 건너뛰고, 다음 반복으로 진행  

---

### 3) `else` 블록과 반복문 (반복문이 정상 종료될 때 실행)
```python
for i in range(5):
    print(i)
else:
    print("반복문이 정상적으로 종료되었습니다.")
```

✔ `break`로 중단되지 않은 경우, 반복문이 종료된 후 `else` 블록 실행  

---

## 🎯 정리

✔ **조건문**  
- `if`, `if-else`, `if-elif-else`를 활용하여 조건 분기 가능  
- 논리 연산자(`and`, `or`, `not`)를 이용한 복합 조건 처리  

✔ **반복문**  
- `for` 문을 사용하여 리스트, 문자열, `range()` 등의 요소 순회 가능  
- `while` 문은 조건이 참(`True`)인 동안 실행  

✔ **반복문 제어**  
- `break` → 반복문 즉시 종료  
- `continue` → 현재 반복을 건너뛰고 다음 반복으로 진행  
- `else` → 반복문이 정상 종료된 경우 실행  
