<?php
// 세션 시작
session_start();

// 세션 변수 설정
if ($_SERVER["REQUEST_METHOD"] == "POST") {
    $_SESSION["username"] = htmlspecialchars($_POST["username"]);
    $_SESSION["user_age"] = htmlspecialchars($_POST["user_age"]);

    // 쿠키 설정 (1시간 동안 유지)
    setcookie("username", $_SESSION["username"], time() + 3600, "/");
    setcookie("user_age", $_SESSION["user_age"], time() + 3600, "/");
}

// 세션 & 쿠키 삭제
if (isset($_GET["logout"])) {
    session_unset(); // 모든 세션 변수 제거
    session_destroy(); // 세션 삭제

    // 쿠키 삭제
    setcookie("username", "", time() - 3600, "/");
    setcookie("user_age", "", time() - 3600, "/");

    header("Location: " . $_SERVER["PHP_SELF"]);
    exit;
}
?>

<!DOCTYPE html>
<html lang="ko">
<head>
    <meta charset="UTF-8">
    <title>PHP 세션과 쿠키 예제</title>
</head>
<body>
    <h2>세션과 쿠키 설정</h2>
    <form action="" method="post">
        이름: <input type="text" name="username" required><br>
        나이: <input type="number" name="user_age" required><br>
        <input type="submit" value="세션 & 쿠키 저장">
    </form>

    <h2>세션 데이터</h2>
    <?php if (isset($_SESSION["username"]) && isset($_SESSION["user_age"])) { ?>
        <p>이름: <?php echo $_SESSION["username"]; ?></p>
        <p>나이: <?php echo $_SESSION["user_age"]; ?></p>
        <a href="?logout=true">세션 & 쿠키 삭제</a>
    <?php } else { ?>
        <p>세션 데이터가 없습니다.</p>
    <?php } ?>

    <h2>쿠키 데이터</h2>
    <?php if (isset($_COOKIE["username"]) && isset($_COOKIE["user_age"])) { ?>
        <p>쿠키 저장된 이름: <?php echo $_COOKIE["username"]; ?></p>
        <p>쿠키 저장된 나이: <?php echo $_COOKIE["user_age"]; ?></p>
    <?php } else { ?>
        <p>쿠키 데이터가 없습니다.</p>
    <?php } ?>
</body>
</html>
