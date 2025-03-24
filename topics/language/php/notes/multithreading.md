# ⚙️ PHP 멀티스레딩 및 병렬 처리

PHP는 기본적으로 **요청당 하나의 프로세스(싱글 스레드)** 로 동작합니다.  
하지만 상황에 따라 **병렬 처리나 비동기 처리**가 필요할 수 있으며, 이를 구현하기 위한 다양한 기법이 존재합니다.

---

## 1️⃣ PHP는 기본적으로 싱글 스레드

- PHP는 하나의 요청에서 한 줄씩 코드를 실행합니다.  
- 동시에 여러 작업을 처리하려면 **멀티프로세스, 비동기 처리, 병렬 요청 분산**이 필요합니다.

---

## 2️⃣ `pthreads` (PHP 7 이하)

PHP 7 이하에서만 사용 가능한 **진짜 스레드 기반 확장 모듈**입니다.

> ⚠️ 현재는 PHP 8부터는 지원되지 않으며, 대체로 `parallel` 확장을 사용합니다.

```php
<?php
class Task extends Thread {
    public function run() {
        echo "작업 실행 중...\n";
    }
}

$thread = new Task();
$thread->start();
?>
```

✔ `Thread` 클래스를 상속받아 사용합니다.  
✔ PHP CLI 전용이며, `php.ini`에서 pthreads 확장 설정 필요  

---

## 3️⃣ `parallel` 확장 (PHP 7.2 이상)

`parallel`은 최신 PHP에서 **병렬 실행을 구현**할 수 있는 확장입니다.

#### 1. 설치 (CLI 모드 전용)

```bash
pecl install parallel
```

#### 2. 예제 코드

```php
<?php
use parallel\Runtime;

$runtime = new Runtime();
$future = $runtime->run(function () {
    return "백그라운드 처리 완료";
});

echo $future->value();
?>
```

✔ `Runtime()` → 새로운 스레드 생성  
✔ `run()` → 전달된 함수는 병렬로 실행됨  
✔ `value()` → 결과를 기다리며 반환  

---

## 4️⃣ `curl_multi_exec()` — 병렬 HTTP 요청 처리

여러 URL을 동시에 호출해야 할 때 매우 유용한 방식입니다.

```php
<?php
$urls = [
  "https://example.com/api/1",
  "https://example.com/api/2"
];

$mh = curl_multi_init();
$chs = [];

foreach ($urls as $i => $url) {
    $chs[$i] = curl_init($url);
    curl_setopt($chs[$i], CURLOPT_RETURNTRANSFER, true);
    curl_multi_add_handle($mh, $chs[$i]);
}

do {
    $status = curl_multi_exec($mh, $running);
} while ($running);

foreach ($chs as $ch) {
    $response = curl_multi_getcontent($ch);
    echo $response;
    curl_multi_remove_handle($mh, $ch);
}

curl_multi_close($mh);
?>
```

✔ **비동기 병렬 요청**을 구현할 수 있는 실전적 기법  
✔ API 호출을 동시에 실행할 수 있어 응답 속도 향상에 매우 유리

---

## 5️⃣ `proc_open()` / `shell_exec()` — 외부 프로세스 실행

병렬로 **외부 명령어 또는 PHP 스크립트를 실행**하는 방법입니다.

```php
<?php
$process = proc_open('php child.php', [
    1 => ['pipe', 'w']
], $pipes);

$output = stream_get_contents($pipes[1]);
fclose($pipes[1]);
proc_close($process);

echo $output;
?>
```

✔ `child.php` → 같은 별도 PHP 파일을 **동시 실행**할 수 있습니다.  
✔ `shell_exec("php other.php &")` 방식으로 **백그라운드 실행**도 가능  

---

## 6️⃣ `pcntl_fork()` — 유닉스 계열에서만 가능

PHP에서 **프로세스 분기(복제)** 를 구현할 수 있는 함수입니다.  
리눅스 서버에서만 사용 가능하며, CLI 환경에서 실행해야 합니다.

```php
<?php
$pid = pcntl_fork();

if ($pid == -1) {
    die("fork 실패");
} elseif ($pid) {
    echo "부모 프로세스 실행\n";
} else {
    echo "자식 프로세스 실행\n";
}
?>
```

✔ 부모/자식 프로세스를 나누어 병렬 실행할 수 있음  
⚠️ 고급 시스템 프로그래밍에 적합하며 주의가 필요합니다    

---

## 7️⃣ 병렬 작업 설계 시 주의사항

| 항목               | 설명                                                 |
|--------------------|------------------------------------------------------|
| 공유 자원 접근     | 여러 프로세스가 같은 파일/DB를 동시에 수정하지 않도록 주의 |
| 에러 처리          | 개별 스레드/프로세스의 예외 처리 필수               |
| 실행 환경 제한     | 대부분 CLI 환경에서만 가능 (`parallel`, `pcntl`)    |
| 스레드 안전성      | PHP 내장 함수 대부분은 스레드 세이프하지 않음        |
| 로그/출력 순서     | 병렬 실행 시 출력 순서가 뒤섞일 수 있음               |

---

## 🎯 정리

✔ PHP는 기본적으로 싱글 스레드 기반이지만, 여러 병렬 처리 방법을 제공합니다.  
✔ PHP 7 이하: `pthreads`, PHP 7.2 이상: `parallel` 확장을 통해 스레딩 구현 가능  
✔ 외부 API 호출 시 `curl_multi_exec()`는 매우 실용적인 병렬 처리 방식입니다.  
✔ 리눅스 서버에서는 `pcntl_fork()`로 프로세스 분기가 가능하며 CLI에서만 실행됩니다.  
✔ 병렬 처리 시 공유 자원 충돌 및 예외 처리에 주의하셔야 합니다.

