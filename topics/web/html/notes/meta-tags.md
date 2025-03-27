# 📘 HTML - 메타 태그(meta tag)

**메타 태그**는 HTML 문서의 `<head>` 안에 작성되며, 문서에 대한 **정보(메타데이터)를 브라우저, 검색 엔진, SNS 등에 전달**하는 역할을 합니다.

---

## 1️⃣ 메타 태그 기본 구조

```html
<head>
  <meta charset="UTF-8">
  <meta name="description" content="이 페이지는 HTML 메타 태그를 설명합니다.">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>
```

---

## 2️⃣ 주요 메타 태그 목록

| 태그 | 설명 | 예시 |
|------|------|------|
| `charset` | 문자 인코딩 설정 | `<meta charset="UTF-8">` |
| `viewport` | 모바일 화면 대응 | `<meta name="viewport" content="width=device-width, initial-scale=1.0">` |
| `description` | 페이지 설명 (SEO에 중요) | `<meta name="description" content="페이지 요약">` |
| `keywords` | 페이지 키워드 (요즘은 SEO 영향 ↓) | `<meta name="keywords" content="HTML, meta, SEO">` |
| `author` | 작성자 정보 | `<meta name="author" content="홍길동">` |
| `robots` | 검색 엔진 색인 설정 | `<meta name="robots" content="index, follow">` |
| `http-equiv` | HTTP 응답 헤더 설정 | `<meta http-equiv="refresh" content="5;url=next.html">` |
| `theme-color` | 모바일 브라우저 툴바 색상 설정 (PWA 등) | `<meta name="theme-color" content="#ffffff">` |

---

## 3️⃣ Open Graph (OG) 메타 태그 (SNS 공유용)

- 페이스북, 카카오톡, 디스코드 등에서 **링크 미리보기용**
- `<meta property="og:...">` 형태로 사용

```html
<meta property="og:title" content="문서 제목">
<meta property="og:description" content="문서 설명">
<meta property="og:image" content="https://example.com/image.jpg">
<meta property="og:url" content="https://example.com/page.html">
```

---

## 4️⃣ Twitter Card 메타 태그

- 트위터에 공유될 때 표시될 정보 지정

```html
<meta name="twitter:card" content="summary_large_image">
<meta name="twitter:title" content="문서 제목">
<meta name="twitter:description" content="문서 설명">
<meta name="twitter:image" content="https://example.com/image.jpg">
```

---

## 5️⃣ 메타 태그 활용 예시

```html
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta name="description" content="웹 접근성과 HTML 구조를 설명하는 문서입니다.">
  <meta name="author" content="홍길동">
  <meta name="robots" content="index, follow">

  <meta property="og:title" content="접근성과 메타 정보">
  <meta property="og:description" content="HTML의 접근성과 메타 정보를 정리한 문서">
  <meta property="og:image" content="https://example.com/thumb.png">
</head>
```

---

## 6️⃣ 기타 참고

| 설정 목적 | 메타 태그 예시 |
|-----------|----------------|
| 5초 뒤 다른 페이지로 이동 | `<meta http-equiv="refresh" content="5;url=next.html">` |
| 웹 앱 툴바 색상 지정 | `<meta name="theme-color" content="#222222">` |
| 검색엔진 색인 금지 | `<meta name="robots" content="noindex, nofollow">` |

---

## 🎯 정리

✔ `<meta>` 태그는 **문서 정보 전달**에 사용됨 (브라우저, 검색엔진, SNS 등)  
✔ SEO, 반응형 웹, 접근성, SNS 공유 등에 필수  
✔ `<meta>`는 **빈 태그**이며 항상 `<head>` 안에서 사용됨  
