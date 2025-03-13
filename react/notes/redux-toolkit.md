# 🔄 Redux Toolkit 개요

**Redux Toolkit (RTK)**는 **Redux의 보일러플레이트 코드를 줄이고 쉽게 상태 관리를 할 수 있도록 도와주는 공식 라이브러리**입니다.  
기존 Redux보다 **더 간결한 문법, 강력한 기능, 내장 미들웨어 지원**을 제공합니다.

---

## 1. Redux Toolkit의 필요성

✅ **기존 Redux의 단점 해결**  
   - `createStore`, `combineReducers`, `applyMiddleware` 등 설정이 복잡함  
   - 액션과 리듀서를 분리해야 해서 코드가 많아짐  
   - `useSelector()`, `useDispatch()`를 사용할 때 코드가 길어짐  

✅ **Redux Toolkit의 장점**  
   - `configureStore()` → **스토어 설정 간소화**  
   - `createSlice()` → **액션과 리듀서를 한 번에 작성 가능**  
   - `createAsyncThunk()` → **비동기 처리 쉽게 구현 가능**  
   - Redux DevTools 및 미들웨어 자동 설정  

---

## 2. Redux Toolkit 설치

```sh
npm install @reduxjs/toolkit react-redux
```
✅ `@reduxjs/toolkit` → Redux Toolkit 라이브러리  
✅ `react-redux` → React에서 Redux를 사용할 수 있도록 연결  

---

## 3. Redux Toolkit의 주요 기능

### 3-1. `configureStore()` - 스토어 설정
```jsx
import { configureStore } from "@reduxjs/toolkit";
import counterReducer from "./counterSlice";

const store = configureStore({
  reducer: {
    counter: counterReducer,
  },
});

export default store;
```
✅ 기존 Redux의 `createStore()`보다 간결하게 **스토어 설정 가능**  
✅ `reducer` 객체에 여러 개의 slice 리듀서를 추가 가능  

---

### 3-2. `createSlice()` - 액션 & 리듀서 한 번에 정의
```jsx
import { createSlice } from "@reduxjs/toolkit";

const counterSlice = createSlice({
  name: "counter",
  initialState: { count: 0 },
  reducers: {
    increment: (state) => { state.count += 1; },
    decrement: (state) => { state.count -= 1; },
  },
});

export const { increment, decrement } = counterSlice.actions;
export default counterSlice.reducer;
```
✅ `reducers` 안에서 **객체 형태로 액션 & 리듀서 동시 정의**  
✅ 기존 Redux에서는 **액션과 리듀서를 따로 관리했지만, RTK에서는 `createSlice()`로 합칠 수 있음**  
✅ **불변성 관리를 자동 처리 (`immer` 사용)** → `state.count += 1` 가능  

---

### 3-3. Redux Store를 React에 연결 (`Provider`)
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
✅ `<Provider store={store}>` → **Redux Store를 React 애플리케이션에 제공**  

---

### 3-4. Redux 상태 가져오기 (`useSelector`) & 업데이트 (`useDispatch`)
```jsx
import { useSelector, useDispatch } from "react-redux";
import { increment, decrement } from "./counterSlice";

function Counter() {
  const count = useSelector((state) => state.counter.count);
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
✅ `useSelector(state => state.counter.count)` → **Redux Store에서 상태 가져오기**  
✅ `useDispatch()` → **Redux 액션(dispatch) 실행**  

---

## 4. 비동기 작업 (`createAsyncThunk`)

### 4-1. `createAsyncThunk()` - API 호출 처리
Redux에서 **비동기 작업(API 호출 등)을 쉽게 처리할 수 있도록 지원**합니다.

```jsx
import { createSlice, createAsyncThunk } from "@reduxjs/toolkit";

export const fetchUser = createAsyncThunk("user/fetchUser", async () => {
  const response = await fetch("https://jsonplaceholder.typicode.com/users/1");
  return response.json();
});

const userSlice = createSlice({
  name: "user",
  initialState: { user: null, status: "idle" },
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchUser.pending, (state) => {
        state.status = "loading";
      })
      .addCase(fetchUser.fulfilled, (state, action) => {
        state.user = action.payload;
        state.status = "succeeded";
      })
      .addCase(fetchUser.rejected, (state) => {
        state.status = "failed";
      });
  },
});

export default userSlice.reducer;
```
✅ `createAsyncThunk("액션명", 비동기 함수)` → API 요청을 정의  
✅ `extraReducers` → **비동기 요청 상태 (`pending`, `fulfilled`, `rejected`) 관리**  

---

### 4-2. API 호출 데이터 표시 (`useDispatch` 활용)
```jsx
import { useEffect } from "react";
import { useSelector, useDispatch } from "react-redux";
import { fetchUser } from "./userSlice";

function UserProfile() {
  const { user, status } = useSelector((state) => state.user);
  const dispatch = useDispatch();

  useEffect(() => {
    dispatch(fetchUser());
  }, [dispatch]);

  if (status === "loading") return <p>로딩 중...</p>;
  if (status === "failed") return <p>에러 발생</p>;

  return <p>사용자 이름: {user?.name}</p>;
}

export default UserProfile;
```
✅ `dispatch(fetchUser())`를 실행하면 자동으로 API 요청이 진행됨  
✅ `status`를 활용하여 로딩, 성공, 실패 상태를 관리  

---

## 5. Redux Toolkit vs 기존 Redux 비교

| 기능 | Redux | Redux Toolkit |
|------|--------|---------------|
| 설치 패키지 | `redux` + `react-redux` | `@reduxjs/toolkit` + `react-redux` |
| 스토어 설정 | `createStore() + combineReducers()` | `configureStore()` |
| 액션 정의 | `actions.js`에서 따로 정의 | `createSlice().actions`에서 자동 생성 |
| 리듀서 정의 | 여러 개의 `switch(action.type)` 필요 | `createSlice().reducers`로 간편 관리 |
| 비동기 작업 | `redux-thunk` 설치 필요 | `createAsyncThunk()` 기본 제공 |
| 코드량 | 많음 (보일러플레이트 많음) | 적음 (간결한 코드) |

✅ **Redux Toolkit은 기존 Redux보다 코드가 짧고 직관적**  
✅ **비동기 작업을 쉽게 관리 가능 (`createAsyncThunk`)**  

---

## 6. 공식 문서 및 추가 자료
- [Redux Toolkit 공식 문서](https://redux-toolkit.js.org/)
- [React Redux 공식 문서](https://react-redux.js.org/)
- [Redux vs Redux Toolkit 비교](https://redux.js.org/tutorials/fundamentals/part-8-modern-redux)

---

🚀 **Redux Toolkit을 익혔다면, 이제 React Router를 배워봅시다!**  
다음 개념: [React Router](./react-router.md) →
