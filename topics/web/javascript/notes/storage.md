# 💾 JavaScript - 웹 스토리지 (Web Storage)

- 웹 스토리지는 브라우저에 데이터를 **키-값 쌍(Key-Value Pair)** 형태로 저장할 수 있는 기능입니다.  
- 서버와 통신하지 않고 **클라이언트 측에서 데이터 유지**가 가능합니다.

---

## 1️⃣ 종류

| 종류              | 설명                                 | 저장 기간      | 접근 범위       |
|-------------------|--------------------------------------|----------------|-----------------|
| `localStorage`    | 브라우저에 **영구 저장**              | 영구           | 같은 출처(origin) |
| `sessionStorage`  | **브라우저 탭/세션이 유지되는 동안 저장** | 탭이 닫힐 때까지 | 같은 탭, 같은 출처 |

---

## 2️⃣ 기본 메서드

두 스토리지 모두 아래의 **공통 메서드**를 사용합니다:

```js
// 저장
localStorage.setItem("key", "value");

// 불러오기
const value = localStorage.getItem("key");

// 삭제
localStorage.removeItem("key");

// 전체 삭제
localStorage.clear();
```

✔ 값은 **문자열(string)** 으로 저장됨. 객체 저장 시 JSON.stringify / JSON.parse 필요

---

## 3️⃣ `localStorage` 예제

```js
localStorage.setItem("user", "Alice");

const user = localStorage.getItem("user");
console.log(user);  // "Alice"

localStorage.removeItem("user");

localStorage.clear(); // 모든 키 삭제
```

✔ 데이터는 브라우저를 껐다 켜도 유지됩니다.

---

## 4️⃣ `sessionStorage` 예제

```js
sessionStorage.setItem("temp", "123");

const temp = sessionStorage.getItem("temp");
console.log(temp); // "123"

sessionStorage.removeItem("temp");
```

✔ 탭을 닫거나 새로 열면 데이터가 사라집니다.

---

## 5️⃣ 객체 저장 & 불러오기

객체를 저장하려면 `JSON.stringify`, 꺼낼 땐 `JSON.parse`를 사용해야 합니다:

```js
const user = {
  name: "Jane",
  age: 25
};

localStorage.setItem("user", JSON.stringify(user));

const storedUser = JSON.parse(localStorage.getItem("user"));
console.log(storedUser.name); // "Jane"
```

---

## 6️⃣ localStorage vs sessionStorage 요약

| 항목              | `localStorage`        | `sessionStorage`     |
|-------------------|------------------------|------------------------|
| 지속 시간         | 영구 (삭제 전까지 유지) | 세션 종료 시 삭제      |
| 탭/창 공유        | 동일 출처이면 공유     | **탭마다 개별 저장소**  |
| 저장 용량         | 보통 5~10MB            | 보통 5~10MB            |
| 사용 용도         | 로그인 상태, 설정 저장 등 | 임시 데이터, 단일 작업 등 |

---

## 🎯 정리

✔ `localStorage`: 탭을 닫아도 데이터가 남음  
✔ `sessionStorage`: 탭을 닫으면 데이터 사라짐  
✔ 둘 다 문자열만 저장 가능 → 객체는 JSON으로 변환  
✔ 서버 없이도 클라이언트에서 간단한 상태 유지 가능
