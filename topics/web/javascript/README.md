# 💻 JavaScript (for Web) - 브라우저 환경 중심 JS 정리

이 디렉토리는 브라우저 환경에서 사용하는 JavaScript 기능들을 정리합니다.  
언어 문법이 아닌, **DOM 조작, 이벤트 처리, 브라우저 API, 비동기 처리**와 같은 웹 중심 기능에 초점을 맞춥니다.

---

### 🌳 DOM (Document Object Model)
| 주제 | 파일명 | 설명 |
|------|--------|------|
| DOM이란? | [what-is-dom.md](./what-is-dom.md) | 브라우저에서 HTML 문서를 객체로 표현하는 구조 |
| DOM 선택자 | [selectors.md](./selectors.md) | getElementById, querySelector 등 요소 선택 |
| DOM 조작 | [dom-manipulation.md](./dom-manipulation.md) | 텍스트, 속성, 클래스 추가/수정 |
| DOM 노드 트리 | [dom-tree.md](./dom-tree.md) | 부모-자식-형제 관계 이해 |

---

### 🖱️ 이벤트 처리
| 주제 | 파일명 | 설명 |
|------|--------|------|
| 이벤트 모델 | [event-model.md](./event-model.md) | 캡처링, 버블링, 이벤트 위임 개념 |
| 이벤트 리스너 등록 | [event-listener.md](./event-listener.md) | addEventListener와 인라인 방식 비교 |
| 마우스 & 키보드 이벤트 | [mouse-keyboard.md](./mouse-keyboard.md) | click, keydown, input 등 주요 이벤트 |
| 이벤트 객체 | [event-object.md](./event-object.md) | target, currentTarget, preventDefault 등 |

---

### 🌀 비동기 & 타이머
| 주제 | 파일명 | 설명 |
|------|--------|------|
| setTimeout / setInterval | [timer.md](./timer.md) | 타이머 기반 비동기 처리 |
| 이벤트 루프 이해 | [event-loop.md](./event-loop.md) | 태스크 큐, 콜스택, 마이크로태스크 |
| fetch & XMLHttpRequest | [fetch.md](./fetch.md) | Ajax 요청 방식과 예제 |
| async/await 기초 | [async-await.md](./async-await.md) | 비동기 코드를 동기처럼 작성하는 방식 |

---

### 🌐 브라우저 API & 객체 모델
| 주제 | 파일명 | 설명 |
|------|--------|------|
| BOM이란? | [what-is-bom.md](./what-is-bom.md) | window, navigator, screen, location 객체 |
| localStorage / sessionStorage | [storage.md](./storage.md) | 브라우저 저장소 API |
| location & history 객체 | [location-history.md](./location-history.md) | URL 조작 및 브라우저 히스토리 관리 |
| alert, prompt, confirm | [dialog.md](./dialog.md) | 브라우저 기본 팝업 처리 방법 |
