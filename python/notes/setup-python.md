# 🛠 Python 개발 환경 설정

Python 개발을 시작하기 위한 환경 설정 방법을 정리합니다.

---

## 1. Python 설치

Python 설치 방법에 대해 설명합니다.

| 운영체제 | 설치 방법 |
|----------|----------------|
| Windows | [Python 공식 사이트](https://www.python.org/downloads/)에서 설치 |
| macOS | `brew install python` 명령어 사용 |
| Linux | `sudo apt install python3` 명령어 사용 |

설치 완료 후, 터미널에서 Python 버전 확인:
```sh
python --version
```

---

## 2. 가상 환경 (Python 프로젝트별 독립적인 환경 구성)

### `venv` (Python 기본 가상 환경)
Python 내장 모듈 `venv`를 사용하여 가상 환경을 설정합니다.

```sh
# 가상 환경 생성
python -m venv myenv
```

#### ✅ 가상 환경 활성화 
- **Windows (CMD / PowerShell)**
  ```sh
  myenv\Scripts\activate
  ```
- **macOS / Linux**
  ```sh
  source myenv/bin/activate
  ```

#### ❌ 가상 환경 비활성화
```sh
deactivate
```

---

## 3. 패키지 관리 (`pip` 사용)

| 기능 | 명령어 |
|----------|----------------|
| 패키지 설치 | `pip install package_name` |
| 패키지 제거 | `pip uninstall package_name` |
| 설치된 패키지 목록 확인 | `pip list` |

#### `requirements.txt` 사용 방법
```sh
# 현재 환경의 패키지 목록 저장
pip freeze > requirements.txt

# 패키지 목록 기반으로 설치
pip install -r requirements.txt
```

---

## 4. 플랫폼 지원

Python은 다양한 운영체제에서 실행할 수 있습니다.

| 플랫폼 | 지원 상태 | 주요 특징 |
|--------|---------|----------|
| Windows | ✅ 완전 지원 | 네이티브 실행 가능 |
| macOS | ✅ 완전 지원 | Homebrew를 통한 간편 설치 |
| Linux | ✅ 완전 지원 | 다양한 패키지 관리자를 통한 설치 가능 |

---

## 5. Python 기본 문법

Python의 기초적인 문법을 소개합니다.

#### ✅ 올바른 Python 코드
```python
def greet():
    print("Hello, Python!")

greet()
```

#### ❌ 잘못된 코드 (SyntaxError 발생)
```python
def greet()
    print("Hello, Python!")
```
> **해결 방법**  
```python
def greet():
    print("Hello, Python!")
```

---

## 🎯 정리

✔ Python 설치 및 환경 설정을 위한 필수 과정 정리  
✔ `venv`를 활용한 가상 환경 설정 방법  
✔ `pip`을 활용한 패키지 관리 방법  
