# 🌍 Python HTTP 요청 (HTTP Request)

Python에서 **HTTP 요청을 보내는 방법**을 설명합니다.  
웹 서버와 데이터를 주고받을 때 `requests` 라이브러리를 사용하여 **GET, POST, PUT, DELETE** 등의 요청을 수행할 수 있습니다.

---

## 1. HTTP 요청이란?

- **HTTP 요청(Request)** 은 클라이언트가 웹 서버에 데이터를 요청하는 방식입니다.
- 요청 방식(Method)에 따라 수행하는 작업이 달라집니다.

### 주요 HTTP 요청 메서드

| HTTP 메서드 | 설명 |
|------------|----------------------------------|
| `GET` | 데이터를 가져올 때 사용 (읽기) |
| `POST` | 데이터를 전송할 때 사용 (생성) |
| `PUT` | 데이터를 업데이트할 때 사용 |
| `DELETE` | 데이터를 삭제할 때 사용 |

---

## 2. `requests` 라이브러리 설치 및 기본 사용법

Python에서 HTTP 요청을 보낼 때는 `requests` 라이브러리를 사용합니다.

### `requests` 설치
```sh
pip install requests
```

### 기본 GET 요청 (`requests.get()`)
```python
import requests

response = requests.get("https://jsonplaceholder.typicode.com/posts/1")

print(response.status_code)  # 응답 상태 코드 출력 (200: 성공)
print(response.text)  # 응답 본문 출력
```

✔ `requests.get(url)` → 웹 서버에 GET 요청을 보냄  
✔ `response.status_code` → 응답 상태 코드 확인  
✔ `response.text` → 응답 본문(HTML, JSON 등) 출력  

---

## 3. GET 요청 (Query Parameters 포함)

### URL에 파라미터 추가 (`params` 사용)
```python
params = {"userId": 1}
response = requests.get("https://jsonplaceholder.typicode.com/posts", params=params)

print(response.url)  # 요청된 URL 확인
print(response.json())  # JSON 데이터 출력
```

✔ `params={"key": "value"}` → GET 요청에 쿼리 문자열 추가  
✔ `response.url` → 최종 요청된 URL 확인  
✔ `response.json()` → JSON 응답을 파이썬 딕셔너리로 변환  

---

## 4. POST 요청 (데이터 전송)

### `requests.post()` 사용
```python
data = {"title": "Hello", "body": "This is a test", "userId": 1}
response = requests.post("https://jsonplaceholder.typicode.com/posts", json=data)

print(response.status_code)  # 응답 코드 출력 (201: 생성됨)
print(response.json())  # 응답 데이터 확인
```

✔ `json=data` → JSON 형식으로 데이터 전송  
✔ `response.status_code == 201`이면 성공적으로 생성됨  

---

## 5. PUT 요청 (데이터 수정)

### `requests.put()` 사용
```python
data = {"title": "Updated Title", "body": "Updated body", "userId": 1}
response = requests.put("https://jsonplaceholder.typicode.com/posts/1", json=data)

print(response.status_code)  # 응답 코드 (200: 성공)
print(response.json())  # 수정된 데이터 확인
```

✔ `PUT` 요청은 기존 데이터를 수정할 때 사용됨  

---

## 6. DELETE 요청 (데이터 삭제)

### `requests.delete()` 사용
```python
response = requests.delete("https://jsonplaceholder.typicode.com/posts/1")

print(response.status_code)  # 응답 코드 (200: 성공, 204: 삭제 완료)
```

✔ `DELETE` 요청을 보내면 서버에서 해당 데이터를 삭제  

---

## 7. 요청 헤더 설정

일부 API는 **헤더(Headers)** 를 요구할 수 있습니다.  
예를 들어, `User-Agent`, `Authorization`, `Content-Type` 등의 정보를 포함해야 할 수 있습니다.

### `headers` 사용 예제
```python
headers = {
    "User-Agent": "Mozilla/5.0",
    "Authorization": "Bearer YOUR_ACCESS_TOKEN"
}

response = requests.get("https://jsonplaceholder.typicode.com/posts", headers=headers)
print(response.status_code)
```

✔ `headers={"키": "값"}` → 요청 헤더 추가  

---

## 8. 응답(Response) 다루기

### 응답 데이터 확인 방법
```python
response = requests.get("https://jsonplaceholder.typicode.com/posts/1")

print(response.status_code)  # 상태 코드
print(response.text)  # 응답 본문 (문자열)
print(response.json())  # JSON 데이터 변환 (딕셔너리)
print(response.headers)  # 응답 헤더 출력
```

✔ `response.text` → 문자열 형태로 응답 데이터 확인  
✔ `response.json()` → JSON 형식의 응답 데이터를 딕셔너리로 변환  
✔ `response.headers` → 응답 헤더 출력  

---

## 9. 예외 처리 (`try-except` 사용)

요청이 실패할 경우 예외 처리를 추가하여 안정적인 코드를 작성할 수 있습니다.

### `try-except`로 예외 처리
```python
try:
    response = requests.get("https://jsonplaceholder.typicode.com/posts/1", timeout=5)
    response.raise_for_status()  # HTTP 에러 발생 시 예외 처리
    print(response.json())
except requests.exceptions.HTTPError as e:
    print(f"HTTP 오류 발생: {e}")
except requests.exceptions.ConnectionError:
    print("연결 오류 발생")
except requests.exceptions.Timeout:
    print("요청 시간이 초과되었습니다.")
except requests.exceptions.RequestException as e:
    print(f"요청 오류 발생: {e}")
```

✔ `raise_for_status()` → HTTP 에러(4xx, 5xx) 발생 시 예외 처리  
✔ `ConnectionError` → 인터넷 연결 문제 발생 시 처리  
✔ `Timeout` → 요청이 일정 시간 내 응답하지 않으면 처리  

---

## 10. 파일 업로드 및 다운로드

### 파일 업로드 (`files` 사용)
```python
files = {"file": open("example.txt", "rb")}

response = requests.post("https://httpbin.org/post", files=files)
print(response.json())  # 업로드된 파일 정보 확인
```

✔ `files={"file": open("파일명", "rb")}` → 파일 업로드  

---

### 파일 다운로드
```python
response = requests.get("https://example.com/image.jpg")

with open("downloaded_image.jpg", "wb") as file:
    file.write(response.content)  # 파일 저장
```

✔ `response.content` → 바이너리 데이터 다운로드 후 저장  

---

## 🎯 정리

✔ **GET 요청** → `requests.get(url, params=...)`  
✔ **POST 요청** → `requests.post(url, json=data)`  
✔ **PUT 요청** → `requests.put(url, json=data)`  
✔ **DELETE 요청** → `requests.delete(url)`  
✔ **요청 헤더 추가** → `headers={"User-Agent": "...", "Authorization": "..."}`  
✔ **응답 데이터 처리** → `response.json()`, `response.text`, `response.status_code`  
✔ **예외 처리** → `try-except`를 사용하여 안정적인 요청 구현  
✔ **파일 업로드/다운로드** → `files`, `response.content` 활용  
