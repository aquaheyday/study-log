# 🛡️ HSTS (HTTP Strict Transport Security)

**HSTS**는 웹사이트가 오직 **HTTPS로만 접근**되도록 강제하는 보안 메커니즘입니다.  
사용자가 실수로 `http://`로 접속하거나, 공격자가 HTTP 요청을 가로채려 할 때도 **자동으로 HTTPS로 전환**되므로 중간자 공격(MITM)을 방지할 수 있습니다.

---

## 1️⃣ HSTS란?

HTTP 응답 헤더를 통해, 브라우저가 **앞으로 이 사이트는 무조건 `HTTPS`로만 접속해야 한다**고 기억하게 만드는 보안 정책

---

## 2️⃣ 사용 목적

- 사용자가 실수로 `HTT`P로 접속해도 자동으로 `HTTPS`로 리디렉션
- `HTTPS`가 아닌 요청을 아예 차단 (클릭도 안 됨)
- **SSL Stripping** 같은 중간자 공격 방지

---

## 3️⃣ 설정 방법

### 응답 헤더 예시:
```
Strict-Transport-Security: max-age=31536000; includeSubDomains; preload
```
| 디렉티브 | 설명 |
|----------|------|
| `max-age=초` | HTTPS를 유지할 기간 (초 단위, 1년 = 31536000) |
| `includeSubDomains` | 서브도메인도 동일하게 적용 |
| `preload` | 브라우저에 미리 HTTPS 강제 목록으로 등록 요청 (옵션)

---

## 4️⃣ 작동 방식

1. 브라우저가 `HTTPS` 접속 시 `HSTS` 헤더를 응답에서 수신
2. 해당 도메인을 **브라우저가 기억**
3. 이후 사용자가 `http://example.com` 으로 접속 시 → 브라우저가 **자동으로** `https://example.com`**로 전환**

---

## 5️⃣ 보안 효과

| 효과 | 설명 |
|------|------|
| SSL Stripping 방지 | HTTP 요청을 가로채서 HTTPS를 제거하는 공격 차단 |
| 클릭 보안 | HTTPS 연결 안 되는 경우 접속 자체 불가 |
| 서브도메인 보호 | `includeSubDomains` 설정 시 효과 확장 가능 |
| 완전한 HTTPS 환경 구축 | HTTP 노출 가능성 최소화

---

## 6️⃣ preload란?

주요 브라우저(Chrome, Firefox 등)는 `HSTS`를 **미리 내장된 preload list**로 관리합니다.

- [https://hstspreload.org](https://hstspreload.org) 에서 등록 가능
- 조건:
  - `max-age >= 31536000` (1년 이상)
  - `includeSubDomains` 필수
  - `preload` 디렉티브 포함

#### 사용 예시
```
Strict-Transport-Security: max-age=31536000; includeSubDomains; preload
```

---

## 7️⃣ 주의 사항

- 일단 설정하면 **브라우저 캐시에 남기 때문에 롤백 어려움**
- HTTPS가 완벽히 구축되지 않으면 접근 불가가 될 수 있음
- HSTS 적용 전에 모든 리소스가 HTTPS에서 잘 작동하는지 확인 필수

---

## 🎯 정리

✔ HSTS는 브라우저에게 "이 사이트는 HTTPS로만 접속해" 라고 알려주는 장치  
✔ 중간자 공격(SSL Stripping)과 실수로 인한 HTTP 접속을 원천 차단  
✔ `max-age`, `includeSubDomains`, `preload` 세 가지로 구성  
✔ 완전한 HTTPS 전환을 마쳤을 때 적용하는 것이 안전
