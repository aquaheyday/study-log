# 🧩 PHP 네임스페이스와 오토로딩

PHP에서 규모가 큰 프로젝트를 개발할 때는 **클래스 이름의 충돌을 방지**하고  
**코드의 구조를 체계적으로 유지**하기 위해 **네임스페이스(Namespace)** 와 **오토로딩(Autoloading)** 기능을 활용합니다.

---

## 1️⃣ 네임스페이스(Namespace)란?

네임스페이스는 클래스, 함수, 상수의 **이름 충돌을 방지하기 위한 논리적인 구역**입니다.  
**패키지 구조처럼 디렉토리별로 코드 영역을 구분**할 수 있습니다.

#### 기본 사용법

```php
<?php
namespace App\Models;

class User {
    public function getName() {
        return "홍길동";
    }
}
?>
```

✔ `namespace App\Models;` 구문은 이 파일의 클래스가 **`App\Models` 네임스페이스**에 속한다는 의미입니다.  
✔ 네임스페이스는 일반적으로 **디렉토리 구조와 매칭**하는 것이 좋습니다.  

---

## 2️⃣ 네임스페이스 클래스 사용하기

네임스페이스가 선언된 클래스를 사용할 때는 **전체 경로(FQCN: Fully Qualified Class Name)** 로 참조하거나,  
`use` 키워드를 사용하여 간단하게 불러올 수 있습니다.

### 1) 전체 경로로 접근

```php
<?php
require 'App/Models/User.php';

$user = new \App\Models\User();
echo $user->getName();
?>
```

---

### 2) `use` 키워드로 가져오기

```php
<?php
require 'App/Models/User.php';

use App\Models\User;

$user = new User();
?>
```

✔ `use` 키워드를 사용하면 긴 네임스페이스를 매번 쓰지 않아도 됩니다.  

---

## 3️⃣ 중첩 네임스페이스와 별칭 사용

### 1) 중첩 구조
`App\Controllers\Admin` 처럼 중첩 네임스페이스도 가능합니다.

```php
namespace App\Controllers\Admin;
```

---

### 2) 별칭(alias) 사용
`as` 키워드를 사용하면 다른 이름으로 사용할 수 있어 네임 충돌을 피할 수 있습니다.

```php
use App\Models\User as Member;

$member = new Member();
```

---

## 4️⃣ 오토로딩(Autoloading) 이란?

오토로딩은 **클래스를 사용할 때 자동으로 필요한 파일을 불러오는 기능**입니다.  
매번 `require` 나 `include`를 직접 작성하지 않아도 됩니다.

---

## 5️⃣ PSR-4 오토로딩 (Composer 기준)

**PSR-4**는 PHP의 **표준 클래스 자동 로딩 규칙**입니다.  
주로 **Composer**를 통해 오토로딩을 설정합니다.

#### 1. 예시 디렉토리 구조

```
project/
├── src/
│   └── Models/
│       └── User.php
├── vendor/
└── composer.json
```

#### 2. src/Models/User.php

```php
<?php
namespace App\Models;

class User {
    public function sayHi() {
        return "안녕하세요!";
    }
}
?>
```

---

## 6️⃣ composer.json 설정
- `"App\\"` 네임스페이스는 `src/` 디렉토리와 연결됩니다.
- 반드시 `composer dump-autoload` 명령으로 적용해야 합니다.

```json
{
  "autoload": {
    "psr-4": {
      "App\\": "src/"
    }
  }
}
```

---

## 7️⃣ Autoload 테스트 코드

```php
<?php
require 'vendor/autoload.php';

use App\Models\User;

$user = new User();
echo $user->sayHi();  // 안녕하세요!
?>
```

✔ 클래스를 사용하는 시점에 자동으로 로드되므로 `require` 문이 필요하지 않습니다.

---

## 8️⃣ `spl_autoload_register()` (수동 오토로딩)

Composer 없이 오토로딩을 직접 구현하고 싶을 때는 `spl_autoload_register()` 함수를 사용할 수 있습니다.

```php
<?php
spl_autoload_register(function ($class) {
    $path = str_replace("\\", "/", $class) . ".php";
    require $path;
});

$user = new App\Models\User();
?>
```

✔ 네임스페이스를 경로로 변환해 자동으로 파일을 불러옵니다.  
✔ 실무에서는 Composer를 사용하는 것이 일반적입니다.  

---

## 🎯 정리

✔ `namespace` 키워드를 사용하여 코드 영역을 논리적으로 구분할 수 있습니다.  
✔ `use`를 통해 네임스페이스 클래스의 경로를 짧게 사용할 수 있으며, `as` 키워드로 별칭도 지정 가능합니다.  
✔ 오토로딩은 클래스를 사용하는 시점에 자동으로 파일을 불러오는 기능이며
