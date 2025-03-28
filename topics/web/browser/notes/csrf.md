# 🛡️ CSRF (Cross-Site Request Forgery)

사용자가 인증된 상태를 악용해 **원하지 않는 요청을 사용자의 의도 없이 보내게 만드는 공격**입니다.  
대부분 쿠키 기반 인증을 사용하는 시스템에서 발생하며, 사용자의 권한으로 악의적인 요청이 실행됩니다.

---

## 1️⃣ CSRF란?

#### Cross-Site Request Forgery (사이트 간 요청 위조)

- 사용자가 로그인된 상태라는 점을 악용
- 공격자가 **피해자의 브라우저로** 특정 요청을 보내게 유도
- 서버는 **사용자가 보낸 정상 요청**으로 오인해 처리함

---

## 2️⃣ 공격 흐름 예시

1. 사용자가 `bank.com`에 로그인하고 세션이 유지됨 (쿠키 보유 상태)
2. 공격자가 만든 사이트 `evil.com`에 접속
3. 그 페이지 안에 다음과 같은 요청이 숨겨져 있음:
```html
<img src="https://bank.com/transfer?to=attacker&amount=10000">
```
4. 피해자의 브라우저는 이 이미지 요청을 자동으로 실행함
5. 세션 쿠키가 자동으로 포함되어 서버로 전송됨
6. 서버는 인증된 사용자로 인식하고 송금 처리함

---

## 3️⃣ CSRF 특징

| 항목 | 설명 |
|------|------|
| 브라우저 기반 공격 | 요청은 피해자 브라우저에서 전송됨 |
| 자동 인증 쿠키 사용 | 쿠키/세션 기반 인증을 악용 |
| 모든 HTTP 메서드 대상 | GET, POST, PUT 등 모두 가능 |
| 사용자 모르게 실행 | UI 없이도 공격 가능 (img, form, script 등) |

---

## 4️⃣ 공격 예시

#### 1. GET 방식
```html
<img src="https://example.com/delete-account">
```

#### 2. POST 방식
```html
<form action="https://example.com/update-email" method="POST">
  <input type="hidden" name="email" value="attacker@example.com">
</form>

<script>document.forms[0].submit();</script>
```

---

## 5️⃣ 방어 방법

| 방법 | 설명 |
|------|------|
| ✅ CSRF Token 사용 | 요청 시 고유 토큰을 포함, 서버에서 검증 |
| ✅ SameSite 쿠키 설정 | 외부 사이트에서 쿠키 전송 제한 |
| ✅ Referer / Origin 검사 | 요청 출처 검사로 외부 요청 차단 |
| ✅ 사용자 재인증 | 민감한 요청 전에 비밀번호 재확인 등 |
| ✅ CORS 제한 | 신뢰된 출처만 서버에 접근 허용

---

## 6️⃣ CSRF Token 사용 예시

#### 1. 서버가 HTML에 토큰 삽입
```html
<input type="hidden" name="csrf_token" value="abc123xyz">
```

#### 2. 요청 시 토큰 포함
```http
POST /update-password
csrf_token=abc123xyz
```

#### 3. 서버는 세션의 토큰과 비교해 일치 여부 검증

---

## 7️⃣ SameSite 쿠키 예시
```http
Set-Cookie: session=abc123; SameSite=Strict; HttpOnly; Secure
```

| 옵션 | 설명 |
|------|------|
| `Strict` | 외부 사이트 요청 시 쿠키 전송 안 됨 (가장 안전) |
| `Lax` | GET 같은 안전한 요청만 쿠키 포함 |
| `None` | 제약 없음 (`Secure` 옵션과 함께 사용 필수) |

---

## 🎯 요약

✔ CSRF는 **로그인된 사용자의 인증 정보를 악용한 요청 위조 공격**  
✔ 서버는 이 요청을 **정상적인 사용자 요청으로 오해**  
✔ CSRF Token, SameSite 쿠키, Origin 검사 등을 통해 방어 가능  
✔ 민감한 기능에는 **항상 추가 인증 또는 토큰 검사**가 필요
