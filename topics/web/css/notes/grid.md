# 🧱 CSS Grid Layout 정리

**CSS Grid**는 웹 페이지의 콘텐츠를 **행(row)과 열(column)** 기준으로 **2차원 레이아웃**을 구성할 수 있는 가장 강력한 레이아웃 시스템입니다.

✔ `Flexbox`가 1차원(가로 or 세로)이라면, `Grid`는 **2차원(가로 + 세로)** 배치가 가능  

---

## 1️⃣ 기본 문법

```css
.container {
  display: grid;
}
```

✔ `Grid`를 적용하려면 부모 요소에 `display: grid` 또는 `inline-grid` 설정  

---

## 2️⃣ 주요 속성 (컨테이너용)

| 속성 | 설명 | 예시 |
|------|------|------|
| `display` | 그리드 컨테이너 지정 | `display: grid` |
| `grid-template-columns` | 열의 개수와 크기 정의 | `grid-template-columns: 1fr 1fr 1fr;` |
| `grid-template-rows` | 행의 개수와 크기 정의 | `grid-template-rows: 100px 200px;` |
| `gap` | 셀 간 간격 (간단히) | `gap: 20px;` |
| `row-gap`, `column-gap` | 행, 열 간격 따로 설정 | `row-gap: 10px;` |
| `grid-auto-rows` / `grid-auto-columns` | 자동 생성되는 행/열 크기 | `grid-auto-rows: 100px;` |
| `justify-items` | 셀 내부의 수평 정렬 | `start`, `center`, `end`, `stretch` |
| `align-items` | 셀 내부의 수직 정렬 | `start`, `center`, `end`, `stretch` |
| `place-items` | `justify-items` + `align-items` 단축 | `place-items: center;` |

---

## 3️⃣ 주요 속성 (아이템용)

| 속성 | 설명 | 예시 |
|------|------|------|
| `grid-column` | 열 영역 설정 | `grid-column: 1 / 3;` |
| `grid-row` | 행 영역 설정 | `grid-row: 2 / 4;` |
| `grid-area` | 셀 하나에 이름 부여하거나 범위 지정 | `grid-area: header;` |
| `justify-self` | 셀 내부에서 수평 정렬 | `center`, `start` 등 |
| `align-self` | 셀 내부에서 수직 정렬 | `center`, `end` 등 |
| `place-self` | 위 둘을 단축 | `place-self: center;` |

---

## 4️⃣ 예제

```html
<div class="grid-container">
  <div class="item a">A</div>
  <div class="item b">B</div>
  <div class="item c">C</div>
</div>
```

```css
.grid-container {
  display: grid;
  grid-template-columns: 1fr 2fr 1fr;
  grid-template-rows: 100px 100px;
  gap: 10px;
}

.a {
  grid-column: 1 / 3;
}

.b {
  grid-row: 1 / 3;
}
```

---

## 5️⃣ fr 단위란?

`fr`은 **grid에서만 사용하는 유연한 비율 단위**  
- `1fr` = 사용 가능한 공간의 1단위  
- 예: `grid-template-columns: 1fr 2fr` → 전체 공간을 3등분해서 1:2 비율로 분배

---

## 6️⃣ `Grid` vs `Flexbox`

| 항목 | `Grid` | `Flexbox` |
|------|------|---------|
| 레이아웃 | 2차원 (행 + 열) | 1차원 (가로 또는 세로) |
| 배치 대상 | 전체 영역 기반 | 콘텐츠 흐름 기반 |
| 정렬 방식 | 명확한 셀 위치 | 순차적 배치 |
| 사용 예 | 전체 페이지 레이아웃 | 내부 아이템 정렬 |

✔ 복잡한 레이아웃은 **`Grid`**, 아이템 정렬은 **`Flex`**가 더 적합  

---

## 🎯 정리

✔ `display: grid`로 2차원 레이아웃 시작  
✔ `grid-template-columns` / `rows`로 구조 정의  
✔ `gap`으로 셀 간격 조절  
✔ `grid-column`, `grid-row`으로 셀 위치 지정  
✔ `fr`, `minmax()`, `repeat()` 등으로 유연한 레이아웃 설계 가능  
