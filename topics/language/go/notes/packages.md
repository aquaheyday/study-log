# 📦 Go 언어 패키지와 모듈

Go는 **패키지(package)** 를 기반으로 코드를 구조화하고, **모듈(module)** 을 통해 외부 패키지 의존성과 버전을 관리합니다.  

---

## 1️⃣ 패키지(Package)란?

- Go의 기본 코드 단위 (모든 파일은 `package` 선언으로 시작)
- 하나의 디렉토리가 하나의 패키지를 의미
- 파일 상단에 `package` 키워드로 정의

```go
// 파일: utils/math.go
package utils

func Add(a, b int) int {
    return a + b
}
```

#### 패키지 가져오기 (import)

```go
// 파일: main.go
package main

import (
    "fmt"
    "myapp/utils"
)

func main() {
    result := utils.Add(3, 4)
    fmt.Println(result)
}
```

✔ 동일 모듈 내에 있는 패키지는 상대 경로로 import 가능  
✔ 외부 패키지는 `go get`을 통해 설치  

---

## 2️⃣ 모듈(Module) 이란?

- **Go 1.11+** 에 도입된 **의존성 관리 시스템**
- `go.mod` 파일로 프로젝트 경로, 패키지 버전 등을 관리

### 1) 모듈 초기화

```bash
go mod init myapp
```

✔ 현재 디렉토리에 `go.mod` 생성  
✔ `myapp`은 모듈 경로 (예: `github.com/username/myapp`)  

---

### 2) 모듈 구조 예시

```text
myapp/
├── go.mod
├── go.sum
├── main.go
└── utils/
    └── math.go  (package utils)
```

---

## 3️⃣ 외부 패키지 설치

```bash
go get github.com/gin-gonic/gin
```

✔ `go.mod`에 버전 추가됨  
✔ `go.sum`에 무결성 해시 저장됨  

---

## 4️⃣ 주요 `go mod` 명령어

| 명령어 | 설명 |
|--------|------|
| `go mod init` | 모듈 생성 |
| `go get 패키지` | 외부 패키지 설치 |
| `go mod tidy` | 불필요한 의존성 정리 |
| `go mod download` | 의존성 다운로드 |
| `go list -m all` | 모든 모듈 목록 확인 |

---

## 5️⃣ 모듈 버전 지정

```bash
go get github.com/gin-gonic/gin@v1.8.2
```

- 원하는 버전 명시 가능  
- 필요 시 `replace`로 로컬 경로 대체 가능

```go
// go.mod
replace mylib => ../local/mylib
```

---

## 6️⃣ 내부 패키지 분리 예시

```text
myapp/
├── go.mod
├── main.go
├── service/
│   └── user.go       (package service)
└── repository/
    └── user_repo.go  (package repository)
```

✔ 각 폴더는 고유 패키지 이름 사용 (`package service`, `package repository`)  
✔ import 시 `myapp/service`, `myapp/repository` 형식 사용  

---

## 7️⃣ 모듈 경로 주의사항

- 모듈 경로는 보통 **GitHub 경로를 사용** (`github.com/username/project`)
- 모듈을 **다른 프로젝트에서 import** 하려면 반드시 올바른 **모듈명 + 공개 저장소** 경로여야 함  

---

## 8️⃣ `go install`과 빌드

```bash
go install ./...
```

✔ 모듈 내 전체 패키지를 컴파일하고 실행 파일은 `$GOBIN`에 저장  
✔ 모듈 기반 프로젝트는 `$GOPATH` 설정 없이도 관리 가능

---

## 🎯 정리

✔ 패키지는 Go 코드의 구성 단위이며, 모듈은 프로젝트와 의존성을 관리하는 상위 개념입니다.  
✔ `go mod init`, `go get`, `go mod tidy` 등으로 외부 패키지를 깔끔하게 관리할 수 있습니다.  
✔ 패키지를 잘게 나누고 모듈화하면 **코드 재사용성과 유지보수성**이 크게 향상됩니다.

