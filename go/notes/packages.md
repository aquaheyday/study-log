# Go (Golang) 패키지와 모듈 정리

Go에서 **패키지(Package)** 와 **모듈(Module)** 은 코드 관리 및 재사용성을 높이기 위한 핵심 개념입니다.  
패키지는 관련된 코드 파일을 그룹화하는 단위이며, 모듈은 패키지를 포함하는 더 큰 개념으로 프로젝트 단위를 의미합니다.

---

## 1. 패키지 (Package)

### 1.1 패키지란?
Go에서 패키지는 **코드를 논리적으로 그룹화하는 단위**입니다.  
각 Go 파일은 **반드시 하나의 패키지에 속해야 하며**, `package` 키워드를 사용하여 패키지를 선언합니다.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```
✅ `package main`은 실행 가능한 프로그램의 진입점이 되는 패키지입니다.  
✅ 실행 가능한 Go 프로그램을 만들려면 `main` 패키지가 반드시 필요합니다.  

---

### 1.2 패키지 생성 및 사용 방법
Go 프로젝트에서 사용자 정의 패키지를 생성하고 사용하는 방법을 살펴보겠습니다.

#### 1.2.1 패키지 생성
Go에서는 각 패키지를 **별도의 폴더로 관리**합니다.  
예를 들어, `mathutil`이라는 패키지를 만든다고 가정해 보겠습니다.

📂 프로젝트 구조:
```
project/
│── main.go
│── mathutil/
│   │── mathutil.go
```

`mathutil/mathutil.go` 파일을 생성하고, 다음과 같이 작성합니다.

```go
package mathutil

// Add 함수는 두 숫자의 합을 반환합니다.
func Add(a, b int) int {
    return a + b
}
```
✅ 패키지 이름은 폴더 이름과 일치해야 합니다.  
✅ **함수명이 대문자로 시작하면(`Add`), 다른 패키지에서 사용할 수 있습니다 (공개 함수).**  

---

### 1.3 패키지 사용
이제 `main.go` 파일에서 `mathutil` 패키지를 사용해 보겠습니다.

```go
package main

import (
    "fmt"
    "project/mathutil" // 사용자 정의 패키지 가져오기
)

func main() {
    result := mathutil.Add(5, 10)
    fmt.Println("결과:", result) // 출력: 15
}
```

✅ 사용자 정의 패키지를 사용하려면 `import` 문을 사용해야 합니다.  
✅ `mathutil.Add(5, 10)`처럼 패키지명을 앞에 붙여서 호출해야 합니다.  

---

### 1.4 `init()` 함수
Go의 `init()` 함수는 패키지가 로드될 때 자동으로 실행됩니다.

```go
package mathutil

import "fmt"

func init() {
    fmt.Println("mathutil 패키지 초기화됨")
}

func Add(a, b int) int {
    return a + b
}
```
✅ `init()` 함수는 **자동으로 실행되며**, 여러 개 존재할 수도 있습니다.  
✅ 패키지가 **import될 때마다 실행**되므로 초기화 작업에 유용합니다.  

---

## 2. 모듈 (Module)

### 2.1 모듈이란?
**모듈(Module)** 은 **Go 패키지를 관리하는 단위**로, Go에서 의존성을 효과적으로 관리할 수 있도록 도와줍니다.  
Go 1.11 이상에서는 **Go Modules**을 사용하여 프로젝트를 관리합니다.

---

### 2.2 모듈 초기화 (`go mod init`)
모듈을 사용하려면 프로젝트 루트 디렉토리에서 `go mod init` 명령어를 실행해야 합니다.

```sh
go mod init example.com/myproject
```
이렇게 하면 `go.mod` 파일이 생성됩니다.

---

### 2.3 `go.mod` 파일 구조
`go.mod` 파일은 프로젝트의 메타데이터를 포함합니다.

```
module example.com/myproject

go 1.20
```
✅ `module` 키워드 뒤에 모듈 이름이 옵니다.  
✅ `go 1.20`은 Go 버전을 나타냅니다.  

---

### 2.4 외부 패키지 추가 (`go get`)
Go에서는 `go get` 명령어를 사용하여 외부 패키지를 설치할 수 있습니다.

```sh
go get github.com/google/uuid
```
✅ `go get`을 실행하면 `go.mod`와 `go.sum` 파일이 자동으로 업데이트됩니다.  
✅ `go.sum` 파일은 다운로드된 패키지의 체크섬을 저장하여 무결성을 보장합니다.  

---

### 2.5 외부 패키지 사용
설치한 패키지는 `import` 문을 사용하여 사용할 수 있습니다.

```go
package main

import (
    "fmt"
    "github.com/google/uuid"
)

func main() {
    id := uuid.New()
    fmt.Println("생성된 UUID:", id)
}
```
✅ `uuid.New()`를 호출하여 새로운 UUID를 생성할 수 있습니다.  

---

## 3. 패키지와 모듈의 차이점

| 구분 | 패키지 (Package) | 모듈 (Module) |
|------|-----------------|--------------|
| **정의** | 관련된 Go 파일들의 묶음 | 패키지를 관리하는 상위 개념 |
| **파일 구조** | 하나의 폴더로 관리 | 여러 개의 패키지를 포함 |
| **사용 목적** | 코드 재사용과 논리적 그룹화 | 패키지 관리 및 의존성 해결 |
| **예제** | `mathutil.Add()` | `go get`으로 외부 패키지 설치 |

---

## 4. `go mod tidy` (불필요한 의존성 정리)

프로젝트를 관리하다 보면 사용하지 않는 패키지가 생길 수 있습니다.  
이때 `go mod tidy` 명령어를 실행하면 **불필요한 의존성을 정리**할 수 있습니다.

```sh
go mod tidy
```

✅ 사용되지 않는 패키지를 자동으로 정리해 줍니다.  

---

## 5. `go install` vs `go build` vs `go run`

| 명령어 | 설명 |
|--------|----------------------|
| `go run` | Go 파일을 실행 (컴파일하지 않음) |
| `go build` | Go 파일을 컴파일하여 실행 파일 생성 |
| `go install` | 빌드된 실행 파일을 `$GOPATH/bin`에 설치 |

```sh
go build main.go   # main.go 파일을 실행 파일로 빌드
go run main.go     # main.go 파일을 직접 실행
go install         # 실행 파일을 시스템에 설치
```

✅ `go build`는 실행 파일을 생성하지만 설치하지 않음.  
✅ `go install`은 `$GOPATH/bin`에 실행 파일을 저장하여 다른 곳에서도 실행 가능.  

---

## 6. 모듈과 패키지의 관계

1. **모듈(Module)은 패키지를 포함하는 개념입니다.**  
2. **패키지는 폴더 단위로 구분되며, `import` 문을 사용하여 가져올 수 있습니다.**  
3. **모듈을 초기화(`go mod init`)하면 `go.mod` 파일이 생성되며, 패키지의 의존성을 관리할 수 있습니다.**  

---

## 7. 결론
- **패키지(Package)**는 **코드를 논리적으로 그룹화**하는 단위입니다.
- **모듈(Module)**은 **패키지를 포함하며, 의존성 관리를 위한 도구**입니다.
- `go mod init`을 실행하면 **Go Modules**이 활성화되며, 패키지를 효율적으로 관리할 수 있습니다.
- `go get`을 사용하여 **외부 패키지를 설치**하고, `go mod tidy`로 **불필요한 패키지를 정리**할 수 있습니다.
