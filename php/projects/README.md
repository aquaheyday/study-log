# 🚀 PHP Projects

이 폴더는 PHP 학습 과정에서 진행한 **실습 프로젝트**를 보관하는 공간입니다.  
기능별 단위 예제는 `examples/`에서 관리하고, 이곳은 **조금 더 완성도 있는 작은 프로젝트**를 저장합니다.  
각 프로젝트는 독립 실행 가능하며, 프로젝트별 구성/설명은 각각의 `README.md`에서 확인할 수 있습니다.

---

## 📋 프로젝트 목록

| 번호 | 프로젝트명 | 설명 | 링크 |
|---|---|---|---|
| 01 | 미니 게시판 | 기본 CRUD 기능 구현 | [mini-board](./mini-board) |
| 02 | 간단한 TODO 앱 | 세션 기반 할일 관리 | [todo-app](./todo-app) |
| 03 | 회원가입 & 로그인 | 세션/쿠키 활용 로그인 구현 | [auth-system](./auth-system) |
| 04 | 파일 업로드 기능 | 업로드/저장/확장자 검증 | [file-uploader](./file-uploader) |
| 05 | 외부 API 연동 | 공공 데이터 API 연동 실습 | [api-client](./api-client) |

---

## 📁 프로젝트 디렉터리 구조 예시

```text
projects/
├── mini-board/
│   ├── index.php
│   ├── view.php
│   ├── write.php
│   ├── delete.php
│   ├── db.php
│   └── README.md
├── todo-app/
│   ├── index.php
│   ├── add.php
│   ├── delete.php
│   ├── session.php
│   └── README.md
├── auth-system/
│   ├── login.php
│   ├── register.php
│   ├── logout.php
│   ├── auth.php
│   └── README.md
├── file-uploader/
│   ├── upload.php
│   ├── form.php
│   ├── file_check.php
│   └── README.md
├── api-client/
│   ├── fetch_data.php
│   ├── config.php
│   ├── process.php
│   └── README.md
├── laravel/
|   ├── project/
|
