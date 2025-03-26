# 🧱 JavaScript - 객체 & 프로토타입

자바스크립트에서 **객체(Object)** 는 데이터와 동작(메서드)을 저장하는 핵심 구조이며, 모든 객체는 **프로토타입(Prototype)** 을 통해 다른 객체로부터 속성과 메서드를 상속받습니다.

---

## 1️⃣ 객체(Object)란?

- **키(key) - 값(value)** 쌍으로 이루어진 데이터 집합
- 자바스크립트에서 거의 모든 것이 객체
- 배열, 함수도 객체의 일종

```js
const user = {
  name: "Alice",
  age: 25,
  greet: function() {
    console.log(`Hi, I'm ${this.name}`);
  }
};

user.greet(); // Hi, I'm Alice
```

---

## 2️⃣ 객체 생성 방법

### 1) 객체 리터럴

```js
const obj = { key: "value" };
```

### 2) 생성자 함수

```js
function Person(name) {
  this.name = name;
}

const p1 = new Person("Bob");
console.log(p1.name); // Bob
```

### 3) 클래스 문법 (ES6)

```js
class Animal {
  constructor(type) {
    this.type = type;
  }

  speak() {
    console.log(`${this.type} makes a sound.`);
  }
}

const dog = new Animal("Dog");
dog.speak(); // Dog makes a sound.
```

---

## 3️⃣ 프로토타입(Prototype)이란?

- 자바스크립트는 **프로토타입 기반 언어**
- 객체는 **자신의 부모 역할을 하는 객체**를 가리키는 **[[Prototype]]** 을 가지고 있음
- 이를 통해 **상속**이 일어남

```js
// 생성자 함수 정의
// 새로운 객체가 생성되고 그 객체의 this가 함수 내부에 바인딩됨 this.name = name → 생성된 객체에 name 속성 추가됨
// 예: new Person("Charlie") → { name: "Charlie" }라는 객체 생성됨
function Person(name) {
  this.name = name;
}

// 프로토타입 메서드 정의
// 생성자 함수인 Person에 sayHello 메서드를 추가
// Person으로 만들어지는 모든 객체가 공유하는 함수
Person.prototype.sayHello = function() {
  console.log(`Hello, I'm ${this.name}`);
};

// 즉, sayHello 함수는 메모리에 하나만 존재하고,
// user1.sayHello()를 실행하면 user1 → Person.prototype → sayHello 를 찾아서 실행함
const user1 = new Person("Charlie");
user1.sayHello(); // Hello, I'm Charlie
```

✔ 이 구조를 "프로토타입 체인"이라고 부름

---

## 4️⃣ 프로토타입 체인

- 객체 → 부모 객체 → 부모의 부모 ... → `Object.prototype` → `null`

```js
console.log(user1.toString()); 
// user1 → Person.prototype → Object.prototype → toString()
```

---

## 5️⃣ __proto__ vs prototype

| 구분 | 설명 |
|------|------|
| `__proto__` | 모든 객체가 가진 내부 참조, 실제 상속 연결 |
| `prototype` | 생성자 함수에만 존재, 인스턴스의 `__proto__`로 연결됨 |

```js
user1.__proto__ === Person.prototype; // true
```

---

## 6️⃣ Object.create()

- 특정 객체를 프로토타입으로 갖는 새 객체 생성

```js
const parent = {
  greet() {
    console.log("Hello from parent");
  }
};

const child = Object.create(parent);
child.greet(); // Hello from parent
```

---

## 7️⃣ 클래스와 프로토타입

- 클래스와 프로토타입은 완전히 다른 것이 아니라, 클래스는 프로토타입을 감싸는 문법적 설탕(syntax sugar) 입니다.
- 즉, 클래스 = 프로토타입 기반 상속을 쉽게 쓰게 해주는 문법

```js
class Car {
  drive() {
    console.log("Driving...");
  }
}

const myCar = new Car();
myCar.drive(); // Driving...

console.log(myCar.__proto__ === Car.prototype); // true
```

---

## 🎯 정리

✔ 객체는 key-value 형태로 데이터와 메서드를 저장  
✔ 객체는 프로토타입을 통해 다른 객체로부터 상속받음  
✔ 함수에는 prototype이 있고, 인스턴스의 `__proto__`는 이를 참조  
✔ 클래스 문법도 결국은 프로토타입 기반  
✔ 프로토타입 체인을 통해 상속과 재사용이 가능  

