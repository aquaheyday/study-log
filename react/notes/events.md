# 🔲 이벤트 처리

React에서는 **DOM 이벤트를 다루기 위해 이벤트 핸들러를 설정**할 수 있습니다.  
JSX 문법을 사용하여 **HTML과 유사한 방식**으로 이벤트를 바인딩할 수 있지만, 일부 차이점이 있습니다.

---

## 1. React 이벤트의 특징

### CamelCase 문법 사용
`onclick` → `onClick`, `onchange` → `onChange`

### 이벤트 핸들러에 함수 전달 (문자열 X)
HTML: `<button onclick="handleClick()">클릭</button>` ❌  
React: `<button onClick={handleClick}>클릭</button>` ✅

### SyntheticEvent 사용
React는 브라우저의 네이티브 이벤트를 감싸서 `SyntheticEvent` 객체로 제공  

✔ 이벤트 핸들러 내부에서 `this` 문제 해결 필요 (클래스형 컴포넌트의 경우)  

---

## 2. 기본 이벤트 핸들링

### 함수형 컴포넌트에서 이벤트 처리
```jsx
function ButtonClick() {
  const handleClick = () => {
    alert("버튼이 클릭되었습니다!");
  };

  return <button onClick={handleClick}>클릭</button>;
}

export default ButtonClick;
```
✔ `onClick={handleClick}` → 함수 이름을 직접 전달해야 함 (괄호 없음)  

---

### 클래스형 컴포넌트에서 이벤트 처리
```jsx
import React, { Component } from "react";

class ButtonClick extends Component {
  handleClick() {
    alert("버튼이 클릭되었습니다!");
  }

  render() {
    return <button onClick={this.handleClick}>클릭</button>;
  }
}

export default ButtonClick;
```
✔ 클래스형 컴포넌트에서는 `this.handleClick`을 직접 바인딩  

---

## 3. 이벤트 객체 사용 (`event`)

이벤트 핸들러에서 **이벤트 객체를 활용**할 수 있습니다.

```jsx
function InputField() {
  const handleChange = (event) => {
    console.log("입력값:", event.target.value);
  };

  return <input type="text" onChange={handleChange} />;
}
```
✔ `event.target.value` → 입력된 값을 가져올 수 있음  

---

## 4. 이벤트 핸들러에 매개변수 전달

매개변수를 전달할 경우 화살표 함수 또는 `bind()`를 사용해야 합니다.

### 함수형 컴포넌트에서 매개변수 전달
```jsx
function GreetingButton() {
  const sayHello = (name) => {
    alert(`안녕하세요, ${name}님!`);
  };

  return <button onClick={() => sayHello("Alice")}>클릭</button>;
}
```
✔ 화살표 함수 사용 → `onClick={() => sayHello("Alice")}`  

---

### 클래스형 컴포넌트에서 매개변수 전달
```jsx
import React, { Component } from "react";

class GreetingButton extends Component {
  sayHello(name) {
    alert(`안녕하세요, ${name}님!`);
  }

  render() {
    return <button onClick={() => this.sayHello("Alice")}>클릭</button>;
  }
}

export default GreetingButton;
```
✔ 화살표 함수 사용하여 매개변수 전달 (`onClick={() => this.sayHello("Alice")}`)  

---

## 5. `this` 바인딩 문제 (클래스형 컴포넌트)

클래스형 컴포넌트에서는 `this`를 명시적으로 바인딩해야 합니다.

```jsx
import React, { Component } from "react";

class Counter extends Component {
  constructor(props) {
    super(props);
    this.state = { count: 0 };
    this.handleClick = this.handleClick.bind(this); // this 바인딩
  }

  handleClick() {
    this.setState({ count: this.state.count + 1 });
  }

  render() {
    return (
      <div>
        <p>카운트: {this.state.count}</p>
        <button onClick={this.handleClick}>+1</button>
      </div>
    );
  }
}

export default Counter;
```
✔ `this.handleClick = this.handleClick.bind(this);` → 생성자에서 `this` 바인딩 필요  
✔ 클래스형 컴포넌트에서는 화살표 함수 (`handleClick = () => {}`)를 사용하면 `this` 바인딩이 필요 없음!  

```jsx
class Counter extends Component {
  state = { count: 0 };

  handleClick = () => {
    this.setState({ count: this.state.count + 1 });
  };

  render() {
    return (
      <div>
        <p>카운트: {this.state.count}</p>
        <button onClick={this.handleClick}>+1</button>
      </div>
    );
  }
}
```

---

## 6. 기본 이벤트 방지 (`preventDefault`)

이벤트 기본 동작을 막으려면 `event.preventDefault()`를 사용합니다.

```jsx
function Form() {
  const handleSubmit = (event) => {
    event.preventDefault(); // 기본 제출 방지
    alert("폼이 제출되었습니다.");
  };

  return (
    <form onSubmit={handleSubmit}>
      <button type="submit">제출</button>
    </form>
  );
}
```
✔ `event.preventDefault()` → 폼 제출 기본 동작 방지  

---

## 7. 이벤트 버블링 중지 (`stopPropagation`)

이벤트 버블링을 막으려면 `event.stopPropagation()`을 사용합니다.

```jsx
function EventExample() {
  const handleParentClick = () => alert("부모 div 클릭됨");
  const handleChildClick = (event) => {
    event.stopPropagation(); // 부모로의 이벤트 전파 방지
    alert("자식 버튼 클릭됨");
  };

  return (
    <div onClick={handleParentClick} style={{ padding: "20px", background: "#ddd" }}>
      <button onClick={handleChildClick}>버튼</button>
    </div>
  );
}
```
✔ `event.stopPropagation()` → 부모의 `onClick` 이벤트 실행 방지  

---

## 8. 여러 이벤트 핸들러 사용 (onMouseEnter, onFocus 등)

React에서는 다양한 이벤트를 지원합니다.

```jsx
function MouseEventExample() {
  const handleMouseEnter = () => console.log("마우스가 올라갔습니다.");
  const handleFocus = () => console.log("입력 필드가 포커스를 얻었습니다.");

  return (
    <div>
      <button onMouseEnter={handleMouseEnter}>마우스 올려보기</button>
      <input type="text" onFocus={handleFocus} placeholder="클릭하세요" />
    </div>
  );
}
```
✔ `onMouseEnter`, `onFocus` 등 다양한 이벤트 처리 가능  

---

## 🎯 정리
✔ React에서 이벤트 핸들링은 JSX 문법을 활용하여 설정  
✔ `onClick`, `onChange` 등 CamelCase 이벤트 속성 사용  
✔ 클래스형 컴포넌트에서는 this 바인딩 문제 해결 필요  
✔ `event.preventDefault()` → 기본 이벤트 방지  
✔ `event.stopPropagation()` → 이벤트 버블링 방지  
✔ 다양한 이벤트(`onMouseEnter`, `onFocus` 등) 활용 가능  
