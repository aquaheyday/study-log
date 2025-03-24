# 🛠️ Go 언어 애플리케이션 빌드

Go는 **단일 실행 파일**로 애플리케이션을 빌드할 수 있는 언어입니다.  
이 문서에서는 Go의 빌드 방식과 옵션, 크로스 컴파일, 배포 준비 등을 정리합니다.

---

## 1️⃣ 기본 빌드 명령

```bash
go build
```

✔ 현재 디렉토리의 `main.go` 또는 `package main`을 찾아 실행 파일 생성  
✔ 기본적으로 현재 OS/아키텍처용 바이너리 생성

```bash
go build -o myapp
```

✔ 실행 파일 이름 지정  

---

## 2️⃣ 실행 없이 확인만 (컴파일 체크)

```bash
go build -n
```

✔ 어떤 명령이 실행될지 출력만  
✔ `-x`: 실제 컴파일 과정 출력

---

## 3️⃣ 빌드 + 실행

```bash
go run main.go
```

✔ `go build` + `./main`을 한 번에  
✔ 임시 실행용 (실행 파일 생성하지 않음)

---

## 4️⃣ 크로스 컴파일 (다른 플랫폼용 빌드)

```bash
GOOS=linux GOARCH=amd64 go build -o myapp-linux
GOOS=windows GOARCH=amd64 go build -o myapp.exe
```

| 환경변수 | 설명 |
|----------|------|
| `GOOS`   | 대상 운영체제 (linux, windows, darwin 등) |
| `GOARCH` | 아키텍처 (amd64, arm64, 386 등) |

✔ 크로스 컴파일도 별도 툴 없이 지원  
✔ 예: macOS에서 리눅스용 실행파일 만들기 가능

---

## 5️⃣ 릴리즈용 최적화 빌드

```bash
go build -ldflags="-s -w" -o myapp
```

| 옵션 | 설명 |
|------|------|
| `-s` | 심볼 정보 제거 |
| `-w` | 디버그 정보 제거 |

✔ 결과 바이너리 크기 줄어듦  
✔ `upx`를 통해 추가 압축도 가능

```bash
upx myapp
```

---

## 6️⃣ 빌드 시 버전 정보 삽입

```go
// main.go
var version = "dev"

func main() {
    fmt.Println("Version:", version)
}
```

```bash
go build -ldflags "-X main.version=1.0.0"
```

✔ `-X` 옵션으로 변수 값 삽입 가능  
✔ 릴리즈 자동화에 자주 활용

---

## 7️⃣ 빌드 캐시 관리

```bash
go clean -cache    # 빌드 캐시 삭제
go clean -modcache # 모듈 캐시 삭제
```

✔ 캐시 문제로 의심되는 경우 사용  
✔ 불필요한 캐시 정리로 디스크 절약

---

## 8️⃣ 디렉토리 구조 예시

```
myapp/
├── cmd/
│   └── main.go       // 진입점
├── internal/
│   └── logic.go
├── go.mod
```

```bash
cd cmd/
go build -o ../bin/myapp
```

✔ 실무에서는 `cmd/`, `internal/`, `pkg/` 등의 구조 많이 사용  
✔ `cmd/` 폴더는 여러 실행 진입점(main.go) 관리용

---

## 🎯 정리

✔ `Makefile`, `Taskfile.yaml`로 빌드 자동화 추천  
✔ 릴리즈 자동화에는 `goreleaser`, `GoReleaser Action`, `GitHub Actions` 많이 씀  
✔ 내부 라이브러리는 `internal/`로, 외부 공유 코드는 `pkg/`로 구분  
✔ 여러 바이너리 프로젝트는 `cmd/` 아래 디렉토리별 `main.go` 관리  
