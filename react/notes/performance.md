# 🚀 성능 최적화

React 애플리케이션의 **렌더링 성능을 최적화**하여 불필요한 연산을 줄이고, 사용자 경험을 향상시키는 방법을 정리합니다.

---

## 1. 성능 최적화가 필요한 이유

- 불필요한 렌더링을 방지하여 성능 개선
- 빠른 상태 업데이트 및 렌더링 최적화
- 사용자 경험 향상 (UX 개선)
- 메모리 및 CPU 사용량 절약

---

## 2. 불필요한 렌더링 방지

### React.memo` (컴포넌트 메모이제이션)
`React.memo`를 사용하면 **동일한 `props`가 전달될 때 컴포넌트의 불필요한 재렌더링을 방지**할 수 있습니다.

```jsx
import React from "react";

const MyComponent = React.memo(({ name }) => {
  console.log("렌더링됨");
  return <div>{name}</div>;
});

export default MyComponent;
```

✔ props가 변경되지 않으면 이전 렌더링 결과를 재사용  
⚠ 객체, 배열, 함수 등 참조형 데이터 변경 시 리렌더링 발생 → useMemo, useCallback 사용 필요

---

## 3. 연산 비용이 큰 작업 최적화

### `useMemo` (값의 계산 결과 캐싱)
`useMemo`는 연산량이 큰 작업을 캐싱하여 불필요한 재계산을 방지합니다.

```jsx
import { useMemo } from "react";

const MyComponent = ({ numbers }) => {
  const sum = useMemo(() => {
    console.log("합계를 계산하는 중...");
    return numbers.reduce((acc, curr) => acc + curr, 0);
  }, [numbers]);

  return <div>합계: {sum}</div>;
};

export default MyComponent;
```

✔ `numbers`가 변경되지 않으면 기존 계산 결과를 재사용  
⚠ 불필요한 `useMemo` 사용은 오히려 성능 저하 초래 가능  

---

## 4. 함수 재생성 방지

### `useCallback` (함수 메모이제이션)
`useCallback`은 컴포넌트가 리렌더링될 때 동일한 함수가 매번 새로 생성되는 문제를 방지합니다.

```jsx
import { useCallback } from "react";

const MyComponent = ({ onClick }) => {
  const handleClick = useCallback(() => {
    console.log("클릭됨");
  }, []);

  return <button onClick={handleClick}>클릭</button>;
};

export default MyComponent;
```

✔ 이벤트 핸들러, 콜백 함수 최적화 가능  
⚠ 의존성 배열이 변경되면 함수가 새로 생성됨  

---

## 5. 리스트 렌더링 최적화

### `key` 속성 최적화
리스트를 렌더링할 때 고유한 `key` 값을 지정하면 렌더링 성능이 향상됩니다.

```jsx
const items = ["Apple", "Banana", "Cherry"];

return (
  <ul>
    {items.map((item) => (
      <li key={item}>{item}</li> // 고유한 값 사용
    ))}
  </ul>
);
```

⚠ 배열의 index를 key로 사용하면 리스트 변경 시 성능 저하 발생 가능  

### 가상화 (Virtualization) 적용
긴 리스트를 효율적으로 렌더링하기 위해 `react-window`, `react-virtualized` 등의 라이브러리를 활용할 수 있습니다.

```jsx
npm install react-window
```

```jsx
import { FixedSizeList as List } from "react-window";

const Row = ({ index, style }) => <div style={style}>Item {index}</div>;

const MyList = () => (
  <List height={300} itemCount={1000} itemSize={35} width={"100%"}>
    {Row}
  </List>
);

export default MyList;
```

✔ 보이는 항목만 렌더링하여 성능 최적화 가능

---

## 6. 불필요한 리렌더링 방지

### `shouldComponentUpdate` (클래스 컴포넌트)
클래스형 컴포넌트에서 `shouldComponentUpdate`를 사용하면 특정 조건에서만 리렌더링을 수행할 수 있습니다.

```jsx
class MyComponent extends React.Component {
  shouldComponentUpdate(nextProps) {
    return this.props.value !== nextProps.value; // 값이 변경될 때만 렌더링
  }

  render() {
    return <div>{this.props.value}</div>;
  }
}

export default MyComponent;
```

### `PureComponent` 활용
`PureComponent`는 자동으로 `shouldComponentUpdate`를 구현하여 성능을 최적화합니다.

```jsx
import React, { PureComponent } from "react";

class MyComponent extends PureComponent {
  render() {
    console.log("렌더링됨");
    return <div>{this.props.value}</div>;
  }
}

export default MyComponent;
```

✔ 객체, 배열 등의 참조형 데이터가 변경될 때는 useMemo와 함께 사용 권장

---

## 7. React 성능 최적화 비교

| 기법 | 사용 목적 | 주요 특징 |
|------|----------|----------|
| `React.memo` | 컴포넌트 리렌더링 방지 | `props`가 변경되지 않으면 이전 렌더링 결과 재사용 |
| `useMemo` | 연산량이 많은 값 캐싱 | 특정 값이 변경될 때만 연산 수행 |
| `useCallback` | 함수 재생성 방지 | 이벤트 핸들러, 콜백 함수 최적화 |
| `shouldComponentUpdate` | 클래스형 컴포넌트 최적화 | 상태 변화 감지 후 렌더링 여부 결정 |
| `PureComponent` | `shouldComponentUpdate` 자동 적용 | 클래스형 컴포넌트 성능 개선 |
| `react-window` | 가상화(Virtualization) 적용 | 긴 리스트 렌더링 최적화 |

---

## 🎯 정리
✔ `React.memo` → `props` 변경이 없을 때 렌더링 방지  
✔ `useMemo` → 연산량이 많은 값의 재계산 방지  
✔ `useCallback` → 이벤트 핸들러 및 콜백 함수의 재생성 방지  
✔ `react-window` → 긴 리스트의 렌더링 성능 최적화 (가상화 적용)  
✔ `shouldComponentUpdate`, `PureComponent` → 클래스형 컴포넌트의 리렌더링 최적화  
