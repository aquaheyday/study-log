# PHP 배열과 연관 배열

## 1. 배열(Array) 생성 및 사용
```php
// 기본 배열 생성
$numbers = array(1, 2, 3, 4, 5);
$fruits = ["Apple", "Banana", "Cherry"];

// 배열 요소 접근
echo $numbers[0]; // 1
echo $fruits[1]; // Banana
```

## 2. 배열 요소 추가 및 제거
```php
$fruits[] = "Orange"; // 배열에 요소 추가
array_push($fruits, "Grapes"); // 끝에 추가
array_unshift($fruits, "Mango"); // 앞에 추가

array_pop($fruits); // 마지막 요소 제거
array_shift($fruits); // 첫 번째 요소 제거
```

## 3. 연관 배열(Associative Array)
```php
// 키-값 형태의 연관 배열 생성
$person = [
    "name" => "Alice",
    "age" => 25,
    "city" => "New York"
];

// 연관 배열 요소 접근
echo $person["name"]; // Alice
```

## 4. 배열 순회
```php
// 인덱스 배열 반복문
foreach ($fruits as $fruit) {
    echo $fruit . "\n";
}

// 연관 배열 반복문
foreach ($person as $key => $value) {
    echo "$key: $value\n";
}
```

## 5. 배열 함수 활용
```php
$numbers = [5, 1, 8, 3, 9];

sort($numbers); // 오름차순 정렬
rsort($numbers); // 내림차순 정렬

$keys = array_keys($person); // 키 배열 가져오기
$values = array_values($person); // 값 배열 가져오기

if (in_array(25, $values)) {
    echo "Value exists!";
}
```

## 6. 다차원 배열(Multidimensional Array)
```php
$students = [
    ["name" => "Alice", "age" => 20],
    ["name" => "Bob", "age" => 22],
    ["name" => "Charlie", "age" => 21]
];

echo $students[1]["name"]; // Bob
```
