# 🎨 Browser - 렌더링 파이프라인

**렌더링 파이프라인(Rendering Pipeline)** 은 브라우저가 HTML, CSS, JS를 해석하고 **최종적으로 화면에 그리기까지의 과정**을 말합니다.

---

## 1️⃣ 렌더링 파이프라인 전체 흐름

1. **HTML 파싱** → **DOM 트리 생성**
2. **CSS 파싱** → **CSSOM 트리 생성**
3. **DOM + CSSOM** → **Render Tree(렌더 트리)** 생성
4. **Layout (Reflow)** → 각 요소의 **크기와 위치 계산**
5. **Paint** → 요소들을 **픽셀로 변환**
6. **Compositing (합성)** → 레이어들을 **하나의 화면으로 출력**

---

## 2️⃣ 각 단계 상세 설명

### 1) HTML 파싱 → DOM 생성
- HTML을 파싱해 **문서 객체 모델(DOM)** 트리 생성
- JavaScript도 실행되며 DOM 조작 가능

---

### 2) CSS 파싱 → CSSOM 생성
- CSS를 파싱해 **CSSOM(CSS Object Model)** 트리 생성
- 스타일 정보 계산

---

### 3) Render Tree 생성
- DOM + CSSOM → **렌더 트리(Render Tree)** 구성
- **시각적으로 보여지는 요소만 포함**
- `display: none` 요소는 포함되지 않음

---

### 4) Layout 단계 (Reflow)
- 각 요소의 **크기와 위치 계산**
- 뷰포트와 폰트, 패딩, 마진 등 기반으로 레이아웃 결정

---

### 5) Paint 단계
- 배경색, 텍스트, 박스 그림자 등 **스타일 정보를 픽셀로 변환**
- 각 요소의 시각적 표현 준비

---

### 6) Composite 단계
- Paint된 결과를 **레이어별로 합성**하여 **하나의 최종 화면**으로 출력
- GPU가 사용될 수 있음 (특히 `transform`, `opacity`)

---

## 3️⃣ 렌더링 파이프라인 요약 흐름

```text
HTML → DOM
↓
CSS  → CSSOM
↓
DOM + CSSOM → Render Tree
↓
Layout (Reflow)
↓
Paint
↓
Composite → 최종 화면
```

---

## 4️⃣ 성능 최적화와 연관된 단계

| 단계 | 최적화 키워드 |
|------|----------------|
| Layout | Reflow 최소화 (`transform`, `opacity` 사용) |
| Paint | Repaint 최소화 (컬러 변경 최소화) |
| Composite | GPU-friendly 속성 사용 (`will-change`) |
| 전체 | Lazy Loading, 비동기 로딩, DOM 조작 최소화 |

---

## 5️⃣ 관련 용어 요약

| 용어 | 설명 |
|------|------|
| **DOM** | HTML로 구성된 문서 구조 트리 |
| **CSSOM** | CSS로 구성된 스타일 정보 트리 |
| **Render Tree** | 실제 화면에 보여질 요소의 정보 |
| **Reflow** | 레이아웃 계산 (크기·위치) |
| **Repaint** | 시각적 스타일 다시 그림 |
| **Composite** | 그려진 요소들을 합쳐 출력 (GPU 가능) |

---

## 🎯 정리

✔ 렌더링 파이프라인은 **6단계**  
✔ DOM → CSSOM → Render Tree → Layout → Paint → Composite  
✔ 성능 최적화를 위해 Reflow / Repaint 최소화 필요  
✔ `transform`, `opacity`, `will-change` 사용으로 GPU 활용 유도
