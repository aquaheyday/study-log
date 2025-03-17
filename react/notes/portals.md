# 🚀 Portals

React의 **Portals(포털)** 은 **컴포넌트를 렌더링할 때, 부모 컴포넌트의 DOM 계층 구조를 벗어나 다른 DOM 노드에 렌더링할 수 있도록 해주는 기능**입니다.  
기본적으로 React 컴포넌트는 부모 요소 내부에 렌더링되지만, **Portals를 사용하면 특정 DOM 요소로 렌더링을 이동**할 수 있습니다.

---

## 1. React Portals란?
- 컴포넌트의 렌더링 위치를 변경할 수 있도록 하는 기능
- 모달, 드롭다운, 툴팁 등 특정 UI 요소를 body의 최상위에 렌더링할 때 유용
- 이벤트 전파는 기존 부모 컴포넌트와 동일하게 동작
- 기존 React 렌더링 구조를 유지하면서도 다른 DOM 위치에 렌더링 가능

---

## 2. React Portals를 사용하는 이유
- 모달, 드롭다운, 툴팁 등의 UI 요소를 `body` 태그 아래에 렌더링할 때 필요
- CSS `z-index` 문제를 방지하여 레이아웃 깨짐 없이 최상위에서 렌더링 가능
- 부모 요소의 `overflow: hidden` 속성 영향을 받지 않고 독립적으로 표시 가능
- 렌더링 위치를 변경하면서도 이벤트 전파는 기존 React 계층을 유지

---

## 3. Portals의 주요 개념

### `ReactDOM.createPortal()`
- `ReactDOM.createPortal(child, container)` 형식으로 사용  
- `child`: 렌더링할 React 요소  
- `container`: React 요소를 렌더링할 DOM 노드 (예: `document.getElementById("portal-root")`)  

### Portals의 이벤트 전파 특징
- 렌더링 위치는 이동하지만, 이벤트 전파는 기존 React 트리 구조를 유지
- 이벤트 버블링(Bubbling)은 여전히 부모 컴포넌트에서 처리됨
- 이벤트를 차단하려면 `event.stopPropagation()`을 활용

---

## 4. Portals 사용 사례

### 모달 창 (Modal)
- 부모 요소의 `overflow: hidden` 영향을 받지 않고 독립적으로 렌더링 가능  

### 드롭다운 (Dropdown)
- 특정 UI 요소를 원하는 위치에 렌더링 가능  

### 툴팁 (Tooltip)
- 부모 요소가 `position: relative`여도 원하는 위치에 툴팁 배치 가능  

---

## 5. Portals 사용 시 주의할 점

⚠ Portals는 부모 컴포넌트의 DOM 계층을 벗어나지만, React의 이벤트 전파는 그대로 유지됨  
⚠ 모달 창이나 드롭다운을 만들 때 `stopPropagation()`을 사용하여 이벤트 버블링을 막을 수 있음  
⚠ Portals를 사용할 때는 `document.getElementById()`를 통해 대상 DOM을 확인해야 함  

---

## 🎯 정리
✔ Portals는 React 컴포넌트를 기존 DOM 계층을 벗어나 특정 DOM 요소에 렌더링하는 기능  
✔ `ReactDOM.createPortal(child, container)` → 원하는 container에 컴포넌트 렌더링  
✔ 모달, 드롭다운, 툴팁 같은 UI 요소를 `body` 최상위에서 렌더링할 때 유용  
✔ 이벤트 버블링은 기존 부모 컴포넌트와 동일하게 동작  
✔ `stopPropagation()`을 활용하면 불필요한 이벤트 전파 방지 가능  
