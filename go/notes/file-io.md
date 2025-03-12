# Go 언어: 파일 입출력(File I/O)

Go 언어에서 파일 입출력은 `os` 및 `io/ioutil` 패키지를 사용하여 간단하게 처리할 수 있습니다.

---

## 1. 파일 쓰기 (Write to File)
### 1.1 기본 파일 쓰기
```go
package main

import (
    "fmt"
    "os"
)

func main() {
    content := "Hello, Go File I/O!"
    err := os.WriteFile("example.txt", []byte(content), 0644)
    if err != nil {
        fmt.Println("Error:", err)
    }
}
```

### 1.2 파일을 열어 쓰기 (os.OpenFile)
```go
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.OpenFile("example.txt", os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    _, err = file.WriteString("Appending some text\n")
    if err != nil {
        fmt.Println("Error:", err)
    }
}
```

### 주요 패턴
- `os.WriteFile(filename, []byte, perm)`: 한 번에 파일 작성
- `os.OpenFile(filename, flags, perm)`: 파일을 열어 여러 방식으로 조작 가능
- `file.WriteString("text")`: 문자열 쓰기

---

## 2. 파일 읽기 (Read from File)
### 2.1 파일 전체 읽기
```go
package main

import (
    "fmt"
    "os"
)

func main() {
    data, err := os.ReadFile("example.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("File Content:", string(data))
}
```

### 2.2 파일을 한 줄씩 읽기
```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
    
    if err := scanner.Err(); err != nil {
        fmt.Println("Error:", err)
    }
}
```

### 주요 패턴
- `os.ReadFile(filename)`: 전체 파일 읽기
- `bufio.NewScanner(file)`: 파일을 한 줄씩 읽기

---

## 3. 파일 존재 여부 확인
```go
package main

import (
    "fmt"
    "os"
)

func fileExists(filename string) bool {
    _, err := os.Stat(filename)
    return !os.IsNotExist(err)
}

func main() {
    filename := "example.txt"
    if fileExists(filename) {
        fmt.Println("File exists")
    } else {
        fmt.Println("File does not exist")
    }
}
```

### 주요 패턴
- `os.Stat(filename)`: 파일 정보 조회
- `os.IsNotExist(err)`: 파일 존재 여부 확인

---

## 4. 파일 삭제
```go
package main

import (
    "fmt"
    "os"
)

func main() {
    err := os.Remove("example.txt")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("File deleted successfully")
    }
}
```

### 주요 패턴
- `os.Remove(filename)`: 파일 삭제

---

## 5. 디렉터리 생성 및 삭제
### 5.1 디렉터리 생성
```go
os.Mkdir("testdir", 0755) // 단일 디렉터리 생성
os.MkdirAll("parent/child", 0755) // 중첩된 디렉터리 생성
```

### 5.2 디렉터리 삭제
```go
os.Remove("testdir") // 단일 디렉터리 삭제
os.RemoveAll("parent") // 하위 디렉터리까지 삭제
```

---

## 6. 파일 복사
```go
package main

import (
    "fmt"
    "io"
    "os"
)

func copyFile(src, dst string) error {
    sourceFile, err := os.Open(src)
    if err != nil {
        return err
    }
    defer sourceFile.Close()

    destFile, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer destFile.Close()

    _, err = io.Copy(destFile, sourceFile)
    return err
}

func main() {
    err := copyFile("example.txt", "copy.txt")
    if err != nil {
        fmt.Println("Copy failed:", err)
    } else {
        fmt.Println("File copied successfully")
    }
}
```

### 주요 패턴
- `io.Copy(dst, src)`: 파일 내용을 복사

---

## 7. 파일 입출력 베스트 프랙티스
**파일을 열었으면 `defer file.Close()`로 안전하게 닫기**
**에러 체크(`if err != nil { }`)를 철저히 하기**
**버퍼를 사용하여 효율적으로 파일 읽기/쓰기 (`bufio` 활용)**
**파일이 존재하는지 확인 후 작업 진행**
