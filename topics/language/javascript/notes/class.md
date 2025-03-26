# 🏗 JavaScript - 클래스 & 상속

자바스크립트에서 클래스는 객체지향 프로그래밍(OOP)을 쉽게 구현할 수 있는 문법입니다.  
클래스를 사용하면 **생성자 함수 + 프로토타입** 기반 코드를 더 직관적으로 쓸 수 있습니다.

---

## 1️⃣ 클래스 정의 (ES6 도입)

- `constructor()` 는 인스턴스를 초기화하는 함수
- `sayHello()` 는 프로토타입 메서드로 생성됨 (모든 인스턴스가 공유)

```js
class Person {
  constructor(name) {
    this.name = name;
  }

  sayHello() {
    console.log(`Hi, I'm ${this.name}`);
  }
}

const user = new Person("Alice");
user.sayHello(); // Hi, I'm Alice
```

---

## 2️⃣ 클래스 = 함수

- 클래스는 내부적으로 여전히 **함수 기반 + 프로토타입 구조**
- `class` 문법은 **프로토타입 기반 상속의 문법적 설탕(Syntax Sugar)**

```js
console.log(typeof Person); // "function"
console.log(Person.prototype.sayHello); // sayHello 함수 있음
```

---

## 3️⃣ 클래스 상속: `extends`

- `extends` 키워드로 부모 클래스를 상속
- 자식 클래스에서 같은 메서드를 정의하면 **오버라이딩(덮어쓰기)** 됨

```js
class Animal {
  constructor(name) {
    this.name = name;
  }

  speak() {
    console.log(`${this.name} makes a sound`);
  }
}

class Dog extends Animal {
  speak() {
    console.log(`${this.name} barks`);
  }
}

const dog = new Dog("Buddy");
dog.speak(); // Buddy barks
```

---

## 4️⃣ `super` 키워드

- 부모 클래스의 생성자나 메서드에 접근할 때 사용

```js
class Cat extends Animal {
  constructor(name, color) {
    super(name); // 부모의 constructor 호출
    this.color = color;
  }

  speak() {
    super.speak(); // 부모의 speak() 호출
    console.log(`${this.name} meows`);
  }
}

const cat = new Cat("Kitty", "white");
cat.speak();
// Kitty makes a sound
// Kitty meows
```

---

## 5️⃣ 클래스 필드와 정적 메서드

### 1) 클래스 필드 (속성)

```js
class Counter {
  count = 0; // 클래스 필드 (public field)

  increase() {
    this.count++;
    console.log(this.count);
  }
}
```

### 2) 정적(`static`) 메서드

- `static` 메서드는 클래스 이름으로 직접 호출하는 메서드이며, 인스턴스를 생성하지 않아도 사용할 수 있습니다.

```js
class MathUtil {
  static add(x, y) {
    return x + y;
  }
}

console.log(MathUtil.add(2, 3)); // 5
```

#### ✅ 호출 방식

```js
MathUtil.add(2, 3);
```

#### ❌ 호출 방식

```js
const m = new MathUtil();
m.add(2, 3); // ❌ TypeError: m.add is not a function
```

---

## 6️⃣ instanceof 연산자

- 객체가 어떤 클래스(또는 부모 클래스)에서 생성되었는지 확인할 수 있음

```js
console.log(dog instanceof Dog);     // true
console.log(dog instanceof Animal);  // true
console.log(dog instanceof Object);  // true
```

---

## 7️⃣ 클래스 내부의 `this`

- 클래스의 메서드 안에서 `this`는 **해당 인스턴스**를 가리킴
- 주의: 메서드를 꺼내서 단독 실행하면 `this`가 사라질 수 있음 (bind 필요)

```js
const fn = dog.speak;
fn(); // ❌ this가 undefined일 수 있음
```

---

## 🎯 정리

| 개념 | 설명 |
|------|------|
| `class` | 객체 생성 & 상속을 위한 문법 |
| `constructor` | 인스턴스 초기화 메서드 |
| `extends` | 클래스 상속 |
| `super` | 부모 클래스의 생성자/메서드 호출 |
| `static` | 클래스 자체에서 호출하는 메서드 |
| `this` | 클래스 내부에서는 인스턴스를 가리킴 |

#### 클래스도 결국 프로토타입 기반

```js
class A {}
const a = new A();

console.log(a.__proto__ === A.prototype); // true
```

✔ 즉, 클래스도 내부적으로는 기존의 **프로토타입 상속**을 활용하고 있습니다.

