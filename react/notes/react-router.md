# 🌍 React Router

**React Router**는 **React 애플리케이션에서 페이지 간 이동(라우팅)을 처리하는 라이브러리**입니다.  
싱글 페이지 애플리케이션(SPA)에서도 **클라이언트 측에서 URL을 변경하여 다른 페이지처럼 동작**하게 해줍니다.

---

## 1. React Router의 필요성

- 싱글 페이지 애플리케이션(SPA)에서 페이지 전환을 가능하게 함  
- 컴포넌트 기반으로 동적 라우팅 구현 가능  
- URL을 통해 특정 페이지 상태를 유지  
- 중첩 라우트, 보호된 라우트 등 다양한 기능 제공  

---

## 2. React Router 설치

```sh
npm install react-router-dom
```
✔ `react-router-dom` → 브라우저 환경에서 사용하는 React Router 패키지  

---

## 3. React Router 기본 사용법

### `BrowserRouter`로 감싸기
`BrowserRouter`를 사용하여 라우팅 기능을 활성화합니다.

```jsx
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Home from "./Home";
import About from "./About";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/about" element={<About />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
```
✔ `<BrowserRouter>` → 애플리케이션을 감싸서 라우팅 활성화  
✔ `<Routes>` → 여러 개의 라우트(Route)를 포함하는 컨테이너  
✔ `<Route>` → 개별 페이지를 정의 (예: `/` → `Home` 컴포넌트)  

---

## 4. `Link`, `useNavigate` (페이지 이동)

### `<Link>` 컴포넌트 사용 (기본)
`<Link>`를 사용하여 페이지 이동이 가능합니다.

```jsx
import { Link } from "react-router-dom";

function Navbar() {
  return (
    <nav>
      <Link to="/">홈</Link>
      <Link to="/about">소개</Link>
    </nav>
  );
}

export default Navbar;
```
✔ `<a>` 태그 대신 `<Link to="/">홈</Link>` 사용 → 새로고침 없이 페이지 전환  

---

### `useNavigate()` 훅 사용 (프로그래밍 방식으로 이동)
`useNavigate()`를 사용하면 특정 이벤트에서 페이지 이동이 가능합니다.

```jsx
import { useNavigate } from "react-router-dom";

function Home() {
  const navigate = useNavigate();

  return (
    <div>
      <h1>홈 페이지</h1>
      <button onClick={() => navigate("/about")}>소개 페이지로 이동</button>
    </div>
  );
}

export default Home;
```
✔ `navigate("/about")` → 버튼 클릭 시 `/about` 페이지로 이동  

---

## 5. `useParams` (동적 라우팅)

### `:id` (URL 파라미터 설정 사용)
```jsx
import { BrowserRouter, Routes, Route } from "react-router-dom";
import ProductDetail from "./ProductDetail";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/product/:id" element={<ProductDetail />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
```
✔ `/product/:id` → `id` 값을 동적으로 받는 페이지  

---

### `useParams()`로 URL 파라미터 가져오기
```jsx
import { useParams } from "react-router-dom";

function ProductDetail() {
  const { id } = useParams();

  return <h1>상품 상세 페이지 - 상품 ID: {id}</h1>;
}

export default ProductDetail;
```
✔ `useParams()` → 현재 URL의 `:id` 값을 가져옴  

---

## 6. `useSearchParams` (쿼리 스트링)

### `?key=value` (URL 쿼리 스트링 사용)
```jsx
import { useSearchParams } from "react-router-dom";

function SearchPage() {
  const [searchParams, setSearchParams] = useSearchParams();
  const keyword = searchParams.get("keyword");

  return (
    <div>
      <h1>검색 결과</h1>
      <p>검색어: {keyword}</p>
      <button onClick={() => setSearchParams({ keyword: "React" })}>
        React 검색
      </button>
    </div>
  );
}

export default SearchPage;
```
✔ `useSearchParams()` → URL의 쿼리 스트링 값 (`?keyword=React`)을 가져옴  

---

## 7. Nested Routes (중첩 라우트)

### `Outlet`을 사용한 중첩 라우팅
```jsx
import { BrowserRouter, Routes, Route, Outlet } from "react-router-dom";

function Dashboard() {
  return (
    <div>
      <h1>대시보드</h1>
      <Outlet />
    </div>
  );
}

function DashboardHome() {
  return <h2>대시보드 홈</h2>;
}

function DashboardSettings() {
  return <h2>설정</h2>;
}

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/dashboard" element={<Dashboard />}>
          <Route index element={<DashboardHome />} />
          <Route path="settings" element={<DashboardSettings />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
```
✔ `Outlet` → 중첩된 라우트의 내용을 렌더링  
✔ `/dashboard` → `DashboardHome` 표시  
✔ `/dashboard/settings` → `DashboardSettings` 표시  

---

## 8. `Protected Routes` (보호된 라우트)

로그인한 사용자만 접근 가능한 페이지를 만들려면 **보호된 라우트(Protected Route)** 를 설정해야 합니다.

```jsx
import { Navigate } from "react-router-dom";

function ProtectedRoute({ children }) {
  const isAuthenticated = false; // 로그인 여부 (예제)

  return isAuthenticated ? children : <Navigate to="/login" />;
}

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/dashboard" element={<ProtectedRoute><Dashboard /></ProtectedRoute>} />
        <Route path="/login" element={<Login />} />
      </Routes>
    </BrowserRouter>
  );
}
```
✔ `Navigate to="/login"` → 로그인하지 않으면 자동으로 로그인 페이지로 이동  

---

## 9. `HashRouter` vs `BrowserRouter`

| 라우터 | 설명 | URL 예시 |
|--------|------|---------|
| `BrowserRouter` | HTML5 `history` API 사용 | `https://example.com/about` |
| `HashRouter` | `#`을 사용한 라우팅 (서버 설정 불필요) | `https://example.com/#/about` |

✔ `BrowserRouter`는 서버에서 `404 에러`가 발생할 수 있으므로, 서버 설정이 필요  
✔ `HashRouter`는 `#`을 이용해 경로를 유지하기 때문에 서버 설정 없이 사용 가능  

---

## 🎯 정리
✔ React Router는 싱글 페이지 애플리케이션(SPA)에서 페이지 이동을 관리하는 라이브러리  
✔ <BrowserRouter>로 라우팅 활성화, <Routes>와 <Route>로 경로 설정  
✔ <Link to="/"> → 새로고침 없이 페이지 이동, useNavigate() → 프로그래밍 방식으로 이동  
✔ useParams() → URL 파라미터 (/product/:id) 가져오기  
✔ useSearchParams() → 쿼리 스트링 (?keyword=React) 관리  
✔ <Outlet>을 사용하면 중첩 라우트(Nested Routes) 구현 가능  
✔ ProtectedRoute를 활용해 인증된 사용자만 특정 페이지 접근 가능  
