# 🖼️ Go 언어 HTML 템플릿

Go는 `html/template` 패키지를 통해 **동적 HTML 페이지**를 안전하고 쉽게 생성할 수 있습니다.  
HTML 템플릿은 서버 사이드 렌더링이 필요한 웹 앱에서 자주 사용됩니다.

---

## 1️⃣ 기본 템플릿 렌더링

```go
package main

import (
    "html/template"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.New("index").Parse("<h1>Hello, {{.}}!</h1>"))
    tmpl.Execute(w, "Go 사용자")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

✔ `template.New().Parse()` 로 템플릿 생성  
✔ `Execute(w, data)` 로 HTML에 데이터 바인딩  

---

## 2️⃣ 파일로 템플릿 분리하기

```go
// templates/index.html
<h1>Hello, {{.Name}}!</h1>
<p>Your age is {{.Age}}.</p>
```

```go
type User struct {
    Name string
    Age  int
}

func handler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/index.html"))
    data := User{Name: "Go", Age: 10}
    tmpl.Execute(w, data)
}
```

✔ 템플릿 파일을 따로 분리해 관리  
✔ 구조체 데이터를 넘겨서 `.필드명` 으로 접근  

---

## 3️⃣ 조건문 (`if`) 사용

```html
{{if .LoggedIn}}
    <p>Welcome back!</p>
{{else}}
    <p>Please log in.</p>
{{end}}
```

✔ `.LoggedIn` 값이 true일 경우만 출력  
✔ `else`, `end` 필수  

---

## 4️⃣ 반복문 (`range`) 사용

```html
<ul>
{{range .Items}}
    <li>{{.}}</li>
{{end}}
</ul>
```

```go
data := struct {
    Items []string
}{
    Items: []string{"Apple", "Banana", "Cherry"},
}
```

✔ `range`는 슬라이스, 배열, 맵 반복 가능  

---

## 5️⃣ 템플릿 구성 요소 나누기 (`define` / `template`)

```html
<!-- templates/base.html -->
<html>
  <body>
    {{template "content" .}}
  </body>
</html>

<!-- templates/content.html -->
{{define "content"}}
    <h1>{{.Title}}</h1>
{{end}}
```

```go
tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/content.html"))
tmpl.Execute(w, struct{ Title string }{"Dynamic Title"})
```

✔ `define` 블록은 **부분 템플릿**으로 정의  
✔ `template` 호출로 **중첩 렌더링** 가능  

---

## 6️⃣ HTML 이스케이프 처리 (XSS 방지)

```go
template.New("safe").Parse("입력: {{.}}")
```

```go
tmpl.Execute(w, "<script>alert('XSS')</script>")
```

✔ 기본적으로 HTML 특수문자는 **자동 이스케이프됨**  
✔ `<script>` 태그는 출력 시 `<script>` → `&lt;script&gt;` 로 변환됨  
✔ 신뢰된 HTML을 출력하려면 `template.HTML` 타입을 명시적으로 사용해야 함 (주의)

---

## 7️⃣ 에러 처리 및 템플릿 캐싱

```go
tmpl, err := template.ParseFiles("index.html")
if err != nil {
    http.Error(w, "Template error", http.StatusInternalServerError)
    return
}
tmpl.Execute(w, data)
```

✔ 템플릿 파싱 시 에러 처리를 반드시 해야 안정성 확보  
✔ 반복 렌더링 시 `template.Must()` 로 **미리 파싱해 캐싱**하는 방식 권장  

---

## 8️⃣ 템플릿 문법 요약

| 문법 | 설명 |
|------|------|
| `{{.}}` | 현재 값 출력 |
| `{{.Field}}` | 구조체 필드 출력 |
| `{{if .Cond}}...{{end}}` | 조건문 |
| `{{range .List}}...{{end}}` | 반복문 |
| `{{template "name" .}}` | 부분 템플릿 삽입 |
| `{{define "name"}}...{{end}}` | 템플릿 정의 |

---

## 9️⃣ 템플릿 디렉토리 구성 예시

```
project/
├── main.go
└── templates/
    ├── base.html
    ├── index.html
    └── layout.html
```

✔ `templates` 폴더로 파일 분리 관리  
✔ 여러 템플릿을 조합할 땐 `template.ParseFiles(...)` 사용  

---

## 🎯 정리

✔ Go의 `html/template`은 **자동 이스케이프 + 타입 안전성** 제공  
✔ 템플릿 분리는 구조화에 필수 (base + content 구조)  
✔ XSS 방지 기본 내장 (직접 HTML 넣을 땐 `template.HTML` 주의)  
✔ 반복, 조건, 중첩 템플릿 등 기능은 강력하지만 문법은 간결함  
✔ 실제 HTML 렌더링뿐 아니라 이메일 템플릿, 문서 생성에도 활용 가능  
