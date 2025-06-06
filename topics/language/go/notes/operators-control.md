# 🔢 Go 언어 연산자와 제어문

Go 언어는 다른 C 계열 언어들과 비슷한 연산자와 제어문을 가지고 있으며, 간결하고 직관적인 문법으로 작성할 수 있습니다.

---

## 1️⃣ 연산자 종류

### 1) 산술 연산자

| 연산자 | 설명         | 예시 (`a = 10`, `b = 3`) |
|--------|--------------|--------------------------|
| `+`    | 덧셈         | `a + b` → `13`           |
| `-`    | 뺄셈         | `a - b` → `7`            |
| `*`    | 곱셈         | `a * b` → `30`           |
| `/`    | 나눗셈       | `a / b` → `3` (정수 나눗셈) |
| `%`    | 나머지       | `a % b` → `1`            |

✔ 정수형끼리 나누면 **소수점 이하 버림** (실수로 나누려면 `float64` 변환 필요)  

---

### 2)  비교 연산자

| 연산자 | 설명          |
|--------|---------------|
| `==`   | 같음          |
| `!=`   | 같지 않음     |
| `>`    | 크다          |
| `<`    | 작다          |
| `>=`   | 크거나 같다   |
| `<=`   | 작거나 같다   |

```go
fmt.Println(a == b)  // false
fmt.Println(a >= 10) // true
```

---

### 3) 논리 연산자

| 연산자 | 설명         |
|--------|--------------|
| `&&`   | AND (그리고) |
| `||`   | OR (또는)    |
| `!`    | NOT (부정)   |

```go
if a > 5 && b < 5 {
    fmt.Println("조건 만족")
}
```

---

### 4) 대입 및 복합 대입 연산자

| 연산자 | 설명             |
|--------|------------------|
| `=`    | 대입             |
| `+=`   | 덧셈 후 대입     |
| `-=`   | 뺄셈 후 대입     |
| `*=`   | 곱셈 후 대입     |
| `/=`   | 나눗셈 후 대입   |
| `%=`   | 나머지 후 대입   |

---

### 5) 증감 연산자

| 연산자 | 설명       | 주의 |
|--------|------------|------|
| `++`   | 1 증가     | **문장(statement)로만 사용 가능** |
| `--`   | 1 감소     | 할당식 내 사용 불가 |

```go
i := 0
i++ // ✅
j := i++ // ❌ 오류: 사용 불가
```

---

### 6) 비트 연산자

| 연산자 | 설명        |
|--------|-------------|
| `&`    | AND         |
| `|`    | OR          |
| `^`    | XOR         |
| `&^`   | AND NOT     |
| `<<`   | 왼쪽 시프트 |
| `>>`   | 오른쪽 시프트 |

---

## 2️⃣ 제어문

### 1) `if` 문

```go
if score >= 90 {
    fmt.Println("A 학점")
} else if score >= 80 {
    fmt.Println("B 학점")
} else {
    fmt.Println("재시험")
}
```

✔ 조건식은 괄호 `()` 없이 작성  
✔ 중괄호 `{}`는 반드시 사용  

---

### 2) if 문 안에서 변수 선언

```go
if err := doSomething(); err != nil {
    fmt.Println("오류 발생:", err)
}
```

---

### 3) `switch` 문

```go
switch day {
case "Mon":
    fmt.Println("월요일")
case "Tue":
    fmt.Println("화요일")
default:
    fmt.Println("기타")
}
```

✔ `break` 없이도 자동으로 빠져나옴 (C 언어와 다름)  
✔ 여러 조건 묶기 가능: `case "Mon", "Tue":`  

---

### 4) `switch type` 문 (인터페이스 타입 비교)

```go
var x interface{} = "hello"

switch v := x.(type) {
case int:
    fmt.Println("int 타입:", v)
case string:
    fmt.Println("string 타입:", v)
default:
    fmt.Println("기타 타입")
}
```

---

### 5) `for` 문 (Go의 유일한 반복문)

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

---

### 6) while처럼 사용

```go
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
```

---

### 7) 무한 루프

```go
for {
    fmt.Println("무한 반복")
}
```

---

## 3️⃣ 제어 키워드

| 키워드    | 설명                       |
|-----------|----------------------------|
| `break`   | 루프 또는 switch 종료     |
| `continue`| 다음 반복으로 건너뜀       |
| `goto`    | 지정한 라벨로 이동 (사용 비추천) |
| `return`  | 함수 종료 및 값 반환       |

```go
for i := 1; i <= 5; i++ {
    if i == 3 {
        continue // 3 건너뜀
    }
    fmt.Println(i)
}
```

---

## 🎯 정리

✔ Go는 C 계열과 유사한 연산자와 문법을 가지며, `if`, `switch`, `for`가 핵심 제어문입니다.  
✔ `++`, `--` 연산자는 **식(expression)이 아닌 문장(statement)** 으로만 사용 가능합니다.  
✔ `switch`는 `break` 없이 자동 종료되며, 타입 비교에도 활용할 수 있습니다.  
✔ `for`만으로 모든 반복문 패턴을 처리할 수 있어 문법이 단순하고 명확합니다.

