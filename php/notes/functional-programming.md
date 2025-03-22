# 🧠 PHP 고급 함수형 프로그래밍

PHP는 함수형 언어는 아니지만,  
**1급 함수(First-Class Functions)** 및 **익명 함수, 클로저, 콜백 함수, 고차 함수** 등을 통해  
**함수형 프로그래밍 스타일**을 활용할 수 있습니다.

---

## 1️⃣ 1급 함수 (First-Class Function)

PHP에서는 **함수를 변수에 할당하거나, 인자로 전달하거나, 반환할 수 있습니다.**

```php
<?php
$square = function($x) {
    return $x * $x;
};

echo $square(4); // 16
?>
```

✔ 함수도 변수처럼 다룰 수 있습니다.  
✔ 익명 함수(anonymous function)를 변수에 할당 가능  

---

## 2️⃣ 고차 함수 (Higher-Order Function)

**함수를 인자로 받거나 반환하는 함수**를 의미합니다.

```php
<?php
function apply($callback, $value) {
    return $callback($value);
}

echo apply(fn($x) => $x + 10, 5); // 15
?>
```

✔ `apply()`는 전달받은 함수를 실행합니다.  
✔ `fn()`은 PHP 7.4+에서 도입된 **화살표 함수(short closure)** 입니다.

---

## 3️⃣ 클로저 (Closure)

클로저는 **외부 변수를 캡처해서 사용할 수 있는 함수**입니다.

```php
<?php
function makeAdder($n) {
    return function($x) use ($n) {
        return $x + $n;
    };
}

$add5 = makeAdder(5);
echo $add5(10); // 15
?>
```

✔ `use` 키워드로 외부 변수를 클로저 안으로 가져옵니다.  
✔ 외부 변수는 복사됩니다 (참조하고 싶다면 `&$var` 사용)

---

## 4️⃣ 배열 고차 함수 활용 (`array_map`, `array_filter`, `array_reduce`)

### 1) `array_map()` — 배열의 각 요소를 변환

```php
<?php
$data = [1, 2, 3];
$squares = array_map(fn($x) => $x * $x, $data);
print_r($squares); // [1, 4, 9]
?>
```

---

### 2) `array_filter()` — 조건에 맞는 요소만 추출

```php
<?php
$even = array_filter($data, fn($x) => $x % 2 === 0);
print_r($even); // [1 => 2]
?>
```

---

### 3) `array_reduce()` — 누적 처리

```php
<?php
$sum = array_reduce($data, fn($carry, $x) => $carry + $x, 0);
echo $sum; // 6
?>
```

---

## 5️⃣ 커링(Currying) 흉내내기

PHP는 기본적으로 커링을 지원하지 않지만, 함수를 반환하는 방식으로 **흉내낼 수 있습니다.**

```php
<?php
function multiply($a) {
    return function($b) use ($a) {
        return $a * $b;
    };
}

$double = multiply(2);
echo $double(5); // 10
?>
```

✔ 커링은 **부분 적용(Partial Application)** 을 가능하게 합니다.

---

## 6️⃣ 콜백 함수와 call_user_func

PHP에서 함수 이름 또는 클로저를 **문자열이나 변수로 전달**해 호출할 수 있습니다.

```php
<?php
function greet($name) {
    return "Hello, $name";
}

echo call_user_func("greet", "Alice"); // Hello, Alice
?>
```

✔ `call_user_func()` → 동적으로 함수를 호출할 수 있도록 해줍니다.  
✔ 객체의 메서드 호출도 가능: `call_user_func([$obj, 'method'])`  

---

## 7️⃣ 재귀 함수 (Recursive Function)

함수가 **자기 자신을 호출**할 수 있습니다.

```php
<?php
function factorial($n) {
    if ($n <= 1) return 1;
    return $n * factorial($n - 1);
}

echo factorial(5); // 120
?>
```

✔ 조건문으로 종료 조건을 잘 설정해주어야 합니다.  
✔ 꼬리 재귀 최적화(TCO)는 PHP에서 지원되지 않지만, 작은 범위에서는 충분히 유용합니다.

---

## 8️⃣ 함수형 프로그래밍의 이점

| 장점                           | 설명                                         |
|--------------------------------|----------------------------------------------|
| **가독성 향상**                | 선언적 코드 스타일로 코드가 더 명확해짐      |
| **재사용성 증가**              | 작은 순수 함수를 조합하여 다양한 동작 구성 가능 |
| **버그 발생 가능성 감소**     | 상태 변경을 줄이고 부작용(Side Effect) 최소화 |
| **테스트 용이성**              | 입력과 출력이 명확하므로 테스트가 쉬움       |

---

## 🎯 정리

✔ PHP는 1급 함수, 클로저, 고차 함수 등을 통해 함수형 스타일을 구현할 수 있습니다.  
✔ `array_map`, `array_filter`, `array_reduce`는 가장 실용적인 함수형 도구입니다.  
✔ 커링, 클로저, call_user_func 같은 기법을 활용하면 유연한 구조를 만들 수 있습니다.  
✔ 함수형 프로그래밍은 **작고 순수한 함수**들의 조합으로 복잡한 로직을 단순화할 수 있습니다.

