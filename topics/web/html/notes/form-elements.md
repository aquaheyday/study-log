# 📝 HTML - 폼(Form) 요소 개요

HTML 폼은 사용자의 **입력 데이터를 수집**하고 서버에 전송하기 위해 사용됩니다.  
예: 회원가입, 로그인, 검색, 설문 등

---

## 1️⃣ 기본 구조

```html
<form action="/submit" method="post">
  <input type="text" name="username">
  <button type="submit">제출</button>
</form>
```

| 태그 | 설명 |
|------|------|
| `<form>` | 폼 전체를 감싸는 태그 |
| `action` | 데이터를 전송할 서버 주소 |
| `method` | 전송 방식: `GET` or `POST` |

---

## 2️⃣ 주요 폼 요소 목록

| 태그 | 용도 | 비고 |
|------|------|------|
| `<input>` | 사용자 입력 필드 | 다양한 `type` 존재 |
| `<textarea>` | 여러 줄 텍스트 입력 | 댓글, 소개글 등 |
| `<select>` | 드롭다운 목록 | 옵션 선택 |
| `<option>` | 드롭다운의 항목 | `<select>` 내부에 사용 |
| `<label>` | 입력 요소에 대한 설명 | 접근성 향상 |
| `<button>` | 클릭 가능한 버튼 | `type` 지정 가능 |
| `<fieldset>` | 관련 입력 요소 그룹화 | 시각적 구분 |
| `<legend>` | 필드셋의 제목 | `<fieldset>` 내부에서 사용 |

---

## 3️⃣ `<input>` 태그의 다양한 타입

| 태그 | 타입 | 설명 |
|------|------|------|
| `<input type="text">`| `text` | 한 줄 텍스트 |
| `<input type="password">` | `password` | 비밀번호 입력 (숨김 처리) |
| `<input type="email">` | `email` | 이메일 형식 검증 |
| `<input type="number">` | `number` | 숫자만 입력 가능 |
| `<input type="checkbox">` | `checkbox` | 체크박스 (다중 선택) |
| `<input type="radio">` | `radio` | 라디오 버튼 (단일 선택) |
| `<input type="file">` | `file` | 파일 업로드 |
| `<input type="date">` | `date` | 날짜 선택 |
| `<input type="submit">` | `submit` | 제출 버튼 |
| `<input type="reset">` | `reset` | 입력 초기화 |

---

## 4️⃣ 회원가입 폼 예제

```html
<form action="/signup" method="post">
  <label for="username">아이디:</label>
  <input type="text" id="username" name="username" required>

  <label for="email">이메일:</label>
  <input type="email" id="email" name="email">

  <label for="pw">비밀번호:</label>
  <input type="password" id="pw" name="password">

  <label>성별:</label>
  <input type="radio" name="gender" value="male"> 남
  <input type="radio" name="gender" value="female"> 여

  <label>관심 분야:</label>
  <input type="checkbox" name="interest" value="html"> HTML
  <input type="checkbox" name="interest" value="css"> CSS

  <button type="submit">가입하기</button>
</form>
```

---

## 🎯 정리

✔ `<form>` 태그는 사용자 데이터를 서버로 전송  
✔ `method` 속성: `GET` (주소 표시), `POST` (보안성↑)  
✔ 다양한 `<input type="...">`으로 사용자 정보 입력받기  
✔ `<label>`과 `for` 속성을 함께 써서 **접근성 향상**  
✔ `<button>` 또는 `<input type="submit">`로 제출  
✔ `required`, `placeholder`, `maxlength`, `min`, `step` 등의 **속성**도 적극 활용  
✔ `name` 속성은 서버에서 데이터를 식별할 때 **필수**  
✔ 자바스크립트와 함께 사용하면 **실시간 유효성 검사, 제출 제어 등** 가능

