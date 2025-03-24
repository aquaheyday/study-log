# 🛠️ Python 단위 테스트 (Unit Testing)

단위 테스트(Unit Testing)는 **코드의 개별 함수나 모듈이 정상적으로 동작하는지 검증하는 과정**입니다.  
Python에서는 **`unittest`**, **`pytest`**, **`doctest`** 등의 라이브러리를 사용하여 단위 테스트를 수행할 수 있습니다.

---

## 1️⃣ 단위 테스트란?

- **프로그램의 개별 구성 요소(함수, 클래스, 메서드)를 독립적으로 테스트하는 방법**
- **버그를 조기에 발견하여 유지보수 비용을 절감**
- **CI/CD(지속적 통합 및 배포) 파이프라인에서 자동화 가능**
- **주요 단위 테스트 프레임워크**:
  - `unittest` (Python 표준 라이브러리)
  - `pytest` (간결하고 강력한 기능 제공)
  - `doctest` (문서 문자열에서 테스트)

---

## 2️⃣ 단위 테스트 작성 방법 (`unittest` 활용)

### 1) `unittest` 모듈 가져오기
```python
import unittest
```
✔ `unittest` 는 Python 내장 라이브러리로 별도 설치 불필요  

---

### 2) 기본적인 테스트 코드 작성
```python
import unittest

def add(x, y):
    return x + y

class TestMathFunctions(unittest.TestCase):
    
    def test_add(self):
        self.assertEqual(add(2, 3), 5)  # 2 + 3 = 5 검증
        self.assertEqual(add(-1, 1), 0) # -1 + 1 = 0 검증

if __name__ == "__main__":
    unittest.main()
```
✔ `unittest.TestCase` 상속받아 테스트 클래스 생성  
✔ `assertEqual(a, b)` → `a == b` 인지 검증  

---

### 3) 다양한 단언(Assertion) 메서드

| 메서드 | 설명 |
|--------|------|
| `assertEqual(a, b)` | a == b |
| `assertNotEqual(a, b)` | a != b |
| `assertTrue(x)` | x가 `True`인지 |
| `assertFalse(x)` | x가 `False`인지 |
| `assertIs(a, b)` | a와 b가 동일 객체인지 (`is`) |
| `assertIsNone(x)` | x가 `None`인지 |
| `assertRaises(Exception, func, *args)` | 특정 예외 발생 검증 |

---

## 3️⃣ 테스트 실행하기

#### 1.`python` 명령어로 실행
```sh
python test_sample.py
```

#### 2. `unittest` 모듈 직접 실행
```sh
python -m unittest test_sample.py
```

#### ✅ 성공 시
```
.
----------------------------------------------------------------------
Ran 1 test in 0.001s

OK
```

#### ❌ 실패 시
```
F
======================================================================
FAIL: test_add (__main__.TestMathFunctions)
----------------------------------------------------------------------
Traceback (most recent call last):
  File "test_sample.py", line 6, in test_add
    self.assertEqual(add(2, 3), 6)  # 잘못된 예상 결과
AssertionError: 5 != 6

----------------------------------------------------------------------
Ran 1 test in 0.001s

FAILED (failures=1)
```

---

## 4️⃣ 테스트 실행 전후 설정 (`setUp` / `tearDown`)

### 1) 테스트 시작 전/후에 실행할 코드 정의
```python
import unittest

class TestExample(unittest.TestCase):

    def setUp(self):
        """각 테스트 전에 실행"""
        self.data = [1, 2, 3]

    def tearDown(self):
        """각 테스트 후에 실행"""
        self.data.clear()

    def test_length(self):
        self.assertEqual(len(self.data), 3)

    def test_first_element(self):
        self.assertEqual(self.data[0], 1)

if __name__ == "__main__":
    unittest.main()
```
✔ `setUp()` → 테스트 시작 전 실행 (테스트 초기화)  
✔ `tearDown()` → 테스트 종료 후 실행 (리소스 정리)  

---

## 5️⃣ 예외 처리 테스트 (`assertRaises`)

#### 예외 발생 확인
```python
def divide(x, y):
    if y == 0:
        raise ValueError("0으로 나눌 수 없습니다.")
    return x / y

class TestDivision(unittest.TestCase):
    
    def test_divide_by_zero(self):
        with self.assertRaises(ValueError):
            divide(10, 0)

if __name__ == "__main__":
    unittest.main()
```
✔ `assertRaises(ValueError)` → **ValueError 발생 여부 테스트**  

---

## 6️⃣ `mock`을 활용한 테스트 (외부 의존성 제거)

#### `unittest.mock`을 활용한 가짜(Mock) 객체 만들기
```python
from unittest.mock import MagicMock
import unittest

class PaymentGateway:
    def charge(self, amount):
        # 실제 결제 처리 코드 (테스트 X)
        pass

class TestPayment(unittest.TestCase):

    def test_charge(self):
        payment = PaymentGateway()
        payment.charge = MagicMock(return_value="결제 성공")  # Mock 설정

        result = payment.charge(1000)
        payment.charge.assert_called_once_with(1000)  # 특정 인자로 호출 확인
        self.assertEqual(result, "결제 성공")

if __name__ == "__main__":
    unittest.main()
```
✔ `MagicMock()` → 실제 기능 대신 **가짜(Mock) 함수 실행**  
✔ `assert_called_once_with(값)` → 특정 값으로 **정확히 한 번 호출**되었는지 검증  

---

## 7️⃣ `pytest`를 활용한 간편한 단위 테스트

#### 1. `pytest` 설치
```sh
pip install pytest
```

#### 2. `pytest`로 테스트 실행
```sh
pytest test_sample.py
```

#### 3. `pytest`를 활용한 간단한 테스트 코드
```python
def add(x, y):
    return x + y

def test_add():
    assert add(2, 3) == 5
    assert add(-1, 1) == 0
```
✔ `pytest`는 **클래스 없이 간결한 테스트 코드 작성 가능**  

---

## 🎯 정리

✔ **단위 테스트(Unit Test)** → 코드의 개별 단위(함수, 클래스)를 검증  
✔ **`unittest` 모듈 사용** → Python 기본 제공 테스트 프레임워크  
✔ **Assertion 활용** → `assertEqual()`, `assertRaises()` 등 다양한 검증 메서드 지원  
✔ **Mock 객체 사용** → `MagicMock()` 을 활용하여 외부 API 호출 없이 테스트 가능  
✔ **`pytest` 활용 가능** → `unittest` 보다 간결한 코드 작성 지원  
