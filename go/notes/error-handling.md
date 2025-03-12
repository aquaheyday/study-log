# Go 언어: 에러 처리 패턴

Go 언어는 에러 처리를 예외(Exception) 기반이 아닌 **오류 반환(Error Return) 방식**을 사용합니다. 

---

## 1. 기본적인 에러 처리
Go에서 `error` 타입을 반환하여 에러를 처리하는 기본적인 패턴입니다.

```go
package main

import (
    "errors"
    "fmt"
)

// 에러를 반환하는 함수 정의
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result:", result)
}
```

### 주요 패턴
- `errors.New("message")`: 새로운 에러 객체 생성
- 함수에서 `error` 타입을 반환하여 호출자가 처리하도록 함
- `if err != nil` 패턴을 사용하여 에러를 확인 후 처리

---

## 2. fmt.Errorf 를 활용한 에러 포맷팅
Go에서는 `fmt.Errorf`를 활용하여 상세한 에러 메시지를 만들 수 있습니다.

```go
package main

import (
    "fmt"
)

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("invalid division: %f / %f", a, b)
    }
    return a / b, nil
}
```

### 주요 패턴
- `fmt.Errorf("message: %v", value)`: 에러 메시지에 추가 정보를 포함 가능

---

## 3. errors 패키지의 `errors.Is` 와 `errors.As`
Go 1.13 이후부터는 `errors.Is`와 `errors.As`를 이용하여 에러를 비교하거나 특정 타입으로 변환할 수 있습니다.

```go
package main

import (
    "errors"
    "fmt"
)

var ErrDivideByZero = errors.New("cannot divide by zero")

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, ErrDivideByZero
    }
    return a / b, nil
}

func main() {
    _, err := divide(10, 0)
    if errors.Is(err, ErrDivideByZero) {
        fmt.Println("Handled divide by zero error!")
    }
}
```

### 주요 패턴
- `errors.Is(err, targetErr)`: 특정 에러인지 확인
- `errors.As(err, &targetType)`: 특정 타입의 에러로 변환 가능

---

## 4. 사용자 정의 에러 타입
Go에서는 `error` 인터페이스를 구현하여 사용자 정의 에러를 만들 수 있습니다.

```go
package main

import (
    "fmt"
)

// 사용자 정의 에러 타입
type MyError struct {
    Code    int
    Message string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func doSomething() error {
    return &MyError{Code: 404, Message: "Not Found"}
}

func main() {
    err := doSomething()
    if err != nil {
        fmt.Println("Custom Error:", err)
    }
}
```

### 주요 패턴
- `error` 인터페이스를 구현하는 구조체 생성
- `func (e *MyError) Error() string` 메서드 구현
- 특정한 오류 정보를 포함하는 사용자 정의 에러 반환 가능

---

## 5. 패닉(panic)과 복구(recover)
Go에서는 치명적인 오류가 발생하면 `panic`을 발생시키고, 이를 `recover`로 복구할 수 있습니다.

```go
package main

import "fmt"

func riskyFunction() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()
    
    fmt.Println("Causing panic!")
    panic("something went wrong")
}

func main() {
    riskyFunction()
    fmt.Println("Program continues...")
}
```

### 주요 패턴
- `panic("message")`: 프로그램 실행 중단 및 패닉 발생
- `recover()`: `defer`를 활용하여 패닉 복구 가능

**주의:** `panic`은 일반적인 에러 처리 대신, 예상치 못한 상황에서만 사용해야 합니다.

---

## 6. Go 에러 처리 베스트 프랙티스
**가능하면 명확한 에러 메시지 반환**
**에러 발생 시 즉시 반환 (`if err != nil { return err }` 패턴 사용)**
**`errors.Is` 및 `errors.As`를 사용하여 에러 비교**
**패닉(panic) 대신 에러를 반환하여 처리하도록 설계**
**사용자 정의 에러 타입을 활용하여 상세한 에러 정보 제공**
