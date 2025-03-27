# 📱 CSS - 뷰포트 & 미디어쿼리

웹 페이지가 **다양한 화면 크기(PC, 태블릿, 모바일 등)** 에서도 잘 보이도록 스타일을 조정하려면 **뷰포트 설정**과 **미디어 쿼리**가 꼭 필요합니다.

---

## 1️⃣ 뷰포트(Viewport)란?

**뷰포트**는 사용자가 보는 **브라우저 화면의 영역**을 의미합니다.  
화면 영역은 데스크탑 기준으로 표시되기 때문에, 뷰포트를 조정하지 않으면 모바일 환경에서는 깨진 레이아웃이 나타날 수 있습니다.

---

### 뷰포트 설정 메타 태그

```html
<meta name="viewport" content="width=device-width, initial-scale=1.0">
```

| 속성 | 설명 |
|------|------|
| `width=device-width` | 기기의 실제 화면 너비에 맞춤 |
| `initial-scale=1.0` | 초기 확대/축소 비율 설정 |

---

## 2️⃣ 미디어 쿼리(Media Query)란?

**미디어 쿼리**는 화면 크기, 해상도, 기기 유형 등에 따라 **다른 CSS 스타일을 적용**할 수 있게 해주는 조건문 같은 도구입니다.

---

### 기본 문법

```css
@media (조건) {
  /* 조건이 만족될 때 적용할 CSS */
}
```

---

## 3️⃣ 자주 사용되는 미디어쿼리 예시

#### 조건 종류

| 조건 | 설명 | 예시 |
|------|------|------|
| `max-width` | 최대 너비 이하 | `(max-width: 768px)` |
| `min-width` | 최소 너비 이상 | `(min-width: 1024px)` |
| `orientation` | 방향 (가로/세로) | `(orientation: portrait)` |
| `prefers-color-scheme` | 사용자 테마 | `(prefers-color-scheme: dark)` |

### 1) 모바일 전용 스타일

```css
@media (max-width: 767px) {
  body {
    background-color: lightyellow;
  }
}
```

✔ `767px 이하` → 스마트폰 기준  

---

### 2) 태블릿, 데스크탑

```css
@media (min-width: 768px) and (max-width: 1023px) {
  .container {
    padding: 40px;
  }
}
```

---

### 3) 데스크탑 이상

```css
@media (min-width: 1024px) {
  .grid {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
  }
}
```

---

## 4️⃣ 뷰포트 너비 기준 (예시)

| 기기 | 화면 크기 (가이드용) |
|------|-----------------------|
| 모바일 | ~767px |
| 태블릿 | 768px ~ 1023px |
| 데스크탑 | 1024px 이상 |

⚠️ 실제 기기 해상도는 다르지만, **스타일 적용 기준으로 자주 사용되는 값** 입니다.

---

## 5️⃣ 미디어쿼리 적용 방식 2가지

### 1) CSS 안에 직접 작성

```css
@media (max-width: 600px) {
  h1 {
    font-size: 20px;
  }
}
```

---

### 2) 링크된 CSS 파일 분리

```html
<link rel="stylesheet" href="mobile.css" media="(max-width: 600px)">
```

---

## 🎯 정리

✔ `<meta name="viewport">` → 뷰포트 설정  
✔ `@media (조건)` → 미디어 쿼리를 사용해 화면 크기별 스타일 적용  
✔ `min-width`, `max-width` 조합으로 반응형 웹 구현  
