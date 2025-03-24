# 🧵 Go 언어 고루틴(Goroutine)과 채널(Channel)

Go는 **경량 스레드**인 고루틴(goroutine)과 **동기화 및 통신 수단**인 채널(channel)을 통해 간단하고 효율적인 **동시성 프로그래밍**을 제공합니다.

---

## 1️⃣ 고루틴(Goroutine) 이란?

- Go의 **경량 스레드(lightweight thread)**  
- 함수 앞에 `go` 키워드를 붙이면 고루틴으로 실행됨
- 수천~수만 개의 고루틴 생성 가능 (OS 스레드보다 가볍고 빠름)

```go
package main

import (
    "fmt"
    "time"
)

func sayHello() {
    fmt.Println("Hello from goroutine")
}

func main() {
    go sayHello() // 고루틴으로 실행
    time.Sleep(time.Second) // 메인 고루틴 종료 대기
}
```

✔ `go` 키워드로 새 고루틴 시작  
✔ `main` 함수가 먼저 종료되면 고루틴도 같이 종료됨 (그래서 `Sleep` 필요)  

---

## 2️⃣ 고루틴 특징

- **비동기 실행**: 고루틴은 독립적으로 실행되며, 메인 함수가 끝나면 모든 고루틴 종료
- **스택 메모리 사용량 적음** (초기 2KB 수준, 필요 시 자동 증가)
- **스케줄러**: Go 런타임이 고루틴을 자동으로 분산 처리

---

## 3️⃣ 채널(Channel)이란?

- 고루틴 간 **데이터 통신을 위한 파이프**
- 채널은 **타입이 지정된 통로**
- `chan T`는 T 타입의 값을 주고받는 채널

```go
package main

import "fmt"

func main() {
    ch := make(chan int) // int형 채널 생성

    go func() {
        ch <- 42 // 채널에 값 전송
    }()

    val := <-ch // 채널에서 값 수신
    fmt.Println("받은 값:", val)
}
```

✔ `<-` 연산자를 사용해 **채널 송수신**  
✔ 채널은 기본적으로 **동기적(synchronous)** 이라, 송신과 수신이 **둘 다 준비되어야 통신**이 이뤄짐  

---

## 4️⃣ 채널 버퍼(Buffer) 사용

- 기본 채널은 **비버퍼(동기식)**  
- 버퍼를 주면 **비동기 전송 가능**

```go
func main() {
    ch := make(chan string, 2) // 버퍼 크기 2

    ch <- "hello"
    ch <- "world"

    fmt.Println(<-ch)
    fmt.Println(<-ch)
}
```

✔ 버퍼 크기 만큼은 고루틴 없이도 송신 가능  
✔ 버퍼가 꽉 차면 송신 시 대기, 비어있으면 수신 시 대기  

---

## 5️⃣ 채널 닫기와 range 수신

- `close(channel)`로 채널을 닫을 수 있음  
- 닫힌 채널은 더 이상 값을 보낼 수 없음
- `range`를 사용하면 채널에서 값이 올 때까지 반복

```go
func main() {
    ch := make(chan int)

    go func() {
        for i := 1; i <= 3; i++ {
            ch <- i
        }
        close(ch)
    }()

    for val := range ch {
        fmt.Println(val)
    }
}
```

✔ `range`는 채널이 닫힐 때까지 수신 반복  
✔ 닫힌 채널에서 수신하면 **제로 값 반환**, 송신하면 **panic** 발생  

---

## 6️⃣ select 문

- 여러 채널을 **동시에 처리**할 수 있는 switch-like 문법
- 먼저 준비된 채널이 실행됨

```go
func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        ch1 <- "one"
    }()

    go func() {
        ch2 <- "two"
    }()

    select {
    case msg1 := <-ch1:
        fmt.Println("ch1:", msg1)
    case msg2 := <-ch2:
        fmt.Println("ch2:", msg2)
    }
}
```

✔ 가장 먼저 준비된 채널이 실행됨  
✔ 모든 채널이 블로킹이면 select도 대기함  
✔ `default:` 문을 넣으면 비동기 폴링 가능  

---

## 🎯 정리

| 개념 | 설명 |
|------|------|
| 고루틴 | `go` 키워드로 생성하는 경량 스레드 |
| 채널 | 고루틴 간 통신 수단, `<-` 연산자로 사용 |
| 버퍼 채널 | 비동기 전송 가능, 버퍼 초과 시 대기 |
| close | 채널 닫기, range 수신에 사용 |
| select | 다중 채널 동시 처리 |

✔ 고루틴은 간단하게 비동기 실행을 가능하게 해줌  
✔ 채널은 공유 메모리 대신 **메시지를 통해 통신**하는 구조  
✔ Go는 **동시성(concurrency)** 을 직관적으로 다룰 수 있게 설계되어 있음  
