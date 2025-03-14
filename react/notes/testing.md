# 🛠️ React 컴포넌트 테스트 with Jest & React Testing Library

React에서 **Jest와 React Testing Library(RTL)** 를 활용하여 컴포넌트를 효과적으로 테스트하는 방법을 정리했습니다.

---

## 1. Jest & React Testing Library란?

✅ **Jest**: JavaScript 테스팅 프레임워크로 React와 함께 사용하기 좋음  
✅ **React Testing Library (RTL)**: React 컴포넌트를 실제 사용자 관점에서 테스트하는 라이브러리  
✅ **단위(Unit) 테스트, 이벤트 핸들링 테스트, 비동기 API 테스트 등에 활용 가능**  
✅ **Jest의 `expect()`과 함께 사용하여 다양한 테스트 작성 가능**  

---

## 2. Jest & React Testing Library 설치하기

### ✅ Jest & RTL 설치
```sh
npm install --save-dev jest @testing-library/react @testing-library/jest-dom
```

### ✅ `package.json` 설정 추가
```json
{
  "scripts": {
    "test": "jest"
  }
}
```

---

## 3. React Testing Library 기본 개념

✅ **`render()`** - 컴포넌트를 가상 DOM에 렌더링  
✅ **`screen`** - 렌더링된 요소를 찾을 때 사용  
✅ **`fireEvent` / `userEvent`** - 이벤트를 시뮬레이션  
✅ **`waitFor()`** - 비동기 작업이 완료될 때까지 대기  
✅ **`expect()`** - Jest의 단언문으로 테스트 결과 검증  

---

## 4. 기본적인 컴포넌트 테스트

### 4-1. 렌더링 테스트
```jsx
// components/Hello.js
import React from 'react';

const Hello = ({ name }) => <h1>Hello, {name}!</h1>;

export default Hello;
```

```jsx
// __tests__/Hello.test.js
import { render, screen } from '@testing-library/react';
import Hello from '../components/Hello';

test('Hello 컴포넌트가 올바르게 렌더링되는지 테스트', () => {
  render(<Hello name="React" />);
  expect(screen.getByText('Hello, React!')).toBeInTheDocument();
});
```

---

## 5. 이벤트 핸들링 테스트

### 5-1. 버튼 클릭 테스트
```jsx
// components/Counter.js
import React, { useState } from 'react';

const Counter = () => {
  const [count, setCount] = useState(0);

  return (
    <div>
      <p>Count: {count}</p>
      <button onClick={() => setCount(count + 1)}>Increase</button>
    </div>
  );
};

export default Counter;
```

```jsx
// __tests__/Counter.test.js
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import Counter from '../components/Counter';

test('버튼 클릭 시 카운트가 증가하는지 테스트', async () => {
  render(<Counter />);
  const button = screen.getByText('Increase');
  
  await userEvent.click(button);
  expect(screen.getByText('Count: 1')).toBeInTheDocument();
});
```

---

## 6. 비동기 코드 테스트

✅ **API 호출을 포함한 컴포넌트 테스트 방법**  
✅ **`waitFor()`를 사용하여 비동기 작업이 완료될 때까지 대기**  

```jsx
// components/User.js
import React, { useEffect, useState } from 'react';

const User = ({ userId }) => {
  const [user, setUser] = useState(null);

  useEffect(() => {
    fetch(`https://jsonplaceholder.typicode.com/users/${userId}`)
      .then((res) => res.json())
      .then((data) => setUser(data));
  }, [userId]);

  if (!user) return <p>Loading...</p>;

  return <h2>{user.name}</h2>;
};

export default User;
```

```jsx
// __tests__/User.test.js
import { render, screen, waitFor } from '@testing-library/react';
import User from '../components/User';

global.fetch = jest.fn(() =>
  Promise.resolve({
    json: () => Promise.resolve({ name: 'John Doe' }),
  })
);

test('API 데이터를 불러와서 렌더링하는지 테스트', async () => {
  render(<User userId={1} />);
  expect(screen.getByText('Loading...')).toBeInTheDocument();

  await waitFor(() => expect(screen.getByText('John Doe')).toBeInTheDocument());
});
```

---

## 7. React Router와 Redux 상태 테스트

### 7-1. React Router 테스트
```jsx
// components/Home.js
import React from 'react';
import { Link } from 'react-router-dom';

const Home = () => (
  <div>
    <h1>Home</h1>
    <Link to="/about">Go to About</Link>
  </div>
);

export default Home;
```

```jsx
// __tests__/Home.test.js
import { render, screen } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';
import Home from '../components/Home';

test('Home 페이지가 올바르게 렌더링되는지 테스트', () => {
  render(
    <MemoryRouter>
      <Home />
    </MemoryRouter>
  );

  expect(screen.getByText('Home')).toBeInTheDocument();
  expect(screen.getByText('Go to About')).toBeInTheDocument();
});
```

---

## 8. Jest & RTL을 활용한 테스트 팁

✅ **가능한 한 사용자의 실제 행동을 시뮬레이션**  
✅ **React Testing Library의 쿼리 메서드를 적절히 활용 (`getByText`, `getByRole`)**  
✅ **비동기 테스트 시 `waitFor()`를 활용하여 데이터 로딩을 기다림**  
✅ **Jest의 `jest.fn()`을 사용하여 모의(Mock) 함수를 테스트**  
✅ **React Router, Redux 등의 상태 관리 라이브러리와도 함께 사용 가능**  

---

## 9. 공식 문서 및 추가 자료
- [Jest 공식 문서](https://jestjs.io/)
- [React Testing Library 공식 문서](https://testing-library.com/docs/react-testing-library/intro)

---

🚀 **React Jest & React Testing Library를 익혔다면, 이제 React 애플리케이션 배포를 배워봅시다!**  
다음 개념: [React 애플리케이션 배포](./deployment.md) →
