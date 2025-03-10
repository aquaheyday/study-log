# PHP 에러 처리와 예외

## 1. 에러 종류

PHP에서 발생하는 주요 에러 유형은 다음과 같습니다:
- **E_NOTICE**: 실행에 영향을 미치지 않는 경고 (예: 정의되지 않은 변수 사용)
- **E_WARNING**: 실행은 계속되지만 경고 메시지 출력 (예: `include` 파일 없음)
- **E_ERROR**: 치명적인 오류로 실행 중단 (예: 잘못된 함수 호출)

---

## 2. 기본 에러 출력 설정

```php
ini_set('display_errors', 1);
ini_set('display_startup_errors', 1);
error_reporting(E_ALL); // 모든 에러 출력
```

---

## 3. `error_log`를 사용한 에러 기록

```php
error_log("이것은 사용자 정의 에러 로그입니다.", 3, "errors.log");
```

---

## 4. 사용자 정의 에러 핸들러
```php
function customErrorHandler($errno, $errstr, $errfile, $errline) {
    echo "에러 [$errno]: $errstr - 파일: $errfile (줄: $errline)";
    error_log("에러 [$errno]: $errstr - 파일: $errfile (줄: $errline)", 3, "errors.log");
}

set_error_handler("customErrorHandler");

// 테스트 에러
echo $undefinedVar;
```

---

## 5. 예외 처리 (`try-catch`)

```php
try {
    throw new Exception("예외 발생!");
} catch (Exception $e) {
    echo "예외 메시지: " . $e->getMessage();
}
```

---

## 6. 예외의 상세 정보 가져오기

```php
try {
    throw new Exception("파일을 찾을 수 없습니다.", 404);
} catch (Exception $e) {
    echo "예외 코드: " . $e->getCode() . "<br>";
    echo "예외 메시지: " . $e->getMessage() . "<br>";
    echo "예외 발생 파일: " . $e->getFile() . "<br>";
    echo "예외 발생 줄: " . $e->getLine();
}
```

---

## 7. 사용자 정의 예외 클래스

```php
class CustomException extends Exception {
    public function errorMessage() {
        return "오류 발생: " . $this->getMessage();
    }
}

try {
    throw new CustomException("사용자 정의 예외입니다.");
} catch (CustomException $e) {
    echo $e->errorMessage();
}
```

---

## 8. `finally` 블록 사용

```php
try {
    echo "작업 수행 중...<br>";
    throw new Exception("문제가 발생했습니다.");
} catch (Exception $e) {
    echo "예외 발생: " . $e->getMessage() . "<br>";
} finally {
    echo "이 부분은 항상 실행됩니다.";
}
```

---

## 9. PHP 7+ 예외 처리 (`ErrorException`)

PHP 7부터는 `Error` 클래스를 이용해 보다 강력한 에러 처리가 가능합니다.

```php
try {
    intdiv(10, 0); // 0으로 나누기 에러 발생
} catch (Error $e) {
    echo "오류 발생: " . $e->getMessage();
}
```
