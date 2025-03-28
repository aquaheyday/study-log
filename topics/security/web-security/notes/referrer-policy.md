# 🔐 Referrer-Policy

Referrer-Policy 는 브라우저가 외부 사이트로 요청을 보낼 때, `Referer` 헤더에 포함되는 **출처(URL)** 정보를 **얼마나, 어떻게 보낼지** 결정하는 보안 정책입니다.

---

## 1️⃣ `Referer`란?

사용자가 A → B 페이지로 이동할 때, 브라우저가 B 에게 **"이 사용자는 A에서 왔어요"** 라고 알려주는 HTTP 헤더

```
Referer: https://example.com/page1
```

---

## 2️⃣ 왜 문제가 되나?

- 전체 URL 경로가 노출될 수 있음
  - 예: 쿼리 파라미터에 민감 정보가 있는 경우  
    `/payment?card=1234-5678` 같은 경로 노출 ❌
- 외부 도메인에게 **내 서비스 구조나 내부 주소**가 노출될 수 있음
- **광고 차단 우회**나 **트래픽 추적**에 악용될 수 있음

---

## 3️⃣ Referrer-Policy 설정 방법

#### 1. HTTP 응답 헤더로 설정
```
Referrer-Policy: no-referrer
```

#### 2. HTML 메타태그로 설정
```html
<meta name="referrer" content="no-referrer">
```

---

## 4️⃣ 정책 옵션별 동작 차이

| 옵션 | 설명 |
|------|------|
| `no-referrer` | Referer 아예 보내지 않음 |
| `no-referrer-when-downgrade` | (기본값) HTTPS → HTTP 이동 시에는 보내지 않음 |
| `origin` | 도메인만 전송 (`https://example.com`) |
| `origin-when-cross-origin` | 같은 출처면 전체 URL, 다르면 origin만 |
| `strict-origin` | HTTPS → HTTPS 시 origin만, HTTPS → HTTP 시 안 보냄 |
| `same-origin` | 같은 출처일 때만 전체 URL 전송 |
| `strict-origin-when-cross-origin` | 최신 브라우저 기본값, 상황에 따라 URL/Origin/무전송 결정 |
| `unsafe-url` | 항상 전체 URL 전송 (❌ 비추천) |

---

## 5️⃣ 사용 예시

#### 1. 개인정보 노출 방지
로그인 정보, 결제 정보 등 URL에 포함된 민감 데이터를 외부에 노출하지 않음
```
Referrer-Policy: no-referrer
```

#### 2. 트래픽 추적 제한
광고, 외부 링크 추적 방지
```
Referrer-Policy: origin
``` 

#### 3. 안전한 기본값 추천
같은 도메인엔 전체 URL, 외부 도메인에는 origin만 전달 (보안 + 유용성에 좋음)
```
Referrer-Policy: strict-origin-when-cross-origin
```

---

## 6️⃣ 브라우저 지원

- 대부분의 최신 브라우저는 모든 정책 지원
- `strict-origin-when-cross-origin`은 최신 크롬/파이어폭스에서 기본값으로 채택됨

---

## 🎯 정리

✔ Referrer-Policy는 외부 요청 시 내 사이트 정보가 **어떻게 노출되는지 제어**  
✔ 민감한 정보가 URL에 있다면 `no-referrer` 또는 `origin`이 안전  
✔ `strict-origin-when-cross-origin`은 가장 권장되는 균형 잡힌 정책  
✔ 메타 태그나 응답 헤더로 간단히 설정 가능
