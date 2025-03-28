# 🔒 Same-Origin Policy (동일 출처 정책)

**Same-Origin Policy(SOP)** 는 브라우저가 보안을 위해 **다른 출처(origin)에서 로드된 문서나 스크립트 간의 접근을 제한**하는 정책입니다.

✔️ 즉, **출처가 다른 웹 페이지끼리는 서로의 데이터에 접근할 수 없음**  

---

## 1️⃣ 출처(Origin)란?

**출처(Origin)** 는 아래 3가지 요소로 구성됩니다

```text
Origin = 프로토콜 + 호스트(도메인) + 포트번호
```

| 요소 | 예시 |
|------|------|
| 프로토콜 | `http`, `https` |
| 호스트 | `example.com`, `localhost` |
| 포트 | `80`, `443`, `3000` 등 |

✔ 세 요소 중 하나라도 다르면 → "다른 출처"  

---

## 2️⃣ Same-Origin 판단 예시

| URL | 같은 출처인가? | 이유 |
|-----|----------------|------|
| `https://example.com/page1`, `https://example.com/page2` | ✅ 같음 | 경로는 출처에 영향 없음 |
| `http://example.com` , `https://example.com` | ❌ 다름 | 프로토콜이 다름 |
| `https://example.com`, `https://api.example.com` | ❌ 다름 | 서브도메인이 다름 |
| `https://example.com:443`, `https://example.com:8443` | ❌ 다름 | 포트가 다름 |

---

## 3️⃣ Same-Origin Policy 왜 필요할까?

- **악의적인 사이트(X)가 사용자의 정보가 있는 사이트(Y)의 정보를 몰래 가져오는 것을 방지**
- 예: 로그인된 사이트의 쿠키, 로컬스토리지, DOM 등에 무단 접근 차단

---

## 4️⃣ Same-Origin Policy 의 제한 대상

| 대상 | 설명 |
|------|------|
| DOM 접근 | 다른 출처의 iframe/document에 접근 불가 |
| 쿠키/로컬스토리지 | 다른 출처의 쿠키/스토리지 접근 불가 |
| AJAX 요청 응답 | JS로 다른 출처에 요청은 가능하지만, 응답 데이터 접근은 제한됨 |

```js
// 다른 출처에 요청은 되지만, 브라우저가 응답을 차단함
fetch("https://other-site.com/data")
  .then(res => res.json()) // ❌ SOP 위반 에러 발생
```

---

## 5️⃣ 교차 출처(Cross-Origin)란?

현재 웹 페이지의 출처와 다른 출처의 리소스에 접근하는 것

---

## 6️⃣ 예외 처리: CORS (Cross-Origin Resource Sharing)

서버에서 **명시적으로 다른 출처를 허용**하는 HTTP 헤더를 보내면 브라우저가 교차 출처 요청을 허용함

#### Access-Control-Allow-Origin 헤더 예시

```http
Access-Control-Allow-Origin: https://example.com
```

✔ CORS를 통해 **제한된 리소스 공유 가능**  
✔ 단, 서버와 클라이언트가 모두 **명시적 설정** 필요 (서버와 클라이언트(브라우저) 가 각자 명시적으로 CORS 관련 설정을 해야함)

---

## 7️⃣ 기타 예외 / 우회 방식

| 방식 | 설명 |
|------|------|
| `document.domain` 설정 | 동일 서브도메인 간 제한 완화 (거의 사용 안 함) |
| 프록시 서버 | 백엔드에서 요청 중계 |
| JSONP | `<script>` 태그로 우회 (구형 방식) |
| iframe + postMessage | 서로 통신할 수 있는 안전한 방식 제공 |

---

## 🎯 정리

✔ Same-Origin Policy는 **브라우저 보안의 핵심 정책**  
✔ **다른 출처(origin)** 의 리소스는 기본적으로 접근 불가  
✔ 출처 = `프로토콜 + 도메인 + 포트`  
✔ 교차 출처 = 현재 웹 페이지의 출처와 다른 출처의 리소스에 접근하는 것  
✔ 예외적으로 CORS 등으로 제한 완화 가능
