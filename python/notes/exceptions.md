# Python 예외 처리 (Exception Handling) 정리

Python에서 예외(Exception)는 실행 중 발생할 수 있는 오류를 의미합니다. 예외 처리를 통해 프로그램이 중단되지 않도록 할 수 있습니다.

---

## 1. 기본 예외 처리 (`try-except`)

Python에서는 `try-except` 블록을 사용하여 예외를 처리할 수 있습니다.

```python
try:
    x = 10 / 0  # 0으로 나누기 예외 발생
except ZeroDivisionError:
    print("0으로 나눌 수 없습니다!")
```

---

## 2. 여러 개의 예외 처리

여러 개의 예외를 한꺼번에 처리할 수도 있습니다.

```python
try:
    num = int("abc")  # ValueError 발생
    result = 10 / 0   # ZeroDivisionError 발생
except ZeroDivisionError:
    print("0으로 나눌 수 없습니다!")
except ValueError:
    print("숫자로 변환할 수 없습니다!")
```

또는 한 번에 여러 예외를 처리할 수도 있습니다.

```python
try:
    num = int("abc")
except (ValueError, TypeError):
    print("잘못된 입력입니다.")
```

---

## 3. 예외 객체 받기 (`as` 사용)

예외 정보를 얻기 위해 `as` 키워드를 사용할 수 있습니다.

```python
try:
    x = 10 / 0
except ZeroDivisionError as e:
    print("예외 발생:", e)
```

---

## 4. `finally` 블록 (항상 실행)

`finally` 블록은 예외 발생 여부와 관계없이 실행됩니다.

```python
try:
    file = open("test.txt", "r")
    content = file.read()
except FileNotFoundError:
    print("파일을 찾을 수 없습니다.")
finally:
    print("이 문장은 항상 실행됩니다.")
```

---

## 5. `else` 블록 (예외가 없을 때 실행)

`else` 블록은 예외가 발생하지 않을 경우 실행됩니다.

```python
try:
    num = int("10")
except ValueError:
    print("변환 오류 발생")
else:
    print("예외 없이 성공적으로 실행됨:", num)
```

---

## 6. 사용자 정의 예외

Python에서는 `Exception` 클래스를 상속받아 사용자 정의 예외를 만들 수 있습니다.

```python
class CustomError(Exception):
    def __init__(self, message):
        super().__init__(message)

try:
    raise CustomError("사용자 정의 예외 발생!")
except CustomError as e:
    print("예외 처리됨:", e)
```

---

## 7. 예외 발생시키기 (`raise` 사용)

`raise` 키워드를 사용하여 강제로 예외를 발생시킬 수 있습니다.

```python
def check_age(age):
    if age < 18:
        raise ValueError("미성년자는 접근할 수 없습니다.")
    return "접속 허용"

try:
    print(check_age(15))
except ValueError as e:
    print("예외 발생:", e)
```
