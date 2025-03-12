# Go 언어: 맵(Map)과 구조체(Struct)

## 1. 맵(Map)
Go의 맵(Map)은 키-값(Key-Value) 형태의 데이터를 저장하는 자료구조입니다.

### 1.1 맵 선언 및 초기화
```go
// 방법 1: make 함수 사용
var myMap = make(map[string]int)

// 방법 2: 맵 리터럴 사용
myMap := map[string]int{
    "apple":  5,
    "banana": 3,
}
```

### 1.2 맵 값 추가 및 삭제
```go
// 값 추가 및 업데이트
myMap["orange"] = 7  // 새로운 키 추가
myMap["apple"] = 10   // 기존 키 값 업데이트

// 값 삭제
delete(myMap, "banana")
```

### 1.3 맵 값 조회
```go
value, exists := myMap["apple"]
if exists {
    fmt.Println("Value:", value)
} else {
    fmt.Println("Key does not exist")
}
```

### 1.4 맵 순회
```go
for key, value := range myMap {
    fmt.Println("Key:", key, "Value:", value)
}
```

### 1.5 맵의 길이 확인
```go
fmt.Println("Map length:", len(myMap))
```

---

## 2. 구조체(Struct)
Go의 구조체는 여러 개의 필드를 그룹화하여 사용자 정의 타입을 만들 때 사용됩니다.

### 2.1 구조체 선언 및 초기화
```go
type Person struct {
    Name string
    Age  int
    City string
}

// 초기화 방법 1: 필드 순서대로 입력
p1 := Person{"Alice", 25, "Seoul"}

// 초기화 방법 2: 필드명을 명시적으로 입력
p2 := Person{
    Name: "Bob",
    Age:  30,
    City: "Busan",
}
```

### 2.2 구조체 필드 접근 및 수정
```go
fmt.Println(p1.Name) // Alice
p1.Age = 26          // 값 변경
```

### 2.3 구조체 포인터 사용
```go
p3 := &Person{Name: "Charlie", Age: 28, City: "Incheon"}
fmt.Println(p3.Name) // Charlie
p3.Age = 29          // 포인터를 통해 값 변경 가능
```

### 2.4 구조체 메서드
```go
type Rectangle struct {
    Width  float64
    Height float64
}

// 구조체 메서드 선언
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

rect := Rectangle{Width: 5, Height: 10}
fmt.Println("Area:", rect.Area()) // 50
```
