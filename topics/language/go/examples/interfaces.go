package main

import "fmt"

// 인터페이스 정의
type Speaker interface {
	Speak() string
}

// 구조체와 인터페이스 구현
type Person struct {
	name string
}

func (p Person) Speak() string {
	return "안녕하세요, 제 이름은 " + p.name + "입니다."
}

type Dog struct {
	breed string
}

func (d Dog) Speak() string {
	return "멍멍! 나는 " + d.breed + "입니다."
}

// 인터페이스를 매개변수로 사용하는 함수
func introduce(s Speaker) {
	fmt.Println(s.Speak())
}

// 빈 인터페이스 (모든 타입을 받을 수 있음)
func describe(i interface{}) {
	fmt.Printf("값: %v, 타입: %T\n", i, i)
}

func main() {
	// 인터페이스 구현체 생성
	person := Person{name: "Alice"}
	dog := Dog{breed: "골든 리트리버"}

	// 인터페이스를 사용한 함수 호출
	introduce(person)
	introduce(dog)

	// 빈 인터페이스 활용
	describe("문자열")
	describe(42)
	describe(3.14)
	describe(person)
}
