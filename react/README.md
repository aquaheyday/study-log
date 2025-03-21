# ⚛️ React

이 폴더는 **React 학습 과정**에서 정리한 자료, 예제 코드, 프로젝트를 보관하는 공간입니다.  
기본 개념부터 실습 예제, 프로젝트 구성까지 단계적으로 학습한 내용을 체계적으로 정리합니다.

---

## 📂 디렉토리 구성

| 폴더명 | 설명 |
|---|---|
| [notes](./notes) | 개념 정리 및 학습 노트 |
| [examples](./examples) | 주요 기능별 예제 코드 (컴포넌트, 상태관리 등) |
| [projects](./projects) | 미니 프로젝트 및 실습 프로젝트 |

---

## 📋 개념 정리 목록

### 📌 기본 개념
| 번호 | 주제 | 파일명 | 설명 |
|---|---|---|---|
| 01 | React 개요 | [react-intro.md](./notes/react-intro.md) | React란 무엇인가? 특징과 장점 |
| 02 | 프로젝트 설정 | [setup-react.md](./notes/setup-react.md) | CRA, Vite로 프로젝트 생성, 환경 설정 |
| 03 | JSX 문법 | [jsx-basics.md](./notes/jsx-basics.md) | JSX 기본 문법, 표현식, 조건부 렌더링 |
| 04 | 컴포넌트 기초 | [components.md](./notes/components.md) | 함수형 컴포넌트, Props 사용법 |
| 05 | 상태 관리 | [state.md](./notes/state.md) | useState 훅, 상태 변경과 렌더링 |

### 🔲 컴포넌트 심화
| 번호 | 주제 | 파일명 | 설명 |
|---|---|---|---|
| 06 | 이벤트 처리 | [events.md](./notes/events.md) | onClick, onChange, 이벤트 핸들링 |
| 07 | 폼과 입력 관리 | [forms.md](./notes/forms.md) | controlled/uncontrolled 컴포넌트 |
| 08 | useEffect 훅 | [useeffect.md](./notes/useeffect.md) | 마운트/업데이트/언마운트 시 실행 |
| 09 | 컴포넌트 라이프사이클 | [lifecycle.md](./notes/lifecycle.md) | 클래스 컴포넌트와 라이프사이클 이해 |

### 🔄 상태 관리
| 번호 | 주제 | 파일명 | 설명 |
|---|---|---|---|
| 10 | Context API | [context-api.md](./notes/context-api.md) | 전역 상태 관리, useContext 사용법 |
| 11 | Redux 기본 개념 | [redux-basics.md](./notes/redux-basics.md) | Redux 개념, store, reducer, action |
| 12 | Redux Toolkit | [redux-toolkit.md](./notes/redux-toolkit.md) | Redux Toolkit을 활용한 상태 관리 |

### 🌍 네트워크 및 데이터 관리
| 번호 | 주제 | 파일명 | 설명 |
|---|---|---|---|
| 13 | React Router | [react-router.md](./notes/react-router.md) | 페이지 이동, 동적 라우팅, useParams |
| 14 | 데이터 Fetching | [fetching-data.md](./notes/fetching-data.md) | Fetch API, Axios, SWR, React Query |
| 15 | 비동기 처리 | [async-await.md](./notes/async-await.md) | async/await, useEffect와 API 호출 |

### 🚀 고급 개념
| 번호 | 주제 | 파일명 | 설명 |
|---|---|---|---|
| 16 | 성능 최적화 | [performance.md](./notes/performance.md) | React.memo, useMemo, useCallback |
| 17 | 커스텀 훅 | [custom-hooks.md](./notes/custom-hooks.md) | 재사용 가능한 커스텀 훅 만들기 |
| 18 | Refs 사용법 | [useRef.md](./notes/useRef.md) | DOM 조작, useRef 활용법 |
| 19 | Portals | [portals.md](./notes/portals.md) | 모달을 루트 밖에 렌더링하는 방법 |

### 🛠️ 테스트 및 배포
| 번호 | 주제 | 파일명 | 설명 |
|---|---|---|---|
| 20 | Jest와 React Testing Library | [testing.md](./notes/testing.md) | 컴포넌트 테스트 방법 |
| 21 | React 애플리케이션 배포 | [deployment.md](./notes/deployment.md) | Netlify, Vercel, Firebase 배포 | 

---

## 📋 예제 목록

| 번호 | 주제 | 파일명 | 설명 |
|---|---|---|---|
| 01 | 컴포넌트 기초 | [basic-component.jsx](./examples/basic-component.jsx) | 함수형 컴포넌트 구성법 |
| 02 | 이벤트 핸들링 | [event-handling.jsx](./examples/event-handling.jsx) | 버튼 클릭 이벤트 처리 |
| 03 | 리스트 렌더링 | [list-rendering.jsx](./examples/list-rendering.jsx) | map()을 활용한 리스트 렌더링 |
| 04 | 상태관리 기초 | [simple-state.jsx](./examples/simple-state.jsx) | useState 기본 활용 |
| 05 | Props 전달 | [props-example.jsx](./examples/props-example.jsx) | 부모→자식 데이터 전달 패턴 |
| 06 | 폼 입력 처리 | [form-handling.jsx](./examples/form-handling.jsx) | Form 상태 관리 및 제출 처리 |
| 07 | API 호출 및 데이터 표시 | [fetch-data.jsx](./examples/fetch-data.jsx) | useEffect+fetch로 데이터 요청 및 표시 |

---

## 📋 프로젝트 목록

| 번호 | 프로젝트명 | 폴더명 | 설명 |
|---|---|---|---|
| 01 | 관리자 사이트 | [vite-ts-scss-admin-dashboard](./projects/vite-ts-scss-admin-dashboard) | CRUD 기능을 갖춘 관리자 사이트  |

---

## 📚 참고 자료
- [React 공식 문서](https://react.dev/) - 최신 React 공식 문서 및 가이드  
- [React API 레퍼런스](https://react.dev/reference/react) - React의 주요 API 및 Hook 정리  
- [Create React App](https://create-react-app.dev/) - React 프로젝트 생성 및 설정 가이드  
- [React TypeScript 가이드](https://react-typescript-cheatsheet.netlify.app/) - TypeScript와 함께 React 사용법 정리  

---

## 📢 업데이트 로그
- 2025-03-06: 초기 구성
- 2025-03-13: notes 초기 구성
- 2025-03-14: notes 내용 추가
- 2025-03-17: notes 내용 정리 추가
- 2025-03-19: 디렉토리 구성 수정, projects > vite-ts-scss-admin-dashboard 추가
- 2025-03-21: notes 템플릿 수정 
- 새로운 예제 추가 시 목록 업데이트 예정
