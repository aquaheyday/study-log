# 🐘 PHP

이 폴더는 **PHP 학습 과정**에서 정리한 자료, 예제 코드, 프로젝트를 보관하는 공간입니다.  
기본 개념부터 실습 예제, 프로젝트 구성까지 단계적으로 학습한 내용을 체계적으로 정리합니다.

---

## 📂 디렉토리 구성

| 폴더명 | 설명 |
|---|---|
| [notes](./notes) | 개념 정리 및 학습 노트 |
| [examples](./examples) | 주요 기능별 예제 코드 |
| [projects](./projects) | 미니 프로젝트 및 실습 프로젝트 |

---

## 📋 개념 정리 목록

| 번호 | 주제 | 파일명 | 설명 |
|---|---|---|---|
| 01 | 기본 문법 | [basic-syntax.md](./notes/basic-syntax.md) | 변수, 자료형, 제어문 기초 정리 |
| 02 | 함수와 스코프 | [functions.md](./notes/functions.md) | 함수 선언, 매개변수, 반환값, 스코프 개념 |
| 03 | 배열과 연관 배열 | [arrays.md](./notes/arrays.md) | 배열 생성, 다차원 배열, 연관 배열 활용 |
| 04 | 폼 처리와 슈퍼글로벌 | [forms.md](./notes/forms.md) | GET/POST, $_GET, $_POST 사용법 |
| 05 | 파일 입출력 | [file-handling.md](./notes/file-handling.md) | 파일 열기, 쓰기, 읽기, 삭제 |
| 06 | 세션과 쿠키 | [session-cookie.md](./notes/session-cookie.md) | 로그인 유지 등 상태 관리 기법 |
| 07 | 데이터베이스 연동 | [database.md](./notes/database.md) | PDO 기반 데이터베이스 연결 및 쿼리 실행 |
| 08 | 에러 처리와 예외 | [error-handling.md](./notes/error-handling.md) | 에러 로그, try-catch, 사용자 정의 예외 |
| 09 | 간단한 MVC 패턴 | [mvc.md](./notes/mvc.md) | 기본적인 MVC 구조 설명 |
| 10 | 기타 유용한 팁 | [tips.md](./notes/tips.md) | 보안 관련 팁, PHP 설정 팁 등 |

---

## 📋 예제 목록

| 번호 | 주제 | 파일명 | 설명 |
|---|---|---|---|
| 01 | 변수와 자료형 | [variables.php](./examples/variables.php) | 변수 선언, 자료형 변환 등 기초 예제 |
| 02 | 연산자 | [operators.php](./examples/operators.php) | 산술, 비교, 논리 연산자 활용 예제 |
| 03 | 제어문 | [control-flow.php](./examples/control-flow.php) | if문, switch문, 반복문 예제 |
| 04 | 함수 | [functions.php](./examples/functions.php) | 함수 정의, 매개변수, 반환값 예제 |
| 05 | 배열과 연관 배열 | [arrays.php](./examples/arrays.php) | 배열 생성, 다차원 배열 등 |
| 06 | 파일 입출력 | [file-handling.php](./examples/file-handling.php) | 파일 열기, 쓰기, 읽기 |
| 07 | 폼 데이터 처리 | [form-handling.php](./examples/form-handling.php) | GET/POST 처리 예제 |
| 08 | 세션과 쿠키 | [session-cookie.php](./examples/session-cookie.php) | 세션 관리 및 쿠키 설정 예제 |
| 09 | 예외 처리 | [exceptions.php](./examples/exceptions.php) | try-catch, 사용자 정의 예외 |
| 10 | 데이터베이스 연동 | [database.php](./examples/database.php) | PDO로 DB 연결 및 쿼리 실행 예제 |

---

## 📋 프로젝트 목록

| 번호 | 프로젝트명 |  폴더명 |설명 |
|---|---|---|---|
| 01 | 대기표 발급 및 호출 시스템 | [queue-ticket](./projects/queue-ticket) | 키오스크에서 대기표를 접수하고, 관리자가 고객번호를 호출하면 대기실 모니터 화면에 호출된 번호가 표시되는 시스템 |
| 02 | 메뉴 접수 API | [menu-order-api](./projects/menu-order-api) | 방을 생성하여 개인 메뉴를 접수하는 RESTful API |

---

## 📢 업데이트 로그
- 2025-03-06: 초기 구성 완료
- 2025-03-09: 초기 notes 추가
- 2025-03-10: 초기 examples 추가
- 2025-03-19: 구성 수정
- 새로운 예제 추가 시 목록 업데이트 예정
