# 📁 PHP 파일 입출력

PHP는 파일 시스템에 접근하여 **파일을 읽고, 쓰고, 삭제하는 작업**을 할 수 있습니다.  
파일 입출력은 주로 **로그 기록, 설정 파일 관리, 사용자 업로드 처리** 등에 활용됩니다.

---

## 1️⃣ 파일 열기 (`fopen()`)

```php
<?php
$file = fopen("sample.txt", "r");
?>
```

| 모드 | 설명 |
|---|---|
| `r` | 읽기 전용. 파일이 **존재해야 하며**, 없으면 오류 발생 |
| `w` | 쓰기 전용. 파일이 **없으면 생성**, 있으면 **내용 덮어씀** |
| `a` | 추가 전용. 파일이 없으면 생성, 있으면 **끝에 내용 추가** |
| `x` | 새 파일 쓰기 전용. **파일이 이미 존재하면 오류** 발생 |
| `r+` | 읽기 + 쓰기. 파일이 **존재해야 하며**, 덮어쓰지 않음 |
| `w+` | 읽기 + 쓰기. 파일이 없으면 생성, 있으면 **내용 모두 덮어씀** |
| `a+` | 읽기 + 추가. 파일 끝에 내용 추가. 없으면 생성 |

---

## 2️⃣ 파일 읽기

### 1) `fread()` — 지정한 바이트만큼 읽기

```php
<?php
$file = fopen("sample.txt", "r");
$content = fread($file, filesize("sample.txt"));
fclose($file);
echo $content;
?>
```

---

### 2) `file_get_contents()` — 파일 전체 읽기 (간편함)

```php
<?php
$content = file_get_contents("sample.txt");
echo $content;
?>
```

✔ 파일 크기나 구조가 단순할 경우 `file_get_contents()`를 사용하는 것이 더 효율적입니다.  

---

## 3️⃣ 파일 한 줄씩 읽기
- `fgets()`는 줄 단위로 읽습니다.  
- `feof()`는 파일 끝(EOF: End of File)에 도달했는지 확인합니다.

```php
<?php
$file = fopen("sample.txt", "r");

while (!feof($file)) {
    $line = fgets($file);
    echo $line . "<br>";
}

fclose($file);
?>
```

---

## 4️⃣ 파일 쓰기

### 1) `fwrite()` — 문자열을 파일에 작성

```php
<?php
$file = fopen("write.txt", "w");
fwrite($file, "Hello, PHP File!");
fclose($file);
?>
```

---

### 2) `file_put_contents()` — 간단한 쓰기 방법

```php
<?php
file_put_contents("write.txt", "빠르게 쓰기!");
?>
```

---

## 5️⃣ 파일 존재 여부 확인

#### `file_exists()`는 파일 또는 디렉토리 존재 여부를 확인하는 데 사용됩니다.
```php
<?php
if (file_exists("sample.txt")) {
    echo "파일이 존재합니다.";
} else {
    echo "파일이 없습니다.";
}
?>
```

---

## 6️⃣ 파일 삭제 (`unlink()`)

#### `unlink()`는 파일을 삭제하는 함수입니다.  
```php
<?php
if (file_exists("delete.txt")) {
    unlink("delete.txt");
    echo "파일 삭제 완료";
}
?>
```
✔ 삭제 전 `file_exists()`로 존재 여부를 확인하는 것이 좋습니다.  

---

## 7️⃣ 파일 정보 확인
- `filesize()` — 파일 크기(바이트 단위)  
- `filemtime()` — 마지막 수정 시간 (타임스탬프 반환)

```php
<?php
$size = filesize("sample.txt");
$time = filemtime("sample.txt");

echo "파일 크기: $size 바이트<br>";
echo "수정 시각: " . date("Y-m-d H:i:s", $time);
?>
```

---

## 8️⃣ 디렉토리 생성 및 삭제

```php
<?php
mkdir("newFolder");       // 폴더 생성
rmdir("newFolder");       // 폴더 삭제 (비어 있어야 함)
?>
```

✔ 이미 있는 디렉토리를 만들면 오류가 발생하므로 `is_dir()` 또는 `file_exists()`로 사전 체크가 필요합니다.

---

## 9️⃣ 업로드된 파일 저장 (기초 예제)
- `$_FILES["upload"]`는 `<input type="file" name="upload">`로 전송된 파일 정보입니다.  
- `move_uploaded_file()`은 업로드된 임시 파일을 지정한 위치로 이동시킵니다.

```php
<?php
if ($_FILES["upload"]["error"] == 0) {
    move_uploaded_file($_FILES["upload"]["tmp_name"], "uploads/" . $_FILES["upload"]["name"]);
    echo "업로드 성공!";
}
?>
```

---

## 🎯 정리

✔ `fopen()`으로 파일 열고, `fread()`/`fwrite()`로 읽고 씁니다.  
✔ `file_get_contents()` / `file_put_contents()`는 간단한 작업에 매우 유용합니다.  
✔ `file_exists()`로 파일 존재 여부를 확인한 후 `unlink()`로 삭제할 수 있습니다.  
✔ `filesize()`, `filemtime()` 등을 통해 파일 정보를 확인할 수 있습니다.  
✔ 업로드된 파일은 `$_FILES` 배열을 통해 접근하고, 보안상 반드시 유효성 검사를 함께 진행하셔야 합니다.

