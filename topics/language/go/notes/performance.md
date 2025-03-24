# ⚡ Go 언어 성능 최적화

Go는 본래 빠르고 효율적인 언어지만, **코드 작성 방식이나 구조에 따라 성능 차이**가 발생할 수 있습니다.  
이 문서는 **슬라이스, 메모리, 병렬 처리, GC, 프로파일링** 등 다양한 관점에서 최적화 기법을 다룹니다.

---

## 1️⃣ 프로파일링 도구 활용

### 1) CPU 사용량 분석

```bash
go test -bench=. -cpuprofile=cpu.out
go tool pprof cpu.out
```

---

### 2) 메모리 사용량 분석

```bash
go test -bench=. -memprofile=mem.out
go tool pprof mem.out
```

✔ 병목 지점, 메모리 누수, GC 문제 등을 시각화 가능  
✔ `net/http/pprof`로 웹 기반 실시간 분석도 가능

---

## 2️⃣ 슬라이스 최적화

### 1) make로 용량(cap) 미리 지정

```go
nums := make([]int, 0, 1000)
```

✔ append() 중간에 메모리 재할당이 일어나지 않도록 용량(cap) 설정

---

### 2) 슬라이스 복사 vs 참조

```go
copy := append([]int(nil), original...) // 안전 복사
ref := original[:3] // 원본과 메모리 공유 (주의)
```

✔ 데이터 변경 의도가 없으면 복사 추천

---

## 3️⃣ 문자열 처리 최적화

### 문자열 연결은 `strings.Builder` 사용

```go
var sb strings.Builder
sb.WriteString("Hello")
sb.WriteString("World")
result := sb.String()
```

✔ 문자열 `+` 연산 반복은 메모리 낭비 큼 → `Builder`가 훨씬 빠름

---

## 4️⃣ 메모리 할당과 GC 줄이기

### 1) 포인터 남용 X

- **작은 struct**는 값 복사
- 포인터는 힙 할당 → GC 부담 증가 = 가비지 컬렉터(Garbage Collector) 가 더 자주 일하게 됨

```go
user := User{Name: "Go"}        // 값으로 사용
userPtr := &User{Name: "Lang"}  // 필요할 때만 포인터
```

---

### 2) `sync.Pool`로 객체 재사용

```go
var pool = sync.Pool{
    New: func() any {
        return new(bytes.Buffer)
    },
}

buf := pool.Get().(*bytes.Buffer)
buf.Reset()
pool.Put(buf)
```

✔ 자주 생성/삭제되는 객체 재활용  
✔ GC 부하 줄이고 성능 향상

---

## 5️⃣ map 최적화

- 초기 용량 설정: `make(map[string]int, 1000)`
- 존재 확인: `v, ok := m[key]`
- 짧고 고유한 키 사용 (해시 충돌 줄이기)

---

## 6️⃣ 고루틴과 채널 최적화

- 고루틴은 가볍지만 **과도한 생성은 오히려 성능 저하**
- 채널은 **버퍼 설정**으로 블로킹 줄이기

```go
ch := make(chan int, 100)
```

✔ 작업자 패턴(Worker Pool) 구현 시 `sync.WaitGroup`도 활용

---

## 7️⃣ 컴파일 최적화

```bash
go build -ldflags="-s -w"
```

✔ 디버그 심볼 제거 → 바이너리 크기 최소화  
✔ `upx` 등으로 압축 가능

---

## 8️⃣ Escape Analysis & Inlining 확인

```bash
go build -gcflags="-m main.go"
```

✔ `... escapes to heap` → 힙 할당 발생  
✔ 값 타입 → 스택에서 처리 → 빠르고 GC 영향 없음

---

## 9️⃣ GC 튜닝

```bash
GOGC=100   # 기본값 (100%)
GOGC=200   # GC 빈도 줄임 → CPU ↓, 메모리 ↑
GOGC=50    # GC 자주 돌림 → CPU ↑, 메모리 ↓
```

✔ `GOGC`는 GC의 민감도를 조절하는 환경 변수  
✔ GC 로그 보기: `GODEBUG=gctrace=1`

---

## 🔟 벤치마크 테스트 작성

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = 1 + 2
    }
}
```

```bash
go test -bench=. -benchmem
```

✔ 성능 측정 + 메모리 할당 확인  
✔ `testing.B`는 내부적으로 반복 횟수 조절함

## 🎯 정리

✔ 최적화는 **측정 후에** 진행  
✔ `pprof`, `benchstat`, `go test`, `trace` 등을 적극 활용  
✔ 코드 가독성 ↔ 성능 트레이드오프를 항상 고려  
