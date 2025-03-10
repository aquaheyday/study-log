# PHP 파일 입출력

## 1. 파일 열기 (`fopen`)
```php
$file = fopen("example.txt", "r"); // 읽기 모드로 파일 열기
if ($file) {
    echo "파일을 성공적으로 열었습니다.";
    fclose($file);
} else {
    echo "파일을 열 수 없습니다.";
}
```

---

## 2. 파일 읽기 (`fread`, `fgets`, `file_get_contents`)
```php
// 전체 파일 읽기
$file = fopen("example.txt", "r");
if ($file) {
    $content = fread($file, filesize("example.txt"));
    echo $content;
    fclose($file);
}
```

```php
// 한 줄씩 읽기
$file = fopen("example.txt", "r");
if ($file) {
    while (($line = fgets($file)) !== false) {
        echo $line . "<br>";
    }
    fclose($file);
}
```

```php
// `file_get_contents`로 전체 파일 읽기
$content = file_get_contents("example.txt");
echo $content;
```

---

## 3. 파일 쓰기 (`fwrite`, `file_put_contents`)
```php
// 파일에 쓰기
$file = fopen("example.txt", "w");
if ($file) {
    fwrite($file, "Hello, PHP!\n");
    fclose($file);
}
```

```php
// `file_put_contents` 사용
file_put_contents("example.txt", "Hello, PHP!");
```

---

## 4. 파일 추가 쓰기 (`a` 모드)
```php
$file = fopen("example.txt", "a");
if ($file) {
    fwrite($file, "추가된 내용\n");
    fclose($file);
}
```

---

## 5. 파일 삭제 (`unlink`)
```php
if (file_exists("example.txt")) {
    unlink("example.txt");
    echo "파일이 삭제되었습니다.";
} else {
    echo "파일이 존재하지 않습니다.";
}
```

---

## 6. 파일 존재 여부 확인 (`file_exists`)
```php
if (file_exists("example.txt")) {
    echo "파일이 존재합니다.";
} else {
    echo "파일이 없습니다.";
}
```

---

## 7. 파일 크기 확인 (`filesize`)
```php
if (file_exists("example.txt")) {
    echo "파일 크기: " . filesize("example.txt") . " 바이트";
}
```

---

## 8. 파일 복사 및 이동 (`copy`, `rename`)
```php
// 파일 복사
copy("example.txt", "copy_example.txt");

// 파일 이동 (이름 변경)
rename("copy_example.txt", "moved_example.txt");
```

---

## 9. 파일 및 디렉토리 권한 변경 (`chmod`)
```php
chmod("example.txt", 0644); // 읽기/쓰기 권한 변경
```

---

## 10. 디렉토리 생성 및 삭제 (`mkdir`, `rmdir`)
```php
// 디렉토리 생성
mkdir("new_folder");

// 디렉토리 삭제
rmdir("new_folder");
```
