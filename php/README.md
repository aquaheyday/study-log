# 🐘 PHP Archive

이 폴더는 **PHP 학습 및 실습 자료**를 보관하는 공간입니다.  
기본 문법, 주요 기능, 예제 코드, 실습 프로젝트를 단계적으로 정리하며, 학습하면서 새로 알게 된 내용과 트러블슈팅 경험도 함께 기록합니다.

---

## 📂 디렉토리 구성

| 폴더명 | 설명 |
|---|---|
| [notes](./notes) | PHP 개념 정리 및 학습 노트 |
| [examples](./examples) | 문법/기능별 예제 코드 모음 |
| [projects](./projects) | 실습 프로젝트 (미니 프로젝트, 클론코딩 등) |
| [troubleshooting](./troubleshooting) | 에러 및 트러블슈팅 기록 |
| [templates](./templates) | 새 파일 작성 시 사용하는 템플릿 모음 |

---

## 📖 추천 학습 흐름

| 단계 | 학습 주제 | 링크 |
|---|---|---|
| 1️⃣ | 기본 문법 (변수, 자료형, 제어문) | [basic-syntax.md](./notes/basic-syntax.md) |
| 2️⃣ | 함수와 스코프 | [functions.md](./notes/functions.md) |
| 3️⃣ | 배열과 연관 배열 | [arrays.md](./notes/arrays.md) |
| 4️⃣ | 폼 처리와 슈퍼글로벌 | [forms.md](./notes/forms.md) |
| 5️⃣ | 파일 입출력 | [file-handling.md](./notes/file-handling.md) |
| 6️⃣ | 세션과 쿠키 | [session-cookie.md](./notes/session-cookie.md) |
| 7️⃣ | 데이터베이스 연동 (PDO) | [database.md](./notes/database.md) |
| 8️⃣ | 에러 처리와 예외 | [error-handling.md](./notes/error-handling.md) |
| 9️⃣ | 간단한 MVC 패턴 실습 | [mvc.md](./notes/mvc.md) |
| 🔟 | 트러블슈팅 모음 | [troubleshooting.md](./troubleshooting/troubleshooting.md) |

---

## 📂 기본 폴더 구조 예시

```text
php/
├── README.md                    # 현재 파일
├── notes/                        # 개념 정리
│   ├── basic-syntax.md
│   ├── functions.md
│   └── ...
├── examples/                     # 문법별 예제 코드
│   ├── variables.php
│   ├── loops.php
│   └── ...
├── projects/                     # 실습 프로젝트
│   ├── mini-board/
│   │   ├── README.md
│   │   ├── index.php
│   │   └── ...
│   ├── simple-crud/
│   │   ├── README.md
│   │   ├── index.php
│   │   └── ...
├── troubleshooting/              # 문제 해결 기록
│   ├── file-upload-error.md
│   ├── session-timeout.md
│   └── ...
└── templates/                    # 작성 템플릿
    ├── note-template.md
    ├── example-template.md
    ├── project-readme-template.md
    └── troubleshooting-template.md

