# 🌐 JavaScript - DOM 조작

- **DOM(Document Object Model)** 은 웹 페이지의 HTML 요소들을 **자바스크립트로 제어**할 수 있게 만든 객체 구조입니다.
- DOM을 이용해 **HTML 요소를 선택, 변경, 추가, 삭제, 이벤트 등록** 등을 할 수 있습니다.

---

## 1️⃣ DOM 요소 선택

| 메서드 | 설명 |
|--------|------|
| `document.getElementById(id)` | ID로 요소 선택 |
| `document.getElementsByClassName(class)` | 클래스 이름으로 선택 (HTMLCollection) |
| `document.getElementsByTagName(tag)` | 태그 이름으로 선택 (HTMLCollection) |
| `document.querySelector(selector)` | CSS 선택자 사용 (첫 번째 요소) |
| `document.querySelectorAll(selector)` | CSS 선택자로 모두 선택 (NodeList) |

```js
const title = document.getElementById("main-title");
const items = document.querySelectorAll(".item");
```

---

## 2️⃣ DOM 요소 내용 변경

| 속성 | 설명 |
|------|------|
| `textContent` | 순수 텍스트 변경 |
| `innerHTML`   | HTML 태그 포함하여 변경 |

```js
title.textContent = "안녕하세요!";
title.innerHTML = "<span style='color:red;'>Hello</span>";
```

---

## 3️⃣ 속성(attribute) 조작

```js
const link = document.querySelector("a");

link.getAttribute("href");         // 속성값 가져오기
link.setAttribute("href", "https://naver.com"); // 속성값 변경
link.removeAttribute("target");    // 속성 제거
```

---

## 4️⃣ 스타일 조작

- **JS의 스타일 속성은 camelCase로 씀** (`background-color` → `backgroundColor`)
  
```js
const box = document.querySelector(".box");

box.style.color = "blue";
box.style.backgroundColor = "lightgray";
```

---

## 5️⃣ 클래스 조작

```js
box.classList.add("active");
box.classList.remove("hidden");
box.classList.toggle("dark-mode");
box.classList.contains("active"); // true/false
```

---

## 6️⃣ 요소 추가 / 삭제

```js
const list = document.querySelector("ul");

// 새 요소 만들기
const newItem = document.createElement("li");
newItem.textContent = "새 항목";

// 추가
list.appendChild(newItem);

// 삭제
list.removeChild(list.children[0]);
```

---

## 7️⃣ 이벤트 등록

```js
const btn = document.querySelector("button");

btn.addEventListener("click", function() {
  alert("버튼이 클릭됨!");
});
```

| 이벤트 종류 | 설명 |
|-------------|------|
| `"click"`   | 클릭 |
| `"mouseover"` | 마우스 올림 |
| `"keydown"` | 키 누름 |
| `"submit"` | 폼 제출 |

---

## 8️⃣ 인라인 vs addEventListener

#### 인라인 예제
```html
<!-- 인라인 방식 (권장 ❌) -->
<button onclick="alert('클릭!')">Click me</button>
```

#### addEventListener 예제
```js
// 권장 방식 (addEventListener)
btn.addEventListener("click", () => {
  console.log("클릭됨");
});
```

---

## 🎯 정리

| 할 일            | 방법                         |
|------------------|------------------------------|
| 요소 선택         | `getElementById`, `querySelector` 등 |
| 내용 변경         | `textContent`, `innerHTML` |
| 속성 변경         | `getAttribute`, `setAttribute` |
| 스타일 변경       | `element.style.속성`        |
| 클래스 조작       | `classList.add/remove/toggle` |
| 요소 추가/삭제    | `createElement`, `appendChild`, `removeChild` |
| 이벤트 연결       | `addEventListener`          |

