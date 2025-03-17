# 🔲 컴포넌트 라이프사이클

React 컴포넌트는 **생성(Mount), 업데이트(Update), 소멸(Unmount)** 단계를 거칩니다.  
React에서는 **함수형 컴포넌트 + `useEffect` 훅**을 사용하여 라이프사이클을 제어할 수 있으며,  
클래스형 컴포넌트에서는 **라이프사이클 메서드**를 사용합니다.

---

## 1. React 컴포넌트의 라이프사이클 단계

| 단계 | 설명 | 함수형 컴포넌트 (`useEffect`) | 클래스형 컴포넌트 |
|------|------|-----------------------------|------------------|
| **Mount (마운트)** | 컴포넌트가 생성되어 DOM에 추가됨 | `useEffect(() => {...}, [])` | `componentDidMount()` |
| **Update (업데이트)** | 상태(State)나 Props가 변경될 때 | `useEffect(() => {...}, [deps])` | `componentDidUpdate()` |
| **Unmount (언마운트)** | 컴포넌트가 제거될 때 | `useEffect(() => {... return () => {...}}, [])` | `componentWillUnmount()` |

---

## 2. Mount (마운트: 컴포넌트가 처음 렌더링될 때)

- 컴포넌트가 생성되고 DOM에 추가됨  
- API 호출, 이벤트 리스너 등록 등의 초기화 작업 수행  

### 함수형 컴포넌트에서 `useEffect` 사용
```jsx
import { useEffect } from "react";

function Component() {
  useEffect(() => {
    console.log("✅ 컴포넌트가 마운트됨 (처음 렌더링)");

    return () => {
      console.log("❌ 컴포넌트가 언마운트됨 (제거됨)");
    };
  }, []);

  return <h1>React 컴포넌트</h1>;
}

export default Component;
```
✔ `useEffect(() => {...}, [])` → 처음 한 번만 실행됨  

---

### 클래스형 컴포넌트에서 `componentDidMount`
```jsx
import React, { Component } from "react";

class ComponentExample extends Component {
  componentDidMount() {
    console.log("✅ 컴포넌트가 마운트됨");
  }

  render() {
    return <h1>React 클래스형 컴포넌트</h1>;
  }
}

export default ComponentExample;
```
✔ `componentDidMount()` → 컴포넌트가 처음 마운트될 때 실행됨  

---

## 3. Update (업데이트: 상태 또는 Props 변경 시 실행)

- State 또는 Props가 변경되면 실행됨  
- 렌더링이 다시 수행되며, 변경된 값이 반영됨  

### 함수형 컴포넌트에서 `useEffect` 사용
```jsx
import { useState, useEffect } from "react";

function Counter() {
  const [count, setCount] = useState(0);

  useEffect(() => {
    console.log(`🔄 카운트 변경됨: ${count}`);
  }, [count]);

  return (
    <div>
      <p>카운트: {count}</p>
      <button onClick={() => setCount(count + 1)}>+1</button>
    </div>
  );
}

export default Counter;
```
✔ `[count]` → `count` 값이 변경될 때마다 `useEffect` 실행됨  

---

### 클래스형 컴포넌트에서 `componentDidUpdate`
```jsx
import React, { Component } from "react";

class Counter extends Component {
  state = { count: 0 };

  componentDidUpdate(prevProps, prevState) {
    if (prevState.count !== this.state.count) {
      console.log(`🔄 카운트 변경됨: ${this.state.count}`);
    }
  }

  render() {
    return (
      <div>
        <p>카운트: {this.state.count}</p>
        <button onClick={() => this.setState({ count: this.state.count + 1 })}>+1</button>
      </div>
    );
  }
}

export default Counter;
```
✔ `componentDidUpdate()` → **이전 상태(prevState)와 현재 상태 비교 가능**  

---

## 4. Unmount (언마운트: 컴포넌트가 제거될 때)

- 컴포넌트가 DOM에서 제거될 때 실행됨  
- 이벤트 리스너 제거, 타이머 정리, API 구독 해제 등의 작업 수행  

### 함수형 컴포넌트에서 `useEffect` 정리 함수 사용
```jsx
import { useEffect } from "react";

function Timer() {
  useEffect(() => {
    const interval = setInterval(() => {
      console.log("⏳ 1초마다 실행");
    }, 1000);

    return () => {
      console.log("❌ 타이머 정리됨");
      clearInterval(interval);
    };
  }, []);

  return <h1>타이머 실행 중...</h1>;
}

export default Timer;
```
✔ `return () => {}` → **컴포넌트가 제거될 때 실행됨**  

---

### 클래스형 컴포넌트에서 `componentWillUnmount`
```jsx
import React, { Component } from "react";

class Timer extends Component {
  componentDidMount() {
    this.interval = setInterval(() => {
      console.log("⏳ 1초마다 실행");
    }, 1000);
  }

  componentWillUnmount() {
    console.log("❌ 타이머 정리됨");
    clearInterval(this.interval);
  }

  render() {
    return <h1>타이머 실행 중...</h1>;
  }
}

export default Timer;
```
✔ `componentWillUnmount()` → **컴포넌트가 제거될 때 실행됨**  

---

## 5. `useEffect`와 라이프사이클 메서드 비교

| 단계 | 클래스형 컴포넌트 | 함수형 컴포넌트 (`useEffect`) |
|------|------------------|-----------------------------|
| 마운트 | `componentDidMount()` | `useEffect(() => {...}, [])` |
| 업데이트 | `componentDidUpdate()` | `useEffect(() => {...}, [deps])` |
| 언마운트 | `componentWillUnmount()` | `useEffect(() => {... return () => {...}}, [])` |

✔ 클래스형 컴포넌트의 라이프사이클 메서드는 `useEffect`로 대체 가능!  

---

## 🎯 정리
✔ React 컴포넌트는 마운트(Mount) → 업데이트(Update) → 언마운트(Unmount) 단계를 거침  
✔ `useEffect(() => {...}, []) (componentDidMount)` → 마운트 시 실행  
✔ `useEffect(() => {...}, [state]) (componentDidUpdate)` → 업데이트 시 실행  
✔ `useEffect(() => {... return () => {...}}, []) (componentWillUnmount)` →  언마운트 시 실행 (정리 작업)  
