# ⚙️ Browser - 렌더링 엔진 개요

웹 브라우저의 **렌더링 엔진(Rendering Engine)** 은 `HTML`, `CSS`, `JavaScript`를 해석해 **화면에 웹 페이지를 그려주는 핵심 모듈**입니다.

---

## 1️⃣ 렌더링 엔진이 하는 일

| 단계 | 작업 | 설명 |
|------|------|------|
| 1 | HTML 파싱 | HTML → DOM(Document Object Model) 트리 생성 |
| 2 | CSS 파싱 | CSS → CSSOM(CSS Object Model) 트리 생성 |
| 3 | Render Tree 생성 | DOM + CSSOM → 실제 보이는 요소 트리 |
| 4 | Layout (Reflow) | 각 요소의 크기와 위치 계산 |
| 5 | Paint | 화면에 각 요소를 그리기 |
| 6 | Composite | 레이어들을 합쳐 최종 화면 생성 (GPU 사용 가능) |

✔ 이 과정을 **렌더링 파이프라인(Rendering Pipeline)** 이라고 부름  

---

## 2️⃣ 주요 렌더링 엔진 종류

| 엔진 | 사용 브라우저 |
|------|----------------|
| **Blink** | Chrome, Edge, Opera 등 |
| **WebKit** | Safari, iOS 브라우저 |
| **Gecko** | Firefox |
| **Trident (구형)** | IE (Internet Explorer) |
| **EdgeHTML (구형)** | 구 버전 Microsoft Edge |

✔ 대부분의 최신 브라우저는 **Blink** 또는 **WebKit** 기반을 사용  

---

## 3️⃣ Blink 엔진 (by Google)

- Chrome, Edge, Opera, Brave 등에서 사용
- WebKit에서 분기되어 개발됨
- **V8 JavaScript 엔진과 긴밀히 통합**

####  엔진 구성 예시 (Blink 기준)

| 구성 요소 | 설명 |
|-----------|------|
| HTML Parser | DOM 생성 |
| CSS Parser | CSSOM 생성 |
| Style Engine | 스타일 계산 |
| Layout Engine | 크기/위치 계산 |
| Painting | 픽셀 정보 계산 |
| Compositor | GPU로 최종 출력 |

---

## 4️⃣ 렌더링 최적화 방법

- DOM/CSSOM 접근 최소화
- `transform`, `opacity`로 애니메이션 처리 (GPU 활용)
- Reflow/Repaint 줄이기
- Lazy Loading 활용

### 1) 리플로우(Reflow)란?

요소의 **위치나 크기**가 변경되어 **레이아웃을 다시 계산**해야 하는 과정

#### 발생 예시

- 요소의 `width`, `height`, `margin`, `padding`, `position`, `display` 변경
- DOM 추가/제거 (`appendChild`, `removeChild`)
- 브라우저 창 크기 조절
- 글자 수가 달라지거나 이미지가 로드됨

---

### 2) 리페인트(Repaint)란?

요소의 위치는 그대로지만, **시각적 속성(색상, 배경, 테두리 등)** 이 변경될 때 발생하는 과정

#### 발생 예시

- `color`, `background-color`, `border-color`
- `visibility`, `box-shadow`, `outline` 변경
- CSS 클래스 토글 등

---

### 3) Lazy Loading 이란?

- 웹 페이지에서 모든 리소스를 한 번에 불러오는 것이 아니라, **사용자가 필요한 시점에만 로드하는 기법**입니다.
- 대표적인 예는 **이미지, 비디오, iframe** 요소의 지연 로딩입니다.


#### 1. HTML 속성만으로 구현 (`loading="lazy"`)

```html
<img src="image.jpg" loading="lazy" alt="지연 로딩 이미지">
<iframe src="video.html" loading="lazy"></iframe>
```

✔ 지원 브라우저: Chrome, Edge, Firefox, Safari (일부)  

#### 2. JavaScript + IntersectionObserver API

```html
<img data-src="image.jpg" class="lazy" alt="지연 로딩 이미지">
```

```js
const images = document.querySelectorAll('img.lazy');

const observer = new IntersectionObserver((entries, obs) => {
  entries.forEach(entry => {
    if (entry.isIntersecting) {
      const img = entry.target;
      img.src = img.dataset.src;
      img.classList.remove('lazy');
      obs.unobserve(img);
    }
  });
});

images.forEach(img => observer.observe(img));
```

✔ 뷰포트에 **요소가 등장할 때만 로드**됨  
✔ 이미지 외에도 **무한 스크롤, 콘텐츠 로딩** 등에 활용 가능  

---

## 🎯 정리

✔ 렌더링 핵심 = DOM + CSSOM → Render Tree  
✔ Reflow, Repaint (성능 부담)  
✔ `transform`과 `opacity`는 GPU 활용으로 부담 적음  
✔ Lazy Loading = 초기 속도 ↑, 트래픽 ↓  
