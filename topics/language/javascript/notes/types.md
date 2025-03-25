# 🧾 JavaScript - 데이터 타입

JavaScript는 동적 타입 언어입니다. → 변수에 저장되는 **값의 타입은 실행 중에 결정**되며, 언제든 변경될 수 있습니다.

---

## 1️⃣ 데이터 타입 분류

| 분류 | 종류 | 설명 |
|------|------|------|
| 기본(원시) 타입 | `string`, `number`, `boolean`, `undefined`, `null`, `symbol`, `bigint` | 값 자체를 저장 |
| 참조 타입 | `object`, `array`, `function`, `date`, ... | 참조 주소를 저장 |

---

## 2️⃣ 기본(원시) 타입

### 1) `string`

문자열을 표현 (작은따옴표, 큰따옴표, 백틱 사용 가능)

```js
let name = "Alice";
let greeting = `Hello, ${name}`;
```

---

### 2) `number`

정수, 실수, NaN, Infinity 등 모두 포함

```js
let a = 10;
let b = 3.14;
let c = Infinity;
let d = NaN; // Not a Number
```

---

### 3) `boolean`

참(`true`), 거짓(`false`)

```js
let isActive = true;
let isDeleted = false;
```

---

### 4) `undefined`

값이 할당되지 않은 변수의 기본 값

```js
let x;
console.log(x); // undefined
```

---

### 5) `null`

"값이 없음"을 명시적으로 나타냄

```js
let y = null;
```

---

### 6) `symbol` (ES6)

유일하고 변경 불가능한 값

```js
let sym1 = Symbol("id");
```

---

### 7) `bigint` (ES2020)

큰 정수를 안전하게 표현 (`n` 접미사 사용)

```js
let big = 1234567890123456789012345678901234567890n;
```

---

## 3️⃣ 참조(객체) 타입

### 1) `object`

키-값 쌍으로 이루어진 컬렉션

```js
let user = {
  name: "Tom",
  age: 30
};
```

---

### 2) `array`

순서가 있는 데이터 목록 (객체의 특수 형태)

```js
let fruits = ["apple", "banana", "cherry"];
```

---

### 3) `function`

함수도 객체의 한 종류

```js
function greet(name) {
  return "Hello " + name;
}
```

---

## 4️⃣ typeof 연산자

변수의 타입을 문자열로 반환

```js
typeof 123        // "number"
typeof "hello"    // "string"
typeof true       // "boolean"
typeof null       // ❗ "object" (자바스크립트의 오래된 버그)
typeof undefined  // "undefined"
typeof {}         // "object"
typeof []         // "object"
typeof function(){} // "function"
```

### typeof null == "object" ?

⚠️ 버그입니다. null은 원시 타입이지만, 오래된 구현상 "object"로 나옵니다.

---

## 🎯 정리

✔ JavaScript에는 **7가지 원시 타입**과 **참조 타입**이 존재  
✔ `null`, `undefined`의 차이 이해 중요  
✔ `typeof`를 통해 데이터 타입 확인 가능  
✔ 기본 타입은 **값 자체**, 참조 타입은 **메모리 주소**를 저장  
✔ `null`이 "object"로 나오는 건 JS의 유서 깊은 버그!

