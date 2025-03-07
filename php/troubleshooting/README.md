# 🐛 PHP Troubleshooting

이 폴더는 PHP 개발 중 발생한 **에러 및 문제 상황**과 그에 대한 원인 분석, 해결 방법을 정리하는 공간입니다.  
하나의 문제마다 독립적인 파일로 관리하며, 각 파일은 문제 상황과 해결 과정을 쉽게 파악할 수 있도록 `문제 → 원인 → 해결 → 추가 학습` 흐름으로 작성합니다.

---

## 📋 문제 목록

| 번호 | 문제 요약 | 파일명 | 링크 |
|---|---|---|---|
| 01 | 파일 업로드 시 Permission 에러 | [file-upload-permission.md](./file-upload-permission.md) | [바로가기](./file-upload-permission.md) |
| 02 | 한글 파일명 인코딩 깨짐 | [filename-encoding.md](./filename-encoding.md) | [바로가기](./filename-encoding.md) |
| 03 | 세션이 유지되지 않는 문제 | [session-loss.md](./session-loss.md) | [바로가기](./session-loss.md) |
| 04 | PDO 접속 에러 | [pdo-connection-error.md](./pdo-connection-error.md) | [바로가기](./pdo-connection-error.md) |
| 05 | 헤더 이미 전송됨 에러 | [headers-already-sent.md](./headers-already-sent.md) | [바로가기](./headers-already-sent.md) |

---

## 📑 작성 가이드
- 파일명은 `스네이크케이스`로 작성 (ex: `file_upload_permission.md`)
- 각 문제는 템플릿(`troubleshooting-template.md`) 기반으로 작성
- 로그, 에러 메시지 등은 가능하면 원본 그대로 포함
- 참고한 자료 링크는 하단에 정리
- 재발 방지를 위한 교훈까지 기록

---

## 📂 디렉토리 구성 예시
```text
troubleshooting/
├── README.md                            # 현재 파일
├── file-upload-permission.md            # 문제별 정리 파일
├── filename-encoding.md
├── session-loss.md
├── pdo-connection-error.md
└── headers-already-sent.md
```

---

## 📚 참고 자료
- 


