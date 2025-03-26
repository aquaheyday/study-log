# ⏳ JavaScript - async/await

`async/await`는 **비동기 코드를 마치 동기 코드처럼 작성**할 수 있도록 도와주는 ES2017(ES8) 문법입니다.  
기존의 `Promise.then()` 체이닝보다 가독성이 뛰어나고 에러 처리도 간편합니다.

---

## 1️⃣ 기본 개념

### 1) `async` 함수

- 함수 앞에 `async` 키워드를 붙이면 **무조건 Promise를 반환**
- 내부에서 `await`를 사용할 수 있음

```js
async function greet() {
  return "Hello";
}

greet().then(msg => console.log(msg)); // Hello
```

---

### 2) `await` 키워드

- **Promise가 처리(resolve)될 때까지 기다림**
- `await`은 **오직 async 함수 내부에서만 사용 가능**

```js
async function fetchData() {
  const response = await fetch("https://api.example.com/data");
  const data = await response.json();
  console.log(data);
}
```

---

## 2️⃣ `async`/`await`로 비동기 흐름 제어

```js
function delay(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

async function run() {
  console.log("시작");
  await delay(1000); // 1초 대기
  console.log("1초 후 실행");
}

run();
```

#### 결과
```
시작
(1초 후)
1초 후 실행
```

---

## 3️⃣ `try`/`catch` 로 에러 처리

```js
async function getUser() {
  try {
    const res = await fetch("https://api.invalid-url");
    const data = await res.json();
    console.log(data);
  } catch (err) {
    console.error("에러 발생:", err.message);
  }
}
```

- `await`가 실패하면 **reject 상태가 되어 `catch`로 이동**

---

## 4️⃣ 병렬 실행: Promise.all()과 함께 사용

```js
async function loadAll() {
  const [user, posts] = await Promise.all([
    fetchUser(),
    fetchPosts()
  ]);

  console.log(user, posts);
}
```

- 병렬로 처리 → 효율적인 성능
- 개별 `await`로 순차 처리하면 **느려질 수 있음**

---

## 5️⃣ `async`/`await` vs `Promise`

| 항목            | Promise                    | async/await                     |
|-----------------|----------------------------|----------------------------------|
| 스타일          | `.then()`, `.catch()` 체인 | `await`, `try/catch` 구조         |
| 가독성          | 중첩될 수 있음             | 깔끔하고 동기 코드처럼 작성 가능 |
| 에러 처리       | `.catch()` 사용            | `try/catch` 사용                 |
| 학습 난이도     | 낮음                       | 약간 더 높지만 직관적            |

---

## 6️⃣ `await`는 일반 값도 기다릴까?

```js
async function demo() {
  const result = await 123;
  console.log(result); // 123
}
```

- `await`는 **`Promise`가 아니어도** 값을 그대로 반환함
- `await 123` → `Promise.resolve(123)`처럼 동작

---

## 7️⃣ `async` 함수는 항상 `Promise`를 반환

```js
async function foo() {
  return "bar";
}

foo().then(res => console.log(res)); // bar
```

- 반환값이 일반 값이더라도 **Promise로 감싸져 반환됨**

---

## 📌 사용 시 주의사항

- `await`는 **async 함수 내부에서만 사용 가능**
- 병렬 작업은 `Promise.all()`로 처리해야 더 빠름
- 에러 처리는 항상 `try/catch`로 감싸주는 습관!

---

## 🎯 정리

| 개념         | 설명 |
|--------------|------|
| `async`      | 함수가 Promise를 반환하도록 만듦 |
| `await`      | Promise가 처리될 때까지 대기 |
| `try/catch`  | 에러를 동기식처럼 처리 가능 |
| 사용 목적    | 가독성 있는 비동기 코드 작성 |
| 조합         | `Promise.all()`과 함께 사용 가능 |
