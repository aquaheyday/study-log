# 🔁 Browser - 이벤트 루프 & 큐 

자바스크립트는 **싱글 스레드(Single Thread)** 언어이지만, 비동기 작업을 처리하기 위해 브라우저는 **이벤트 루프(Event Loop)** 구조를 사용합니다.

---

## 1️⃣ 이벤트 루프(Event Loop)란?

**콜스택(Call Stack)**, **태스크 큐(Task Queue)**, **마이크로태스크 큐(Microtask Queue)** 를 관리하며 **비동기 코드가 실행되는 순서를 조절**하는 메커니즘

---

## 2️⃣ 이벤트 루프의 작동 흐름

#### Microtask > Task 순서로 항상 `Microtask`가 먼저 처리됨  

1. 콜스택(Call Stack) 이 비면,
2. 마이크로태스크 큐(Microtask Queue) 의 모든 작업을 처리
3. 그 다음 태스크 큐(Task Queue) 에서 작업을 하나 꺼내 실행
4. 다시 콜스택이 비면 → 2번부터 반복

#### 각 구성 요소 설명

| 구성 요소 | 설명 | 예시 |
|-----------|------|------|
| **Call Stack** | 현재 실행 중인 코드 (함수 호출 스택) | `console.log()` |
| **Task Queue (Macro Task)** | 타이머, 이벤트 등 비동기 큐 | `setTimeout`, `setInterval`, `DOM 이벤트` |
| **Microtask Queue** | Task보다 먼저 실행되는 비동기 큐 | `Promise.then`, `MutationObserver`, `queueMicrotask` |

---

## 3️⃣ 실행 순서 예제

#### 예제 코드
```js
console.log("1");

setTimeout(() => {
  console.log("2");
}, 0);

Promise.resolve().then(() => {
  console.log("3");
});

console.log("4");
```

#### 출력 순서
```
1
4
3   // (Microtask)
2   // (Task)
```

✔ `Promise.then`은 Microtask → 현재 루프 끝에서 실행  
✔ `setTimeout`은 Task → 다음 루프에서 실행  

#### 주의사항

- **Promise.then** 은 **setTimeout보다 먼저 실행됨**
- **Microtask가 너무 많으면** Task 실행이 **지연**될 수 있음
- 무한 루프 조심 (예: `Promise.then().then().then()` 무한 중첩)

---

## 4️⃣ 렌더링은 언제 발생할까?

- 브라우저는 렌더링이 필요한 경우, 이벤트 루프 사이클 중에 렌더링을 수행합니다.
- 일반적으로는 Microtask 처리 이후, Task 실행 전에 렌더링이 발생할 수 있습니다.
- 단, **Microtask가 너무 많으면** 렌더링 타이밍이 밀릴 수 있습니다.
- 따라서 UI 업데이트가 필요한 작업은 너무 무거운 Microtask로 처리하지 않는 것이 좋습니다.


```js
Promise.resolve().then(() => {
  while (true) {} // 렌더링 블로킹 발생 가능
});
```

✔ Microtask는 짧고 가볍게 유지하는 것이 좋습니다.

---

## 5️⃣ `queueMicrotask()`

- ES2019부터 지원되는 **마이크로태스크 등록 함수**
- `Promise.then()`과 동일한 우선순위를 가짐

```js
queueMicrotask(() => {
  console.log("microtask 실행");
});
```

✔ 비동기지만 **현재 이벤트 루프 사이클 끝에서 바로 실행**됨  
✔ **가볍고 짧은 작업**에 적합함

---

## 6️⃣ `requestAnimationFrame`

**브라우저의 다음 리페인트(화면 그리기) 직전에 지정한 콜백 함수를 실행**하는 API입니다.

✔️ 주로 **부드럽고 효율적인 애니메이션**을 만들 때 사용됩니다.

#### 동작 원리

- 브라우저는 보통 **1초에 60번(60fps)** 화면을 다시 그립니다. (16.6ms 간격)
- `requestAnimationFrame()`은 **그 다음 프레임 렌더링 직전에 콜백 함수를 실행** (Microtask 이후, 렌더 직전)
- 콜백 함수 내부에서  화면 변경 작업을 하면, 브라우저는 가장 효율적으로 렌더링할 타이밍에 맞춰서 그립니다.

#### 예제

```js
function moveBox(timestamp) {
  box.style.transform = `translateX(${timestamp / 10}px)`;
  requestAnimationFrame(moveBox);
}

requestAnimationFrame(moveBox);
```

- `timestamp`는 **고해상도 시간 (ms)** 이 자동으로 들어옴
- 위 예제는 프레임마다 위치를 조금씩 바꿔 **부드럽게 움직임**

#### 팁

- 브라우저 탭이 **비활성화되면 실행 중단**됨 (배터리 절약 & CPU 자원 절약)

---

## 7️⃣ 보너스 팁

- `setTimeout(fn, 0)`도 **0초가 아님** → 브라우저마다 **최소 지연 시간**이 있음 (보통 4ms 이상)
- DevTools의 **Performance 탭**을 통해 **이벤트 루프 실행 타이밍 확인** 가능

---

## 🎯 정리

| 구성 요소 | 설명 |
|-----------|------|
| **Call Stack** | 현재 실행 중인 코드 (동기 처리) |
| **Microtask Queue** | `Promise.then`, `queueMicrotask` 등 |
| **렌더링 타이밍** | 일반적으로 Microtask 후, Task 전 (필요 시) |
| **Task Queue** | `setTimeout`, `DOM 이벤트`, `fetch` 등 |

✔ 이벤트 루프는 **이 구조를 따라 루프를 돌며**  
✔ 코드 실행 순서를 **논리적으로, 효율적으로, 시각적으로** 처리함

