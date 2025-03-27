# 🧾 HTML - `<input>` 태그 속성

`<input>` 태그는 사용자로부터 데이터를 입력받기 위한 HTML 요소입니다.  
속성(Attribute)을 통해 다양한 형태와 동작을 설정할 수 있어요.

---

## 1️⃣ 주요 속성 목록

| 속성 | 설명 | 예시 |
|------|------|------|
| `type` | 입력 필드 유형 지정 | `type="text"`, `type="email"` 등 |
| `name` | 서버로 전송될 데이터의 이름 | `name="username"` |
| `id` | 고유 식별자, `<label for="">` 연결 시 사용 | `id="user-id"` |
| `value` | 초기값 설정 | `value="홍길동"` |
| `placeholder` | 입력 전 안내 텍스트 (회색 글씨) | `placeholder="이름을 입력하세요"` |
| `required` | 필수 입력 항목 | `required` |
| `readonly` | 읽기 전용(수정 불가) | `readonly` |
| `disabled` | 비활성화(입력 불가, 전송 제외) | `disabled` |
| `maxlength` | 최대 글자 수 제한 | `maxlength="10"` |
| `min` / `max` | 숫자/날짜 입력의 최소/최대값 | `min="1" max="100"` |
| `step` | 증가 단위 (숫자 입력 시) | `step="5"` |
| `autocomplete` | 자동완성 여부 (`on` / `off`) | `autocomplete="off"` |
| `autofocus` | 페이지 로드시 자동 커서 포커스 | `autofocus` |
| `checked` | 체크 상태 기본 설정 (`checkbox`, `radio`) | `checked` |
| `pattern` | 정규 표현식 검증 | `pattern="[A-Za-z]{3,}"` |
| `multiple` | 여러 파일/이메일 선택 가능 | `multiple` |
| `accept` | 파일 업로드 허용 형식 | `accept="image/png, image/jpeg"` |

---

## 2️⃣ 다양한 속성 활용 예제

```html
<form>
  <label>이름:
    <input type="text" name="username" placeholder="이름을 입력하세요" required>
  </label>

  <label>비밀번호:
    <input type="password" name="pw" minlength="6" required>
  </label>

  <label>나이:
    <input type="number" name="age" min="1" max="100" step="1">
  </label>

  <label>이메일:
    <input type="email" name="email" autocomplete="off">
  </label>

  <label>프로필 사진:
    <input type="file" name="avatar" accept="image/*">
  </label>

  <label>성별:
    <input type="radio" name="gender" value="male" checked> 남
    <input type="radio" name="gender" value="female"> 여
  </label>

  <button type="submit">제출</button>
</form>
```

---

## 🎯 정리

✔ `type`, `name`, `value`는 **폼 전송 시 핵심 속성**  
✔ `placeholder`, `required`, `maxlength` 등은 **입력 UX 향상**  
✔ `readonly`, `disabled`는 **입력 제어**  
✔ `min`, `max`, `step`은 **숫자나 날짜 등 숫자 기반 입력에 사용**  
✔ `pattern`으로 **정규식 기반 유효성 검증**도 가능  
✔ `required`, `type="email"` 등으로 **기본적인 유효성 검사**는 HTML만으로도 가능  
✔ 자바스크립트와 함께 쓰면 **동적 폼 처리 및 커스텀 검증** 가능  
✔ `autocomplete="off"`를 통해 민감한 정보 입력 필드를 보호 가능  
