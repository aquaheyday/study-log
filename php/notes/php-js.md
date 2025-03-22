# 🧩 PHP와 JavaScript 연동

PHP는 서버에서 실행되고, JavaScript는 브라우저에서 실행되는 언어입니다.  
이 둘을 함께 사용하면 **동적인 웹 애플리케이션**을 만들 수 있습니다.

PHP와 JS는 직접 통신하지 않지만, HTML을 매개로 하거나 `Ajax` / `Fetch API` / `Form` 전송을 통해 간접적으로 데이터를 주고받을 수 있습니다.

---

## 1️⃣ PHP 변수 → JavaScript로 전달

PHP 변수를 JavaScript로 넘길 때는 **HTML에 출력**하여 JS가 읽을 수 있게 합니다.

### 1) 직접 JS 코드에 출력
PHP의 `echo`를 통해 값을 JS 코드 안에 삽입합니다.  


```php
<?php
$phpName = "홍길동";
?>
<script>
  const name = "<?php echo $phpName; ?>";
  console.log("사용자 이름:", name);
</script>
```

⚠️ 문자열은 따옴표로 감싸는 것에 주의해주세요.  

---

### 2) JSON으로 전달

배열이나 객체는 `json_encode()`로 JS에서 그대로 사용할 수 있습니다.

```php
<?php
$user = ["name" => "Alice", "age" => 30];
?>
<script>
  const user = <?php echo json_encode($user); ?>;
  console.log(user.name); // "Alice"
</script>
```

✔ `json_encode()` → PHP 배열을 JS 객체로 변환해주는 함수입니다.

---

## 2️⃣ JavaScript → PHP로 데이터 전송

JS에서 데이터를 PHP로 보내려면 `form` 전송 / `fetch()` / `XMLHttpRequest()` 를 사용하여 **서버에 요청**을 보내야 합니다.

### 1) 폼(`form`) 전송

```html
<form method="POST" action="save.php">
  <input type="text" name="username" value="홍길동">
  <button type="submit">저장</button>
</form>
```

✔ HTML 폼을 전송하면 `save.php`에서 `$_POST["username"]`으로 값 수신 가능  

---

### 2) JavaScript Fetch API (`Ajax` 요청)

```html
<script>
  fetch("save.php", {
    method: "POST",
    headers: {
      "Content-Type": "application/x-www-form-urlencoded"
    },
    body: "username=" + encodeURIComponent("홍길동")
  })
  .then(res => res.text())
  .then(data => {
    console.log("응답:", data);
  });
</script>
```

#### PHP에서 데이터 받기

```php
<?php
$username = $_POST["username"];
echo "서버에서 받은 이름: $username";
?>
```

✔ `Content-Type`이 `application/x-www-form-urlencoded`인 경우, PHP는 `$_POST`로 데이터를 받습니다.

---

## 3️⃣ JS로 PHP 응답 데이터 받아 처리 (API 스타일)

```html
<script>
  fetch("get-user.php")
    .then(res => res.json())
    .then(data => {
      console.log("사용자 정보:", data);
    });
</script>
```

#### `get-user.php` 예시:

```php
<?php
$user = ["id" => 1, "name" => "홍길동"];
header("Content-Type: application/json");
echo json_encode($user);
?>
```

✔ PHP에서 `Content-Type: application/json` 헤더를 지정하면, JS `fetch().json()`으로 바로 파싱할 수 있습니다.

---

## 4️⃣ PHP에서 JS 함수 호출은 불가능하지만…

PHP → JS 직접 호출은 불가능하지만 **조건에 따라 JS를 삽입**할 수는 있습니다.

```php
<?php
if ($isLoginFailed) {
    echo "<script>alert('로그인 실패');</script>";
}
?>
```

✔ PHP가 HTML 문서로 JS 코드를 출력하여 **조건에 따라 JS 동작을 유도할 수 있습니다.**

---

## 5️⃣ 보안 주의사항

| 항목            | 설명                                                   |
|------------------|--------------------------------------------------------|
| XSS 방지         | 사용자 입력을 `htmlspecialchars()`로 출력해야 안전합니다 |
| CSRF 보호        | JS에서 POST 요청 시 토큰 확인 필요 (CSRF Token 사용)     |
| 데이터 검증       | JS에서 보내는 값은 PHP에서 **반드시 서버 측 검증** 필요   |
| JSON 헤더 설정    | PHP에서 JSON 응답 시 `Content-Type: application/json` 필수 |

---

## 🎯 정리

✔ PHP는 서버 측, JavaScript는 클라이언트 측 언어이므로 **직접 통신은 불가능**합니다.  
✔ PHP 데이터를 JS로 넘기려면 `echo` 또는 `json_encode()`를 활용합니다.  
✔ JS에서 PHP로 전송하려면 `form`, `fetch()`, `Ajax` 요청 등을 사용합니다.  
✔ PHP는 조건에 따라 JS 코드를 출력할 수 있지만, JS에서 PHP 함수를 직접 호출할 수는 없습니다.  
✔ 실무에서는 **API 형식(JSON 통신)** 으로 데이터를 주고받는 구조가 많이 사용됩니다.

