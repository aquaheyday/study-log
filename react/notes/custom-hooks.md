# 🚀 React 커스텀 훅

React에서 **커스텀 훅(Custom Hooks)** 은 **로직을 재사용하고 컴포넌트를 더욱 간결하게 유지하는 방법**입니다.  
**반복적인 상태 로직을 재사용**할 수 있도록 **기존의 React 훅(`useState`, `useEffect`, `useMemo` 등)을 활용하여 만든 함수**입니다.

---

## 1️⃣ 커스텀 훅이란?

- 반복적인 로직을 모듈화하여 재사용 가능
- 컴포넌트의 복잡성을 줄이고 가독성을 향상
- React의 기존 훅(`useState`, `useEffect` 등)을 기반으로 작성
- 이름이 반드시 `use`로 시작해야 함 (`useFetch`, `useCounter` 등)

---

## 2️⃣ 커스텀 훅을 사용해야 하는 경우

#### ✅ 언제 사용하면 좋을까?
- 컴포넌트 간에 동일한 상태 관리 로직이 반복될 때
- API 호출, 이벤트 핸들링, 폼 상태 관리 등과 같은 로직을 분리할 때
- 컴포넌트의 코드가 길어지고 복잡해질 때
- React의 내장 훅을 조합하여 새로운 기능을 만들고 싶을 때

---

## 3️⃣ 커스텀 훅 작성 및 사용법

#### 기본적인 구조
- `useState`, `useEffect`, `useMemo` 등 기본 훅을 활용하여 새로운 훅을 정의
- 필요한 데이터를 반환(`return`)하여 다른 컴포넌트에서 활용 가능  

```tsx
function useCustomHook() {
  // 내부적으로 React 훅 사용 가능
  const [state, setState] = useState(null);

  useEffect(() => {
    // 로직 수행
  }, []);

  return state; // 필요한 값 반환
}
```

---

## 4️⃣ 대표적인 커스텀 훅 예제 개념

### 1) `useCounter` (카운터 훅)
간단한 카운터 기능을 제공하는 커스텀 훅

```jsx
import { useState } from "react";

function useCounter(initialValue = 0) {
  const [count, setCount] = useState(initialValue);

  const increment = () => setCount((prev) => prev + 1);
  const decrement = () => setCount((prev) => prev - 1);
  const reset = () => setCount(initialValue);

  return { count, increment, decrement, reset };
}
```

#### 사용법

```jsx
const { count, increment, decrement, reset } = useCounter(0);
```

---

### 2) `useFetch` (API 호출을 처리하는 훅)
비동기 API 요청을 관리하는 커스텀 훅

```jsx
import { useState, useEffect } from "react";

function useFetch(url) {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch(url);
        const result = await response.json();
        setData(result);
      } catch (err) {
        setError(err);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, [url]);

  return { data, loading, error };
}
```

#### 사용법

```jsx
const { data, loading, error } = useFetch("https://jsonplaceholder.typicode.com/posts/1");
```

---

### 3) `useLocalStorage` (로컬 스토리지 활용 훅)
로컬 스토리지를 쉽게 다룰 수 있는 커스텀 훅

```jsx
import { useState, useEffect } from "react";

function useLocalStorage(key, initialValue) {
  const [value, setValue] = useState(() => {
    const storedValue = localStorage.getItem(key);
    return storedValue ? JSON.parse(storedValue) : initialValue;
  });

  useEffect(() => {
    localStorage.setItem(key, JSON.stringify(value));
  }, [key, value]);

  return [value, setValue];
}
```

#### 사용법

```jsx
const [theme, setTheme] = useLocalStorage("theme", "light");
```

---

## 5️⃣ 커스텀 훅 사용 시 주의할 점

### ✅ 훅의 규칙
- 반드시 함수 이름이 `use`로 시작해야 React에서 훅으로 인식  
- 훅 내부에서 다른 훅(`useState`, `useEffect` 등)을 사용할 수 있음  
- 조건문, 반복문 안에서 훅을 호출하지 말 것 (`useEffect` 내부에서는 가능)  
- 반환 값은 객체 또는 배열을 사용하여 다룰 데이터와 함수들을 관리  

#### ❌ 상태를 직접 수정하지 말 것
  
```jsx
function useCounter() {
  let count = 0; // ❌ 상태 관리 안됨 (useState 필요)
  return count;
}
```

✔ `useState` → 필수로 사용하여 상태를 관리해야 함

---

## 6️⃣ 커스텀 훅을 활용한 성능 최적화

- `useMemo` → 연산량이 큰 작업을 캐싱하여 불필요한 연산 방지  
- `useCallback` → 함수가 새로 생성되지 않도록 최적화  
- `useRef` → DOM 접근 및 불필요한 렌더링 방지  

### 1) `useDebounce` (디바운스 기능을 제공하는 훅)
입력값이 일정 시간 후에만 업데이트되도록 디바운싱 적용

```jsx
import { useState, useEffect } from "react";

function useDebounce(value, delay = 500) {
  const [debouncedValue, setDebouncedValue] = useState(value);

  useEffect(() => {
    const handler = setTimeout(() => setDebouncedValue(value), delay);
    return () => clearTimeout(handler);
  }, [value, delay]);

  return debouncedValue;
}
```

#### 사용법

```jsx
const debouncedSearch = useDebounce(searchTerm, 300);
```

---

## 7️⃣ 커스텀 훅을 활용하는 라이브러리

| 라이브러리 | 설명 |
|------------|---------------------------------------------------|
| React Query | 비동기 데이터 관리 (`useQuery`, `useMutation`) |
| SWR | 데이터 가져오기 및 캐싱 (`useSWR`) |
| Redux Toolkit | 상태 관리를 위한 커스텀 훅 (`useSelector`, `useDispatch`) |

---

## 🎯 정리
✔ 커스텀 훅(Custom Hook)은 반복적인 로직을 재사용할 수 있도록 만든 React 훅  
✔ `useState`, `useEffect`, `useMemo` 등의 내장 훅을 조합하여 새로운 기능을 구현 가능  
✔ `useCounter`, `useFetch`, `useLocalStorage` 같은 자주 사용되는 로직을 커스텀 훅으로 만들면 코드 재사용성이 높아짐  
✔ 훅의 규칙 → `use`로 시작해야 하고, 조건문/반복문 내부에서 호출하지 않도록 주의  
✔ `useMemo`, `useCallback`, `useDebounce` 등을 활용하면 성능 최적화도 가능  
