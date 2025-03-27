# 🧱 CSS - z-index & Stacking

**`z-index`** 는 HTML 요소들이 **겹칠 때 어떤 요소가 위에 올지**를 결정하는 속성입니다.  
이는 **"스택(stack)"**, 즉 **쌓이는 순서**를 제어하는 도구입니다.

---

## 1️⃣ 기본 개념

- 기본적으로 HTML 요소는 **나중에 작성된 요소가 위에** 올라옴
- `position`이나 `display`, `transform` 등이 설정되면 **쌓임 맥락(Stacking Context)** 이 생기고 그 안에서 `z-index`로 **우선순위**를 조절할 수 있음

---

## 2️⃣ z-index 문법

- 숫자가 클수록 **더 위에 위치**
- **기본값은 `auto`** (부모 맥락 따라감)

```css
.element {
  position: relative; /* 또는 absolute, fixed, sticky */
  z-index: 10;
}
```

---

## 3️⃣ stacking context(쌓임 맥락)이란?

겹침 순서를 관리하는 **독립된 레이어 그룹** `z-index`는 **같은 stacking context 안에서만 비교됨**

---

## 4️⃣ stacking context가 생기는 조건

| 조건 | 예시 |
|------|------|
| `position: relative / absolute / fixed / sticky` + `z-index` 지정 | ✅ |
| `opacity < 1` | ✅ |
| `transform` 사용 (`transform: scale(1)`) | ✅ |
| `filter`, `perspective`, `will-change` 등 | ✅ |
| `isolation: isolate` | ✅ |

✔ 새로운 stacking context 안에서는 **z-index가 외부와 비교되지 않음**  

---

## 5️⃣ 단순 z-index 예제

```html
<div class="box1">Box 1</div>
<div class="box2">Box 2</div>
```

```css
.box1 {
  position: absolute;
  z-index: 1;
}

.box2 {
  position: absolute;
  z-index: 2; /* 더 위에 표시됨 */
}
```

---

## 6️⃣ stacking context 내부 예제

```html
<div class="parent">
  <div class="child1">Child 1</div>
  <div class="child2">Child 2</div>
</div>
```

```css
.parent {
  position: relative;
  z-index: 100;
}

.child1 {
  position: absolute;
  z-index: 10;
}

.child2 {
  position: absolute;
  z-index: 20;
}
```

✔ 둘 다 **같은 부모(context)** 안에 있으므로 `child2`가 더 위에 위치함

---

## 7️⃣ 주의할 점

- `z-index`는 **position 또는 transform 등이 적용된 요소에만 동작**
- stacking context가 생성되면 **그 안에서만 z-index 우선순위 비교**
- 부모 context보다 자식의 z-index가 커도 **부모 위로 못 나감**

---

## 🎯 정리

✔ `z-index` → 요소의 **겹침 우선순위**를 정하는 속성  
✔ `z-index` 의 높은 값이 더 위에 표시됨  
✔ `z-index` 는 `position`, `transform` 등 있어야 적용됨  
✔ `stacking context` → **독립된 레이어 그룹**, 내부 요소 끼리 z-index 비교
