# 🐍 Python

이 폴더는 **Python 학습 과정**에서 정리한 자료, 예제 코드, 프로젝트를 보관하는 공간입니다.  
기본 개념부터 실습 예제, 프로젝트 구성까지 단계적으로 학습한 내용을 체계적으로 정리합니다.

---

## 📂 디렉토리 구성

| 폴더명 | 설명 |
|---|---|
| [notes](./notes) | 개념 정리 및 학습 노트 |
| [examples](./examples) | 주요 기능별 예제 코드 (컴포넌트, 상태관리 등) |
| [projects](./projects) | 미니 프로젝트 및 실습 프로젝트 |

---

## 📋 Python 개념 정리 목록

### 📌 기본 개념
| 번호 | 주제 | 파일명 | 설명 |
|---|---|---|---|
| 01 | Python 개요 | [python-intro.md](./notes/python-intro.md) | Python이란? 특징과 장점 |
| 02 | 개발 환경 설정 | [setup-python.md](./notes/setup-python.md) | Python 설치, 가상 환경, 패키지 관리 |
| 03 | 변수와 자료형 | [variables.md](./notes/variables.md) | 숫자, 문자열, 리스트, 딕셔너리 등 |
| 04 | 연산자와 표현식 | [operators.md](./notes/operators.md) | 산술, 비교, 논리, 비트 연산자 |
| 05 | 조건문과 반복문 | [control-flow.md](./notes/control-flow.md) | if, for, while 문법과 사용법 |

### 🎯 함수와 객체지향 프로그래밍
| 번호 | 주제 | 파일명 | 설명 |
|---|---|---|---|
| 06 | 함수 사용법 | [functions.md](./notes/functions.md) | 함수 정의, 매개변수, 반환값, 람다 |
| 07 | 모듈과 패키지 | [modules.md](./notes/modules.md) | 모듈 가져오기, 패키지 구조 이해 |
| 08 | 클래스와 객체 | [oop.md](./notes/oop.md) | 클래스, 객체, 상속, 다형성 |
| 09 | 예외 처리 | [exceptions.md](./notes/exceptions.md) | try-except, raise, 사용자 정의 예외 |

### 🔄 데이터 처리
| 번호 | 주제 | 파일명 | 설명 |
|---|---|---|---|
| 10 | 파일 입출력 | [file-io.md](./notes/file-io.md) | 텍스트 파일, CSV, JSON 다루기 |
| 11 | 데이터베이스 | [database.md](./notes/database.md) | SQLite, MySQL, PostgreSQL 활용 |
| 12 | 웹 스크래핑 | [web-scraping.md](./notes/web-scraping.md) | BeautifulSoup, Selenium 사용법 |

### 🌍 네트워크와 API 연동
| 번호 | 주제 | 파일명 | 설명 |
|---|---|---|---|
| 13 | HTTP 요청 | [http-requests.md](./notes/http-requests.md) | requests 라이브러리, API 호출 |
| 14 | 비동기 프로그래밍 | [async.md](./notes/async.md) | async/await, asyncio 활용법 |
| 15 | 웹 개발 | [web-frameworks.md](./notes/web-frameworks.md) | Flask, FastAPI, Django 기본 |

### 🚀 고급 개념
| 번호 | 주제 | 파일명 | 설명 |
|---|---|---|---|
| 16 | 정규 표현식 | [regex.md](./notes/regex.md) | re 모듈, 패턴 매칭 활용 |
| 17 | 멀티스레딩 | [multithreading.md](./notes/multithreading.md) | threading, multiprocessing 비교 |
| 18 | 데코레이터 | [decorators.md](./notes/decorators.md) | 함수형 프로그래밍, 고차 함수 |
| 19 | 제너레이터와 이터레이터 | [generators.md](./notes/generators.md) | yield, lazy evaluation |

### 🛠️ 데이터 과학 및 머신러닝
| 번호 | 주제 | 파일명 | 설명 |
|---|---|---|---|
| 20 | NumPy 기본 | [numpy.md](./notes/numpy.md) | 배열 생성, 연산, 슬라이싱 |
| 21 | Pandas 기본 | [pandas.md](./notes/pandas.md) | 데이터프레임 생성, 필터링, 그룹화 |
| 22 | 데이터 시각화 | [visualization.md](./notes/visualization.md) | Matplotlib, Seaborn 그래프 그리기 |
| 23 | 머신러닝 개요 | [ml-basics.md](./notes/ml-basics.md) | Scikit-learn을 활용한 기본 모델 |
| 24 | 딥러닝 기초 | [deep-learning.md](./notes/deep-learning.md) | TensorFlow, PyTorch 활용 |

### 🛠️ 테스트 및 배포
| 번호 | 주제 | 파일명 | 설명 |
|---|---|---|---|
| 25 | 단위 테스트 | [testing.md](./notes/testing.md) | unittest, pytest 활용법 |
| 26 | 패키징과 배포 | [packaging.md](./notes/packaging.md) | pip, setuptools, PyPI 배포 |

---

## 📋 예제 목록

| 번호 | 주제 | 파일명 | 설명 |
|---|---|---|---|
| 01 | 입출력 기본 | [input_output.py](./examples/input_output.py) | 표준 입력 및 출력 예제 |
| 02 | 변수 | [variables.py](./examples/variables.py) | 변수 선언 및 사용 예제 |
| 03 | 연산자 | [operators.py](./examples/operators.py) | 산술, 비교, 논리 연산자 활용 예제 |
| 04 | 문자열 | [strings.py](./examples/strings.py) | 문자열 다루기 및 메서드 활용 |
| 05 | 리스트 | [lists.py](./examples/lists.py) | 리스트 생성, 조작 및 메서드 활용 |
| 06 | 튜플 | [tuples.py](./examples/tuples.py) | 변경 불가능한 튜플의 활용 |
| 07 | 집합 | [sets.py](./examples/sets.py) | 집합 연산 및 활용 예제 |
| 08 | 딕셔너리 | [dictionaries.py](./examples/dictionaries.py) | 키-값 쌍을 저장하는 딕셔너리 활용 |
| 09 | 조건문 | [conditionals.py](./examples/conditionals.py) | if-elif-else 문 활용 예제 |
| 10 | 반복문 | [loops.py](./examples/loops.py) | for 및 while 반복문 예제 |
| 11 | 함수 | [functions.py](./examples/functions.py) | 함수 정의, 매개변수, 반환값 예제 |
| 12 | 리스트 컴프리헨션 | [list_comprehension.py](./examples/list_comprehension.py) | 리스트 컴프리헨션 활용 예제 |

---

## 📋 프로젝트 목록

| 번호 | 프로젝트명 | 설명 | 링크 |
|---|---|---|---|
| 01 | 택배/요금제/문의 관리 | 택배사 정보, 요금제, 문의사항을 관리하는 FastAPI 기반 서비스 | [delivery-api-server](./projects/delivery-api-server) |

---

## 📢 업데이트 로그
- 2025-03-06: 초기 디렉토리 구성
- 2025-03-07: projects > dilivery-api-server 추가
- 2025-03-10: 초기 notes 구성
- 2025-03-19: 디렉토리 구성 수정
- 새로운 예제 추가 시 목록 업데이트 예정
