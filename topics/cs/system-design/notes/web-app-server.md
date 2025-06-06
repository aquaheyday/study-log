# 🌐 웹 서버 & 애플리케이션 서버

**웹 서버(Web Server)** 와 **애플리케이션 서버(Application Server)** 는 웹 기반 시스템에서 **클라이언트 요청을 처리하는 핵심 구성 요소**입니다.

---

## 1️⃣ 웹 서버란?

| 항목       | 설명 |
|------------|------|
| **정의**   | 클라이언트(브라우저)의 HTTP 요청을 처리하고, 정적 파일(HTML, CSS, JS, 이미지 등)을 반환하는 서버 |
| **역할**   | 요청을 받아 정적 리소스를 직접 응답하거나, 동적 처리 요청은 애플리케이션 서버로 전달 |
| **예시**   | Nginx, Apache HTTPD, Caddy |

---

## 2️⃣ 애플리케이션 서버란?

| 항목       | 설명 |
|------------|------|
| **정의**   | 비즈니스 로직(데이터 처리, DB 접근 등)을 수행하고 동적 콘텐츠를 생성하는 서버 |
| **역할**   | 클라이언트의 요청을 분석하고, 데이터베이스와 상호작용하여 처리 결과를 생성 |
| **예시**   | Node.js, Spring Boot, Django, Express, Flask, Ruby on Rails |

---

## 3️⃣ 역할 비교

| 구분               | 웹 서버                             | 애플리케이션 서버                     |
|--------------------|--------------------------------------|----------------------------------------|
| **주요 목적**       | 정적 콘텐츠 제공                    | 동적 콘텐츠 처리 (비즈니스 로직)       |
| **처리 방식**       | HTML, CSS, JS, 이미지 직접 반환     | API 로직 실행, DB 접근, 결과 가공 등   |
| **속도**           | 빠름 (정적 리소스 캐싱 가능)        | 상대적으로 느림 (복잡한 처리 포함)     |
| **요청 처리**       | HTTP 기반 요청 처리 및 프록시 역할  | 내부 로직 실행 후 JSON/HTML 생성       |
| **보통의 구조**     | 클라이언트 → 웹 서버 → 앱 서버 → DB | 클라이언트 → 앱 서버 (웹 서버 생략 가능) |

---

## 4️⃣ 실제 사용 구조 예시

```
[사용자 브라우저]
       ↓ (HTTP 요청)
[웹 서버 - Nginx]
       ↓ (API 요청 프록시)
[애플리케이션 서버 - Node.js, Spring 등]
       ↓
[데이터베이스, 외부 API 등]
```

- 정적 파일: 웹 서버에서 바로 응답 (빠름, 캐싱 가능)
- API 요청: 앱 서버로 전달되어 로직 처리 후 결과 반환

---

## 5️⃣ 같이 사용하는 이유?

| 이유 | 설명 |
|------|------|
| **역할 분리** | 정적/동적 요청을 나눠 처리하여 효율성 ↑ |
| **보안성 향상** | 웹 서버가 외부 노출, 앱 서버는 내부에서만 동작 |
| **성능 최적화** | 정적 리소스는 캐싱, 로드밸런싱 효율 ↑ |
| **유지보수 용이** | 아키텍처 분리로 문제 원인 추적이 쉬움 |

---

## 6️⃣ 단독 사용 가능성?

| 서버 종류       | 단독 사용 가능 여부 | 설명 |
|----------------|--------------------|------|
| **웹 서버만**     | ❌ (동적 로직 불가)   | 정적 파일만 처리, API 등은 불가능 |
| **앱 서버만**     | ⭕ 가능                | 정적 파일도 함께 처리 가능 (속도는 느릴 수 있음) |

---

## 🎯 정리 요약

✔ **웹 서버**: 정적 파일 처리 + 요청 분배 (빠르고 가볍다)  
✔ **애플리케이션 서버**: 비즈니스 로직 처리 + DB 접근 (핵심 로직 담당)  
✔ 둘을 **함께 쓰면 성능, 보안, 구조 분리가 쉬워진다**  
✔ 흔한 조합: **Nginx + Node.js**, **Apache + Django**, **Nginx + Spring Boot**

