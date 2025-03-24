# 🚀 PHP 캐싱과 성능 최적화

PHP 애플리케이션의 **응답 속도를 향상**시키고 **서버 부하를 줄이기 위한 대표적인 방법**은  
- 캐싱(Caching)  
- 코드 최적화  
- DB 쿼리 최적화  
- OPcache 활용

등이 있습니다.

---

## 1️⃣ OPcache 활성화 (코드 실행 캐시)

`OPcache`는 PHP 코드를 **미리 컴파일한 결과를 메모리에 저장**해서 다음 요청부터는 컴파일 없이 빠르게 실행되도록 도와줍니다.

#### 1. `php.ini` 설정 예시

```ini
opcache.enable=1
opcache.memory_consumption=128
opcache.interned_strings_buffer=8
opcache.max_accelerated_files=10000
opcache.validate_timestamps=0
```

✔ `validate_timestamps=0`: 성능 향상, 단 코드 수정 즉시 반영 안 됨 (개발환경에서는 1)  
✔ `memory_consumption`: 캐시 메모리 크기(MB)  

#### 2.  상태 확인

```php
<?php
phpinfo(); // opcache 설정 확인 가능
?>
```

---

## 2️⃣ 데이터 캐싱 (DB 쿼리 결과 저장)

#### 파일 기반 캐싱 예시

```php
<?php
$cacheFile = "cache/products.json";

if (file_exists($cacheFile) && time() - filemtime($cacheFile) < 300) {
    $data = json_decode(file_get_contents($cacheFile), true);
} else {
    $data = getProductsFromDB(); // DB에서 가져옴
    file_put_contents($cacheFile, json_encode($data));
}
?>
```

✔ **파일 캐시를 통해 5분간 DB 조회를 생략**할 수 있습니다.

---

## 3️⃣ 메모리 기반 캐싱 (Redis / Memcached)

#### Redis 사용 예시 (Predis 라이브러리 필요)

```bash
composer require predis/predis
```

```php
<?php
$redis = new Predis\Client();

if ($redis->exists("user:1")) {
    $user = json_decode($redis->get("user:1"), true);
} else {
    $user = getUserFromDB(1);
    $redis->setex("user:1", 600, json_encode($user)); // 10분 캐싱
}
?>
```

✔ 빠른 조회가 필요한 사용자 정보, 설정 값 등에 적합  
✔ 대규모 서비스에서는 **DB보다 Redis 접근이 훨씬 빠릅니다**  

---

## 4️⃣ HTTP 캐싱 (브라우저/프록시 캐시)

브라우저나 CDN에 **정적 파일 또는 API 응답을 캐싱**할 수 있습니다.

```php
<?php
header("Cache-Control: max-age=3600"); // 1시간 동안 캐시
header("Expires: " . gmdate("D, d M Y H:i:s", time() + 3600) . " GMT");
?>
```

✔ 이미지, CSS, JS 또는 API 응답 등에 적용 가능  
✔ `etag`, `last-modified` 헤더도 함께 사용하면 더 정밀한 캐싱이 가능  

---

## 5️⃣ 쿼리 최적화 (INDEX / LIMIT 등)

```sql
SELECT * FROM users WHERE email = 'user@example.com';
```

#### 최적화 팁

- WHERE 절 컬럼에 **INDEX** 생성
- 필요 없는 `SELECT *` 지양 → 필요한 컬럼만 선택
- `LIMIT`, `OFFSET` 사용으로 페이징 처리
- 반복 쿼리 줄이고 JOIN 최소화

---

## 6️⃣ 코드 최적화 팁

| 항목                    | 설명 |
|-------------------------|------|
| 함수 호출 최소화        | 루프 안에서 함수 반복 호출 피하기 |
| 불필요한 include 제거   | 필요할 때만 파일 로드 (`require_once`) |
| 반복문 안에서 DB 쿼리 금지| 루프 밖에서 미리 쿼리 수행 후 결과 활용 |

---

## 7️⃣ 출력 버퍼링 (Output Buffering)

페이지 전체를 버퍼에 저장한 후 한번에 전송하면 성능이 향상될 수 있습니다.

```php
<?php
ob_start();
// ... 페이지 출력 ...
$content = ob_get_clean();
echo $content;
?>
```

✔ 기본적으로 PHP는 echo, print 등을 사용하면 즉시 브라우저로 데이터를 전송합니다. 하지만 ob_start()로 출력 버퍼링을 시작하면 출력 내용이 메모리 상 버퍼에 잠시 저장되었다가 한 번에 클라이언트로 전송됩니다.  
✔ 특히 **템플릿 처리나 PDF 생성 시** 유용합니다.

---

## 8️⃣ 지연 로딩 / 비동기 처리 활용

- Ajax를 사용해 **페이지 로딩 후 데이터 요청 분리**
- 큰 이미지/데이터는 **필요할 때 로딩** (지연 로딩, Lazy Loading)
- 백그라운드 작업은 **Queue 또는 Cron으로 분리 처리**

---

## 🎯 정리

✔ `OPcache`로 PHP 코드 실행 속도 향상  
✔ 데이터는 파일, Redis, Memcached 등을 사용해 캐싱  
✔ DB 쿼리는 인덱스, 페이징, 최소 컬럼 조회로 최적화  
✔ 브라우저 캐시는 HTTP 헤더로 제어 가능 (`Cache-Control`, `Expires`)  
✔ 반복 쿼리나 불필요한 코드 제거로 전체 처리 속도 개선 가능  
✔ 필요 시 **백그라운드 처리나 JS 비동기 요청**으로 사용자 체감 속도 개선

