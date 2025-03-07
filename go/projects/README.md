# 🚀 Go Projects

이 폴더는 **Go 언어 학습 과정에서 진행한 실습 프로젝트**를 모아둔 공간입니다.  
각 프로젝트는 학습 주제를 실제로 적용해보는 형태로 구성되며,  
프로젝트별로 독립적인 폴더와 `README.md`를 갖습니다.

---

## 📋 프로젝트 목록

| 번호 | 프로젝트명 | 폴더명 | 설명 |
|---|---|---|---|
| 01 | CLI TODO 앱 | [cli-todo](./cli-todo) | 터미널 기반 할 일 관리 앱 |
| 02 | 파일 업로더 서버 | [file-uploader](./file-uploader) | 파일 업로드 및 저장 서버 구현 |
| 03 | 간단 REST API 서버 | [simple-rest-api](./simple-rest-api) | CRUD 기능을 갖춘 REST API 서버 |
| 04 | Goroutine 스크래퍼 | [concurrent-scraper](./concurrent-scraper) | Goroutine과 채널을 활용한 웹 스크래퍼 |
| 05 | JSON Config 파서 | [json-config-parser](./json-config-parser) | 설정 파일 파싱 및 검증 툴 |
| 06 | 유닛 테스트 연습 프로젝트 | [test-driven-project](./test-driven-project) | 테스트 주도 개발(TDD)로 기능 구현 연습 |
| 07 | HTTP 클라이언트 툴 | [http-client](./http-client) | 외부 API 요청 및 응답 처리 도구 |
| 08 | Go 모듈 매니저 | [go-module-manager](./go-module-manager) | 모듈 리스트 분석 및 관리 툴 |

---

## 📑 작성 가이드
- 각 프로젝트는 **독립 폴더**로 관리
- 폴더 내에 다음 파일 구성 권장
    - `main.go` (프로그램 진입점)
    - `README.md` (프로젝트 설명 및 실행 방법)
    - 필요 시 `internal`, `pkg` 등 구조 활용
- 프로젝트별 `README.md`에는 다음 정보 포함 추천:
    - 프로젝트 개요
    - 기능 및 실행 방법
    - 학습 포인트
    - 관련 참고 자료

---

## 📚 참고 자료
- [Go 공식 문서](https://go.dev/doc/)
- [Awesome Go 프로젝트 모음](https://github.com/avelino/awesome-go)
- [Go by Example](https://gobyexample.com/)

