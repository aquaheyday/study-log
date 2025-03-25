# 📄 Sample Project Name

Sample Project 설명

---

## 📌 프로젝트 소개

### 주요 기능
| 기능 | 설명 |
|---|---|
| Sample 기능 | Sample 설명 |

---

## 🛠️ 기술 스택

| 구분 | 사용 기술 |
|---|---|
| Backend | Laravel |
| Database | MySQL |
| API 문서 | Swagger |

---

## 📂 프로젝트 폴더 구조

```text
sample-project/
│── config/          # 애플리케이션 환경 설정 (CORS, 설정 파일 등)
│── database/        # 데이터베이스 연결 및 마이그레이션 관련 코드
│── docs/            # API 문서 및 스웨거(Swagger) 문서화 관련 파일
│── handlers/        # HTTP 핸들러 (컨트롤러) 계층, 요청 처리 담당
│── middleware/      # 미들웨어 (인증)
│── models/          # 데이터 모델 정의
```

---

## 📑 API 문서 (Swagger)

Swagger를 사용하여 API 테스트 및 문서를 제공합니다.

- **Swagger UI 접근 방법**
  ```bash
  http://localhost/api/v1/swagger/index.html
  ```

Swagger 문서는 `swag init`을 통해 자동 생성할 수 있습니다.

- **Swagger 문서 생성**
  ```bash
  swag init
  ```

---

## RESTful API 상태 코드 (HTTP Status Codes)
아래는 주요 HTTP 상태 코드 및 설명입니다.

### 1xx: 정보 응답 (Informational Responses)

| 상태 코드 | 설명 |
|-----------|------------------------------------------|
| `100 Continue` | 요청이 정상적으로 접수되었으며 계속 진행 가능 |
| `101 Switching Protocols` | 클라이언트의 프로토콜 변경 요청을 수락함 |

### 2xx: 성공 (Success)

| 상태 코드 | 설명 |
|-----------|------------------------------------------|
| `200 OK` | 요청이 성공적으로 처리됨 (GET, PUT, DELETE 요청) |
| `201 Created` | 리소스가 성공적으로 생성됨 (POST 요청) |
| `204 No Content` | 요청이 성공했지만, 응답 본문이 없음 (DELETE 요청) |

### 3xx: 리디렉션 (Redirection)

| 상태 코드 | 설명 |
|-----------|------------------------------------------|
| `301 Moved Permanently` | 요청한 리소스가 영구적으로 이동됨 |
| `302 Found` | 요청한 리소스가 일시적으로 이동됨 |
| `304 Not Modified` | 캐시된 데이터를 사용할 수 있음 (클라이언트의 캐시 활용) |

### 4xx: 클라이언트 오류 (Client Errors)

| 상태 코드 | 설명 |
|-----------|------------------------------------------|
| `400 Bad Request` | 요청이 잘못되었거나 형식이 올바르지 않음 |
| `401 Unauthorized` | 인증이 필요함 (JWT 토큰 없음/만료됨) |
| `403 Forbidden` | 권한이 없어 요청을 수행할 수 없음 |
| `404 Not Found` | 요청한 리소스를 찾을 수 없음 |
| `405 Method Not Allowed` | 허용되지 않은 HTTP 메서드 사용 |

### 5xx: 서버 오류 (Server Errors)

| 상태 코드 | 설명 |
|-----------|------------------------------------------|
| `500 Internal Server Error` | 서버 내부 오류 (예상치 못한 에러) |
| `502 Bad Gateway` | 게이트웨이 서버가 잘못된 응답을 받음 |
| `503 Service Unavailable` | 서버가 현재 요청을 처리할 수 없음 (과부하 상태) |

---

## 📌 API 엔드포인트 목록

| HTTP Method | 엔드포인트 | 설명 |
|------------|-----------|-------------------------------|
| `GET` | `/api/v1/sample-url` | Sample 설명 |

---

## 📸 기능 시연 이미지

### 1. Swagger API 문서
- Swagger UI에서 API 목록을 확인하고 테스트할 수 있습니다.
<img src="docs/assets/images/sample.png" width="1000">

---

## 📊 Database Table Structure

### 1. Sample 테이블 (Sample 테이블 설명)

```sql
CREATE TABLE `sample` (
	`id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
	`email` VARCHAR(50) NOT NULL,
	`created_at` DATETIME NULL DEFAULT current_timestamp(),
	`updated_at` DATETIME NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
	PRIMARY KEY (`id`) USING BTREE,
	UNIQUE INDEX `email` (`email`) USING BTREE
);
```

---

## 📦 설치 및 실행 방법

### 1. 클론 및 환경 설정
```bash
git clone https://github.com/aquaheyday/language-archive.git language-archive/sample-path/
cd sample-name
```

### 2. 환경 변수 설정 (`.env`)
`.env` 파일을 생성하고 데이터베이스 및 JWT 설정을 입력합니다.

```env
DB_USERNAME=sample_user
DB_PASSWORD=secret
DB_HOST=localhost
DB_PORT=3306
DB_NAME=sample_db

JWT_SECRET=mysecretkey

CORS_ALLOWED_ORIGINS="http://localhost:*"
```

### 3. 애플리케이션 실행
```bash
sample run
```
