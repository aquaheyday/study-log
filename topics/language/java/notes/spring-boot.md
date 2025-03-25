# 🌱 Java - Spring Boot 시작하기

**Spring Boot**는 복잡한 설정 없이 빠르게 Spring 애플리케이션을 시작할 수 있도록 도와주는 프레임워크입니다.  
간단한 웹 서비스부터 대규모 백엔드까지 빠르게 구축할 수 있습니다.

---

## 1️⃣ Spring Boot 특징

- 설정 파일 최소화 (Convention over Configuration)  
- 내장 톰캣 제공 → 별도 WAS 설치 필요 없음  
- REST API, 웹 MVC, DB 연동, 보안 등 손쉽게 통합  
- 실무에서도 가장 많이 사용되는 Java 웹 프레임워크

---

## 2️⃣ 프로젝트 생성

#### 방법 1: [Spring Initializr](https://start.spring.io)  
- 프로젝트: Gradle / Maven  
- 언어: Java  
- Dependencies: `Spring Web`, `Spring Boot DevTools`, `Lombok`, `Spring Data JPA`, `H2 Database` 등

#### 방법 2: IntelliJ 에서 `Spring Initializr` 프로젝트 생성

---

## 3️⃣ 프로젝트 구조

```bash
src/
 ┣ main/
 ┃ ┣ java/com/example/demo/
 ┃ ┃ ┗ DemoApplication.java   // 진입점 (main 메서드)
 ┃ ┣ resources/
 ┃ ┃ ┣ application.yml        // 설정 파일
 ┃ ┃ ┗ static/, templates/    // 정적 리소스 & 뷰 템플릿
 ┗ test/                      // 테스트 코드
```

---

## 4️⃣ `DemoApplication.java`

```java
@SpringBootApplication
public class DemoApplication {
    public static void main(String[] args) {
        SpringApplication.run(DemoApplication.class, args);
    }
}
```

✔ `@SpringBootApplication` =  `@Configuration` + `@EnableAutoConfiguration` + `@ComponentScan` (3개 애노테이션의 묶음)

---

## 5️⃣ 간단한 REST API 만들기

```java
@RestController
public class HelloController {

    @GetMapping("/hello")
    public String hello() {
        return "Hello, Spring Boot!";
    }
}
```

✔ 실행 후 `http://localhost:8080/hello`에서 텍스트 응답 확인

---

## 6️⃣ application.yml 설정 예시

```yaml
server:
  port: 8081

spring:
  application:
    name: demo-app
```

✔ `application.properties`도 사용 가능하지만, `yml` 가독성이 더 좋음

---

## 7️⃣ 의존성 관리 (Gradle)

```groovy
plugins {
    id 'org.springframework.boot' version '3.x.x'
    id 'io.spring.dependency-management' version '1.1.0'
    id 'java'
}

dependencies {
    implementation 'org.springframework.boot:spring-boot-starter-web'
    testImplementation 'org.springframework.boot:spring-boot-starter-test'
}
```

---

## 8️⃣ 실행 방법

```bash
./gradlew bootRun
```

또는 IDE에서 `DemoApplication.main()` 직접 실행

---

## 9️⃣ DevTools (핫 리로드)

- 개발 중 변경 사항을 자동으로 반영

#### 핫 리로드 사용을 위해 의존성 추가:
```groovy
developmentOnly 'org.springframework.boot:spring-boot-devtools'
```

---

## 🔟 기본 포트 변경

#### `application.yml` 또는 `application.properties`에서 설정:

```yaml
server:
  port: 8081
```

---

## 🎯 정리

✔ Spring Boot는 **설정 최소화 + 실행 편리성 극대화**  
✔ `@SpringBootApplication` 으로 구성 시작  
✔ 내장 톰캣으로 실행 → `main()`만 실행하면 바로 웹 서버  
✔ REST API는 `@RestController`, `@GetMapping` 등으로 간단하게 구현  
✔ 설정은 `application.yml` 사용 추천  
✔ `Spring Initializr`로 빠르게 프로젝트 시작 가능  
✔ `Gradle` 또는 `Maven`으로 빌드 & 실행  
✔ 실무에서는 JPA, Spring Security, Validation 등과 함께 사용

