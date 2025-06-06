# 🛡️ Secure Coding - 입력 검증(Input Validation)과 정제(Input Sanitization)

입력 검증과 정제는 보안 코딩의 핵심 중 하나로, 외부에서 들어오는 데이터에 대한 신뢰를 배제하고, **의도하지 않은 동작이나 공격(XSS, SQL Injection 등)** 을 방지하는 데 필수적인 기법입니다.

---

## 1️⃣ 입력 검증(Input Validation) vs 입력 정제(Input Sanitization)

| 구분           | 설명                                                                 | 예시                                    |
|----------------|----------------------------------------------------------------------|-----------------------------------------|
| **입력 검증**   | 사용자 입력이 **허용된 값인지 확인** (화이트리스트 기반 권장)         | 이메일 형식 체크, 숫자 범위 확인 등     |
| **입력 정제**   | 입력값에 포함된 **악성 코드 제거 또는 무력화**                        | 스크립트 태그 제거, HTML 인코딩 등      |

---

## 2️⃣ 입력 검증 원칙

1. **화이트리스트 기반 검증**
   - 허용된 형식만 통과시키고, 나머지는 거부
   - 예: `^[a-zA-Z0-9_]{4,20}$` (ID)

2. **서버 측 검증은 필수**
   - 클라이언트 검증만으로는 보안 취약
   - 백엔드에서 모든 입력 검증 적용

3. **길이/형식/범위 제한**
   - 예: 나이(1~100), 이름(20자 이하), 날짜 형식 등

4. **필수값 여부 확인**
   - 빈 값, null, undefined 여부 체크

---

## 3️⃣ 입력 정제 주요 방식

| 정제 대상      | 정제 방법 예시                                           |
|----------------|----------------------------------------------------------|
| HTML/스크립트  | `<`, `>` 등 특수문자 HTML 엔티티로 인코딩 (`&lt;`)         |
| SQL 문자열     | Prepared Statement 또는 ORM 사용                          |
| 파일 경로      | 경로 탐색(`../`) 차단, 확장자 필터링                       |
| URL 파라미터   | 인코딩 처리 (`encodeURIComponent`)                        |

---

## 4️⃣ 코드 예시

### 1) 입력 검증 – JavaScript (이메일)

```js
function isValidEmail(email) {
  const regex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  return regex.test(email);
}
```

---

### 2) 입력 정제 – Node.js (HTML 인코딩)

```js
function escapeHtml(unsafe) {
  return unsafe
    .replace(/&/g, "&amp;")
    .replace(/</g, "&lt;")
    .replace(/>/g, "&gt;")
    .replace(/"/g, "&quot;")
    .replace(/'/g, "&#039;");
}
```

---

## 5️⃣ 체크리스트

- [ ] 모든 사용자 입력에 검증 로직 적용
- [ ] 클라이언트 + 서버 양쪽 모두 검증 수행
- [ ] 정규표현식 또는 전용 유효성 검사 함수 사용
- [ ] HTML/URL/SQL에 대한 인코딩 또는 정제 처리
- [ ] 숫자/날짜/ID 등에 허용 범위 설정
- [ ] API, 폼, 쿼리스트링, 헤더 등 모든 입력 경로 점검

---

## 🎯 정리

✔ 입력 검증: 허용된 값만 통과시키는 **입장 허가**  
✔ 입력 정제: 악성 또는 불필요한 입력을 **무력화/제거**  
✔ 두 개념을 **함께 적용**해야 실질적인 보안 효과가 있음
