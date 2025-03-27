# 🧭 Browser Internals - 브라우저 동작 원리

이 디렉토리는 웹 브라우저의 내부 구조와 작동 방식을 정리합니다.  
HTML/CSS/JS 파싱 및 렌더링 흐름, 이벤트 루프, 브라우저 보안 정책, 멀티 프로세스 아키텍처 등 웹 개발자가 꼭 알아야 할 브라우저 지식을 포함합니다.

---

### 🧱 브라우저 기본 구조
| 주제 | 파일명 | 설명 |
|------|--------|------|
| 브라우저의 구성 요소 | [browser-components.md](./notes/browser-components.md) | 렌더링 엔진, 자바스크립트 엔진, 네트워크, UI 등 |
| 렌더링 엔진 개요 | [rendering-engine.md](./notes/rendering-engine.md) | Chromium, Gecko 등 파서와 렌더링 흐름 |
| 브라우저 멀티 프로세스 구조 | [multi-process.md](./notes/multi-process.md) | 브라우저, 렌더러, 네트워크, GPU 프로세스 |

---

### 🎨 렌더링 과정 & DOM 구성
| 주제 | 파일명 | 설명 |
|------|--------|------|
| 렌더링 파이프라인 | [rendering-pipeline.md](./notes/rendering-pipeline.md) | DOM → CSSOM → Render Tree → Layout → Paint → Composite |
| DOM & CSSOM | [dom-cssom.md](./notes/dom-cssom.md) | 브라우저가 HTML/CSS를 객체로 구성하는 방식 |
| Reflow & Repaint | [reflow-repaint.md](./notes/reflow-repaint.md) | 레이아웃 계산과 화면 다시 그리기의 차이점 |

---

### ⚙️ 자바스크립트 & 이벤트 루프
| 주제 | 파일명 | 설명 |
|------|--------|------|
| 자바스크립트 엔진 개요 | [js-engine.md](./notes/js-engine.md) | V8 등 JS 코드를 처리하는 엔진 구조 |
| 이벤트 루프와 큐 | [event-loop.md](./notes/event-loop.md) | 콜스택, 태스크 큐, 마이크로태스크 큐 동작 원리 |
| setTimeout vs Promise | [event-loop-example.md](./notes/event-loop-example.md) | 비동기 코드의 실행 순서 비교 |

---

### 🔐 브라우저 보안과 제약
| 주제 | 파일명 | 설명 |
|------|--------|------|
| Same-Origin Policy | [same-origin.md](./notes/same-origin.md) | 보안상의 출처 제약 개념 |
| CORS 작동 방식 | [cors.md](./notes/cors.md) | 교차 출처 요청의 허용/제한 원리 |
| sandbox & iframe 보안 | [sandbox.md](./notes/sandbox.md) | iframe 분리, 속성 설정으로 보안 격리 |
| 브라우저 저장소 보안 | [storage-security.md](./notes/storage-security.md) | localStorage, sessionStorage, 쿠키 보안 이슈 |

---

### 🛠️ 기타
| 주제 | 파일명 | 설명 |
|------|--------|------|
| 브라우저 캐시 동작 | [browser-cache.md](./notes/browser-cache.md) | HTTP 캐시 정책과 브라우저 캐싱 메커니즘 |
| 사용자 입력 처리 흐름 | [input-processing.md](./notes/input-processing.md) | 키 입력, 클릭 등이 DOM에 반영되는 과정 |
| Web API와 Window 객체 | [web-api.md](./notes/web-api.md) | 브라우저가 제공하는 API와 BOM 개요 |
