# 🔲 useEffect 훅

`useEffect`는 **React에서 컴포넌트의 생명주기(lifecycle)를 관리하는 훅**입니다.  
렌더링 후 실행되어 **데이터 패칭, DOM 조작, 구독(subscription) 설정** 등에 사용됩니다.

---

## 1. `useEffect`란?

- 컴포넌트가 렌더링될 때 실행되는 함수  
- 클래스 컴포넌트의 `componentDidMount`, `componentDidUpdate`, `componentWillUnmount`를 대체  
- Side Effect(부작용) 처리: API 요청, 타이머 설정, 이벤트 리스너 등록 등  
- 렌더링 후 실행되며, 특정 값이 변경될 때 다시 실행 가능  

```jsx
import { useEffect } from "react";

function Example() {
  useEffect(() => {
    console.log("컴포넌트가 렌더링됨!");
  });

  return <h1>Hello, useEffect!</h1>;
}
```
✔ `useEffect(() => { ... })` → 기본적으로 렌더링 후 실행됨  

---

## 2. `useEffect`의 실행 타이밍과 의존성 배열

### 컴포넌트가 렌더링될 때마다 실행 (기본)
```jsx
useEffect(() => {
  console.log("렌더링됨!");
});
```
✔ 렌더링이 발생할 때마다 실행됨 (의존성 배열 없음)

---

### 마운트(처음 렌더링)될 때만 실행
```jsx
useEffect(() => {
  console.log("한 번만 실행됨!");
}, []);
```
✔ `[]` → 빈 배열을 전달하면 처음 렌더링될 때만 실행됨  

---

### 특정 값이 변경될 때 실행
```jsx
import { useState, useEffect } from "react";

function Counter() {
  const [count, setCount] = useState(0);

  useEffect(() => {
    console.log(`count가 변경됨: ${count}`);
  }, [count]);

  return (
    <div>
      <p>카운트: {count}</p>
      <button onClick={() => setCount(count + 1)}>+1</button>
    </div>
  );
}
```
✔ `[count]` → `count` 값이 변경될 때만 실행됨  

---

## 3. `useEffect`에서 정리(cleanup) 함수 사용하기

`useEffect`는 **컴포넌트가 언마운트되거나 의존성이 변경될 때 정리(cleanup) 작업을 수행할 수 있음**  

### 이벤트 리스너 제거
```jsx
import { useEffect } from "react";

function WindowResize() {
  useEffect(() => {
    const handleResize = () => {
      console.log(`창 크기 변경됨: ${window.innerWidth}px`);
    };

    window.addEventListener("resize", handleResize);

    return () => {
      console.log("이벤트 리스너 제거됨");
      window.removeEventListener("resize", handleResize);
    };
  }, []);

  return <h1>창 크기를 조절해보세요!</h1>;
}
```
✔ `return () => {}` → 컴포넌트가 언마운트될 때 실행됨  
✔ 이벤트 리스너를 제거하여 메모리 누수 방지  

---

### 인터벌(Interval) 정리
```jsx
import { useState, useEffect } from "react";

function Timer() {
  const [seconds, setSeconds] = useState(0);

  useEffect(() => {
    const interval = setInterval(() => {
      setSeconds((prev) => prev + 1);
    }, 1000);

    return () => {
      console.log("타이머 정리됨");
      clearInterval(interval);
    };
  }, []);

  return <p>타이머: {seconds}초</p>;
}
```
✔ `setInterval()`을 사용하면 매초마다 `seconds`가 증가  
✔ 컴포넌트가 언마운트되면 `clearInterval()`을 호출하여 메모리 누수를 방지  

---

## 4. `useEffect`의 주요 활용 사례

### API 데이터 불러오기 (Fetching Data)
```jsx
import { useState, useEffect } from "react";

function FetchData() {
  const [data, setData] = useState(null);

  useEffect(() => {
    fetch("https://jsonplaceholder.typicode.com/todos/1")
      .then((response) => response.json())
      .then((json) => setData(json));
  }, []);

  return <pre>{JSON.stringify(data, null, 2)}</pre>;
}
```
✔ `fetch()`를 사용하여 데이터 요청 후, 응답을 `setData()`로 저장  
✔ 의존성 배열 `[]`을 추가하여 한 번만 실행되도록 설정  

---

### 다크 모드 상태 저장 (로컬 스토리지 활용)
```jsx
import { useState, useEffect } from "react";

function DarkModeToggle() {
  const [darkMode, setDarkMode] = useState(
    localStorage.getItem("darkMode") === "true"
  );

  useEffect(() => {
    localStorage.setItem("darkMode", darkMode);
  }, [darkMode]);

  return (
    <button onClick={() => setDarkMode(!darkMode)}>
      {darkMode ? "🌙 다크 모드" : "☀️ 라이트 모드"}
    </button>
  );
}
```
✔ `localStorage.getItem()` → 다크 모드 설정을 저장하고 불러오기  
✔ `[darkMode]` → `darkMode` 값이 변경될 때마다 `localStorage` 업데이트  

---

## 5. `useEffect` 사용 시 주의할 점

| 주의사항 | 해결 방법 |
|---------|----------|
| 무한 루프 발생 가능 | 의존성 배열(`[]`)을 올바르게 설정 |
| 의존성 배열 생략 시 불필요한 렌더링 발생 | `useEffect` 실행 조건을 명확히 지정 |
| 이벤트 리스너 제거 필요 | `return () => {}`을 사용하여 정리(cleanup) |

---

## 🎯 정리
✔ useEffect는 React의 Side Effect(부작용) 처리를 위한 훅  
✔ `useEffect(() => {...})` → 렌더링될 때마다 실행  
✔ `useEffect(() => {...}, [])` → 마운트(처음 렌더링) 시 한 번만 실행  
✔ `useEffect(() => {...}, [count])` → count 값이 변경될 때만 실행  
✔ Cleanup 함수 (`return () => {...}`) 사용하여 이벤트 리스너 및 인터벌 제거 가능  
