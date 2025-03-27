# 🧪 CSS 변수 (Custom Properties)

**CSS 변수**는 값을 재사용할 수 있는 사용자 정의 속성입니다.  
중복되는 색상, 크기, 여백 등을 변수로 선언하면 유지보수가 쉬워지고 테마 설정도 편리해집니다.

---

## 1️⃣ 기본 문법

### 1) 변수 선언

- `--변수이름` 형식으로 선언  
- `:root`는 전역(Global) 범위

```css
:root {
  --main-color: #3498db;
  --padding-base: 16px;
}
```

---

### 2) 변수 사용

-  `var(--변수이름)` 으로 값을 불러올 수 있음

```css
.button {
  background-color: var(--main-color);
  padding: var(--padding-base);
}
```

---

## 2️⃣ 변수의 범위 (스코프)

- 변수는 선언한 선택자 기준으로 **하위 요소에만 적용됨**
- `:root`에 선언하면 **전역 사용 가능**

```css
:root {
  --font-color: black;
}

.card {
  --font-color: white; /* 지역 변수 재정의 */
  color: var(--font-color);
}
```

---

## 3️⃣ 기본값 설정

`var()` 함수에서 두 번째 인자를 사용하면 **기본값(대체값)** 을 지정 가능

```css
color: var(--title-color, #333);
```

✔ `--title-color`가 없으면 `#333`을 대신 사용함  

---

## 4️⃣ 실전 예제

```css
:root {
  --primary: #4CAF50;
  --font-size-base: 16px;
  --spacing: 1rem;
}

body {
  font-size: var(--font-size-base);
  color: var(--text-color, #444);
}

.button {
  background: var(--primary);
  margin-bottom: var(--spacing);
}
```

---

## 5️⃣ 다크 모드 적용 예시

```css
:root {
  --bg-color: white;
  --text-color: black;
}

@media (prefers-color-scheme: dark) {
  :root {
    --bg-color: #111;
    --text-color: #eee;
  }
}

body {
  background-color: var(--bg-color);
  color: var(--text-color);
}
```

---

## 6️⃣ CSS 변수 vs SASS 변수

| 비교 항목 | CSS 변수 | SASS 변수 (`$`) |
|-----------|----------|-----------------|
| 선언 위치 | 런타임 (브라우저에서 계산) | 컴파일 타임 |
| 동적 변경 | 가능 (JS, 미디어쿼리 등에서) | 불가능 |
| 브라우저 지원 | 대부분 현대 브라우저 지원 | 별도 컴파일러 필요 |
| 우선순위 | CSS 상속 규칙 따름 | 상속 개념 없음 |

---

## 🎯 정리

✔ `--변수이름: 값;` → **변수 선언 방법**  
✔ `var(--변수이름)` → **변수 사용 방법**  
✔ 선언된 선택자 범위 내에서만 적용 → **스코프(범위)**  
✔ `var(--x, 기본값)` → **기본값 설정 방법**  
