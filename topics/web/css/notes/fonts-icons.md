# ✍️ CSS - 커스텀 폰트 & 아이콘

웹에서는 기본 시스템 폰트 외에도 **커스텀 폰트**, **아이콘 폰트**, **SVG 아이콘**을 활용해 다양한 스타일을 적용할 수 있습니다.

---

## 1️⃣ 커스텀 웹 폰트 사용하기

### 1) `@font-face` 직접 사용

```css
@font-face {
  font-family: 'MyFont';
  src: url('fonts/MyFont.woff2') format('woff2'),
       url('fonts/MyFont.woff') format('woff');
  font-weight: normal;
  font-style: normal;
}

body {
  font-family: 'MyFont', sans-serif;
}
```

✔ `woff2`, `woff`는 웹에서 가장 권장되는 형식  

---

### 2) Google Fonts 불러오기

```html
<!-- HTML에 링크 추가 -->
<link href="https://fonts.googleapis.com/css2?family=Roboto&display=swap" rel="stylesheet">
```

```css
body {
  font-family: 'Roboto', sans-serif;
}
```

✔ 가볍고 쉽게 쓸 수 있는 무료 웹폰트 서비스  

---

### 3) 커스텀 폰트 팁

- `font-display: swap;` 속성을 사용하면 폰트 지연 시 기본 폰트 먼저 보여줌
- 폰트 크기 단위는 `rem`, `%` 등을 사용해서 반응형 대응하기

---

## 2️⃣ 아이콘 넣는 방법 2가지

### 1) 아이콘 폰트(Font Icon)

#### 1. 대표 라이브러리

- [Font Awesome](https://fontawesome.com/)
- [Material Icons](https://fonts.google.com/icons)

#### 2. 사용 예시 (Font Awesome)

```html
<!-- 링크 추가 -->
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.5.0/css/all.min.css">

<!-- 사용 -->
<i class="fas fa-home"></i>
<i class="fa-solid fa-user"></i>
```

✔ 아이콘을 `폰트처럼` 처리 → 크기, 색상, 정렬 등이 자유로움  

---

### 2) SVG 아이콘

#### 1. 사용 방법 1: `img` 태그로 삽입

```html
<img src="icon/search.svg" alt="검색 아이콘" width="24" height="24">
```

#### 2. 사용 방법 2: 인라인 삽입

```html
<svg width="24" height="24" fill="currentColor" viewBox="...">
  <path d="..." />
</svg>
```

✔ `fill="currentColor"`로 텍스트 색상 따라 색상 변경 가능  

---

### 3) 아이콘 스타일링 예시

```css
.icon {
  width: 24px;
  height: 24px;
  color: #333;
  margin-right: 8px;
  vertical-align: middle;
}
```

---

## 3️⃣ 폰트 아이콘 vs SVG 아이콘 차이

| 항목 | 폰트 아이콘 | SVG 아이콘 |
|------|-------------|-------------|
| 품질 | 확대 시 흐림 가능 | 벡터라 선명 |
| 스타일링 | 폰트 스타일처럼 적용 | CSS 다양하게 제어 가능 |
| 접근성 | `aria-hidden` 등 설정 필요 | `<title>`, `aria-label` 가능 |
| 사용 방식 | 클래스 추가 | 인라인 or `<img>` 사용 |
| 제어력 | 제한적 | 매우 유연 |

✔ 최신 웹에서는 **SVG 아이콘**이 더 많이 쓰이는 추세  

---

## 🎯 정리

✔ `@font-face` 또는 Google Fonts로 커스텀 폰트 사용  
✔ Font Awesome 등으로 아이콘 폰트 사용 가능  
✔ SVG 아이콘은 선명하고 커스터마이징 유리  
✔ 아이콘 접근성도 함께 고려하기 (`alt`, `aria-label`, `role="img"` 등)
