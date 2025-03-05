# 🐛 Python Troubleshooting

이 폴더는 Python 개발 중 발생한 **에러 및 문제 상황**과 그에 대한 원인 분석, 해결 방법을 정리하는 공간입니다.  
하나의 문제마다 독립적인 파일로 관리하며, 각 파일은 문제 상황과 해결 과정을 쉽게 파악할 수 있도록  
`문제 상황 → 원인 분석 → 해결 방법 → 추가 학습` 흐름으로 작성합니다.

---

## 📋 문제 목록

| 번호 | 문제 요약 | 파일명 | 링크 |
|---|---|---|---|
| 01 | 모듈 임포트 에러 (ModuleNotFoundError) | [module-not-found.md](./module-not-found.md) | [바로가기](./module-not-found.md) |
| 02 | 파일 인코딩 문제 (UnicodeDecodeError) | [file-encoding-error.md](./file-encoding-error.md) | [바로가기](./file-encoding-error.md) |
| 03 | 가상환경 패키지 불일치 | [venv-package-mismatch.md](./venv-package-mismatch.md) | [바로가기](./venv-package-mismatch.md) |
| 04 | 데이터베이스 연결 실패 (sqlite/mysql) | [db-connection-fail.md](./db-connection-fail.md) | [바로가기](./db-connection-fail.md) |
| 05 | 터미널 출력 깨짐 (print 한글깨짐) | [print-encoding-error.md](./print-encoding-error.md) | [바로가기](./print-encoding-error.md) |

---

## 📑 작성 가이드

- 파일명은 `스네이크케이스`로 작성 (ex: `module_not_found.md`)
- 각 문제는 템플릿(`troubleshooting-template.md`) 기반으로 작성
- 에러 로그, 트레이스백 등은 원본 그대로 포함하는 것을 권장
- 해결 과정에서 참고한 자료는 하단에 링크 정리
- 재발 방지를 위한 팁이나 교훈까지 기록하면 더 좋음

---

## 📂 디렉터리 구성 예시

```text
troubleshooting/
├── README.md                         # 현재 파일
├── module-not-found.md                # 문제별 정리 파일
├── file-encoding-error.md
├── venv-package-mismatch.md
├── db-connection-fail.md
└── print-encoding-error.md

