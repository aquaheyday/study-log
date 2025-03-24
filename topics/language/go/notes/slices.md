# Go (Golang) 배열과 슬라이스 정리

Go에서 **배열(Array)** 과 **슬라이스(Slice)** 는 데이터를 저장하는 자료구조입니다.  
배열은 **고정된 크기**, 슬라이스는 **동적 크기 조정 가능**한 차이가 있습니다.

---

## 1. 배열 (Array)

### 1.1 배열 선언 및 초기화
배열은 **고정된 크기**를 가진 동일한 타입의 데이터 모음입니다.

```go
package main

import "fmt"

func main() {
    var arr1 [3]int              // 크기가 3인 정수형 배열 선언
    arr2 := [3]int{1, 2, 3}       // 크기가 3인 배열 선언과 초기화
    arr3 := [...]int{10, 20, 30}  // 크기를 자동으로 설정

    fmt.Println(arr1) // [0 0 0] (초기값: 0)
    fmt.Println(arr2) // [1 2 3]
    fmt.Println(arr3) // [10 20 30]
}
```
✅ 배열의 크기는 고정되어 있으며, 한 번 선언되면 변경할 수 없습니다.  
✅ `[...]int{}`를 사용하면 크기를 자동으로 설정할 수 있습니다.

---

### 1.2 배열 요소 접근 및 수정
배열의 요소는 **인덱스(index)를 사용하여 접근 및 수정**할 수 있습니다.

```go
package main

import "fmt"

func main() {
    arr := [3]int{100, 200, 300}

    fmt.Println(arr[0]) // 첫 번째 요소 출력 (100)

    arr[1] = 500        // 두 번째 요소 변경
    fmt.Println(arr)    // [100 500 300]
}
```
✅ 배열의 인덱스는 **0부터 시작**하며, `arr[n]`을 사용하여 접근합니다.  
✅ 요소 값을 변경하려면 `arr[index] = 값`을 사용합니다.

---

### 1.3 배열 길이 확인 (`len()`)
배열의 길이는 `len()` 함수를 사용하여 확인할 수 있습니다.

```go
package main

import "fmt"

func main() {
    arr := [5]int{1, 2, 3, 4, 5}
    fmt.Println("배열의 길이:", len(arr)) // 출력: 5
}
```
✅ 배열의 크기는 선언할 때 지정되며, `len(arr)`을 사용하여 크기를 확인할 수 있습니다.

---

### 1.4 배열 반복문 (`for`)
배열은 `for` 루프를 사용하여 반복할 수 있습니다.

```go
package main

import "fmt"

func main() {
    arr := [3]string{"Apple", "Banana", "Cherry"}

    // 일반적인 for 문
    for i := 0; i < len(arr); i++ {
        fmt.Println(arr[i])
    }

    // range 를 사용한 반복
    for index, value := range arr {
        fmt.Printf("Index: %d, Value: %s\n", index, value)
    }
}
```
✅ `for` 문을 사용하여 배열의 모든 요소를 순회할 수 있습니다.  
✅ `range` 키워드를 사용하면 인덱스와 값을 동시에 가져올 수 있습니다.

---

## 2. 슬라이스 (Slice)

### 2.1 슬라이스란?
슬라이스는 **배열보다 더 유연한 크기 조정이 가능한 자료구조**입니다.  
배열과 다르게 크기를 동적으로 변경할 수 있습니다.

```go
package main

import "fmt"

func main() {
    var slice1 []int                // 빈 슬라이스 선언
    slice2 := []int{10, 20, 30, 40} // 슬라이스 선언 및 초기화

    fmt.Println(slice1) // []
    fmt.Println(slice2) // [10 20 30 40]
}
```
✅ 슬라이스는 배열과 다르게 크기를 동적으로 변경할 수 있습니다.  
✅ `[]int{}`처럼 선언하면 슬라이스가 자동으로 생성됩니다.

---

### 2.2 슬라이스의 길이와 용량 (`len()` & `cap()`)
슬라이스의 크기는 `len()`을 사용하여 확인할 수 있으며, **용량(capacity)은 기본 배열의 크기를 나타냅니다.**

```go
package main

import "fmt"

func main() {
    slice := []int{1, 2, 3, 4, 5}

    fmt.Println("길이:", len(slice)) // 5
    fmt.Println("용량:", cap(slice)) // 5 (초기 용량은 길이와 동일)
}
```
✅ `len(slice)` → 현재 슬라이스의 요소 개수  
✅ `cap(slice)` → 슬라이스의 **최대 저장 가능 크기**  

---

### 2.3 `make()`를 사용한 슬라이스 생성
`make()` 함수를 사용하여 슬라이스를 생성할 수 있습니다.

```go
package main

import "fmt"

func main() {
    slice := make([]int, 3, 5) // 길이 3, 용량 5인 슬라이스 생성
    fmt.Println(slice)         // [0 0 0]
    fmt.Println("길이:", len(slice), "용량:", cap(slice)) // 길이: 3, 용량: 5
}
```
✅ `make([]타입, 길이, 용량)` 형식으로 슬라이스를 생성합니다.  
✅ `make()`로 생성한 슬라이스는 기본값으로 0이 들어갑니다.

---

### 2.4 슬라이스 요소 추가 (`append()`)
`append()` 함수를 사용하면 슬라이스에 새로운 요소를 추가할 수 있습니다.

```go
package main

import "fmt"

func main() {
    slice := []int{1, 2, 3}
    slice = append(slice, 4, 5) // 요소 추가

    fmt.Println(slice) // [1 2 3 4 5]
}
```
✅ `append()`는 새로운 요소를 추가하며, 필요 시 자동으로 용량을 증가시킵니다.  

---

### 2.5 슬라이싱 (Slicing)
슬라이싱을 사용하면 기존 슬라이스에서 **부분적으로 잘라내어 새로운 슬라이스를 만들 수 있습니다.**

```go
package main

import "fmt"

func main() {
    nums := []int{10, 20, 30, 40, 50}

    slice1 := nums[1:4] // 1번 인덱스부터 4번 인덱스 전까지 (20, 30, 40)
    slice2 := nums[:3]  // 처음부터 3번 인덱스 전까지 (10, 20, 30)
    slice3 := nums[2:]  // 2번 인덱스부터 끝까지 (30, 40, 50)

    fmt.Println(slice1) // [20 30 40]
    fmt.Println(slice2) // [10 20 30]
    fmt.Println(slice3) // [30 40 50]
}
```
✅ **`slice[start:end]`** 형식으로 **start부터 end-1까지의 요소를 포함**합니다.  
✅ **`[:end]`** → 처음부터 `end-1`까지 포함  
✅ **`[start:]`** → `start`부터 끝까지 포함  

---

## 3. 배열과 슬라이스의 차이점

| 구분 | 배열 (Array) | 슬라이스 (Slice) |
|------|------------|----------------|
| **크기** | 고정됨 | 동적 크기 조정 가능 |
| **메모리** | 선언된 크기만큼 메모리 할당 | 필요에 따라 자동 확장 |
| **초기화 방법** | `var arr [5]int` | `var slice []int` 또는 `make([]int, 5)` |
| **메모리 효율성** | 메모리 낭비 발생 가능 | 유연하게 크기 조절 가능 |

---

## 4. 결론
- **배열(Array)**: **고정 크기**를 가지며 선언 시 크기를 변경할 수 없습니다.
- **슬라이스(Slice)**: **동적으로 크기 조정 가능**하며 배열보다 유연합니다.
- **슬라이스의 `append()`, `make()`, `slicing`을 활용하면 보다 유연한 데이터 관리가 가능합니다.**
