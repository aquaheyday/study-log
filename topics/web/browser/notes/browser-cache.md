# 🗂️ Browser - 캐시 동작

## 1️⃣ 브라우저 캐시란?

- **웹 리소스(HTML, CSS, JS, 이미지 등)** 를 브라우저에 저장해두고, **같은 리소스를 다시 요청할 때 더 빠르게 응답**하는 방식
- 서버의 응답 속도 향상, 트래픽 감소, UX 개선 효과

---

## 2️⃣ 캐시 저장 위치

| 저장소 | 설명 |
|--------|------|
| Memory Cache | 현재 탭에서만 유효, 페이지 리로드 시 유지됨 |
| Disk Cache | 디스크에 저장, 브라우저 재시작 후에도 유지됨 |
| Service Worker Cache (Cache API) | 개발자가 직접 컨트롤, 오프라인 지원 |
| Application Cache (AppCache) | ❌ 사용 중단 (deprecated) |

---

## 3️⃣ 캐시 동작 간략한 흐름

1. 브라우저가 요청 전에 **캐시 저장소 확인**
2. 캐시 정책에 따라:
   - 리소스를 **그대로 사용**
   - 서버에 **조건부 요청**
   - **완전히 새로 요청**
3. 서버 응답 또는 캐시 사용 결과를 렌더링

---

## 4️⃣ 캐시 관련 주요 HTTP 헤더

### 1) `Cache-Control`

```http
Cache-Control: public, max-age=3600
```

| 디렉티브 | 설명 |
|-----------|------|
| `public` | 모든 사용자에게 캐시 가능 |
| `private` | 사용자 전용, 프록시 캐시 금지 |
| `no-cache` | 캐시 저장은 가능하지만 매번 서버 검증 필요 |
| `no-store` | 아예 캐시 금지 |
| `max-age=초` | 지정된 시간 동안 캐시 유지 |

---

### 2) `Expires` (HTTP/1.0 방식)

- 절대 시간을 기준으로 캐시 만료 시점 설정
- `Cache-Control`이 있으면 무시됨

```http
Expires: Wed, 30 Mar 2025 12:00:00 GMT
```

---

### 3) 조건부 요청: `ETag`, `Last-Modified`

- 브라우저가 "캐시에 있는 리소스를 계속 써도 될지" 서버한테 물어보는 요청
- 서버는 "변경 없음"이라고 알려주면 다시 다운로드하지 않고 기존 리소스를 사용함

| 항목 | 설명 |
|------|------|
| ETag | 파일의 고유 ID (내용이 바뀌면 값도 바뀜) |
| Last-Modified | 마지막 수정 시간 |
| If-None-Match |	클라이언트가 갖고 있는 ETag 값 |
| If-Modified-Since |	클라이언트가 갖고 있는 수정 시간 |
| 304 Not Modified |	리소스가 바뀌지 않았음을 의미. 캐시 사용 OK |

#### 1. 브라우저에서 서버에 조건부 요청
```http
If-None-Match: "abc123"       → ETag 기반
If-Modified-Since: Wed, 25 Mar 2025 10:00:00 GMT
```

#### 2. 서버는 해당 리소스의 현재 상태와 비교, 만약 내용이 안 바뀌었으면 아래와 같이 응답함

```http
304 Not Modified
```

 #### 3. 브라우저는 기존에 캐시된 파일을 그대로 사용


---

## 5️⃣ 캐시 전략 비교

| 전략 | 설명 | 예시 |
|------|------|------|
| No Cache | 캐시 안 함 | `Cache-Control: no-store` |
| Reload Validation | 매번 서버에 확인 | `no-cache` + `ETag` |
| Expiration | 일정 시간 캐시 유지 | `Cache-Control: max-age=86400` |
| Versioning | 파일명에 버전 붙여서 갱신 관리 | `/app.v2.js` |

---

## 6️⃣ 실무에서 자주 쓰는 패턴

| 리소스 유형 | 캐시 전략 |
|-------------|-----------|
| HTML 문서 | `no-cache`, 항상 최신 상태 유지 |
| CSS/JS (빌드 파일) | `Cache-Control: public, max-age=31536000`, 파일명에 해시 포함 |
| 이미지/폰트 | 장기 캐싱 (`max-age`, `immutable`) |
| API 응답 | `no-store` or `private`, `max-age=0` (민감도에 따라 조정) |

---

## 7️⃣ Service Worker 캐시 (추가)

- PWA(Progressive Web App)에서 사용되는 **Cache API**
- 네트워크 요청을 가로채서 캐시에서 응답하거나, 동적으로 저장 가능
- 개발자가 직접 캐시 컨트롤 → 매우 유연

```js
caches.open('v1').then((cache) => {
  cache.addAll(['/index.html', '/style.css']);
});
```

## 8️⃣ 캐시 동작 흐름

### 1) 첫 요청

#### 1. 브라우저가 처음 리소스를 요청할 때

```http
GET /main.css
```

#### 2. 서버 응답

```http
200 OK
ETag: "abc123"
Last-Modified: Wed, 25 Mar 2025 10:00:00 GMT
Cache-Control: no-cache
```

✔ 브라우저는 `ETag`와 `Last-Modified` 값을 저장함

----

### 2) 다음 요청 (조건부 요청)

#### 1. 브라우저는 서버에 리소스를 재요청하면서 조건을 함께 보냄

```http
GET /main.css
If-None-Match: "abc123"
If-Modified-Since: Wed, 25 Mar 2025 10:00:00 GMT
```

#### 2. 서버의 판단 (서버는 해당 리소스의 현재 상태와 비교)


##### 만약 내용이 안 바뀌었으면 아래와 같이 응답함

```http
304 Not Modified
```

✔ 브라우저는 기존에 캐시된 파일을 그대로 사용

##### 만약 내용이 바뀌었으면 변경된 파일을 보냄

```http
HTTP/1.1 200 OK
Content-Type: text/css
ETag: "xyz789"
Last-Modified: Thu, 28 Mar 2025 10:00:00 GMT
Cache-Control: no-cache

/* 실제 리소스 본문 */
body {
  background: black;
}
```

---

## 🎯 정리

✔ 브라우저 캐시는 웹 성능 최적화의 핵심  
✔ 캐시 정책은 **헤더 (`Cache-Control`, `ETag`, `Expires`) 조합으로 제어**  
✔ 민감한 데이터는 캐시하지 않도록 설정 (`no-store`, `private`)  
✔ 정적 리소스는 **장기 캐시 + 파일 해싱** 전략이 일반적
