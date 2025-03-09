# PHP 세션과 쿠키

## 1. 세션(Session) 개요
세션은 사용자의 정보를 서버에 저장하여 일정 기간 동안 유지하는 기능입니다.

### 1.1 세션 시작 (`session_start()`)
```php
session_start(); // 세션 시작
$_SESSION["username"] = "Alice";
echo "세션이 설정되었습니다.";
```

### 1.2 세션 값 가져오기
```php
session_start(); // 세션 시작
if (isset($_SESSION["username"])) {
    echo "사용자: " . $_SESSION["username"];
} else {
    echo "세션이 존재하지 않습니다.";
}
```

### 1.3 세션 삭제 (`unset`, `session_destroy`)
```php
session_start();
unset($_SESSION["username"]); // 특정 세션 변수 제거
session_destroy(); // 모든 세션 삭제
echo "세션이 삭제되었습니다.";
```

---

## 2. 쿠키(Cookie) 개요
쿠키는 사용자의 브라우저에 데이터를 저장하는 기능입니다.

### 2.1 쿠키 설정 (`setcookie`)
```php
setcookie("user", "Alice", time() + 3600, "/"); // 1시간 유지
echo "쿠키가 설정되었습니다.";
```

### 2.2 쿠키 값 가져오기
```php
if (isset($_COOKIE["user"])) {
    echo "쿠키 값: " . $_COOKIE["user"];
} else {
    echo "쿠키가 설정되지 않았습니다.";
}
```

### 2.3 쿠키 삭제
```php
setcookie("user", "", time() - 3600, "/"); // 쿠키 삭제
echo "쿠키가 삭제되었습니다.";
```

---

## 3. 세션과 쿠키 비교
| 항목  | 세션(Session) | 쿠키(Cookie) |
|--------|--------------|--------------|
| 저장 위치 | 서버 | 클라이언트(브라우저) |
| 지속 시간 | 브라우저 종료 시 기본 삭제 (설정 가능) | 만료 시간 설정 가능 |
| 보안 | 상대적으로 안전 | 보안에 취약 (암호화 필요) |
| 용도 | 로그인 정보, 사용자 데이터 유지 | 간단한 사용자 정보 저장 |

---

### 🔹 세션과 쿠키를 함께 활용하는 예제
```php
session_start();
setcookie("user", "Alice", time() + 3600, "/");
$_SESSION["username"] = "Alice";

echo "세션과 쿠키가 설정되었습니다.";
```
