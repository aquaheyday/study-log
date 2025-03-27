# 🌐 브라우저의 구성 요소

웹 브라우저는 사용자가 입력한 URL을 통해 웹 페이지를 요청하고, **렌더링하여 보여주는 복합 소프트웨어**입니다.

---

## 1️⃣ 브라우저의 주요 구성 요소

| 구성 요소 | 설명 |
|-----------|------|
| **User Interface (UI)** | 주소창, 탭, 북마크, 뒤로가기 버튼 등 브라우저 외형 |
| **Browser Engine** | 렌더링 엔진과 UI 간의 **중간 조정자** 역할 |
| **Rendering Engine** | HTML, CSS를 읽고 **화면에 그리는 핵심 엔진** (예: Blink, WebKit) |
| **Networking** | HTTP 요청/응답 처리 (요청 전송, 캐싱 등) |
| **JavaScript Engine** | JS 코드 실행 엔진 (예: V8, SpiderMonkey) |
| **UI Backend** | 버튼, 입력창 등 기본 컴포넌트를 렌더링하는 모듈 |
| **Data Storage** | 쿠키, 로컬스토리지, 세션, IndexedDB 등 클라이언트 저장소 |

---

## 2️⃣ 브라우저 렌더링 흐름

#### 1. HTML 파싱 → DOM 트리 생성  
#### 2. CSS 파싱 → CSSOM (CSS Object Model) 생성  
#### 3. DOM + CSSOM → 렌더 트리(Render Tree) 구성  
#### 4. Layout 단계 → 각 요소의 위치 계산  
#### 5. Paint 단계 → 실제 픽셀로 그리기  
#### 6. Composite 단계 → 화면에 출력

✔ 이 과정을 **렌더링 파이프라인(Rendering Pipeline)** 이라고 부름  

---

## 3️⃣ 주요 렌더링 엔진

| 엔진 이름 | 사용 브라우저 |
|-----------|----------------|
| Blink | Chrome, Edge, Opera 등 |
| WebKit | Safari |
| Gecko | Firefox |

---

## 4️⃣ 주요 JavaScript 엔진

| 엔진 | 브라우저 |
|------|----------|
| V8 | Chrome, Edge |
| SpiderMonkey | Firefox |
| JavaScriptCore | Safari |

---

## 5️⃣ 브라우저 저장소 구성

| 저장소 | 설명 |
|--------|------|
| 쿠키 | 작은 데이터 (요청 시 자동 포함) |
| localStorage | 지속 저장, 도메인별로 보존 |
| sessionStorage | 탭 종료 시 삭제 |
| IndexedDB | 구조화된 대용량 저장소 (비동기 API) |

---

## 🎯 정리

✔ UI: 사용자 인터페이스  
✔ Engine: 렌더링 + 브라우저 엔진  
✔ JS Engine: 스크립트 해석 및 실행  
✔ Network: 요청/응답 처리  
✔ Storage: 클라이언트 측 저장소  
✔ 렌더링 과정: `DOM` → `CSSOM` → `Render Tree` → `Layout` → `Paint` → `Composite`
