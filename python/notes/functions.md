# Python 함수와 스코프 정리

## 1. 함수(Function)란?
Python에서 함수는 특정 작업을 수행하는 코드 블록입니다. 함수는 코드의 재사용성을 높이고, 유지보수를 쉽게 하며, 프로그램을 구조화하는 데 도움을 줍니다.

### 1.1 함수 정의와 호출
```python
# 함수 정의
def greet(name):
    return f"Hello, {name}!"

# 함수 호출
print(greet("Python"))  # Hello, Python!
```

### 1.2 매개변수와 반환값
```python
# 두 개의 숫자를 더하는 함수
def add(a, b):
    return a + b

print(add(3, 5))  # 8
```

### 1.3 기본값 매개변수 (Default Parameter)
```python
def greet(name="Guest"):
    return f"Hello, {name}!"

print(greet())       # Hello, Guest!
print(greet("Alice"))  # Hello, Alice!
```

### 1.4 키워드 인자 (Keyword Arguments)
```python
def introduce(name, age):
    return f"My name is {name} and I am {age} years old."

print(introduce(age=25, name="Alice"))  
# My name is Alice and I am 25 years old.
```

### 1.5 가변 인자 (*args)
```python
def add_all(*args):
    return sum(args)

print(add_all(1, 2, 3, 4, 5))  # 15
```

### 1.6 가변 키워드 인자 (**kwargs)
```python
def introduce(**kwargs):
    for key, value in kwargs.items():
        print(f"{key}: {value}")

introduce(name="Alice", age=25, city="Seoul")
# name: Alice
# age: 25
# city: Seoul
```

---

## 2. 함수의 스코프(Scope)

Python에서 변수의 스코프(범위)는 변수가 어디에서 유효한지를 결정합니다.

### 2.1 지역 변수(Local Variable)
함수 내부에서 선언된 변수는 함수 내에서만 유효합니다.
```python
def my_function():
    local_var = "I am local"
    print(local_var)

my_function()
# print(local_var)  # 오류 발생! (local_var는 함수 밖에서 접근 불가)
```

### 2.2 전역 변수(Global Variable)
함수 바깥에서 선언된 변수는 프로그램 전체에서 접근할 수 있습니다.
```python
global_var = "I am global"

def my_function():
    print(global_var)  # 함수 내부에서도 접근 가능

my_function()
print(global_var)
```

### 2.3 `global` 키워드 사용
함수 내부에서 전역 변수를 수정하려면 `global` 키워드를 사용해야 합니다.
```python
count = 0

def increase():
    global count
    count += 1

increase()
print(count)  # 1
```

### 2.4 `nonlocal` 키워드 사용
중첩 함수(함수 안의 함수)에서 바깥 함수의 변수를 수정하려면 `nonlocal` 키워드를 사용합니다.
```python
def outer():
    x = 10

    def inner():
        nonlocal x
        x += 5
        print("Inner x:", x)

    inner()
    print("Outer x:", x)

outer()
# Inner x: 15
# Outer x: 15
```

---

## 3. 람다 함수 (Lambda Function)
람다 함수는 `lambda` 키워드를 사용하여 간단한 익명 함수를 만들 때 사용합니다.
```python
square = lambda x: x ** 2
print(square(5))  # 25
```

람다 함수는 여러 개의 매개변수를 받을 수도 있습니다.
```python
add = lambda a, b: a + b
print(add(3, 5))  # 8
```

람다 함수는 `map()`, `filter()`, `sorted()` 등의 함수와 함께 자주 사용됩니다.
```python
numbers = [1, 2, 3, 4, 5]
squared = list(map(lambda x: x ** 2, numbers))
print(squared)  # [1, 4, 9, 16, 25]

even_numbers = list(filter(lambda x: x % 2 == 0, numbers))
print(even_numbers)  # [2, 4]
```

---

## 4. 함수의 응용

### 4.1 함수 안에서 함수 정의 (중첩 함수)
```python
def outer_function(msg):
    def inner_function():
        print("Message:", msg)
    
    inner_function()

outer_function("Hello, World!")  
# Message: Hello, World!
```

### 4.2 클로저(Closure)
클로저는 내부 함수가 바깥 함수의 변수를 기억하는 기능을 의미합니다.
```python
def multiplier(n):
    def inner(x):
        return x * n
    return inner

times3 = multiplier(3)
print(times3(5))  # 15
```

### 4.3 재귀 함수 (Recursion)
```python
def factorial(n):
    if n == 1:
        return 1
    return n * factorial(n - 1)

print(factorial(5))  # 120
```

---

## 5. 함수형 프로그래밍 기법

### 5.1 `map()` 함수
```python
numbers = [1, 2, 3, 4, 5]
squared = list(map(lambda x: x ** 2, numbers))
print(squared)  # [1, 4, 9, 16, 25]
```

### 5.2 `filter()` 함수
```python
even_numbers = list(filter(lambda x: x % 2 == 0, numbers))
print(even_numbers)  # [2, 4]
```

### 5.3 `reduce()` 함수 (from functools)
```python
from functools import reduce
product = reduce(lambda x, y: x * y, [1, 2, 3, 4, 5])
print(product)  # 120
```
