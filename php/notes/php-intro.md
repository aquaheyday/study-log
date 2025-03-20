# 🌐 PHP 개요 (Introduction to PHP)

PHP는 **웹 개발을 위한 서버 사이드 스크립트 언어**로,  
동적인 웹 페이지 및 애플리케이션을 구축하는 데 널리 사용됩니다.

---

## 1. PHP란?

- **PHP(Hypertext Preprocessor)** 는 웹 서버에서 실행되는 스크립트 언어입니다.
- HTML에 포함하여 **동적인 웹 페이지를 생성**할 수 있습니다.
- **MySQL, PostgreSQL, SQLite 등 다양한 데이터베이스와 연동**이 용이합니다.
- **Laravel, CodeIgniter, Symfony 등 강력한 프레임워크 지원**으로 개발 생산성을 향상할 수 있습니다.

---

## 2. PHP의 특징과 장점

| 특징 | 설명 |
|------|------|
| **오픈소스** | 누구나 무료로 사용할 수 있으며 커뮤니티가 활발함 |
| **서버 사이드 언어** | 클라이언트(브라우저)가 아닌 서버에서 코드 실행 |
| **HTML 및 데이터베이스 연동** | HTML 코드 내에서 PHP 사용 가능, MySQL 등과 쉽게 연결 가능 |
| **플랫폼 독립적** | Windows, macOS, Linux에서 실행 가능 |
| **광범위한 지원** | Apache, Nginx, IIS 등 다양한 웹 서버에서 동작 |
| **빠른 실행 속도** | 최신 PHP 버전(8.x)에서 성능 최적화 및 JIT 컴파일 지원 |

✔ PHP는 **초보자도 쉽게 배울 수 있으며**, **대규모 웹 애플리케이션 개발에도 적합**합니다.

---

## 3. PHP 기본 문법 예제

### 1) PHP 코드 작성
```php
<?php
echo "Hello, World!";
?>
```
✔ PHP 코드 블록은 `<?php ... ?>` 사이에 작성  
✔ `echo` 를 사용하여 문자열 출력  

---

### 2) 변수와 데이터 타입
```php
<?php
$name = "Alice";  // 문자열
$age = 25;        // 정수
$price = 9.99;    // 실수
$is_admin = true; // 불리언

echo "이름: $name, 나이: $age";
?>
```
✔ PHP는 동적 타입 언어 → 변수 선언 시 타입 지정 불필요  

---

### 3) 조건문 및 반복문
```php
<?php
$score = 85;
if ($score >= 90) {
    echo "A 학점";
} elseif ($score >= 80) {
    echo "B 학점";
} else {
    echo "C 학점";
}

for ($i = 1; $i <= 5; $i++) {
    echo "번호: $i<br>";
}
?>
```
✔ `if-elseif-else` 문으로 조건 처리  
✔ `for` 문을 이용한 반복 실행  

---

## 4. PHP의 주요 기능

### 1) HTML과 함께 사용 가능
```php
<!DOCTYPE html>
<html>
<body>
    <h1>오늘의 날짜:</h1>
    <p><?php echo date("Y-m-d"); ?></p>
</body>
</html>
```
✔ PHP 코드를 HTML 내에 삽입하여 동적 콘텐츠 생성  

---

### 2) 폼 처리 (GET & POST)
```php
<form method="post">
    이름: <input type="text" name="name">
    <input type="submit" value="제출">
</form>

<?php
if ($_SERVER["REQUEST_METHOD"] == "POST") {
    $name = $_POST["name"];
    echo "입력한 이름: " . htmlspecialchars($name);
}
?>
```
✔ `$_POST` → 폼 데이터 수집  
✔ `htmlspecialchars()` → 보안 강화를 위해 XSS 방지  

---

### 3) 데이터베이스 연동 (MySQL)
```php
<?php
$conn = new mysqli("localhost", "root", "password", "test_db");
if ($conn->connect_error) {
    die("연결 실패: " . $conn->connect_error);
}

$sql = "SELECT * FROM users";
$result = $conn->query($sql);

while ($row = $result->fetch_assoc()) {
    echo "이름: " . $row["name"] . "<br>";
}

$conn->close();
?>
```
✔ `mysqli` 를 사용하여 MySQL 데이터베이스 연결  

---

### 4) 파일 입출력
```php
<?php
$file = fopen("example.txt", "w");
fwrite($file, "Hello, PHP!");
fclose($file);
?>
```
✔ `fopen()`, `fwrite()`, `fclose()` 로 파일 조작 가능  

---

## 5. PHP의 활용 분야

| 활용 분야 | 설명 |
|----------|------|
| **동적 웹사이트** | 로그인 시스템, 사용자 인증, 데이터 관리 |
| **API 개발** | RESTful API를 만들어 모바일 앱 및 프론트엔드와 연동 |
| **CMS(Content Management System)** | WordPress, Drupal, Joomla 같은 CMS 개발 |
| **전자상거래** | WooCommerce, Magento 등 쇼핑몰 구축 |
| **웹 애플리케이션** | Laravel, CodeIgniter 등의 프레임워크로 대형 웹 앱 개발 |

✔ PHP는 **대부분의 웹 애플리케이션에서 사용 가능**하며, **유지보수와 확장성이 뛰어남**.

---

## 6. PHP 실행 방법

### 1) PHP CLI(Command Line Interface) 실행
```sh
php -v    # PHP 버전 확인
php script.php  # script.php 실행
```

### 2) PHP 내장 서버 실행 (테스트용)
```sh
php -S localhost:8000
```
✔ `localhost:8000` 에서 PHP 스크립트 실행 가능  

---

## 7. PHP와 함께 사용하는 기술

| 기술 | 설명 |
|------|------|
| **MySQL, PostgreSQL** | 데이터 저장 및 관리 |
| **Apache, Nginx** | 웹 서버 설정 |
| **Laravel, CodeIgniter** | PHP 프레임워크 |
| **Composer** | PHP 의존성 관리 도구 |
| **JavaScript (AJAX, Vue.js, React)** | 동적 UI 연동 |

✔ PHP는 다양한 기술과 함께 **웹 개발 생태계를 구성**함.

---

## 🎯 정리

✔ **PHP는 서버 사이드 스크립트 언어로 동적 웹 개발에 최적화**  
✔ **HTML과 함께 사용하여 웹 페이지를 동적으로 생성 가능**  
✔ **MySQL 등 데이터베이스와 쉽게 연동 가능**  
✔ **폼 처리, API 개발, 보안 기능, 프레임워크 활용 가능**  
✔ **빠르고 쉬운 개발이 가능하며, 대규모 웹사이트에도 적합**  
