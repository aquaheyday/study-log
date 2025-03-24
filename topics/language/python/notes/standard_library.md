# Python 표준 라이브러리 정리

Python 표준 라이브러리는 다양한 기능을 제공하는 내장 모듈로, 추가 설치 없이 사용할 수 있습니다.

---

## 1. `math` - 수학 관련 모듈
수학 연산을 수행하는 모듈입니다.

```python
import math

print(math.sqrt(16))  # 4.0 (제곱근)
print(math.factorial(5))  # 120 (팩토리얼)
print(math.pi)  # 3.141592653589793 (파이 값)
print(math.sin(math.radians(90)))  # 1.0 (90도의 사인 값)
```

---

## 2. `random` - 난수 생성 모듈
무작위 숫자를 생성하는 모듈입니다.

```python
import random

print(random.random())  # 0.0 ~ 1.0 사이의 난수
print(random.randint(1, 10))  # 1~10 사이의 랜덤 정수
print(random.choice(["A", "B", "C"]))  # 리스트에서 무작위 선택
```

---

## 3. `datetime` - 날짜 및 시간 처리
날짜 및 시간을 다루는 모듈입니다.

```python
import datetime

now = datetime.datetime.now()
print(now)  # 현재 날짜와 시간 출력
print(now.strftime("%Y-%m-%d %H:%M:%S"))  # 포맷 변경
```

---

## 4. `time` - 시간 관련 함수
시간과 관련된 기능을 제공합니다.

```python
import time

print("3초 동안 대기...")
time.sleep(3)  # 3초 동안 대기
print("끝!")
```

---

## 5. `os` - 운영체제 관련 기능
운영체제의 기능을 다룰 수 있는 모듈입니다.

```python
import os

print(os.name)  # 현재 OS 종류
print(os.getcwd())  # 현재 작업 디렉토리 출력
os.mkdir("test_dir")  # 새 디렉토리 생성
os.rmdir("test_dir")  # 디렉토리 삭제
```

---

## 6. `sys` - 시스템 관련 정보
Python 인터프리터와 관련된 정보를 제공하는 모듈입니다.

```python
import sys

print(sys.version)  # Python 버전 정보 출력
print(sys.argv)  # 실행 시 전달된 인수 목록
sys.exit()  # 프로그램 종료
```

---

## 7. `collections` - 자료구조 모듈
고급 자료구조를 제공하는 모듈입니다.

```python
from collections import deque, Counter

# deque: 양방향 큐
dq = deque([1, 2, 3])
dq.append(4)
dq.appendleft(0)
print(dq)  # deque([0, 1, 2, 3, 4])

# Counter: 리스트 요소 개수 세기
cnt = Counter(["apple", "banana", "apple", "orange", "banana"])
print(cnt)  # Counter({'apple': 2, 'banana': 2, 'orange': 1})
```

---

## 8. `itertools` - 반복자 관련 기능
반복 작업을 쉽게 처리할 수 있도록 돕는 모듈입니다.

```python
import itertools

# 순열 (permutations)
perms = itertools.permutations([1, 2, 3])
print(list(perms))  # [(1,2,3), (1,3,2), (2,1,3), ...]

# 조합 (combinations)
combs = itertools.combinations([1, 2, 3], 2)
print(list(combs))  # [(1,2), (1,3), (2,3)]
```

---

## 9. `functools` - 함수 관련 기능
함수를 다룰 때 유용한 기능을 제공하는 모듈입니다.

```python
import functools

# reduce: 리스트 요소 누적 합
result = functools.reduce(lambda x, y: x + y, [1, 2, 3, 4])
print(result)  # 10
```

---

## 10. `operator` - 연산자 관련 기능
연산을 더 간결하게 수행할 수 있도록 돕는 모듈입니다.

```python
import operator

a = (1, 2)
b = (3, 4)
print(operator.add(a, b))  # (1, 2, 3, 4)
print(operator.mul(3, 4))  # 12
```

---

## 11. `re` - 정규 표현식 (Regular Expression)
문자열에서 특정 패턴을 찾거나 변환할 때 사용됩니다.

```python
import re

text = "My phone number is 010-1234-5678"
match = re.search(r"\d{3}-\d{4}-\d{4}", text)
print(match.group())  # 010-1234-5678
```

---

## 12. `json` - JSON 데이터 처리
JSON 형식의 데이터를 다루는 모듈입니다.

```python
import json

data = {"name": "Alice", "age": 25, "city": "Seoul"}
json_data = json.dumps(data)  # Python -> JSON 문자열
print(json_data)

parsed_data = json.loads(json_data)  # JSON 문자열 -> Python
print(parsed_data["name"])  # Alice
```

---

## 13. `csv` - CSV 파일 처리
CSV(Comma-Separated Values) 파일을 읽고 쓸 때 사용됩니다.

```python
import csv

# CSV 파일 쓰기
with open("data.csv", "w", newline="") as file:
    writer = csv.writer(file)
    writer.writerow(["Name", "Age", "City"])
    writer.writerow(["Alice", 25, "Seoul"])

# CSV 파일 읽기
with open("data.csv", "r") as file:
    reader = csv.reader(file)
    for row in reader:
        print(row)
```

---

## 14. `shutil` - 파일 및 디렉토리 관리
파일 및 디렉토리를 복사, 이동, 삭제하는 기능을 제공합니다.

```python
import shutil

shutil.copy("data.csv", "backup.csv")  # 파일 복사
shutil.move("backup.csv", "archive.csv")  # 파일 이동
shutil.rmtree("test_dir")  # 디렉토리 삭제
```

---

## 15. `hashlib` - 해싱(Hash) 함수
SHA256, MD5 등의 해시 알고리즘을 지원합니다.

```python
import hashlib

text = "Hello, Python!"
hashed = hashlib.sha256(text.encode()).hexdigest()
print(hashed)  # SHA-256 해시 출력
```
