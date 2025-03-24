# PHP 기타 유용한 팁

## 1. 삼항 연산자 (Ternary Operator)
```php
// 일반적인 if 문
$age = 20;
$status = ($age >= 18) ? "성인" : "미성년자";
echo $status; // 성인
```

## 2. Null 병합 연산자 (??)
```php
$name = $_GET['name'] ?? 'Guest';
echo "Hello, " . $name; // name이 없으면 'Guest' 출력
```

## 3. 문자열 내 변수 사용 (Herodoc & Nowdoc)
```php
$name = "Alice";

// Heredoc
$text = <<<EOT
안녕하세요, $name 님!
PHP를 배우는 중입니다.
EOT;

echo $text;
```

## 4. 배열 정렬
```php
$numbers = [4, 2, 8, 6];
sort($numbers); // 오름차순 정렬
print_r($numbers);

rsort($numbers); // 내림차순 정렬
print_r($numbers);
```

## 5. JSON 데이터 다루기
```php
$data = ["name" => "Alice", "age" => 25];
$json = json_encode($data); // 배열을 JSON 문자열로 변환
echo $json;

$decoded = json_decode($json, true); // JSON을 배열로 변환
print_r($decoded);
```

## 6. 다차원 배열에서 값 찾기 (`array_column`)
```php
$users = [
    ["id" => 1, "name" => "Alice"],
    ["id" => 2, "name" => "Bob"]
];
$names = array_column($users, "name");
print_r($names); // ["Alice", "Bob"]
```

## 7. 특정 문자열 포함 여부 확인 (`strpos`)
```php
$haystack = "Hello, world!";
$needle = "world";
if (strpos($haystack, $needle) !== false) {
    echo "문자열이 포함되어 있습니다.";
}
```

## 8. 랜덤 문자열 생성
```php
function randomString($length = 8) {
    return substr(str_shuffle("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"), 0, $length);
}

echo randomString(10); // 랜덤한 10자리 문자열 출력
```

## 9. 현재 URL 가져오기
```php
$currentUrl = "http://" . $_SERVER['HTTP_HOST'] . $_SERVER['REQUEST_URI'];
echo $currentUrl;
```

## 10. 간단한 캐싱 구현 (`ob_start`)
```php
ob_start();
echo "이 페이지는 캐싱되었습니다.";
$content = ob_get_contents();
file_put_contents("cache.html", $content);
ob_end_flush();
```

## 11. PHP에서 실행 시간 측정
```php
$start = microtime(true);
// 실행할 코드
sleep(1);
$end = microtime(true);
echo "실행 시간: " . ($end - $start) . "초";
```

## 12. 대소문자 구분 없이 문자열 비교 (`strcasecmp`)
```php
if (strcasecmp("Hello", "hello") == 0) {
    echo "같은 문자열입니다.";
}
```

## 13. URL 리다이렉트
```php
header("Location: https://example.com");
exit();
```

## 14. 배열을 쿼리 스트링으로 변환 (`http_build_query`)
```php
$params = ["name" => "Alice", "age" => 25];
$queryString = http_build_query($params);
echo $queryString; // name=Alice&age=25
```

## 15. 대용량 파일을 줄 단위로 읽기
```php
$file = fopen("large_file.txt", "r");
while (!feof($file)) {
    echo fgets($file) . "<br>";
}
fclose($file);
``
