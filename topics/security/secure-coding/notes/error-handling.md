# 🛡️ Secure Coding - 에러 처리 보안 (Secure Error Handling)

**에러 처리 보안**은 애플리케이션에서 발생하는 예외나 오류를 **공격자에게 유용한 정보를 노출하지 않고**,  
사용자에게는 **안전하고 필요한 수준의 메시지만 전달**하는 기법입니다.

---

## 1️⃣ 왜 에러 처리를 신경 써야 할까?

| 잘못된 에러 처리 | 보안 위험 |
|------------------|-----------|
| 스택 트레이스 노출 | 내부 구조/기술 스택 유출 (경로, 파일명, DB 쿼리 등) |
| 에러 코드 그대로 노출 | SQL 에러 → 쿼리 추측 가능, 인증 로직 유추 가능 |
| 디버그 모드 활성화 | 전체 로그 및 변수 상태 노출 가능성 |
| 상세 오류 메시지 클라이언트에 출력 | 공격자가 시스템을 역추적 가능 |

---

## 2️⃣ 에러 처리 보안 원칙

1. **사용자에게는 최소한의 정보만 제공**
   - 예: “알 수 없는 오류가 발생했습니다”  
   - X: “SQL syntax error at line 1: SELECT ...”

2. **내부 로깅은 상세하게**
   - 에러 코드, 요청 정보, 사용자 정보 등은 서버에만 저장

3. **예외는 반드시 catch**
   - 예외가 노출되지 않도록 `try-catch`, `error boundary`, `middleware` 활용

4. **디버그/개발 모드 비활성화**
   - `debug=false`, `NODE_ENV=production` 등의 설정 필수

5. **HTTP 상태 코드 적절히 사용**
   - `401`, `403`, `500`, `400` 등을 상황에 맞게 구분

---

## 3️⃣ 언어별 코드 예시

### 1) Node.js (Express)

```js
// 오류 핸들러 미들웨어
app.use((err, req, res, next) => {
  console.error(err.stack); // 내부 로그
  res.status(500).json({ message: '서버에 오류가 발생했습니다.' }); // 사용자 메시지
});
```

---

### 2) Python (Flask)

```python
@app.errorhandler(500)
def server_error(e):
    app.logger.error(f'Server Error: {e}, Path: {request.path}')
    return "서버 오류가 발생했습니다.", 500
```

---

### 3) Java (Spring)

```java
@ExceptionHandler(Exception.class)
public ResponseEntity<String> handleException(Exception e) {
    log.error("Internal error", e); // 내부 로깅
    return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR)
                         .body("서버 오류가 발생했습니다.");
}
```

---

## 4️⃣ 체크리스트

- [ ] 사용자에게 에러 상세 정보 노출 방지
- [ ] 내부 로그에는 상세한 스택 트레이스/요청 정보 기록
- [ ] 모든 예외를 처리할 전역 핸들러 구성
- [ ] 운영 환경에서 디버그 설정 비활성화
- [ ] 에러 발생 시 의도한 HTTP 상태 코드 반환
- [ ] 에러 코드에 민감한 정보 포함되지 않도록 주의

---

## 🎯 정리

✔ 에러는 발생할 수 있다. 하지만 **노출되면 공격자의 단서가 된다.**  
✔ 사용자에게는 **최소 정보**, 내부에는 **정확한 로그**  
✔ 운영 환경에서는 항상 **디버그 모드 OFF**  
