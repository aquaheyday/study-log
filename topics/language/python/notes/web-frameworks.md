# 🌍 Python 웹 개발 (Web Development)

Python을 사용하여 웹 애플리케이션을 개발하는 방법을 설명합니다.  
주요 웹 프레임워크로 **Flask** 와 **Django** 를 사용하며, 웹 서버와 API 구축을 위한 개념을 포함합니다.

---

## 1️⃣ 웹 개발이란?

웹 개발(Web Development)이란, **웹 애플리케이션(Web Application)** 을 만들고 실행하는 과정입니다.  
Python을 사용하면 백엔드 서버를 구축하고 클라이언트 요청을 처리할 수 있습니다.

### 웹 애플리케이션 구성 요소
- **프론트엔드(Frontend)**: HTML, CSS, JavaScript로 UI(사용자 인터페이스) 구성
- **백엔드(Backend)**: 서버에서 비즈니스 로직을 처리하고 데이터베이스 연동
- **데이터베이스(Database)**: 사용자 정보, 게시글, 제품 데이터 등을 저장

Python은 주로 **백엔드 개발**에 사용되며, 대표적인 프레임워크로 **Flask, Django** 가 있습니다.

---

## 2️⃣ Flask vs Django 비교

| 프레임워크 | 특징 |
|-----------|------------------------------------------------|
| **Flask** | 가벼운 마이크로 웹 프레임워크, 단순하고 빠른 개발 가능 |
| **Django** | 강력한 기능과 보안이 포함된 풀스택 웹 프레임워크 |

✔ **Flask** → 간단한 API, 소규모 프로젝트에 적합  
✔ **Django** → 대규모 애플리케이션, 보안이 중요한 프로젝트에 적합  

---

## 3️⃣ Flask 기본 사용법

### 1) Flask 설치
```sh
pip install flask
```

---

### 2) 기본 웹 서버 만들기
```python
from flask import Flask

app = Flask(__name__)

@app.route("/")
def home():
    return "Hello, Flask!"

if __name__ == "__main__":
    app.run(debug=True)
```
✔ `app = Flask(__name__)` → Flask 애플리케이션 생성  
✔ `@app.route("/")` → 루트 URL (`/`)에 대한 요청 처리  
✔ `app.run(debug=True)` → 개발 서버 실행 (변경 사항 자동 반영)  

---

### 3) 라우팅 (URL 관리)
```python
@app.route("/hello/<name>")
def hello(name):
    return f"Hello, {name}!"
```
✔ `/<name>` → 동적 URL을 사용하여 변수 전달 가능  

---

### 4) HTML 템플릿 사용 (`render_template`)
```sh
mkdir templates
```
**`templates/index.html`** 파일 생성:
```html
<!DOCTYPE html>
<html>
<head><title>Flask App</title></head>
<body>
    <h1>Hello, {{ name }}!</h1>
</body>
</html>
```
#### Flask에서 HTML 렌더링
```python
from flask import render_template

@app.route("/greet/<name>")
def greet(name):
    return render_template("index.html", name=name)
```
✔ `render_template("index.html", name=name)` → HTML 파일과 데이터 연결  

---

## 4️⃣ Django 기본 사용법

### 1) Django 설치
```sh
pip install django
```

---

### 2) 프로젝트 생성 및 실행
```sh
django-admin startproject myproject
cd myproject
python manage.py runserver
```
✔ `startproject` → Django 프로젝트 생성  
✔ `runserver` → 개발 서버 실행 (http://127.0.0.1:8000/)  

---

### 3) 앱 생성 및 라우팅 설정
```sh
python manage.py startapp myapp
```
**`myapp/views.py`**
```python
from django.http import HttpResponse

def home(request):
    return HttpResponse("Hello, Django!")
```
**`myproject/urls.py`**
```python
from django.contrib import admin
from django.urls import path
from myapp.views import home

urlpatterns = [
    path("admin/", admin.site.urls),
    path("", home),  # 루트 URL에 home 뷰 연결
]
```
✔ `path("", home)` → `/` URL을 `home` 함수와 연결  

---

### 4) 모델 및 데이터베이스 설정

Django는 기본적으로 **SQLite 데이터베이스**를 사용합니다.  
모델을 정의하면 자동으로 데이터베이스 테이블이 생성됩니다.

**`myapp/models.py`**
```python
from django.db import models

class Post(models.Model):
    title = models.CharField(max_length=100)
    content = models.TextField()
    created_at = models.DateTimeField(auto_now_add=True)
```
✔ `CharField`, `TextField`, `DateTimeField` 등 다양한 필드 제공  

#### 마이그레이션 및 데이터베이스 적용
```sh
python manage.py makemigrations
python manage.py migrate
```
✔ `makemigrations` → 변경된 모델을 마이그레이션 파일로 생성  
✔ `migrate` → 데이터베이스에 적용  

---

## 5️⃣ REST API 구축

웹 애플리케이션에서 **REST API** 를 제공하려면 Flask 또는 Django REST Framework를 사용할 수 있습니다.

### 1) Flask REST API (`Flask-RESTful` 사용)
```sh
pip install flask flask-restful
```

---

#### API 엔드포인트 만들기
```python
from flask import Flask, jsonify
from flask_restful import Api, Resource

app = Flask(__name__)
api = Api(app)

class HelloWorld(Resource):
    def get(self):
        return jsonify({"message": "Hello, API!"})

api.add_resource(HelloWorld, "/api/hello")

if __name__ == "__main__":
    app.run(debug=True)
```
✔ `GET /api/hello` 요청 시 JSON 응답 반환  

---

### 2) Django REST API (`Django REST Framework` 사용)
```sh
pip install djangorestframework
```

#### `myapp/views.py`
```python
from rest_framework.response import Response
from rest_framework.decorators import api_view

@api_view(["GET"])
def hello_api(request):
    return Response({"message": "Hello, Django REST API!"})
```
#### `myproject/urls.py`
```python
from django.urls import path
from myapp.views import hello_api

urlpatterns = [
    path("api/hello/", hello_api),
]
```
✔ `GET /api/hello/` 요청 시 JSON 응답 반환  

---

## 6️⃣ 웹 서버 배포

### 1) Flask 배포 (Gunicorn + Nginx)
```sh
pip install gunicorn
gunicorn -w 4 -b 0.0.0.0:5000 app:app
```
✔ Gunicorn을 사용하여 Flask 애플리케이션을 프로덕션 환경에서 실행  

---

### 2) Django 배포 (Gunicorn + Nginx)
```sh
pip install gunicorn
gunicorn myproject.wsgi:application --bind 0.0.0.0:8000
```
✔ `wsgi.py`를 사용하여 Django 애플리케이션 실행  

---

## 🎯 정리

✔ **Flask** → 가볍고 빠른 웹 프레임워크 (소규모 프로젝트에 적합)  
✔ **Django** → 강력한 기능 제공 (대규모 프로젝트에 적합)  
✔ **HTML 템플릿** → `render_template()` (Flask), `templates/` (Django)  
✔ **REST API 구축** → `Flask-RESTful`, `Django REST Framework` 활용  
✔ **배포** → Gunicorn, Nginx를 사용하여 프로덕션 서버 운영  
