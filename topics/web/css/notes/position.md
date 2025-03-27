# 📌 CSS - Position

`position` 속성은 **요소를 어떤 기준으로 어디에 배치할지**를 정의합니다.  
이를 통해 **정적인 흐름에서 벗어난 위치 지정**이 가능해집니다.

---

## 1️⃣ 주요 `position` 값 종류

| 값 | 설명 | 기준 기준 요소 |
|-----|------|----------------|
| `static` | 기본값, 위치 지정 불가 | 없음 (문서 흐름) |
| `relative` | 요소 자신을 기준으로 이동 | 자기 자신 |
| `absolute` | 가장 가까운 **position이 설정된 조상**을 기준으로 위치 지정 | 부모 요소 (or body) |
| `fixed` | **브라우저 뷰포트**를 기준으로 고정 | 뷰포트 |
| `sticky` | 스크롤에 따라 고정되는 하이브리드 | 스크롤 위치 기준 |

---

## 2️⃣ 각 `position` 설명

### 1) `static` (기본값)

- 기본값, 위치 이동 불가  
- `top`, `left` 등의 속성 **무시됨**

```css
.box {
  position: static;
}
```

---

### 2) `relative` (상대 위치)

- 원래 위치를 기준으로 **상대적으로 이동**  
- 공간은 **기존 위치 그대로 차지함**

```css
.box {
  position: relative;
  top: 10px;
  left: 20px;
}
```

---

### 3) `absolute` (절대 위치)

- **조상 중 `position`이 지정된 요소를 기준**으로 위치 결정  
- 없다면 `body` 기준  
- 원래 위치 공간은 **제거됨**

```css
.box {
  position: absolute;
  top: 0;
  right: 0;
}
```

---

### 4) `fixed` (화면 고정)

- **뷰포트(브라우저 창)를 기준**으로 위치 고정  
- 스크롤해도 **항상 같은 위치 유지**  
- 팝업, 플로팅 버튼 등에 자주 사용

```css
.box {
  position: fixed;
  bottom: 10px;
  right: 10px;
}
```

---

### 5) `sticky` (스크롤 고정)

- **일정 위치에 도달하면 고정**되고  
- 그 전에는 `relative`처럼 작동  
- 스크롤 가능한 부모 요소가 있어야 정상 작동

```css
.box {
  position: sticky;
  top: 0;
}
```

---

## 3️⃣ `top` / `left` / `right` / `bottom`

`static`을 제외한 다른 position 값에서는 속성을 이용해 위치를 조정할 수 있음

```css
top: 20px;    /* 위에서 20px 떨어짐 */
left: 10%;    /* 왼쪽에서 10% 떨어짐 */
right: 0;     /* 오른쪽 끝에 붙음 */
bottom: 30px; /* 아래에서 30px 떨어짐 */
```

---

## 🎯 정리

✔ `static`: 기본값, 위치 이동 안 됨  
✔ `relative`: 기존 위치 기준으로 살짝 이동  
✔ `absolute`: 부모 요소 기준 위치 (레이아웃에서 빠짐)  
✔ `fixed`: 브라우저 기준 고정 (스크롤 무시)  
✔ `sticky`: 스크롤에 따라 고정되는 하이브리드  
