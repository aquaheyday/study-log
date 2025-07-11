# 🛡️ Auth Service (Spring Boot + Kotlin + JWT)

이 프로젝트는 마이크로서비스 아키텍처 기반의 사용자 인증 서비스입니다.

## 기술 스택
- Kotlin
- Spring Boot 3.x
- Spring Security + JWT
- PostgreSQL + JPA
- Docker / Docker Compose

## 기능
- ✅ 회원가입 (username, password)
- ✅ 로그인 (JWT 발급)
- ✅ JWT 인증 필터 구현
- ✅ Bcrypt 암호화

## 실행 방법

### 로컬 실행
1. PostgreSQL 도커 실행:
```bash
   docker-compose up -d
```

2. 애플리케이션 실행:
```bash
./gradlew bootRun
```

### API 예시
```http
POST /api/auth/signup
Content-Type: application/json

{
  "username": "testuser",
  "password": "testpass"
} 
```

```http
POST /api/auth/login
Content-Type: application/json

{
  "username": "testuser",
  "password": "testpass"
}
```
응답:
```json
{
  "token": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6..."
}
```


---

### ✅ 3. 올릴 때 주의할 점

- `.gitignore`에 `build/`, `.idea/`, `.gradle/`, `*.iml` 등 추가
- `application.yml`에서 비밀 키는 숨기기 (예: `application-example.yml`로 대체)
- 커밋 메시지 깔끔하게 정리

---

### ✅ 깃허브 업로드 예시 커밋 순서

```bash
git init
git add .
git commit -m "feat: Add JWT-based auth service with signup/login"
git remote add origin https://github.com/yourname/auth-service.git
git push -u origin main
