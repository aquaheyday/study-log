# 🔄 Go 언어 JSON 처리 (encoding/json)

Go는 표준 패키지 `encoding/json`을 통해 **JSON 데이터를 손쉽게 인코딩(직렬화)** 하거나 **디코딩(파싱)** 할 수 있습니다.

---

## 1️⃣ 구조체 → JSON (Marshal)

```go
package main

import (
    "encoding/json"
    "fmt"
)

type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    u := User{Name: "Alice", Age: 30}
    jsonBytes, _ := json.Marshal(u)
    fmt.Println(string(jsonBytes))
}
```

✔ `json.Marshal(v)`는 구조체를 JSON 문자열로 변환  
✔ 반환값은 `[]byte` → `string()`으로 출력  

---

## 2️⃣ JSON → 구조체 (Unmarshal)

```go
data := []byte(`{"name": "Bob", "age": 25}`)

var u User
err := json.Unmarshal(data, &u)
fmt.Println(u.Name, u.Age)
```

✔ `json.Unmarshal(data, &obj)`  
✔ JSON 키는 구조체 태그(`json:"key"`)와 일치해야 함  

---

## 3️⃣ JSON 필드 태그

```go
type Product struct {
    ID    int    `json:"id"`
    Name  string `json:"name,omitempty"` // 값 없으면 생략
    Price int    `json:"-"`              // JSON에서 제외
}
```

| 태그        | 의미                     |
|-------------|--------------------------|
| `json:"name"`       | 필드 이름 지정            |
| `json:"name,omitempty"` | 값이 없으면 생략         |
| `json:"-"`          | JSON 변환에서 무시        |

---

## 4️⃣ JSON 인덴트 출력 (`MarshalIndent`)

```go
b, _ := json.MarshalIndent(u, "", "  ")
fmt.Println(string(b))
```

✔ 보기 좋게 들여쓰기된 JSON 생성  
✔ 첫 번째 인자는 prefix, 두 번째는 들여쓰기 문자  

---

## 5️⃣ 익명 구조체 사용

```go
data := struct {
    Title string `json:"title"`
    Done  bool   `json:"done"`
}{
    Title: "할 일",
    Done:  true,
}

jsonStr, _ := json.Marshal(data)
fmt.Println(string(jsonStr))
```

✔ 임시적으로 JSON 만들 때 편리  

---

## 6️⃣ 맵 또는 인터페이스로 유연하게 처리

```go
var m map[string]interface{}
json.Unmarshal([]byte(`{"id": 1, "active": true}`), &m)
fmt.Println(m["id"], m["active"])
```

✔ 구조체가 아닌 **동적 JSON 처리**가 필요할 때  
✔ 타입 단언 또는 변환 필요 (`value.(string)`, `value.(float64)` 등)

---

## 7️⃣ 배열/슬라이스 처리

```go
var names []string
json.Unmarshal([]byte(`["Alice", "Bob"]`), &names)
fmt.Println(names)
```

✔ 슬라이스, 배열도 JSON 배열로 직렬화/파싱 가능  

---

## 8️⃣ 중첩 구조체 처리

```go
type Post struct {
    Title string `json:"title"`
    Author struct {
        Name string `json:"name"`
    } `json:"author"`
}
```

✔ 구조체 안에 구조체를 중첩해서 사용 가능  
✔ JSON도 중첩된 구조로 표현됨  

---

## 9️⃣ 디코딩 스트림 방식 (Decoder)

```go
r := strings.NewReader(`{"name":"Eve","age":22}`)
decoder := json.NewDecoder(r)

var u User
decoder.Decode(&u)
```

✔ 큰 JSON을 스트림처럼 읽고 싶을 때 사용  
✔ `io.Reader` 인터페이스 사용 가능 (`os.File`, `http.Body`, 등)

---

## 🔟 인코딩 스트림 방식 (Encoder)

```go
encoder := json.NewEncoder(os.Stdout)
encoder.SetIndent("", "  ")
encoder.Encode(u)
```

✔ JSON을 스트림에 바로 출력 가능  
✔ 파일, 네트워크 응답 등과 연결해서 사용  

---

## 🎯 정리

✔ JSON 키가 구조체 필드명과 다르면 반드시 태그(`json:"..."`) 추가  
✔ 값이 없는 필드 제외하려면 `omitempty` 사용  
✔ 숫자는 `Unmarshal` 시 `float64`로 들어오는 경우 많음 (형변환 필요)  
✔ API 응답 처리할 땐 `Decoder`, 내부 처리할 땐 `Unmarshal`이 편함  
✔ JSON으로 구조체 초기화 시에도 태그는 적용됨
✔ Go는 JSON 직렬화/역직렬화가 매우 빠르고 안전합니다.    
