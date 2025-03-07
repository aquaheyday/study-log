# 📝 Go Notes

이 폴더는 Go 언어 학습 과정에서 정리한 **개념 노트**를 보관하는 공간입니다.  
기본 문법부터 실무에서 자주 쓰이는 패턴, 성능 최적화, 동시성 제어 등 다양한 주제를 체계적으로 정리합니다.  
각 노트는 독립적으로 읽을 수 있도록 작성하며, 필요 시 예제 코드는 `examples/`와 연결할 수 있습니다.

---

## 📋 목차

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

## 📝 작성 가이드
- 파일명은 `케밥 케이스`로 작성 (ex: `basic-syntax.md`)
- 각 노트는 다음 형식으로 작성 권장:
    - 개념 정의
    - 기본 문법 및 예제 코드
    - 실전 팁 & 주의사항
    - 관련 링크 (공식 문서, 블로그 등)
- 코드 예제는 `examples/` 폴더에 별도 파일로 저장하고, 노트에서 링크로 연결

---

## 📚 참고 자료
- [Go 공식 문서](https://go.dev/doc/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Awesome Go](https://github.com/avelino/awesome-go)
