# 📌 Python 변수와 자료형

Python에서 변수를 선언하고 다양한 자료형을 다루는 방법을 정리합니다.

---

## 1️⃣ 변수 (Variable)

변수는 데이터를 저장하는 **이름이 있는 메모리 공간**입니다.

### 변수 선언 및 할당
```python
# 변수 선언 및 값 할당
name = "Python"
version = 3.11
pi = 3.14159
is_active = True
```

✔ Python에서는 변수의 자료형을 **명시적으로 선언하지 않아도 됨** (동적 타이핑)  
✔ 변수 이름은 **알파벳, 숫자, `_`** 를 포함할 수 있으며 숫자로 시작할 수 없음  
✔ 대소문자를 구분 (`age`와 `Age`는 다른 변수)

---

## 2️⃣ 기본 자료형 (Primitive Data Types)

| 자료형 | 설명 | 예제 |
|--------|----------|----------------|
| `int` | 정수형 | `num = 42` |
| `float` | 실수형 | `pi = 3.14` |
| `str` | 문자열 | `text = "Hello"` |
| `bool` | 논리형 (참/거짓) | `is_active = True` |

자료형 확인 방법:
```python
print(type(42))        # <class 'int'>
print(type(3.14))      # <class 'float'>
print(type("Hello"))   # <class 'str'>
print(type(True))      # <class 'bool'>
```

---

## 3️⃣ 문자열 (String)

Python에서 문자열은 작은따옴표(`'`) 또는 큰따옴표(`"`)로 감쌉니다.

### 문자열 기본 사용법
```python
text = "Hello, Python"
print(text[0])   # 'H' (인덱싱)
print(text[-1])  # 'n' (뒤에서 첫 번째 문자)
print(text[:5])  # 'Hello' (슬라이싱)
print(len(text)) # 문자열 길이
```

### 문자열 연산
```python
a = "Hello"
b = "Python"
print(a + " " + b)  # "Hello Python" (문자열 연결)
print(a * 3)        # "HelloHelloHello" (반복)
```

---

## 4️⃣ 숫자형 (Numbers)

Python에서 정수(`int`)와 실수(`float`)를 다룰 수 있습니다.

### 기본 연산
```python
a = 10
b = 3

print(a + b)  # 덧셈 (13)
print(a - b)  # 뺄셈 (7)
print(a * b)  # 곱셈 (30)
print(a / b)  # 나눗셈 (3.3333...)
print(a // b) # 몫 (3)
print(a % b)  # 나머지 (1)
print(a ** b) # 거듭제곱 (1000)
```

---

## 5️⃣ 논리형 (Boolean)

논리형(`bool`)은 `True` 또는 `False` 값을 가집니다.

```python
is_python_fun = True
is_learning_hard = False

print(is_python_fun and is_learning_hard)  # False
print(is_python_fun or is_learning_hard)   # True
print(not is_python_fun)                   # False
```

`bool()` 함수를 사용하여 값을 논리형으로 변환할 수 있음
```python
print(bool(0))    # False
print(bool(42))   # True
print(bool(""))   # False
print(bool("text"))  # True
```

---

## 6️⃣ 컬렉션 자료형 (Collection Data Types)

Python에서는 여러 개의 값을 저장할 수 있는 컬렉션 자료형을 제공합니다.

| 자료형 | 설명 | 예제 |
|--------|-----------------|----------------|
| `list` | 순서가 있는 변경 가능한 배열 | `[1, 2, 3]` |
| `tuple` | 순서가 있는 변경 불가능한 배열 | `(1, 2, 3)` |
| `set` | 순서가 없고 중복을 허용하지 않는 집합 | `{1, 2, 3}` |
| `dict` | 키-값 쌍으로 이루어진 데이터 구조 | `{"name": "Alice", "age": 25}` |

✔ 자료형 확인:
```python
print(type([1, 2, 3]))  # <class 'list'>
print(type((1, 2, 3)))  # <class 'tuple'>
print(type({1, 2, 3}))  # <class 'set'>
print(type({"name": "Alice"}))  # <class 'dict'>
```

---

## 7️⃣ 자료형 변환 (Type Casting)

Python에서는 자료형을 자유롭게 변환할 수 있습니다.

### 정수, 실수, 문자열 변환
```python
num = 42
pi = 3.14
text = "123"

print(int(pi))     # 3 (실수를 정수로 변환)
print(float(num))  # 42.0 (정수를 실수로 변환)
print(str(num))    # '42' (숫자를 문자열로 변환)
print(int(text))   # 123 (문자열을 정수로 변환)
```

### 리스트, 튜플, 집합 변환
```python
my_list = [1, 2, 3]
my_tuple = tuple(my_list)  # 리스트 → 튜플
my_set = set(my_list)      # 리스트 → 집합
my_list2 = list(my_tuple)  # 튜플 → 리스트

print(my_tuple)  # (1, 2, 3)
print(my_set)    # {1, 2, 3}
print(my_list2)  # [1, 2, 3]
```

⚠️ **주의:** 문자열을 숫자로 변환할 때 숫자가 아닌 문자가 포함되어 있으면 오류 발생
```python
print(int("123a"))  # ValueError 발생!
```

---

## 🎯 정리

✔ **변수**는 값을 저장하는 메모리 공간  
✔ **자료형**에는 `int`, `float`, `str`, `bool`, `list`, `tuple`, `set`, `dict` 등이 있음  
✔ **문자열**, **숫자형**, **논리형**, **컬렉션 자료형**을 자유롭게 변환 가능  
✔ `print(type(value))`을 사용하여 자료형 확인 가능  
