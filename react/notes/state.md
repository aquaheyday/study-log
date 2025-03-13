# ⚛️ React 상태 관리 (State)

React에서 **State(상태)** 는 **컴포넌트 내부에서 관리되는 동적인 데이터**를 의미합니다.  
State를 사용하면 **사용자의 입력, API 데이터, UI 상태** 등을 관리할 수 있습니다.

---

## 1. State란?
- State는 **컴포넌트 내부에서 변경될 수 있는 값**입니다.
- `useState()` 훅을 사용하여 **함수형 컴포넌트에서 State를 관리**할 수 있습니다.
- State가 변경되면 **자동으로 UI가 다시 렌더링**됩니다.

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
✅ `useState(0)` → 초기 값 `0` 설정  
✅ `setCount(count + 1)` → 버튼 클릭 시 `count` 값 증가  

---

## 2. useState() 훅 사용법

### 2-1. 기본적인 State 사용
```jsx
import { useState } from "react";

function Example() {
  const [text, setText] = useState("초기 상태");

  return (
    <div>
      <p>{text}</p>
      <button onClick={() => setText("변경된 상태")}>변경</button>
    </div>
  );
}
```
✅ `useState("초기 상태")` → 초기 값 설정  
✅ `setText("변경된 상태")` → 버튼 클릭 시 State 변경  

---

### 2-2. 객체 형태의 State 관리
State가 **객체인 경우**, `setState()`를 사용할 때 **기존 값을 유지**해야 합니다.

```jsx
import { useState } from "react";

function UserProfile() {
  const [user, setUser] = useState({ name: "Alice", age: 25 });

  return (
    <div>
      <p>이름: {user.name}</p>
      <p>나이: {user.age}</p>
      <button onClick={() => setUser({ ...user, age: user.age + 1 })}>
        나이 증가
      </button>
    </div>
  );
}
```
✅ `setUser({ ...user, age: user.age + 1 })` → 기존 값 유지 후 `age`만 변경  

---

### 3-3. 배열 형태의 State 관리
State가 **배열인 경우**, `map()`, `filter()` 등을 사용하여 업데이트합니다.

```jsx
import { useState } from "react";

function TodoList() {
  const [todos, setTodos] = useState(["React 배우기", "JS 복습"]);

  const addTodo = () => {
    setTodos([...todos, "새로운 할 일"]);
  };

  return (
    <div>
      <ul>
        {todos.map((todo, index) => (
          <li key={index}>{todo}</li>
        ))}
      </ul>
      <button onClick={addTodo}>할 일 추가</button>
    </div>
  );
}
```
✅ `setTodos([...todos, "새로운 할 일"])` → 기존 배열 유지 후 새로운 요소 추가  

---

## 3. State 업데이트 주의사항

### ❌ 직접 변경하면 안 됨
State는 반드시 **`setState()` 함수를 사용하여 변경**해야 합니다.

```jsx
function Counter() {
  const [count, setCount] = useState(0);

  // 잘못된 방식 ❌
  const wrongUpdate = () => {
    count = count + 1; // 직접 수정하면 렌더링되지 않음
  };

  // 올바른 방식 ✅
  const correctUpdate = () => {
    setCount(count + 1);
  };

  return (
    <div>
      <p>카운트: {count}</p>
      <button onClick={wrongUpdate}>잘못된 업데이트</button>
      <button onClick={correctUpdate}>올바른 업데이트</button>
    </div>
  );
}
```
✅ `setCount(count + 1)`을 사용해야 UI가 자동 업데이트됨  

---

## 4. State 업데이트가 비동기적으로 작동함
State 업데이트는 **즉시 반영되지 않고, 비동기적으로 처리**됩니다.

```jsx
function Example() {
  const [count, setCount] = useState(0);

  const handleClick = () => {
    setCount(count + 1);
    console.log(count); // 이전 값이 출력됨
  };

  return (
    <div>
      <p>카운트: {count}</p>
      <button onClick={handleClick}>+1</button>
    </div>
  );
}
```
✅ `console.log(count)` → 이전 값이 출력됨  
✅ 최신 값을 반영하려면 **`setState(prev => prev + 1)` 방식 사용**

```jsx
setCount(prevCount => prevCount + 1);
```

---

## 5. 여러 개의 State 관리하기

컴포넌트 내에서 여러 개의 `useState()`를 사용할 수 있습니다.

```jsx
import { useState } from "react";

function MultiStateExample() {
  const [name, setName] = useState("Alice");
  const [age, setAge] = useState(25);

  return (
    <div>
      <p>이름: {name}</p>
      <p>나이: {age}</p>
      <button onClick={() => setAge(age + 1)}>나이 증가</button>
    </div>
  );
}
```
✅ 여러 개의 `useState()`를 사용하여 독립적인 상태 관리 가능  

---

## 6. State vs Props 차이점

| 특징 | State | Props |
|------|-------|-------|
| 변경 가능 여부 | ✅ 변경 가능 (`setState`) | ❌ 변경 불가 (읽기 전용) |
| 사용 대상 | 컴포넌트 내부 데이터 | 부모 → 자식 데이터 전달 |
| 초기값 설정 | `useState(initialValue)` | 부모 컴포넌트에서 전달 |
| 예제 사용 | `useState()` | `<Child name="Alice" />` |

✅ **State는 컴포넌트 내부에서 변경 가능**  
✅ **Props는 부모 컴포넌트에서 자식으로 전달**  

---

## 7. 공식 문서 및 추가 자료
- [React 공식 문서 - State](https://react.dev/reference/react/useState)
- [React 상태 관리 가이드](https://react.dev/learn/state-a-components-memory)

---

🚀 **React의 State를 이해했다면, 이제 이벤트 핸들링을 배워봅시다!**  
다음 개념: [React 이벤트 처리](./events.md) →
