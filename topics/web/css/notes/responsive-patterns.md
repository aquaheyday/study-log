# 📐 반응형 레이아웃 패턴

**반응형 웹 디자인**은 화면 크기(모바일, 태블릿, 데스크탑 등)에 따라 레이아웃이 유연하게 변화하는 웹 디자인 기법입니다.
실무에서 자주 쓰이는 **반응형 레이아웃 패턴**을 유형별로 정리합니다.

---

## 1️⃣ Fluid Grid (유동형 그리드)

- 모든 너비를 `%`나 `fr` 단위로 지정
- **뷰포트 크기에 따라 자연스럽게 늘어나고 줄어듦**

```css
.container {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 20px;
}
```

```css
@media (max-width: 768px) {
  .container {
    grid-template-columns: 1fr;
  }
}
```

---

## 2️⃣ **Column Drop (열 나누기 → 줄바꿈)**

- 데스크탑에서는 열(column) 구조  
- 모바일에서는 **수직 스택**으로 변경

```css
.layout {
  display: flex;
  gap: 20px;
}

.sidebar {
  width: 250px;
}

@media (max-width: 768px) {
  .layout {
    flex-direction: column;
  }

  .sidebar {
    width: 100%;
  }
}
```

---

## 3️⃣ Off-canvas Navigation (모바일 메뉴 숨기기)

- 데스크탑에서는 메뉴가 보임  
- 모바일에서는 숨기고 버튼으로 열고 닫음

✔ 메뉴는 `position: absolute` 또는 `fixed`로 설정  

```css
@media (max-width: 768px) {
  .nav {
    display: none;
  }

  .menu-button {
    display: block;
  }
}
```

---

## 4️⃣ Mostly Fluid (헤더/사이드바는 고정, 본문은 유동)

- `header`, `footer`, `sidebar`는 **고정 너비**
- `main`은 **flex-grow: 1** 로 유동적으로 확장

```css
.page {
  display: flex;
}

.sidebar {
  width: 300px;
}

.main {
  flex: 1;
}
```

---

## 5️⃣ Layout Shifter (구조 변경형)

- 화면 크기에 따라 레이아웃 **자체가 완전히 바뀜**
- 데스크탑: 사이드바 + 본문  
- 모바일: 상하 배치

```html
<div class="page">
  <aside class="sidebar">사이드바</aside>
  <main class="content">본문 콘텐츠</main>
</div>
```

```css
@media (max-width: 768px) {
  .page {
    flex-direction: column;
  }
}
```

---

## 6️⃣ Breakpoint-based Swapping (구간별 컴포넌트 교체)

- 특정 해상도에서 **다른 컴포넌트나 레이아웃 적용**

```html
<div class="desktop-nav">...</div>
<div class="mobile-nav">...</div>
```

```css
.mobile-nav {
  display: none;
}

@media (max-width: 768px) {
  .desktop-nav {
    display: none;
  }

  .mobile-nav {
    display: block;
  }
}
```

---

## 7️⃣ 반응형 디자인 팁

#### 뷰포트 설정 필수
  ```html
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  ```

#### 미디어쿼리 기준 예시

| 기기 | 기준 해상도 |
|------|-------------|
| 모바일 | ~767px |
| 태블릿 | 768px ~ 1023px |
| 데스크탑 | 1024px 이상 |

✔ `rem`, `%`, `fr`, `vw` 등 **상대 단위** 사용 권장  

---

## 🎯 정리

| 패턴 | 설명 |
|------|------|
| Fluid Grid | 유동적인 열 너비 사용 |
| Column Drop | 수평 → 수직 구조 변경 |
| Off-canvas Menu | 모바일에서 메뉴 숨기고 토글 |
| Mostly Fluid | 일부 고정, 나머지 유동 |
| Layout Shifter | 전체 레이아웃 구조 변경 |
| Breakpoint Swapping | 특정 뷰포트에서 다른 UI 적용 |

