# 🚀 Python 정규 표현식 (Regular Expressions)

Python에서 **정규 표현식(Regular Expressions, RegEx)** 은 **문자열에서 특정 패턴을 검색, 추출, 치환하는 기능**을 제공합니다.  
이를 위해 **`re` 모듈**을 사용합니다.

---

## 1️⃣ 정규 표현식이란?

- **특정 문자열 패턴을 찾기 위한 기법**입니다.
- 이메일 주소, 전화번호, URL 등의 패턴을 검출하는 데 유용합니다.
- `re` 모듈을 사용하여 문자열 검색, 매칭, 치환을 수행할 수 있습니다.

#### `re` 모듈 불러오기
```python
import re
```

---

## 2️⃣ 정규 표현식 기본 패턴

| 정규식 | 설명 | 예제 매칭 |
|--------|------|----------|
| `.`    | 아무 문자 한 개 | `"a", "b", "1", "@"` |
| `^`    | 문자열의 시작 | `"^Hello"` → `"Hello world"` |
| `$`    | 문자열의 끝 | `"world$"` → `"Hello world"` |
| `*`    | 0개 이상 반복 | `"a*"` → `"a", "aaa", ""` |
| `+`    | 1개 이상 반복 | `"a+"` → `"a", "aaa"` (빈 문자열 X) |
| `?`    | 0개 또는 1개 | `"colou?r"` → `"color", "colour"` |
| `{n}`  | 정확히 n개 반복 | `"a{3}"` → `"aaa"` |
| `{n,}` | n개 이상 반복 | `"a{2,}"` → `"aa", "aaa", "aaaa"` |
| `{n,m}` | n개 이상, m개 이하 반복 | `"a{2,4}"` → `"aa", "aaa", "aaaa"` |
| `[]`   | 문자 클래스(여러 문자 중 하나) | `"[abc]"` → `"a", "b", "c"` |
| `\d`   | 숫자 `[0-9]` | `"123"` |
| `\D`   | 숫자가 아닌 문자 | `"abc", "@"` |
| `\w`   | 영문자, 숫자, `_` | `"abc123_"` |
| `\W`   | 영문자, 숫자, `_`가 아닌 문자 | `"@#$%"` |
| `\s`   | 공백 문자 (스페이스, 탭, 개행) | `" "` |
| `\S`   | 공백이 아닌 문자 | `"abc123"` |

---

## 3️⃣ 정규 표현식 기본 사용법 (`re` 모듈)

### 1) `re.search()` - 패턴 검색
```python
import re

text = "Hello, my email is example@mail.com"
match = re.search(r"\w+@\w+\.\w+", text)

if match:
    print("이메일 주소:", match.group())  # example@mail.com
```
✔ `re.search(pattern, text)` → **처음으로 매칭된 결과 반환**  
✔ `.group()` → 매칭된 문자열 출력  

---

### 2) `re.findall()` - 모든 패턴 찾기
```python
text = "My numbers are 123-4567 and 987-6543"
matches = re.findall(r"\d{3}-\d{4}", text)
print(matches)  # ['123-4567', '987-6543']
```
✔ `re.findall(pattern, text)` → **모든 매칭 결과를 리스트로 반환**  

---

### 3) `re.match()` - 문자열의 시작 부분에서 매칭 확인
```python
text = "Python is great!"
match = re.match(r"Python", text)

if match:
    print("매칭됨:", match.group())  # Python
```
✔ `re.match(pattern, text)` → **문자열의 시작 부분에서만 매칭 확인**  

---

### 4) `re.sub()` - 문자열 치환 (Replace)
```python
text = "I love Python!"
new_text = re.sub(r"Python", "Regex", text)
print(new_text)  # I love Regex!
```
✔ `re.sub(pattern, replace, text)` → 특정 패턴을 다른 문자열로 변경  

---

## 4️⃣ 정규 표현식 그룹핑 (`()` 사용)

#### 그룹핑된 데이터 가져오기
```python
text = "My phone number is 010-1234-5678"
match = re.search(r"(\d{3})-(\d{4})-(\d{4})", text)

if match:
    print("전체 번호:", match.group())   # 010-1234-5678
    print("첫 번째 그룹:", match.group(1))  # 010
    print("두 번째 그룹:", match.group(2))  # 1234
    print("세 번째 그룹:", match.group(3))  # 5678
```
✔ **`()` 안의 패턴을 그룹으로 인식**  
✔ `.group(1)`, `.group(2)` → 각각의 그룹에 접근 가능  

---

## 5️⃣ `re.compile()` - 정규식 패턴 미리 컴파일

`re.compile()`을 사용하면 정규식을 **미리 컴파일하여 성능을 향상**시킬 수 있습니다.

#### 정규식 컴파일 후 재사용
```python
pattern = re.compile(r"\d{3}-\d{4}-\d{4}")
text = "My number is 010-1234-5678"

match = pattern.search(text)
if match:
    print("전화번호:", match.group())  # 010-1234-5678
```
✔ `re.compile(pattern)` → 정규식을 **객체로 저장하여 반복 사용 가능**  

---

## 6️⃣ 정규 표현식 활용 예제

### 1) 이메일 주소 검증
```python
email = "example@mail.com"
pattern = r"^\w+@\w+\.\w+$"

if re.match(pattern, email):
    print("유효한 이메일 주소입니다.")
else:
    print("잘못된 이메일 주소입니다.")
```
✔ 이메일 형식을 검사하여 올바른지 확인  

---

### 2) 비밀번호 강도 검사 (대문자, 소문자, 숫자 포함)
```python
password = "StrongPass123!"
pattern = r"^(?=.*[a-z])(?=.*[A-Z])(?=.*\d).{8,}$"

if re.match(pattern, password):
    print("안전한 비밀번호입니다.")
else:
    print("비밀번호는 최소 8자 이상이며, 대문자/소문자/숫자를 포함해야 합니다.")
```
✔ `(?=.*[a-z])` → 소문자 포함  
✔ `(?=.*[A-Z])` → 대문자 포함  
✔ `(?=.*\d)` → 숫자 포함  

---

### 3) 한글만 포함된 문자열 검사
```python
text = "안녕하세요"
pattern = r"^[가-힣]+$"

if re.match(pattern, text):
    print("한글만 포함된 문자열입니다.")
else:
    print("한글 외의 문자가 포함되었습니다.")
```
✔ `^[가-힣]+$` → 한글만 포함된 문자열 매칭  

---

## 🎯 정리

✔ **정규 표현식 패턴** → `\d`, `\w`, `[]`, `*`, `+`, `{}` 등 사용  
✔ **문자열 검색** → `re.search()`, `re.match()`, `re.findall()`  
✔ **문자열 치환** → `re.sub()`  
✔ **그룹핑** → `()`를 사용하여 특정 부분 추출 가능  
✔ **`re.compile()`** → 패턴을 미리 컴파일하여 성능 향상  
✔ **이메일, 전화번호, 비밀번호 검증** 등 다양한 활용 가능  

