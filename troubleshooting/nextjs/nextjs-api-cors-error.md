# 🐛 Next.js + Apache 리버스 프록시 환경에서 CORS 오류 발생

## ⚠️ 문제 상황
- 발생 날짜: 2025-04-17  
- 발생 환경 linux, next.js, go, apache, docker
- 재현 방법 
1. 브라우저에서 Next.js fetch()로 Go API 서버 호출
2. 도메인: `http://nextjs.com` → `http://go-api.com`
3. 응답에서 CORS 헤더 누락 → 브라우저에서 CORS 차단됨

```
Access to fetch at 'http://nextjs.com/...' from origin 'http://go-api.com'
has been blocked by CORS policy: No 'Access-Control-Allow-Origin' header...
```
---

## 🔍 원인 분석
- 기존 React에서는 단순 `GET` 요청만 보내므로 브라우저가 **CORS preflight 요청(OPTIONS)** 을 생략 → 문제 없음
- 반면 Next.js에서 `fetch()`를 사용할 때 `credentials: 'include'` 옵션을 함께 쓰면, 브라우저가 **OPTIONS 요청을 먼저 전송 (preflight)**
- 이 OPTIONS 요청을 Apache가 자체적으로 `200 OK`로 응답하면서, **CORS 헤더가 빠진 응답**이 발생 → 브라우저가 차단

```apache
<VirtualHost *:80>
    ServerAdmin webmaster@localhost
    ServerName go-api.com
    ServerAlias go-api.com

    ProxyPreserveHost On
    ProxyPass / http://127.0.0.1:port/
    ProxyPassReverse / http://127.0.0.1:port/
</VirtualHost>
```

---

## 🛠 해결 방법
- Apache에 CORS 헤더를 수동으로 추가
- Next.js + fetch + credentials	브라우저가 OPTIONS 요청 먼저 보냄 → Apache 에서 OPTIONS를 Go API 로 전달 → Go 에서 OPTIONS는 200과 헤더만 리턴해줌
- 기본적으로 OPTIONS는 "preflight 요청" = 브라우저가 본 요청 전에 서버가 CORS를 지원하는지 확인하기 위해 먼저 보내는 예비 요청

### Apache 설정 (추가 내용)
```apache
<VirtualHost *:80>
    ServerAdmin webmaster@localhost
    ServerName go-api.com
    ServerAlias go-api.com

    ProxyPreserveHost On
    ProxyPass / http://127.0.0.1:port/
    ProxyPassReverse / http://127.0.0.1:port/

    <IfModule mod_headers.c>
        Header always set Access-Control-Allow-Origin "*" // 보안상 도메인을 특정하는게 좋음
        Header always set Access-Control-Allow-Credentials "true"
        Header always set Access-Control-Allow-Methods "GET, POST, OPTIONS, PUT, DELETE"
        Header always set Access-Control-Allow-Headers "Content-Type, Authorization"
    </IfModule>
</VirtualHost>

```

---

## 🚀 결과
- 브라우저의 OPTIONS → GET 요청 흐름이 모두 정상적으로 처리됨
- Next.js에서 fetch() + credentials 조합으로도 Go API와 정상 통신 가능
