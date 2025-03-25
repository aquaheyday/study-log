# 🌐 Java - 간단한 웹 서버 만들기

Java에서도 별도의 프레임워크 없이 **HTTP 웹 서버를 직접 구현**할 수 있습니다.  
간단한 실습용 API나 테스트 서버를 만들 때 유용합니다.

---

## 1️⃣ 사용할 클래스: `com.sun.net.httpserver.HttpServer`

- JDK에 내장된 경량 HTTP 서버  
- 실무보단 테스트/교육용으로 적합  
- 톰캣 없이 간단하게 HTTP 요청을 처리할 수 있음

✔ 단, `com.sun.*` 패키지는 비공식 API이므로 실무에서는 Spring, Tomcat 권장

---

## 2️⃣ 의존성 없이 기본 웹 서버 만들기

```java
import com.sun.net.httpserver.*;

import java.io.IOException;
import java.io.OutputStream;
import java.net.InetSocketAddress;

public class SimpleServer {
    public static void main(String[] args) throws IOException {
        HttpServer server = HttpServer.create(new InetSocketAddress(8080), 0);

        server.createContext("/", new MyHandler());

        server.setExecutor(null); // 기본 executor 사용
        server.start();

        System.out.println("서버 실행 중: http://localhost:8080/");
    }

    static class MyHandler implements HttpHandler {
        @Override
        public void handle(HttpExchange exchange) throws IOException {
            String response = "Hello, Java Web Server!";
            exchange.sendResponseHeaders(200, response.getBytes().length);

            OutputStream os = exchange.getResponseBody();
            os.write(response.getBytes());
            os.close();
        }
    }
}
```

---

## 3️⃣ 실행 방법

```bash
javac SimpleServer.java
java SimpleServer
```

✔ 브라우저에서 `http://localhost:8080/` 에서서 확인 가능

---

## 4️⃣ 간단한 REST API 만들기

```java
server.createContext("/hello", exchange -> {
    String response = "{\"message\":\"Hello JSON\"}";
    exchange.getResponseHeaders().set("Content-Type", "application/json");
    exchange.sendResponseHeaders(200, response.getBytes().length);

    try (OutputStream os = exchange.getResponseBody()) {
        os.write(response.getBytes());
    }
});
```

---

## 5️⃣ GET, POST 분기 처리

```java
server.createContext("/echo", exchange -> {
    String method = exchange.getRequestMethod();

    if (method.equalsIgnoreCase("GET")) {
        String response = "GET 요청 수신!";
        exchange.sendResponseHeaders(200, response.length());
        exchange.getResponseBody().write(response.getBytes());
        exchange.close();
    } else if (method.equalsIgnoreCase("POST")) {
        String body = new String(exchange.getRequestBody().readAllBytes());
        String response = "POST 데이터: " + body;
        exchange.sendResponseHeaders(200, response.length());
        exchange.getResponseBody().write(response.getBytes());
        exchange.close();
    }
});
```

---

## 6️⃣ 포트 변경 & Executor 설정

```java
HttpServer server = HttpServer.create(new InetSocketAddress(8000), 0);
server.setExecutor(Executors.newFixedThreadPool(4)); // 병렬 처리
```

---

## 7️⃣ 톰캣/스프링과 비교

| 항목 | 내장 HttpServer | 톰캣/스프링 |
|------|------------------|-------------|
| 설치 | 필요 없음 (JDK 내장) | 별도 설정 필요 |
| 코드량 | 매우 간단 | 구조적, 복잡함 |
| 목적 | 학습/테스트용 | 실무/대규모 서비스 |
| 성능/보안 | 제한적 | 안정적, 확장 가능 |

---

## 8️⃣ 주의사항

- `com.sun.net.httpserver`는 JDK 내부 API로 **공식 표준 아님**
- TLS/HTTPS, 필터, 세션 등은 **직접 구현** 필요
- 고부하 환경엔 적합하지 않음 (Spring Boot, Netty 추천)

---

## 🎯 정리

✔ Java에는 `HttpServer` 클래스를 활용한 **간단한 웹 서버 기능** 내장  
✔ 별도 라이브러리 없이도 GET/POST 응답 처리 가능  
✔ API 라우팅, 응답 헤더, JSON 처리도 구현 가능  
✔ 실무보단 **테스트, 로컬 API 서버 용도**로 사용 적합  
✔ 멀티스레드 처리는 `setExecutor()`로 가능  
✔ 실무에서는 **Spring Boot** 또는 **Tomcat** 기반 서버 사용 권장

