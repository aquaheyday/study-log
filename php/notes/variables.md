# 📦 PHP 변수와 데이터 타입

PHP는 동적 타이핑 언어로, 변수를 선언할 때 타입을 명시하지 않아도 됩니다.  
모든 값은 자동으로 타입이 할당되며, 다양한 데이터 타입과 관련 함수들을 제공합니다.

---

## 1️⃣ PHP 변수 개념

### 1) 변수 선언 방법
PHP에서 변수는 `$` 기호로 시작합니다.

```php
<?php
$name = "Alice";     // 문자열
$age = 30;           // 정수
$price = 19.99;      // 실수
$isValid = true;     // 불리언
?>
```

✔ 변수명은 문자 또는 `_`로 시작해야 하며 숫자로 시작할 수 없습니다.
✔ PHP는 **타입 선언 없이 변수 생성**이 가능합니다.

---

## 2️⃣ 주요 데이터 타입

PHP에서 자주 사용되는 **기본 타입(primitive types)** 들입니다.

| 타입       | 설명                  | 예시             |
|------------|-----------------------|------------------|
| string     | 문자열                 | `"hello"`        |
| int        | 정수                   | `42`             |
| float      | 실수 (`double`도 사용) | `3.14`           |
| bool       | 불리언 (참/거짓)        | `true`, `false`  |
| array      | 배열                   | `[1, 2, 3]`      |
| object     | 객체                   | `new ClassName()` |
| null       | 값 없음                | `null`           |

---

## 3️⃣ 타입 확인 및 디버깅 함수

#### 1. `var_dump()` - 값과 타입을 함께 출력

```php
<?php
$name = "Alice";
var_dump($name);      // string(5) "Alice"
?>
```

#### 2. `gettype()` - 타입만 문자열로 반환

```php
<?php
$name = "Alice";
echo gettype($name);  // string
?>
```

---

## 4️⃣ 명시적 형 변환 (Casting)

다른 타입으로 명시적으로 변환할 수 있습니다.

```php
<?php
$val = "10";
$intVal = (int)$val;       // 문자열 → 정수
$floatVal = (float)$val;   // 문자열 → 실수
$boolVal = (bool)$val;     // 문자열 → 불리언
?>
```

- `(int)`, `(float)`, `(bool)` 등의 형식으로 형 변환

---

## 5️⃣ 배열 (Array)

#### 1. 인덱스 배열

```php
<?php
$colors = ["red", "green", "blue"];
echo $colors[1]; // green
?>
```

#### 2. 연관 배열

```php
<?php
$user = ["name" => "Alice", "age" => 30];
echo $user["name"]; // Alice
?>
```

- 배열은 키-값 쌍으로 구성할 수도 있고, 인덱스를 자동으로 부여할 수도 있습니다.

---

## 6️⃣ 객체 (Object)

사용자 정의 클래스를 기반으로 객체를 생성할 수 있습니다.

```php
<?php
class Person {
    public $name = "Tom";
}

$p = new Person();
echo $p->name;  // Tom
?>
```

✔ `class`로 정의하고, `new` 키워드로 인스턴스를 생성합니다.  

---

## 7️⃣ NULL

변수에 아무런 값도 할당되지 않은 상태입니다.

```php
<?php
$val = null;
var_dump($val);  // NULL
?>
```

✔ 초기화하지 않은 변수, 혹은 `unset()`된 변수도 null 취급됩니다.  

---

## 8️⃣ 타입 관련 함수

| 함수                    | 설명                         |
|-------------------------|------------------------------|
| `is_string($var)`       | 문자열 여부 확인             |
| `is_int($var)`          | 정수 여부 확인               |
| `is_float($var)`        | 실수 여부 확인               |
| `is_bool($var)`         | 불리언 여부 확인             |
| `is_array($var)`        | 배열 여부 확인               |
| `is_object($var)`       | 객체 여부 확인               |
| `is_null($var)`         | null 여부 확인               |
| `gettype($var)`         | 타입 문자열 반환             |
| `settype($var, "type")` | 변수 타입을 직접 지정         |

---

## 🎯 정리

✔ PHP 변수는 `$` 기호로 선언되며 타입 선언이 필요 없음  
✔ 문자열, 정수, 실수, 불리언, 배열, 객체, null 등의 타입 제공  
✔ `var_dump()`와 `gettype()`으로 타입과 값을 쉽게 확인  
✔ 배열은 인덱스형과 연관형 둘 다 사용 가능  
✔ 객체는 클래스로 생성하며, `new`로 인스턴스를 만듦  
✔ `is_*()` 함수로 타입 검사 가능  

