# Go 언어 테스트 및 벤치마크

## 1. Go 테스트 기본

Go에서는 `testing` 패키지를 사용하여 유닛 테스트를 작성할 수 있습니다.

### 1.1 테스트 파일 구조
- 테스트 파일은 `_test.go` 접미사를 가져야 합니다.
- 테스트 함수명은 `Test`로 시작해야 합니다.
- `testing.T` 타입의 포인터를 매개변수로 받아야 합니다.

### 1.2 기본 테스트 예제
```go
package main

import (
	"testing"
)

func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Expected 5, but got %d", result)
	}
}
```

### 1.3 테스트 실행
```sh
go test
```

## 2. Go 벤치마크

Go에서는 `testing.B`를 사용하여 성능 테스트(벤치마크)를 수행할 수 있습니다.

### 2.1 벤치마크 함수 구조
- 벤치마크 함수명은 `Benchmark`로 시작해야 합니다.
- `testing.B` 타입의 포인터를 매개변수로 받아야 합니다.
- 반복 실행을 위해 `b.N`을 사용해야 합니다.

### 2.2 기본 벤치마크 예제
```go
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(2, 3)
	}
}
```

### 2.3 벤치마크 실행
```sh
go test -bench .
```

## 3. 추가 기능

### 3.1 테이블 기반 테스트
테스트 케이스가 여러 개일 경우 테이블 기반 테스트를 사용하면 효율적입니다.

```go
func TestAddTableDriven(t *testing.T) {
	tests := []struct {
		a, b   int
		expect int
	}{
		{1, 1, 2},
		{2, 3, 5},
		{10, 20, 30},
	}

	for _, tt := range tests {
		result := Add(tt.a, tt.b)
		if result != tt.expect {
			t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expect)
		}
	}
}
```

### 3.2 벤치마크 결과 비교
벤치마크 실행 시 `-benchmem` 옵션을 사용하면 메모리 할당 정보를 확인할 수 있습니다.

```sh
go test -bench . -benchmem
```

### 3.3 커버리지 측정
테스트 코드의 커버리지를 확인하려면 `-cover` 옵션을 사용할 수 있습니다.

```sh
go test -cover
```

## 4. 결론
- `testing` 패키지를 사용하여 간단한 테스트와 벤치마크를 작성할 수 있습니다.
- 테이블 기반 테스트를 활용하면 가독성과 유지보수성이 향상됩니다.
- 벤치마크 실행 시 `-benchmem` 옵션을 사용하면 성능 분석에 유용합니다.
