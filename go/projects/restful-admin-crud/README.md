# 🛠️ Admin RESTful CRUD API

Go 언어 기반으로 개발된 관리자용 RESTful CRUD API 서비스입니다.  
Swagger 문서를 제공하여 API 테스트 및 연동이 용이합니다.

---

## 📌 프로젝트 소개

### 주요 기능
| 기능 | 설명 |
|---|---|
| 사용자 관리 | 사용자 계정 생성, 수정, 삭제, 조회 |
| 권한 관리 | 관리자 및 일반 사용자 권한 부여 |
| 데이터 CRUD | 리소스 생성(Create), 조회(Read), 수정(Update), 삭제(Delete) |
| 인증 및 인가 | JWT 기반 사용자 인증 및 권한 관리 |
| API 문서 제공 | Swagger UI를 활용한 API 문서 제공 |

---

## 🛠️ 기술 스택

| 구분 | 사용 기술 |
|---|---|
| Backend | Go (Golang), Gin |
| Database | PostgreSQL / MySQL |
| API 문서 | Swagger (swaggo/gin-swagger) |
| 인증 | JWT (jsonwebtoken) |
| 기타 | Docker, Docker Compose |

---

## 📑 API 문서 (Swagger)

Swagger를 사용하여 API 테스트 및 문서를 제공합니다.

- **Swagger UI 접근 방법**
  ```bash
  http://localhost:8080/swagger/index.html
  ```

Swagger 문서는 `swag init`을 통해 자동 생성할 수 있습니다.

- **Swagger 문서 생성**
  ```bash
  swag init
  ```

---

## 📸 기능 시연 이미지

### 1. Swagger API 문서
- Swagger UI에서 API 목록을 확인하고 테스트할 수 있습니다.


### 2. 사용자 목록 조회 API
- `/api/users` 엔드포인트를 통해 사용자 정보를 조회합니다.


### 3. 사용자 생성 API
- `/api/users` 엔드포인트를 통해 새 사용자를 등록합니다.


### 4. 인증 (JWT 로그인)
- 로그인 후 발급된 JWT 토큰을 통해 API 요청을 수행할 수 있습니다.


---

## 📦 설치 및 실행 방법

### 1. 클론 및 환경 설정
```bash
git clone https://github.com/aquaheyday/admin-restful-crud.git
cd admin-restful-crud
```

### 2. 환경 변수 설정 (`.env`)
`.env` 파일을 생성하고 데이터베이스 및 JWT 설정을 입력합니다.

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=admin
DB_PASS=secret
DB_NAME=admin_db
JWT_SECRET=mysecretkey
```

### 3. Docker 컨테이너 실행
```bash
docker-compose up -d
```

### 4. 애플리케이션 실행
```bash
go run main.go
```

---

## 📌 API 엔드포인트

| 메서드 | 엔드포인트 | 설명 |
|---|---|---|
| **POST** | `/api/login` | 사용자 로그인 (JWT 토큰 발급) |
| **GET** | `/api/users` | 모든 사용자 조회 |
| **POST** | `/api/users` | 새 사용자 등록 |
| **PUT** | `/api/users/:id` | 사용자 정보 수정 |
| **DELETE** | `/api/users/:id` | 사용자 삭제 |
