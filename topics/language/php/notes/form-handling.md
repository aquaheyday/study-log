# 🌐 PHP 폼 처리 및 입력 검증

HTML 폼(Form)을 통해 사용자가 입력한 데이터를 PHP로 받아서 처리하는 기능은  
**회원가입, 로그인, 게시판 작성, 검색 등 거의 모든 웹 애플리케이션의 핵심 요소**입니다.  
또한 **입력값 검증과 보안 처리**는 매우 중요합니다.

---

## 1️⃣ 기본 폼 구조 (HTML)
- `method="post"`: 데이터를 POST 방식으로 전송  
- `action="process.php"`: 입력 데이터를 처리할 PHP 파일 지정  
- `name` 속성이 있어야 PHP에서 데이터를 받을 수 있습니다.

```html
<form method="post" action="process.php">
  <input type="text" name="username" placeholder="이름 입력">
  <input type="email" name="email" placeholder="이메일 입력">
  <button type="submit">제출</button>
</form>
```

---

## 2️⃣ 데이터 수신 (`$_POST`, `$_GET`)

```php
<?php
$username = $_POST["username"];
$email = $_POST["email"];

echo "이름: $username<br>";
echo "이메일: $email";
?>
```

- `$_POST["name"]`: POST 방식으로 전달된 값  
- `$_GET["name"]`: URL 쿼리로 전달된 값 (예: `?name=value`)  
- 값이 존재하지 않으면 에러가 발생할 수 있으므로 `isset()`으로 체크하는 것이 좋습니다.

---

## 3️⃣ 입력값 존재 여부 확인
- `empty()`: 값이 비어 있는지 확인 (`null`, 빈 문자열, 0 등 모두 포함)  
- `isset()`: 해당 변수가 존재하는지만 확인

```php
<?php
if (isset($_POST["username"])) {
    echo "username 변수가 존재하지 않습니다.";
}

if (empty($_POST["username"])) {
    echo "이름을 입력해주세요.";
}
?>
```

---

## 4️⃣ 입력값 유효성 검증

### 1) 예제: 이름 + 이메일 + 나이 검증

```php
<?php
$errors = [];

if (empty($_POST["username"])) {
    $errors[] = "이름을 입력해주세요.";
}

if (!filter_var($_POST["email"], FILTER_VALIDATE_EMAIL)) {
    $errors[] = "유효한 이메일 주소를 입력해주세요.";
}

if (!is_numeric($_POST["age"]) || $_POST["age"] < 0) {
    $errors[] = "나이는 양의 숫자로 입력해주세요.";
}

if (count($errors) > 0) {
    foreach ($errors as $msg) {
        echo "<p style='color:red;'>$msg</p>";
    }
} else {
    echo "입력 성공!";
}
?>
```

✔ `filter_var()` 함수는 다양한 필터로 입력값을 검증할 수 있습니다.  
✔ 숫자 검사에는 `is_numeric()` 또는 `ctype_digit()`을 사용할 수 있습니다.  

---

## 5️⃣ 보안 처리 (XSS / SQL Injection 방지)

### 1) 출력 시 XSS 방지: `htmlspecialchars()`

```php
<?php
echo htmlspecialchars($_POST["username"]);
?>
```

✔ 사용자의 입력값에 `<script>` 같은 HTML 태그가 있을 경우 **무력화시켜 출력**합니다.  

---

### 2) DB 입력 시 SQL Injection 방지: Prepared Statement 사용 예시

```php
$stmt = $pdo->prepare("INSERT INTO users (name, email) VALUES (?, ?)");
$stmt->execute([$_POST["username"], $_POST["email"]]);
```

✔ SQL 구문에 직접 값을 연결하지 않고 **바인딩하여 보안성 향상**  
✔ `mysqli`나 `PDO`에서 모두 지원됩니다.  

---

## 6️⃣ 폼 재전송 방지 (PRG 패턴)

POST 요청 후 새로고침 시 중복 전송을 막기 위해 **Post-Redirect-Get 패턴**을 사용합니다.

```php
<?php
// 처리 완료 후 리디렉션
header("Location: result.php");
// header() 함수로 다른 페이지로 이동
// result.php에서는 성공 메시지 또는 결과 화면 표시
exit;
?>
```

✔ 반드시 `exit` 또는 `die()` 호출로 후속 코드 실행 방지  

---

## 7️⃣ 파일 업로드 입력 예시

```html
<form method="post" action="upload.php" enctype="multipart/form-data">
  <input type="file" name="upload">
  <button type="submit">업로드</button>
</form>
```

```php
<?php
if ($_FILES["upload"]["error"] === 0) {
    move_uploaded_file($_FILES["upload"]["tmp_name"], "uploads/" . $_FILES["upload"]["name"]);
    echo "업로드 성공!";
}
?>
```

✔ 파일 업로드 시 `enctype="multipart/form-data"`를 반드시 추가해야 합니다.

---

## 🎯 정리

✔ HTML 폼은 `method`와 `action`을 설정하고, 각 필드에 `name` 속성이 있어야 PHP에서 접근할 수 있습니다.  
✔ `$_POST`, `$_GET`을 통해 전달된 값을 받고, `isset()`과 `empty()`로 검증합니다.  
✔ 유효성 검사에는 `filter_var()` 등의 검증 함수를 사용합니다.  
✔ 보안 강화를 위해 출력 시 `htmlspecialchars()`, DB 연동 시 **Prepared Statement** 사용이 필수입니다.  
✔ PRG 패턴으로 폼 중복 전송을 방지하고, 파일 업로드 시 `$_FILES` 배열을 사용합니다.

