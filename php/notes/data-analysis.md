# 📊 PHP 데이터 분석 정리

PHP는 웹 프로그래밍 언어지만, **파일 데이터 처리**, **간단한 수치 계산**, **외부 API 활용** 등을 통해 **기초적인 데이터 분석 및 리포팅 작업**도 수행할 수 있습니다.

---

## 1️⃣ 데이터 분석에 적합한 PHP 환경
- PHP 7.4 이상  
- Composer 사용 가능  
- CLI 환경에서 실행 (ex: `php analysis.php`)  
- 확장 라이브러리 설치 가능 (예: `league/csv`, `php-spreadsheet` 등)

---

## 2️⃣ CSV 데이터 분석 (가장 많이 쓰임)

### 1) CSV 읽기

```php
<?php
$rows = array_map("str_getcsv", file("data.csv"));
$headers = array_shift($rows); // 첫 줄은 헤더

foreach ($rows as $row) {
    $record = array_combine($headers, $row);
    echo $record["name"] . " - " . $record["score"] . "\n";
}
?>
```

✔ `array_map()`과 `array_combine()`을 사용하면 **CSV → 연관 배열**로 변환할 수 있습니다.  

---

### 2) CSV 분석 예: 평균 점수 계산

```php
<?php
$total = 0;
$count = 0;

foreach ($rows as $row) {
    $record = array_combine($headers, $row);
    $total += (int) $record["score"];
    $count++;
}

$average = $total / $count;
echo "평균 점수: $average";
?>
```

---

## 3️⃣ JSON 데이터 분석

```php
<?php
$json = file_get_contents("data.json");
$data = json_decode($json, true);

foreach ($data["users"] as $user) {
    echo $user["name"] . ": " . $user["age"] . "세\n";
}
?>
```

✔ `json_decode($json, true)` 로 **연관 배열 형태**로 분석 (API 응답 분석에도 자주 사용)  

---

## 4️⃣ 간단한 통계 계산 (평균, 최댓값, 최솟값 등)

```php
<?php
$scores = [75, 82, 91, 64, 88];

$avg = array_sum($scores) / count($scores);
$max = max($scores);
$min = min($scores);

echo "평균: $avg / 최고점: $max / 최저점: $min";
?>
```

---

## 5️⃣ 텍스트 데이터 분석 (단어 수, 빈도수 등)

```php
<?php
$text = file_get_contents("article.txt");
$words = str_word_count(strtolower($text), 1);

$counts = array_count_values($words);
arsort($counts);

foreach (array_slice($counts, 0, 5) as $word => $freq) {
    echo "$word: $freq 회\n";
}
?>
```

✔ `str_word_count()` + `array_count_values()` 조합으로 **단어 빈도 분석** 가능  

---

## 6️⃣ 시각화 (외부 라이브러리 활용)

PHP는 기본적으로 그래프 시각화에 특화되어 있지 않지만, **GD**, **Image_Graph**, **Chart.js(JSON 연동)** 등의 방식으로 시각화 가능합니다.

### Bar 그래프 예시 (`GD`)

```php
// PHP GD로 막대그래프 생성 (간단 예)
$img = imagecreatetruecolor(300, 200);
$white = imagecolorallocate($img, 255, 255, 255);
$blue = imagecolorallocate($img, 0, 0, 255);
imagefill($img, 0, 0, $white);

$values = [100, 80, 60, 40];
$x = 20;

foreach ($values as $v) {
    imagefilledrectangle($img, $x, 200, $x + 30, 200 - $v, $blue);
    $x += 40;
}

header("Content-Type: image/png");
imagepng($img);
imagedestroy($img);
```

✔ 또는 결과 데이터를 JS 라이브러리(Chart.js)에 JSON으로 넘겨 클라이언트에서 시각화 가능

---

## 7️⃣ Excel / Spreadsheet 분석

### `PhpSpreadsheet` (by Composer)

#### `Compoer` 로 `PhpSpreadsheet` 설치
```bash
composer require phpoffice/phpspreadsheet
```

#### 예제
```php
<?php
use PhpOffice\PhpSpreadsheet\IOFactory;

$spreadsheet = IOFactory::load("data.xlsx");
$sheet = $spreadsheet->getActiveSheet();

$data = $sheet->toArray();
print_r($data);
?>
```

✔ 엑셀 데이터를 PHP 배열로 불러와 통계 및 필터링 가능  

---

## 8️⃣ 외부 API 연동 데이터 분석

```php
<?php
$response = file_get_contents("https://api.exchangerate-api.com/v4/latest/USD");
$data = json_decode($response, true);

echo "KRW 환율: " . $data["rates"]["KRW"];
?>
```

✔ PHP로 외부 JSON API를 호출하여 **환율, 날씨, 주식 등 실시간 분석 가능**  

---

## 🎯 정리

✔ PHP로 CSV/JSON/Excel 파일을 읽고 배열로 변환하여 분석할 수 있습니다.  
✔ `array_map`, `array_sum`, `array_count_values`, `json_decode` 등 기본 함수만으로도 충분한 처리가 가능합니다.  
✔ GD, PhpSpreadsheet, 외부 API 등을 활용하면 **시각화와 외부 데이터 활용**도 가능합니다.  
✔ CLI 또는 웹 기반 리포트 자동화에도 적합합니다.

