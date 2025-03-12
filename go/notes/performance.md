# Go 언어 성능 최적화 팁

Go 언어에서 성능을 최적화하는 다양한 방법을 정리합니다.

## 1. 데이터 구조 최적화
### 1.1 배열과 슬라이스 사용 최적화
- 슬라이스의 크기를 미리 할당하여 동적 할당 비용을 줄입니다.
- `make` 함수를 활용하여 용량을 지정합니다.

```go
s := make([]int, 0, 1000) // 1000개의 용량을 미리 예약
```

### 1.2 맵(Map) 최적화
- 맵 크기를 예상할 수 있다면 `make`를 사용하여 초기 용량을 설정합니다.
- 충돌을 줄이기 위해 키의 해시 분포가 균등한지 확인합니다.

```go
m := make(map[string]int, 1000) // 1000개의 엔트리를 예상
```

## 2. 메모리 관리 최적화
### 2.1 객체 풀링 (sync.Pool)
- 반복적으로 생성되는 객체는 `sync.Pool`을 사용하여 재사용합니다.

```go
import "sync"

var pool = sync.Pool{
	New: func() interface{} {
		return new(MyStruct)
	},
}

obj := pool.Get().(*MyStruct)
defer pool.Put(obj) // 사용 후 반환
```

### 2.2 불필요한 메모리 할당 줄이기
- 문자열과 바이트 변환 시 `[]byte` 또는 `string` 변환을 최소화합니다.

```go
// 비효율적인 변환
b := []byte("hello")
s := string(b)
```

- `strings.Builder`를 사용하여 문자열 연결 성능을 개선합니다.

```go
import "strings"

var sb strings.Builder
sb.WriteString("Hello, ")
sb.WriteString("World!")
result := sb.String()
```

## 3. 병렬 처리 최적화
### 3.1 고루틴 효율적 사용
- 너무 많은 고루틴을 생성하면 컨텍스트 스위칭 비용이 증가할 수 있습니다.
- `sync.WaitGroup`과 `worker pool` 패턴을 활용합니다.

```go
import (
	"sync"
)

var wg sync.WaitGroup

for i := 0; i < 10; i++ {
	wg.Add(1)
	go func(i int) {
		defer wg.Done()
		// 작업 수행
	}(i)
}

wg.Wait()
```

### 3.2 채널(Buffer) 활용
- 채널을 사용할 때는 `buffered channel`을 사용하여 블로킹을 줄입니다.

```go
ch := make(chan int, 100) // 100개의 버퍼
```

## 4. 함수 및 코드 최적화
### 4.1 반복문 최적화
- `range`를 활용하여 슬라이스를 순회합니다.

```go
for i, v := range arr {
	// i: 인덱스, v: 값
}
```

### 4.2 인라인 함수 사용
- 작은 함수는 컴파일러가 인라인 최적화를 수행할 수 있도록 단순하게 작성합니다.

## 5. 성능 분석 및 벤치마킹
### 5.1 벤치마킹
- `testing.B`를 활용하여 성능을 측정합니다.

```go
func BenchmarkFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Function()
	}
}
```

### 5.2 프로파일링
- `pprof`를 활용하여 CPU 및 메모리 사용량을 분석합니다.

```sh
go test -bench . -benchmem -cpuprofile cpu.prof
```

```go
import _ "net/http/pprof"
```

## 6. 결론
- 데이터 구조를 최적화하고, 메모리 사용을 최소화합니다.
- 고루틴과 채널을 적절히 활용하여 병렬 처리 성능을 높입니다.
- 벤치마킹 및 프로파일링을 통해 성능 병목을 파악하고 최적화합니다.
