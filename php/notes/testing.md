# ✅ PHP 단위 테스트 (Unit Testing)

**단위 테스트(Unit Test)** 는 애플리케이션의 기능을 작은 단위로 나누어 **각 기능이 정확히 동작하는지 자동으로 검증**하는 과정입니다.  
PHP에서는 대표적으로 **PHPUnit**을 사용하며, 클래스/함수의 입력과 출력이 올바른지 확인할 수 있습니다.  

---

## 1️⃣ PHPUnit 설치

#### 1. Composer로 설치

```bash
composer require --dev phpunit/phpunit
```

#### 2. 설치 확인

```bash
./vendor/bin/phpunit --version
```

✔ 글로벌 설치 대신 **프로젝트별 설치**가 권장됩니다.  

---

## 2️⃣ 테스트 파일 구조

일반적으로 `tests/` 디렉토리에 테스트 파일을 구성합니다.  

```text
project/
├── src/
│   └── Calculator.php
└── tests/
    └── CalculatorTest.php
```

---

## 3️⃣ 테스트 클래스 기본 구조

```php
<?php
use PHPUnit\Framework\TestCase;

class CalculatorTest extends TestCase
{
    public function testAddition()
    {
        $this->assertEquals(4, 2 + 2);
    }
}
?>
```

✔ 모든 테스트 메서드는 `test`로 시작해야 인식됩니다.  
✔ `assert*` 메서드로 결과를 검증합니다.  

---

## 4️⃣ 테스트 실행 방법

```bash
./vendor/bin/phpunit tests
```

또는 특정 파일만:

```bash
./vendor/bin/phpunit tests/CalculatorTest.php
```

---

## 5️⃣ 주요 Assertion 메서드

| 메서드                          | 설명                                  |
|----------------------------------|---------------------------------------|
| `assertEquals($a, $b)`          | 두 값이 같은지 비교                   |
| `assertTrue($cond)`             | 조건이 true 인지 확인                |
| `assertFalse($cond)`            | 조건이 false 인지 확인               |
| `assertNull($val)`              | 값이 null 인지 확인                  |
| `assertCount($count, $array)`   | 배열/컬렉션 요소 수 비교             |
| `assertInstanceOf($class, $obj)`| 객체가 특정 클래스의 인스턴스인지 확인 |

---

## 6️⃣ 테스트 대상 클래스 예시

```php
// src/Calculator.php
<?php
class Calculator {
    public function add($a, $b) {
        return $a + $b;
    }
}
?>
```

#### 테스트 클래스

```php
// tests/CalculatorTest.php
<?php
use PHPUnit\Framework\TestCase;

require_once __DIR__ . '/../src/Calculator.php';

class CalculatorTest extends TestCase {
    public function testAdd() {
        $calc = new Calculator();
        $this->assertEquals(5, $calc->add(2, 3));
    }
}
?>
```

---

## 7️⃣ 테스트 자동 실행 (watch 모드)

테스트를 자동화하려면 **`phpunit-watcher`** 같은 도구 사용도 가능합니다.

```bash
composer require --dev spatie/phpunit-watcher
./vendor/bin/phpunit-watcher watch
```

✔ 코드 수정 시마다 자동으로 테스트 재실행

---

## 8️⃣ 테스트 커버리지 측정

PHPUnit은 코드 커버리지 보고서도 지원합니다 (Xdebug 필요).

```bash
./vendor/bin/phpunit --coverage-html coverage
```

✔ `coverage/` 디렉토리에 **HTML 보고서 생성**

---

## 🎯 정리

✔ 단위 테스트는 클래스/함수 단위의 자동화된 검증 도구입니다.  
✔ PHPUnit을 사용하면 `assertEquals()`, `assertTrue()` 등으로 **결과를 손쉽게 비교**할 수 있습니다.  
✔ `tests/` 폴더에 테스트 클래스를 만들어 실행하고, 커버리지 확인도 가능합니다.  
✔ CI/CD 파이프라인에 포함하면 **자동화된 품질 관리**가 가능해집니다.

