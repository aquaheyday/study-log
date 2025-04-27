# 📦 Laravel Application with Docker Compose

Docker로 컨테이너화하고 Docker Compose를 사용해 관리하는 Laravel 애플리케이션 환경입니다.

---

## 🚀 실행 방법

### 1. Docker & Docker Compose 설치 확인

```bash
docker --version
docker-compose --version
```

> Docker와 Docker Compose가 설치되어 있어야 합니다.

### 2. 프로젝트 클론 및 환경 설정

```bash
git clone https://github.com/aquaheyday/study-log.git study-log/projects/docker/laravel-mysql
cd laravel-mysql
cp .env.example .env
```

- `.env` 파일을 생성한 후, 데이터베이스 및 APP_KEY 관련 환경변수를 설정합니다.

### 3. 컨테이너 빌드 및 실행

```bash
docker-compose up -d
```

- `laravel_app`, `laravel_webserver`, `laravel_mysql` 등 컨테이너가 생성됩니다.
- `localhost`로 접속하여 Laravel 애플리케이션을 확인할 수 있습니다.

---

## ⚙️ 주요 서비스 구성

### app (PHP-FPM)

- PHP 8.1 기반 `php:8.1-fpm` 이미지 사용
- 소스 코드 경로: `/var/www/html`
- Laravel 애플리케이션 구동

### webserver (Nginx)

- `nginx:stable-alpine` 이미지 사용
- 포트 80 매핑 (`localhost:80`)
- `nginx.conf` 파일을 통해 PHP 요청을 `app` 컨테이너로 프록시

### mysql (MySQL)

- `mysql:8.0` 이미지 사용
- 포트 3306 매핑
- 환경변수로 DB 설정
- 볼륨을 통해 데이터 지속성 보장 (`db_data`)

### composer

- `composer:latest` 이미지 사용
- 의존성 설치 (`composer install`)

### artisan

- PHP-FPM을 기반으로 artisan 명령어 실행

---

## 🌐 Nginx 설정 (`nginx.conf`)

```nginx
server {
    listen 80;
    server_name localhost;

    root /var/www/html/public;
    index index.php index.html;

    location / {
        try_files $uri $uri/ /index.php?$query_string;
    }

    location ~ \.php$ {
        include fastcgi_params;
        fastcgi_pass app:9000;
        fastcgi_index index.php;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
    }

    location ~ /\.ht {
        deny all;
    }
}
```

- `/public` 디렉토리를 DocumentRoot로 설정
- `.php` 파일은 `app` 컨테이너(PHP-FPM)로 전달
- `.ht*` 파일 접근 차단

---

## ✨ 추가 작업 (필수)

- `.env` 파일에서 `DB_HOST=mysql`로 설정해야 MySQL 컨테이너와 연결됩니다.
- `.env`에 `APP_KEY` 생성 필요:

```bash
docker-compose run --rm artisan key:generate
```

- Storage 링크 설정:

```bash
docker-compose run --rm artisan storage:link
```
