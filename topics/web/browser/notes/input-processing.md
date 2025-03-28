# 🎮 Browser - 사용자 입력 처리 흐름

## 1️⃣ 입력 처리란?

- 사용자가 **키보드, 마우스, 터치, 포인터 등으로 상호작용**했을 때, 브라우저가 이를 감지하고 **이벤트를 생성/전파/처리**하는 전체 흐름

---

## 2️⃣ 전체 흐름 요약

```text
[사용자 입력 발생]
       ↓
[브라우저가 이벤트 생성]
       ↓
[DOM 트리 위에서 이벤트 전파 (캡처링 → 타겟 → 버블링)]
       ↓
[JS에서 등록한 이벤트 핸들러 호출]
       ↓
[기본 동작 수행 or 차단 (preventDefault)]
```

---

## 3️⃣ 이벤트 생성 단계

- 키보드 입력: `keydown`, `keypress`, `keyup`
- 마우스 입력: `click`, `mousedown`, `mouseup`, `mousemove`
- 터치 입력: `touchstart`, `touchmove`, `touchend`
- 기타: `focus`, `input`, `change` 등

---

## 4️⃣ 이벤트 전파 단계

### 3단계 이벤트 전파

#### 1. 캡처링(Capturing) 단계
- 최상위(`window`, `document`) → 이벤트 타겟까지 내려옴

#### 2. 타겟(Target) 단계
- 실제 이벤트가 발생한 DOM 요소 도달

#### 3. 버블링(Bubbling) 단계
- 타겟 → 상위 DOM으로 이벤트가 다시 올라감

---

## 5️⃣ 이벤트 등록 예시

```js
const btn = document.querySelector("#myBtn");

// 버블링 단계 (기본)
btn.addEventListener("click", (e) => {
  console.log("버튼 클릭됨");
});

// 캡처링 단계
btn.addEventListener("click", (e) => {
  console.log("캡처링 단계");
}, true);
```

#### `addEventListener`의 세 번째 인자: `useCapture`

| 값 | 의미 |
|------|------|
| true	| 캡처링 단계에서 이벤트 실행 |
| false	| 버블링 단계에서 실행 (기본값) |

---

## 6️⃣ 이벤트 전파 제어

| 메서드 | 설명 |
|--------|------|
| `e.stopPropagation()` | 부모로의 전파 중단 |
| `e.stopImmediatePropagation()` | 동일 요소의 다른 핸들러도 중단 |
| `e.preventDefault()` | 기본 동작 차단 (ex: 폼 제출, 링크 이동 등) |

---

## 7️⃣ 입력값 처리 흐름 (예: input 태그)


#### 사용자 입력 → input 이벤트 발생 → 이벤트 핸들러에서 상태 반영 → DOM/UI 업데이트

```js
document.querySelector("input").addEventListener("input", (e) => {
  console.log("입력된 값:", e.target.value);
});
```

---

## 8️⃣ React 등 프레임워크의 역할

- 프레임워크는 브라우저 이벤트를 **랩핑하여 내부 시스템으로 전달**
- 성능 최적화 & 일관된 동작 제공
- 예: React의 Synthetic Event

```js
function MyInput() {
  return <input onChange={(e) => console.log(e.target.value)} />;
}
```

---

## 🎯 정리

✔ 브라우저는 입력 시 **이벤트 객체를 생성하고, DOM을 따라 전파**  
✔ 이벤트는 **캡처링 → 타겟 → 버블링** 순서로 흐름  
✔ JS는 원하는 지점에서 **이벤트 핸들러를 등록하고 전파 제어 가능**  
✔ UI 업데이트는 입력 이벤트를 **상태나 DOM에 반영하면서 발생**

