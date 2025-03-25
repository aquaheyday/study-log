# 🔁 JavaScript - 반복문

JavaScript에서 반복문은 **동일한 작업을 여러 번 실행**할 때 사용됩니다.

---

## 1️⃣ `for` 문

가장 기본적인 반복문 (초기식 → 조건식 → 증감식 순서로 반복)

```js
for (let i = 0; i < 5; i++) {
  console.log("i =", i);
}
```

---

## 2️⃣ `while` 문

조건이 **참인 동안 계속 반복** (조건을 먼저 검사 → 조건이 거짓이면 한 번도 실행되지 않을 수 있음)

```js
let count = 0;
while (count < 3) {
  console.log("count =", count);
  count++;
}
```

---

## 3️⃣ `do...while` 문

**최소 한 번은 실행**됨 (조건을 나중에 검사함)

```js
let num = 0;
do {
  console.log("num =", num);
  num++;
} while (num < 3);
```

---

## 4️⃣ `for...of` 문

**배열/문자열 등 iterable 객체**를 순회.

```js
const fruits = ["🍎", "🍌", "🍇"];
for (const fruit of fruits) {
  console.log(fruit);
}
```

✔ 문자열도 순회 가능

```js
for (const char of "hello") {
  console.log(char);
}
```

---

## 5️⃣ `for...in` 문

**객체의 key(속성 이름)** 을 순회

```js
const user = { name: "Alice", age: 25 };

for (const key in user) {
  console.log(key, ":", user[key]);
}
```

✔ 객체 순회 시 `for...in` 사용, 배열에는 `for...of` 를 권장

---

## 6️⃣ `break` / `continue`

- `break`: 반복문을 **즉시 종료**
- `continue`: 현재 반복을 **건너뛰고 다음 반복으로 진행**

```js
for (let i = 1; i <= 5; i++) {
  if (i === 3) continue; // 3 건너뜀
  if (i === 5) break;    // 5에서 종료
  console.log(i);
}
// 출력: 1, 2, 4
```

---

## 🎯 정리

| 구문 | 특징 | 사용 예 |
|------|------|----------|
| `for` | 일반적인 반복 | 카운터 기반 루프 |
| `while` | 조건 기반 반복 | 조건이 불명확할 때 |
| `do...while` | 최소 1회 실행 | 메뉴 반복 등 |
| `for...of` | iterable 순회 | 배열, 문자열 등 |
| `for...in` | 객체 속성 순회 | key 탐색 |

✔ 배열에는 `for`, `forEach`, `for...of`  
✔ 객체에는 `for...in`  

