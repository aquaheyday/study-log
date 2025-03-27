# 🎯 CSS - 선택자(Selector)

CSS 선택자는 **HTML 요소를 선택해서 스타일을 적용**할 때 사용합니다.  

---

## 1️⃣ 기본 선택자

| 선택자 | 설명 | 예시 |
|--------|------|------|
| `*` | 전체 선택자 (모든 요소 선택) | `* { margin: 0; }` |
| `태그명` | 태그 선택 | `p { color: red; }` |
| `.클래스` | 클래스 선택 | `.box { background: yellow; }` |
| `#아이디` | ID 선택 | `#header { font-size: 20px; }` |

---

## 2️⃣ 계층(구조) 선택자

| 선택자 | 설명 | 예시 |
|--------|------|------|
| `A B` | A 요소 내부의 B (하위 선택자) | `div p {}` |
| `A > B` | A의 **직계 자식** B | `ul > li {}` |
| `A + B` | A 바로 다음에 오는 형제 B | `h1 + p {}` |
| `A ~ B` | A 뒤에 나오는 **모든 형제 B** | `h1 ~ p {}` |

---

## 3️⃣ 속성(attribute) 선택자

| 선택자 | 설명 | 예시 |
|--------|------|------|
| `[속성]` | 해당 속성이 있는 요소 | `[disabled] { opacity: 0.5; }` |
| `[속성=값]` | 특정 값을 가진 속성 | `[type="text"]` |
| `[속성^=값]` | 해당 값으로 **시작하는** 속성 | `[href^="https"]` |
| `[속성$=값]` | 해당 값으로 **끝나는** 속성 | `[src$=".jpg"]` |
| `[속성*=값]` | 해당 값을 **포함하는** 속성 | `[title*="팁"]` |

---

## 4️⃣ 가상 클래스 선택자

마치 클래스처럼 작동하지만, HTML에 직접 안 써도 되는 "가상의 조건"

| 선택자 | 설명 | 예시 |
|--------|------|------|
| `:hover` | 마우스 오버 시 | `a:hover { color: red; }` |
| `:active` | 클릭 중일 때 | `button:active { background: gray; }` |
| `:focus` | 포커스 되었을 때 | `input:focus { border: 2px solid blue; }` |
| `:first-child` | 첫 번째 자식 요소 | `li:first-child` |
| `:last-child` | 마지막 자식 요소 | `li:last-child` |
| `:nth-child(n)` | n번째 자식 요소 | `li:nth-child(2)` |
| `:not(선택자)` | 특정 선택자를 제외 | `div:not(.active)` |

---

## 5️⃣ 가상 요소 선택자

실제 DOM에는 없지만, CSS를 통해 생성된 시각적 요소

| 선택자 | 설명 | 예시 |
|--------|------|------|
| `::before` | 요소 앞에 콘텐츠 삽입 | `p::before { content: "👉"; }` |
| `::after` | 요소 뒤에 콘텐츠 삽입 | `p::after { content: "✔"; }` |
| `::placeholder` | 입력 필드 placeholder 스타일 | `input::placeholder { color: gray; }` |

#### 예시

```css
p::before {
  content: "👉 ";
  color: orange;
}

p::after {
  content: " ✔";
  color: green;
}
```

```html
<p>이 문장은 예시입니다.</p>
```

#### 출력

👉 이 문장은 예시입니다. ✔

---

## 6️⃣ 조합 예시

```css
/* 클래스가 box이면서 active도 가진 경우 */
.box.active {
  border: 1px solid black;
}

/* id가 nav인 요소 아래의 모든 a 태그 */
#nav a {
  text-decoration: none;
}

/* 홀수 번째 리스트 항목 */
li:nth-child(odd) {
  background-color: #f9f9f9;
}
```

---

## 🎯 정리

| 분류 | 설명 |
|------|------|
| **기본 선택자** | 태그명, 클래스(`.`), ID(`#`), 전체(`*`) |
| **계층 선택자** | 요소 간 부모-자식-형제 관계 |
| **속성 선택자** | HTML 속성 기반 선택 |
| **가상 클래스** | 상태 기반 선택 (`:hover`, `:first-child` 등) |
| **가상 요소** | 시각적으로 추가되는 콘텐츠 (`::before`, `::after`) |
