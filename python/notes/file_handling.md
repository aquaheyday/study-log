# Python 파일 입출력 (I/O) 정리

Python에서 파일을 읽고 쓰는 방법을 정리합니다.

---

## 1. 파일 열기와 닫기

파일을 열 때 `open()` 함수를 사용하며, 모드를 지정해야 합니다.

| 모드 | 설명 |
|------|----------------|
| `"r"` | 읽기 모드 (기본값, 파일이 존재해야 함) |
| `"w"` | 쓰기 모드 (파일이 존재하면 덮어쓰기, 없으면 생성) |
| `"a"` | 추가 모드 (파일 끝에 내용 추가) |
| `"b"` | 바이너리 모드 (예: `"rb"`, `"wb"`) |

```python
# 파일 열기
file = open("example.txt", "w")
file.write("Hello, Python!")  # 파일에 쓰기
file.close()  # 파일 닫기
```

---

## 2. 파일 쓰기 (Write)
파일을 쓰기 모드(`"w"`)로 열면, 기존 내용이 삭제됩니다.

```python
with open("example.txt", "w") as file:
    file.write("첫 번째 줄\n")
    file.write("두 번째 줄\n")
```

### 2.1 여러 줄 쓰기
```python
lines = ["Python\n", "File I/O\n", "Example\n"]

with open("example.txt", "w") as file:
    file.writelines(lines)  # 리스트를 한 번에 쓰기
```

---

## 3. 파일 읽기 (Read)
파일을 읽을 때 `"r"` 모드를 사용합니다.

### 3.1 `read()` - 파일 전체 읽기
```python
with open("example.txt", "r") as file:
    content = file.read()
    print(content)  # 파일 전체 출력
```

### 3.2 `readline()` - 한 줄씩 읽기
```python
with open("example.txt", "r") as file:
    line = file.readline()
    print(line.strip())  # 개행 문자 제거
```

### 3.3 `readlines()` - 모든 줄을 리스트로 반환
```python
with open("example.txt", "r") as file:
    lines = file.readlines()
    for line in lines:
        print(line.strip())  # 개행 문자 제거 후 출력
```

---

## 4. 파일 추가 쓰기 (Append)
파일을 `"a"` 모드로 열면 기존 내용을 유지하면서 새로운 내용을 추가할 수 있습니다.

```python
with open("example.txt", "a") as file:
    file.write("추가된 내용\n")
```

---

## 5. 바이너리 파일 읽기 및 쓰기
이미지, 오디오, 동영상 등의 바이너리 파일을 다룰 때는 `"b"` 모드를 사용합니다.

### 5.1 바이너리 파일 쓰기
```python
with open("binary_file.bin", "wb") as file:
    file.write(b"Hello, Binary!")
```

### 5.2 바이너리 파일 읽기
```python
with open("binary_file.bin", "rb") as file:
    content = file.read()
    print(content)
```

---

## 6. 파일 예외 처리
파일을 다룰 때는 예외 처리를 통해 오류를 방지할 수 있습니다.

```python
try:
    with open("nonexistent.txt", "r") as file:
        content = file.read()
except FileNotFoundError:
    print("파일이 존재하지 않습니다!")
```
