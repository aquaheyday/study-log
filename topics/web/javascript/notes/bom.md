# 🌐 JavaScript - 브라우저 BOM (Browser Object Model)

- **BOM (Browser Object Model)** 은 웹 브라우저와 관련된 객체들의 집합으로, **브라우저 창을 제어하거나 정보에 접근할 수 있게 해주는 모델**입니다.
- DOM이 웹 페이지의 요소를 제어하는 것이라면, BOM은 **브라우저 자체를 제어**합니다.

---

## 1️⃣ window 객체

- `window`는 **브라우저 창 자체를 나타내는 최상위 객체**로, BOM의 핵심입니다.  
- 모든 전역 객체와 함수는 `window`의 프로퍼티로 포함됩니다.

```js
console.log(window.innerWidth);  // 창의 너비
console.log(window.innerHeight); // 창의 높이
```

✔ `window` 객체는 생략 가능: `window.alert()` → `alert()`  

---

## 2️⃣ `alert`, `confirm`, `prompt`

사용자와 **간단한 상호작용**을 제공하는 메서드들입니다.

```js
alert("안녕하세요!");  // 단순 메시지 출력

const result = confirm("정말 삭제하시겠습니까?");
console.log(result);  // true 또는 false

const name = prompt("당신의 이름은?");
console.log(name);  // 사용자가 입력한 값
```

✔ confirm: 확인(true) / 취소(false)  
✔ prompt: 사용자 입력 받기  

---

## 3️⃣ `setTimeout` / `setInterval`

비동기적으로 코드를 실행할 수 있게 하는 타이머 함수입니다.

```js
setTimeout(function() {
  console.log("3초 후 실행됩니다.");
}, 3000);

const intervalId = setInterval(function() {
  console.log("1초마다 실행됩니다.");
}, 1000);

// 멈추기
clearInterval(intervalId);
```

✔ `setTimeout`: 한 번만 실행  
✔ `setInterval`: 주기적으로 반복 실행

---

## 4️⃣ `location` 객체

현재 문서의 URL 정보를 담고 있으며, **페이지 이동이나 새로고침 등에 사용**됩니다.

```js
console.log(location.href);  // 현재 URL 전체
location.href = "https://www.google.com";  // 페이지 이동

console.log(location.hostname); // 호스트 이름
console.log(location.pathname); // 경로
console.log(location.search);   // 쿼리 문자열
```

✔ `location.reload()` → 페이지 새로고침  
✔ `location.assign(url)` → 다른 페이지로 이동  

---

## 5️⃣ `navigator` 객체

브라우저와 운영체제에 대한 정보를 제공합니다.

```js
console.log(navigator.userAgent); // 브라우저 정보
console.log(navigator.language);  // 언어 설정
console.log(navigator.onLine);    // 인터넷 연결 여부
```

✔ 보안상 제약이 많아, 일부 정보는 정확하지 않을 수 있음  

---

## 6️⃣ `screen` 객체

사용자의 화면(모니터)에 대한 정보를 제공합니다.

```js
console.log(screen.width);   // 전체 화면 너비
console.log(screen.height);  // 전체 화면 높이
console.log(screen.availWidth);  // 실제 사용 가능한 너비
```

✔ 다양한 디바이스 환경에 따라 반응형 UI 설계 시 참고 가능  

---

## 7️⃣ DOM vs BOM 차이

| 항목 | DOM (Document Object Model) | BOM (Browser Object Model) |
|---|---|---|
| 의미 |	HTML 문서를 객체로 표현한 모델 | 브라우저 자체를 객체로 표현한 모델 |
| 목적 | 웹 페이지 내용 조작 (요소 생성, 변경 등) | 브라우저 기능 제어 (주소, 팝업, 타이머 등) |
| 포함 객체 | `document`, `element`, `text`, `attribute` 등 | `window`, `location`, `navigator`, `screen` 등 |
| 주요 기능	| HTML 구조 변경, 요소 스타일/클래스 조작, 이벤트 등록 등 |	페이지 이동, 브라우저 정보 접근, 타이머 등 |
| 관계 | `document` 객체는 `window` 객체의 프로퍼티임 |	BOM은 `window` 중심으로 작동, DOM 포함함 |
| 예시 | `document.querySelector("h1")` | `window.alert("Hi!"), location.href` |

---

## 🎯 정리

✔ **window 객체**: 모든 BOM 객체의 최상위, 생략 가능  
✔ **alert/confirm/prompt**: 사용자와 간단한 대화창  
✔ **setTimeout/setInterval**: 비동기 타이머 함수  
✔ **location**: URL 정보와 페이지 이동 관련  
✔ **navigator**: 브라우저 및 OS 정보  
✔ **screen**: 디바이스 화면 정보
