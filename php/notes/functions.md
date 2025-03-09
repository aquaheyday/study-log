# PHP 함수와 스코프

## 1. 함수 정의 및 호출
```php
function sayHello($name) {
    echo "Hello, $name!\n";
}

sayHello("Alice"); // Hello, Alice!
```

## 2. 지역 변수(Local Scope)
```php
function localScopeExample() {
    $localVar = "I am local";
    echo $localVar . "\n";
}

localScopeExample();
// echo $localVar; // 오류 발생: $localVar는 함수 내부에서만 접근 가능
```

## 3. 전역 변수(Global Scope)
```php
$globalVar = "I am global";

function globalScopeExample() {
    global $globalVar; // global 키워드를 사용하여 접근 가능
    echo $globalVar . "\n";
}

globalScopeExample();
```

## 4. 정적 변수(Static Scope)
```php
function staticExample() {
    static $count = 0; // 함수 호출 시 값을 유지함
    $count++;
    echo "Count: $count\n";
}

staticExample(); // Count: 1
staticExample(); // Count: 2
staticExample(); // Count: 3
```

## 5. 함수 내 변수 범위 확인
```php
function checkScope() {
    $insideVar = "Inside Function";
    echo "Inside Function: $insideVar\n";
}

checkScope();
// echo $insideVar; // 오류 발생: 함수 외부에서는 접근 불가
```

## 6. 익명 함수(Anonymous Function) 및 클로저(Closure)
```php
$greet = function($name) {
    return "Hello, $name!";
};

echo $greet("Bob") . "\n";
```

## 7. 클로저(Closure)에서 외부 변수 사용 (use 키워드 활용)
```php
$message = "Good morning";

$customGreet = function($name) use ($message) {
    return "$message, $name!";
};

echo $customGreet("Charlie") . "\n";
```
