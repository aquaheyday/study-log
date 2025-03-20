# 🔲 Python 예외 처리 (Exception Handling)

Python에서 **예외(Exception)** 는 프로그램 실행 중 발생할 수 있는 오류입니다.  
**예외 처리(Exception Handling)** 를 사용하면 프로그램이 예외 발생 시 중단되지 않고 정상적으로 동작하도록 만들 수 있습니다.

---

## 1️⃣ 예외(Exception)란?

- 예외는 **코드 실행 중 발생하는 오류**입니다.
- 예외가 발생하면 프로그램이 즉시 중단되므로, 이를 **처리**하는 것이 중요합니다.
- 예외는 `try-except` 문을 사용하여 처리할 수 있습니다.

#### 예외 발생 예제
```python
print(10 / 0)  # ZeroDivisionError 발생 (0으로 나눌 수 없음)
```

#### 예외 발생 시 프로그램이 중단됨
```
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
ZeroDivisionError: division by zero
```

---

## 2️⃣ 기본 예외 처리 (`try-except`)

```python
try:
    x = 10 / 0  # 예외 발생 가능 코드
except ZeroDivisionError:
    print("0으로 나눌 수 없습니다.")  # 예외 처리
```

✔ `try` 블록 안에서 오류가 발생하면, `except` 블록이 실행됨  
✔ 예외가 발생하지 않으면 `except` 블록은 실행되지 않음  

---

## 3️⃣ 복수의 예외 처리

### 1) 여러 개의 예외 처리 (`except` 여러 개)
```python
try:
    num = int(input("숫자를 입력하세요: "))
    result = 10 / num
except ZeroDivisionError:
    print("0으로 나눌 수 없습니다.")
except ValueError:
    print("숫자가 아닌 값을 입력했습니다.")
```

✔ 여러 개의 예외를 개별적으로 처리 가능  

---

### 2) 하나의 `except`에서 여러 예외 처리 (`except (예외1, 예외2)`)
```python
try:
    num = int(input("숫자를 입력하세요: "))
    result = 10 / num
except (ZeroDivisionError, ValueError):
    print("잘못된 입력입니다.")
```

✔ `except (예외1, 예외2):` 형태로 여러 예외를 한 번에 처리 가능  

---

## 4️⃣ 모든 예외 처리 (`Exception` 사용)

```python
try:
    x = int("abc")  # ValueError 발생
except Exception as e:
    print(f"예외 발생: {e}")
```

✔ `Exception`을 사용하면 모든 예외를 포괄적으로 처리할 수 있음  
✔ `as e`를 사용하면 예외 메시지를 출력할 수 있음  

---

## 5️⃣ `finally` 블록 (항상 실행되는 코드)

```python
try:
    file = open("test.txt", "r")
    data = file.read()
except FileNotFoundError:
    print("파일을 찾을 수 없습니다.")
finally:
    print("이 코드는 항상 실행됩니다.")
```

✔ `finally` 블록은 예외 발생 여부와 상관없이 항상 실행됨  
✔ 파일 닫기, 리소스 정리 등에 유용하게 사용  

---

## 6️⃣ `else` 블록 (예외가 발생하지 않았을 때 실행)

```python
try:
    num = int(input("숫자를 입력하세요: "))
    result = 10 / num
except ZeroDivisionError:
    print("0으로 나눌 수 없습니다.")
else:
    print(f"결과: {result}")  # 예외가 발생하지 않았을 때 실행
```

✔ `else` 블록은 `try` 블록이 **정상 실행된 경우에만** 실행됨  

---

## 7️⃣ 사용자 정의 예외 (`raise` 사용)

### 1) 직접 예외 발생 (`raise`)
```python
def check_age(age):
    if age < 18:
        raise ValueError("18세 이상만 가능합니다.")
    print("입장 가능합니다.")

try:
    check_age(15)
except ValueError as e:
    print(f"예외 발생: {e}")
```

✔ `raise`를 사용하면 특정 조건에서 강제로 예외 발생 가능  

---

### 2) 사용자 정의 예외 클래스
```python
class NegativeNumberError(Exception):
    """음수 입력 시 발생하는 사용자 정의 예외"""
    pass

def check_number(num):
    if num < 0:
        raise NegativeNumberError("음수는 입력할 수 없습니다.")
    print("올바른 숫자입니다.")

try:
    check_number(-5)
except NegativeNumberError as e:
    print(f"사용자 정의 예외 발생: {e}")
```

✔ 사용자 정의 예외는 `Exception`을 상속받아 정의 가능  

---

## 8️⃣ 주요 예외 종류

| 예외 클래스 | 설명 |
|------------|--------------------------------|
| `ZeroDivisionError` | 0으로 나누었을 때 발생 |
| `ValueError` | 잘못된 값이 입력되었을 때 발생 |
| `TypeError` | 연산 또는 함수에 잘못된 타입이 사용되었을 때 발생 |
| `IndexError` | 리스트나 튜플에서 잘못된 인덱스를 참조했을 때 발생 |
| `KeyError` | 딕셔너리에 존재하지 않는 키를 참조했을 때 발생 |
| `FileNotFoundError` | 존재하지 않는 파일을 열려고 할 때 발생 |
| `ImportError` | 모듈을 찾을 수 없을 때 발생 |
| `NameError` | 정의되지 않은 변수를 참조했을 때 발생 |

---

## 🎯 정리

✔ **`try-except`** 문을 사용하여 예외 처리  
✔ **여러 개의 예외**를 개별적으로 처리하거나, 한 번에 처리 가능  
✔ **`finally` 블록** → 항상 실행되는 코드  
✔ **`else` 블록** → 예외 없이 실행될 때만 실행  
✔ **`raise`를 사용하여 직접 예외 발생 가능**  
✔ **사용자 정의 예외**를 만들어 특정 조건에서 예외 처리 가능  
