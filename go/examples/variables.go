package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 1. 변수 선언 및 초기화
	var name string = "Alice"     // 명시적 선언
	age := 25                    // 암시적 선언 (타입 추론)
	var isStudent bool = true     // 논리형

	// 2. 여러 변수 선언
	var x, y int = 10, 20
	a, b := 3.5, "GoLang"

	// 3. 기본 자료형
	var integer int = 100
	var floating float64 = 3.1415
	var boolean bool = false
	var text string = "Hello, Go!"

	// 4. 기본값 (Zero Values)
	var zeroInt int       // 0
	var zeroFloat float64 // 0.0
	var zeroBool bool     // false
	var zeroString string // ""

	// 5. 타입 변환 (Type Conversion)
	var num1 int = 10
	var num2 float64 = float64(num1) // int → float64 변환
	var num3 int = int(num2)         // float64 → int 변환 (소수점 버림)

	// 6. 상수 (Constants)
	const Pi float64 = 3.141592
	const AppName = "GoApp"

	// 출력
	fmt.Println("이름:", name, "| 나이:", age, "| 학생:", isStudent)
	fmt.Println("x:", x, "| y:", y, "| a:", a, "| b:", b)
	fmt.Println("정수:", integer, "| 실수:", floating, "| 논리:", boolean, "| 문자열:", text)
	fmt.Println("기본값 (Zero Values) →", zeroInt, zeroFloat, zeroBool, `"`+zeroString+`"`)
	fmt.Println("타입 변환 → int:", num1, "| float64:", num2, "| 다시 int:", num3)
	fmt.Println("상수 → Pi:", Pi, "| 앱 이름:", AppName)

	// 변수 타입 출력 (reflect 패키지 사용)
	fmt.Println("\n 타입 확인:")
	fmt.Println("name:", reflect.TypeOf(name))
	fmt.Println("age:", reflect.TypeOf(age))
	fmt.Println("floating:", reflect.TypeOf(floating))
	fmt.Println("boolean:", reflect.TypeOf(boolean))
	fmt.Println("num2:", reflect.TypeOf(num2))
}
