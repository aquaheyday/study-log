# 🎨 CSS란?

**CSS**는 **Cascading Style Sheets(종속 스타일 시트)** 의 약자로, HTML 요소의 **디자인(색상, 배치, 폰트 등 스타일)** 을 지정하는 언어입니다.
HTML이 **구조(내용)** 를 담당한다면, CSS는 **표현(스타일)** 을 담당합니다.

---

## 1️⃣ CSS의 주요 역할

- 글꼴, 색상, 배경, 여백 등 시각적 스타일 설정
- 레이아웃 구성 (그리드, 플렉스 등)
- 반응형 웹 구현 (미디어 쿼리 등)
- 애니메이션, 전환 효과 등 인터랙션 표현

---

## 2️⃣ CSS 문법 기본

```css
선택자 {
  속성: 값;
}
```

#### 예시

```css
h1 {
  color: blue;
  font-size: 2rem;
}
```

---

## 3️⃣ CSS 적용 방법 3가지

| 방법 | 설명 | 예시 |
|------|------|------|
| 인라인 스타일 | HTML 태그 안에 직접 작성 | `<p style="color:red;">텍스트</p>` |
| 내부 스타일 | `<style>` 태그 안에 작성 | `<style>p { color: red; }</style>` |
| 외부 스타일 | `.css` 파일로 분리 | `<link rel="stylesheet" href="style.css">` |

✔ 대규모 프로젝트에선 외부 스타일 방식 권장

---

## 4️⃣ 자주 쓰는 CSS 속성들

| 속성 | 설명 | 예시 |
|------|------|------|
| `color` | 글자 색상 | `color: red;` |
| `background-color` | 배경색 | `background-color: yellow;` |
| `font-size` | 글자 크기 | `font-size: 16px;` |
| `margin` | 바깥 여백 | `margin: 10px;` |
| `padding` | 안쪽 여백 | `padding: 5px;` |
| `border` | 테두리 | `border: 1px solid black;` |
| `width`, `height` | 너비/높이 | `width: 100px;` |
| `display` | 레이아웃 방식 | `display: flex;` |
| `position` | 위치 지정 방식 | `position: absolute;` |
| `text-align` | 텍스트 정렬 | `text-align: center;` |

---

## 5️⃣ 선택자(Selectors) 예시

| 선택자 | 설명 | 예시 |
|--------|------|------|
| `*` | 모든 요소 | `* { margin: 0; }` |
| `태그` | 태그 선택 | `p {}` |
| `.클래스` | 클래스 선택 | `.box {}` |
| `#아이디` | ID 선택 | `#header {}` |
| `요소 요소` | 하위 요소 | `div p {}` |
| `요소 > 요소` | 자식 요소 | `ul > li {}` |

---

## 6️⃣ CSS 관련 개념

| 개념 | 설명 |
|------|------|
| **Cascading(종속성)** | 여러 스타일이 겹칠 경우 우선순위에 따라 적용됨 |
| **Box Model** | 모든 요소는 박스 형태로 여백(margin), 경계(border), 안쪽 여백(padding), 콘텐츠(content)로 구성 |
| **Responsive Web** | 화면 크기에 따라 레이아웃이 유동적으로 변함 (미디어 쿼리 사용) |
| **Flexbox / Grid** | 레이아웃을 구성하는 최신 방식 |

---

## 🎯 정리

✔ CSS = HTML의 스타일을 설정하는 언어  
✔ `선택자 { 속성: 값; }` 형태로 작성  
✔ 적용 방식: 인라인 / 내부 / 외부  
✔ 자주 쓰는 속성: `color`, `font-size`, `margin`, `display`, 등  
✔ 핵심 개념: 박스 모델, 선택자, 반응형, 우선순위  
