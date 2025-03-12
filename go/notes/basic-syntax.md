# 📌 Go (Golang) 기본 문법 정리

Go는 **간결하고 빠른 컴파일 언어**로, **웹 서버, CLI 도구, 마이크로서비스** 등 다양한 분야에서 사용됩니다.

---

## 1. Go 프로그램 구조

Go 코드는 `package` 단위로 작성되며, `main` 패키지가 실행 진입점이 됩니다.

```go
package main

import "fmt" // 표준 라이브러리 패키지 가져오기

func main() {
    fmt.Println("Hello, Go!") // 출력
}
```
Go 프로그램을 실행하려면 `package main`을 포함해야 하며, `main()` 함수가 프로그램의 시작점이 됩니다.

---

## 2. 변수 선언

```go
package main

import "fmt"

func main() {
    var a int = 10      // 명시적 선언
    b := 20             // 암시적 선언 (타입 추론)
    var c, d string = "Hello", "Go" // 여러 변수 선언

    fmt.Println(a, b, c, d)
}
```
변수는 `var` 키워드를 사용하여 선언할 수 있으며, `:=`를 사용하면 타입을 추론하여 선언할 수 있습니다.

---

## 3. 데이터 타입

| 타입 | 설명 |
|------|------|
| `bool` | `true` 또는 `false` 값을 가집니다. |
| `string` | 문자열 타입입니다. |
| `int`, `int8`, `int16`, `int32`, `int64` | 정수 타입이며 크기가 다릅니다. |
| `uint`, `uint8`, `uint16`, `uint32`, `uint64` | 부호 없는 정수 타입입니다. |
| `float32`, `float64` | 실수 타입입니다. |
| `complex64`, `complex128` | 복소수 타입입니다. |
| `byte` | `uint8`과 동일합니다. |
| `rune` | `int32`와 동일하며, 유니코드 문자를 저장할 때 사용됩니다. |

Go는 정수, 실수, 문자열, 논리형 등 다양한 데이터 타입을 제공합니다.

---

## 4. 상수 (Constants)

```go
package main

import "fmt"

const Pi float64 = 3.14159

func main() {
    fmt.Println("Pi:", Pi)
}
```
상수는 `const` 키워드를 사용하여 선언하며, 변경할 수 없습니다.

---

## 5. 조건문 & 반복문

### if 문
```go
package main

import "fmt"

func main() {
    age := 20

    if age >= 18 {
        fmt.Println("성인입니다.")
    } else {
        fmt.Println("미성년자입니다.")
    }
}
```
Go의 `if` 문에서는 조건을 괄호 없이 사용하며, `{}` 중괄호를 반드시 포함해야 합니다.

### switch 문
```go
package main

import "fmt"

func main() {
    day := "Monday"

    switch day {
    case "Monday":
        fmt.Println("월요일입니다.")
    case "Friday":
        fmt.Println("금요일입니다.")
    default:
        fmt.Println("다른 요일입니다.")
    }
}
```
`switch` 문에서는 `break`를 사용하지 않아도 자동으로 종료됩니다.

### for 문 (Go의 유일한 반복문)
```go
package main

import "fmt"

func main() {
    for i := 1; i <= 5; i++ {
        fmt.Println("반복:", i)
    }
}
```
Go에는 `while` 문이 없으며, `for` 문을 사용하여 모든 반복을 처리합니다.

---

## 6. 배열 & 슬라이스

### 배열 (고정 크기)
```go
package main

import "fmt"

func main() {
    var arr [3]int = [3]int{1, 2, 3}
    fmt.Println(arr[0]) // 첫 번째 요소 출력
}
```
배열은 고정된 크기를 가지며, 선언할 때 크기를 지정해야 합니다.

### 슬라이스 (동적 크기)
```go
package main

import "fmt"

func main() {
    nums := []int{10, 20, 30}
    nums = append(nums, 40) // 요소 추가
    fmt.Println(nums)
}
```
슬라이스는 크기가 동적으로 변하며, `append()` 함수를 사용하여 요소를 추가할 수 있습니다.

---

## 7. 맵 (Map, 딕셔너리)

```go
package main

import "fmt"

func main() {
    person := map[string]int{
        "Alice": 25,
        "Bob":   30,
    }

    fmt.Println(person["Alice"]) // 25 출력

    person["Charlie"] = 35 // 새로운 요소 추가
    fmt.Println(person)
}
```
맵은 `key-value` 형태로 데이터를 저장하는 자료구조입니다.

---

## 8. 함수 (Functions)

```go
package main

import "fmt"

// 함수 정의
func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(5, 10)
    fmt.Println("결과:", result)
}
```
Go에서 함수를 정의할 때 `func` 키워드를 사용하며, 반환 타입을 명시해야 합니다.

---

## 9. 구조체 (Struct)

```go
package main

import "fmt"

// 구조체 정의
type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{Name: "Alice", Age: 25}
    fmt.Println(p.Name, p.Age)
}
```
구조체는 여러 개의 필드를 그룹화하여 사용자 정의 타입을 만들 때 사용됩니다.

---

## 10. 고루틴 (Goroutines) - 동시성 처리

```go
package main

import (
    "fmt"
    "time"
)

// 고루틴 실행 함수
func sayHello() {
    for i := 0; i < 3; i++ {
        fmt.Println("Hello, Go!")
        time.Sleep(time.Second)
    }
}

func main() {
    go sayHello() // 고루틴 실행
    time.Sleep(3 * time.Second) // 메인 함수 종료 방지
}
```
`go` 키워드를 사용하면 **비동기 함수(고루틴)**를 실행할 수 있습니다.
