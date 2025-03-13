# ⚛️ React 개요

## React란 무엇인가?
**React**는 Facebook(현재 Meta)에서 개발한 **사용자 인터페이스(UI) 라이브러리**입니다.  
컴포넌트 기반 아키텍처를 사용하여 **재사용 가능한 UI를 효율적으로 개발**할 수 있습니다.

- **빠르고 효율적인 렌더링**: Virtual DOM을 활용하여 최소한의 업데이트 수행
- **컴포넌트 기반 구조**: UI를 독립적인 컴포넌트로 분리하여 개발
- **단방향 데이터 흐름**: 데이터가 일관되게 관리됨
- **강력한 생태계**: Redux, React Router, Next.js 등과 함께 확장 가능

---

## React의 특징

### ✅ Virtual DOM
- **Virtual DOM**이란 **실제 DOM을 복제한 가상의 DOM**을 의미합니다.
- 변경이 발생하면 **Virtual DOM과 실제 DOM을 비교(diffing)**하여 필요한 부분만 업데이트합니다.
- 이 방식은 **렌더링 성능을 최적화**하는 데 도움을 줍니다.

### ✅ 컴포넌트 기반 개발
- React는 **재사용 가능한 컴포넌트**로 UI를 구성합니다.
- 각 컴포넌트는 **독립적으로 관리**되며, **재사용 가능**하여 유지보수가 용이합니다.

```jsx
function Greeting() {
  return <h1>Hello, React!</h1>;
}
```

### ✅ 단방향 데이터 흐름
- 데이터는 부모 → 자식 방향으로만 전달됩니다 (Props 사용).
- 단방향 데이터 흐름을 유지하면 **예측 가능한 상태 관리**가 가능합니다.

### ✅ 상태 관리 (State)
- `useState()`를 사용하여 **컴포넌트의 상태**를 관리할 수 있습니다.

```jsx
import { useState } from "react";

function Counter() {
  const [count, setCount] = useState(0);
  return (
    <div>
      <p>현재 카운트: {count}</p>
      <button onClick={() => setCount(count + 1)}>증가</button>
    </div>
  );
}
```

---

## React와 다른 프레임워크 비교

| 특징 | React | Vue.js | Angular |
|------|-------|--------|---------|
| 개발사 | Meta (Facebook) | Evan You | Google |
| 데이터 흐름 | 단방향 (Props, State) | 양방향 바인딩 | 단방향 |
| 상태 관리 | useState, Redux, Context | Vuex, Pinia | RxJS, NgRx |
| 컴포넌트 방식 | 함수형, 클래스형 | 옵션 API, Composition API | 모듈 기반 |
| 학습 난이도 | 중 | 쉬움 | 어려움 |

---

## React의 주요 사용처
✔ **SPA (Single Page Application)**: 페이지 전환 없이 동적 데이터 렌더링  
✔ **대규모 웹 애플리케이션**: 성능 최적화가 필요한 서비스  
✔ **모바일 앱 개발**: React Native를 사용하여 iOS/Android 앱 개발  
✔ **대시보드 및 관리자 페이지**: 빠른 개발과 유지보수 용이  

---

## React를 배우면 좋은 이유
✅ **컴포넌트 기반 개발**로 유지보수 용이  
✅ **Virtual DOM**을 이용한 성능 최적화  
✅ **풍부한 생태계**로 다양한 라이브러리 활용 가능  
✅ **대기업 및 스타트업에서 널리 사용** (Facebook, Netflix, Airbnb 등)  

---

## 공식 문서 및 학습 자료
- [React 공식 문서](https://react.dev/)
- [React 공식 튜토리얼](https://react.dev/learn)
- [React GitHub 저장소](https://github.com/facebook/react)

---

🚀 **이제 React의 기본 개념을 이해했으니, 실습을 시작해 보세요!**  
다음 개념: [React 프로젝트 설정](./setup-react.md) →
