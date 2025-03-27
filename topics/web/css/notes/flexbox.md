# 📦 CSS - Flexbox

**Flexbox(Flexible Box Layout)** 는 1차원(가로 또는 세로)의 요소 배치와 정렬을 쉽고 유연하게 할 수 있도록 도와주는 CSS 레이아웃 방식입니다.

---

## 1️⃣ Flexbox 핵심 개념

- 부모 요소: **Flex 컨테이너**
- 자식 요소: **Flex 아이템**
- 주 축(main axis) vs 교차 축(cross axis)

```css
.container {
  display: flex;
}
```

---

## 2️⃣ Flex 구성요소

- **컨테이너(부모 요소)**: `display: flex` 또는 `inline-flex`
- **아이템(자식 요소)**: 정렬 대상

---

## 3️⃣ 컨테이너용 속성

| 속성 | 설명 | 예시 |
|------|------|------|
| `display` | Flex 컨테이너 지정 | `display: flex;` |
| `flex-direction` | 주 축 방향 지정 | `row`, `row-reverse`, `column`, `column-reverse` |
| `flex-wrap` | 줄바꿈 허용 여부 | `nowrap`(기본), `wrap`, `wrap-reverse` |
| `justify-content` | 주 축 정렬 | `flex-start`, `center`, `space-between` 등 |
| `align-items` | 교차 축 정렬 | `stretch`, `center`, `flex-start`, `flex-end` |
| `align-content` | 여러 줄 정렬 (줄 바꿈 있을 때만 적용) | `space-between`, `center`, etc. |

---

## 4️⃣ 아이템용 속성

| 속성 | 설명 | 예시 |
|------|------|------|
| `flex` | 성장/축소/기본크기 단축 속성 | `flex: 1` = `flex-grow: 1; flex-shrink: 1; flex-basis: 0;` |
| `flex-grow` | 남은 공간을 얼마나 차지할지 | `flex-grow: 2` → 2배 더 큼 |
| `flex-shrink` | 공간 부족 시 줄어드는 비율 | `flex-shrink: 0` → 줄어들지 않음 |
| `flex-basis` | 기본 크기 설정 | `flex-basis: 200px` |
| `align-self` | 특정 아이템만 교차 축 정렬 다르게 | `align-self: center` |
| `order` | 아이템 순서 변경 | `order: 1`, `order: -1` 등 |

---

## 5️⃣ 예제

```html
<div class="container">
  <div class="item">A</div>
  <div class="item">B</div>
  <div class="item">C</div>
</div>
```

```css
.container {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
}

.item {
  flex: 1;
  padding: 10px;
  background-color: #eee;
}
```

---

## 🎯 정리

✔ **display: flex** → 부모 요소를 flex 컨테이너로 만듦  
✔ **가로/세로 정렬, 간격 배분**이 쉬워짐  
✔ 1차원 레이아웃에서 **아이템 정렬을 유연하게 제어**  
✔ 자식 요소의 **크기 비율, 순서, 정렬 방식**까지 제어 가능  
