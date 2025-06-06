# 🐹 Go

이 폴더는 **Go 학습 과정**에서 정리한 자료, 예제 코드, 프로젝트를 보관하는 공간입니다.  
기본 개념부터 실습 예제, 프로젝트 구성까지 단계적으로 학습한 내용을 체계적으로 정리합니다.

---

## 📂 디렉토리 구성

| 폴더명 | 설명 |
|---|---|
| [notes](./notes) | 개념 정리 및 학습 노트 |
| [examples](./examples) | 주요 기능별 예제 코드 |

---

## 📋 Go 언어 개념 정리 목록

### 📌 기본 개념
| 주제 | 파일명 | 설명 |
|---|---|---|
| Go 언어 개요 | [go-intro.md](./notes/go-intro.md) | Go의 특징, 장점, 사용 사례 |
| 개발 환경 설정 | [setup-go.md](./notes/setup-go.md) | Go 설치, GOPATH, 모듈 관리 |
| 변수와 데이터 타입 | [variables.md](./notes/variables.md) | 기본 타입, 상수, 타입 추론 |
| 연산자와 제어문 | [operators-control.md](./notes/operators-control.md) | 조건문, 반복문, switch 등 |
| 배열, 슬라이스, 맵 | [collections.md](./notes/collections.md) | 컬렉션 타입 및 사용법 |

### 🧱 함수와 구조체
| 주제 | 파일명 | 설명 |
|---|---|---|
| 함수 선언과 호출 | [functions.md](./notes/functions.md) | 다중 반환값, defer, 클로저 |
| 구조체와 메서드 | [structs.md](./notes/structs.md) | 구조체 정의, 임베딩, 메서드 |
| 인터페이스 | [interfaces.md](./notes/interfaces.md) | 인터페이스 개념과 사용법 |
| 패키지와 모듈 | [packages.md](./notes/packages.md) | 패키지 분리, go mod 관리 |

### 🔄 고급 기능
| 주제 | 파일명 | 설명 |
|---|---|---|
| 포인터와 메모리 | [pointers.md](./notes/pointers.md) | 포인터, new, make, 주소 연산 |
| 고루틴과 채널 | [concurrency.md](./notes/concurrency.md) | 병렬 처리, 동기화, select |
| 에러 처리 | [errors.md](./notes/errors.md) | error 인터페이스, custom error |
| 제네릭 (Go 1.18+) | [generics.md](./notes/generics.md) | 제네릭 함수 및 타입 정의 |

### 🌍 웹 개발
| 주제 | 파일명 | 설명 |
|---|---|---|
| 기본 웹 서버 | [http-server.md](./notes/http-server.md) | net/http로 웹 서버 구현 |
| 라우팅과 핸들러 | [router.md](./notes/router.md) | URL 라우팅, 핸들러 분리 |
| REST API 개발 | [rest-api.md](./notes/rest-api.md) | JSON 응답, CRUD API 만들기 |
| HTML 템플릿 | [html-template.md](./notes/html-template.md) | 템플릿 렌더링, XSS 보호 |

### 🗃️ 데이터 처리
| 주제 | 파일명 | 설명 |
|---|---|---|
| 파일 입출력 | [file-io.md](./notes/file-io.md) | os, ioutil, bufio 활용 |
| 데이터베이스 연동 | [database.md](./notes/database.md) | sql 패키지, MySQL, PostgreSQL |
| JSON 처리 | [json.md](./notes/json.md) | 인코딩/디코딩, struct 태그 |

### 🔐 보안 및 최적화
| 주제 | 파일명 | 설명 |
|---|---|---|
| 인증과 세션 관리 | [auth-session.md](./notes/auth-session.md) | JWT, 쿠키, 미들웨어 |
| 암호화와 해싱 | [crypto.md](./notes/crypto.md) | bcrypt, SHA256, TLS 인증 |
| 성능 최적화 | [performance.md](./notes/performance.md) | 프로파일링, GOMAXPROCS |

### 🛠️ 테스트 및 배포
| 주제 | 파일명 | 설명 |
|---|---|---|
| 단위 테스트 | [testing.md](./notes/testing.md) | `testing` 패키지, mock, benchmark |
| 애플리케이션 빌드 | [build.md](./notes/build.md) | go build, cross compile, 버전 관리 |
| 패키징과 배포 | [deployment.md](./notes/deployment.md) | Docker, systemd, CI/CD 배포 |

---

## 📋 예제 목록

| 주제 | 파일명 | 설명 |
|---|---|---|
| 변수와 자료형 | [variables.go](./examples/variables.go) | 변수 선언, 상수, 타입 변환 |
| 제어문 | [control_flow.go](./examples/control_flow.go) | 조건문(if/switch) & 반복문(for) |
| 함수 | [functions.go](./examples/functions.go) | 함수 선언, 다중 리턴, 익명 함수 |
| 포인터 | [pointers.go](./examples/pointers.go) | 포인터 개념과 활용법 |
| 슬라이스 | [slices.go](./examples/slices.go) | 슬라이스 생성, 추가, 복사 |
| 맵(Map) | [maps.go](./examples/maps.go) | 맵 생성, 조회, 수정, 삭제 |
| 구조체 | [structs.go](./examples/structs.go) | 구조체 정의 및 활용 |
| 인터페이스 | [interfaces.go](./examples/interfaces.go) | 인터페이스 개념 및 다형성 |
| 기본 입출력 | [basic_io.go](./examples/basic_io.go) | 터미널에서 입력/출력 처리 |
| 파일 입출력 | [file_io.go](./examples/file_io.go) | 파일 읽기/쓰기, 에러 처리 |
| JSON 처리 | [json_parsing.go](./examples/json_parsing.go) | JSON 인코딩/디코딩 예제 |
| Goroutine | [goroutine.go](./examples/goroutine.go) | 기본 Goroutine 사용법 |
| 채널 | [channels.go](./examples/channels.go) | 채널을 활용한 동시성 제어 |
| WaitGroup | [waitgroup.go](./examples/waitgroup.go) | 여러 Goroutine 동기화 |
| HTTP 서버 | [http_server.go](./examples/http_server.go) | 기본 HTTP 서버 구현 |
| 유닛 테스트 | [unit_example_test.go](./examples/unit_example_test.go) | 기본 테스트 코드 작성 |
| 벤치마크 테스트 | [benchmark_test.go](./examples/benchmark_test.go) | 성능 테스트 예제 |

---

## 📚 참고 자료
- [Go 공식 문서](https://go.dev/doc/) - Go 언어의 공식 문서 및 가이드  
- [Go API 레퍼런스](https://pkg.go.dev/std) - Go 표준 라이브러리 및 API 문서  
- [A Tour of Go](https://go.dev/tour/) - Go의 핵심 개념을 배울 수 있는 공식 튜토리얼  
- [Effective Go](https://go.dev/doc/effective_go) - Go 스타일 가이드 및 모범 사례  
