# ⚙️ Go 개발 환경 설정

Go 언어는 **간단한 설치**와 **빠른 실행 환경**을 제공하여 빠르게 개발을 시작할 수 있습니다.

---

## 1️⃣ Go 설치

#### 1. 공식 다운로드

- Go 공식 사이트: [https://go.dev/dl](https://go.dev/dl)

#### 2. 운영체제별 설치

| OS | 설치 방법 |
|----|-----------|
| Windows | `.msi` 설치 파일 다운로드 후 실행 |
| macOS   | `brew install go` 또는 공식 `.pkg` 설치 |
| Linux   | `.tar.gz` 압축 파일을 `/usr/local/`에 풀기 |

#### 3. 설치 확인

```bash
go version
```

---

## 2️⃣ 기본 환경 변수 설정

설치 후 `go env` 명령어로 환경 정보를 확인할 수 있습니다.

```bash
go env
```

| 환경 변수 | 설명 |
|-----------|------|
| `GOROOT`  | Go 설치 디렉토리 (보통 자동 설정됨) |
| `GOPATH`  | Go 작업 공간 디렉토리 (기본값: `$HOME/go`) |
| `GOBIN`   | Go 실행 파일 저장 위치 (`$GOPATH/bin`) |

✔ 대부분의 경우 `GOPATH`는 따로 설정하지 않고 **Go Modules**을 사용합니다.  

---

## 3️⃣ Go Modules 사용 (Go 1.11+)

Go Modules는 **Go의 의존성 관리 시스템**입니다.  
현재는 **GOPATH 대신 Go Modules 방식이 표준**입니다.

#### 1. 모듈 초기화

```bash
go mod init myapp
```

✔ 명령어 실행 시 `go.mod` 파일 생성됨

#### 2. 패키지 설치

```bash
go get github.com/gin-gonic/gin
```

✔ 의존성 자동으로 `go.sum`에 기록되고 관리됨

---

## 4️⃣ 프로젝트 구조 예시

```text
myapp/
├── go.mod
├── go.sum
├── main.go
└── handlers/
    └── user.go
```

✔ `go.mod` / `go.sum`: 모듈과 의존성 정보  
✔ `main.go`: 애플리케이션 진입점  
✔ `handlers/`: 기능별 모듈 분리

---

## 5️⃣ 코드 실행 & 빌드

```bash
go run main.go        # 실행
go build              # 바이너리 빌드
go install            # 실행 파일을 $GOBIN 에 설치
```

---

## 6️⃣ 추천 개발 도구

| 도구 / IDE | 설명 |
|------------|------|
| VS Code    | 경량 IDE, Go 확장 기능 풍부 (Go 공식 확장 설치 필요) |
| GoLand     | JetBrains의 Go 전문 IDE (강력한 자동완성, 리팩토링) |
| LiteIDE    | Go 전용 무료 IDE |
| Vim/Neovim | Go 플러그인과 함께 사용 가능 (gopls, gofmt 등) |

---

### VS Code 필수 확장

- `Go` by Google
- 자동 포맷팅, IntelliSense, 테스트 실행, 디버깅 지원

---

## 7️⃣ Go 개발에 유용한 명령어

| 명령어 | 설명 |
|--------|------|
| `go mod tidy` | 사용하지 않는 의존성 제거 및 누락된 패키지 추가 |
| `go fmt`      | 코드 포맷 정리 (자동 들여쓰기, 스타일 통일) |
| `go test`     | 테스트 실행 |
| `go doc`      | 패키지 문서 출력 |
| `go install`  | 실행 파일 컴파일 후 `$GOBIN`에 저장 |

---

## 8️⃣ 디버깅 및 Lint 도구 추천

| 도구        | 설명 |
|-------------|------|
| `dlv`       | Go 공식 디버거 (`go install github.com/go-delve/delve/cmd/dlv@latest`) |
| `golangci-lint` | 코드 정적 분석 및 스타일 체크 |
| `gopls`     | Go 공식 언어 서버 (IntelliSense, 타입 분석) |

---

## 🎯 정리

✔ Go는 설치가 간단하며, `go mod`를 사용해 현대적인 모듈 기반 개발을 지원합니다.  
✔ `VS Code` 또는 `GoLand`를 사용하면 강력한 자동완성, 디버깅, 테스트 기능 활용 가능  
✔ `go run`, `go build`, `go test`, `go mod tidy` 등 명령어로 프로젝트를 쉽고 빠르게 관리할 수 있습니다.  
✔ GOPATH 대신 `Go Modules` 기반 프로젝트 구조가 기본이며, CI/CD 및 Docker 배포 환경과도 자연스럽게 연동됩니다.
