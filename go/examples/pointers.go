package main

import "fmt"

// 1. 포인터 기본 개념
func pointerBasic() {
	// 변수 선언 및 초기화
	var a int = 10
	var p *int // 포인터 변수 선언 (*int는 int 타입의 주소를 저장하는 변수)

	// 변수의 주소를 포인터에 저장
	p = &a

	// 출력
	fmt.Println("변수 a의 값:", a)    // 10
	fmt.Println("변수 a의 주소:", &a) // 메모리 주소 출력
	fmt.Println("포인터 p의 값:", p)  // a의 주소
	fmt.Println("포인터 p가 가리키는 값:", *p) // 10 (*p는 포인터가 가리키는 값을 의미)
}

// 2. 포인터를 매개변수로 사용하는 함수
func modifyValue(x *int) {
	// 포인터를 통해 원본 값을 변경
	*x = *x * 2
}

// 3. 포인터를 반환하는 함수
func getPointer() *int {
	num := 100  // 지역 변수 선언
	return &num // 변수의 주소 반환
}

// 4. 구조체 포인터 사용 예제
type Person struct {
	name string
	age  int
}

func structPointer() {
	// 구조체 선언
	p := Person{"Alice", 25}
	ptr := &p // 구조체 포인터 선언

	// 포인터를 사용하여 구조체 필드 접근
	fmt.Println("이름:", ptr.name) // 포인터를 통해 구조체 필드 접근
	fmt.Println("나이:", ptr.age)

	// 포인터를 사용하여 구조체 값 변경
	ptr.age = 30
	fmt.Println("변경된 나이:", ptr.age)
}

// 5. 배열과 슬라이스에서 포인터 사용
func slicePointer() {
	numbers := []int{1, 2, 3, 4, 5}
	p := &numbers[0] // 배열 또는 슬라이스의 첫 번째 요소 주소 저장

	fmt.Println("슬라이스 첫 번째 값:", *p) // 1
}

func main() {
	// 1. 포인터 기본 개념
	fmt.Println("\n 포인터 기본 개념")
	pointerBasic()

	// 2. 포인터를 활용한 함수 호출
	fmt.Println("\n 포인터를 활용한 함수 호출")
	num := 10
	fmt.Println("변경 전 값:", num)
	modifyValue(&num) // 포인터를 통해 값 변경
	fmt.Println("변경 후 값:", num)

	// 3. 포인터를 반환하는 함수 호출
	fmt.Println("\n 포인터를 반환하는 함수 호출")
	ptr := getPointer()
	fmt.Println("포인터가 가리키는 값:", *ptr)

	// 4. 구조체 포인터 사용
	fmt.Println("\n 구조체 포인터 사용")
	structPointer()

	// 5. 배열과 슬라이스에서 포인터 사용
	fmt.Println("\n 배열과 슬라이스에서 포인터 사용")
	slicePointer()
}
