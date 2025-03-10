# Python 테스트와 디버깅 정리

Python에서 코드를 안정적으로 실행하고, 오류를 수정하기 위해 **테스트(Test)** 와 **디버깅(Debugging)** 이 필요합니다.  
여기서는 Python에서 자주 사용하는 테스트 방법과 디버깅 기법을 정리합니다.

---

## 1. 디버깅(Debugging)

디버깅은 코드의 오류를 찾고 수정하는 과정입니다. Python에서 디버깅하는 주요 방법은 다음과 같습니다.

### 1.1 `print()`를 이용한 디버깅
가장 기본적인 방법으로, 특정 변수를 출력하여 확인하는 방식입니다.

```python
def add(a, b):
    result = a + b
    print(f"Debug: a={a}, b={b}, result={result}")
    return result

add(3, 5)
```

### 1.2 `assert` 문을 이용한 간단한 오류 검출
`assert` 문은 특정 조건이 참(`True`)인지 확인하며, 거짓이면 프로그램이 중단됩니다.

```python
def divide(a, b):
    assert b != 0, "0으로 나눌 수 없습니다!"
    return a / b

print(divide(10, 2))  # 정상 실행
# print(divide(10, 0))  # AssertionError 발생
```

### 1.3 `try-except`를 이용한 예외 처리
예외 발생 시 프로그램이 멈추지 않도록 처리하는 방법입니다.

```python
try:
    x = 10 / 0
except ZeroDivisionError as e:
    print("예외 발생:", e)
```

### 1.4 `logging` 모듈을 이용한 디버깅
`print()` 대신 `logging` 모듈을 사용하면 더 체계적인 디버깅이 가능합니다.

```python
import logging

logging.basicConfig(level=logging.DEBUG)

def add(a, b):
    logging.debug(f"add() called with {a}, {b}")
    return a + b

add(3, 5)
```

---

## 2. Python 디버거 (`pdb`)

Python의 내장 디버거 **pdb**를 사용하면 코드 실행을 단계별로 확인할 수 있습니다.

### 2.1 `pdb` 사용 예제
```python
import pdb

def add(a, b):
    pdb.set_trace()  # 디버거 중단 지점 설정
    result = a + b
    return result

add(3, 5)
```
> 실행 후 `c`(계속), `n`(다음 줄 실행), `p 변수명`(변수 출력) 등의 명령어를 사용할 수 있습니다.

### 2.2 `python -m pdb script.py` 명령어 사용
터미널에서 `python -m pdb script.py` 명령어를 실행하여 코드 디버깅을 수행할 수도 있습니다.

---

## 3. 단위 테스트(Unit Test)

단위 테스트는 코드의 특정 기능이 올바르게 작동하는지 확인하는 테스트입니다.  
Python에서는 **`unittest` 모듈**을 사용하여 단위 테스트를 수행할 수 있습니다.

### 3.1 기본적인 `unittest` 예제
```python
import unittest

def add(a, b):
    return a + b

class TestMathFunctions(unittest.TestCase):
    def test_add(self):
        self.assertEqual(add(3, 5), 8)
        self.assertEqual(add(-1, 1), 0)

if __name__ == "__main__":
    unittest.main()
```

### 3.2 `setUp()`과 `tearDown()` 활용
테스트 실행 전과 후에 필요한 작업을 수행할 수 있습니다.

```python
import unittest

class TestExample(unittest.TestCase):
    def setUp(self):
        self.num1 = 10
        self.num2 = 5

    def test_add(self):
        self.assertEqual(self.num1 + self.num2, 15)

    def tearDown(self):
        print("테스트 종료")

if __name__ == "__main__":
    unittest.main()
```

---

## 4. `pytest`를 이용한 테스트

`pytest`는 `unittest`보다 간단하고 강력한 기능을 제공합니다.  
설치:  
```
pip install pytest
```

### 4.1 `pytest` 기본 예제
```python
def add(a, b):
    return a + b

def test_add():
    assert add(3, 5) == 8
    assert add(-1, 1) == 0
```
> `pytest` 실행:  
> ```
> pytest test_script.py
> ```

---

## 5. 모의 객체(Mock) 테스트 (`unittest.mock`)

외부 API나 DB와 연동된 코드를 테스트할 때 **Mocking**을 활용할 수 있습니다.

### 5.1 `unittest.mock`을 이용한 테스트
```python
from unittest.mock import MagicMock

class Database:
    def get_data(self):
        pass

db = Database()
db.get_data = MagicMock(return_value={"id": 1, "name": "Alice"})

print(db.get_data())  # {'id': 1, 'name': 'Alice'}
```

---

## 6. 코드 커버리지 측정 (`coverage`)

코드 테스트가 얼마나 수행되었는지 확인하려면 **`coverage` 모듈**을 사용합니다.

설치:
```
pip install coverage
```

사용:
```
coverage run -m unittest test_script.py
coverage report
coverage html  # HTML 리포트 생성
```

---

## 7. 코드 스타일 검사 (`flake8`, `black`)

코드 스타일을 검사하고 자동으로 정리할 수 있습니다.

### 7.1 `flake8` - 코드 스타일 검사
설치:
```
pip install flake8
```
사용:
```
flake8 script.py
```

### 7.2 `black` - 코드 자동 정리
설치:
```
pip install black
```
사용:
```
black script.py
```

---

## 8. 정리

| 기법 | 설명 |
|------|------|
| `print()` | 변수 값을 출력하여 디버깅 |
| `assert` | 특정 조건이 만족하는지 확인 |
| `try-except` | 예외 발생 시 처리 |
| `logging` | 체계적인 디버깅 |
| `pdb` | Python 내장 디버거 |
| `unittest` | 기본적인 단위 테스트 |
| `pytest` | 간결한 테스트 프레임워크 |
| `mock` | 외부 API 또는 DB 연동 테스트 |
| `coverage` | 테스트 코드 커버리지 측정 |
| `flake8` | 코드 스타일 검사 |
| `black` | 코드 자동 포맷팅 |
