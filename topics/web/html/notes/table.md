# 📊 HTML - 테이블(Table) 태그

HTML에서 **표(테이블)** 는 데이터를 행과 열로 정리해서 보여줄 때 사용합니다.

---

## 1️⃣ 기본 구조

```html
<table>
  <thead>
    <tr>
      <th>이름</th>
      <th>나이</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>홍길동</td>
      <td>25</td>
    </tr>
    <tr>
      <td>김영희</td>
      <td>30</td>
    </tr>
  </tbody>
</table>
```

---

## 2️⃣ 주요 태그 설명

| 태그 | 설명 |
|------|------|
| `<table>` | 표의 전체 컨테이너 |
| `<thead>` | 표의 헤더 영역 (한 번만 사용) |
| `<tbody>` | 실제 데이터 영역 |
| `<tfoot>` | 표의 하단 요약 영역 (합계 등) |
| `<tr>` | 한 행(row)을 의미 |
| `<th>` | 열 제목(헤더 셀), 기본적으로 **굵게 & 가운데 정렬** |
| `<td>` | 일반 데이터 셀 (table data) |

---

## 3️⃣ 속성 (HTML or CSS에서 사용 가능)

| 속성/스타일 | 설명 | 예시 |
|-------------|------|------|
| `border` | 테두리 (CSS로 설정 권장) | `style="border: 1px solid"` |
| `colspan` | 열 병합 (가로 병합) | `<td colspan="2">` |
| `rowspan` | 행 병합 (세로 병합) | `<td rowspan="2">` |
| `width`, `height` | 셀 크기 조정 | `<td width="100">` |

---

## 4️⃣ `colspan`, `rowspan` 병합 예제

```html
<table border="1">
  <tr>
    <th rowspan="2">이름</th>
    <th colspan="2">연락처</th>
  </tr>
  <tr>
    <td>이메일</td>
    <td>전화번호</td>
  </tr>
  <tr>
    <td>홍길동</td>
    <td>hong@example.com</td>
    <td>010-1234-5678</td>
  </tr>
</table>
```

---

## 5️⃣ 스타일링 예시 (CSS)

```html
<table style="border-collapse: collapse; width: 100%;">
  <thead>
    <tr style="background-color: #f2f2f2;">
      <th>과목</th>
      <th>점수</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>수학</td>
      <td>95</td>
    </tr>
    <tr>
      <td>영어</td>
      <td>88</td>
    </tr>
  </tbody>
</table>
```

---

## 6️⃣ 표 제목 `<caption>`

`<caption>` 태그를 사용하면 표 제목도 달 수 있습니다.

#### 예제

```html
<caption>성적표</caption>
```

---

## 🎯 정리

| 태그 | 설명 |
|------|------|
| `<table>` | 전체 표 영역 |
| `<tr>` | 행(Row) |
| `<th>` | 제목 셀 (굵고 가운데 정렬) |
| `<td>` | 일반 데이터 셀 |
| `<thead>` | 헤더 영역 |
| `<tbody>` | 데이터 본문 영역 |
| `<tfoot>` | 요약/푸터 영역 |
| `colspan` / `rowspan` | 셀 병합에 사용 |

✔ **표를 그릴 때는 항상 구조를 먼저 생각**하고 `<thead>`, `<tbody>`를 나눠주는 것이 좋습니다.
