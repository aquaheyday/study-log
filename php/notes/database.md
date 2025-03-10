# PHP 데이터베이스 연동 (PDO)

## 1. PDO 개요

PHP Data Objects(PDO)는 데이터베이스와 상호작용하는 PHP의 객체 지향 방식입니다.

---

## 2. 데이터베이스 연결 (MySQL 예제)

```php
$dsn = "mysql:host=localhost;dbname=testdb;charset=utf8mb4";
$username = "root";
$password = "";

try {
    $pdo = new PDO($dsn, $username, $password, [
        PDO::ATTR_ERRMODE => PDO::ERRMODE_EXCEPTION,
        PDO::ATTR_DEFAULT_FETCH_MODE => PDO::FETCH_ASSOC
    ]);
    echo "데이터베이스 연결 성공";
} catch (PDOException $e) {
    echo "연결 실패: " . $e->getMessage();
}
```

---

## 3. 데이터 삽입 (INSERT)

```php
$sql = "INSERT INTO users (name, email) VALUES (:name, :email)";
$stmt = $pdo->prepare($sql);
$stmt->execute([
    'name' => 'Alice',
    'email' => 'alice@example.com'
]);

echo "데이터 삽입 성공";
```

---

## 4. 데이터 조회 (SELECT)

```php
$sql = "SELECT * FROM users WHERE email = :email";
$stmt = $pdo->prepare($sql);
$stmt->execute(['email' => 'alice@example.com']);
$user = $stmt->fetch();

if ($user) {
    echo "사용자 이름: " . $user['name'];
} else {
    echo "사용자를 찾을 수 없습니다.";
}
```

---

## 5. 데이터 수정 (UPDATE)

```php
$sql = "UPDATE users SET name = :name WHERE email = :email";
$stmt = $pdo->prepare($sql);
$stmt->execute([
    'name' => 'Alice Updated',
    'email' => 'alice@example.com'
]);

echo "데이터 수정 성공";
```

---

## 6. 데이터 삭제 (DELETE)

```php
$sql = "DELETE FROM users WHERE email = :email";
$stmt = $pdo->prepare($sql);
$stmt->execute(['email' => 'alice@example.com']);

echo "데이터 삭제 성공";
```

---

## 7. 모든 데이터 조회 (반복문 사용)

```php
$sql = "SELECT * FROM users";
$stmt = $pdo->query($sql);

foreach ($stmt as $row) {
    echo "이름: " . $row['name'] . ", 이메일: " . $row['email'] . "<br>";
}
```

---

## 8. 트랜잭션 사용하기

```php
try {
    $pdo->beginTransaction();
    
    $sql1 = "UPDATE users SET name = 'Bob' WHERE email = 'bob@example.com'";
    $pdo->exec($sql1);
    
    $sql2 = "DELETE FROM users WHERE email = 'charlie@example.com'";
    $pdo->exec($sql2);
    
    $pdo->commit();
    echo "트랜잭션 완료";
} catch (Exception $e) {
    $pdo->rollBack();
    echo "트랜잭션 실패: " . $e->getMessage();
}
```
