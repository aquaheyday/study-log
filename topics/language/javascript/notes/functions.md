# 🔧 JavaScript 함수 선언 정리

JavaScript에서 **함수(Function)** 는 **재사용 가능한 코드 블록**입니다.  
입력(매개변수)을 받아 동작을 수행하고, 결과를 반환합니다.

---

## 1️⃣ 호이스팅이란?

선언이 코드의 맨 위로 끌어올려지는 것처럼 동작하는 자바스크립트의 특성입니다.

- 자바스크립트는 코드를 실행하기 전에 변수 선언과 함수 선언을 먼저 처리합니다.
- 이 과정에서 변수와 함수의 선언부만 위로 끌어올리는 것처럼 동작하게 되는것을 호이스팅 이라고합니다.

---

## 2️⃣ 함수 선언 방식

### 1) 함수 선언식 (Function Declaration)

- 가장 일반적인 선언 방식
- ✅ **호이스팅** → 코드 위에서 호출 가능

```js
function sayHello(name) {
  return `Hello, ${name}!`;
}

console.log(sayHello("Alice")); // Hello, Alice!
```

---

### 2) 함수 표현식 (Function Expression)

- 변수에 함수를 **값처럼 할당**
- ❌ **호이스팅** → 선언 이후에만 사용 가능

```js
const sayBye = function(name) {
  return `Bye, ${name}!`;
};

console.log(sayBye("Bob")); // Bye, Bob!
```

---

### 3) 화살표 함수 (Arrow Function, ES6)

- `this`를 **상속**받음 (일반 함수와 차이)
- 한 줄 리턴은 `{}` 생략 가능
- 콜백, 간단한 표현에 자주 사용됨
- ❌ **호이스팅** → 함수 자체는 호이스팅되지 않고, 변수 선언만 호이스팅되고 초기화는 안 되기 때문에 사용할 수 없음

```js
// 두 숫자 a, b를 받아 더한 값을 반환하는 함수
const add = (a, b) => a + b;

// 문자열 name을 받아 "Hi, 이름!" 형태로 인사하는 문자열을 반환하는 함수
const greet = (name) => {
  return `Hi, ${name}!`;
};
```

---

## 3️⃣ 매개변수 & 기본값

- 매개변수 생략 시 `undefined`
- `=`로 기본값 설정 가능

```js
function multiply(x, y = 1) {
  return x * y;
}

console.log(multiply(5));      // 5
console.log(multiply(5, 2));   // 10
```

---

## 4️⃣ 나머지 매개변수 (...rest)

- 전달된 모든 인자를 배열로 받음
- `...` 문법: 나머지 매개변수 또는 스프레드

```js
// 나머지 매개변수(rest parameter)와 배열의 reduce 함수를 활용한 가변 인자 합계 함수
function sumAll(...numbers) {
  return numbers.reduce((a, b) => a + b, 0);
  // reduce는 배열을 순회하면서 값을 누적합니다.
  // (a, b) => a + b: 두 값을 더해서 누적
  // 0: 누적의 초기값
}

console.log(sumAll(1, 2, 3, 4)); // 10
```

---

## 5️⃣ 콜백 함수

함수를 **다른 함수의 인자로 전달**해서 나중에 실행

```js
function greetUser(callback) {
  const name = "Alice";
  callback(name);
}

greetUser(function(name) {
  console.log("Hi,", name);
});
```

---

## 6️⃣ 즉시 실행 함수 (IIFE, Immediately Invoked Function Expression)

정의하자마자 바로 실행되는 함수

```js
(function() {
  console.log("즉시 실행!");
})();
```

---

## 🎯 정리

✔ 함수는 선언식, 표현식, 화살표 형태로 정의 가능  
✔ 화살표 함수는 간결하지만 `this`가 다르게 동작  
✔ 매개변수 기본값, 나머지 매개변수, 콜백 함수 등 다양하게 활용 가능  
✔ 자바스크립트에서 함수는 **1급 객체** → 변수에 저장, 인자로 전달, 리턴 가능

