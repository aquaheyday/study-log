# 📦 Go 언어 패키징과 배포

Go는 컴파일된 **단일 바이너리**로 애플리케이션을 패키징할 수 있어 CLI 도구나 백엔드 서버 등을 간편하게 배포할 수 있습니다.  
이 문서는 **Go 모듈 구조, 빌드, 압축, 릴리즈 자동화, 배포 방법** 등을 다룹니다.

---

## 1️⃣ Go 모듈 구조

```text
myapp/
├── go.mod
├── main.go
├── cmd/
│   └── myapp/
│       └── main.go
├── internal/
│   └── logic.go
└── pkg/
    └── helper/
        └── util.go
```

| 디렉토리 | 역할 |
|----------|------|
| `cmd/`     | 실행 바이너리 별 진입점 (여러 CLI/서버 지원 시) |
| `internal/` | 외부에서 접근 불가능한 내부 코드 |
| `pkg/`      | 다른 프로젝트에서 가져다 쓸 수 있는 재사용 코드 |

---

## 2️⃣ 실행 파일 빌드

```bash
go build -o myapp ./cmd/myapp
```

✔ `-o`로 실행 파일 이름 지정  

#### 크로스 컴파일 예시
```bash
GOOS=linux GOARCH=amd64 go build -o myapp-linux
```

---

## 3️⃣ 빌드 최적화 (배포용)

```bash
go build -ldflags="-s -w" -o myapp
```

| 옵션 | 설명 |
|------|------|
| `-s` | 심볼 정보 제거 |
| `-w` | 디버그 정보 제거 |

✔ 결과 파일 크기 줄어듦  
✔ `upx`로 압축도 가능:

```bash
upx myapp
```

---

## 4️⃣ 버전 정보 삽입 (릴리즈용)

```go
var version = "dev"

func main() {
    fmt.Println("Version:", version)
}
```

```bash
go build -ldflags "-X main.version=1.0.0"
```

✔ `-X 패키지.변수=값`으로 컴파일 시점에 값 삽입 가능  
✔ Git 커밋, 빌드 날짜 등도 삽입 가능

---

## 5️⃣ 압축 및 패키징

### ZIP / TAR 압축

```bash
zip myapp.zip myapp README.md
tar -czvf myapp.tar.gz myapp README.md
```

✔ 릴리즈에 필요한 파일만 포함  
✔ 플랫폼별로 압축본 따로 생성

---

## 6️⃣ 자동화 배포 (Goreleaser)

#### 1. 설치

```bash
brew install goreleaser
```

#### 2. 설정 (goreleaser.yaml)

```yaml
project_name: myapp
builds:
  - main: ./cmd/myapp
    goos: [linux, windows, darwin]
    goarch: [amd64, arm64]
archives:
  - format: tar.gz
    files:
      - README.md
release:
  github:
    owner: yourname
    name: myapp
```

#### 3. 릴리즈 실행

```bash
goreleaser release --clean
```

✔ GitHub 릴리즈에 자동 업로드  
✔ 플랫폼별 실행 파일 + 압축본 자동 생성

---

## 7️⃣ 배포 방식

| 방식 | 설명 |
|------|------|
| GitHub Release | Goreleaser로 릴리즈 자동화 |
| AWS S3, GCS | CLI 도구 업로드 / 공유 가능 |
| Docker | 컨테이너 이미지 빌드 후 배포 |
| SCP / FTP | 서버에 직접 복사 배포 |
| CI/CD | GitHub Actions, GitLab CI, Jenkins 등과 연동 |

---

## 8️⃣ 실전 배포 예시 (Linux 서버에 복사)

```bash
GOOS=linux GOARCH=amd64 go build -o myapp
scp myapp user@host:/usr/local/bin/
```

✔ 서버에 직접 복사하여 CLI/서버 실행  
✔ systemd로 서비스 등록도 가능

---

## 🎯 정리

✔ 릴리즈 전 `go mod tidy`, `go vet`, `go test`로 상태 점검  
✔ 실행 파일 하나만 있으면 배포가 간편한 게 Go의 큰 장점  
✔ 여러 플랫폼 배포는 Goreleaser로 자동화하면 생산성이 크게 향상됨  
✔ 버전 정보는 `-ldflags` 로 삽입하는 방식이 실무에서 가장 널리 사용됨
