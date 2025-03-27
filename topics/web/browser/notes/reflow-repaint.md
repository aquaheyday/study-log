# 🔄 Browser - Reflow & Repaint

웹 브라우저는 화면에 요소를 렌더링할 때 **스타일 변경에 따라 레이아웃과 시각적 요소를 다시 계산하거나 그리게 됩니다.**  
이때 발생하는 작업이 바로 **Reflow** 와 **Repaint**입니다.

---

## 1️⃣ Reflow (리플로우)

요소의 **위치, 크기, 레이아웃**에 변경이 생겼을 때, 브라우저가 **레이아웃을 다시 계산하는 과정**

#### 대표 발생 원인

- DOM 추가/삭제/이동
- `width`, `height`, `margin`, `padding`, `border`, `position`, `display` 등 변경
- 창 크기 변경 (resize)
- 폰트 크기 변경, 이미지 로드 등으로 콘텐츠 양이 변함

#### 성능 비용
- **비쌈** → 계산된 모든 요소의 **위치와 크기 재계산**
- 경우에 따라 **전체 문서가 다시 그려짐**

---

## 2️⃣ Repaint (리페인트)

요소의 레이아웃은 그대로지만, **색상, 배경, 그림자 등 시각적 속성이 변경**되었을 때, 브라우저가 시각적 속성을 **다시 그리는 과정**

#### 대표 발생 원인

- `color`, `background-color`, `border-color`, `box-shadow` 등 스타일 변경
- `visibility`, `outline`, `opacity` 변경
- 클래스 변경으로 인해 색상만 바뀌는 경우

#### 성능 비용
- **Reflow보다 가볍지만**, 여전히 리소스를 사용함
- 여러 요소가 동시에 영향을 받을 수 있음

---

## 3️⃣ Reflow vs Repaint 비교

| 항목 | Reflow | Repaint |
|------|--------|---------|
| 의미 | 레이아웃 재계산 | 다시 그리기 |
| 발생 조건 | 위치, 크기 등 변경 | 색상, 배경 등 변경 |
| 성능 비용 | 크다 (전체 트리 영향 가능) | 중간 (그림만 다시 그림) |
| 발생 시점 | 레이아웃 계산 필요 시 | 시각적 속성만 변경 시 |
| 리페인트 발생 여부 | 항상 발생 | 리플로우 없으면 단독 발생 가능 |

✔ **Reflow가 일어나면 Repaint도 반드시 발생**함  

---

## 4️⃣ 성능 최적화 팁

#### DOM 접근 최소화

- DOM 읽기 → 쓰기 순으로 묶기 (읽기/쓰기 혼합 X)
- `documentFragment`로 한 번에 DOM 추가
- `display: none` → 수정 → 다시 표시 방식 활용

#### 애니메이션 최적화

- `transform`, `opacity` 속성만으로 애니메이션 → GPU 활용
- `top`, `left` 등은 Reflow 유발 → 피하는 것이 좋음

```css
/* ✅ 좋은 애니메이션 방식 */
.element {
  transform: translateX(100px);
  opacity: 0.5;
  transition: all 0.3s ease;
}
```

#### layout 강제 트리거 피하기

```js
// ❌ 아래와 같은 속성은 Reflow를 강제로 유발
element.offsetWidth;
element.offsetHeight;
getComputedStyle(element);
```

✔ 이런 속성들은 레이아웃 계산을 강제로 발생시킴  

---

#### 예시

```js
// ❌ 비효율적인 DOM 조작 (Reflow 100번 발생)
for (let i = 0; i < 100; i++) {
  box.style.width = i + 'px';
}

// ✅ 성능 개선: 클래스 한 번만 변경
box.classList.add('expand');
```

---

## 🎯 정리

✔ Reflow → 레이아웃 재계산 (크기/위치 변경)  
✔ Repaint → 시각적 스타일만 다시 그림  
✔ 비용 → Reflow > Repaint  
✔ 최적화 전략 → transform/opacity 사용, DOM 조작 최소화, 레이아웃 강제 트리거 방지  

