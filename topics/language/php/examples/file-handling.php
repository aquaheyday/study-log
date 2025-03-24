<?php
// 1. 파일 쓰기 (Writing to a File)
$file = "example.txt";
$content = "안녕하세요, PHP 파일 입출력 예제입니다.\n";

// 파일 열기 (쓰기 모드 "w" : 기존 내용 삭제 후 새로 작성)
$fileHandle = fopen($file, "w");

if ($fileHandle) {
    fwrite($fileHandle, $content); // 파일에 내용 쓰기
    fclose($fileHandle); // 파일 닫기
    echo "파일이 성공적으로 작성되었습니다.\n";
} else {
    echo "파일을 열 수 없습니다.\n";
}

// 2. 파일 읽기 (Reading from a File)
$fileHandle = fopen($file, "r");

if ($fileHandle) {
    echo "파일 내용:\n";
    while (!feof($fileHandle)) { // 파일 끝(EOF)까지 읽기
        echo fgets($fileHandle); // 한 줄씩 읽기
    }
    fclose($fileHandle);
} else {
    echo "파일을 열 수 없습니다.\n";
}

// 3. 파일 존재 여부 확인
if (file_exists($file)) {
    echo "파일이 존재합니다.\n";
} else {
    echo "파일이 존재하지 않습니다.\n";
}

// 4. 파일 추가 쓰기 (Appending to a File)
$fileHandle = fopen($file, "a");

if ($fileHandle) {
    fwrite($fileHandle, "이 줄은 추가된 내용입니다.\n");
    fclose($fileHandle);
    echo "파일에 추가 내용이 저장되었습니다.\n";
} else {
    echo "파일을 열 수 없습니다.\n";
}

// 5. 파일 전체 읽기 (file_get_contents)
$fileContents = file_get_contents($file);
echo "파일 전체 내용:\n$fileContents";

// 6. 파일 삭제 (Deleting a File)
if (file_exists($file)) {
    unlink($file);
    echo "파일이 삭제되었습니다.\n";
} else {
    echo "삭제할 파일이 없습니다.\n";
}
?>
