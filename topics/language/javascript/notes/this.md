# 🔗 JavaScript - `this` 바인딩

JavaScript에서 `this`는 **함수가 호출될 때 결정되는 동적 바인딩**입니다.  
정의할 때가 아니라 **"호출될 때" 어떤 방식으로 불렸는지**에 따라 `this`가 결정됩니다.

---

## 1️⃣ 기본 바인딩 (전역/일반 함수 호출)

- **strict mode**가 아닌 일반 환경에서는 `this`는 `window` (전역 객체)
- **strict mode**에서는 `undefined`

```js
function showThis() {
  console.log(this);
}

showThis(); // 브라우저에선 window, strict mode에선 undefined
```

### 🔐 strict 모드 란?

'엄격한 모드' — JavaScript를 좀 더 "엄격하게" 실행하게 만들어서 실수나 오류가 발생할 수 있는 코드를 사전에 차단

#### 사용 방법

```js
'user strict';
```

파일 또는 함수의 맨 위에 선언하면 `strict` 모드가 활성화됨

#### 일반 vs strict moded

```js
// 일반 모드 
function foo() {
  x = 10; // 선언 안 했는데도 됨 (암묵적 전역)
  console.log(x);
}

foo(); // 10


// strict 모드
'use strict';

function foo() {
  x = 10; // ❌ ReferenceError: x is not defined
}

foo();


// 일반 모드 (this)
function show() {
  console.log(this); // window
}

show();


// strict 모드 (this)
'use strict';

function show() {
  console.log(this); // undefined
}

show();
```

✔ 일반 함수에서 this는 자동으로 전역 객체(window)를 참조하지만 strict mode에서는 자동 바인딩이 안 됨 → undefined

#### 어디서 사용하는 게 좋을까?

- 모던 자바스크립트는 기본적으로 strict 모드를 사용하는게 일반적
- 모듈(.mjs) 나 클래슨는자동으로 strict 모드가 활성화됨 

---

## 2️⃣ 객체의 메서드 호출 (암시적 바인딩)

```js
const user = {
  name: "Alice",
  sayHi() {
    console.log(this.name);
  }
};

user.sayHi(); // "Alice"
```

- `this`는 `.` 앞의 **객체 (`user`)** 를 참조

---

## 3️⃣ 명시적 바인딩 (`call`, `apply`, `bind`)

- `call`/`apply`는 즉시 실행
- `bind`는 `this`를 고정한 새 함수를 반환

```js
function greet() {
  console.log(`Hello, I'm ${this.name}`);
}

const person = { name: "Bob" };

greet.call(person);  // Hello, I'm Bob
greet.apply(person); // Hello, I'm Bob

const boundFn = greet.bind(person);
boundFn(); // Hello, I'm Bob
```

---

## 4️⃣ new 바인딩 (생성자 호출)

- `new`를 쓰면 `this`는 **새로 생성되는 객체**를 가리킴

```js
function Person(name) {
  this.name = name;
}

const p = new Person("Charlie");
console.log(p.name); // "Charlie"
```

---

## 5️⃣ 화살표 함수의 바인딩 (`Lexical this`)

```js
const obj = {
  name: "Arrow",
  sayHi: () => {
    console.log(this.name);
  }
};

obj.sayHi(); // undefined (this는 상위 스코프를 참조)
```

- 화살표 함수는 **자신만의 `this`를 갖지 않음**
- 외부(상위) 스코프의 `this`를 **"그대로 사용"**

```js
function Timer() {
  this.seconds = 0;

  setInterval(() => {
    this.seconds++;
    console.log(this.seconds);
  }, 1000);
}
```

- 위 예제에서 `this.seconds`는 `Timer` 인스턴스를 가리킴

---

## 6️⃣ 이벤트 핸들러에서의 `this`

### 일반 함수

```js
const btn = document.querySelector("button");
btn.addEventListener("click", function () {
  console.log(this); // 버튼 요소
});
```

### 화살표 함수

```js
btn.addEventListener("click", () => {
  console.log(this); // 상위 스코프의 this (대개 window or undefined)
});
```

---

## 🎯 정리

#### `this` 바인딩 규칙 요약표
| 호출 방식           | this가 가리키는 대상                          |
|----------------------|-----------------------------------------------|
| 일반 함수 호출       | `window` (non-strict), `undefined` (strict)   |
| 메서드 호출          | 해당 메서드를 호출한 **객체**                |
| `call`, `apply`, `bind` | **명시한 객체**                              |
| 생성자 함수 (`new`)  | 새로 생성된 객체                              |
| 화살표 함수          | **외부 스코프의 this**                        |
| DOM 이벤트 핸들러     | 일반 함수: 해당 요소 / 화살표 함수: 외부 this |

#### 헷갈릴 때는?

- **호출부를 보자!** → 함수가 어떻게 호출되었는지를 보면 `this`가 보인다
- **화살표 함수는 this 바인딩 안 됨** → 외부 스코프 따라간다
