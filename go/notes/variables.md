# 🔤 Go 언어 변수와 데이터 타입

Go는 **정적 타입 언어(static typed language)**로, 변수를 선언할 때 타입을 명시하거나 컴파일러가 자동 추론합니다.

---

## 1️⃣ 변수 선언 방식

### 1) 기본 선언

```go
var name string = "Gopher"
var age int = 30
```

---

### 2) 타입 생략 (타입 추론)

```go
var city = "Seoul"  // 타입 자동 추론 → string
```

---

### 3) 짧은 선언 (함수 내에서만 가능)

```go
message := "Hello, Go!"  // 타입 추론 + 선언 + 초기화
```

✔ `:=`은 **함수 내에서만 사용 가능**, 전역에서는 `var`만 사용 가능

---

## 2️⃣ 상수 선언 (`const`)

```go
const pi = 3.14
const appName string = "MyApp"
```

✔ `const`는 **컴파일 타임에 값이 결정되는 상수**를 선언할 때 사용  
✔ 상수는 **재할당 불가**

---

## 3️⃣ 주요 데이터 타입

### 1) 숫자형

| 타입 | 설명 |
|------|------|
| `int`, `int8`, `int16`, `int32`, `int64` | 정수형 |
| `uint`, `uint8`, `uint16`, `uint32`, `uint64` | 부호 없는 정수형 |
| `float32`, `float64` | 실수형 |
| `complex64`, `complex128` | 복소수형 |

```go
var a int = 10
var b float64 = 3.14
```

---

### 2) 문자열

```go
var s string = "Hello"
s = s + " World"
fmt.Println(len(s))        // 문자열 길이
fmt.Println(s[0])          // 문자열 인덱스 접근 → byte 값
```

✔ Go의 문자열은 **UTF-8 인코딩**, 불변(immutable)  
✔ **`rune` 타입**을 사용하면 유니코드 문자 처리 가능

---

### 3) 불리언 (`bool`)

```go
var isActive bool = true
```

✔ `true` / `false` 값만 가질 수 있음  
✔ `if`, `for` 등에서 조건식으로 사용

---

### 4) 배열과 슬라이스

```go
var arr [3]int = [3]int{1, 2, 3}
s := []string{"Go", "Rust", "Python"}
```

| 타입 | 설명 |
|------|------|
| `[n]T` | 고정 길이 배열 |
| `[]T` | 가변 길이 슬라이스 |

✔ 슬라이스는 `append()`, `copy()` 등으로 유연하게 사용 가능

---

### 5) 맵 (Map)

```go
m := map[string]int{
    "apple":  3,
    "banana": 5,
}
fmt.Println(m["apple"]) // 3
```

✔ 키-값 형태의 자료구조  
✔ `make(map[string]int)` 으로도 생성 가능

---

### 6) 구조체 (Struct)

```go
type User struct {
    Name string
    Age  int
}

user := User{Name: "Alice", Age: 28}
```

✔ 구조체는 객체 지향처럼 **데이터 그룹화**에 사용

---

## 4️⃣ 타입 변환 (Type Conversion)

Go는 **암시적 형 변환을 지원하지 않으며**, **명시적 타입 변환**을 해야 합니다.

```go
var a int = 10
var b float64 = float64(a)

var s string = strconv.Itoa(a)  // int → string
```

✔ 문자열 ↔ 숫자 변환은 `strconv` 패키지 사용

---

## 5️⃣ Zero Value (초기값)

Go는 변수 선언 시 값을 지정하지 않으면 **타입에 따라 기본값(zero value)**이 설정됩니다.

| 타입 | 기본값 |
|------|--------|
| 숫자형 | `0` |
| 문자열 | `""` |
| bool | `false` |
| 포인터, 슬라이스, 맵 등 | `nil` |

---

## 🎯 정리

✔ Go는 정적 타입 언어로, 변수는 `var`, `:=`, `const`로 선언합니다.  
✔ 숫자, 문자열, 불리언, 배열, 맵, 구조체 등 다양한 기본 타입을 제공합니다.  
✔ 자동 타입 추론은 편리하지만, 명시적 타입 지정이 더 명확한 경우도 많습니다.  
✔ 타입 변환은 항상 **명시적으로 수행**해야 하며, `strconv` 패키지가 자주 사용됩니다.

