# 🐘 PHP 기본 문법 정리

PHP 프로그래밍의 기초가 되는 문법들을 정리한 문서입니다.  
기본적인 파일 구성, 변수 선언, 자료형, 제어문까지 학습합니다.

---

## 1. PHP 기본 구조

### 파일 확장자
- PHP 파일은 `.php` 확장자 사용

### 기본 구조 예시
```php
<?php
// PHP 코드는 <?php로 시작
echo "Hello, PHP!";
?>
```

---

## 2. 출력문 (echo, print)

### 기본 출력

```php
echo "안녕하세요!";
print "PHP 공부 시작!";
```

### HTML 출력도 가능

```php
echo "<h1>Welcome to PHP!</h1>";
```
---

## 3. 변수 선언과 사용

```php
$name = "홍길동";
$age = 25;
echo "이름: $name, 나이: $age 세";
```
---

## 4. 자료형

| 자료형 | 설명 | 예시 |
|---|---|---|
| 문자열 | 텍스트 데이터 | `$str = "Hello";` |
| 정수 | 정수형 숫자 | `$num = 100;` |
| 실수 | 소수점 있는 숫자 | `$pi = 3.14;` |
| 불리언 | 참/거짓 | `$isActive = true;` |
| 배열 | 여러 값 저장 | `$colors = ["red", "green", "blue"];` |
| 객체 | 클래스 인스턴스 | `$user = new User();` |
| NULL | 값 없음 | `$value = null;` |

---

## 5. 연산자

# 🐘 PHP 연산자 정리표

| 종류 | 기호 | 설명 | 예시 |
|---|---|---|---|
| 대입 | = | 값 대입 | `$a = 10;` |
| 산술 | +, -, *, /, % | 사칙연산 | `$sum = $a + $b;` |
| 비교 | ==, ===, !=, <, > | 값 비교 | `if ($a == $b)` |
| 논리 | &&, \|\|, ! | 논리 연산 | `if ($a > 0 && $b > 0)` |
| 증감 | ++, -- | 값 증가/감소 | `$i++` |
| 문자열 | . | 문자열 연결 | `$fullName = $firstName . " " . $lastName` |
| 삼항 | ?: | 조건부 값 반환 | `$result = ($score >= 60) ? "Pass" : "Fail"` |

---

## 6. 제어문

### if 조건문

```php
$score = 85;

if ($score >= 90) {
    echo "A학점";
} elseif ($score >= 80) {
    echo "B학점";
} else {
    echo "C학점";
}
```

### switch 조건문

```php
$menu = "coffee";

switch ($menu) {
    case "coffee":
        echo "커피 주문";
        break;
    case "tea":
        echo "차 주문";
        break;
    default:
        echo "메뉴 없음";
}
```

---

## 7. 반복문

### for

```php
for ($i = 1; $i <= 5; $i++) {
    echo "숫자: $i<br>";
}
```

### while

```php
$count = 1;

while ($count <= 3) {
    echo "Count: $count<br>";
    $count++;
}
```

---

## 8. 주석

```php
// 한 줄 주석

/*
여러 줄 주석
여러 줄 설명 가능
*/
```

---

## 9. 팁 & 주의사항

- 반드시 <?php 로 시작
- 파일 마지막 ?>는 가급적 생략 (빈줄 출력 문제 방지)
- 문자열 출력 시 큰따옴표 " " 사용 추천
- 한 줄 끝에는 세미콜론 ; 필수
