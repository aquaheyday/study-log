# 📦 JavaScript - Promise

**Promise(프로미스)** 는 자바스크립트에서 **비동기 작업의 성공/실패 결과를 표현**하는 객체입니다.  
콜백 지옥을 피하고, 더 깔끔한 비동기 흐름 제어를 위해 사용됩니다.

---

## 1️⃣ `Promise`란?

**"미래에 완료될 수도 있는 비동기 작업의 결과"** 를 나타내는 객체

#### `Promise`는 3가지 상태를 가집니다:

| 상태        | 설명                                |
|-------------|-------------------------------------|
| `pending`   | 대기 중 (아직 완료되지 않음)         |
| `fulfilled` | 작업 성공 (`resolve` 호출됨)        |
| `rejected`  | 작업 실패 (`reject` 호출됨)         |

---

## 2️⃣ `Promise` 기본 구조

```js
const promise = new Promise((resolve, reject) => {
  // 비동기 작업 수행
  if (성공) {
    resolve("성공 결과");
  } else {
    reject("에러 메시지");
  }
});
```

✔ `resolve(value)` → 성공 상태로 변경  
✔ `reject(error)` → 실패 상태로 변경  

---

## 3️⃣ `Promise` 사용 예

```js
const promise = new Promise((resolve, reject) => {
  setTimeout(() => {
    resolve("Done!");
  }, 1000);
});

promise
  .then(result => {
    console.log("성공:", result);
  })
  .catch(error => {
    console.error("실패:", error);
  });
```

---

## 4️⃣ `Promise` 체이닝 (`then` 연결)

- `then()`은 **성공했을 때 실행되는 콜백**
- `catch()`는 **중간에 실패한 경우 에러 처리**
- 에러가 발생하면 아래 `then()`들은 건너뜀

```js
doSomething()
  .then(result => doNext(result))
  .then(nextResult => doFinal(nextResult))
  .then(finalResult => {
    console.log("완료:", finalResult);
  })
  .catch(err => {
    console.error("에러 발생:", err);
  });
```

---

## 5️⃣ `Promise` 에러 처리

- `reject()`나 예외 발생 시 `catch()`로 이동

```js
new Promise((resolve, reject) => {
  throw new Error("문제 발생!");
})
  .then(() => {
    console.log("성공");
  })
  .catch(err => {
    console.error("에러:", err.message);
  });
```

---

## 6️⃣ `Promise` 유틸 메서드

### 1) `Promise.resolve(value)`

- 이미 성공한 Promise를 만듦

```js
Promise.resolve(42).then(console.log); // 42
```

---

### 2. `Promise.reject(error)`

- 실패한 Promise를 만듦

```js
Promise.reject("에러").catch(console.error); // 에러
```

---

### 3. `Promise.all([ ... ])`

- 여러 개의 Promise가 **모두 성공해야** 다음 단계 실행

```js
Promise.all([p1, p2])
  .then(results => console.log(results))
  .catch(err => console.error(err));
```

---

### 4. `Promise.race([ ... ])`

- 가장 먼저 끝나는 Promise의 결과 반환

```js
Promise.race([p1, p2])
  .then(result => console.log(result))
  .catch(err => console.error(err));
```

---

## 7️⃣ `async`/`await`와 함께 사용

- `await`는 `Promise`가 해결될 때까지 기다림
- `try/catch`로 에러 처리

```js
async function fetchData() {
  try {
    const data = await getData(); // getData는 Promise 반환
    console.log(data);
  } catch (err) {
    console.error("에러:", err);
  }
}
```

---

## 🎯 정리

| 개념         | 설명 |
|--------------|------|
| `Promise`     | 비동기 결과를 나타내는 객체 |
| `resolve`     | 성공 시 호출 (fulfilled) |
| `reject`      | 실패 시 호출 (rejected) |
| `then()`      | 성공 시 실행되는 함수 |
| `catch()`     | 실패 시 실행되는 함수 |
| `finally()`   | 성공/실패와 상관없이 마지막에 실행 |
| `Promise.all` | 모든 Promise 성공 시 실행 |
| `Promise.race`| 가장 빨리 끝난 Promise 반환 |
| `async/await` | Promise를 더 깔끔하게 쓰는 문법 |

✔ Promise는 **한 번 상태가 정해지면 변하지 않음**  
✔ 비동기 흐름을 깔끔하게 연결하고, 에러를 한 곳에서 처리 가능  

