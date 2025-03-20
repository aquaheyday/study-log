# 🔄 Python 파일 입출력 (File I/O)

Python에서는 파일을 읽고 쓰는 기능을 제공하여 데이터를 저장하고 불러올 수 있습니다.  
이 문서에서는 **텍스트 파일, CSV, JSON 파일 처리 방법**을 정리합니다.

---

## 1️⃣ 파일 열기 및 닫기

Python에서 파일을 열고 작업하는 기본 방식은 `open()` 함수를 사용하는 것입니다.

### `open()` 함수 기본 문법
```python
file = open("example.txt", "r")  # 파일 열기
file.close()  # 파일 닫기
```

✔ 파일 작업이 끝나면 `close()`를 호출하여 반드시 닫아야 합니다.

---

## 2️⃣ 파일 모드 (File Mode)

파일을 열 때 사용할 수 있는 **모드(Mode)** 는 다음과 같습니다.

| 모드 | 설명 |
|------|------------------------------|
| `"r"` | 읽기 모드 (기본값, 파일이 존재해야 함) |
| `"w"` | 쓰기 모드 (파일이 없으면 생성, 기존 내용 삭제) |
| `"a"` | 추가 모드 (파일이 없으면 생성, 기존 내용 유지) |
| `"x"` | 새 파일 생성 (이미 존재하면 오류 발생) |
| `"b"` | 바이너리 모드 (`rb`, `wb` 등과 함께 사용) |

✔ `"r+"`, `"w+"`, `"a+"` → 읽기와 쓰기 동시 지원

---

## 3️⃣ 파일 읽기 (Read)

### 1) `read()` - 전체 내용 읽기
```python
with open("example.txt", "r") as file:
    content = file.read()
    print(content)  # 파일 전체 내용 출력
```

✔ `with open()`을 사용하면 `close()`를 호출할 필요 없음  

---

### 2) `readline()` - 한 줄씩 읽기
```python
with open("example.txt", "r") as file:
    line = file.readline()
    print(line)  # 첫 번째 줄 출력
```

✔ `readline()`은 한 줄씩 읽으며, **개행 문자(`\n`) 포함됨**  

---

### 3) `readlines()` - 모든 줄을 리스트로 반환
```python
with open("example.txt", "r") as file:
    lines = file.readlines()
    print(lines)  # ['첫 번째 줄\n', '두 번째 줄\n', '세 번째 줄\n']
```

✔ `readlines()`를 사용하면 각 줄이 리스트의 요소로 저장됨  

---

## 4️⃣ 파일 쓰기 (Write)

### 1) `write()` - 새로운 내용 쓰기
```python
with open("example.txt", "w") as file:
    file.write("Hello, Python!\n")
```

✔ `"w"` 모드는 **기존 내용을 덮어씀**  

---

### 2) `writelines()` - 여러 줄 쓰기
```python
lines = ["첫 번째 줄\n", "두 번째 줄\n", "세 번째 줄\n"]

with open("example.txt", "w") as file:
    file.writelines(lines)
```

✔ 리스트의 문자열을 한 번에 파일에 저장  

---

### 3) `append` 모드 (`"a"`) - 기존 내용 유지하며 추가
```python
with open("example.txt", "a") as file:
    file.write("추가된 내용\n")
```

✔ `"a"` 모드는 **기존 내용을 유지한 채 새로운 내용을 추가**  

---

## 5️⃣ `with` 문 사용 (자동 파일 닫기)

Python에서는 `with open()`을 사용하면 **자동으로 파일이 닫히므로** 더 안전합니다.

```python
with open("example.txt", "r") as file:
    content = file.read()
    print(content)  # 파일 읽기

# 파일이 자동으로 닫힘
```

✔ `with` 문을 사용하면 `close()`를 명시적으로 호출할 필요 없음  

---

## 6️⃣ `rb` - 바이너리 파일 읽기

이미지, 오디오, 동영상 같은 **바이너리 파일**을 읽거나 저장할 때 사용합니다.

```python
with open("image.jpg", "rb") as file:
    binary_data = file.read()
    print(binary_data[:20])  # 처음 20바이트 출력
```

✔ `"rb"` 모드를 사용하면 **텍스트 파일이 아닌 이진 파일**을 다룰 수 있음  

---

## 7️⃣ CSV 파일 처리

Python에서는 `csv` 모듈을 사용하여 CSV 파일을 쉽게 다룰 수 있습니다.

### 1) `csv.reader` - CSV 파일 읽기
```python
import csv

with open("data.csv", "r", encoding="utf-8") as file:
    reader = csv.reader(file)
    for row in reader:
        print(row)  # 한 줄씩 리스트로 출력
```

✔ `csv.reader()`는 각 행을 리스트로 반환  

---

### 2) `csv.writer` - CSV 파일 쓰기
```python
import csv

data = [["이름", "나이"], ["Alice", 25], ["Bob", 30]]

with open("data.csv", "w", newline="", encoding="utf-8") as file:
    writer = csv.writer(file)
    writer.writerows(data)
```

✔ `writerows()`를 사용하여 여러 줄을 한 번에 저장  

---

## 8️⃣ JSON 파일 처리

Python에서는 `json` 모듈을 사용하여 JSON 데이터를 읽고 저장할 수 있습니다.

### 1) `json.dump` - JSON 파일 저장
```python
import json

data = {"name": "Alice", "age": 25, "city": "New York"}

with open("data.json", "w", encoding="utf-8") as file:
    json.dump(data, file, indent=4, ensure_ascii=False)
```

✔ `indent=4` → 보기 좋은 포맷으로 저장  
✔ `ensure_ascii=False` → 한글 깨짐 방지  

---

### 2) `json.load` - JSON 파일 읽기
```python
import json

with open("data.json", "r", encoding="utf-8") as file:
    data = json.load(file)

print(data)  # {'name': 'Alice', 'age': 25, 'city': 'New York'}
```

✔ `json.load()`를 사용하여 JSON 파일을 Python 딕셔너리로 변환  

---

## 9️⃣ 파일 존재 여부 확인

파일이 존재하는지 확인하려면 `os` 또는 `pathlib` 모듈을 사용할 수 있습니다.

### 1) `os.path.exists()`
```python
import os

if os.path.exists("example.txt"):
    print("파일이 존재합니다.")
else:
    print("파일이 없습니다.")
```

---

### 2) `pathlib.Path.exists()`
```python
from pathlib import Path

file_path = Path("example.txt")
if file_path.exists():
    print("파일이 존재합니다.")
```

✔ `os.path.exists()`와 `Path.exists()` 모두 파일 존재 여부를 확인할 수 있음  

---

## 🎯 정리

✔ **파일 열기 및 닫기** → `open("파일명", "모드")` 사용  
✔ **파일 읽기** → `read()`, `readline()`, `readlines()`  
✔ **파일 쓰기** → `write()`, `writelines()`  
✔ **`with` 문 사용** → 자동으로 파일 닫기 (`close()` 생략 가능)  
✔ **CSV/JSON 파일 처리** → `csv`, `json` 모듈 활용  
✔ **파일 존재 여부 확인** → `os.path.exists()` 또는 `Path.exists()`  
