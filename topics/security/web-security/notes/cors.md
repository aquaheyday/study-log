# 🌐 Browser - CORS 작동 방식

## 1️⃣ CORS란?

- **CORS (Cross-Origin Resource Sharing)** 는 웹 브라우저가 **다른 출처(origin)의 리소스에 대한 요청을 허용할 수 있도록 서버가 명시적으로 허용 정책을 정의하는 방식**입니다.
- 기본적으로 브라우저는 **Same-Origin Policy**에 따라 다른 출처로부터의 응답을 **차단**합니다. 하지만 CORS를 통해 **서버가 허용한 경우에만** 응답을 사용할 수 있습니다.

---

## 2️⃣ 기본 작동 흐름

1. 브라우저에서 **다른 출처로 요청**을 보냄 (`fetch`, `XHR` 등)
2. 브라우저는 요청 방식에 따라:
   - 단순 요청(Simple Request)
   - 사전 요청(Preflight Request)을 보낼지 결정
3. 서버는 CORS 관련 헤더를 응답에 포함시킴
4. 브라우저가 **응답에 포함된 헤더를 검사**
   → 조건에 맞으면 응답을 **정상적으로 전달**  
   → 조건 미달이면 **CORS 오류로 차단**

---

## 3️⃣ CORS 관련 주요 HTTP 헤더

| 헤더 | 설명 |
|------|------|
| `Access-Control-Allow-Origin` | 허용된 출처 (`*` 또는 특정 origin) |
| `Access-Control-Allow-Methods` | 허용된 HTTP 메서드 (`GET`, `POST`, 등) |
| `Access-Control-Allow-Headers` | 허용된 커스텀 헤더 |
| `Access-Control-Allow-Credentials` | 쿠키/세션 허용 여부 (`true`) |
| `Access-Control-Expose-Headers` | JS에서 접근 가능한 응답 헤더 명시 |

---

## 4️⃣ Simple Request vs Preflight Request

### 1) Simple Request (단순 요청)

브라우저가 **바로 요청을 보내고**, 서버가 응답 시 CORS 헤더만 확인함

#### 조건

- 메서드가 `GET`, `POST`, `HEAD`
- `Content-Type`이 단순한 값 (`application/x-www-form-urlencoded`, `multipart/form-data`, `text/plain`)
- 커스텀 헤더 없음

```js
fetch("https://api.example.com/data");
```

---

### 2) Preflight Request (사전 요청)

Simple Request 조건을 넘는 요청일 경우, 브라우저는 먼저 `OPTIONS` 메서드로 **"이 요청을 보내도 되는지"** 서버에 물어봄

```http
OPTIONS /data HTTP/1.1
Origin: https://myapp.com
Access-Control-Request-Method: POST
Access-Control-Request-Headers: Content-Type
```

→ 서버가 `200 OK`와 함께 `Access-Control-Allow-*` 헤더로 응답하면  
→ 본 요청이 **실제로 전송됨**

---

## 5️⃣ 브라우저가 차단하는 경우 (CORS 오류 발생)

| 상황 | 설명 |
|------|------|
| 서버 응답에 `Access-Control-Allow-Origin` 없음 | 다른 출처 → 브라우저 차단 |
| 클라이언트가 `credentials: include` 요청했지만 서버에 `Access-Control-Allow-Credentials: true` 없음 | 쿠키 접근 거부 |
| 서버가 `OPTIONS` 요청에 허용된 메서드/헤더를 명시하지 않음 | Preflight 실패 |
| 서버는 허용했지만 브라우저가 검사 결과 불일치 | 응답 차단됨 |

---

## 6️⃣ credentials 옵션 설명

```js
fetch("https://api.com/data", {
  credentials: "include"
});
```

#### 쿠키, 인증 헤더 등을 포함한 요청을 보낼 경우 서버도 반드시 아래와 같이 응답해야 함

```http
Access-Control-Allow-Credentials: true
Access-Control-Allow-Origin: https://your-app.com
```

✔ 이때, `Access-Control-Allow-Origin: *` 은 사용할 수 없음  
(`*` 는 "누구든지 접근 허용" 이라는 뜻, 거기에 쿠키, 세션, 인증 정보를 포함해서 전송하는건 명백한 보안 문제)

---

## 🎯 정리

✔ CORS는 **브라우저의 보안 정책을 우회하려는 게 아님**  
✔ **서버가 명시적으로 허용해야만** 브라우저가 응답을 처리함  
✔ 요청이 **단순(Simple)** 하면 바로 요청,  
✔ **단순 조건을 넘으면 Preflight(OPTIONS)** 확인 절차 발생  
✔ 브라우저는 **응답이 허용 범위 내에 있는지 검사하고, 아니면 무조건 차단**  
