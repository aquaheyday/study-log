# 📝 PHP Notes

이 폴더는 PHP 학습 과정에서 정리한 **개념 노트**를 보관하는 공간입니다.  
기본 문법부터 실무에서 자주 쓰이는 패턴, 심화 내용까지 점진적으로 정리합니다.  
각 노트는 독립적으로 읽을 수 있도록 작성하며, 필요 시 예제 코드는 `examples/`와 연결할 수 있습니다.

---

## 📋 목차

| 번호 | 주제 | 파일명 | 설명 |
|---|---|---|---|
| 01 | 기본 문법 | [basic-syntax.md](./basic-syntax.md) | 변수, 자료형, 제어문 기초 정리 |
| 02 | 함수와 스코프 | [functions.md](./functions.md) | 함수 선언, 매개변수, 반환값, 스코프 개념 |
| 03 | 배열과 연관 배열 | [arrays.md](./arrays.md) | 배열 생성, 다차원 배열, 연관 배열 활용 |
| 04 | 폼 처리와 슈퍼글로벌 | [forms.md](./forms.md) | GET/POST, $_GET, $_POST 사용법 |
| 05 | 파일 입출력 | [file-handling.md](./file-handling.md) | 파일 열기, 쓰기, 읽기, 삭제 |
| 06 | 세션과 쿠키 | [session-cookie.md](./session-cookie.md) | 로그인 유지 등 상태 관리 기법 |
| 07 | 데이터베이스 연동 | [database.md](./database.md) | PDO 기반 데이터베이스 연결 및 쿼리 실행 |
| 08 | 에러 처리와 예외 | [error-handling.md](./error-handling.md) | 에러 로그, try-catch, 사용자 정의 예외 |
| 09 | 간단한 MVC 패턴 | [mvc.md](./mvc.md) | 기본적인 MVC 구조 설명 |
| 10 | 기타 유용한 팁 | [tips.md](./tips.md) | 보안 관련 팁, PHP 설정 팁 등 |

---

## 📝 작성 가이드

- 파일명은 `케밥 케이스`로 작성 (ex: `basic-syntax.md`)
- 각 노트는 다음 형식을 추천:
    - 개념 정의
    - 기본 문법 예제
    - 실전 팁 & 주의사항
    - 관련 링크 (공식 문서, 블로그 등)
- 코드 예제는 가능하면 별도 파일(`examples/`)로 관리하고, 노트에서는 링크로 연결
