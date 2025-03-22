# 🔐 PHP 보안 기본 가이드

웹 애플리케이션 보안은 **데이터 보호와 시스템 안정성**을 위해 필수입니다.  

---

## 1️⃣ XSS (Cross-Site Scripting) 방지

**XSS**는 악의적인 사용자가 `<script>` 태그 등으로 **사용자 브라우저에서 스크립트를 실행시키는 공격**입니다.

### 방어 방법: 출력 시 `htmlspecialchars()` 사용

```php
<?php
$name = $_GET["name"];
echo htmlspecialchars($name, ENT_QUOTES, "UTF-8");
?>
```

| 함수              | 설명                           |
|-------------------|--------------------------------|
| `htmlspecialchars`| HTML 태그 문자들을 엔티티로 변환 |
| `ENT_QUOTES`      | 작은따옴표/큰따옴표 모두 변환     |
| `UTF-8`           | 문자 인코딩 지정                |

---

## 2️⃣ SQL Injection 방지

SQL 인젝션은 사용자 입력을 조작하여 **데이터베이스 쿼리를 변경하거나 탈취하는 공격**입니다.

### 방어 방법: Prepared Statement 사용

```php
<?php
$stmt = $pdo->prepare("SELECT * FROM users WHERE email = ?");
$stmt->execute([$_POST["email"]]);
?>
```

✔ `prepare()` + `execute()` 조합은 **입력값을 자동으로 이스케이프** 처리해줍니다.  

---

## 3️⃣ CSRF (Cross-Site Request Forgery) 방지

**CSRF**는 사용자의 인증된 세션을 이용하여 **원하지 않는 요청을 강제로 실행**시키는 공격입니다.

### 방어 방법: CSRF 토큰 사용

#### 1. HTML 폼에 토큰 삽입

```php
<?php
session_start();
$token = bin2hex(random_bytes(32));
$_SESSION["csrf_token"] = $token;
?>
<form method="POST">
  <input type="hidden" name="csrf_token" value="<?= $token ?>">
</form>
```

#### 2. 서버에서 토큰 검증

```php
if ($_POST["csrf_token"] !== $_SESSION["csrf_token"]) {
    die("CSRF 공격 차단됨");
}
```

---

## 4️⃣ 파일 업로드 보안

사용자 업로드 파일을 악용하면 서버가 **공격당할 수 있으므로** 반드시 검증이 필요합니다.

### 방어 방법

```php
$allowedTypes = ['image/jpeg', 'image/png'];
if (!in_array($_FILES["upload"]["type"], $allowedTypes)) {
    die("허용되지 않는 파일 형식입니다.");
}

$filename = basename($_FILES["upload"]["name"]);
move_uploaded_file($_FILES["upload"]["tmp_name"], "uploads/" . $filename);
```

✔ MIME 타입, 확장자, 용량 등을 모두 검증  
✔ 업로드 디렉토리에 `.htaccess`로 PHP 실행 방지 권장  

```
# .htaccess 예시
php_flag engine off
```

---

## 5️⃣ 세션 보안

세션은 사용자 인증 정보가 담기는 중요한 공간이므로 **세션 탈취를 막기 위한 보안 설정**이 필요합니다.

### 안전한 세션 처리

```php
session_start();
session_regenerate_id(true); // 로그인 시 세션 재발급
```

✔ HTTPS 연결 필수 (`session.cookie_secure = On`)  
✔ 브라우저 종료 시 삭제 (`session.cookie_lifetime = 0`)  
✔ `HttpOnly`, `SameSite=Strict` 쿠키 옵션 설정 권장  

---

## 6️⃣ 입력값 검증 및 필터링

```php
$email = filter_input(INPUT_POST, "email", FILTER_VALIDATE_EMAIL);
$age = filter_input(INPUT_POST, "age", FILTER_VALIDATE_INT);
```

✔  `filter_input()` 함수는 안전하고 구조적인 검증 도구입니다.  
✔  정규 표현식도 함께 사용 가능  

---

## 7️⃣ 에러 메시지 노출 방지

```php
// 운영 서버에서는
ini_set("display_errors", 0);
error_reporting(0);
```

✔ 개발 환경에서는 오류를 표시하되, 운영 환경에서는 **사용자에게 내부 정보 노출을 차단**해야 합니다.  

---

## 8️⃣ 패스워드 안전하게 저장하기

### 비밀번호 해싱 (`password_hash`, `password_verify`)

```php
$hash = password_hash($password, PASSWORD_DEFAULT); // 저장 시
```

```php
if (password_verify($input, $hash)) {
    echo "로그인 성공";
}
```

✔ 단방향 해시 사용 (복호화 불가)  
✔ `bcrypt` 또는 `argon2` 기반으로 자동 적용됨  
⚠️ 절대로 `md5()`나 `sha1()` 등은 사용하지 마세요  

---

## 🎯 정리

✔ XSS는 `htmlspecialchars()`로 출력값 보호  
✔ SQL Injection은 Prepared Statement로 차단  
✔ CSRF는 토큰을 생성하여 검증  
✔ 파일 업로드 시 MIME, 확장자, 실행 권한 철저히 제한  
✔ 세션은 HTTPS + ID 재생성으로 보호  
✔ 비밀번호는 반드시 `password_hash()`로 해시 처리  
✔ 운영환경에서는 에러 출력 제거 및 디버깅 정보 노출 방지

