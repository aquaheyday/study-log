# 🧱 PHP 클래스와 객체

PHP는 객체지향 프로그래밍(OOP)을 지원하는 언어입니다.  
객체지향은 프로그램을 구성할 때 **현실 세계처럼 "객체"로 모델링**하고, 그 객체들이 서로 상호작용하도록 만드는 방식입니다.

PHP에서는 클래스(class)와 객체(object)를 통해 이러한 구조를 구현할 수 있습니다.

---

## 1️⃣ 클래스와 객체의 기본

### 1) 클래스 정의 및 객체 생성
- `class` 키워드를 사용하여 클래스를 정의합니다.  
- `new` 키워드를 사용하여 클래스로부터 객체를 생성할 수 있습니다.  
- `$this`는 현재 객체 자신을 가리킬 때 사용합니다.

```php
<?php
class Car {
    public $brand = "Hyundai";

    public function drive() {
        echo "Driving {$this->brand}";
    }
}

$myCar = new Car();
$myCar->drive(); // Driving Hyundai
?>
```

---

## 2️⃣ 생성자 (Constructor)

생성자는 객체가 생성될 때 자동으로 호출되는 메서드입니다.  
주로 **초기값을 설정할 때 사용**합니다.

```php
<?php
class User {
    public $name;

    public function __construct($name) {
        $this->name = $name;
    }
}

$user = new User("Alice");
echo $user->name; // Alice
?>
```

✔ `__construct()` 메서드를 통해 객체 초기화가 가능합니다.

---

## 3️⃣ 접근 제어자 (Access Modifiers)

클래스 내부의 속성과 메서드에 대해 **외부에서 접근 가능한 범위를 지정**할 수 있습니다.

| 제어자        | 설명                                               |
|---------------|----------------------------------------------------|
| `public`      | 어디서든 접근 가능                                 |
| `protected`   | 현재 클래스 및 상속받은 클래스 내에서만 접근 가능 |
| `private`     | 현재 클래스 내부에서만 접근 가능                   |

```php
<?php
class Example {
    public $a = 1;
    protected $b = 2;
    private $c = 3;

    public function showAll() {
        echo $this->a; // 접근 가능
        echo $this->b; // 접근 가능
        echo $this->c; // 접근 가능
    }
}
?>
```

✔ 클래스 외부에서는 `protected`, `private` 속성에 직접 접근할 수 없습니다.

---

## 4️⃣ Getter & Setter

`private`으로 선언된 속성은 외부에서 직접 접근할 수 없기 때문에, **Getter와 Setter 메서드를 통해 안전하게 값을 다룹니다.**

```php
<?php
class Product {
    private $price;

    public function setPrice($p) {
        $this->price = $p;
    }

    public function getPrice() {
        return $this->price;
    }
}

$p = new Product();
$p->setPrice(1000);
echo $p->getPrice(); // 1000
?>
```

✔ 객체의 속성을 외부에서 제어할 때 **보안성과 유연성**을 높이기 위해 사용됩니다.

---

## 5️⃣ 상속 (Inheritance)

상속은 **기존 클래스의 기능을 다른 클래스가 물려받을 수 있는 기능**입니다.

```php
<?php
class Animal {
    public function speak() {
        echo "동물의 소리";
    }
}

class Dog extends Animal {
    public function bark() {
        echo "멍멍";
    }
}

$d = new Dog();
$d->speak(); // 동물의 소리
$d->bark();  // 멍멍
?>
```

✔ `extends` 키워드를 사용하여 상속을 구현합니다.  
✔ 자식 클래스는 부모 클래스의 `public`, `protected` 속성과 메서드를 사용할 수 있습니다.

---

## 6️⃣ 메서드 오버라이딩 (Method Overriding)

상속받은 메서드를 자식 클래스에서 **다시 정의(재정의)** 할 수 있습니다.

```php
<?php
class Animal {
    public function sound() {
        echo "동물 소리";
    }
}

class Cat extends Animal {
    public function sound() {
        echo "야옹";
    }
}

$c = new Cat();
$c->sound(); // 야옹
?>
```

✔ 부모 클래스의 기능을 **특정 목적에 맞게 바꾸고 싶을 때** 유용합니다.

---

## 7️⃣ static 키워드

`static` 키워드를 사용하면 **객체를 생성하지 않고도 클래스 단위로 메서드나 속성에 접근**할 수 있습니다.

```php
<?php
class Math {
    public static $pi = 3.14;

    public static function square($n) {
        return $n * $n;
    }
}

echo Math::$pi;          // 3.14
echo Math::square(4);    // 16
?>
```

✔ 클래스 내부에서는 `self::` 키워드를 사용하여 접근합니다.  
✔ `$this`는 사용할 수 없습니다 (객체 인스턴스가 없기 때문입니다).

---

## 8️⃣ 추상 클래스 (abstract class)

추상 클래스는 **공통된 구조를 정의하고, 구현은 자식 클래스에게 맡기는 클래스**입니다.  
직접 객체를 만들 수 없고, 반드시 상속해서 사용해야 합니다.

```php
<?php
abstract class Shape {
    abstract public function area();
}

class Circle extends Shape {
    public function area() {
        return 3.14 * 5 * 5;
    }
}
?>
```

✔ 자식 클래스는 반드시 추상 메서드를 구현해야 합니다.

---

## 9️⃣ 인터페이스 (interface)

인터페이스는 클래스가 **반드시 구현해야 하는 메서드의 목록(계약서)** 를 정의합니다.

```php
<?php
interface Logger {
    public function log($message);
}

class FileLogger implements Logger {
    public function log($message) {
        echo "Log: $message";
    }
}
?>
```

✔ `implements` 키워드를 사용하여 구현합니다.  
✔ 하나의 클래스는 여러 인터페이스를 구현할 수 있습니다.

---

## 🔟 클래스 상수 및 final 키워드

### 1) 클래스 상수

변경할 수 없는 고정된 값을 정의할 때 사용합니다.

```php
<?php
class Config {
    const VERSION = "1.0";
}

echo Config::VERSION;
?>
```

✔ 클래스 상수는 `::` 연산자를 통해 접근합니다.

---

### 2) final 키워드

`final`을 사용하면 **클래스의 상속이나 메서드의 오버라이딩을 막을 수 있습니다.**

```php
<?php
final class Base {} // 상속 불가

class Sample {
    final public function run() {
        echo "오버라이딩 불가";
    }
}
?>
```

✔ 변경되면 안 되는 기능에 대해 의도적으로 제한을 둘 수 있습니다.

---

## 🎯 정리

✔ `class` 키워드로 클래스를 정의하고 `new` 키워드로 객체를 생성합니다.  
✔ `__construct()`를 사용하여 객체 생성 시 초기화할 수 있습니다.  
✔ `public`, `protected`, `private` 접근 제어자를 통해 속성과 메서드의 접근 범위를 제어할 수 있습니다.  
✔ Getter/Setter 메서드를 통해 `private` 속성에 안전하게 접근합니다.  
✔ 상속(`extends`)과 오버라이딩을 통해 재사용성과 확장성을 확보할 수 있습니다.  
✔ `static`은 클래스 자체에서 직접 접근 가능한 속성/메서드를 만들 때 사용합니다.  
✔ 추상 클래스(`abstract`)와 인터페이스(`interface`)는 구조 설계 시 유용하게 사용됩니다.  
✔ `final` 키워드를 통해 더 이상 상속하거나 재정의하지 못하게 제한할 수 있습니다.

