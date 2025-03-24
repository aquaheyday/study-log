package main

import "fmt"

func main() {
	// 슬라이스 선언 및 초기화
	var nums []int
	chars := []string{"A", "B", "C"} // 리터럴을 사용한 초기화

	fmt.Println("빈 슬라이스:", nums)
	fmt.Println("문자 슬라이스:", chars)

	// make() 함수로 슬라이스 생성
	numbers := make([]int, 5)       // 길이 5, 용량 5인 슬라이스
	numbersWithCap := make([]int, 3, 10) // 길이 3, 용량 10

	fmt.Println("make()로 생성한 슬라이스:", numbers)
	fmt.Println("용량 지정한 슬라이스 (길이 3, 용량 10):", numbersWithCap)

	// append() 함수로 요소 추가
	numbers = append(numbers, 10, 20, 30) // 요소 추가
	fmt.Println("append() 후:", numbers)

	// 슬라이스의 부분 추출 (Slicing)
	slice := numbers[1:4] // 인덱스 1부터 3까지 (4는 포함되지 않음)
	fmt.Println("부분 슬라이스 (1:4):", slice)

	// 슬라이스의 길이와 용량 확인
	fmt.Println("슬라이스 길이:", len(numbers))
	fmt.Println("슬라이스 용량:", cap(numbers))

	// 슬라이스 복사 (copy)
	src := []int{1, 2, 3, 4, 5}
	dest := make([]int, len(src)) // src와 같은 길이의 슬라이스 생성
	copy(dest, src)               // src 내용을 dest로 복사

	fmt.Println("원본 슬라이스:", src)
	fmt.Println("복사된 슬라이스:", dest)

	// 슬라이스의 요소 삭제 (append를 활용)
	indexToRemove := 2 // 3번째 요소(인덱스 2) 삭제
	numbers = append(numbers[:indexToRemove], numbers[indexToRemove+1:]...)
	fmt.Println("삭제 후 슬라이스:", numbers)

	// 슬라이스의 반복문 사용
	for i, val := range numbers {
		fmt.Printf("인덱스 %d: 값 %d\n", i, val)
	}
}
