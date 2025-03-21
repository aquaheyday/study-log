# 🔁 PHP 제어문 기본

제어문(Control Structures)은 프로그램의 흐름을 제어하기 위해 사용됩니다.  
PHP에서는 조건에 따라 코드를 실행하거나 반복문을 통해 같은 작업을 여러 번 수행할 수 있습니다.

---

## 1️⃣ `if` / `else` / `elseif` 조건문

조건에 따라 다른 코드를 실행할 수 있는 **기본 조건문**입니다.

```php
<?php
$score = 85;

if ($score >= 90) {
    echo "A학점";
} elseif ($score >= 80) {
    echo "B학점";
} else {
    echo "C학점 이하";
}
?>
```

✔ `if`: 조건이 참일 때 실행  
✔ `elseif`: 이전 조건이 거짓이고 새로운 조건이 참이면 실행  
✔ `else`: 모든 조건이 거짓일 때 실행

---

## 2️⃣ `switch` 조건문

하나의 변수에 대해 **여러 값을 비교**할 때 사용합니다.

```php
<?php
$day = "tue";

switch ($day) {
    case "mon":
        echo "월요일";
        break;
    case "tue":
        echo "화요일";
        break;
    default:
        echo "기타 요일";
        break;
}
?>
```

✔ `case`: 일치하는 조건일 때 실행  
✔ `break`: 이후 case 실행을 막기 위해 반드시 사용  
✔ `default`: 어느 case에도 해당하지 않을 경우 실행

---

## 3️⃣ `while` 반복문

조건이 참일 동안 계속해서 코드를 반복 실행합니다.

```php
<?php
$i = 1;

while ($i <= 3) {
    echo $i;
    $i++;
}
?>
```

✔ **사전 조건 검사 방식**  
✔ 조건이 처음부터 거짓이면 한 번도 실행되지 않음

---

## 4️⃣ `do...while` 반복문

**한 번은 무조건 실행 후** 조건을 검사하는 반복문입니다.

```php
<?php
$i = 1;

do {
    echo $i;
    $i++;
} while ($i <= 3);
?>
```

✔ 최소 **한 번은 실행**되는 반복문  
✔ `while`과의 차이는 실행 순서

---

## 5️⃣ `for` 반복문

**초기값, 조건, 증감식**을 한 줄에 작성하는 반복문입니다.

```php
<?php
for ($i = 1; $i <= 3; $i++) {
    echo $i;
}
?>
```

✔ 반복 횟수가 명확할 때 가장 많이 사용

---

## 6️⃣ `foreach` 반복문

배열이나 컬렉션을 순회할 때 사용하는 반복문입니다.

#### 1. 값만 순회

```php
<?php
$colors = ["red", "green", "blue"];

foreach ($colors as $color) {
    echo $color;
}
?>
```

#### 2. 키-값 쌍 순회

```php
<?php
$user = ["name" => "Alice", "age" => 25];

foreach ($user as $key => $value) {
    echo "$key: $value";
}
?>
```

✔ 배열에 특화된 반복문  
✔ 객체 순회도 가능

---

## 7️⃣ `break` / `continue`

반복문 안에서 **제어 흐름을 변경**할 때 사용합니다.

```php
<?php
for ($i = 1; $i <= 5; $i++) {
    if ($i == 3) continue; // 3은 건너뜀
    if ($i == 5) break;    // 5에서 반복 종료
    echo $i;
}
?>
```

✔ `break`: 반복문 즉시 종료  
✔ `continue`: 현재 루프 건너뛰고 다음 반복 진행

---

## 8️⃣ `match` 표현식 (PHP 8+)

`switch`문보다 간결하고 **엄격한 비교(===)** 를 사용하는 새로운 분기 방식입니다.

```php
<?php
$status = 404;

echo match($status) {
    200 => "OK",
    404 => "Not Found",
    500 => "Server Error",
    default => "Unknown",
};
?>
```

✔ **표현식이므로 바로 값을 반환 가능**  
✔ `case`가 아닌 `조건 => 결과` 구조

---

## 🎯 정리

✔ 조건 분기: `if`, `else`, `elseif`, `switch`, `match (PHP 8+)`  
✔ 반복문: `while`, `do...while`, `for`, `foreach`  
✔ 루프 제어: `break`, `continue`  
✔ `foreach`는 배열이나 객체 순회에 가장 효율적  
✔ `match`는 간결하고 타입에 민감
