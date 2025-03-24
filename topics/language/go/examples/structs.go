package main

import "fmt"

// 구조체 정의
type Person struct {
	name string
	age  int
	city string
}

// 구조체를 반환하는 함수
func newPerson(name string, age int, city string) Person {
	return Person{name, age, city}
}

// 구조체 포인터를 반환하는 함수
func newPersonPointer(name string, age int, city string) *Person {
	return &Person{name, age, city}
}

// 구조체 메서드 정의
func (p Person) introduce() {
	fmt.Printf("안녕하세요, 제 이름은 %s이고, 나이는 %d살이며 %s에 삽니다.\n", p.name, p.age, p.city)
}

// 구조체 메서드 (포인터 리시버로 값 변경)
func (p *Person) birthday() {
	p.age++
}

func main() {
	// 구조체 선언 및 초기화
	var person1 Person
	person1.name = "Alice"
	person1.age = 25
	person1.city = "Seoul"

	fmt.Println("기본 구조체:", person1)

	// 구조체 리터럴 초기화
	person2 := Person{name: "Bob", age: 30, city: "Busan"}
	fmt.Println("리터럴 초기화:", person2)

	// 필드 순서로 초기화
	person3 := Person{"Charlie", 28, "Incheon"}
	fmt.Println("순서 초기화:", person3)

	// 구조체 반환 함수 사용
	person4 := newPerson("David", 35, "Daegu")
	fmt.Println("함수를 통한 구조체 생성:", person4)

	// 구조체 포인터 사용
	personPtr := newPersonPointer("Eve", 22, "Gwangju")
	fmt.Println("포인터를 통한 구조체 생성:", *personPtr)

	// 구조체 메서드 호출
	person1.introduce()
	person2.introduce()

	// 구조체 값 변경 (포인터 리시버 사용)
	person2.birthday()
	fmt.Println("생일이 지나 나이 증가:", person2)
}
