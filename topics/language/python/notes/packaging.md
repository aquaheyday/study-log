# 📦 Python 패키징과 배포 (Packaging & Deployment)

Python 애플리케이션을 다른 환경에서도 쉽게 설치하고 실행할 수 있도록 **패키징과 배포**하는 방법을 정리합니다.  
주요 내용:
- **패키징(Packaging)** → 코드 구조화 및 배포 가능하게 만들기
- **PyPI(Python Package Index) 배포** → `pip install`로 설치 가능하도록 업로드
- **실행 파일(.exe) 생성** → Python 없이 실행 가능하게 빌드

---

## 1. 패키징(Packaging)이란?

- **Python 프로젝트를 하나의 모듈 또는 라이브러리로 묶어 배포하는 과정**
- **패키징을 통해 다른 개발자가 쉽게 설치 가능 (`pip install` 활용)**
- **PyPI(Python Package Index)** 에 업로드하면 전 세계에서 패키지를 설치할 수 있음

---

## 2. Python 패키지 기본 구조

### 패키지 디렉토리 구조 예제
```
my_project/
│── my_package/
│   ├── __init__.py  # 패키지 초기화 파일
│   ├── module1.py   # 첫 번째 모듈
│   ├── module2.py   # 두 번째 모듈
│── setup.py  # 패키지 설정 파일
│── README.md  # 설명 파일
│── requirements.txt  # 패키지 종속성
```

✔ **`__init__.py`** → 패키지를 인식할 수 있도록 하는 파일  
✔ **`setup.py`** → 패키지를 설치하고 PyPI에 업로드하기 위한 설정 파일  
✔ **`requirements.txt`** → 의존성(필수 패키지) 관리  

---

## 3. `setup.py` 작성 (패키지 정보 설정)

### `setup.py` 파일 예제
```python
from setuptools import setup, find_packages

setup(
    name="my_package",
    version="0.1.0",
    author="Your Name",
    author_email="your_email@example.com",
    description="이 패키지는 예제 패키지입니다.",
    long_description=open("README.md").read(),
    long_description_content_type="text/markdown",
    url="https://github.com/yourusername/my_package",
    packages=find_packages(),
    install_requires=[
        "numpy", "pandas"  # 패키지 설치 시 필요한 종속성
    ],
    classifiers=[
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
    ],
    python_requires=">=3.6",
)
```

✔ **패키지 정보** (`name`, `version`, `author`, `description`) 설정  
✔ **`find_packages()`** → 패키지 내부의 모든 모듈을 자동 검색  
✔ **`install_requires`** → 패키지 설치 시 필요한 라이브러리 자동 설치  

---

## 4. 패키지 빌드 및 설치

### 패키지 빌드 (배포 전 준비)
```sh
pip install setuptools wheel
python setup.py sdist bdist_wheel
```
✔ `sdist` → 소스 코드 배포용 압축 파일 생성  
✔ `bdist_wheel` → 바이너리 배포 파일 생성 (더 빠른 설치 가능)  

### 로컬에서 패키지 설치 테스트
```sh
pip install .
```
✔ 현재 디렉토리에 있는 패키지를 직접 설치하여 테스트 가능  

---

## 5. PyPI(Python Package Index)에 패키지 배포

### Twine 설치 (PyPI 업로드 도구)
```sh
pip install twine
```

### PyPI 계정 생성
- [PyPI 공식 사이트](https://pypi.org/)에서 계정 생성

### 패키지 업로드
```sh
twine upload dist/*
```
✔ `dist/` 폴더 내의 빌드된 패키지를 PyPI에 업로드  
✔ 이후 `pip install my_package` 로 설치 가능  

---

## 6. 실행 파일(.exe)로 패키징 (`PyInstaller` 사용)

### PyInstaller 설치
```sh
pip install pyinstaller
```

### 단일 실행 파일(.exe) 생성
```sh
pyinstaller --onefile main.py
```
✔ `--onefile` → 단일 실행 파일 생성  
✔ `dist/main.exe` 에 실행 파일 생성됨  

### 아이콘 추가 및 창 없이 실행
```sh
pyinstaller --onefile --noconsole --icon=app.ico main.py
```
✔ `--noconsole` → 콘솔 창 없이 실행 (GUI 앱 용)  
✔ `--icon=app.ico` → 실행 파일 아이콘 설정  

---

## 7. `requirements.txt` 로 패키지 종속성 관리

### 프로젝트의 종속성 저장
```sh
pip freeze > requirements.txt
```
✔ 현재 프로젝트에서 사용 중인 패키지를 `requirements.txt` 에 저장  

### 종속성 설치
```sh
pip install -r requirements.txt
```
✔ `requirements.txt` 기반으로 필요한 패키지 자동 설치  

---

## 8. 가상 환경 설정 (`venv`)

패키징 및 배포 전에 **가상 환경을 설정하면 충돌 없이 패키지 관리 가능**.

### 가상 환경 생성
```sh
python -m venv venv
```

### 가상 환경 활성화
```sh
# Windows
venv\Scripts\activate

# Mac/Linux
source venv/bin/activate
```

### 가상 환경 비활성화
```sh
deactivate
```

✔ **가상 환경 사용 시, 패키징과 배포 테스트를 안전하게 수행 가능**  

---

## 9. Docker를 활용한 배포 (컨테이너화)

패키지를 **컨테이너화(Dockerization)** 하여 배포 가능.

### `Dockerfile` 작성
```dockerfile
FROM python:3.9

WORKDIR /app
COPY . .

RUN pip install -r requirements.txt
CMD ["python", "main.py"]
```

### Docker 이미지 빌드 & 실행
```sh
docker build -t my_python_app .
docker run my_python_app
```

✔ **Docker를 사용하면 실행 환경이 통일되어 어디서든 동일한 실행 가능**  

---

## 🎯 정리

✔ **Python 패키징** → `setup.py` 를 활용하여 배포 가능한 패키지 제작  
✔ **PyPI 배포** → `twine`을 이용하여 `pip install`로 설치 가능하게 업로드  
✔ **실행 파일 생성** → `pyinstaller` 로 `.exe` 또는 `.app` 빌드  
✔ **가상 환경 활용** → `venv` 를 사용하여 패키지 종속성 관리  
✔ **Docker 배포** → Python 애플리케이션을 컨테이너로 배포  
