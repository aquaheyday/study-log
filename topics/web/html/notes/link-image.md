# 🔗 HTML - 링크 & 이미지 태그

HTML에서 **링크와 이미지**는 웹 콘텐츠의 핵심입니다  
- `<a>`: 하이퍼링크(클릭 가능한 텍스트/이미지)
- `<img>`: 이미지를 표시할 때 사용

---

## 1️⃣ `<a>` - 링크(Anchor) 태그

웹 페이지 간 이동, 파일 다운로드, 외부 사이트 연결 등  
**모든 하이퍼링크**는 `<a>` 태그로 구성됩니다.

```html
<a href="https://example.com">사이트 방문</a>
```

### 1) 주요 속성

| 속성 | 설명 | 예시 |
|------|------|------|
| `href` | 이동할 URL 또는 파일 경로 | `href="page.html"` |
| `target` | 링크 열기 방식 (`_blank`: 새 창) | `target="_blank"` |
| `download` | 클릭 시 다운로드 (파일 저장) | `download="file.pdf"` |
| `rel` | 링크 간의 관계 명시 (SEO, 보안) | `rel="noopener noreferrer"` |

✔ 외부 링크는 항상 `target="_blank"` + `rel="noopener noreferrer"`로 보안 강화  

---

### 2) 예제

```html
<a href="https://naver.com" target="_blank" rel="noopener noreferrer">
  네이버 새창 열기
</a>

<a href="/files/report.pdf" download>PDF 다운로드</a>
```

---

## 2️⃣ `<img>` - 이미지 태그

웹 페이지에 **이미지를 삽입**할 때 사용합니다.

```html
<img src="cat.jpg" alt="고양이 사진">
```

### 1) 주요 속성

| 속성 | 설명 | 예시 |
|------|------|------|
| `src` | 이미지 파일 경로 또는 URL | `src="logo.png"` |
| `alt` | 이미지 대체 텍스트 (접근성) | `alt="회사 로고"` |
| `width`, `height` | 이미지 크기 지정 (px, %) | `width="200"` |
| `title` | 마우스 올리면 툴팁 표시 | `title="고양이"` |
| `loading` | 지연 로딩 여부 (`lazy`) | `loading="lazy"` |

✔ `alt`는 **이미지 설명**을 적는 것이지 키워드 스팸이 아님  

### 2) 예제

```html
<img src="https://via.placeholder.com/150" alt="샘플 이미지" width="150" loading="lazy">
```

---

## 3️⃣ 이미지 링크로 만들기

이미지에 링크를 걸 수도 있어요:

```html
<a href="https://example.com">
  <img src="banner.png" alt="배너 광고">
</a>
```

---

## 🎯 정리

| 태그 | 역할 | 필수 속성 |
|------|------|------------|
| `<a>` | 하이퍼링크 생성 | `href` |
| `<img>` | 이미지 삽입 | `src`, `alt` |

✔ `<a>`는 내부/외부 페이지, 파일, 이메일, 전화번호 등 모두 연결 가능  
✔ `<img>`는 콘텐츠 이미지 외에도 아이콘, 배너 등 다양한 용도로 사용  
✔ `alt`는 접근성과 SEO에 매우 중요
