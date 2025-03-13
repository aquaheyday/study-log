# 🏗️ Go Examples

이 폴더는 **Go 언어 학습 과정에서 작성한 기능별 예제 코드**를 모아둔 공간입니다.  
기본 문법, 주요 기능, 실무에서 자주 쓰이는 패턴까지 폭넓게 다룹니다.

---

## 📋 예제 목록

| 번호 | 주제 | 파일명 | 설명 |
|---|---|---|---|
| 01 | 변수와 자료형 | [variables.go](./variables.go) | 변수 선언, 상수, 타입 변환 |
| 02 | 제어문 | [control-flow.go](./control-flow.go) | 조건문(if/switch) & 반복문(for) |
| 03 | 함수 | [functions.go](./functions.go) | 함수 선언, 다중 리턴, 익명 함수 |
| 04 | 포인터 | [pointers.go](./pointers.go) | 포인터 개념과 활용법 |
| 05 | 슬라이스 | [slices.go](./slices.go) | 슬라이스 생성, 추가, 복사 |
| 06 | 맵(Map) | [maps.go](./maps.go) | 맵 생성, 조회, 수정, 삭제 |
| 07 | 구조체 | [structs.go](./structs.go) | 구조체 정의 및 활용 |
| 08 | 인터페이스 | [interfaces.go](./interfaces.go) | 인터페이스 개념 및 다형성 |
| 09 | 기본 입출력 | [basic-io.go](./basic-io.go) | 터미널에서 입력/출력 처리 |
| 10 | 파일 입출력 | [file-io.go](./file-io.go) | 파일 읽기/쓰기, 에러 처리 |
| 11 | JSON 처리 | [json-parsing.go](./json-parsing.go) | JSON 인코딩/디코딩 예제 |
| 12 | Goroutine | [goroutine.go](./goroutine.go) | 기본 Goroutine 사용법 |
| 13 | 채널 | [channels.go](./channels.go) | 채널을 활용한 동시성 제어 |
| 14 | WaitGroup | [waitgroup.go](./waitgroup.go) | 여러 Goroutine 동기화 |
| 15 | HTTP 서버 | [http-server.go](./http-server.go) | 기본 HTTP 서버 구현 |
| 16 | 유닛 테스트 | [unit-test-example_test.go](./unit-test-example_test.go) | 기본 테스트 코드 작성 |
| 17 | 벤치마크 테스트 | [benchmark_test.go](./benchmark_test.go) | 성능 테스트 예제 |

---

## 📝 작성 가이드
- 파일명은 **스네이크 케이스 + `.go`** 로 작성 (ex: `basic_io.go`)
- 각 예제는 **단일 main 패키지로 구성** 해서 독립 실행 가능하도록 작성
- 필요 시 관련 노트(`notes/`)와 상호 링크 추가
- 주석을 충분히 작성해, 해당 예제의 포인트를 명확하게 설명

---

## 📚 참고 자료
- [Go 공식 문서](https://go.dev/doc/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example](https://gobyexample.com/)
