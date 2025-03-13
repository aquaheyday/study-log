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

## 📊 Database Table Structure

### 1. admin_users 테이블 (관리자 사이트 사용자 정보)

```sql
CREATE TABLE `admin_users` (
	`id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
	`username` VARCHAR(50) NOT NULL,
	`email` VARCHAR(100) NOT NULL,
	`password` VARCHAR(255) NOT NULL,
	`role` VARCHAR(50) NULL DEFAULT 'admin',
	`is_active` TINYINT(1) NULL DEFAULT 1,
	`last_login` DATETIME NULL DEFAULT NULL,
	`created_at` DATETIME NULL DEFAULT current_timestamp(),
	`updated_at` DATETIME NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
	PRIMARY KEY (`id`) USING BTREE,
	UNIQUE INDEX `email` (`email`) USING BTREE
);
```

### 2. admin_user_permissions 테이블 (사용자에게 부여된 권한 정보)

```sql
CREATE TABLE `admin_user_permissions` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`email` VARCHAR(100) NOT NULL COMMENT '사용자 ID',
	`permission_id` INT(11) NOT NULL COMMENT '권한 ID',
	`created_at` TIMESTAMP NULL DEFAULT current_timestamp(),
	`updated_at` TIMESTAMP NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
	PRIMARY KEY (`id`) USING BTREE,
	INDEX `email` (`email`) USING BTREE,
	INDEX `permission_id` (`permission_id`) USING BTREE,
	CONSTRAINT `admin_user_permissions_ibfk_1` FOREIGN KEY (`email`) REFERENCES `admin_users` (`email`) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT `admin_user_permissions_ibfk_2` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
);
```

### 3. admin_user_permission_requests 테이블 (사용자가 요청한 권한 정보)

```sql
CREATE TABLE `admin_user_permission_requests` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`email` VARCHAR(100) NOT NULL COMMENT '사용자 ID',
	`permission_id` INT(11) NOT NULL COMMENT '요청된 권한 ID',
	`status` ENUM('pending','approved','rejected') NOT NULL DEFAULT 'pending' COMMENT '요청 상태',
	`requested_at` TIMESTAMP NULL DEFAULT current_timestamp() COMMENT '요청 시간',
	`approved_at` TIMESTAMP NULL DEFAULT NULL COMMENT '승인 시간',
	`rejected_at` TIMESTAMP NULL DEFAULT NULL COMMENT '거절 시간',
	`admin_email` VARCHAR(100) NULL DEFAULT NULL COMMENT '승인/거절 처리한 관리자 ID',
	PRIMARY KEY (`id`) USING BTREE,
	INDEX `email` (`email`) USING BTREE,
	INDEX `permission_id` (`permission_id`) USING BTREE,
	CONSTRAINT `admin_permission_requests_ibfk_1` FOREIGN KEY (`email`) REFERENCES `admin_users` (`email`) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT `admin_permission_requests_ibfk_2` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
);
```

### 4. permissions 테이블 (권한 목록 정보)

```sql
CREATE TABLE `permissions` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(50) NOT NULL COMMENT '권한명',
	`type` VARCHAR(50) NULL DEFAULT 'url' COMMENT '권한 타입',
	`role` ENUM('read', 'write') NOT NULL DEFAULT 'read' COMMENT 'read/write',
	`detail` VARCHAR(100) NULL DEFAULT NULL COMMENT '상세 권한 설정',
	`description` TEXT NULL DEFAULT NULL,
	`sort` INT(11) NULL DEFAULT 0 COMMENT '권한 정렬 순서',
	`created_at` TIMESTAMP NULL DEFAULT current_timestamp(),
	`updated_at` TIMESTAMP NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
	PRIMARY KEY (`id`) USING BTREE
);
```

### 5. permission_groups 테이블 (권한 목록을 그룹 단위로 묶은 정보)

```sql
CREATE TABLE `permission_groups` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(50) NOT NULL,
	`description` TEXT NULL DEFAULT NULL,
	`sort` INT(11) NULL DEFAULT 0 COMMENT '권한 정렬 순서',
	`created_at` TIMESTAMP NULL DEFAULT current_timestamp(),
	`updated_at` TIMESTAMP NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
	PRIMARY KEY (`id`) USING BTREE
);
```

### 6. menu 테이블 (메뉴 목록 정보)

```sql
CREATE TABLE `menu` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`group` VARCHAR(30) NOT NULL COMMENT '메뉴 그룹(관리자:admin)',
	`name` VARCHAR(255) NOT NULL COMMENT '메뉴명',
	`parent_id` INT(11) NULL DEFAULT NULL COMMENT '상위 메뉴 id (NULL 이면 1뎁스)',
	`url` VARCHAR(255) NULL DEFAULT NULL COMMENT '메뉴 URL 경로',
	`level` INT(11) NOT NULL DEFAULT 1 COMMENT '메뉴 레벨 (1, 2, 3)',
	`sort` INT(11) NULL DEFAULT 0 COMMENT '메뉴 정렬 순서',
	`is_active` TINYINT(1) NULL DEFAULT 1 COMMENT '활성화 여부 (1: 활성, 0: 비활성)',
	`created_at` TIMESTAMP NULL DEFAULT current_timestamp(),
	`updated_at` TIMESTAMP NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
	PRIMARY KEY (`id`) USING BTREE,
	INDEX `parent_id` (`parent_id`) USING BTREE,
	CONSTRAINT `menu_ibfk_1` FOREIGN KEY (`parent_id`) REFERENCES `menu` (`id`) ON UPDATE RESTRICT ON DELETE CASCADE
);
```

### 7. franchise_user_leads 테이블 (외부 랜딩에서 유입된 데이터 정보)

```sql
CREATE TABLE `franchise_user_leads` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(50) NULL DEFAULT NULL,
	`birth_date` DATE NULL DEFAULT NULL,
	`gender` ENUM('male', 'female', 'other') NULL DEFAULT NULL,
	`phone_number` VARCHAR(20) NULL DEFAULT NULL,
	`branch_location` VARCHAR(50) NULL DEFAULT NULL COMMENT '지점 위치',
	`inquiry` VARCHAR(255) NULL DEFAULT NULL COMMENT '문의내용',
	`status` ENUM('live','delete') NOT NULL DEFAULT 'live' COMMENT '요청 상태',
	`created_at` TIMESTAMP NOT NULL DEFAULT current_timestamp(),
	`delete_at` TIMESTAMP NULL DEFAULT NULL COMMENT '삭제 시간',
	`admin_email` VARCHAR(100) NULL DEFAULT NULL COMMENT '삭제 처리한 관리자 ID',
	PRIMARY KEY (`id`) USING BTREE,
	INDEX `idx_branch_location` (`branch_location`) USING BTREE
);
```

---

## 📑 API 문서 (Swagger)

Swagger를 사용하여 API 테스트 및 문서를 제공합니다.

- **Swagger UI 접근 방법**
  ```bash
  http://localhost:8080/v1/swagger/index.html
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
