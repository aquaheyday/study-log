# 🔄 React Context API

**Context API**는 **전역 상태(Global State)를 관리하는 React 내장 기능**입니다.  
부모 → 자식으로 **Props를 계속 전달하는 문제(Props Drilling)** 를 해결할 수 있습니다.

---

## 1. Context API란?

✅ **Props Drilling 문제 해결**  
✅ **Redux 같은 외부 상태 관리 라이브러리 없이 전역 상태 관리 가능**  
✅ **React 내장 기능으로 별도 설치 없이 사용 가능**  

---

## 2. Context API의 기본 개념

React에서는 **`createContext()`를 사용하여 전역 상태를 생성**합니다.  
🔹 **Provider** → 데이터를 제공하는 컴포넌트  
🔹 **Consumer** → 데이터를 사용하는 컴포넌트  

### 2-1. Props Drilling 문제 (기본 상태 전달 방식)
```jsx
function Parent() {
  const user = "Alice";
  return <Child user={user} />;
}

function Child({ user }) {
  return <GrandChild user={user} />;
}

function GrandChild({ user }) {
  return <p>사용자: {user}</p>;
}
```
✅ **Props를 계속 전달해야 함 → 유지보수 어려움**  

### 2-2. Context API 적용 후 (Props Drilling 해결)
```jsx
import { createContext, useContext } from "react";

const UserContext = createContext();

function Parent() {
  return (
    <UserContext.Provider value="Alice">
      <Child />
    </UserContext.Provider>
  );
}

function Child() {
  return <GrandChild />;
}

function GrandChild() {
  const user = useContext(UserContext);
  return <p>사용자: {user}</p>;
}
```
✅ **`useContext(UserContext)`로 직접 접근 가능 → Props 전달 불필요**  

---

## 3. Context API 기본 사용법

### 3-1. Context 생성 (`createContext`)
```jsx
import { createContext } from "react";

const ThemeContext = createContext("light"); // 기본값 설정
export default ThemeContext;
```
✅ `createContext(defaultValue)` → **기본값을 설정 가능**  

---

### 3-2. Provider로 상태 전달
```jsx
import ThemeContext from "./ThemeContext";

function App() {
  return (
    <ThemeContext.Provider value="dark">
      <Component />
    </ThemeContext.Provider>
  );
}
```
✅ `Provider`에서 **전역으로 상태를 제공**  

---

### 3-3. useContext()로 값 사용
```jsx
import { useContext } from "react";
import ThemeContext from "./ThemeContext";

function Component() {
  const theme = useContext(ThemeContext);
  return <p>현재 테마: {theme}</p>;
}
```
✅ `useContext(ThemeContext)`를 사용하여 값 가져오기  

---

## 4. 상태 업데이트 (전역 State 관리)

`useState()`를 활용하면 **Context에서 전역 상태를 업데이트 가능**합니다.

### 4-1. Context에 `useState()` 적용
```jsx
import { createContext, useState } from "react";

const ThemeContext = createContext();

export function ThemeProvider({ children }) {
  const [theme, setTheme] = useState("light");

  return (
    <ThemeContext.Provider value={{ theme, setTheme }}>
      {children}
    </ThemeContext.Provider>
  );
}

export default ThemeContext;
```
✅ `useState()`를 사용하여 상태를 관리  
✅ `value={{ theme, setTheme }}` → **상태와 변경 함수 전달 가능**  

---

### 4-2. Provider로 감싸기
```jsx
import { ThemeProvider } from "./ThemeContext";
import ThemeSwitcher from "./ThemeSwitcher";

function App() {
  return (
    <ThemeProvider>
      <ThemeSwitcher />
    </ThemeProvider>
  );
}
```
✅ `ThemeProvider`로 **전체 컴포넌트를 감싸서 전역 상태 사용 가능**  

---

### 4-3. Context에서 값 변경 (`setState()` 활용)
```jsx
import { useContext } from "react";
import ThemeContext from "./ThemeContext";

function ThemeSwitcher() {
  const { theme, setTheme } = useContext(ThemeContext);

  return (
    <div>
      <p>현재 테마: {theme}</p>
      <button onClick={() => setTheme(theme === "light" ? "dark" : "light")}>
        테마 변경
      </button>
    </div>
  );
}

export default ThemeSwitcher;
```
✅ `setTheme()`을 통해 **Context 내부의 상태 변경 가능**  

---

## 5. 여러 개의 Context 사용 (중첩 가능)

**여러 개의 Context를 함께 사용할 수 있음**  
```jsx
import { createContext, useContext } from "react";

const UserContext = createContext();
const ThemeContext = createContext();

function App() {
  return (
    <UserContext.Provider value="Alice">
      <ThemeContext.Provider value="dark">
        <Component />
      </ThemeContext.Provider>
    </UserContext.Provider>
  );
}

function Component() {
  const user = useContext(UserContext);
  const theme = useContext(ThemeContext);
  return <p>{user}님의 테마: {theme}</p>;
}
```
✅ `useContext(UserContext)`, `useContext(ThemeContext)`로 개별 접근 가능  

---

## 6. Context API vs Redux 비교

| 특징 | Context API | Redux |
|------|------------|-------|
| 상태 저장 위치 | `useContext + useState` | Redux Store |
| 사용 목적 | **간단한 전역 상태 관리** | **복잡한 상태 관리 (예: 대규모 앱)** |
| 데이터 흐름 | 단방향 | 단방향 (Flux 패턴) |
| 추가 라이브러리 | ❌ 필요 없음 | ✅ Redux 설치 필요 |
| 데이터 변경 방식 | `setState()` | `dispatch(action)` |

✅ **Context API는 소규모 전역 상태 관리에 적합**  
✅ **Redux는 복잡한 상태 관리가 필요한 경우 사용**  

---

## 7. Context API 사용 시 주의할 점

✅ **불필요한 리렌더링 주의**  
   → `useMemo()` 또는 `React.memo()`를 사용하여 최적화 가능  
✅ **Context 값이 자주 변경될 경우 Redux 고려**  
   → 상태 변경이 빈번하면 성능 저하 발생 가능  
✅ **중첩 Provider 피하기**  
   → 너무 많은 Context를 사용하면 코드 복잡도 증가  

---

## 8. 공식 문서 및 추가 자료
- [React 공식 문서 - Context API](https://react.dev/reference/react/createContext)
- [MDN - JavaScript 클로저와 Context](https://developer.mozilla.org/ko/docs/Web/JavaScript/Closures)

---

🚀 **React의 Context API를 이해했다면, 이제 Redux를 배워봅시다!**  
다음 개념: [Redux 기본](./redux-basics.md) →
