# 🖼️ PHP GD와 이미지 처리

PHP는 기본적으로 **GD(Graphics Draw)** 라는 내장 라이브러리를 통해 이미지 생성, 편집, 변환 등의 다양한 기능을 제공합니다.

#### GD는 주로 다음과 같은 용도로 사용됩니다
- 이미지 업로드 후 썸네일 생성  
- 워터마크나 텍스트 삽입  
- 캡차 이미지 생성  
- 포맷 변환 (JPG → PNG 등)

---

## 1️⃣ GD 라이브러리 확인 및 설치

#### 1. 설치 확인

```php
<?php
phpinfo(); // GD Support 항목 확인
?>
```

#### 2. 설치 방법 (Linux)

```bash
sudo apt install php-gd
sudo systemctl restart apache2
```

---

## 2️⃣ 새 이미지 생성 및 출력

```php
<?php
header("Content-Type: image/png");

$image = imagecreatetruecolor(200, 100); // 너비 200, 높이 100
$bgColor = imagecolorallocate($image, 0, 128, 255); // 파란색 배경
imagefill($image, 0, 0, $bgColor);

imagepng($image);
imagedestroy($image);
?>
```

✔ 브라우저에서 이 파일을 열면 **이미지가 출력됩니다.**  

---

## 3️⃣ 기존 이미지 불러오기

```php
<?php
$image = imagecreatefromjpeg("input.jpg"); // JPG 파일 불러오기
```

| 파일 포맷   | 함수                      |
|-------------|---------------------------|
| JPEG        | `imagecreatefromjpeg()`   |
| PNG         | `imagecreatefrompng()`    |
| GIF         | `imagecreatefromgif()`    |

---

## 4️⃣ 텍스트 삽입 (imagefttext)

```php
<?php
$image = imagecreatetruecolor(400, 200);
$bg = imagecolorallocate($image, 255, 255, 255);
$black = imagecolorallocate($image, 0, 0, 0);
imagefill($image, 0, 0, $bg);

$text = "Hello, GD!";
$font = "./fonts/NotoSansKR-Regular.ttf";

imagefttext($image, 20, 0, 50, 100, $black, $font, $text);

header("Content-Type: image/png");
imagepng($image);
imagedestroy($image);
?>
```

✔ `imagefttext()`는 **TrueType 폰트를 사용한 텍스트 출력**을 지원합니다.  
✔ 폰트 파일(`.ttf`)은 직접 지정해야 합니다.  

---

## 5️⃣ 이미지 리사이즈 (썸네일 생성)

```php
<?php
$src = imagecreatefromjpeg("large.jpg");

$thumbWidth = 200;
$thumbHeight = 150;

$thumb = imagecreatetruecolor($thumbWidth, $thumbHeight);
imagecopyresampled($thumb, $src, 0, 0, 0, 0, $thumbWidth, $thumbHeight, imagesx($src), imagesy($src));

imagejpeg($thumb, "thumb.jpg", 90); // 90은 품질(0~100)

imagedestroy($src);
imagedestroy($thumb);
?>
```

✔ 썸네일 생성 시 `imagecopyresampled()` 함수가 더 부드럽고 고품질입니다.  

---

## 6️⃣ 이미지 자르기 (Crop)

```php
<?php
$src = imagecreatefromjpeg("input.jpg");
$crop = imagecrop($src, ['x' => 100, 'y' => 50, 'width' => 200, 'height' => 200]);
if ($crop !== FALSE) {
    imagejpeg($crop, "cropped.jpg");
    imagedestroy($crop);
}
imagedestroy($src);
?>
```

---

## 7️⃣ 워터마크 삽입 (이미지 위에 또 다른 이미지)

```php
<?php
$main = imagecreatefromjpeg("main.jpg");
$logo = imagecreatefrompng("logo.png");

$posX = imagesx($main) - imagesx($logo) - 10;
$posY = imagesy($main) - imagesy($logo) - 10;

imagecopy($main, $logo, $posX, $posY, 0, 0, imagesx($logo), imagesy($logo));
imagejpeg($main, "watermarked.jpg", 90);

imagedestroy($main);
imagedestroy($logo);
?>
```

✔ 위치를 조절하여 **우측 하단 워터마크**를 삽입하는 방식입니다.  

---

## 8️⃣ 이미지 저장 vs 출력

| 함수            | 설명                    |
|-----------------|-------------------------|
| `imagejpeg()`   | JPEG 이미지 저장/출력   |
| `imagepng()`    | PNG 이미지 저장/출력    |
| `imagegif()`    | GIF 이미지 저장/출력    |

```php
imagejpeg($image);              // 브라우저 출력
imagejpeg($image, "result.jpg"); // 파일로 저장
```

---

## 9️⃣ 메모리 해제

```php
imagedestroy($image);
```

✔ GD는 이미지 리소스를 사용하므로 작업 후에는 메모리 해제가 반드시 필요합니다.

---

## 🎯 정리

✔ PHP의 GD 라이브러리는 이미지 생성, 편집, 리사이징, 워터마크 삽입 등 다양한 기능을 제공합니다.  
✔ `imagecreatetruecolor()`, `imagecopyresampled()`, `imagefttext()` 등의 함수를 조합하여 이미지 처리 가능  
✔ 이미지 출력 전에는 반드시 `header("Content-Type: image/포맷")` 선언  
✔ 외부 폰트 사용 시 `imagefttext()`와 `.ttf` 경로 지정 필요  
✔ 작업 후에는 `imagedestroy()`로 메모리 정리 필요

