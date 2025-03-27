# 📚 Browser - DOM & CSSOM

HTML과 CSS가 브라우저에 의해 해석될 때 각각 **DOM**과 **CSSOM**이라는 객체 모델로 변환됩니다.  
이 두 가지는 **렌더링 트리(Render Tree)** 를 만들기 위한 기반입니다.

---

## 1️⃣ DOM (Document Object Model)
HTML 문서를 브라우저가 **계층적 트리 구조**로 파싱한 결과

#### 특징
- 각 HTML 요소가 **노드(Node)** 로 표현됨
- JavaScript에서 `document.querySelector()` 등으로 접근 가능
- 실시간으로 조작 가능 (`createElement`, `appendChild`, `remove()` 등)

#### 예시

```html
<html>
  <body>
    <h1>Hello</h1>
    <p>World</p>
  </body>
</html>
```

#### DOM 트리 구조

```
document
└── html
    └── body
        ├── h1
        │   └── #text: "Hello"
        └── p
            └── #text: "World"
```

---

## 2️⃣ CSSOM (CSS Object Model)

CSS 스타일을 브라우저가 파싱하여 만든 **스타일 정보 트리 구조**

#### 특징
- 모든 CSS 규칙은 객체로 표현됨
- JavaScript로 접근 가능 (`document.styleSheets`, `CSSRule`)
- 각 요소에 적용될 **계산된 스타일**을 포함

#### 예시

```css
h1 {
  color: red;
  font-size: 20px;
}

p {
  display: none;
}
```

#### CSSOM 구조 (개념적으로)

```
Stylesheet
├── Rule: h1
│   ├── color: red
│   └── font-size: 20px
└── Rule: p
    └── display: none
```

---

## 3️⃣ DOM + CSSOM → Render Tree 생성

- 렌더링을 위해 브라우저는 DOM과 CSSOM을 결합하여 **렌더 트리(Render Tree)** 를 생성
- 이 렌더 트리를 기반으로 Layout, Paint, Composite 단계가 수행됨

```
RenderTree
└── body
    └── h1
        └── #text: "Hello"
```

✔ 렌더 트리는 보여지는 요소만 포함  

---

## 4️⃣ DOM vs CSSOM

| 항목 | DOM | CSSOM |
|------|-----|--------|
| 대상 | HTML 문서 | CSS 스타일 |
| 구조 | 노드 기반 트리 | 스타일 객체 트리 |
| 생성 시점 | HTML 파싱 시 | CSS 파싱 시 |
| 사용 목적 | 콘텐츠 구조 표현 | 스타일 표현 |
| 접근 방식 | `document.getElementById()` 등 | `document.styleSheets` 등 |
| 렌더링과의 관계 | 콘텐츠 구조 제공 | 스타일 정보 제공 |
| 변경 영향 | Reflow + Repaint 발생 가능 | Repaint 발생 가능 |

---

## 🎯 정리

✔ DOM: HTML 문서의 구조 표현  
✔ CSSOM: CSS 스타일 규칙을 객체화한 구조  
✔ 둘 다 렌더 트리 생성에 사용됨  
✔ JavaScript로 조작 가능 (단, 성능에 주의)  
✔ 실시간 조작은 Reflow/Repaint 발생 가능 → 성능 최적화 고려 필요
