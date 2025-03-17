# 📌 JSX 문법

JSX(JavaScript XML)는 **JavaScript 코드에서 XML 형식의 문법을 사용할 수 있도록 지원하는 React의 문법 확장**입니다.  

---

## 1. JSX란 무엇인가?
- JSX는 JavaScript 코드 내에서 HTML과 유사한 구문을 작성하는 방법입니다.
- 브라우저에서 직접 실행할 수 없으며, **Babel이 JSX를 순수 JavaScript로 변환**합니다.

#### JSX 예제 (React 컴포넌트)
```jsx
function Greeting() {
  return <h1>Hellow, React!</h1>;
}
```

#### 변환된 JavaScript 코드
```js
function Greeting() {
  return React.createElement("h1", null, "Hellow, React!");
}
```

---

## 2. JSX 기본 문법

### 반드시 하나의 부모 요소로 감싸야 함
JSX에서는 **하나의 부모 요소**로 감싸지 않으면 오류가 발생합니다.

✅ **올바른 JSX 코드**
```jsx
function App() {
  return (
    <div>
      <h1>안녕하세요</h1>
      <p>React JSX 문법을 배우고 있습니다.</p>
    </div>
  );
}
```

❌ **잘못된 JSX 코드 (에러 발생)**
```jsx
function App() {
  return (
    <h1>안녕하세요</h1>
    <p>React JSX 문법을 배우고 있습니다.</p>
  );
}
```
> **해결 방법:** `<> ... </>` **(Fragment 태그 사용 가능)**  
```jsx
function App() {
  return (
    <>
      <h1>안녕하세요</h1>
      <p>React JSX 문법을 배우고 있습니다.</p>
    </>
  );
}
```

---

### JavaScript 표현식 사용 `{}`  
JSX 내부에서는 `{}`를 사용하여 **JavaScript 표현식**을 삽입할 수 있습니다.

```jsx
function Greeting() {
  const name = "Alice";
  return <h1>안녕하세요, {name}!</h1>;
}
```

```jsx
function MathExample() {
  return <p>1 + 1 = {1 + 1}</p>;
}
```

---

### 인라인 스타일 적용  
JSX에서 **스타일을 적용할 때** 객체 형태로 작성해야 합니다.

```jsx
const textStyle = {
  color: "blue",
  fontSize: "20px",
  fontWeight: "bold",
};

function StyledText() {
  return <p style={textStyle}>이 문장은 파란색입니다.</p>;
}
```

---

### class 대신 className 사용  
JSX에서는 `class` 대신 **`className`을 사용**해야 합니다.

```jsx
function Button() {
  return <button className="btn-primary">클릭하세요</button>;
}
```

> **이유:** `class`는 JavaScript에서 예약어이기 때문

---

### 조건부 렌더링 (`&&`, `? :`, 삼항 연산자)  
JSX에서 **조건부 렌더링**을 할 때 `&&`, 삼항 연산자(`? :`)를 사용할 수 있습니다.

#### 논리 연산자 (`&&`)
```jsx
function ShowMessage({ isVisible }) {
  return <>{isVisible && <p>이 메시지는 보입니다!</p>}</>;
}
```

#### 삼항 연산자 (`? :`)
```jsx
function Greeting({ isLoggedIn }) {
  return <h1>{isLoggedIn ? "환영합니다!" : "로그인해주세요."}</h1>;
}
```

---

### 배열을 사용한 리스트 렌더링  
JSX에서는 **배열을 사용하여 여러 개의 요소를 렌더링**할 수 있습니다.

```jsx
function ItemList() {
  const items = ["사과", "바나나", "체리"];
  return (
    <ul>
      {items.map((item, index) => (
        <li key={index}>{item}</li>
      ))}
    </ul>
  );
}
```

> **주의:** 리스트를 렌더링할 때 **반드시 `key` 속성을 설정**해야 성능 최적화 가능

---

### JSX에서 함수 호출  
JSX 내부에서 **함수를 호출할 수도 있습니다.**

```jsx
function getMessage() {
  return "반갑습니다!";
}

function Greeting() {
  return <h1>{getMessage()}</h1>;
}
```

---

## 3. JSX의 장점  
✔ **가독성이 뛰어남**: HTML과 유사한 문법으로 직관적인 코드 작성 가능  
✔ **빠른 렌더링**: Babel을 통해 최적화된 JavaScript로 변환  
✔ **강력한 표현력**: JavaScript 표현식과 동적으로 결합 가능  

---

## 4. JSX 없이 React 사용 (일반 JavaScript)  
JSX를 사용하지 않고 React를 작성할 수도 있지만, **코드가 길고 가독성이 떨어짐**  

```js
const element = React.createElement("h1", null, "Hello, React!");
ReactDOM.render(element, document.getElementById("root"));
```

✔ **JSX를 사용하면 더 직관적으로 작성 가능**
```jsx
const element = <h1>Hello, React!</h1>;
ReactDOM.render(element, document.getElementById("root"));
```

---

## 5. 공식 문서 및 추가 자료
- [React 공식 문서 - JSX](https://react.dev/learn/writing-markup-with-jsx)
- [MDN - JavaScript 표현식](https://developer.mozilla.org/ko/docs/Web/JavaScript/Reference/Operators)
- [Babel JSX 변환기](https://babeljs.io/repl)  
