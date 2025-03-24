# PHP 폼 처리와 슈퍼글로벌

## 1. PHP 폼 처리 개요
PHP에서 폼 데이터를 처리할 때 `$_GET`, `$_POST`, `$_REQUEST`와 같은 슈퍼글로벌 변수를 사용합니다.

---

## 2. HTML 폼 예제
```html
<form action="process.php" method="post">
    <label for="name">이름:</label>
    <input type="text" id="name" name="name">
    <br>
    <label for="email">이메일:</label>
    <input type="email" id="email" name="email">
    <br>
    <input type="submit" value="제출">
</form>
```

---

## 3. `$_POST`를 이용한 폼 데이터 처리
```php
// process.php
if ($_SERVER["REQUEST_METHOD"] == "POST") {
    $name = htmlspecialchars($_POST["name"]);
    $email = htmlspecialchars($_POST["email"]);
    
    echo "이름: " . $name . "<br>";
    echo "이메일: " . $email;
}
```

---

## 4. `$_GET`을 이용한 데이터 전송
```html
<a href="process.php?name=Alice&email=alice@example.com">Alice 정보 전송</a>
```

```php
// process.php
if ($_SERVER["REQUEST_METHOD"] == "GET") {
    $name = htmlspecialchars($_GET["name"] ?? "");
    $email = htmlspecialchars($_GET["email"] ?? "");
    
    echo "이름: " . $name . "<br>";
    echo "이메일: " . $email;
}
```

---

## 5. `$_REQUEST`를 이용한 데이터 처리
`$_REQUEST`는 `$_GET`과 `$_POST`를 모두 포함하는 배열입니다.
```php
if (!empty($_REQUEST["name"])) {
    echo "이름: " . htmlspecialchars($_REQUEST["name"]);
}
```

---

## 6. `$_SERVER`를 이용한 요청 확인
```php
if ($_SERVER["REQUEST_METHOD"] == "POST") {
    echo "폼이 POST 방식으로 제출되었습니다.";
}
```

---

## 7. 파일 업로드 처리 (`$_FILES`)
HTML 폼에서 파일을 업로드하려면 `enctype="multipart/form-data"`를 추가해야 합니다.
```html
<form action="upload.php" method="post" enctype="multipart/form-data">
    <input type="file" name="fileToUpload">
    <input type="submit" value="업로드">
</form>
```

```php
// upload.php
if ($_FILES["fileToUpload"]["error"] == UPLOAD_ERR_OK) {
    $targetDir = "uploads/";
    $targetFile = $targetDir . basename($_FILES["fileToUpload"]["name"]);
    
    if (move_uploaded_file($_FILES["fileToUpload"]["tmp_name"], $targetFile)) {
        echo "파일 업로드 성공: " . $targetFile;
    } else {
        echo "파일 업로드 실패.";
    }
} else {
    echo "파일 업로드 중 오류 발생.";
}
```
