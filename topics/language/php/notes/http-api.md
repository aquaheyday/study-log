# 🔗 PHP HTTP 요청 및 API 연동

PHP를 사용하면 외부 API와 연동하여 데이터를 주고받을 수 있습니다.  
주로 **REST API 호출**, **웹훅 처리**, **외부 시스템과의 통신** 등에 활용됩니다.

PHP에서 HTTP 요청을 보내는 대표적인 방법은 다음과 같습니다:

#### 1. `file_get_contents()` – 간단한 GET 요청  
#### 2. `cURL` – 유연하고 강력한 HTTP 클라이언트  
#### 3. `Guzzle` – Composer 기반의 고급 HTTP 라이브러리

---

## 1️⃣ file_get_contents() — 간단한 GET 요청
가장 기본적인 방식으로, 외부 URL의 내용을 가져옵니다.

- `file_get_contents()`는 간단한 GET 요청에 적합합니다.  
- `json_decode()`를 사용하면 JSON 응답을 배열로 변환할 수 있습니다.

```php
<?php
$response = file_get_contents("https://jsonplaceholder.typicode.com/posts/1");
$data = json_decode($response, true);

echo $data["title"];
?>
```

✔ POST 요청이나 헤더 설정 등은 어렵기 때문에 단순한 호출에만 추천

---

## 2️⃣ cURL 사용법 (GET/POST 지원)

PHP에서 가장 널리 사용되는 HTTP 요청 방법입니다.  
**GET, POST, PUT, DELETE** 등 다양한 방식과 헤더 설정이 가능합니다.

#### 1. GET 요청 예시

```php
<?php
$ch = curl_init("https://jsonplaceholder.typicode.com/posts/1");
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);

$response = curl_exec($ch);
curl_close($ch);

$data = json_decode($response, true);
echo $data["title"];
?>
```

#### 2. POST 요청 예시 (JSON 데이터 전송)

```php
<?php
$data = ["title" => "Test", "body" => "This is a post", "userId" => 1];

$ch = curl_init("https://jsonplaceholder.typicode.com/posts");
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_POST, true);
curl_setopt($ch, CURLOPT_HTTPHEADER, ['Content-Type: application/json']);
curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($data));

$response = curl_exec($ch);
curl_close($ch);

echo $response;
?>
```

- `CURLOPT_RETURNTRANSFER`: 결과를 문자열로 반환  
- `CURLOPT_POSTFIELDS`: 요청 본문 설정  
- `CURLOPT_HTTPHEADER`: 헤더 설정 가능

---

## 3️⃣ 응답 처리 (JSON 디코딩)

API 서버에서 JSON 형식으로 응답을 받을 경우 `json_decode()` 함수로 파싱합니다.

```php
<?php
$json = '{"id":1,"title":"hello"}';
$data = json_decode($json, true); // true 옵션으로 연관 배열 반환

echo $data["title"]; // hello
?>
```

- 두 번째 인자를 `true`로 설정하면 **연관 배열**, 생략하면 **객체**로 반환됩니다.

---

## 4️⃣ HTTP 상태 코드 확인

```php
<?php
$ch = curl_init("https://jsonplaceholder.typicode.com/posts/1");
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_exec($ch);

$httpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);
curl_close($ch);

echo "상태 코드: $httpCode";
?>
```

✔ REST API 사용 시 응답 상태 코드(200, 404, 500 등)를 확인하여 예외 처리를 할 수 있습니다.

---

## 5️⃣ Guzzle HTTP 클라이언트 (Composer 사용)

`Guzzle`은 PHP에서 가장 많이 쓰이는 **고급 HTTP 클라이언트 라이브러리**입니다.  
복잡한 API 호출, 비동기 처리, 예외 처리에 적합합니다.

#### 1. 설치 (Composer 필요)

```bash
composer require guzzlehttp/guzzle
```

#### 2. 사용 예제

```php
<?php
require 'vendor/autoload.php';

use GuzzleHttp\Client;

$client = new Client();
$response = $client->request('GET', 'https://jsonplaceholder.typicode.com/posts/1');

$body = $response->getBody();
$data = json_decode($body, true);

echo $data["title"];
?>
```

✔ Guzzle은 `request(method, url, options)` 형식으로 다양한 요청을 보낼 수 있으며  
✔ 헤더, 본문, 인증, 타임아웃 등도 옵션으로 간편하게 설정할 수 있습니다.

---

## 6️⃣ API 인증 방식 간단 소개

| 방식             | 설명                                      |
|------------------|-------------------------------------------|
| API Key          | 요청 헤더 또는 URL에 키 포함 (`?key=abc`) |
| Bearer Token     | `Authorization: Bearer {token}` 형식으로 전달 |
| Basic Auth       | `Authorization: Basic base64(id:pw)`      |
| OAuth 2.0        | 토큰 기반 인증 (복잡한 인증 플로우 포함)  |

✔ 대부분의 API 서비스는 인증을 요구하므로, **요청 시 적절한 헤더 설정이 필수**입니다.

---

## 🎯 정리

✔ `file_get_contents()`는 간단한 GET 요청에 적합  
✔ `cURL`은 실무에서 가장 많이 쓰이며, 다양한 HTTP 요청 방식과 유연한 설정이 가능합니다  
✔ JSON 응답은 `json_decode()`로 배열 또는 객체로 변환하여 처리합니다  
✔ `Guzzle`은 PHP용 고급 HTTP 클라이언트로, 복잡한 API 연동에 매우 유용합니다  
✔ 외부 API와 통신할 때는 반드시 **HTTP 상태 코드 확인과 예외 처리**를 함께 하셔야 합니다

