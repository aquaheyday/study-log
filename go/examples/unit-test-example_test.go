// unit-test-example.go 파일의 내용이라고 가정한다.
// 파일명과 동일한 *_test.go 파일에서 유닛 테스트 함수 실행이 가능하다.

// unit-test-example.go 파일의 내용
package main

import (
	"fmt"
	"testing"
)

// 두 수를 더하는 함수
func Add(a, b int) int {
	return a + b
}

// 두 수를 곱하는 함수
func Multiply(a, b int) int {
	return a * b
}

// 메인 함수 (테스트 실행과는 관계없음, 단순 예제 실행용)
func main() {
	fmt.Println("Add(2, 3):", Add(2, 3))
	fmt.Println("Multiply(2, 3):", Multiply(2, 3))
}

// unit-test-example_test.go 파일의 내용

// Add 함수 테스트
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; 기대값 %d", result, expected)
	}
}

// Multiply 함수 테스트
func TestMultiply(t *testing.T) {
	result := Multiply(2, 3)
	expected := 6
	if result != expected {
		t.Errorf("Multiply(2, 3) = %d; 기대값 %d", result, expected)
	}
}

// 테이블 기반 테스트 (여러 입력값을 테스트)
func TestTableDriven(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{1, 2, 3},
		{2, 3, 5},
		{10, 20, 30},
		{5, -5, 0},
	}

	for _, tt := range tests {
		result := Add(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Add(%d, %d) = %d; 기대값 %d", tt.a, tt.b, result, tt.expected)
		}
	}
}


// 테스트 실행 방법 터미널에 go test 입력
// 상세 로그를 포함하고 싶으면 go test -v 옵션 사용 가능
