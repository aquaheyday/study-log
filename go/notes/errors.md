# ⚠️ Go 언어 에러 처리

Go는 **예외(exception)** 대신 **명시적인 에러 반환**을 사용해 에러 처리 흐름을 더 명확하게 만듭니다.

---

## 1️⃣ 에러 기본 처리 방식

- Go 함수는 보통 `(결과, 에러)` 형태로 리턴함
- `error`는 Go 내장 인터페이스

```go
package main

import (
    "errors"
    "fmt"
)

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("0으로 나눌 수 없습니다")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("에러 발생:", err)
        return
    }
    fmt.Println("결과:", result)
}
```

✔ `errors.New()`는 간단한 에러 메시지를 생성  
✔ `nil`은 에러 없음의 의미  

---

## 2️⃣ error 인터페이스

```go
type error interface {
    Error() string
}
```

✔ 모든 에러 타입은 `Error()` 메서드를 구현  
✔ 커스텀 에러 타입을 만들 수 있음

---

## 3️⃣ fmt.Errorf로 포맷 포함 에러 만들기

- `fmt.Errorf`를 사용해 문자열 포맷 가능
- `%w`를 사용하면 **wrap** 가능 (Go 1.13+)

```go
import "fmt"

err := fmt.Errorf("파일 읽기 실패: %w", ioErr)
```

✔ `%w`는 에러 체이닝(wrapping)을 지원해 원인 추적 가능  

---

## 4️⃣ errors.Is & errors.As (Go 1.13+)

에러 체이닝이 있을 때 원본 에러를 비교하거나 타입으로 판별 가능

```go
import (
    "errors"
    "fmt"
    "os"
)

func main() {
    err := fmt.Errorf("상위 에러: %w", os.ErrNotExist)

    if errors.Is(err, os.ErrNotExist) {
        fmt.Println("파일 없음 에러입니다")
    }
}
```

✔ `errors.Is(err, target)` → 포함 여부 확인  
✔ `errors.As(err, &target)` → 타입 캐스팅  

---

## 5️⃣ 커스텀 에러 타입 정의

```go
type MyError struct {
    Code int
    Msg  string
}

func (e MyError) Error() string {
    return fmt.Sprintf("에러 %d: %s", e.Code, e.Msg)
}

func doSomething() error {
    return MyError{Code: 404, Msg: "Not Found"}
}

func main() {
    err := doSomething()
    fmt.Println(err)
}
```

✔ `Error()` 메서드 구현 시 `error` 인터페이스 만족  
✔ `errors.As()`로 타입 추출 가능  

---

## 6️⃣ 패닉과 복구 (panic & recover)

- `panic`은 치명적인 에러 상황에 사용  
- `recover`를 통해 패닉 복구 가능 (보통 `defer`와 함께 사용)

```go
func safeDivide(a, b int) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("복구됨:", r)
        }
    }()

    if b == 0 {
        panic("0으로 나눌 수 없습니다")
    }

    fmt.Println("결과:", a/b)
}
```

✔ `panic`은 프로그램을 즉시 종료  
✔ `recover`는 패닉을 잡아내고 프로그램 계속 실행 가능  
✔ 일반적인 에러 처리에는 `panic` 사용 ❌  

---

## 🎯 정리

| 개념 | 설명 |
|------|------|
| error | Go의 기본 에러 타입 (인터페이스) |
| errors.New | 간단한 에러 생성 |
| fmt.Errorf | 포맷 포함 에러 생성 (`%w`로 wrapping 가능) |
| errors.Is / As | 에러 비교 및 타입 추출 |
| panic / recover | 치명적인 에러 처리 방식 (비권장) |

✔ Go는 **명시적 에러 반환**을 지향함 (`result, err := ...`)  
✔ `nil` 체크는 Go 에러 처리의 핵심 패턴  
✔ **복잡한 에러 로직**은 `fmt.Errorf`, `errors.Is`, `custom error`로 관리  
✔ `panic`은 복구 불가능하거나 진짜 치명적일 때만 사용
