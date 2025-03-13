# 🌍 React 데이터 Fetching

**React 에서 데이터 Fetching (데이터 가져오기)** 은 서버나 API에서 데이터를 요청하고 받아오는 과정입니다.  
주로 `fetch()`, `axios`, `React Query`, `SWR` 등을 사용하여 데이터를 가져옵니다.

---

## 1. 데이터 Fetching의 필요성

✅ **클라이언트에서 서버 API 데이터를 요청 및 사용 가능**  
✅ **비동기 요청을 통해 사용자 경험 개선**  
✅ **React 상태 관리와 결합하여 동적 UI 제공**  
✅ **캐싱 및 자동 리패칭으로 성능 최적화 가능 (React Query, SWR 활용 시)**  

---

## 2. Fetch API 사용 (기본 방법)

```jsx
import { useEffect, useState } from "react";

function App() {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetch("https://jsonplaceholder.typicode.com/posts/1")
      .then((response) => {
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        return response.json();
      })
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

✅ `fetch()` → 네트워크 요청을 보내 데이터를 받아옴  
✅ `useEffect()` → 컴포넌트가 마운트될 때 데이터 요청 실행  
✅ `setState()` → 데이터를 상태에 저장 후 화면에 렌더링  

---

## 3. Axios 사용 (더 간결한 HTTP 요청)

```sh
npm install axios
```

```jsx
import { useEffect, useState } from "react";
import axios from "axios";

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

✅ `axios.get()` → 데이터를 요청하여 JSON 형식으로 자동 변환  
✅ `fetch()`보다 간결하고 직관적인 API 제공  

---

## 4. React Query 사용 (비동기 상태 관리)

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

✅ `useQuery()` → 자동으로 데이터 캐싱 및 리패칭 지원  
✅ `isLoading`, `error` 등의 상태 제공  

---

## 5. SWR 사용 (React Query와 유사)

```sh
npm install swr
```

```jsx
import useSWR from "swr";
import axios from "axios";

const fetcher = (url) => axios.get(url).then((res) => res.data);

function App() {
  const { data, error } = useSWR("https://jsonplaceholder.typicode.com/posts/1", fetcher);

  if (!data) return <p>Loading...</p>;
  if (error) return <p>Error: {error.message}</p>;

  return <div>{data.title}</div>;
}

export default App;
```

✅ `useSWR()` → 자동으로 데이터 캐싱 및 리패칭 수행  
✅ `fetcher()` → Axios를 활용해 데이터를 가져옴  

---

## 6. 데이터 Fetching 방식 비교

| 방법 | 캐싱 지원 | 추가 패키지 필요 | 사용 난이도 |
|------|---------|----------------|------------|
| Fetch API | ❌ | ❌ | 보통 |
| Axios | ❌ | ✅ (`axios`) | 쉬움 |
| React Query | ✅ | ✅ (`@tanstack/react-query`) | 중간 |
| SWR | ✅ | ✅ (`swr`) | 중간 |

### 결론
- **간단한 요청**: Fetch API 또는 Axios 사용
- **상태 관리 및 자동 리패칭 필요**: React Query 또는 SWR 사용

---

## 7. 공식 문서 및 추가 자료
- [React 공식 문서](https://react.dev/)
- [React Query 공식 문서](https://tanstack.com/query/latest)
- [SWR 공식 문서](https://swr.vercel.app/)
- [Axios 공식 문서](https://axios-http.com/)

---

🚀 **React 데이터 Fetching을 익혔다면, 이제 비동기 처리를 배워봅시다!**  
다음 개념: [비동기 처리](./async-await.md) →
