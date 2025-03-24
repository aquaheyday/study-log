<?php
// 1. 배열 생성 (Indexed Array)
$fruits = array("사과", "바나나", "체리");

// 배열 요소 출력
echo "첫 번째 과일: " . $fruits[0] . "\n"; // 사과

// 배열 길이 구하기
echo "과일 개수: " . count($fruits) . "\n"; // 3

// 2. 배열 요소 추가 및 수정
$fruits[] = "포도"; // 배열 끝에 추가
$fruits[1] = "오렌지"; // 기존 요소 수정

// 3. 반복문을 사용한 배열 출력 (foreach)
foreach ($fruits as $fruit) {
    echo "과일: $fruit\n";
}

// 4. 연관 배열 (Associative Array)
$person = array(
    "이름" => "김철수",
    "나이" => 25,
    "직업" => "개발자"
);

// 연관 배열 요소 출력
echo "이름: " . $person["이름"] . "\n"; // 김철수
echo "나이: " . $person["나이"] . "\n"; // 25

// 연관 배열 수정
$person["직업"] = "디자이너";
echo "수정된 직업: " . $person["직업"] . "\n"; // 디자이너

// 5. 연관 배열 반복 출력 (foreach)
foreach ($person as $key => $value) {
    echo "$key: $value\n";
}

// 6. 다차원 배열 (Multidimensional Array)
$students = array(
    array("이름" => "영희", "나이" => 20, "점수" => 90),
    array("이름" => "철수", "나이" => 22, "점수" => 85),
    array("이름" => "민수", "나이" => 21, "점수" => 88)
);

// 다차원 배열 요소 출력
echo "첫 번째 학생 이름: " . $students[0]["이름"] . "\n"; // 영희

// 7. 배열 정렬
$numbers = array(5, 2, 8, 1, 3);
sort($numbers); // 오름차순 정렬
echo "오름차순 정렬된 숫자: " . implode(", ", $numbers) . "\n";

rsort($numbers); // 내림차순 정렬
echo "내림차순 정렬된 숫자: " . implode(", ", $numbers) . "\n";

// 8. 키 정렬 (연관 배열)
$ages = array("철수" => 25, "영희" => 22, "민수" => 28);
asort($ages); // 값 기준 오름차순 정렬
echo "나이 오름차순 정렬:\n";
foreach ($ages as $name => $age) {
    echo "$name: $age\n";
}

ksort($ages); // 키 기준 오름차순 정렬
echo "이름 오름차순 정렬:\n";
foreach ($ages as $name => $age) {
    echo "$name: $age\n";
}
?>
