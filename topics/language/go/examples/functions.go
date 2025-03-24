package main

import "fmt"

// 기본 함수 선언
func sayHello() {
	fmt.Println("안녕하세요, Go 함수입니다!")
}

// 매개변수가 있는 함수
func greet(name string) {
	fmt.Println("안녕하세요,", name)
}

// 여러 개의 매개변수를 받는 함수
func add(a int, b int) int {
	return a + b
}

// 여러 개의 값을 반환하는 함수
func swap(x, y string) (string, string) {
	return y, x
}

// 리턴값에 이름을 지정하는 함수 (Named Return)
func divide(a, b float64) (result float64, err error) {
	if b == 0 {
		err = fmt.Errorf("0으로 나눌 수 없습니다.")
		return
	}
	result = a / b
	return
}

// 가변 인자를 받는 함수 (Variadic Function)
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 익명 함수 (Anonymous Function)
var multiply = func(a, b int) int {
	return a * b
}

// 함수 리턴 (고차 함수 - Higher Order Function)
func getMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func main() {
	// 기본 함수 호출
	sayHello()

	// 매개변수 있는 함수 호출
	greet("Alice")

	// 리턴값이 있는 함수 호출
	result := add(10, 5)
	fmt.Println("덧셈 결과:", result)

	// 여러 개의 리턴값 받기
	first, second := swap("Hello", "World")
	fmt.Println("문자열 교환 결과:", first, second)

	// 네임드 리턴값 사용
	divResult, err := divide(10, 2)
	if err != nil {
		fmt.Println("오류:", err)
	} else {
		fmt.Println("나눗셈 결과:", divResult)
	}

	// 가변 인자 함수 호출
	totalSum := sum(1, 2, 3, 4, 5)
	fmt.Println("가변 인자 합계:", totalSum)

	// 익명 함수 호출
	mulResult := multiply(3, 4)
	fmt.Println("곱셈 결과:", mulResult)

	// 함수 리턴 활용
	double := getMultiplier(2)
	triple := getMultiplier(3)
	fmt.Println("2배 함수 결과:", double(5))
	fmt.Println("3배 함수 결과:", triple(5))
}
