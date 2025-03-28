# 🛡️ 보안 헤더 설정 (Security Headers) 가이드

웹 애플리케이션에서 응답 헤더에 보안 관련 설정을 추가하면,  
브라우저가 해당 사이트를 더 안전하게 처리할 수 있도록 도와줍니다.

---

## 1️⃣ X-Frame-Options

- 클릭재킹(Clickjacking) 방지용  
- iframe 내 삽입 여부를 제어
```
X-Frame-Options: DENY
X-Frame-Options: SAMEORIGIN
```
| 옵션 | 설명 |
|------|------|
| `DENY` | 어떤 사이트에서도 iframe 삽입 불가 |
| `SAMEORIGIN` | 동일 출처에서만 iframe 허용 |

---

## 2️⃣ (CSP) Content-Security-Policy

- XSS 및 리소스 로딩 제한  
- 외부 스크립트, 스타일, 이미지 등의 출처를 지정
```
Content-Security-Policy: default-src 'self'; script-src 'self' https://cdn.example.com
```
- `'self'` → 현재 출처 허용
- `'none'` → 아무 것도 허용하지 않음
- `'unsafe-inline'`, `'unsafe-eval'` → 위험하므로 지양
- `frame-ancestors` → iframe 삽입 제어 (X-Frame-Options 대체 가능)

---

## 3️⃣ Strict-Transport-Security (HSTS)

- HTTP → HTTPS 강제 리디렉션  
- 중간자 공격(MITM) 방지
```
Strict-Transport-Security: max-age=31536000; includeSubDomains; preload
```
| 옵션 | 설명 |
|------|------|
| max-age | HTTPS만 사용할 기간 (초 단위) |
| includeSubDomains | 서브도메인도 적용 |
| preload | 브라우저 HSTS preload 목록에 등록 요청 가능 |

✔ 적용 시 반드시 HTTPS가 정상 적용된 상태여야 함  

---

## 4️⃣ X-Content-Type-Options

- 콘텐츠 타입 스니핑 방지  
- 브라우저가 응답 MIME 타입을 **추측하지 못하도록** 설정
```
X-Content-Type-Options: nosniff
```
- 예: HTML 파일을 JS로 실행하는 위험 차단

---

## 5️⃣ Referrer-Policy

- 링크 클릭, 리다이렉션 시 `Referer` 헤더에 어떤 정보가 포함되는지 제어
```
Referrer-Policy: no-referrer
```
| 옵션 | 설명 |
|------|------|
| no-referrer | Referer 아예 전송 안 함 |
| same-origin | 동일 출처에서만 Referer 전송 |
| strict-origin | HTTPS → HTTPS 시만 전송 |
| origin-when-cross-origin | 같은 출처는 전체 URL, 다른 출처는 origin만 |

---

## 6️⃣ Permissions-Policy (기존 Feature-Policy)

- 브라우저 기능 접근 제어 (카메라, 위치, 마이크 등)
```
Permissions-Policy: geolocation=(), camera=()
```
- 위 설정은 **모든 도메인에서 해당 기능 차단**
- `geolocation=(self)` → 현재 사이트에서만 허용

---

## 7️⃣ Cross-Origin Resource Policy (CORP)

- 타 사이트에서 리소스 요청 시, 응답 허용 여부 지정  
- 보통 민감한 이미지/데이터 보호에 사용
```
Cross-Origin-Resource-Policy: same-origin
```
| 옵션 | 설명 |
|------|------|
| same-origin | 동일 출처만 허용 |
| same-site | 같은 사이트의 다른 도메인도 허용 |
| cross-origin | 모두 허용 (주의 필요) |

---

## 8️⃣ 보안 헤더 통합 예시
```
X-Frame-Options: DENY
X-Content-Type-Options: nosniff
Content-Security-Policy: default-src 'self'
Strict-Transport-Security: max-age=63072000; includeSubDomains; preload
Referrer-Policy: no-referrer
Permissions-Policy: geolocation=(), camera=()
Cross-Origin-Resource-Policy: same-origin
```

---

## 🎯 정리

✔ 보안 헤더는 브라우저에 보안 규칙을 전달하는 장치  
✔ 기본적인 헤더 설정만으로도 **XSS, 클릭재킹, 정보 노출** 등 위험 차단 가능  
✔ 실 서비스에서는 **모든 보안 헤더를 조합하여 사용하는 것이 권장됨**
