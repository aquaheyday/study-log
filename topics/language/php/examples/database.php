<?php
// PDO 미사용
// 1. MySQL 데이터베이스 연결
$servername = "localhost"; // 데이터베이스 서버 주소
$username = "root";        // MySQL 사용자명
$password = "";            // MySQL 비밀번호
$dbname = "testdb";        // 사용할 데이터베이스 이름

// MySQL 연결 생성
$conn = new mysqli($servername, $username, $password, $dbname);

// 연결 확인
if ($conn->connect_error) {
    die("연결 실패: " . $conn->connect_error);
}
echo "데이터베이스 연결 성공!\n";

// 2. 테이블 생성 (users)
$sql = "CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)";

if ($conn->query($sql) === TRUE) {
    echo "테이블 생성 완료!\n";
} else {
    echo "테이블 생성 오류: " . $conn->error . "\n";
}

// 3. 데이터 삽입 (INSERT)
$name = "홍길동";
$email = "hong@example.com";

$sql = "INSERT INTO users (name, email) VALUES ('$name', '$email')";

if ($conn->query($sql) === TRUE) {
    echo "데이터 삽입 성공!\n";
} else {
    echo "데이터 삽입 오류: " . $conn->error . "\n";
}

// 4. 데이터 조회 (SELECT)
$sql = "SELECT id, name, email FROM users";
$result = $conn->query($sql);

if ($result->num_rows > 0) {
    echo "등록된 사용자 목록:\n";
    while ($row = $result->fetch_assoc()) {
        echo "ID: " . $row["id"] . " - 이름: " . $row["name"] . " - 이메일: " . $row["email"] . "\n";
    }
} else {
    echo "사용자가 없습니다.\n";
}

// 5. 데이터 업데이트 (UPDATE)
$newEmail = "hong123@example.com";
$sql = "UPDATE users SET email='$newEmail' WHERE name='홍길동'";

if ($conn->query($sql) === TRUE) {
    echo "데이터 업데이트 성공!\n";
} else {
    echo "데이터 업데이트 오류: " . $conn->error . "\n";
}

// 6. 데이터 삭제 (DELETE)
$sql = "DELETE FROM users WHERE name='홍길동'";

if ($conn->query($sql) === TRUE) {
    echo "데이터 삭제 성공!\n";
} else {
    echo "데이터 삭제 오류: " . $conn->error . "\n";
}

// 7. 데이터베이스 연결 닫기
$conn->close();

// PDO 사용
// 1. PDO를 사용한 MySQL 데이터베이스 연결
$dsn = "mysql:host=localhost;dbname=testdb;charset=utf8mb4";
$username = "root";   // MySQL 사용자명
$password = "";       // MySQL 비밀번호

try {
    $pdo = new PDO($dsn, $username, $password, [
        PDO::ATTR_ERRMODE => PDO::ERRMODE_EXCEPTION, // 예외 처리 모드
        PDO::ATTR_DEFAULT_FETCH_MODE => PDO::FETCH_ASSOC, // 연관 배열로 결과 반환
    ]);
    echo "데이터베이스 연결 성공!\n";
} catch (PDOException $e) {
    die("연결 실패: " . $e->getMessage());
}

// 2. 테이블 생성 (users)
$sql = "CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)";
$pdo->exec($sql);
echo "테이블 생성 완료!\n";

// 3. 데이터 삽입 (Prepared Statement 사용)
$name = "홍길동";
$email = "hong@example.com";

$sql = "INSERT INTO users (name, email) VALUES (:name, :email)";
$stmt = $pdo->prepare($sql);
$stmt->bindParam(':name', $name);
$stmt->bindParam(':email', $email);
$stmt->execute();
echo "데이터 삽입 성공!\n";

// 4. 데이터 조회 (SELECT)
$sql = "SELECT * FROM users";
$stmt = $pdo->query($sql);
$users = $stmt->fetchAll();

if ($users) {
    echo "사용자 목록:\n";
    foreach ($users as $user) {
        echo "ID: {$user['id']} - 이름: {$user['name']} - 이메일: {$user['email']}\n";
    }
} else {
    echo "등록된 사용자가 없습니다.\n";
}

// 5. 데이터 업데이트 (UPDATE)
$newEmail = "hong123@example.com";
$sql = "UPDATE users SET email = :email WHERE name = :name";
$stmt = $pdo->prepare($sql);
$stmt->bindParam(':email', $newEmail);
$stmt->bindParam(':name', $name);
$stmt->execute();
echo "데이터 업데이트 성공!\n";

// 6. 데이터 삭제 (DELETE)
$sql = "DELETE FROM users WHERE name = :name";
$stmt = $pdo->prepare($sql);
$stmt->bindParam(':name', $name);
$stmt->execute();
echo "데이터 삭제 성공!\n";

// 7. 데이터베이스 연결 닫기 (PDO는 자동으로 연결 해제됨)
$pdo = null;

?>
