# ✅ HTML - 폼 유효성 검사(Form Validation)

HTML5에서는 자바스크립트 없이도 기본적인 **폼 유효성 검사**를 지원합니다.  
속성만 추가해주면 **자동으로 입력 형식이나 필수 여부 등을 체크**해줍니다.

---

## 1️⃣ 주요 유효성 검사 속성

| 속성 | 설명 | 예시 |
|------|------|------|
| `required` | 필수 입력 필드로 지정 | `<input required>` |
| `minlength` / `maxlength` | 입력 글자 수 제한 | `<input minlength="3" maxlength="10">` |
| `min` / `max` | 숫자나 날짜의 최소/최대값 제한 | `<input type="number" min="1" max="100">` |
| `pattern` | 정규 표현식으로 입력 형식 제한 | `<input pattern="[0-9]{3}-[0-9]{4}">` |
| `type` | 자동 검증되는 입력 유형 설정 | `email`, `url`, `number` 등 |
| `step` | 숫자 입력 증가 단위 설정 | `<input type="number" step="5">` |

---

## 2️⃣ 기본 예제

```html
<form>
  <label>이름:
    <input type="text" name="name" required>
  </label>

  <label>나이:
    <input type="number" name="age" min="1" max="99" required>
  </label>

  <label>이메일:
    <input type="email" name="email" required>
  </label>

  <label>비밀번호 (8자 이상):
    <input type="password" name="pw" minlength="8" required>
  </label>

  <label>전화번호 (정규식 검사):
    <input type="text" pattern="[0-9]{3}-[0-9]{4}-[0-9]{4}" placeholder="010-1234-5678" required>
  </label>

  <button type="submit">제출</button>
</form>
```

---

## 3️⃣ 검증되는 `type` 예시

| 타입 | 검사 항목 | 예시 |
|------|------------|------|
| `email` | 이메일 형식 (`@` 포함 여부 등) | `user@example.com` |
| `url` | URL 형식 | `https://example.com` |
| `number` | 숫자만 허용 | `123` |
| `date` | 날짜 형식 | `2025-03-27` |

---

## 4️⃣ 사용 시 주의사항

- 브라우저에 따라 유효성 검사 동작이 **조금씩 다를 수 있음**
- `pattern`은 **정규 표현식**이므로 주의해서 작성
- `title` 속성을 함께 사용하면 **오류 메시지 커스터마이징** 가능

```html
<input type="text" pattern="[A-Za-z]+" title="영문만 입력해주세요">
```

---

## 5️⃣ 기본 유효성 검사 비활성화
- `<form novalidate>`를 사용하면 **기본 유효성 검사 비활성화** 가능
- 자바스크립트와 함께 쓰면 더 정교한 **커스텀 유효성 검사**도 가능


```html
<form onsubmit="return validateForm();">
  ...
</form>

<script>
  function validateForm() {
    const email = document.querySelector('input[name="email"]');
    if (!email.value.includes('@')) {
      alert("이메일 형식이 올바르지 않습니다.");
      return false;
    }
    return true;
  }
</script>
```

---

## 🎯 정리

| 기능 | 적용 방법 |
|------|-----------|
| 필수 입력 | `required` |
| 글자 수 제한 | `minlength`, `maxlength` |
| 숫자/날짜 범위 제한 | `min`, `max`, `step` |
| 이메일/숫자 등 자동검사 | `type="email"`, `type="number"` |
| 형식 검증 | `pattern` + 정규표현식 |
| 사용자 안내 메시지 | `placeholder`, `title` 활용 |
