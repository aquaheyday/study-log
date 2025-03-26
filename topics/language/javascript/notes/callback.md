# 🔁 JavaScript - 콜백 함수

**다른 함수에 인자로 전달되어, 그 함수 내부에서 나중에 호출되는 함수**를 말합니다.

---

## 1️⃣ 콜백 함수란?

- **함수의 인자로 전달되는 함수**
- 어떤 작업이 끝난 후 **나중에 실행될 함수**
- 비동기 처리, 이벤트 처리, 반복 작업 등에서 자주 사용됨

```js
function greetUser(callback) {
  const name = "Alice";
  callback(name);
}

greetUser(function(name) {
  console.log(`Hi, ${name}!`);
});
// 출력: Hi, Alice!
```

---

## 2️⃣ 왜 콜백 함수를 사용할까?

| 상황 | 예시 |
|------|------|
| 동기적 흐름에서 로직 분리 | 반복문, 조건문 안의 함수 실행 |
| 비동기 처리 | setTimeout, Ajax, 파일 읽기 등 |
| 이벤트 처리 | 버튼 클릭, 폼 제출 등 |

---

## 3️⃣ 대표적인 콜백 함수

### 1) 반복 작업 (`forEach`)

```js
const arr = [1, 2, 3];

arr.forEach(function(item) {
  console.log(item);
});
```

---

### 2) 비동기 처리 (`setTimeout`)

```js
setTimeout(function() {
  console.log("3초 후 실행!");
}, 3000);
```

---

### 3) 이벤트 핸들러

```js
document.querySelector("button").addEventListener("click", function() {
  console.log("버튼 클릭됨!");
});
```

---

## 4️⃣ 콜백 함수를 직접 만들어 보기

```js
function doSomething(task, callback) {
  console.log(`Doing: ${task}`);
  callback();
}

doSomething("청소하기", function() {
  console.log("끝났어요!");
});
```

---

## 5️⃣ 콜백 함수 vs 일반 함수 차이

| 구분 | 일반 함수 | 콜백 함수 |
|------|-----------|------------|
| 실행 시점 | 정의하고 직접 호출 | 다른 함수에 의해 호출 |
| 목적 | 기능을 수행 | 어떤 일이 끝난 후 실행될 동작 전달 |
| 예시 | `greet()` | `setTimeout(() => {}, 1000)` |

---

## 6️⃣ 콜백 지옥 (Callback Hell)

콜백을 너무 많이 중첩해서 **가독성이 떨어지는 코드**를 말함

```js
doSomething1(function() {
  doSomething2(function() {
    doSomething3(function() {
      // 계속 깊어짐...
    });
  });
});
```

---

## 7️⃣ 콜백 지옥 해결 방법

### 1) `Promise`

"미래에 어떤 일이 성공하거나 실패할 수 있는 값" 을 나타내는 객체, 비동기 작업의 성공(resolve), 실패(reject) 상태를 명확하게 관리

#### 예제
```js
const promise = new Promise((resolve, reject) => {
  // 비동기 작업
  if (성공) {
    resolve("성공 결과");
  } else {
    reject("에러 발생");
  }
});

promise
  .then(result => {
    console.log("성공:", result);
  })
  .catch(error => {
    console.error("실패:", error);
  });
```

#### `Promise` 체이닝으로 콜백 지옥 탈출

```js
getUser()
  .then(user => getPosts(user.id))
  .then(posts => getComments(posts[0].id))
  .then(comments => console.log(comments))
  .catch(err => console.error("에러:", err));
```

✔ 각 단계는 .then() 으로 연결하고, 에러는 .catch() 하나로 관리 가능

---

### 2) `async` / `await`

- `Promise`를 더 동기식(순서대로) 으로 보이게 만드는 문법
- `await`는 `Promise`가 처리될 때까지 기다림
- 에러는 `try/catch`로 처리

#### 예제

```js
async function fetchData() {
  try {
    const user = await getUser();
    const posts = await getPosts(user.id);
    const comments = await getComments(posts[0].id);
    console.log(comments);
  } catch (err) {
    console.error("에러 발생:", err);
  }
}

fetchData();
```

✔ 마치 동기 코드처럼 보이지만, 실제로는 비동기 처리

---

## 🎯 정리

✔ 콜백 함수는 **다른 함수에 인자로 전달되는 함수**  
✔ 나중에 특정 시점에 호출되도록 함  
✔ 반복, 이벤트, 비동기 처리에 자주 사용됨  
✔ 중첩되면 콜백 지옥 → **Promise / async/await**로 개선 가능

