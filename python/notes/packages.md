# Python 패키지와 모듈 정리

Python에서 **모듈(Module)** 과 **패키지(Package)** 는 코드의 재사용성을 높이고, 유지보수를 쉽게 하기 위해 사용됩니다.

---

## 1. 모듈(Module)이란?
모듈은 `.py` 확장자를 가진 Python 파일로, 변수, 함수, 클래스 등을 포함할 수 있습니다.  
즉, 하나의 `.py` 파일이 하나의 모듈이 됩니다.

### 1.1 모듈 만들기
`math_operations.py`라는 파일을 생성하여 아래와 같이 정의합니다.

```python
# math_operations.py
def add(a, b):
    return a + b

def subtract(a, b):
    return a - b
```

### 1.2 모듈 불러오기
```python
import math_operations

print(math_operations.add(3, 5))  # 8
print(math_operations.subtract(10, 3))  # 7
```

### 1.3 모듈에서 특정 함수만 가져오기
```python
from math_operations import add

print(add(4, 6))  # 10
```

### 1.4 모듈 별칭(Alias) 사용하기
```python
import math_operations as mo

print(mo.add(2, 3))  # 5
```

### 1.5 모든 함수 가져오기 (`*` 사용)
```python
from math_operations import *

print(add(5, 7))  # 12
print(subtract(10, 4))  # 6
```

---

## 2. 패키지(Package)란?
패키지는 여러 개의 모듈을 포함하는 **디렉토리(폴더)** 입니다.  
패키지는 **`__init__.py`** 파일을 포함해야 하며, 이를 통해 Python이 해당 디렉토리를 패키지로 인식합니다.

### 2.1 패키지 구조
```
mypackage/        # 패키지 폴더
│── __init__.py   # 패키지 초기화 파일
│── math_operations.py  # 모듈 1
│── string_operations.py  # 모듈 2
```

### 2.2 패키지 생성
각 파일의 내용을 작성합니다.

#### `mypackage/__init__.py`
```python
# 빈 파일로 두거나 패키지 정보를 정의할 수 있음
__all__ = ["math_operations", "string_operations"]
```

#### `mypackage/math_operations.py`
```python
def add(a, b):
    return a + b
```

#### `mypackage/string_operations.py`
```python
def to_upper(text):
    return text.upper()
```

---

## 3. 패키지 사용하기

### 3.1 패키지 전체 가져오기
```python
import mypackage.math_operations
import mypackage.string_operations

print(mypackage.math_operations.add(2, 3))  # 5
print(mypackage.string_operations.to_upper("hello"))  # HELLO
```

### 3.2 특정 모듈만 가져오기
```python
from mypackage import math_operations

print(math_operations.add(5, 10))  # 15
```

### 3.3 특정 함수만 가져오기
```python
from mypackage.math_operations import add
from mypackage.string_operations import to_upper

print(add(3, 7))  # 10
print(to_upper("python"))  # PYTHON
```

---

## 4. `__init__.py`의 역할
패키지의 `__init__.py` 파일을 사용하여 패키지를 더욱 편리하게 관리할 수 있습니다.

예를 들어, `mypackage/__init__.py`에 다음 내용을 추가하면 패키지에서 직접 함수를 가져올 수 있습니다.

```python
from .math_operations import add
from .string_operations import to_upper
```

이제 아래와 같이 패키지를 사용할 수 있습니다.

```python
import mypackage

print(mypackage.add(3, 5))  # 8
print(mypackage.to_upper("hello"))  # HELLO
```

---

## 5. 모듈의 실행 (`__name__ == "__main__"`)

Python 파일을 직접 실행할 때와 모듈로 불러올 때의 동작을 다르게 설정할 수 있습니다.

```python
# sample.py
def main():
    print("이 파일은 직접 실행되었습니다.")

if __name__ == "__main__":
    main()
```

### 5.1 직접 실행한 경우
```
python sample.py
```
출력:
```
이 파일은 직접 실행되었습니다.
```

### 5.2 모듈로 불러온 경우
```python
import sample
```
출력 없음 (`main()`이 실행되지 않음).

---

## 6. 표준 라이브러리 모듈 활용

Python은 다양한 내장 모듈을 제공합니다.

### 6.1 `os` 모듈 (운영체제 관련 기능)
```python
import os

print(os.getcwd())  # 현재 작업 디렉토리
os.mkdir("new_folder")  # 새 폴더 생성
```

### 6.2 `sys` 모듈 (시스템 관련 기능)
```python
import sys

print(sys.version)  # Python 버전 출력
print(sys.path)  # 모듈 검색 경로 출력
```

### 6.3 `math` 모듈 (수학 함수)
```python
import math

print(math.sqrt(16))  # 4.0 (제곱근)
print(math.pi)  # 3.141592653589793
```

---

## 7. 외부 패키지 사용 (`pip`)

### 7.1 패키지 설치
```
pip install requests
```

### 7.2 설치된 패키지 사용
```python
import requests

response = requests.get("https://api.github.com")
print(response.status_code)  # 200
```

### 7.3 설치된 패키지 목록 확인
```
pip list
```

---

## 8. 가상 환경 (Virtual Environment)

가상 환경은 프로젝트별로 Python 패키지를 독립적으로 관리할 수 있도록 도와줍니다.

### 8.1 가상 환경 생성
```
python -m venv myenv
```

### 8.2 가상 환경 활성화
#### Windows
```
myenv\Scripts\activate
```
#### macOS/Linux
```
source myenv/bin/activate
```

### 8.3 가상 환경 비활성화
```
deactivate
```

---

## 9. 정리

| 개념 | 설명 |
|------|------|
| **모듈** | 하나의 `.py` 파일, 함수/클래스 포함 |
| **패키지** | 여러 모듈을 포함하는 디렉토리 |
| `import` | 모듈 불러오기 |
| `from ... import` | 특정 함수/클래스만 가져오기 |
| `__init__.py` | 패키지를 Python이 인식하도록 함 |
| `pip` | 외부 패키지 설치 및 관리 |
| 가상 환경 | 프로젝트별 패키지 독립 관리 |
