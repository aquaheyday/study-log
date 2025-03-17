# 🌍 비동기 처리

React에서 **비동기 처리(Asynchronous Handling)** 는 API 요청, 데이터 로딩, 상태 업데이트 등을 효과적으로 처리하는 중요한 개념입니다.  
비동기 처리를 통해 사용자 경험을 향상하고, 불필요한 렌더링을 방지할 수 있습니다.

---

## 1. 비동기 처리의 필요성

- 네트워크 요청을 효율적으로 처리 (API 호출, 데이터 가져오기 등)
- 사용자 인터페이스(UI)와 비동기 작업을 원활하게 연결
- 빠른 응답과 로딩 상태 표시 가능
- 비동기 작업이 끝난 후 UI 업데이트 용이

---

## 2. JavaScript 비동기 처리 기본 개념

JavaScript에서는 비동기 처리를 위해 **콜백 함수(Callback), Promise, async/await**을 사용할 수 있습니다.

### 콜백 함수 (Callback)
콜백 함수는 비동기 작업이 완료된 후 실행되는 함수입니다.

```javascript
function fetchData(callback) {
  setTimeout(() => {
    callback("데이터 로드 완료");
  }, 1000);
}

fetchData((message) => {
  console.log(message); // "데이터 로드 완료"
});
```

✔ 하지만 **콜백 지옥(Callback Hell)** 이 발생할 수 있어 가독성이 떨어질 수 있음  

### `Promise` 사용
`Promise`는 비동기 작업이 성공(`resolve()`)하거나 실패(`reject()`)할 경우의 처리를 제공하는 객체입니다.

```javascript
function fetchData() {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      resolve("데이터 로드 완료");
    }, 1000);
  });
}

fetchData().then((message) => {
  console.log(message); // "데이터 로드 완료"
});
```

✔ `.then()`을 사용해 비동기 작업 완료 후 실행할 로직을 지정 가능  

### 2-3. `async/await` 사용
`async/await`는 `Promise`를 보다 쉽게 다룰 수 있도록 도와줍니다.

```javascript
async function fetchData() {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve("데이터 로드 완료");
    }, 1000);
  });
}

async function loadData() {
  const data = await fetchData();
  console.log(data); // "데이터 로드 완료"
}

loadData();
```

✔ 동기 코드처럼 읽히면서 비동기 처리를 쉽게 할 수 있음  
✔ `try/catch`로 에러 핸들링 가능  

---

## 3. React에서 비동기 처리

### `useEffect`에서 `fetch()` 사용
React에서는 `useEffect`를 사용하여 API 요청을 보낼 수 있습니다.
```jsx
import { useEffect, useState } from "react";

function App() {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetch("https://jsonplaceholder.typicode.com/posts/1")
      .then((response) => response.json())
      .then((data) => {
        setData(data);
        setLoading(false);
      })
      .catch((error) => {
        setError(error);
        setLoading(false);
      });
  }, []);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error.message}</p>;

  return <div>{data.title}</div>;
}

export default App;
```

✔ `useEffect()`를 사용하여 API 호출  
✔ `setState()`를 통해 상태를 업데이트하여 렌더링  

### `useEffect`에서 `async/await` 사용
```jsx
import { useEffect, useState } from "react";

function App() {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch("https://jsonplaceholder.typicode.com/posts/1");
        const result = await response.json();
        setData(result);
      } catch (err) {
        setError(err);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error.message}</p>;

  return <div>{data.title}</div>;
}

export default App;
```

✔ `async` 함수 내부에서 `await fetch()` 사용  
✔ `try/catch`를 사용하여 에러 핸들링  

---

## 4. React에서 비동기 요청 라이브러리 활용

### `Axios` 사용
`fetch()`보다 더 간결하고 직관적인 `axios` 라이브러리를 사용할 수도 있습니다.
```sh
npm install axios
```
```jsx
import axios from "axios";
import { useEffect, useState } from "react";

function App() {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    axios
      .get("https://jsonplaceholder.typicode.com/posts/1")
      .then((response) => {
        setData(response.data);
        setLoading(false);
      })
      .catch((error) => {
        setError(error);
        setLoading(false);
      });
  }, []);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error.message}</p>;

  return <div>{data.title}</div>;
}

export default App;
```

✔ `axios.get()`을 사용해 JSON 변환 없이 바로 데이터 반환  

### React Query 사용
React Query는 **자동 캐싱, 리패칭** 등의 기능을 제공하는 비동기 데이터 관리 라이브러리입니다.

```sh
npm install @tanstack/react-query
```
```jsx
import { useQuery } from "@tanstack/react-query";
import axios from "axios";

const fetchPost = async () => {
  const { data } = await axios.get("https://jsonplaceholder.typicode.com/posts/1");
  return data;
};

function App() {
  const { data, isLoading, error } = useQuery(["post"], fetchPost);

  if (isLoading) return <p>Loading...</p>;
  if (error) return <p>Error: {error.message}</p>;

  return <div>{data.title}</div>;
}

export default App;
```

✔ `useQuery()`를 사용해 API 호출을 자동화  
✔ 자동으로 캐싱 및 데이터 리패칭 수행  

---

## 5. 비동기 처리 방식 비교

| 방식 | 사용 예시 | 주요 특징 |
|------|---------|---------|
| 콜백 함수 | `callback()` | 간단하지만 가독성이 낮음 (콜백 지옥 가능) |
| Promise | `.then().catch()` | 비동기 체이닝 가능, 가독성 개선 |
| async/await | `await fetchData()` | 동기 코드처럼 작성 가능, `try/catch`로 에러 처리 |
| React Query | `useQuery()` | 자동 캐싱, 리패칭, 상태 관리 기능 포함 |

---

## 🎯 정리
✔ 콜백 함수, Promise, async/await을 활용하여 비동기 작업 수행 가능  
✔ `fetch()`, `axios` → API 요청을 처리하는 기본적인 방법 (`useEffect`와 함께 사용)  
✔ `React Query` → 자동 캐싱, 리패칭, 상태 관리 기능 제공 (`useQuery()`)  
✔ 단순 요청에는 `fetch`/`axios`, 상태 관리가 필요한 경우 `React Query` 사용 추천  
