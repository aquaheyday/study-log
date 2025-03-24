# 🛠️ PHP 개발 환경 설정

PHP 개발을 시작하려면 **PHP 실행 환경**, **웹 서버(Apache/Nginx)**, **데이터베이스(MySQL/PostgreSQL)** 등을 설정해야 합니다.  
이 문서에서는 **Windows, macOS, Linux에서 PHP 개발 환경을 설정하는 방법**을 정리합니다.  

---

## 1️⃣ PHP 설치

### 1) Windows에서 PHP 설치

#### 1. XAMPP 설치 (추천)
XAMPP는 **PHP, Apache, MySQL을 한 번에 설치**할 수 있는 패키지입니다.

- [XAMPP 다운로드](https://www.apachefriends.org/index.html) 후 설치
- XAMPP Control Panel에서 **Apache & MySQL 실행**

#### 2. 개별 PHP 설치
1. [PHP 공식 사이트](https://windows.php.net/download/)에서 ZIP 파일 다운로드  
2. `C:\php`에 압축 해제  
3. 환경 변수 설정 (`C:\php`을 시스템 PATH에 추가)  
4. 명령 프롬프트에서 버전 확인
```sh
php -v
```

---

### 2) macOS에서 PHP 설치

macOS에는 기본적으로 PHP가 설치되어 있지만 최신 버전을 사용하는 것이 좋습니다.

#### 1. Homebrew로 PHP 설치
```sh
brew install php
```
✔ `php -v` 로 설치 확인  

#### 2. 내장 PHP 사용 (버전 확인)
```sh
php -v
```
✔ macOS 기본 PHP 버전 확인  

---

### 3) Linux에서 PHP 설치

#### 1. Ubuntu/Debian
```sh
sudo apt update
sudo apt install php php-cli php-mysql php-curl php-xml
```

#### 2. CentOS/RHEL
```sh
sudo yum install php php-cli php-mysql php-curl php-xml
```

✔ 설치 후 `php -v` 로 버전 확인  

---

## 2️⃣ PHP 개발 필수 도구

| 도구 | 설명 |
|------|------|
| **Apache / Nginx** | 웹 서버 (PHP 실행) |
| **MySQL / PostgreSQL** | 데이터베이스 연동 |
| **Composer** | PHP 패키지 관리 |
| **VS Code / PhpStorm** | 추천 PHP 개발 환경 (IDE) |

---

## 3️⃣ PHP 내장 서버 실행 (테스트용)

PHP는 내장 서버를 제공하여 **Apache 없이 로컬 개발 가능**합니다.

#### 1. 내장 서버 실행
```sh
php -S localhost:8000
```
✔ `localhost:8000`에서 PHP 실행 가능  

#### 2. 특정 디렉토리에서 실행
```sh
php -S localhost:8000 -t public
```
✔ `public/` 디렉토리를 웹 루트로 설정  

---

## 4️⃣ PHP 패키지 관리 (`Composer`)

**Composer**는 PHP의 **의존성 관리 도구**로, Laravel, Symfony 같은 프레임워크 설치에 필수적입니다.

#### 1. Composer 설치
Windows
- [Composer 공식 사이트](https://getcomposer.org/)에서 설치 프로그램 실행  

macOS / Linux
```sh
curl -sS https://getcomposer.org/installer | php
sudo mv composer.phar /usr/local/bin/composer
```

#### 2. Composer 버전 확인
```sh
composer -V
```

---

## 5️⃣ PHP 프로젝트 설정 (Composer 활용)

#### 1. 프로젝트 생성
```sh
composer init
```
✔ `composer.json` 파일 생성 (패키지 정보 설정)  

#### 2. 패키지 설치 (예: Laravel)
```sh
composer create-project laravel/laravel myapp
```

#### 3. 프로젝트 종속성 업데이트
```sh
composer update
```

✔ 프로젝트에서 사용 중인 모든 패키지를 최신 버전으로 업데이트  

---

## 6️⃣ PHP 확장 모듈 설치 및 설정

PHP는 다양한 기능을 제공하는 **확장 모듈(Extensions)** 을 지원합니다.

#### 1. PHP 확장 모듈 목록 확인
```sh
php -m
```

#### 2. 특정 모듈 활성화 (`php.ini` 편집)
1. `php.ini` 파일 열기 (예: `C:\php\php.ini` 또는 `/etc/php.ini`)  
2. 다음 줄의 주석(`;`) 제거:
```
;extension=mysqli
;extension=curl
;extension=mbstring
```
✔ `extension=mysqli` → MySQL 연동  
✔ `extension=curl` → HTTP 요청 가능  
✔ `extension=mbstring` → 다국어 문자 처리  

---

## 7️⃣ PHP 디버깅 환경 설정

디버깅을 위해 **Xdebug** 또는 PHP 내장 디버거를 사용할 수 있습니다.

### 1) Xdebug 설치
```sh
pecl install xdebug
```

---

### 2) VS Code에서 PHP Debug 설정
#### 1. **PHP Debug 확장 설치**  
#### 2. `.vscode/launch.json` 설정 추가
```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Listen for Xdebug",
            "type": "php",
            "request": "launch",
            "port": 9003
        }
    ]
}
```

✔ `F5` 키를 눌러 디버깅 시작  

---

## 8️⃣ PHP 테스트 환경 설정

PHP 애플리케이션의 **단위 테스트(Unit Testing)** 를 위해 PHPUnit을 사용할 수 있습니다.

#### 1. PHPUnit 설치
```sh
composer require --dev phpunit/phpunit
```

#### 2. 테스트 코드 예제 (`tests/SampleTest.php`)
```php
<?php
use PHPUnit\Framework\TestCase;

class SampleTest extends TestCase {
    public function testAddition() {
        $this->assertEquals(4, 2 + 2);
    }
}
?>
```

✔ `phpunit` 명령어로 실행 가능  

---

## 9️⃣ PHP 서버 배포 방법
Apache 또는 Nginx에서 PHP 실행 설정 방법

### 1) Apache (`httpd.conf` 설정)
```
<IfModule php_module>
    AddHandler application/x-httpd-php .php
</IfModule>
```

---

### 2) Nginx (`nginx.conf` 설정)
```
location ~ \.php$ {
    fastcgi_pass unix:/run/php/php7.4-fpm.sock;
    include fastcgi_params;
}
```

✔ Apache 또는 Nginx 재시작 후 PHP 실행 가능  

---

## 🔟 PHP 실행 최적화 (성능 개선)

### OPcache 활성화
#### 1. `php.ini` 파일에서 다음 설정 추가:
```
opcache.enable=1
opcache.memory_consumption=128
opcache.max_accelerated_files=4000
```
#### 2. 웹 서버 재시작  

✔ OPcache는 **PHP 코드 실행 속도를 크게 향상**시킴.  

---

## 🎯 정리

✔ **PHP 설치** → XAMPP, Homebrew, apt/yum 사용  
✔ **내장 서버 실행** → `php -S localhost:8000`  
✔ **Composer 활용** → 패키지 설치 및 관리  
✔ **PHP 확장 모듈 설정** → `php.ini` 수정  
✔ **Xdebug 디버깅** → VS Code 연동 가능  
✔ **PHP 배포** → Apache, Nginx 설정  
