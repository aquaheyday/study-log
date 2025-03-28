# 🌐 Browser - Web API & `window` 객체

## 1️⃣ Web API란?

- **Web API**는 브라우저가 제공하는 **자바스크립트에서 사용할 수 있는 기능 모음**
- 브라우저 내장 API로, **자바스크립트 언어 자체에는 없지만**, 브라우저 환경에서는 사용 가능

| 분류 | 대표 API |
|------|----------|
| DOM 조작 | `document`, `getElementById`, `querySelector` |
| 이벤트 처리 | `addEventListener`, `dispatchEvent` |
| 타이머 | `setTimeout`, `setInterval` |
| 네트워크 | `fetch`, `XMLHttpRequest`, `WebSocket` |
| 브라우저 기능 | `localStorage`, `location`, `navigator`, `history` |
| 멀티미디어 | `MediaDevices`, `Canvas`, `Audio`, `Video` |
| 기타 | `Geolocation`, `Notification`, `Clipboard`, `Fullscreen` |

---

## 2️⃣ `window` 객체란?

- 브라우저 환경에서 **모든 전역 객체와 API를 포함하는 최상위 객체**
- 모든 웹 페이지는 자동으로 하나의 `window` 객체를 가짐

```js
console.log(window === globalThis); // true (브라우저에서는)
```

---

## 3️⃣ `window` 객체의 주요 역할

### 1) 전역 객체

`var`로 선언된 변수는 `window`의 속성이 됨

```js
var x = 1;
console.log(window.x); // 1
```

✔ `let`, `const`는 `window`에 등록되지 않음  

---

### 2) 전역 함수도 포함

| 함수 이름 | 설명 |
|-----------|------|
| `alert()` | 경고창 |
| `setTimeout()` | 지정 시간 후 실행 |
| `setInterval()` | 반복 실행 |
| `fetch()` | HTTP 요청 |
| `addEventListener()` | 이벤트 리스너 등록 |

```js
window.alert("Hello"); // alert()과 동일
```

---

### 3) 브라우저 상태/정보

| 속성 | 설명 |
|------|------|
| `window.location` | 현재 URL, 리디렉션 등 |
| `window.history` | 이전/다음 페이지 이동 |
| `window.navigator` | 브라우저, OS 정보 |
| `window.screen` | 해상도, 화면 크기 |
| `window.innerWidth` / `innerHeight` | 뷰포트 크기 |

---

## 4️⃣ `window`와 `document`의 차이

| 객체 | 설명 |
|------|------|
| `window` | 브라우저 창 전체, Web API 포함 |
| `document` | HTML 문서 (DOM) 자체 |

```js
window.document === document; // true
```

✔ `document`는 `window`의 속성 중 하나이다  

---

## 5️⃣ Web API 사용 예시

```js
// 3초 후에 메시지 출력
setTimeout(() => {
  alert("3초 지남!");
}, 3000);

// 현재 URL 출력
console.log(location.href);

// DOM 조작
document.getElementById("title").textContent = "변경됨";
```

---

## 6️⃣ Web API는 어디에 정의돼 있나?

- 브라우저 엔진 (예: Chromium, WebKit, Gecko 등)에 내장된 **C++ 기반의 네이티브 API**
- JavaScript 엔진 (예: V8)은 언어만 처리, Web API는 브라우저가 처리

```txt
JS 실행 컨텍스트 ↔ Web API (콜백 큐, 타이머 등)
```

---

## 7️⃣ JS 실행 컨텍스트 ↔ Web API 흐름이란?

JS는 단일 스레드지만, 브라우저(Web API)는 백그라운드 작업을 맡아서 처리한 뒤, 콜백을 이벤트 루프를 통해 다시 JS로 전달하는 구조

### 1) 흐름 예시
```text
JS 코드 실행
  ↓
Call Stack (JS 싱글 스레드)
  ↓
setTimeout, fetch, event → Web API에게 위임
  ↓
Web API가 완료되면 콜백 → Callback Queue로 이동
  ↓
Call Stack이 비면 → Event Loop가 Queue에서 꺼내서 실행
```

---

### 2) 주요 구성요소

| 구성 요소 |	역할 |
|------|------|
| Call Stack (실행 컨텍스트) |	JS 코드 실행 순서 관리 (함수 호출 시 쌓이고, 끝나면 빠짐) |
| Web API (브라우저 기능)	| setTimeout, DOM, fetch, event listener 등 백그라운드 처리 |
| Callback Queue | 작업이 끝난 콜백을 잠시 대기시키는 곳 |
| Event Loop | Call Stack이 비면,  Queue에서 콜백을 꺼내 실행함 |

---

### 3) 실행 흐름 예제

#### 1. 예시 코드
```js
console.log("1");

setTimeout(() => {
  console.log("2");
}, 0);

console.log("3");
```

#### 2. 출력
```
1
3
2
```

#### 3. 실행 순서
1. "1" → 바로 실행
2. setTimeout() → Web API에 등록됨
3. "3" → 바로 실행
4. Web API에서 0ms 후 콜백 → Callback Queue
5. Call Stack이 비었으므로 → "2" 실행됨

---

## 🎯 정리

✔ Web API는 **브라우저가 제공하는 자바스크립트용 기능 집합**  
✔ `window`는 **모든 전역 변수, 함수, 객체를 포함하는 최상위 객체**  
✔ `document`는 그 안에 포함된 **DOM 전용 객체**  
✔ Web API는 브라우저 환경에서만 동작하며, Node.js 등에는 없음

