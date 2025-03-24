# 📁 Go 언어 파일 입출력 (File I/O)

Go에서는 `os`, `io`, `bufio`, `ioutil` 등 표준 패키지를 통해 텍스트 파일, 바이너리 파일을 손쉽게 읽고 쓸 수 있습니다.

---

## 1️⃣ 파일 열기 & 닫기

```go
file, err := os.Open("example.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()
```

✔ `os.Open()`은 읽기 전용  
✔ 열었으면 반드시 `defer file.Close()`로 닫기  

---

## 2️⃣ 파일 전체 읽기 (`ioutil` / `os.ReadFile`)

```go
data, err := os.ReadFile("example.txt")
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(data))
```

✔ Go 1.16 이후로 `ioutil.ReadFile` 대신 `os.ReadFile` 사용 권장  
✔ 작은 텍스트 파일 읽기에 적합  

---

## 3️⃣ 파일 줄 단위로 읽기 (`bufio.Scanner`)

```go
file, _ := os.Open("example.txt")
defer file.Close()

scanner := bufio.NewScanner(file)
for scanner.Scan() {
    fmt.Println(scanner.Text())
}
```

✔ 큰 파일을 줄 단위로 읽을 때 효율적  
✔ `scanner.Text()`로 한 줄씩 출력  

---

## 4️⃣ 파일 쓰기 (새로 쓰기 & 덮어쓰기)

```go
content := "Hello, Go!\n"
err := os.WriteFile("output.txt", []byte(content), 0644)
if err != nil {
    log.Fatal(err)
}
```

✔ `os.WriteFile()`은 파일이 없으면 생성, 있으면 덮어씀  
✔ 권한 `0644`는 읽기/쓰기 설정  

---

## 5️⃣ 파일 이어쓰기 (`os.OpenFile` + `O_APPEND`)

```go
f, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
    log.Fatal(err)
}
defer f.Close()

f.WriteString("추가된 내용\n")
```

✔ `os.O_APPEND`: 이어쓰기  
✔ `os.O_CREATE`: 파일 없으면 생성  
✔ `os.O_WRONLY`: 쓰기 전용  

---

## 6️⃣ 버퍼 기반 쓰기 (`bufio.Writer`)

```go
f, _ := os.Create("buffered.txt")
defer f.Close()

writer := bufio.NewWriter(f)
writer.WriteString("버퍼에 쓰고\n")
writer.WriteString("Flush로 실제 반영됨\n")
writer.Flush()
```

✔ `bufio.Writer`는 내부 버퍼에 먼저 쓰고 `Flush()`로 반영  
✔ 대용량 출력 시 효율적  

---

## 7️⃣ 파일 복사 (`io.Copy`)

```go
src, _ := os.Open("source.txt")
dst, _ := os.Create("copy.txt")
defer src.Close()
defer dst.Close()

_, err := io.Copy(dst, src)
if err != nil {
    log.Fatal(err)
}
```

✔ `io.Copy(dst, src)` 한 줄로 전체 복사 가능  

---

## 8️⃣ 파일 삭제 / 존재 여부 확인

```go
err := os.Remove("output.txt")
if err != nil {
    log.Fatal(err)
}
```

파일 존재 확인은 이렇게 처리 가능:

```go
if _, err := os.Stat("some.txt"); errors.Is(err, os.ErrNotExist) {
    fmt.Println("파일이 존재하지 않습니다")
}
```

✔ `os.Remove()`로 삭제  
✔ `os.Stat()` + `os.IsNotExist(err)`로 존재 확인  

---

## 9️⃣ 디렉토리 만들기

```go
err := os.Mkdir("mydir", 0755)
```

하위 폴더까지 한 번에 만들기 (예: `a/b/c`):

```go
os.MkdirAll("a/b/c", 0755)
```

✔ `0755`는 읽기/쓰기/실행 권한  
✔ `MkdirAll()`은 부모 디렉토리까지 자동 생성  

---

## 🔟 기능 요약

| 기능 | 방법 |
|------|------|
| 파일 열기 | `os.Open()`, `os.OpenFile()` |
| 전체 읽기 | `os.ReadFile()` |
| 줄 단위 읽기 | `bufio.Scanner` |
| 쓰기 | `os.WriteFile()`, `bufio.Writer` |
| 이어쓰기 | `os.OpenFile(..., O_APPEND)` |
| 복사 | `io.Copy()` |
| 삭제 | `os.Remove()` |
| 존재 확인 | `os.Stat()` + `errors.Is(..., os.ErrNotExist)` |
| 디렉토리 생성 | `os.Mkdir()`, `os.MkdirAll()` |

---

## 🎯 정리

✔ 파일 작업 전 항상 **오류 처리 필수**  
✔ `defer file.Close()`는 빠뜨리면 리소스 누수 발생  
✔ `WriteFile`은 항상 **덮어쓰기**, 이어쓰기할 땐 `OpenFile`  
✔ 큰 파일은 `bufio`로 처리하는 게 메모리 효율적  

