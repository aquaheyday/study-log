# 🗄️ PHP 데이터베이스 연동

PHP는 다양한 데이터베이스(MySQL, PostgreSQL, SQLite 등)와 연동할 수 있으며,  
그 중에서도 가장 널리 사용되는 **MySQL**을 중심으로 데이터베이스 연동 방법을 설명드립니다.

---

## 1️⃣ PHP에서 MySQL 연결 방법

### 1) `mysqli` 확장 사용 (객체 지향 방식)

```php
<?php
$mysqli = new mysqli("localhost", "username", "password", "database"); // new mysqli(호스트, 사용자명, 비밀번호, 데이터베이스명)

if ($mysqli->connect_error) {
    die("연결 실패: " . $mysqli->connect_error);
}
echo "DB 연결 성공!";
?>
```

✔ 연결 실패 시 `$mysqli->connect_error`를 통해 오류 메시지 확인 가능

---

## 2️⃣ SQL 쿼리 실행

### 1) SELECT (조회)
`query()` 메서드로 쿼리 실행  

```php
<?php
$sql = "SELECT id, name FROM users";
$result = $mysqli->query($sql);

while ($row = $result->fetch_assoc()) {
    echo $row["id"] . ": " . $row["name"] . "<br>";
}
?>
```
 
✔ `fetch_assoc()`으로 결과를 **연관 배열**로 가져옴

---

### 2) INSERT / UPDATE / DELETE
성공 여부는 `TRUE` / `FALSE`로 반환됩니다.  
오류 발생 시 `$mysqli->error` 확인가능합니다.  

```php
<?php
$sql = "INSERT INTO users (name, email) VALUES ('Alice', 'alice@example.com')";
if ($mysqli->query($sql) === TRUE) {
    echo "데이터 삽입 성공!";
} else {
    echo "오류: " . $mysqli->error;
}
?>
```

---

## 3️⃣ 사용자 입력값 바인딩 (보안 처리)

SQL Injection을 방지하기 위해 **Prepared Statement**를 사용합니다.
- `prepare()` → 쿼리 구조 생성  
- `bind_param()` → 실제 값 바인딩  
- `execute()` → 실행

```php
<?php
$stmt = $mysqli->prepare("INSERT INTO users (name, email) VALUES (?, ?)");
$stmt->bind_param("ss", $name, $email); // s: string, i: integer 등

$name = "Bob";
$email = "bob@example.com";
$stmt->execute();
$stmt->close();
?>
```

✔ **보안을 위해 사용자 입력은 반드시 prepared statement로 처리하시는 걸 권장**

---

## 4️⃣ 연결 종료

```php
<?php
$mysqli->close();
?>
```

✔ 모든 작업이 끝난 뒤에는 연결을 종료하는 것이 좋습니다.

---

## 5️⃣ 에러 처리 및 트랜잭션

#### 트랜잭션 예제

```php
<?php
$mysqli->begin_transaction();

try {
    $mysqli->query("UPDATE accounts SET balance = balance - 100 WHERE id = 1");
    $mysqli->query("UPDATE accounts SET balance = balance + 100 WHERE id = 2");
    $mysqli->commit();
    echo "이체 성공!";
} catch (Exception $e) {
    $mysqli->rollback();
    echo "이체 실패: " . $e->getMessage();
}
?>
```

✔ `begin_transaction()` → `commit()` 또는 `rollback()` 으로 묶어서 처리 가능  
✔ 예외 발생 시 `rollback()`으로 안전하게 취소할 수 있습니다.

---

## 6️⃣ PDO 사용 예시 (선택)

PHP에서는 `PDO`(PHP Data Objects)를 통해 **다양한 DB를 통합된 방식**으로 사용할 수 있습니다.

```php
<?php
try {
    $pdo = new PDO("mysql:host=localhost;dbname=testdb", "root", "password");
    $pdo->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);

    $stmt = $pdo->prepare("SELECT * FROM users WHERE id = ?");
    $stmt->execute([1]);
    $user = $stmt->fetch();

    echo $user["name"];
} catch (PDOException $e) {
    echo "DB 오류: " . $e->getMessage();
}
?>
```

- `PDO`는 다양한 DB에 대응 가능 (MySQL, PostgreSQL, SQLite 등)  
- 예외 기반 에러 처리와 바인딩이 편리함

---

## 🎯 정리

✔ `mysqli` 또는 `PDO`를 사용하여 MySQL에 연결할 수 있습니다.  
✔ 쿼리는 `query()` 또는 `prepare()`, `bind_param()`을 통해 실행합니다.  
✔ 보안상 사용자 입력은 반드시 prepared statement로 처리하시는 것이 좋습니다.  
✔ 트랜잭션 처리(`begin_transaction`, `commit`, `rollback`)를 통해 데이터 무결성을 유지할 수 있습니다.  
✔ `PDO`는 다양한 DB에 대한 호환성과 예외 기반 에러 처리에 강력한 기능을 제공합니다.

