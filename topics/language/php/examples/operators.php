<?php
// 1. 산술 연산자 (Arithmetic Operators)
$a = 10;
$b = 3;

echo "덧셈: " . ($a + $b) . "\n"; // 13
echo "뺄셈: " . ($a - $b) . "\n"; // 7
echo "곱셈: " . ($a * $b) . "\n"; // 30
echo "나눗셈: " . ($a / $b) . "\n"; // 3.3333
echo "나머지: " . ($a % $b) . "\n"; // 1

// 2. 할당 연산자 (Assignment Operators)
$x = 5;
$x += 3; // $x = $x + 3
echo "할당 연산 (+=): " . $x . "\n"; // 8

// 3. 비교 연산자 (Comparison Operators)
echo "비교 연산자 결과:\n";
var_dump($a == $b);  // false (값이 같으면 true)
var_dump($a != $b);  // true (값이 다르면 true)
var_dump($a > $b);   // true (a가 b보다 크면 true)
var_dump($a < $b);   // false (a가 b보다 작으면 true)

// 4. 논리 연산자 (Logical Operators)
$bool1 = true;
$bool2 = false;

echo "논리 연산자 결과:\n";
var_dump($bool1 && $bool2); // false (AND 연산)
var_dump($bool1 || $bool2); // true (OR 연산)
var_dump(!$bool1);          // false (NOT 연산)

// 5. 문자열 연산자 (String Operators)
$str1 = "Hello";
$str2 = " World";
echo "문자열 연결: " . $str1 . $str2 . "\n"; // "Hello World"

// 6. 삼항 연산자 (Ternary Operator)
$age = 20;
echo ($age >= 18) ? "성인입니다.\n" : "미성년자입니다.\n";

// 7. Null 병합 연산자 (Null Coalescing Operator)
$username = $_GET["username"] ?? "Guest";
echo "사용자: " . $username . "\n";
?>
