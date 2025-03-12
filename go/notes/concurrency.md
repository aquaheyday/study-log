# Go 언어: 동시성(Concurrency) 기초

Go 언어는 **고루틴(Goroutine)**과 **채널(Channel)**을 활용하여 동시성 프로그래밍을 간결하게 구현할 수 있습니다.

---

## 1. 고루틴(Goroutine)
고루틴은 Go의 경량 스레드로, `go` 키워드를 사용하여 함수를 비동기 실행할 수 있습니다.

### 1.1 고루틴 생성
```go
package main

import (
    "fmt"
    "time"
)

func printMessage(msg string) {
    for i := 0; i < 5; i++ {
        fmt.Println(msg, i)
        time.Sleep(time.Millisecond * 500)
    }
}

func main() {
    go printMessage("Goroutine") // 고루틴 실행
    printMessage("Main Function") // 메인 함수 실행
}
```

### 주요 패턴
- `go 함수이름()`: 새로운 고루틴 실행
- 메인 함수(`main()`)가 종료되면 모든 고루틴이 종료됨
- `time.Sleep()`을 사용하여 동작을 지연할 수 있음

---

## 2. 채널(Channel)
채널을 사용하면 고루틴 간 안전한 데이터 공유 및 통신이 가능합니다.

### 2.1 기본 채널 사용
```go
package main

import "fmt"

func sendMessage(ch chan string) {
    ch <- "Hello, Channel!"
}

func main() {
    ch := make(chan string) // 채널 생성
    go sendMessage(ch) // 고루틴 실행
    msg := <-ch // 채널에서 데이터 수신
    fmt.Println(msg)
}
```

### 주요 패턴
- `ch := make(chan 타입)`: 채널 생성
- `ch <- value`: 채널에 데이터 전송
- `value := <-ch`: 채널에서 데이터 수신

---

## 3. 버퍼링된 채널(Buffered Channel)
버퍼 크기를 지정하여 데이터 저장이 가능한 채널입니다.

```go
package main

import "fmt"

func main() {
    ch := make(chan string, 2) // 버퍼 크기 2
    ch <- "Hello"
    ch <- "World"
    fmt.Println(<-ch) // Hello
    fmt.Println(<-ch) // World
}
```

### 주요 패턴
- `make(chan 타입, 버퍼크기)`: 버퍼 크기가 있는 채널 생성
- 버퍼가 가득 차면 **송신자(Sender)는 블록됨**
- 버퍼가 비어 있으면 **수신자(Receiver)는 블록됨**

---

## 4. 채널의 `close()`와 `range`
채널을 `close()` 하면 더 이상 데이터를 보낼 수 없으며, `range`를 사용하여 반복 처리할 수 있습니다.

```go
package main

import "fmt"

func main() {
    ch := make(chan int, 3)
    ch <- 1
    ch <- 2
    ch <- 3
    close(ch) // 채널 닫기
    
    for v := range ch {
        fmt.Println(v)
    }
}
```

### 주요 패턴
- `close(ch)`: 채널을 닫음
- `for v := range ch`: 채널에서 데이터를 반복적으로 읽음

---

## 5. `select`를 활용한 다중 채널 처리
`select`문을 사용하면 여러 채널에서 데이터를 기다릴 수 있습니다.

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    go func() {
        time.Sleep(2 * time.Second)
        ch1 <- "Message from Channel 1"
    }()
    
    go func() {
        time.Sleep(1 * time.Second)
        ch2 <- "Message from Channel 2"
    }()
    
    select {
    case msg := <-ch1:
        fmt.Println(msg)
    case msg := <-ch2:
        fmt.Println(msg)
    }
}
```

### 주요 패턴
- `select`를 사용하여 여러 채널을 동시에 대기
- 먼저 도착한 데이터를 처리

---

## 6. `sync.WaitGroup`을 활용한 동기화
`sync.WaitGroup`을 사용하면 여러 개의 고루틴이 종료될 때까지 기다릴 수 있습니다.

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // 작업 완료 시 WaitGroup 감소
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup
    
    for i := 1; i <= 3; i++ {
        wg.Add(1) // WaitGroup 증가
        go worker(i, &wg)
    }
    
    wg.Wait() // 모든 고루틴이 종료될 때까지 대기
    fmt.Println("All workers completed")
}
```

### 주요 패턴
- `wg.Add(n)`: `n`개의 작업을 대기
- `wg.Done()`: 하나의 작업이 완료됨을 알림
- `wg.Wait()`: 모든 작업이 끝날 때까지 대기

---

## 7. 뮤텍스(Mutex)를 활용한 공유 자원 보호
여러 개의 고루틴이 공유 데이터를 안전하게 사용하도록 `sync.Mutex`를 활용할 수 있습니다.

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type SafeCounter struct {
    mu sync.Mutex
    value int
}

func (c *SafeCounter) Increment() {
    c.mu.Lock()
    c.value++
    c.mu.Unlock()
}

func main() {
    var counter SafeCounter
    var wg sync.WaitGroup
    
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter.Increment()
        }()
    }
    
    wg.Wait()
    fmt.Println("Final Counter Value:", counter.value)
}
```

### 주요 패턴
- `sync.Mutex`를 사용하여 공유 자원을 보호
- `mu.Lock()`으로 임계 영역 보호, `mu.Unlock()`으로 해제

---

## 8. Go 동시성 베스트 프랙티스
**고루틴을 남발하지 말 것!** 필요할 때만 사용
**채널을 사용하여 데이터를 안전하게 교환**
**sync.WaitGroup과 Mutex를 적절히 활용**
**select 문을 사용하여 다중 채널을 효과적으로 처리**
**고루틴이 메모리 누수를 유발하지 않도록 주의**
