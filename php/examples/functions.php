<?php
// 1. 기본 함수 정의 및 호출
function sayHello() {
    echo "Hello, PHP!\n";
}
sayHello();

// 2. 매개변수(Parameter) 있는 함수
function greet($name) {
    echo "안녕하세요, $name 님!\n";
}
greet("철수");

// 3. 반환값(Return) 있는 함수
function add($a, $b) {
    return $a + $b;
}
$result = add(5, 10);
echo "5 + 10 = $result\n";

// 4. 기본값을 가진 매개변수
function introduce($name = "손님") {
    echo "반갑습니다, $name 님!\n";
}
introduce(); // 기본값 사용
introduce("영희");

// 5. 가변 인자 함수 (매개변수 개수가 정해지지 않은 함수)
function sumAll(...$numbers) {
    return array_sum($numbers);
}
echo "모든 숫자의 합: " . sumAll(1, 2, 3, 4, 5) . "\n";

// 6. 참조(Reference) 전달
function increase(&$num) {
    $num++;
}
$value = 10;
increase($value);
echo "참조 전달 후 값: $value\n"; // 11

// 7. 익명 함수 (Anonymous Function)
$greet = function($name) {
    return "Hello, $name!";
};
echo $greet("PHP") . "\n";

// 8. 화살표 함수 (Arrow Function, PHP 7.4 이상)
$multiply = fn($x, $y) => $x * $y;
echo "3 * 4 = " . $multiply(3, 4) . "\n";

// 9. 재귀 함수 (Recursive Function)
function factorial($n) {
    if ($n <= 1) return 1;
    return $n * factorial($n - 1);
}
echo "5! = " . factorial(5) . "\n";

// 10. 전역 변수와 지역 변수
$globalVar = "전역 변수";
function showVar() {
    global $globalVar;
    echo $globalVar . "\n";
}
showVar();

// 11. 정적 변수 (Static Variable)
function counter() {
    static $count = 0;
    $count++;
    echo "카운트: $count\n";
}
counter();
counter();
counter();
?>
