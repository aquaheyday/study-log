# 📐 CSS - `display` 속성

CSS의 `display` 속성은 **HTML 요소의 박스 유형과 레이아웃 방식**을 정의합니다.  
즉, **요소가 어떻게 화면에 표시될지**를 결정하는 중요한 속성입니다.

---

## 1️⃣ 기본 `display` 값 종류

| 값 | 설명 | 특징 |
|-----|------|------|
| `block` | 블록 요소 | 줄바꿈, 전체 너비 차지 (`div`, `p` 등 기본) |
| `inline` | 인라인 요소 | 줄바꿈 없음, 내용만큼만 너비 (`span`, `a` 등 기본) |
| `inline-block` | 인라인처럼 흐르고, 블록처럼 크기 조절 가능 | `width`, `height` 지정 가능 |
| `none` | 표시 안 함 (숨김) | 요소가 렌더링되지 않음 |
| `flex` | 플렉스 컨테이너 | 자식 요소 정렬에 유연함 |
| `grid` | 그리드 컨테이너 | 2차원 레이아웃 구성 가능 |
| `inline-flex` | 인라인 플렉스 박스 | 인라인 흐름 + 플렉스 속성 |
| `inline-grid` | 인라인 그리드 박스 | 인라인 흐름 + 그리드 속성 |
| `table` | 테이블처럼 표시 | HTML `<table>`과 유사한 동작 |
| `contents` | 요소 자체는 사라지고, 자식만 렌더링됨 | 부모 박스는 사라지지만 DOM에는 남음 |

---

## 2️⃣ 블록 vs 인라인 차이

| 속성 | 블록(`block`) | 인라인(`inline`) |
|------|----------------|------------------|
| 줄바꿈 | 있음 | 없음 |
| 너비/높이 지정 | 가능 | ❌ 불가능 (`inline-block` 필요) |
| 박스 모델 적용 | 전체 적용 | 일부만 적용 (padding, margin 좌우만) |

#### 예제
```html
<div style="display: block;">Block</div>
<span style="display: inline;">Inline</span>
```

---

## 3️⃣ 유용한 활용 예시

### 1) `display: none;`

```css
.modal {
  display: none;
}
```

✔ 요소를 숨길 때 사용    
✔ HTML 구조에는 존재하지만 **화면에는 표시되지 않음**  

---

### 2) `display: inline-block;`

```css
.button {
  display: inline-block;
  padding: 10px 20px;
}
```

✔ 버튼처럼 보이는 인라인 요소 만들기  

---

### 3) `display: flex;`

```css
.container {
  display: flex;
  justify-content: space-between;
}
```

✔ 가로 방향 정렬, 반응형 레이아웃에 유용  
✔ `flex-direction`, `gap`, `align-items` 등과 함께 사용  

---

## 4️⃣ 특수한 값

### `display: contents;`

```css
.wrapper {
  display: contents;
}
```

✔ 요소 자체는 렌더링하지 않고, 자식 요소만 그대로 출력됨  
✔ 레이아웃 조정에 유용하지만 **접근성이나 브라우저 호환성 주의**  

---

## 🎯 정리

| 값 | 요약 설명 |
|----|-----------|
| `block` | 한 줄 전체 차지 |
| `inline` | 줄 안에서 흐름 유지 |
| `inline-block` | 인라인 + 크기 조절 가능 |
| `none` | 렌더링에서 제거됨 |
| `flex` / `grid` | 최신 레이아웃 시스템 |
| `contents` | 부모 껍데기 없이 자식만 렌더링 |
