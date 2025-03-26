# 🎯 JavaScript - 이벤트 처리

- **이벤트(Event)** 는 웹 페이지에서 발생하는 **사건**으로, 사용자의 행동(클릭, 키 입력 등)이나 브라우저의 변화(페이지 로드 등)에 의해 발생합니다.
- **이벤트 처리(Event Handling)** 는 이러한 이벤트에 **반응하여 특정 코드를 실행**하는 것을 의미합니다.

---

## 1️⃣ 이벤트 모델

자바스크립트에서 이벤트 처리는 **3가지 단계**로 이루어집니다

1. **캡처링 단계 (Capturing Phase)**: 이벤트가 최상위 요소에서 시작하여 목표 요소까지 내려가는 단계
2. **타겟 단계 (Target Phase)**: 이벤트가 실제 목표 요소에 도달한 단계
3. **버블링 단계 (Bubbling Phase)**: 이벤트가 목표 요소에서 다시 최상위 요소로 전파되는 단계

---

## 2️⃣ 이벤트 리스너 등록

이벤트를 처리하기 위해서는 **이벤트 리스너(Event Listener)** 를 등록해야 합니다.  
대표적인 방법은 `addEventListener` 메서드를 사용하는 것입니다.

```js
const button = document.querySelector("#myButton");

button.addEventListener("click", function(event) {
  console.log("버튼이 클릭되었습니다!");
});
```

✔ `addEventListener`의 세 번째 매개변수로 `true`를 전달하면 **캡처링 단계**에서 이벤트를 처리하고, `false` 또는 생략하면 **버블링 단계**에서 이벤트를 처리합니다.  

---

## 3️⃣ 이벤트 객체

- 이벤트 리스너의 콜백 함수는 **이벤트 객체(Event Object)** 를 매개변수로 받습니다.  
- 이 객체를 통해 이벤트에 대한 다양한 정보를 얻을 수 있습니다.

```js
button.addEventListener("click", function(event) {
  console.log("이벤트 타입:", event.type);
  console.log("이벤트 발생 위치:", event.clientX, event.clientY);
});
```

✔  `event.type`: 이벤트의 종류 (예: "click", "keydown")  
✔  `event.clientX`, `event.clientY`: 이벤트 발생 시 마우스의 좌표  

---

## 4️⃣ 이벤트 전파 제어

이벤트의 전파를 제어하기 위해 다음 메서드를 사용할 수 있습니다

- `event.stopPropagation()`: 현재 이벤트의 전파를 중지하여 부모 요소로의 전파를 막습니다.
- `event.stopImmediatePropagation()`: 현재 이벤트의 전파를 중지하고, 같은 요소에 등록된 다른 이벤트 리스너의 실행도 막습니다.

```js
button.addEventListener("click", function(event) {
  event.stopPropagation();
  console.log("이 이벤트는 버블링되지 않습니다.");
});
```

---

## 5️⃣ 기본 동작 방지

- 일부 요소는 기본 동작을 가지고 있습니다. 예를 들어, `<a>` 태그는 클릭 시 링크로 이동합니다.  
- 이러한 기본 동작을 막기 위해 `event.preventDefault()`를 사용할 수 있습니다.

```js
const link = document.querySelector("#myLink");

link.addEventListener("click", function(event) {
  event.preventDefault();
  console.log("링크의 기본 동작이 방지되었습니다.");
});
```

---

## 6️⃣ 이벤트 위임

- **이벤트 위임(Event Delegation)** 은 상위 요소에 이벤트 리스너를 등록하여 하위 요소의 이벤트를 처리하는 기법입니다.  
- 이는 동적으로 생성된 요소에도 이벤트를 적용할 수 있어 효율적입니다.

```js
const list = document.querySelector("#myList");

list.addEventListener("click", function(event) {
  if (event.target && event.target.nodeName === "LI") {
    console.log("리스트 아이템이 클릭되었습니다:", event.target.textContent);
  }
});
```

---

## 🎯 정리

✔ **이벤트 모델**: 캡처링 → 타겟 → 버블링 단계로 진행  
✔ **이벤트 리스너 등록**: `addEventListener`를 사용하여 이벤트 핸들러를 등록  
✔ **이벤트 객체**: 이벤트에 대한 정보를 제공하며, 핸들러의 매개변수로 전달됨  
✔ **이벤트 전파 제어**: `stopPropagation`, `stopImmediatePropagation`으로 전파를 막을 수 있음  
✔ **기본 동작 방지**: `preventDefault`로 요소의 기본 동작을 막을 수 있음  
✔ **이벤트 위임**: 상위 요소에 이벤트 리스너를 등록하여 하위 요소의 이벤트를 처리하는 기법  
