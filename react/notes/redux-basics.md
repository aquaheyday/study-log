# 🔄 Redux 기본 개념 (Redux Basics)

**Redux**는 **React 애플리케이션에서 상태를 효율적으로 관리하기 위한 상태 관리 라이브러리**입니다.  
Redux를 사용하면 **전역 상태(Global State)를 관리**하고, 컴포넌트 간의 **데이터 흐름을 예측 가능하게** 만들 수 있습니다.

---

## 1. Redux란?

✅ **중앙 집중식 상태 관리** (전역 상태 저장소 `Store` 활용)  
✅ **단방향 데이터 흐름 (Flux 패턴)** → **State → View → Action → Reducer**  
✅ **React와 독립적** → Vue, Angular에서도 사용 가능  

---

## 2. Redux 동작 흐름

Redux는 **3가지 핵심 개념**을 기반으로 작동합니다.

1️⃣ **Store (저장소)**
   - 애플리케이션의 **전역 상태를 저장**
   - `createStore()`를 사용하여 생성

2️⃣ **Action (액션)**
   - 상태(State)를 변경하는 **이벤트 객체**
   - `{ type: "INCREMENT" }` 형태로 사용

3️⃣ **Reducer (리듀서)**
   - `Action`을 받아 **새로운 State를 반환하는 함수**
   - `switch(action.type)`을 사용하여 상태 변경

### 2-1. Redux 동작 개념도
```plaintext
사용자 이벤트 → Action → Reducer → Store 업데이트 → UI 리렌더링
```

---

## 3. Redux 설치 및 설정

### 3-1. Redux 설치
```sh
npm install redux react-redux
```
✅ `redux` → Redux 코어 라이브러리  
✅ `react-redux` → React에서 Redux를 쉽게 사용하기 위한 라이브러리  

---

## 4. Redux 기본 코드 구조

### 4-1. `store.js` - 전역 상태 저장소 생성
```jsx
import { createStore } from "redux";
import counterReducer from "./counterReducer";

const store = createStore(counterReducer);
export default store;
```
✅ `createStore(reducer)` → Redux 스토어 생성  

---

### 4-2. `counterReducer.js` - 리듀서 정의
```jsx
const initialState = { count: 0 };

function counterReducer(state = initialState, action) {
  switch (action.type) {
    case "INCREMENT":
      return { count: state.count + 1 };
    case "DECREMENT":
      return { count: state.count - 1 };
    default:
      return state;
  }
}

export default counterReducer;
```
✅ `action.type`에 따라 **State를 변경**  

---

### 4-3. `actions.js` - 액션 생성
```jsx
export const increment = () => ({ type: "INCREMENT" });
export const decrement = () => ({ type: "DECREMENT" });
```
✔ 액션 객체 `{ type: "INCREMENT" }` 반환  

---

### 4-4. `index.js` - React와 Redux 연결
```jsx
import React from "react";
import ReactDOM from "react-dom";
import { Provider } from "react-redux";
import store from "./store";
import Counter from "./Counter";

ReactDOM.render(
  <Provider store={store}>
    <Counter />
  </Provider>,
  document.getElementById("root")
);
```
✅ `<Provider store={store}>` → **Redux Store를 React에 연결**  

---

### 4-5. `Counter.js` - Redux 상태 사용하기
```jsx
import { useSelector, useDispatch } from "react-redux";
import { increment, decrement } from "./actions";

function Counter() {
  const count = useSelector((state) => state.count);
  const dispatch = useDispatch();

  return (
    <div>
      <p>카운트: {count}</p>
      <button onClick={() => dispatch(increment())}>+1</button>
      <button onClick={() => dispatch(decrement())}>-1</button>
    </div>
  );
}

export default Counter;
```
✅ `useSelector(state => state.count)` → Redux `store`에서 값 가져오기  
✅ `useDispatch()` → `dispatch(action)`을 실행하여 상태 변경  

---

## 5. Redux의 핵심 원칙

| 원칙 | 설명 |
|------|------|
| **Single Source of Truth** | 상태는 **단 하나의 Store**에서 관리됨 |
| **State is Read-Only** | 상태는 직접 수정할 수 없고, **Action을 통해 변경**해야 함 |
| **Changes are Made with Pure Functions** | 상태 변경은 **순수 함수(Reducer)** 를 통해 이루어짐 |

---

## 6. Redux 사용 시 장점과 단점

### ✅ **Redux의 장점**
✔ **컴포넌트 간 상태 공유가 쉬움**  
✔ **상태 변경이 예측 가능 (Reducer + Action 패턴 사용)**  
✔ **Redux DevTools를 활용한 디버깅 가능**  

### ❌ **Redux의 단점**
✔ **작은 프로젝트에서는 불필요한 복잡성 추가 가능**  
✔ **초기 설정이 복잡하고 보일러플레이트 코드가 많음**  
✔ **매번 Action을 생성해야 하므로 코드량이 증가**  

---

## 7. Redux vs Context API 비교

| 특징 | Redux | Context API |
|------|-------|------------|
| 사용 목적 | **복잡한 상태 관리** | **간단한 전역 상태 관리** |
| 상태 저장소 | `createStore()` | `createContext()` |
| 데이터 흐름 | 단방향 (`dispatch(action) → reducer → store`) | 단방향 (`Provider → useContext()`) |
| 리렌더링 최적화 | ✅ 미들웨어 활용 가능 | ❌ Provider 값이 바뀌면 전체 리렌더링 |

✅ **Redux는 상태 변경이 빈번한 대규모 앱에 적합**  
✅ **Context API는 작은 프로젝트에서 간단한 전역 상태 관리에 적합**  

---

## 8. Redux DevTools 설정 (디버깅 도구)

Redux 개발을 편리하게 하기 위해 **Redux DevTools**를 사용하면 **상태 변경 내역을 쉽게 추적 가능**합니다.

### 8-1. Redux DevTools 설치
```sh
npm install redux-devtools-extension
```

### 8-2. `store.js`에서 설정 추가
```jsx
import { createStore } from "redux";
import { composeWithDevTools } from "redux-devtools-extension";
import counterReducer from "./counterReducer";

const store = createStore(counterReducer, composeWithDevTools());
export default store;
```
✔ `composeWithDevTools()`를 추가하여 Redux 상태를 쉽게 확인 가능  

---

## 9. 공식 문서 및 추가 자료
- [Redux 공식 문서](https://redux.js.org/)
- [React Redux 공식 문서](https://react-redux.js.org/)
- [Redux DevTools 사용법](https://github.com/reduxjs/redux-devtools)

---

🚀 **Redux의 기본 개념을 이해했다면, 이제 Redux Toolkit을 배워봅시다!**  
다음 개념: [Redux Toolkit](./redux-toolkit.md) →
