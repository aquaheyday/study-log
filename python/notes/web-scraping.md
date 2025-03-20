# 🔄 Python 웹 스크래핑 (Web Scraping)

Python을 사용하여 웹 페이지의 데이터를 자동으로 수집하는 방법을 설명합니다.  
웹 스크래핑은 **`requests`**, **`BeautifulSoup`**, **`Selenium`** 등의 라이브러리를 활용하여 수행할 수 있습니다.

---

## 1. 웹 스크래핑(Web Scraping)이란?

- **웹 페이지에서 원하는 데이터를 추출하는 과정**을 의미합니다.
- 뉴스, 가격 정보, 상품 리뷰, 통계 데이터 등을 자동으로 수집하는 데 사용됩니다.
- Python에서는 `requests`, `BeautifulSoup`, `Selenium` 등을 활용하여 웹 크롤링을 수행할 수 있습니다.

---

## 2. 웹 스크래핑을 위한 라이브러리 설치

### 기본 라이브러리 설치 (`requests`, `BeautifulSoup`)
```sh
pip install requests beautifulsoup4
```

### 동적 페이지 크롤링을 위한 `Selenium` 설치
```sh
pip install selenium
```

✔ **Selenium을 사용하려면 웹드라이버(Chrome, Firefox 등)를 설치해야 합니다.**  
✔ [Chrome WebDriver 다운로드](https://chromedriver.chromium.org/downloads)

---

## 3. `requests`와 `BeautifulSoup`를 활용한 기본 웹 스크래핑

### `requests` - 웹 페이지 HTML 가져오기
```python
import requests

url = "https://example.com"
response = requests.get(url)

print(response.status_code)  # 200 (정상 요청 확인)
print(response.text)  # HTML 내용 출력
```

✔ `requests.get(url)`을 사용하여 웹 페이지의 HTML을 가져올 수 있음  
✔ `response.status_code == 200`이면 정상적으로 응답을 받은 것  

---

## 4. HTML 파싱 (`BeautifulSoup` 사용)

### `BeautifulSoup`을 사용한 HTML 파싱
```python
from bs4 import BeautifulSoup

html = "<html><body><h1>Hello, Web Scraping!</h1></body></html>"
soup = BeautifulSoup(html, "html.parser")

print(soup.h1.text)  # Hello, Web Scraping!
```

✔ `BeautifulSoup(html, "html.parser")`을 사용하여 HTML을 분석  

---

## 5. 웹 페이지에서 원하는 데이터 추출

### 특정 태그 가져오기 (`find()`, `find_all()`)
```python
html = """
<html>
    <body>
        <h1>웹 스크래핑 예제</h1>
        <p class="content">첫 번째 문장</p>
        <p class="content">두 번째 문장</p>
    </body>
</html>
"""
soup = BeautifulSoup(html, "html.parser")

h1 = soup.find("h1")  # 첫 번째 h1 태그 찾기
print(h1.text)  # "웹 스크래핑 예제"

p_tags = soup.find_all("p", class_="content")  # 모든 p 태그 찾기
for p in p_tags:
    print(p.text)
```

✔ `find("태그")` → 특정 태그 하나 찾기  
✔ `find_all("태그")` → 특정 태그 모두 찾기  

---

### 링크(URL) 가져오기 (`a` 태그)
```python
html = """
<html>
    <body>
        <a href="https://example.com">Example</a>
        <a href="https://google.com">Google</a>
    </body>
</html>
"""
soup = BeautifulSoup(html, "html.parser")

links = soup.find_all("a")  # 모든 링크 찾기
for link in links:
    print(link["href"])  # href 속성 값 출력
```

✔ `soup.find_all("a")` → 모든 `<a>` 태그 찾기  
✔ `link["href"]` → 링크 주소 추출  

---

### 웹 페이지에서 이미지 가져오기 (`img` 태그)
```python
html = """
<html>
    <body>
        <img src="image1.jpg" alt="이미지1">
        <img src="image2.jpg" alt="이미지2">
    </body>
</html>
"""
soup = BeautifulSoup(html, "html.parser")

images = soup.find_all("img")  # 모든 이미지 태그 찾기
for img in images:
    print(img["src"])  # 이미지 경로 출력
```

✔ `soup.find_all("img")` → 모든 이미지 태그 찾기  
✔ `img["src"]` → 이미지 파일 경로 추출  

---

## 6. 동적 웹 페이지 스크래핑 (`Selenium` 활용)

일부 웹사이트는 **JavaScript** 로 데이터를 로딩하므로 `requests` 만으로는 데이터를 가져올 수 없습니다.  
이 경우 **Selenium** 을 사용하여 웹페이지를 자동으로 조작할 수 있습니다.

### `Selenium`을 사용하여 동적 웹 페이지 크롤링
```python
from selenium import webdriver
from selenium.webdriver.common.by import By

# Chrome WebDriver 실행
driver = webdriver.Chrome()

# 웹사이트 접속
driver.get("https://example.com")

# 특정 요소 가져오기 (예: h1 태그)
h1_text = driver.find_element(By.TAG_NAME, "h1").text
print(h1_text)

# 브라우저 종료
driver.quit()
```

✔ `webdriver.Chrome()` → Chrome 브라우저 실행  
✔ `driver.get(url)` → 웹페이지 열기  
✔ `find_element(By.TAG_NAME, "h1")` → 특정 태그 찾기  

---

### `click()` - 버튼 클릭
```python
button = driver.find_element(By.ID, "search-button")
button.click()  # 버튼 클릭
```

✔ `find_element(By.ID, "search-button")` → ID가 "search-button"인 버튼 찾기  
✔ `.click()` → 버튼 클릭  

---

### `send_keys()` - 입력 상자에 텍스트 입력
```python
search_box = driver.find_element(By.NAME, "q")
search_box.send_keys("Python 웹 스크래핑")  # 검색어 입력
search_box.submit()  # 엔터 입력
```

✔ `send_keys("텍스트")` → 입력 상자에 값 입력  
✔ `.submit()` → 폼 제출  

---

## 7. 웹 크롤링 시 주의할 점

### 웹 크롤링의 법적 이슈
- **웹사이트의 `robots.txt` 규칙을 확인**하세요.  
- 데이터를 무단으로 수집하면 **법적 문제**가 발생할 수 있습니다.  
- 트래픽을 과도하게 발생시키면 **IP 차단**될 수도 있습니다.  

### 서버 부하 방지 (`time.sleep()` 활용)
```python
import time

for i in range(5):
    print(f"페이지 {i+1} 크롤링 중...")
    time.sleep(2)  # 2초 대기
```

✔ `time.sleep(2)` → 요청 간격을 조절하여 서버 부하 방지  

---

## 🎯 정리

✔ **정적 웹 페이지 크롤링** → `requests` + `BeautifulSoup` 사용  
✔ **동적 웹 페이지 크롤링** → `Selenium` 사용 (JavaScript 로딩 필요)  
✔ **HTML 요소 추출** → `find()`, `find_all()`, `select()` 활용  
✔ **웹 크롤링 시 주의 사항** → `robots.txt` 확인, 서버 부하 방지  
