# 🎬 CSS - 애니메이션 & 트랜지션

CSS에서 **요소의 스타일이 시간에 따라 변화하도록** 만들려면 `transition`과 `animation` 속성을 사용합니다.

---

## 1️⃣ CSS 트랜지션 (`transition`)

### 1) 개념  
요소의 속성이 **변경될 때 부드럽게 전환**되도록 만드는 속성

### 2) 기본 문법

```css
.element {
  transition: 속성 지속시간 타이밍함수 지연시간;
}
```

### 3) 주요 속성

| 속성 | 설명 | 예시 |
|------|------|------|
| `transition-property` | 어떤 속성을 변화시킬지 | `all`, `background-color` 등 |
| `transition-duration` | 변화 시간 | `0.3s`, `500ms` |
| `transition-timing-function` | 속도 곡선 (처음엔 느리게 → 점점 빨라지게 등) | `ease`, `linear`, `ease-in`, `ease-out`, `cubic-bezier(...)` |
| `transition-delay` | 시작 지연 시간 | `0.2s` 등 |

### 4) 예제

```css
.button {
  background-color: blue;
  transition: background-color 0.3s ease;
}

.button:hover {
  background-color: red;
}
```

---

## 2️⃣ CSS 애니메이션 (`animation`)

### 1) 개념  
**복잡한 연속 동작**을 만들 수 있는 속성, 키프레임(`@keyframes`)을 정의해서 여러 단계를 지정 가능

### 2) 기본 문법

```css
@keyframes 이름 {
  0%   { 스타일A }
  100% { 스타일B }
}

.element {
  animation: 이름 지속시간 타이밍함수 지연 반복 횟수 방향 채우기;
}
```

### 3) 주요 속성

| 속성 | 설명 | 예시 |
|------|------|------|
| `animation-name` | 사용할 키프레임 이름 | `slide-in` |
| `animation-duration` | 지속 시간 | `2s`, `1.5s` |
| `animation-timing-function` | 속도 곡선 | `ease`, `linear` 등 |
| `animation-delay` | 지연 시간 | `1s` |
| `animation-iteration-count` | 반복 횟수 | `1`, `infinite` |
| `animation-direction` | 방향 | `normal`, `reverse`, `alternate` |
| `animation-fill-mode` | 종료 후 상태 유지 | `none`, `forwards`, `backwards`, `both` |

### 4) 예제

```css
@keyframes fade-in {
  from { opacity: 0; }
  to   { opacity: 1; }
}

.box {
  animation: fade-in 1s ease-in-out forwards;
}
```

---

## 3️⃣ 트랜지션 vs 애니메이션 차이

| 항목 | 트랜지션 (transition) | 애니메이션 (animation) |
|------|------------------------|--------------------------|
| 작동 조건 | **상태 변화(trigger)** 필요 (예: hover) | 자동 실행 가능 |
| 복잡한 단계 | 불가능 (시작-끝만) | 가능 (`@keyframes` 이용) |
| 반복 | ❌ 불가 | ✅ 반복 가능 (`infinite`) |
| 제어력 | 간단 | 복잡한 동작 구현 가능 |
| 사용 예 | hover 효과, 버튼 클릭 시 변화 | 배너 슬라이드, 로딩 애니메이션 등 |

---

## 4️⃣ 버튼 호버 + 자동 애니메이션 예시

```css
@keyframes bounce {
  0%, 100% { transform: translateY(0); }
  50%      { transform: translateY(-10px); }
}

.button {
  background: #4CAF50;
  color: white;
  padding: 12px 24px;
  transition: background-color 0.3s ease;
  animation: bounce 2s ease-in-out infinite;
}

.button:hover {
  background-color: #388E3C;
}
```

---

## 🎯 정리

✔ **짧은 효과**: `transition` (hover, focus 등)  
✔ **복잡한 연속 동작**: `animation`  
✔ `animation-fill-mode: forwards` → 애니메이션 끝난 후 상태 유지  
✔ `animation-play-state` → 애니메이션 일시정지 / 재생 제어 가능  

