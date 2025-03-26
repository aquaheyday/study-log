# 🌍 HTML - 전역 속성 (Global Attributes)

**전역 속성(Global Attributes)** 이란 **모든 HTML 요소에 공통적으로 적용 가능한 속성들**을 말합니다.  
(단, 일부 특정 요소에는 무시되거나 의미가 없을 수 있음)

---

## 1️⃣ 대표적인 전역 속성 목록

| 속성명 | 설명 |
|--------|------|
| `id` | 요소의 고유 식별자 (문서 내 유일해야 함) |
| `class` | 요소에 클래스 이름을 부여 (스타일, JS에서 활용) |
| `style` | 인라인 스타일 지정 |
| `title` | 요소에 대한 툴팁(마우스 오버 시 표시) |
| `hidden` | 요소를 숨김 (화면에서 보이지 않음) |
| `tabindex` | 키보드 탭 순서 지정 |
| `contenteditable` | 요소 내용을 사용자 수정 가능하게 만듦 |
| `draggable` | 요소를 드래그 가능하게 설정 |
| `lang` | 요소 콘텐츠의 언어 설정 |
| `dir` | 텍스트 방향 설정 (ltr, rtl) |
| `data-*` | 사용자 정의 데이터 속성 (커스텀 속성) |

---

## 2️⃣ 주요 전역 속성 사용 예제

```html
<div
  id="intro"
  class="box highlight"
  style="color: blue; font-weight: bold;"
  title="소개 박스"
  tabindex="0"
  contenteditable="true"
  data-user-id="123"
>
  전역 속성 테스트
</div>
```

---

## 3️⃣ data-* 속성 (커스텀 데이터 속성)

- `data-` 접두어를 사용하여 **사용자 정의 속성**을 추가할 수 있음
- 자바스크립트에서 `dataset`으로 접근 가능

```html
<div id="item" data-id="42" data-role="admin">사용자</div>

<script>
  const el = document.getElementById("item");
  console.log(el.dataset.id);    // "42"
  console.log(el.dataset.role);  // "admin"
</script>
```

---

## 🎯 정리

✔ 전역 속성은 대부분의 HTML 요소에 공통적으로 사용 가능  
✔ `id`, `class`, `style` 은 가장 자주 사용되는 전역 속성  
✔ `data-*` 속성은 JS와 연동할 수 있는 **커스텀 속성 시스템**  
✔ 접근성, 스타일링, 동적 조작 등을 위해 필수적으로 활용됨
