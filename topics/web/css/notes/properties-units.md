# 🎨 CSS - 속성과 단위

CSS는 HTML 요소의 **스타일(모양, 위치, 글꼴 등)을 설정**할 수 있게 해주는 언어입니다.  
속성(Property)과 단위(Unit)은 CSS의 핵심입니다.

---

## 1️⃣ 주요 CSS 속성 분류

### 1) 텍스트 관련 속성

| 속성 | 설명 | 예시 |
|------|------|------|
| `color` | 글자 색상 | `color: red;` |
| `font-size` | 글자 크기 | `font-size: 16px;` |
| `font-weight` | 글자 굵기 | `font-weight: bold;` |
| `font-family` | 글꼴 지정 | `font-family: Arial, sans-serif;` |
| `text-align` | 텍스트 정렬 | `text-align: center;` |
| `line-height` | 줄 간격 | `line-height: 1.6;` |
| `text-decoration` | 밑줄, 취소선 등 | `text-decoration: underline;` |

---

### 2) 박스 모델 관련 속성

| 속성 | 설명 |
|------|------|
| `width` / `height` | 요소의 너비/높이 |
| `padding` | 안쪽 여백 |
| `margin` | 바깥 여백 |
| `border` | 테두리 |
| `box-sizing` | 박스 크기 계산 방식 (`content-box`, `border-box`) |

---

### 3) 레이아웃 관련 속성

| 속성 | 설명 | 예시 |
|------|------|------|
| `display` | 요소의 표시 형식 | `block`, `inline`, `flex`, `grid` 등 |
| `position` | 요소 위치 지정 방식 | `static`, `relative`, `absolute`, `fixed`, `sticky` |
| `top`, `left`, `right`, `bottom` | 위치 조정 (position이 relative 이상일 때 사용) |
| `z-index` | 쌓이는 순서 (레이어) |
| `float` | 좌우 배치 지정 |
| `clear` | float 해제 |

---

### 4) 배경 / 색상 관련 속성

| 속성 | 설명 |
|------|------|
| `background-color` | 배경색 |
| `background-image` | 배경 이미지 |
| `background-size` | 배경 이미지 크기 |
| `background-repeat` | 반복 여부 |
| `opacity` | 투명도 (`0`~`1`) |

---

### 5) 기타 속성

| 속성 | 설명 |
|------|------|
| `cursor` | 마우스 커서 모양 (`pointer`, `default` 등) |
| `visibility` | 보이기/숨기기 (`visible`, `hidden`) |
| `overflow` | 넘치는 콘텐츠 처리 (`hidden`, `scroll`, `auto`) |
| `transition` | 전환 효과 설정 |
| `transform` | 요소 회전, 이동, 크기 등 조작 |
| `animation` | 애니메이션 효과 |

---

## 2️⃣ CSS 단위 정리

### 1) 절대 단위 (고정 크기)

| 단위 | 의미 |
|------|------|
| `px` | 픽셀 (화면의 점 단위) |
| `pt` | 포인트 (1pt = 1/72인치, 인쇄용) |
| `cm`, `mm`, `in` | 실제 길이 단위 (센티미터, 밀리미터, 인치) |

✔ 반응형 웹에서는 잘 사용하지 않음  

---

### 2) 상대 단위 (반응형 디자인에 유용)

| 단위 | 기준 | 예시 |
|------|------|------|
| `%` | 부모 요소 기준 비율 | `width: 50%` |
| `em` | **부모 요소의 폰트 크기** 기준 | `font-size: 1.5em` (부모가 16px이면 24px) |
| `rem` | **루트(html)의 폰트 크기** 기준 | `font-size: 2rem` |
| `vw` / `vh` | 뷰포트 너비/높이의 % | `width: 50vw` (뷰포트의 50%) |
| `vmin` / `vmax` | 뷰포트의 작은/큰 쪽 기준 | `vmin: 10` |
| `auto` | 자동 계산 (내용 또는 컨텍스트 기반) |

✔ 반응형 웹에서는 `em`, `rem`, `vw`, `%` 등을 자주 사용  

---

## 🎯 정리

✔ **속성(Property)**: HTML 요소에 스타일 적용 (색상, 크기, 위치 등)  
✔ **단위(Unit)**: 스타일 수치의 기준 (px, %, em 등)  
✔ `px`은 고정, `em`/`rem`/`%`는 유동적  
✔ 속성은 목적에 따라 텍스트/레이아웃/애니메이션 등으로 구분됨
