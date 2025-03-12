# 🏗️ Go Examples

이 폴더는 **Go 언어 학습 과정에서 작성한 기능별 예제 코드**를 모아둔 공간입니다.  
기본 문법, 주요 기능, 실무에서 자주 쓰이는 패턴까지 폭넓게 다룹니다.

---

## 📋 예제 목록

| 번호 | 주제 | 파일명 | 설명 |
|---|---|---|---|
| 01 | 기본 입출력 | [basic-io.go](./basic-io.go) | 터미널 입력/출력 처리 예제 |
| 02 | 함수와 리턴 | [functions.go](./functions.go) | 함수 선언, 다중 리턴, 익명 함수 활용 |
| 03 | 슬라이스와 맵 | [slices-maps.go](./slices-maps.go) | 슬라이스 생성, 맵 활용법 예제 |
| 04 | 구조체와 메서드 | [structs-methods.go](./structs-methods.go) | 구조체 정의, 메서드 추가 예제 |
| 05 | 에러 처리 패턴 | [error-handling.go](./error-handling.go) | 에러 반환, wrapping, fmt.Errorf 활용 |
| 06 | Goroutine과 채널 | [concurrency.go](./concurrency.go) | 비동기 처리, 채널 통신 기본 예제 |
| 07 | 파일 입출력 | [file-io.go](./file-io.go) | 파일 읽기/쓰기, 에러 핸들링 예제 |
| 08 | 간단한 HTTP 서버 | [http-server.go](./http-server.go) | net/http 기반 서버 구현 예제 |
| 09 | 유닛 테스트 | [unit-test-example_test.go](./unit-test-example_test.go) | 테스트 작성 및 실행 예제 |
| 10 | JSON 파싱 | [json-parsing.go](./json-parsing.go) | JSON 인코딩/디코딩 예제 |

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
