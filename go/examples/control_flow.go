package main

import "fmt"

func main() {
	// 1. if-else 조건문
	var number int = 10

	if number > 0 {
		fmt.Println("양수입니다.")
	} else if number < 0 {
		fmt.Println("음수입니다.")
	} else {
		fmt.Println("0입니다.")
	}

	// 2. if 문에서 초기 변수 선언 가능
	if x := 20; x > 10 {
		fmt.Println("x는 10보다 큽니다.")
	}

	// 3. switch 문
	day := "월요일"

	switch day {
	case "월요일":
		fmt.Println("한 주가 시작되었습니다.")
	case "금요일":
		fmt.Println("주말이 곧 다가옵니다.")
	default:
		fmt.Println("일반적인 하루입니다.")
	}

	// 4. 반복문 - for 기본형
	for i := 0; i < 5; i++ {
		fmt.Println("반복:", i)
	}

	// 5. for 문에서 조건문만 사용 (while 문처럼 사용 가능)
	count := 0
	for count < 3 {
		fmt.Println("while처럼 반복:", count)
		count++
	}

	// 6. 무한 루프 (break 사용)
	i := 0
	for {
		if i >= 3 {
			break
		}
		fmt.Println("무한 루프 실행 중:", i)
		i++
	}

	// 7. continue 사용 예제
	for j := 0; j < 5; j++ {
		if j%2 == 0 {
			continue // 짝수일 때 건너뛰기
		}
		fmt.Println("홀수 값:", j)
	}

	// 8. range를 사용한 반복문
	numbers := []int{10, 20, 30, 40}
	for index, value := range numbers {
		fmt.Printf("인덱스: %d, 값: %d\n", index, value)
	}

	// 9. map에서 range 사용
	ages := map[string]int{"Alice": 25, "Bob": 30}
	for key, value := range ages {
		fmt.Printf("이름: %s, 나이: %d\n", key, value)
	}
}
