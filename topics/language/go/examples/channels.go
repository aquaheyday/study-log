package main

import (
	"fmt"
	"time"
)

// 기본 채널 예제 (송수신)
func basicChannel() {
	// 정수형 채널 생성
	ch := make(chan int)

	// 고루틴에서 채널에 값 전송
	go func() {
		fmt.Println("고루틴: 채널에 10 전송")
		ch <- 10 // 채널에 값 전달
	}()

	// 채널에서 값 수신
	val := <-ch
	fmt.Println("메인 함수: 채널에서 받은 값:", val)
}

// 버퍼링 채널 예제
func bufferedChannel() {
	// 용량이 2인 버퍼 채널 생성
	ch := make(chan string, 2)

	// 채널에 값 전송
	ch <- "Hello"
	ch <- "Go"

	// 채널에서 값 수신
	fmt.Println("버퍼 채널에서 받은 값:", <-ch)
	fmt.Println("버퍼 채널에서 받은 값:", <-ch)
}

// 채널 동기화 (고루틴이 완료될 때까지 대기)
func worker(done chan bool) {
	fmt.Println("작업 시작...")
	time.Sleep(time.Second) // 작업 시뮬레이션
	fmt.Println("작업 완료!")
	done <- true // 완료 신호 전송
}

// range와 close를 활용한 채널
func rangeAndCloseChannel() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i // 채널에 값 전송
			time.Sleep(time.Millisecond * 500)
		}
		close(ch) // 채널 닫기
	}()

	for val := range ch { // 채널에서 값 수신 (채널이 닫히면 종료)
		fmt.Println("채널에서 받은 값:", val)
	}
}

func main() {
	// 기본 채널 사용
	fmt.Println("\n 기본 채널 예제")
	basicChannel()

	// 버퍼링 채널 사용
	fmt.Println("\n 버퍼링 채널 예제")
	bufferedChannel()

	// 채널을 이용한 동기화
	fmt.Println("\n 채널 동기화 예제")
	done := make(chan bool)
	go worker(done)
	<-done // 채널에서 값 수신 후 메인 함수 종료
	fmt.Println("메인 함수: 모든 작업 완료")

	// range와 close를 활용한 채널
	fmt.Println("\n range와 close 활용 예제")
	rangeAndCloseChannel()
}
