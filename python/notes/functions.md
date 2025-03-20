# 🔲 Python 함수 사용법

Python에서 함수(Function)는 코드의 재사용성을 높이고 유지보수를 쉽게 하기 위해 사용됩니다.  
이 문서에서는 함수 정의, 매개변수, 반환값, 람다 함수, 스코프 등을 정리합니다.

---

## 1️⃣ 함수란?

- 함수(Function)는 특정 작업을 수행하는 코드 블록입니다.
- `def` 키워드를 사용하여 함수를 정의합니다.
- 필요할 경우 매개변수(Arguments)를 받아 실행할 수 있습니다.
- 값을 반환(`return`)할 수도 있습니다.

---

## 2️⃣ 함수 정의 및 호출

### 기본 함수 정의 및 호출
```python
def greet():
    print("Hello, Python!")

greet()  # Hello, Python!
```

✔ 함수는 `def` 키워드로 정의하며, 호출하려면 함수 이름을 사용  
✔ 함수가 호출될 때 `print()` 문이 실행됨  

---

## 3️⃣ 매개변수와 인자 (Parameters & Arguments)

함수는 매개변수를 받아서 동적으로 동작할 수 있습니다.

### 1) 매개변수가 있는 함수
```python
def greet(name):
    print(f"Hello, {name}!")

greet("Alice")  # Hello, Alice!
greet("Bob")    # Hello, Bob!
```

✔ `name`은 함수 내부에서만 사용되는 지역 변수  

---

### 2) 기본값 매개변수 (Default Parameter)
```python
def greet(name="Guest"):
    print(f"Hello, {name}!")

greet()        # Hello, Guest!
greet("Alice") # Hello, Alice!
```

✔ 인자를 제공하지 않으면 기본값(`Guest`)이 사용됨  

---

### 3) 키워드 인자 (Keyword Arguments)
```python
def introduce(name, age):
    print(f"My name is {name}, and I am {age} years old.")

introduce(age=30, name="Alice")  # 순서와 관계없이 인자 전달 가능
```

✔ 키워드를 사용하여 인자 순서를 바꿔도 정확하게 전달 가능  

---

### 4) 가변 인자 (`*args`) - 여러 개의 인자를 받을 때
```python
def add_numbers(*args):
    return sum(args)

print(add_numbers(1, 2, 3))       # 6
print(add_numbers(10, 20, 30, 40)) # 100
```

✔ `*args`는 튜플 형태로 전달되며, 여러 개의 인자를 받을 수 있음  

---

### 5) 키워드 가변 인자 (`**kwargs`) - 여러 개의 키워드 인자를 받을 때
```python
def print_info(**kwargs):
    for key, value in kwargs.items():
        print(f"{key}: {value}")

print_info(name="Alice", age=25, city="New York")
```

✔ `**kwargs`는 딕셔너리 형태로 전달되며, 여러 개의 키워드 인자를 받을 수 있음  

---

## 4️⃣ 반환값 (`return`)

함수는 결과를 반환할 수 있습니다.

### 1) 반환값이 있는 함수
```python
def add(a, b):
    return a + b

result = add(3, 5)
print(result)  # 8
```

✔ `return` 키워드를 사용하여 값을 반환  
✔ 반환된 값은 변수에 저장할 수 있음  

---

### 2) 여러 개의 값 반환 (Tuple 반환)
```python
def get_user():
    name = "Alice"
    age = 30
    return name, age  # 튜플 형태로 반환

user = get_user()
print(user)        # ('Alice', 30)
print(user[0])     # Alice
print(user[1])     # 30
```

✔ 여러 값을 반환할 때는 **튜플(Tuple)** 로 반환됨  

---

## 5️⃣ 람다 함수 (Lambda Function)

람다 함수는 **한 줄짜리 익명 함수**를 만들 때 사용됩니다.

### 1) 기본 람다 함수
```python
add = lambda x, y: x + y
print(add(3, 5))  # 8
```

✔ `lambda` 키워드를 사용하여 함수를 간결하게 정의  

---

### 2) `map()`, `filter()` 와 함께 사용
```python
numbers = [1, 2, 3, 4, 5]

squared = list(map(lambda x: x**2, numbers))  # 모든 요소 제곱
print(squared)  # [1, 4, 9, 16, 25]

evens = list(filter(lambda x: x % 2 == 0, numbers))  # 짝수만 필터링
print(evens)  # [2, 4]
```

✔ `map()` → 모든 요소에 함수 적용  
✔ `filter()` → 특정 조건을 만족하는 요소만 필터링  

---

## 6️⃣ 함수 스코프 (Scope)

함수 내부와 외부에서 변수의 접근 범위를 다룹니다.

### 1) 지역 변수 (Local Variable)
```python
def my_function():
    x = 10  # 함수 내부에서만 사용 가능
    print(x)

my_function()
print(x)  # 오류 발생 (NameError)
```

✔ 함수 내부에서 선언된 변수는 함수 밖에서 접근할 수 없음  

---

### 2) 전역 변수 (Global Variable)
```python
x = 10  # 전역 변수

def my_function():
    global x  # 전역 변수 수정
    x = 20
    print(x)

my_function()
print(x)  # 20
```

✔ `global` 키워드를 사용하면 함수 내부에서 전역 변수 수정 가능  
✔ 하지만 전역 변수 사용은 신중하게 해야 함 (가독성 및 유지보수에 영향)  

---

## 7️⃣ 재귀 함수 (Recursive Function)

재귀 함수는 자기 자신을 호출하는 함수입니다.

### 팩토리얼 계산 (Factorial)
```python
def factorial(n):
    if n == 1:
        return 1
    return n * factorial(n - 1)

print(factorial(5))  # 120
```

✔ 재귀 함수는 종료 조건(`if n == 1`)이 반드시 필요함  

---

## 🎯 정리

✔ **함수 정의**: `def` 키워드를 사용하여 정의  
✔ **매개변수 & 인자**: 기본값, `*args`, `**kwargs` 사용 가능  
✔ **반환값**: `return`을 사용하여 값 반환 (여러 값 반환 가능)  
✔ **람다 함수**: 간단한 익명 함수 작성 가능 (`lambda x: x + 1`)  
✔ **스코프**: 지역 변수(`local`)와 전역 변수(`global`) 개념 이해 필요  
✔ **재귀 함수**: 자기 자신을 호출하는 함수, 종료 조건 필수  
