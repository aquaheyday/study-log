<?php
// 1. if-else 문 (조건문)
$age = 20;

if ($age >= 18) {
    echo "성인입니다.\n";
} else {
    echo "미성년자입니다.\n";
}

// 2. if-elseif-else 문
$score = 85;

if ($score >= 90) {
    echo "학점: A\n";
} elseif ($score >= 80) {
    echo "학점: B\n";
} elseif ($score >= 70) {
    echo "학점: C\n";
} else {
    echo "학점: F\n";
}

// 3. switch 문 (다중 조건문)
$day = "월요일";

switch ($day) {
    case "월요일":
        echo "오늘은 월요일입니다.\n";
        break;
    case "화요일":
        echo "오늘은 화요일입니다.\n";
        break;
    default:
        echo "오늘은 평범한 하루입니다.\n";
}

// 4. for 문 (반복문)
for ($i = 1; $i <= 5; $i++) {
    echo "for 반복문: 숫자 $i\n";
}

// 5. while 문
$count = 1;
while ($count <= 3) {
    echo "while 반복문: 숫자 $count\n";
    $count++;
}

// 6. do-while 문
$num = 1;
do {
    echo "do-while 반복문: 숫자 $num\n";
    $num++;
} while ($num <= 3);

// 7. foreach 문 (배열 반복)
$fruits = array("사과", "바나나", "체리");

foreach ($fruits as $fruit) {
    echo "과일: $fruit\n";
}
?>
