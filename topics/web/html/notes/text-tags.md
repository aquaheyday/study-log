# ✏️ HTML - 텍스트 관련 태그

HTML은 텍스트 콘텐츠를 **구조화하고 의미를 부여하는** 다양한 태그를 제공합니다.  

---

## 1️⃣ 텍스트 구조 태그

| 태그 | 설명 | 예시 |
|------|------|------|
| `<h1>` ~ `<h6>` | 제목(Heading) 태그, 숫자가 작을수록 중요도 높음 | `<h1>제목</h1>` |
| `<p>` | 문단(Paragraph) | `<p>문단입니다.</p>` |
| `<br>` | 줄바꿈(Line break) | `줄1<br>줄2` |
| `<hr>` | 수평선(Horizontal Rule) | `<hr>` |

---

## 2️⃣ 텍스트 강조 및 의미 태그

| 태그 | 설명 | 예시 |
|------|------|------|
| `<strong>` | 중요한 텍스트, 기본적으로 **굵게** 표시 | `<strong>중요</strong>` |
| `<em>` | 강조된 텍스트, 기본적으로 *기울임* 표시 | `<em>강조</em>` |
| `<b>` | 굵은 텍스트 (시각적 강조, 의미 없음) | `<b>굵게</b>` |
| `<i>` | 기울임 텍스트 (의미 없이 스타일만) | `<i>기울임</i>` |
| `<mark>` | 하이라이트 표시 (형광펜 효과) | `<mark>강조된 텍스트</mark>` |
| `<small>` | 작은 글씨 | `<small>작은 텍스트</small>` |
| `<sub>` | 아래첨자 (subscript) | `H<sub>2</sub>O` |
| `<sup>` | 위첨자 (superscript) | `E = mc<sup>2</sup>` |
| `<u>` | 밑줄 (underline, 스타일 목적) | `<u>밑줄</u>` |
| `<del>` | 삭제된 텍스트 (취소선) | `<del>지워짐</del>` |
| `<ins>` | 삽입된 텍스트 (밑줄로 표시됨) | `<ins>추가됨</ins>` |
| `<abbr>` | 약어 설명 | `<abbr title="HyperText Markup Language">HTML</abbr>` |
| `<code>` | 코드 조각 | `<code>let x = 10;</code>` |
| `<pre>` | 공백과 줄바꿈 유지하는 코드 블록 | `<pre>코드나 텍스트 그대로 표시</pre>` |

---

## 3️⃣ 예제

```html
<h1>HTML 텍스트 태그</h1>
<p>이 문장은 <strong>중요</strong>하며 <em>강조</em>됩니다.</p>
<p>물은 H<sub>2</sub>O, 아인슈타인의 공식은 E = mc<sup>2</sup>입니다.</p>
<pre>
function greet() {
  console.log("Hello, world!");
}
</pre>
<code>const x = 5;</code>
```

---

## 🎯 정리

✔ `<h1>`~`<h6>`: 제목 구조 정의  
✔ `<p>`, `<br>`, `<hr>`: 문단과 줄바꿈, 구분선  
✔ `<strong>`, `<em>`: **의미 기반 강조**  
✔ `<b>`, `<i>`: **스타일 강조(의미 없음)**  
✔ `<mark>`, `<small>`, `<code>` 등은 특정 상황에 유용  

