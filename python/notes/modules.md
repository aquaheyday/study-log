# 🔲 Python 모듈과 패키지

Python에서 **모듈(Module)** 과 **패키지(Package)** 를 활용하면 코드를 효율적으로 관리하고 재사용할 수 있습니다.  
이 문서에서는 모듈과 패키지의 개념, 사용법, 표준 라이브러리 등을 정리합니다.

---

## 1️⃣ 모듈(Module) 이란?

- 모듈(Module)은 **여러 기능을 하나의 파일에 정의한 Python 코드 파일**입니다.
- `.py` 확장자를 가지며, 다른 Python 파일에서 가져와 사용할 수 있습니다.
- 모듈을 사용하면 코드의 **재사용성**과 **가독성**을 높일 수 있습니다.

### 1) 모듈 가져오기 (`import`)

Python에서 모듈을 가져오는 방법은 다음과 같습니다.

```python
import math  # math 모듈 가져오기

print(math.sqrt(16))  # 4.0 (제곱근)
print(math.pi)        # 3.141592653589793
```

✔ `import 모듈이름`을 사용하여 모듈을 가져올 수 있음  
✔ 모듈의 기능은 `모듈이름.함수이름` 형태로 호출  

---

### 2) 모듈에서 특정 함수만 가져오기 (`from import`)

```python
from math import sqrt, pi  # sqrt(), pi만 가져오기

print(sqrt(25))  # 5.0
print(pi)        # 3.141592653589793
```

✔ `from 모듈이름 import 함수이름`을 사용하면 `모듈이름.` 없이 직접 사용 가능  

---

### 3) 모듈의 모든 기능 가져오기 (`from import *`)

```python
from math import *

print(sin(0))   # 0.0
print(cos(0))   # 1.0
```

✔ `*`를 사용하면 모듈 내 모든 함수와 변수를 가져올 수 있음  
✔ 하지만 **이 방식은 코드의 가독성을 떨어뜨릴 수 있으므로 주의**  

---

### 4) 모듈에 별칭(Alias) 지정 (`as`)

```python
import math as m  # math 모듈을 'm'으로 별칭 지정

print(m.sqrt(9))  # 3.0
```

✔ 모듈 이름이 길 경우, `as` 키워드를 사용하여 별칭을 줄 수 있음  

---

## 2️⃣ 사용자 정의 모듈 만들기

사용자만의 모듈을 직접 만들 수 있습니다.

### 1) 사용자 모듈 생성 (`my_module.py`)
```python
# my_module.py
def greet(name):
    return f"Hello, {name}!"

def add(a, b):
    return a + b
```

### 2) 사용자 모듈 불러오기 (`main.py`)
```python
import my_module  # my_module.py 가져오기

print(my_module.greet("Alice"))  # Hello, Alice!
print(my_module.add(3, 5))       # 8
```

✔ 같은 디렉터리에 `my_module.py`가 있어야 정상적으로 동작  
✔ 모듈의 함수는 `모듈이름.함수이름` 형태로 호출  

---

## 3️⃣ 패키지(Package)란?

- **패키지(Package)** 는 **여러 모듈을 포함하는 디렉터리(폴더)** 입니다.
- 패키지를 사용하면 **관련된 모듈을 논리적으로 그룹화**할 수 있습니다.
- 패키지 디렉터리에는 **`__init__.py` 파일**이 있어야 함 (Python 3.3 이후 생략 가능)

---

### 1) 패키지 구조 예시

```
my_package/        # 패키지 디렉터리
    ├── __init__.py  # 패키지를 인식하는 파일 (빈 파일 가능)
    ├── module1.py   # 첫 번째 모듈
    ├── module2.py   # 두 번째 모듈
```

---

### 2) 패키지 모듈 사용하기
```python
import my_package.module1
import my_package.module2
```

✔ 패키지 내부 모듈은 `패키지명.모듈명` 형태로 불러옴  

---

### 3) `__init__.py` 파일의 역할

- `__init__.py` 파일이 있는 디렉터리는 **패키지로 인식**됨
- `__init__.py` 파일에서 패키지를 초기화할 수 있음

```python
# my_package/__init__.py
print("my_package가 로드되었습니다.")
```

✔ 패키지를 가져올 때 실행됨  

---

### 4) 패키지에서 특정 함수 가져오기
```python
from my_package.module1 import some_function
```

✔ 패키지 내 특정 모듈의 함수만 가져올 수 있음  

---

## 4️⃣ 표준 라이브러리 (Python Standard Library)

Python에는 기본적으로 제공되는 **표준 라이브러리(Standard Library)** 가 있습니다.

| 모듈 | 설명 | 예제 |
|------|----------------|--------------------------------|
| `math` | 수학 연산 지원 | `math.sqrt(25) → 5.0` |
| `random` | 난수 생성 | `random.randint(1, 10)` |
| `datetime` | 날짜 및 시간 처리 | `datetime.datetime.now()` |
| `os` | 운영체제 기능 | `os.getcwd()` |
| `sys` | 시스템 관련 정보 | `sys.version` |
| `re` | 정규 표현식 처리 | `re.match(r'\d+', '123abc')` |

✔ **표준 라이브러리는 별도의 설치 없이 바로 사용 가능**  

---

## 5️⃣ 외부 라이브러리 (Third-party Libraries)

Python에서는 **외부 라이브러리**를 설치하여 확장 기능을 사용할 수 있습니다.

### 외부 라이브러리 설치 (`pip`)
```sh
pip install requests  # requests 라이브러리 설치
```

✔ 설치된 라이브러리는 `import`하여 사용할 수 있음
```python
import requests
response = requests.get("https://www.google.com")
print(response.status_code)  # 200
```

✔ 설치된 라이브러리 목록 확인
```sh
pip list
```

✔ 특정 라이브러리 삭제
```sh
pip uninstall requests
```

---

## 🎯 정리

✔ **모듈(Module)**: `.py` 파일로 구성된 코드 묶음 (`import`로 가져옴)  
✔ **패키지(Package)**: 여러 모듈을 포함하는 디렉터리 (서브 모듈 관리 가능)  
✔ **표준 라이브러리**: Python 기본 내장 모듈 (`math`, `random`, `os` 등)  
✔ **외부 라이브러리**: `pip`으로 설치하여 확장 기능 사용 가능  
