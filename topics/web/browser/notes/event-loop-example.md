# 🆚 Browser - setTimeout vs Promise.then

`setTimeout()` 과 `Promise.then()` 은 둘 다 **비동기 함수**처럼 보이지만, **이벤트 루프에서 처리되는 방식**과 **실행 타이밍**은 다릅니다.

---

## 1️⃣ 기본 개념

| 항목 | `setTimeout(fn, 0)` | `Promise.then(fn)` |
|------|----------------------|---------------------|
| 목적 | 일정 시간 후 실행 | 비동기 작업 완료 후 실행 |
| 반환값 | Timeout ID (숫자) | Promise 객체 |
| 실행 시점 | **Task Queue** | **Microtask Queue** |
| 우선순위 | 낮음 (다음 루프) | 높음 (현재 루프 끝) |

---

## 2️⃣ 이벤트 루프 우선순위 비교

1. Call Stack (동기 코드 실행)
2. Microtask Queue (Promise.then, queueMicrotask)
3. [렌더링 타이밍]
4. Task Queue (setTimeout, setInterval, fetch 등)

✔ `Promise.then()` 은 **setTimeout보다 먼저 실행됨**

---

## 3️⃣ 실행 순서 예제

```js
console.log("1");

setTimeout(() => {
  console.log("setTimeout");
}, 0);

Promise.resolve().then(() => {
  console.log("promise");
});

console.log("2");
```

### 출력 결과
```
1
2
promise       // Microtask
setTimeout    // Task
```

---

## 4️⃣ 실행 타이밍 차이

| 구분 | 설명 |
|------|------|
| `setTimeout(fn, 0)` | "0초"여도 **현재 루프가 끝난 다음** 실행됨 |
| `Promise.then()` | **콜스택이 비자마자 즉시** Microtask로 실행됨 |

---

## 5️⃣ 실전에서 언제 써야 할까?

| 상황 | 추천 방식 | 이유 |
|------|-------------|------|
| 렌더링 이후에 실행하고 싶다 | `setTimeout` | Task로 밀어낼 수 있음 |
| 콜스택 이후 바로 실행하고 싶다 | `Promise.then` or `queueMicrotask` | Microtask가 즉시 실행됨 |
| 짧고 빠른 비동기 로직 처리 | `Promise.then` | 빠른 실행 보장 |
| 느긋하게 UI 반응 대기 | `setTimeout` | 렌더링 후 타이밍 확보 가능 |

---

## 6️⃣ 추가 비교: Microtask vs Task

| 항목 | Microtask | Task (Macro Task) |
|------|-----------|-------------------|
| 큐 종류 | `Promise.then`, `queueMicrotask` | `setTimeout`, `setInterval`, `fetch`, `click 이벤트` 등 |
| 처리 시점 | 현재 루프 끝나고 즉시 | 다음 루프 사이클에서 실행 |
| 렌더링 타이밍 | Microtask는 렌더링 전에 실행 | Task는 렌더링 이후 실행됨 |
| 예외 처리 | 에러 전파 빠름 | 상대적으로 느림 |

---

## 🎯 정리

✔ `Promise.then` → **Microtask Queue**, 현재 루프 끝에 실행됨  
✔ `setTimeout` → **Task Queue**, 다음 루프 사이클에 실행됨  
✔ 그래서 **`Promise`는 `setTimeout`보다 먼저 실행됨**  
✔ `setTimeout(fn, 0)`은 **실제로 0ms가 아님** → 최소 지연 시간 존재 (보통 4ms 이상)  
✔ 두 함수는 실행 타이밍 목적에 따라 **용도가 다름**
