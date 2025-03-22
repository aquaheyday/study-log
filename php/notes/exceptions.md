# 🛠️ PHP 예외 처리 (Exception Handling)

PHP에서 예외(Exception)는 **코드 실행 중 발생할 수 있는 오류 상황을 제어하기 위한 구조**입니다.  
`try-catch` 블록을 사용하여 에러가 발생하더라도 프로그램이 **비정상 종료되지 않도록 처리**할 수 있습니다.

---

## 1️⃣ `try-catch` 기본 문법

예외가 발생할 가능성이 있는 코드를 `try` 블록에 작성하고,  
예외 발생 시 처리할 로직을 `catch` 블록에 작성합니다.

```php
<?php
try {
    throw new Exception("오류가 발생했습니다!");
} catch (Exception $e) {
    echo "예외 메시지: " . $e->getMessage();
}
?>
```

✔ `throw` 키워드로 예외를 발생시킵니다.  
✔ `Exception` 클래스는 PHP의 기본 예외 클래스입니다.  
✔ `$e->getMessage()`는 예외 메시지를 반환합니다.

---

## 2️⃣ finally 블록

`finally` 블록은 예외 발생 여부와 상관없이 **항상 실행되는 영역**입니다.  
파일 닫기, 연결 종료 등 **정리(clean-up)** 작업에 유용합니다.

```php
<?php
try {
    echo "작업 실행 중...";
    throw new Exception("문제 발생");
} catch (Exception $e) {
    echo "에러: " . $e->getMessage();
} finally {
    echo "리소스 정리 완료";
}
?>
```

✔ `try` → `catch` → `finally` 순서대로 실행됩니다.

---

## 3️⃣ `Exception` 클래스 주요 메서드

| 메서드             | 설명                           |
|--------------------|--------------------------------|
| `getMessage()`     | 예외 메시지 반환               |
| `getCode()`        | 에러 코드 반환                 |
| `getFile()`        | 예외가 발생한 파일 경로        |
| `getLine()`        | 예외가 발생한 줄 번호          |
| `getTrace()`       | 스택 트레이스 배열 반환        |
| `getTraceAsString()` | 스택 트레이스를 문자열로 반환 |

```php
<?php
try {
    throw new Exception("예외 테스트", 404);
} catch (Exception $e) {
    echo "메시지: " . $e->getMessage();
    echo "코드: " . $e->getCode();
    echo "파일: " . $e->getFile();
    echo "라인: " . $e->getLine();
}
?>
```

---

## 4️⃣ 사용자 정의 예외 클래스

필요에 따라 기본 `Exception` 클래스를 상속하여 **커스텀 예외를 정의**할 수 있습니다.

```php
<?php
class MyException extends Exception {}

function divide($a, $b) {
    if ($b === 0) {
        throw new MyException("0으로 나눌 수 없습니다.");
    }
    return $a / $b;
}

try {
    echo divide(10, 0);
} catch (MyException $e) {
    echo "사용자 정의 예외: " . $e->getMessage();
}
?>
```

✔ `MyException` 클래스는 `Exception`을 상속받아 생성되었습니다.  
✔ 커스텀 예외를 통해 보다 **명확하고 구체적인 에러 처리**가 가능합니다.

---

## 5️⃣ 다중 catch 블록

서로 다른 타입의 예외를 구분해서 처리하고 싶을 때 사용할 수 있습니다.

```php
<?php
class DBException extends Exception {}
class APIException extends Exception {}

try {
    throw new DBException("데이터베이스 연결 실패");
} catch (DBException $e) {
    echo "DB 에러: " . $e->getMessage();
} catch (APIException $e) {
    echo "API 에러: " . $e->getMessage();
} catch (Exception $e) {
    echo "일반 예외: " . $e->getMessage();
}
?>
```

✔ 가장 구체적인 예외부터 `catch`하도록 순서를 정하는 것이 좋습니다.

---

## 6️⃣ Throwable 인터페이스 (PHP 7+)

PHP 7부터 `Exception` 외에도 `Error` 객체도 **Throwable 인터페이스**를 통해 처리할 수 있습니다.

```php
<?php
try {
    throw new Error("심각한 에러 발생");
} catch (Throwable $e) {
    echo "모든 오류 처리: " . $e->getMessage();
}
?>
```

✔ `Throwable`은 `Exception`과 `Error`의 부모로, **모든 예외와 오류를 통합 처리**할 수 있습니다.

---

## 7️⃣ 예외 vs 오류 차이점

| 구분    | 설명                              |
|---------|-----------------------------------|
| `Exception` | 예측 가능한 에러, 개발자가 직접 처리 가능 |
| `Error`     | 시스템 수준의 치명적인 문제, PHP 7부터 객체로 처리 가능 |
| `Throwable` | `Exception`과 `Error`를 모두 포함하는 최상위 인터페이스 |

---

## 🎯 정리

✔ `try-catch` 문을 통해 예외 상황을 안전하게 처리할 수 있습니다.  
✔ `finally` 블록은 항상 실행되며, 정리 작업에 적합합니다.  
✔ `Exception` 클래스는 다양한 메서드로 에러 정보를 제공합니다.  
✔ 커스텀 예외 클래스를 만들어 보다 구체적인 처리 로직을 구성할 수 있습니다.  
✔ PHP 7 이상에서는 `Throwable`을 통해 예외와 에러를 통합적으로 관리할 수 있습니다.

