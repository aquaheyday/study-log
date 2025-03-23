# 📦 Go 언어 배열, 슬라이스, 맵

Go는 배열(Array), 슬라이스(Slice), 맵(Map)을 사용하여 효율적인 자료 구조를 구성할 수 있습니다.

---

## 1️⃣ 배열 (Array)

### 1) 배열 기본 문법

```go
var arr [3]int = [3]int{10, 20, 30}
fmt.Println(arr[0]) // 10
```

✔ 고정 길이 배열
✔ 크기(`[3]`)가 타입에 포함되므로 크기가 다르면 다른 타입
✔ 선언된 크기만큼만 사용 가능

### 2) 길이 및 순회

```go
fmt.Println(len(arr)) // 3

for i, v := range arr {
    fmt.Println(i, v)
}
```

---

## 2️⃣ 슬라이스 (Slice)

### 1) 슬라이스 생성

```go
s := []int{1, 2, 3}             // 리터럴
s2 := make([]int, 3)           // 길이 3의 슬라이스 (값은 0)
s3 := arr[1:3]                 // 배열에서 추출 (슬라이싱)
```

✔ 크기 가변적, 내부적으로 배열을 참조
✔ `append()`로 동적 크기 증가 가능

### 2) `append`로 요소 추가

```go
s := []int{1, 2}
s = append(s, 3, 4)
```

✔ `append`는 새 슬라이스를 반환하므로 **반드시 다시 대입**

---

### 3) copy로 복사

```go
src := []int{1, 2, 3}
dst := make([]int, len(src))
copy(dst, src)
```

✔ `copy(dest, src)`는 슬라이스 간 데이터를 복사

---

### 4) 슬라이스의 내부 구조

슬라이스는 배열의 참조 구조이며, `len`과 `cap`을 가짐

```go
s := []int{1, 2, 3}
fmt.Println(len(s)) // 3
fmt.Println(cap(s)) // 3
```

---

## 3️⃣ 맵 (Map)

### 1) 맵 선언 및 초기화

```go
m := map[string]int{
    "apple":  3,
    "banana": 5,
}
```

또는

```go
m := make(map[string]int)
m["orange"] = 4
```

---

### 2) 요소 접근 및 삭제

```go
fmt.Println(m["apple"]) // 3

val, ok := m["kiwi"]
if ok {
    fmt.Println("존재하는 키:", val)
} else {
    fmt.Println("없는 키")
}

delete(m, "banana") // 삭제
```

✔ 키가 없으면 zero value 반환
✔ `val, ok := m[key]` → 존재 여부 체크 (실무에서 중요)

---

### 3) 맵 순회

```go
for key, value := range m {
    fmt.Println(key, value)
}
```

✔ 순서는 **보장되지 않음** (랜덤 순서)

---

### 4) 중첩 맵 예시

```go
users := map[string]map[string]string{
    "alice": {
        "email": "alice@example.com",
        "role":  "admin",
    },
}
```

---

## 4️⃣ 배열 vs 슬라이스 차이점

| 항목        | 배열 `[N]T`            | 슬라이스 `[]T`         |
|-------------|------------------------|-------------------------|
| 크기        | 고정                   | 가변                    |
| 용도        | 정해진 크기의 자료 저장 | 동적 목록               |
| 성능        | 더 빠름                | 메모리 효율적, 유연함  |
| 사용 빈도   | 낮음                   | 매우 높음               |

---

## 🎯 정리

✔ 배열은 고정 길이, 슬라이스는 동적 크기 자료 구조입니다.  
✔ 슬라이스는 `append`, `copy`, `make` 등을 통해 유연하게 사용됩니다.  
✔ 맵은 키-값 구조로 동적으로 요소를 추가하거나 삭제할 수 있으며, `ok` 패턴으로 안전하게 키 존재 여부를 확인하는 것이 중요합니다.

