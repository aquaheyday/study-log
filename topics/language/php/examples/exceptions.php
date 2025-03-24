<?php
// 1. 기본 예외 처리 (try-catch)
function divide($a, $b) {
    if ($b == 0) {
        throw new Exception("0으로 나눌 수 없습니다!");
    }
    return $a / $b;
}

try {
    echo "10 / 2 = " . divide(10, 2) . "\n"; // 정상 실행
    echo "10 / 0 = " . divide(10, 0) . "\n"; // 예외 발생
} catch (Exception $e) {
    echo "예외 발생: " . $e->getMessage() . "\n";
}

// 2. 여러 개의 예외 처리 (다중 catch)
function checkAge($age) {
    if (!is_numeric($age)) {
        throw new InvalidArgumentException("나이는 숫자여야 합니다.");
    }
    if ($age < 18) {
        throw new Exception("18세 이상만 접근할 수 있습니다.");
    }
    return "접속 허용";
}

try {
    echo checkAge(20) . "\n"; // 정상 실행
    echo checkAge(15) . "\n"; // 예외 발생
} catch (InvalidArgumentException $e) {
    echo "입력 오류: " . $e->getMessage() . "\n";
} catch (Exception $e) {
    echo "접근 제한: " . $e->getMessage() . "\n";
}

// 3. finally 블록 사용
try {
    echo "try 블록 실행 중...\n";
    throw new Exception("에러 발생!");
} catch (Exception $e) {
    echo "예외 처리: " . $e->getMessage() . "\n";
} finally {
    echo "finally 블록 실행 (리소스 정리 가능)\n";
}

// 4. 사용자 정의 예외 (Custom Exception)
class MyException extends Exception {
    public function errorMessage() {
        return "사용자 정의 예외: " . $this->getMessage();
    }
}

try {
    throw new MyException("이것은 사용자 정의 예외입니다!");
} catch (MyException $e) {
    echo $e->errorMessage() . "\n";
}

// 5. 오류를 예외로 변환 (Error to Exception)
function errorToException($errno, $errstr) {
    throw new ErrorException($errstr, $errno);
}

// 사용자 정의 오류 핸들러 설정
set_error_handler("errorToException");

try {
    echo $undefinedVariable; // 존재하지 않는 변수 사용 (오류 발생)
} catch (ErrorException $e) {
    echo "오류를 예외로 변환: " . $e->getMessage() . "\n";
}

restore_error_handler(); // 기본 오류 핸들러 복원
?>
