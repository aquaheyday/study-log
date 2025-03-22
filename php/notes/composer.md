# 📦 PHP Composer 패키지 관리

**Composer**는 PHP의 공식적인 패키지 관리 도구입니다.  
라이브러리 설치, 의존성 관리, 오토로딩 자동 설정 등 **현대적인 PHP 개발에 필수적인 도구**입니다.

---

## 1️⃣ Composer 설치 방법

### 1) Windows

- [https://getcomposer.org/](https://getcomposer.org/)에서 설치 프로그램 다운로드 후 실행  
- 설치 중 PHP 경로를 물어보면 `php.exe` 경로 선택

---

### 2) macOS / Linux

```bash
curl -sS https://getcomposer.org/installer | php
sudo mv composer.phar /usr/local/bin/composer
```

---

### 3) 설치 확인

```bash
composer -V
```

---

## 2️⃣ Composer 기본 명령어

| 명령어                          | 설명                                     |
|----------------------------------|------------------------------------------|
| `composer init`                 | `composer.json` 초기 생성                 |
| `composer install`              | 의존성 설치 (`composer.lock` 기반)       |
| `composer update`               | 의존성 최신 버전으로 갱신                |
| `composer require 패키지명`     | 특정 패키지 설치 (`composer.json` 자동 추가) |
| `composer remove 패키지명`      | 패키지 제거                              |
| `composer dump-autoload`        | 오토로딩 정보 갱신                       |

---

## 3️⃣ 프로젝트 초기화 (`composer.json` 생성)

```bash
composer init
```

- 프로젝트 이름, 설명, 라이선스, 패키지 의존성 등을 입력하면 자동으로 `composer.json` 파일이 생성됩니다.

---

## 4️⃣ 패키지 설치 예시

```bash
composer require guzzlehttp/guzzle
```

✔ `vendor/` 디렉토리에 패키지 설치됨  
✔ `composer.json` 및 `composer.lock` 자동 업데이트  

---

## 5️⃣ 오토로딩 (PSR-4 규칙)

Composer는 설치된 패키지뿐만 아니라 **내 코드도 자동 로딩**할 수 있도록 지원합니다.

#### 1. `composer.json`에 등록

```json
{
  "autoload": {
    "psr-4": {
      "App\\": "src/"
    }
  }
}
```

✔ `"App\\": "src/"` → `App\Controller\HomeController` → `src/Controller/HomeController.php` 로 연결  

#### 2. 설정 후 반드시 아래 명령어 실행:

```bash
composer dump-autoload
```

✔ 패키지를 설치하거나, 직접 만든 클래스를 추가했을 경우 **오토로더가 변경 사항을 인식하도록 재생성**하는 것이 필요  

#### 3. 코드 사용 예시

```php
<?php
require 'vendor/autoload.php';

use App\Controller\HomeController;

$controller = new HomeController();
```

---

## 6️⃣ 전역 패키지 설치

CLI 도구나 전역적으로 사용하는 유틸 패키지는 **전역 설치**할 수 있습니다.

```bash
composer global require laravel/installer
```

✔ `~/.composer/vendor/bin` 또는 `%APPDATA%/Composer/vendor/bin` 을 PATH에 추가해야 실행 가능  

---

## 7️⃣ 패키지 검색 및 정보 확인

#### 1. 검색

```bash
composer search laravel
```

#### 2. 특정 패키지 정보 확인

```bash
composer show guzzlehttp/guzzle
```

✔ 버전, 의존성, 설명 등 확인 가능  

---

## 8️⃣ 의존성 관리 파일 설명

| 파일명            | 설명 |
|-------------------|------|
| `composer.json`   | 프로젝트의 의존성 정의 및 기본 설정 파일 |
| `composer.lock`   | 설치된 정확한 패키지 버전 기록 (재현성 보장) |
| `vendor/`         | 실제 패키지가 설치되는 디렉토리 |
| `autoload.php`    | Composer의 오토로더 파일 (자동 포함 필수) |

---

## 9️⃣ .gitignore에 반드시 추가할 항목

```txt
/vendor/
composer.lock
```

✔ 일반적으로 **`vendor/`는 git에 포함하지 않으며**, 동일 환경 구축 시 `composer install`로 재설치합니다.

---

## 🎯 정리

✔ Composer는 PHP의 공식 패키지 관리자이며, `composer.json`으로 의존성을 관리합니다.  
✔ `require`, `update`, `install`, `remove` 명령어로 라이브러리를 설치/제거할 수 있습니다.  
✔ PSR-4 오토로딩 설정을 통해 내 프로젝트 구조도 자동 로딩할 수 있습니다.  
✔ `vendor/autoload.php`를 통해 모든 라이브러리와 클래스가 자동 로딩됩니다.  
✔ 실무에서는 **재현성 확보를 위해 `composer.lock`을 버전 관리**에 포함시키는 것이 좋습니다.

