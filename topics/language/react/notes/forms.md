# 🔲 React 폼과 입력 관리

React에서 폼(Form) 요소는 HTML과 유사하지만,  
입력값을 관리하고 상태(State)를 활용하는 방식이 다릅니다.  

React에서는 **Controlled Component (제어 컴포넌트)** 를 사용하여  
입력 값을 `state`로 관리하고 동적으로 업데이트할 수 있습니다.

---

## 1️⃣ HTML 폼과의 차이점

| 특징 | HTML (기본) | React (Controlled) |
|------|------------|-------------------|
| 데이터 저장 위치 | DOM 요소 (`input.value`) | 컴포넌트 `state` |
| 값 변경 방식 | 사용자가 직접 입력 | `onChange` 핸들러를 통해 `setState()` |
| 데이터 동기화 | 폼 요소에 직접 저장 | `state`와 동기화 |

✔ React에서는 `useState()`를 사용하여 폼 입력값을 관리해야 함  

---

## 2️⃣ Controlled Component (제어 컴포넌트)

React에서는 **폼 요소의 상태를 `state`로 관리**하는 방식을 사용합니다.

### 1) 단일 입력 필드 관리
```jsx
import { useState } from "react";

function TextInput() {
  const [text, setText] = useState("");

  const handleChange = (event) => {
    setText(event.target.value);
  };

  return (
    <div>
      <input type="text" value={text} onChange={handleChange} />
      <p>입력한 값: {text}</p>
    </div>
  );
}

export default TextInput;
```
✔ `value={text}` → 입력 값이 `state`와 동기화됨  
✔ `onChange={handleChange}` → 입력 시 `state` 업데이트  

---

### 2) 여러 개의 입력 필드 관리
여러 개의 `input` 필드를 관리할 경우, `name` 속성을 활용합니다.

```jsx
import { useState } from "react";

function UserForm() {
  const [form, setForm] = useState({ name: "", email: "" });

  const handleChange = (event) => {
    setForm({ ...form, [event.target.name]: event.target.value });
  };

  return (
    <div>
      <input type="text" name="name" value={form.name} onChange={handleChange} placeholder="이름" />
      <input type="email" name="email" value={form.email} onChange={handleChange} placeholder="이메일" />
      <p>이름: {form.name}</p>
      <p>이메일: {form.email}</p>
    </div>
  );
}

export default UserForm;
```
✔ `name` 속성을 활용하여 `setState()`로 여러 필드 동시 관리 가능  

---

### 3) 체크박스 (Checkbox) 입력 관리
```jsx
import { useState } from "react";

function CheckboxExample() {
  const [isChecked, setIsChecked] = useState(false);

  const handleChange = (event) => {
    setIsChecked(event.target.checked);
  };

  return (
    <div>
      <input type="checkbox" checked={isChecked} onChange={handleChange} />
      <p>{isChecked ? "체크됨 ✅" : "체크 안됨 ❌"}</p>
    </div>
  );
}

export default CheckboxExample;
```
✔ `checked={isChecked}` → 체크 여부를 `state`와 동기화  

---

### 4) 라디오 버튼 (Radio) 입력 관리
```jsx
import { useState } from "react";

function RadioExample() {
  const [gender, setGender] = useState("male");

  const handleChange = (event) => {
    setGender(event.target.value);
  };

  return (
    <div>
      <label>
        <input type="radio" name="gender" value="male" checked={gender === "male"} onChange={handleChange} />
        남성
      </label>
      <label>
        <input type="radio" name="gender" value="female" checked={gender === "female"} onChange={handleChange} />
        여성
      </label>
      <p>선택한 성별: {gender}</p>
    </div>
  );
}

export default RadioExample;
```
✔ `checked={gender === "male"}` → 선택된 값과 비교하여 상태 반영  

---

### 5) 드롭다운 (Select) 입력 관리
```jsx
import { useState } from "react";

function SelectExample() {
  const [fruit, setFruit] = useState("apple");

  const handleChange = (event) => {
    setFruit(event.target.value);
  };

  return (
    <div>
      <select value={fruit} onChange={handleChange}>
        <option value="apple">사과</option>
        <option value="banana">바나나</option>
        <option value="cherry">체리</option>
      </select>
      <p>선택한 과일: {fruit}</p>
    </div>
  );
}

export default SelectExample;
```
✔ `value={fruit}` → 선택된 값을 `state`와 동기화  

---

## 3️⃣ Form 제출 이벤트 (`onSubmit`)

React에서는 `onSubmit`을 사용하여 **폼 제출 이벤트를 처리**할 수 있습니다.

```jsx
import { useState } from "react";

function FormSubmit() {
  const [name, setName] = useState("");

  const handleSubmit = (event) => {
    event.preventDefault(); // 기본 폼 제출 방지
    alert(`제출된 이름: ${name}`);
  };

  return (
    <form onSubmit={handleSubmit}>
      <input type="text" value={name} onChange={(e) => setName(e.target.value)} placeholder="이름 입력" />
      <button type="submit">제출</button>
    </form>
  );
}

export default FormSubmit;
```
✔ `event.preventDefault()` → 기본 제출 동작 방지  
✔ `onSubmit={handleSubmit}` → 폼 제출 시 실행  

---

## 4️⃣ Uncontrolled Component (비제어 컴포넌트)
폼 요소의 값을 **React `state`가 아닌 `ref`를 통해 직접 제어**하는 방식입니다.

```jsx
import { useRef } from "react";

function UncontrolledForm() {
  const inputRef = useRef(null);

  const handleSubmit = (event) => {
    event.preventDefault();
    alert(`입력된 값: ${inputRef.current.value}`);
  };

  return (
    <form onSubmit={handleSubmit}>
      <input type="text" ref={inputRef} placeholder="이름 입력" />
      <button type="submit">제출</button>
    </form>
  );
}

export default UncontrolledForm;
```
✔ `useRef()`를 사용하여 `input`의 값을 직접 가져올 수 있음  

---

## 🎯 정리
✔ React에서 폼 입력값을 관리하는 기본 방법은 Controlled Component (제어 컴포넌트) 사용  
✔ `useState()`를 활용하여 입력값을 상태(state)로 저장 및 관리  
✔ `onChange` 핸들러를 사용하여 입력 변경 시 state 업데이트  
✔ 여러 개의 입력 필드는 객체(state)와 `name` 속성을 활용하여 관리 가능  
✔ `onSubmit`을 사용하여 폼 제출 처리, `event.preventDefault()`로 기본 동작 방지  
✔ Uncontrolled Component (비제어 컴포넌트) 에서는 `useRef()`를 활용하여 값 접근 가능  
