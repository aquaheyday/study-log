package main

import (
	"fmt"
	"time"
)

// 단순한 Goroutine 실행 함수
func printMessage(message string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(message, ":", i)
		time.Sleep(time.Millisecond * 500) // 0.5초 대기
	}
}

func main() {
	// 일반 함수 호출 (순차 실행)
	printMessage("일반 실행")

	// Goroutine을 사용한 비동기 실행
	go printMessage("고루틴 실행 1")
	go printMessage("고루틴 실행 2")

	// 익명 함수 Goroutine 실행
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Println("익명 함수 Goroutine 실행:", i)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	// 메인 함수가 종료되지 않도록 대기
	time.Sleep(time.Second * 3)
	fmt.Println("메인 함수 종료")
}
