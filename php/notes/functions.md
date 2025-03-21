# 📄 PHP 함수 사용법

함수(Function)는 **반복되는 코드를 재사용**하고, **코드를 모듈화**하여 유지보수를 쉽게 해줍니다.  
PHP에서는 `function` 키워드를 사용하여 함수를 정의합니다.

---

## 1️⃣ 함수 정의 및 호출
- `function` 키워드로 정의  
- 정의된 함수는 이름으로 호출 가능

#### 기본 함수 정의 및 호출

```php
<?php
function sayHello() {
    echo "Hello, PHP!";
}

sayHello(); // 함수 호출
?>
```

---

## 2️⃣ 매개변수(Parameter) 전달

#### 값 전달

```php
<?php
function greet($name) {
    echo "Hello, $name!";
}

greet("Alice");
?>
```

✔ `$name`은 매개변수(Parameter)  
✔ 호출 시 전달하는 `"Alice"`는 인자(Argument)

---

## 3️⃣ 기본값 매개변수 (Default Parameter)

인자를 전달하지 않을 경우 **기본값** 사용 가능

```php
<?php
function greet($name = "Guest") {
    echo "Hello, $name!";
}

greet();         // Hello, Guest!
greet("David");  // Hello, David!
?>
```

✔ 기본값은 오른쪽 매개변수부터 설정하는 것이 원칙

---

## 4️⃣ 반환값(Return Value)

`return` 키워드를 사용해 값을 반환할 수 있습니다.

```php
<?php
function add($a, $b) {
    return $a + $b;
}

$result = add(3, 4);  // 7
echo $result;
?>
```

✔ 함수는 항상 하나의 값을 반환하며, `return` 이후 코드는 실행되지 않음

---

## 5️⃣ 가변 인자 함수 (Variable-Length Arguments)

매개변수의 수가 **불확정적일 경우** 사용합니다. (PHP 5.6+)

```php
<?php
function sum(...$numbers) {
    return array_sum($numbers);
}

echo sum(1, 2, 3, 4);  // 10
?>
```

✔ `...` 문법으로 가변 인자 사용 가능  
✔ 인자는 배열로 받게 됨

---

## 6️⃣ 참조에 의한 전달 (Call by Reference)

매개변수를 **원본에 영향을 주도록 전달**합니다.

```php
<?php
function addOne(&$num) {
    $num++;
}

$x = 5;
addOne($x);
echo $x;  // 6
?>
```

✔ `&` 기호를 사용하여 참조로 전달  
✔ 함수 내에서 값이 직접 변경됨

---

## 7️⃣ 함수 안의 변수와 스코프

### 1) 지역 변수(Local Variable)

함수 내부에서 선언한 변수는 **외부에서 접근 불가**합니다.

```php
<?php
function test() {
    $x = 10; // 지역 변수
}
echo $x; // 오류
?>
```

---

### 2) 전역 변수(Global Variable)

`global` 키워드를 사용해 함수 내에서 전역 변수 접근

```php
<?php
$x = 10;

function showX() {
    global $x;
    echo $x;
}
?>
```

---

## 8️⃣ 익명 함수 / 클로저

함수를 변수에 할당하거나, 콜백으로 전달 가능

```php
<?php
$greet = function($name) {
    return "Hello, $name!";
};

echo $greet("PHP");
?>
```

✔ PHP에서 함수도 **일급 객체**처럼 사용할 수 있음  
✔ `use` 키워드로 외부 변수 캡처 가능 (클로저)

---

## 9️⃣ 재귀 함수 (Recursive Function)

함수 내부에서 **자기 자신을 호출**하는 함수

```php
<?php
function factorial($n) {
    if ($n <= 1) return 1;
    return $n * factorial($n - 1);
}

echo factorial(5);  // 120
?>
```

✔ 종료 조건이 없으면 **무한 루프** 발생 주의

---

## 🎯 정리

✔ `function` 키워드로 함수 정의, `()`로 호출  
✔ 매개변수는 기본값 지정 가능  
✔ `return`으로 결과 반환  
✔ `...`로 가변 인자 처리 가능  
✔ `&`를 사용하면 참조에 의한 전달  
✔ `global` 키워드로 전역 변수 접근  
✔ 익명 함수(클로저)는 변수에 저장하거나 콜백으로 활용 가능  
✔ 재귀 함수는 반드시 종료 조건을 포함해야 함  

