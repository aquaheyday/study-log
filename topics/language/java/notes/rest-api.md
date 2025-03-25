# 🔗 Java - REST API 설계

REST API는 **자원을 URI로 표현**하고, **HTTP 메서드**로 행위를 나타내는 아키텍처 스타일입니다.  
Spring Boot에서는 REST API를 매우 간단하게 구현할 수 있습니다.

---

## 1️⃣ REST란?

- **RE**presentational **S**tate **T**ransfer  
- 클라이언트와 서버 간의 **자원 기반 통신 방식**  
- 자원(Resource)은 URL로 표현, 행위는 HTTP 메서드로 표현

---

## 2️⃣ HTTP 메서드와 의미

| 메서드 | 의미 | 사용 예 |
|--------|------|---------|
| `GET` | 자원 조회 | `/users` (목록), `/users/1` (단건) |
| `POST` | 자원 생성 | `/users` + JSON 데이터 |
| `PUT` | 자원 전체 수정 | `/users/1` |
| `PATCH` | 자원 일부 수정 | `/users/1` |
| `DELETE` | 자원 삭제 | `/users/1` |

---

## 3️⃣ RESTful URL 설계 원칙

- 명사 기반의 자원 이름 사용  
- 동사는 HTTP 메서드로 표현  
- 계층적 구조로 설계

```http
GET     /articles          → 전체 글 목록
GET     /articles/10       → ID가 10인 글 조회
POST    /articles          → 새 글 작성
PUT     /articles/10       → 글 전체 수정
PATCH   /articles/10       → 글 일부 수정
DELETE  /articles/10       → 글 삭제
```

---

## 4️⃣ Spring Boot REST 컨트롤러 예시

```java
@RestController
@RequestMapping("/users")
public class UserController {

    @GetMapping
    public List<User> getAllUsers() {
        return userService.findAll();
    }

    @GetMapping("/{id}")
    public User getUser(@PathVariable Long id) {
        return userService.findById(id);
    }

    @PostMapping
    public User createUser(@RequestBody User user) {
        return userService.save(user);
    }

    @PutMapping("/{id}")
    public User updateUser(@PathVariable Long id, @RequestBody User user) {
        return userService.update(id, user);
    }

    @DeleteMapping("/{id}")
    public void deleteUser(@PathVariable Long id) {
        userService.delete(id);
    }
}
```

---

## 5️⃣ DTO 사용 예시

```java
public class UserRequest {
    public String name;
    public String email;
}
```

```java
@PostMapping
public ResponseEntity<User> create(@RequestBody UserRequest request) {
    User user = userService.create(request);
    return ResponseEntity.ok(user);
}
```

---

## 6️⃣ 응답 코드 & ResponseEntity

```java
return ResponseEntity.ok(user);                 // 200 OK
return ResponseEntity.status(201).body(user);   // 201 Created
return ResponseEntity.notFound().build();       // 404 Not Found
```

---

## 7️⃣ 예외 처리 (@ExceptionHandler)

```java
@RestControllerAdvice
public class GlobalExceptionHandler {

    @ExceptionHandler(UserNotFoundException.class)
    public ResponseEntity<String> handleUserNotFound(UserNotFoundException e) {
        return ResponseEntity.status(404).body(e.getMessage());
    }
}
```

---

## 8️⃣ 유효성 검사 (Validation)

```java
public class UserRequest {
    @NotBlank
    private String name;

    @Email
    private String email;
}
```

```java
@PostMapping
public ResponseEntity<?> create(@Valid @RequestBody UserRequest req) {
    ...
}
```

> `@Valid`와 `@Validated`, `BindingResult`도 함께 사용할 수 있음

---

## 9️⃣ REST API 문서화 - Swagger

- OpenAPI 기반 자동 문서화 도구  
- 의존성 추가 (SpringDoc or Swagger 3)

```groovy
implementation 'org.springdoc:springdoc-openapi-starter-webmvc-ui:2.0.2'
```

✔ 접속: `http://localhost:8080/swagger-ui.html`

---

## 🔟 REST 설계 주의사항

- 자원 중심의 URL을 지킬 것 (`/createUser` ❌ `/users` ✅)
- HTTP 메서드로 동작 표현 (POST, GET 등)
- 응답 시 적절한 HTTP 상태 코드 사용
- DTO로 요청/응답 분리 (Entity 직접 노출 ❌)

---

## 🎯 정리

✔ REST는 자원을 URI로 표현 + 동작은 HTTP 메서드로 구분  
✔ `@RestController`, `@RequestMapping`, `@Get/Post/Put/DeleteMapping`으로 API 구성  
✔ URL은 **명사형**, **계층 구조**, **복수형 사용** 권장  
✔ 요청은 `@RequestBody`, 경로 변수는 `@PathVariable`  
✔ 응답은 `ResponseEntity`로 상태 코드 포함  
✔ 예외는 `@ExceptionHandler` 또는 `@ControllerAdvice`로 처리  
✔ 유효성 검사는 `@Valid`, `@Validated`  
✔ 문서는 Swagger(SpringDoc)로 자동 생성 가능

