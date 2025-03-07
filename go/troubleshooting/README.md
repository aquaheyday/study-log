# 🐛 Go Troubleshooting

이 폴더는 Go 개발 과정에서 발생한 **에러 및 문제 상황**과 그에 대한 원인 분석, 해결 방법을 정리하는 공간입니다.  
각 문제는 독립적인 파일로 관리하며, 아래 흐름에 맞춰 작성합니다:

**문제 상황 → 원인 분석 → 해결 방법 → 추가 학습 포인트**

---

## 📋 문제 목록

| 번호 | 문제 요약 | 파일명 | 링크 |
|---|---|---|---|
| 01 | Go 모듈 초기화 실패 | [go-mod-init-error.md](./go-mod-init-error.md) | [바로가기](./go-mod-init-error.md) |
| 02 | go run 시 패키지 인식 불가 | [package-not-found.md](./package-not-found.md) | [바로가기](./package-not-found.md) |
| 03 | 포트 바인딩 충돌 | [port-binding-error.md](./port-binding-error.md) | [바로가기](./port-binding-error.md) |
| 04 | JSON 파싱 에러 | [json-parsing-error.md](./json-parsing-error.md) | [바로가기](./json-parsing-error.md) |
| 05 | Goroutine 데드락 | [goroutine-deadlock.md](./goroutine-deadlock.md) | [바로가기](./goroutine-deadlock.md) |
| 06 | 파일 권한 문제 | [file-permission-error.md](./file-permission-error.md) | [바로가기](./file-permission-error.md) |

---

## 📑 작성 가이드
- 파일명은 `케밥 케이스`로 작성 (ex: `go-mod-init-error.md`)
- 각 문제는 템플릿(`troubleshooting-template.md`) 기반으로 작성
- 에러 로그/캡처/메시지는 원본 그대로 포함
- 해결 과정에서 참고한 공식 문서/블로그 링크를 하단에 정리
- **재발 방지를 위한 교훈 및 팁**도 추가

---

## 📚 참고 자료
- [Go 공식 문서](https://go.dev/doc/)
- [Go Issue Tracker](https://github.com/golang/go/issues)
- [Awesome Go](https://github.com/avelino/awesome-go)
