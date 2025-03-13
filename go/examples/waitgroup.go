package main

import (
	"fmt"
	"sync"
	"time"
)

// 작업 함수 (고루틴에서 실행될 함수)
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // 고루틴이 종료될 때 WaitGroup 카운트 감소
	fmt.Printf("작업자 %d 시작\n", id)
	time.Sleep(time.Second) // 작업 시뮬레이션
	fmt.Printf("작업자 %d 완료\n", id)
}

func main() {
	var wg sync.WaitGroup // WaitGroup 생성

	// 3개의 고루틴 실행
	for i := 1; i <= 3; i++ {
		wg.Add(1)       // 고루틴 실행 전 카운트 증가
		go worker(i, &wg) // 고루틴 실행
	}

	wg.Wait() // 모든 고루틴이 완료될 때까지 대기
	fmt.Println("모든 작업 완료")
}
