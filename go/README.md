# 🐹 Go

이 폴더는 **Go 학습 과정**에서 정리한 자료, 예제 코드, 프로젝트를 보관하는 공간입니다.  
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
| 01 | 기본 문법 | [basic-syntax.md](./basic-syntax.md) | 변수, 자료형, 제어문 등 기초 문법 정리 |
| 02 | 함수와 메서드 | [functions.md](./functions.md) | 함수 선언, 리턴값, 메서드 정의 및 활용법 |
| 03 | 패키지와 모듈 | [packages.md](./packages.md) | Go 모듈 관리, 패키지 구성법 |
| 04 | 배열과 슬라이스 | [slices.md](./slices.md) | 배열과 슬라이스 차이, 슬라이스 다루기 |
| 05 | 맵(Map)과 구조체 | [maps-structs.md](./maps-structs.md) | 맵과 구조체 활용법 |
| 06 | 에러 처리 패턴 | [error-handling.md](./error-handling.md) | 기본 에러 처리부터 커스텀 에러까지 |
| 07 | 동시성 기본 | [concurrency.md](./concurrency.md) | Goroutine, 채널, WaitGroup |
| 08 | 파일 입출력 | [file-io.md](./file-io.md) | 파일 읽기/쓰기, 에러 처리 패턴 |
| 09 | HTTP 서버 만들기 | [http-server.md](./http-server.md) | net/http로 간단한 서버 구현 |
| 10 | 테스트와 벤치마크 | [testing.md](./testing.md) | Go 기본 테스트 패턴 및 성능 테스트 |
| 11 | 성능 최적화 팁 | [performance.md](./performance.md) | 메모리 프로파일링, GC 튜닝 등 |

---

## 📋 예제 목록

| 번호 | 주제 | 파일명 | 설명 |
|---|---|---|---|
| 01 | 변수와 자료형 | [variables.go](./variables.go) | 변수 선언, 상수, 타입 변환 |
| 02 | 제어문 | [control_flow.go](./control_flow.go) | 조건문(if/switch) & 반복문(for) |
| 03 | 함수 | [functions.go](./functions.go) | 함수 선언, 다중 리턴, 익명 함수 |
| 04 | 포인터 | [pointers.go](./pointers.go) | 포인터 개념과 활용법 |
| 05 | 슬라이스 | [slices.go](./slices.go) | 슬라이스 생성, 추가, 복사 |
| 06 | 맵(Map) | [maps.go](./maps.go) | 맵 생성, 조회, 수정, 삭제 |
| 07 | 구조체 | [structs.go](./structs.go) | 구조체 정의 및 활용 |
| 08 | 인터페이스 | [interfaces.go](./interfaces.go) | 인터페이스 개념 및 다형성 |
| 09 | 기본 입출력 | [basic_io.go](./basic_io.go) | 터미널에서 입력/출력 처리 |
| 10 | 파일 입출력 | [file_io.go](./file_io.go) | 파일 읽기/쓰기, 에러 처리 |
| 11 | JSON 처리 | [json_parsing.go](./json_parsing.go) | JSON 인코딩/디코딩 예제 |
| 12 | Goroutine | [goroutine.go](./goroutine.go) | 기본 Goroutine 사용법 |
| 13 | 채널 | [channels.go](./channels.go) | 채널을 활용한 동시성 제어 |
| 14 | WaitGroup | [waitgroup.go](./waitgroup.go) | 여러 Goroutine 동기화 |
| 15 | HTTP 서버 | [http_server.go](./http_server.go) | 기본 HTTP 서버 구현 |
| 16 | 유닛 테스트 | [unit_example_test.go](./unit_example_test.go) | 기본 테스트 코드 작성 |
| 17 | 벤치마크 테스트 | [benchmark_test.go](./benchmark_test.go) | 성능 테스트 예제 |

---

## 📋 프로젝트 목록

| 번호 | 프로젝트명 | 폴더명 | 설명 |
|---|---|---|---|
| 01 | RESTful 관리자 CRUD API + Swagger | [restful-admin-crud](./restful-admin-crud) | 관리자 CRUD + 스웨거 적용한 RESTful API |

---

## 📢 업데이트 로그
- 2025-03-07: 초기 구성
- 2025-03-12: notes 초기 내용
- 2025-03-13: examples 초기 내용, projects > restful-admin-crud 내용
- 2025-03-19: 디렉토리 구성 수정
- 새로운 예제 추가 시 목록 업데이트 예정
