# 🌈 CSS - 색상 & 배경

HTML 요소의 **글자 색상, 배경색, 배경 이미지** 등을 지정할 때 사용하는 CSS의 기본이자 시각 디자인에서 매우 중요한 속성들입니다.

---

## 1️⃣ 색상 관련 속성

| 속성 | 설명 | 예시 |
|------|------|------|
| `color` | 텍스트 색상 | `color: red;` |
| `background-color` | 배경색 | `background-color: #f0f0f0;` |
| `border-color` | 테두리 색상 | `border: 1px solid blue;` |
| `outline-color` | 외곽선 색상 | `outline: 1px solid orange;` |
| `text-decoration-color` | 밑줄/취소선 색상 | `text-decoration-color: gray;` |

---

## 2️⃣ 색상 지정 방법

| 방법 | 예시 | 설명 |
|------|------|------|
| **이름(name)** | `red`, `blue`, `green` | 기본 색상명 |
| **16진수(hex)** | `#ff0000`, `#333` | 빨강, 검정 등 (짧은 3자리도 가능) |
| **RGB** | `rgb(255, 0, 0)` | 빨강 |
| **RGBA** | `rgba(255, 0, 0, 0.5)` | 빨강 + 50% 투명 |
| **HSL** | `hsl(120, 100%, 50%)` | 색상, 채도, 명도 |
| **HSLA** | `hsla(120, 100%, 50%, 0.5)` | + 투명도 |

---

## 3️⃣ 배경 관련 속성

| 속성 | 설명 | 예시 |
|------|------|------|
| `background-color` | 배경 색상 | `background-color: #eee;` |
| `background-image` | 배경 이미지 삽입 | `background-image: url('bg.jpg');` |
| `background-repeat` | 배경 반복 여부 | `repeat`, `no-repeat`, `repeat-x`, `repeat-y` |
| `background-position` | 배경 이미지 위치 | `center`, `top right`, `20px 40px` |
| `background-size` | 배경 이미지 크기 조절 | `cover`, `contain`, `100px 200px` |
| `background-attachment` | 스크롤 시 배경 고정 여부 | `scroll`, `fixed` |
| `background` | 배경 관련 속성 일괄 지정 | 예: `background: #000 url(bg.jpg) no-repeat center/cover;` |

#### 배경 속성 예시

```css
body {
  background-color: #f9f9f9;
  background-image: url('images/pattern.png');
  background-repeat: repeat-x;
  background-position: top left;
  background-size: 100px auto;
  background-attachment: fixed;
}
```

#### 단축 속성 예시 (`background`)

```css
/* 순서: color image repeat position / size attachment */
background: #fff url("bg.jpg") no-repeat center/cover fixed;
```

---

## 🧠 꿀팁

✔ `background-size: cover` → 이미지가 요소를 꽉 채움 (잘림 가능)  
✔ `background-size: contain` → 이미지가 전부 보이도록 조절 (여백 생길 수 있음)  
✔ `rgba()`, `hsla()` → **투명도**를 적용할 수 있음
