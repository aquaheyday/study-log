# 📌 컴포넌트 기초

React에서 **컴포넌트(Component)** 는 UI를 구성하는 기본 단위입니다.  
컴포넌트를 사용하면 **재사용성, 유지보수성, 가독성이 뛰어난 코드**를 작성할 수 있습니다.

---

## 1. 컴포넌트란?
- **컴포넌트(Component)** 는 React에서 **독립적인 UI 블록**입니다.
- 하나의 애플리케이션은 여러 개의 컴포넌트로 구성됩니다.
- 컴포넌트는 **재사용 가능**하며, **props를 통해 데이터를 전달**할 수 있습니다.

---

## 2. 컴포넌트의 종류
React에서는 **함수형 컴포넌트**와 **클래스형 컴포넌트**를 사용할 수 있습니다.  
(최근에는 **함수형 컴포넌트 + React Hooks** 사용을 권장합니다.)

---

### 함수형 컴포넌트 (Functional Component)
- **더 짧고 간결한 코드** 작성 가능
- **React Hooks (`useState`, `useEffect` 등) 사용 가능**
- **클래스형보다 성능이 우수**하고 메모리 사용량이 적음

```jsx
function Greeting() {
  return <h1>안녕하세요, React!</h1>;
}

export default Greeting;
```
✔ `Greeting` 컴포넌트는 `<h1>` 요소를 반환합니다.

---

### 클래스형 컴포넌트 (Class Component)
- `class` 키워드를 사용하여 정의  
- `render()` 메서드가 반드시 필요함  
- `this.state`를 사용하여 **상태(State)** 관리 가능

```jsx
import React, { Component } from "react";

class Greeting extends Component {
  render() {
    return <h1>안녕하세요, React!</h1>;
  }
}

export default Greeting;
```

> ✔ React 16.8 이후부터는 **함수형 컴포넌트와 Hooks**를 주로 사용하며, **클래스형 컴포넌트는 거의 사용되지 않습니다.**

---

## 3. 컴포넌트 사용법 (Import & Render)
컴포넌트를 사용하려면 `import`하여 렌더링해야 합니다.

### 기본 컴포넌트 불러오기
```jsx
import Greeting from "./Greeting";

function App() {
  return (
    <div>
      <Greeting />
    </div>
  );
}

export default App;
```
✔ `<Greeting />` 컴포넌트를 **`App` 컴포넌트에서 사용**

---

### 여러 개의 컴포넌트 조합
컴포넌트는 **다른 컴포넌트 안에서 사용 가능**합니다.

```jsx
function Header() {
  return <h1>My Website</h1>;
}

function Footer() {
  return <footer>© 2024 My Website</footer>;
}

function App() {
  return (
    <div>
      <Header />
      <p>환영합니다!</p>
      <Footer />
    </div>
  );
}

export default App;
```

✔ `Header`, `Footer` 컴포넌트를 조합하여 `App` 컴포넌트 구성

---

## 4. Props (컴포넌트 간 데이터 전달)
`props`는 **부모 컴포넌트에서 자식 컴포넌트로 데이터를 전달하는 방법**입니다.

### Props 기본 사용법
```jsx
function User(props) {
  return <h1>안녕하세요, {props.name}님!</h1>;
}

function App() {
  return <User name="Alice" />;
}

export default App;
```
✔ `User` 컴포넌트에 `name` 값을 전달 → `"안녕하세요, Alice님!"` 출력

---

### Props 기본값 설정 (`defaultProps`)
```jsx
function User(props) {
  return <h1>안녕하세요, {props.name}님!</h1>;
}

User.defaultProps = {
  name: "게스트",
};

export default User;
```
✔ `name` 값이 전달되지 않으면 **기본값("게스트")** 을 사용

---

### Props 비구조화 할당 (Destructuring)
```jsx
function User({ name }) {
  return <h1>안녕하세요, {name}님!</h1>;
}
```
✔ `props`를 **비구조화 할당**하여 더 깔끔한 코드 작성 가능

---

## 5. State (컴포넌트 내부 데이터 관리)
`State`는 **컴포넌트 내부에서 관리되는 동적인 데이터**입니다.  
React에서는 `useState()` 훅을 사용하여 State를 관리할 수 있습니다.

### `useState()`를 사용한 State 관리
```jsx
import { useState } from "react";

function Counter() {
  const [count, setCount] = useState(0);

  return (
    <div>
      <p>카운트: {count}</p>
      <button onClick={() => setCount(count + 1)}>+1</button>
    </div>
  );
}

export default Counter;
```
✔ `useState(0)` → 초기 값 `0` 설정  
✔ `setCount(count + 1)` → 버튼 클릭 시 `count` 값 증가  

---

## 6. 컴포넌트의 주요 특징
✔ 재사용성 → 동일한 컴포넌트를 여러 곳에서 재사용 가능  
✔ 독립성 → 각 컴포넌트는 독립적으로 동작  
✔ 조합 가능 → 여러 컴포넌트를 조합하여 복잡한 UI 구성 가능  
✔ 단방향 데이터 흐름 → `props`를 사용하여 부모 → 자식으로 데이터 전달  

---

## 7. 공식 문서 및 추가 자료
- [React 공식 문서 - Components](https://react.dev/learn/your-first-component)
- [React Hooks - useState](https://react.dev/reference/react/useState)
