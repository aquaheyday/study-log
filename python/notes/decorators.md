# 🚀 Python 데코레이터 (Decorator)

Python에서 **데코레이터(Decorator)** 는 **기존 함수의 동작을 변경하거나 확장하는 기능**을 제공합니다.  
이를 통해 **코드를 더 효율적이고 재사용 가능**하게 만들 수 있습니다.

---

## 1️⃣ 데코레이터란?

- **함수를 감싸서 실행 전/후에 추가 기능을 수행하는 함수**입니다.
- **기존 코드 수정 없이 기능을 확장**할 수 있음.
- `@decorator_name` 문법을 사용하여 적용 가능.

---

## 2️⃣ 기본적인 데코레이터 구현

#### 함수형 데코레이터 예제
```python
def my_decorator(func):
    def wrapper():
        print("함수 실행 전")
        func()
        print("함수 실행 후")
    return wrapper

@my_decorator
def hello():
    print("Hello, World!")

hello()
```
✔ `wrapper()` 함수가 `func()`을 감싸서 실행  
✔ `@my_decorator`를 붙이면 `hello()` 호출 시 자동으로 `wrapper()` 실행  

#### 출력
```
함수 실행 전
Hello, World!
함수 실행 후
```

---

## 3️⃣ 데코레이터의 내부 원리 (`functools.wraps` 사용)

Python에서는 **`functools.wraps`** 를 사용하여 **데코레이터가 원래 함수의 정보를 유지**하도록 할 수 있습니다.

#### 예제
```python
import functools

def my_decorator(func):
    @functools.wraps(func)
    def wrapper(*args, **kwargs):
        print(f"함수 {func.__name__} 실행 전")
        result = func(*args, **kwargs)
        print(f"함수 {func.__name__} 실행 후")
        return result
    return wrapper

@my_decorator
def add(a, b):
    return a + b

print(add(3, 5))
```
✔ `@functools.wraps(func)` → 원래 함수 이름, docstring 유지  

#### 출력
```
함수 add 실행 전
함수 add 실행 후
8
```

---

## 4️⃣ 인자를 받는 데코레이터

데코레이터에 인자를 추가하려면 **함수를 한 번 더 감싸는 구조**로 만들어야 합니다.

#### 데코레이터에 인자 전달 예제
```python
import functools

def repeat(n):
    def decorator(func):
        @functools.wraps(func)
        def wrapper(*args, **kwargs):
            for _ in range(n):
                func(*args, **kwargs)
        return wrapper
    return decorator

@repeat(3)  # 함수 실행을 3번 반복
def greet():
    print("Hello!")

greet()
```
✔ `repeat(n)` → **데코레이터에 인자를 전달**  
✔ `@repeat(3)` → `greet()` 함수가 3번 실행됨  

#### 출력
```
Hello!
Hello!
Hello!
```

---

## 5️⃣ 여러 개의 데코레이터 적용

#### 여러 개의 데코레이터 중첩 적용 예제제
```python
def decorator1(func):
    def wrapper(*args, **kwargs):
        print("데코레이터 1 실행")
        return func(*args, **kwargs)
    return wrapper

def decorator2(func):
    def wrapper(*args, **kwargs):
        print("데코레이터 2 실행")
        return func(*args, **kwargs)
    return wrapper

@decorator1
@decorator2
def say_hello():
    print("Hello, Python!")

say_hello()
```
✔ **데코레이터 적용 순서** → `@decorator1`이 **가장 바깥쪽**에서 실행됨.  

#### 출력
```
데코레이터 1 실행
데코레이터 2 실행
Hello, Python!
```

---

## 6️⃣ 클래스 데코레이터 (`__call__` 사용)

Python에서는 **클래스도 데코레이터로 사용할 수 있음** (`__call__` 메서드 사용).

#### 클래스 데코레이터 구현 예제
```python
class Timer:
    def __call__(self, func):
        import time
        def wrapper(*args, **kwargs):
            start = time.time()
            result = func(*args, **kwargs)
            end = time.time()
            print(f"{func.__name__} 실행 시간: {end - start:.4f}초")
            return result
        return wrapper

@Timer()  # 클래스 데코레이터 사용
def slow_function():
    import time
    time.sleep(2)
    print("작업 완료!")

slow_function()
```
✔ `__call__()` → **함수처럼 동작하도록 설정 가능**  

#### 출력
```
작업 완료!
slow_function 실행 시간: 2.0001초
```

---

## 7️⃣ 실전 활용 예제

#### 실행 시간 측정 데코레이터 예제
```python
import time
import functools

def timing(func):
    @functools.wraps(func)
    def wrapper(*args, **kwargs):
        start = time.time()
        result = func(*args, **kwargs)
        end = time.time()
        print(f"{func.__name__} 실행 시간: {end - start:.4f}초")
        return result
    return wrapper

@timing
def slow_function():
    time.sleep(2)
    print("작업 완료!")

slow_function()
```
✔ 특정 함수의 **실행 시간을 자동으로 측정**하는 데 유용  

---

#### 로그인 확인 데코레이터 (Flask 스타일)
```python
import functools

def login_required(func):
    @functools.wraps(func)
    def wrapper(user, *args, **kwargs):
        if not user.get("is_authenticated"):
            print("⛔ 접근 불가: 로그인이 필요합니다.")
            return
        return func(user, *args, **kwargs)
    return wrapper

@login_required
def view_profile(user):
    print(f"👤 사용자 프로필: {user['name']}")

user1 = {"name": "Alice", "is_authenticated": True}
user2 = {"name": "Bob", "is_authenticated": False}

view_profile(user1)  # 정상 실행
view_profile(user2)  # 접근 불가
```
✔ `user.get("is_authenticated")` → **로그인 여부 체크 후 실행**  

#### 출력 결과
```
👤 사용자 프로필: Alice
⛔ 접근 불가: 로그인이 필요합니다.
```

---

## 🎯 정리

✔ **데코레이터(Decorator)** → 함수 실행 전/후에 추가 기능을 적용  
✔ **기본 구조** → `wrapper()` 내부 함수로 원본 함수 감싸기  
✔ **`@functools.wraps(func)`** → 함수 정보를 유지  
✔ **인자 전달 가능** → 데코레이터 함수 내부에 함수 추가  
✔ **클래스 데코레이터** → `__call__()`을 사용하여 동작 가능  
✔ **실전 활용** → 실행 시간 측정, 로그인 체크, API 인증 등에 활용  
