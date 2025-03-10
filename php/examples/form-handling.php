<?php
session_start(); // 세션 시작

// 변수 초기화
$name = $email = $uploadMessage = "";
$nameErr = $emailErr = "";
$fileUploaded = false;

// 폼 데이터 처리 (POST 방식)
if ($_SERVER["REQUEST_METHOD"] == "POST") {
    // 이름 입력 확인
    if (empty($_POST["name"])) {
        $nameErr = "이름을 입력하세요.";
    } else {
        $name = htmlspecialchars($_POST["name"]);
        $_SESSION["name"] = $name; // 세션 저장
    }

    // 이메일 유효성 검사
    if (empty($_POST["email"])) {
        $emailErr = "이메일을 입력하세요.";
    } elseif (!filter_var($_POST["email"], FILTER_VALIDATE_EMAIL)) {
        $emailErr = "유효한 이메일을 입력하세요.";
    } else {
        $email = htmlspecialchars($_POST["email"]);
    }

    // 파일 업로드 처리
    if (!empty($_FILES["fileToUpload"]["name"])) {
        $target_dir = "uploads/";
        $target_file = $target_dir . basename($_FILES["fileToUpload"]["name"]);

        if (move_uploaded_file($_FILES["fileToUpload"]["tmp_name"], $target_file)) {
            $uploadMessage = "파일이 업로드되었습니다: " . $target_file;
            $fileUploaded = true;
        } else {
            $uploadMessage = "파일 업로드 실패!";
        }
    }
}

// GET 방식 데이터 받기
$getName = isset($_GET["name"]) ? htmlspecialchars($_GET["name"]) : "없음";
$getEmail = isset($_GET["email"]) ? htmlspecialchars($_GET["email"]) : "없음";
?>

<!DOCTYPE html>
<html lang="ko">
<head>
    <meta charset="UTF-8">
    <title>PHP 폼 데이터 처리</title>
</head>
<body>
    <h2>폼 입력</h2>
    <form action="" method="post" enctype="multipart/form-data">
        이름: <input type="text" name="name" value="<?php echo $name; ?>">
        <span style="color:red;"><?php echo $nameErr; ?></span><br>

        이메일: <input type="email" name="email" value="<?php echo $email; ?>">
        <span style="color:red;"><?php echo $emailErr; ?></span><br>

        파일 업로드: <input type="file" name="fileToUpload"><br>

        <input type="submit" value="제출">
    </form>

    <?php if ($name && $email) { ?>
        <h3>입력된 정보:</h3>
        <p>이름: <?php echo $name; ?></p>
        <p>이메일: <?php echo $email; ?></p>
    <?php } ?>

    <?php if ($fileUploaded) { ?>
        <p style="color:green;"><?php echo $uploadMessage; ?></p>
    <?php } ?>

    <h2>세션 데이터</h2>
    <p>저장된 세션 이름: <?php echo isset($_SESSION["name"]) ? $_SESSION["name"] : "세션 값 없음"; ?></p>

    <h2>GET 방식 데이터</h2>
    <a href="?name=홍길동&email=test@example.com">GET 방식으로 데이터 전달</a>
    <p>이름: <?php echo $getName; ?></p>
    <p>이메일: <?php echo $getEmail; ?></p>
</body>
</html>
