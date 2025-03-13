package main

import (
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

// 벤치마크 테스트 - Add 함수
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ { // b.N 만큼 반복 실행 (Go가 자동 조정)
		Add(10, 20)
	}
}

// 벤치마크 테스트 - Multiply 함수
func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiply(10, 20)
	}
}

// 벤치마크 테스트 실행 시 `go test -bench .` 명령어를 사용
