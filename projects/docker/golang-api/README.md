# 📦 Golang API with Docker Compose

Docker로 컨테이너화하고 Docker Compose를 사용해 관리하는 간단한 Golang API 프로젝트입니다.

---

## 🚀 실행 방법

### 1. Docker & Docker Compose 설치 확인

```bash
docker --version
docker-compose --version
```

> Docker와 Docker Compose가 설치되어 있어야 합니다.

### 2. 프로젝트 클론

```bash
git clone https://github.com/aquaheyday/study-log.git study-log/projects/docker/golang-api/
cd golang-api
```

### 3. 컨테이너 빌드 및 실행

```bash
docker-compose up -d
```

- `golang-api-production` 이라는 이름의 컨테이너가 생성됩니다.
- `localhost:8080` 에서 API를 확인할 수 있습니다.

### 4. 서버 접속

```bash
curl http://localhost:8080
```

응답 예시:

```text
Hello, Golang with Docker Compose!
```

---

## ⚙️ 주요 설정

- **Docker Compose**:
  - `docker-compose.yml` 파일을 통해 서비스 및 네트워크를 정의합니다.
  - `restart: unless-stopped` 옵션으로 서버 자동 복구 설정을 합니다.
- **Dockerfile**:
  - `golang:1.20` 이미지를 사용해 빌드 후, 경량 `alpine` 이미지를 이용해 멀티스테이지 빌드로 최적화합니다.
- **Environment Variables**:
  - `GIN_MODE=release`로 설정하여 배포 모드로 실행합니다.
- **Networking**:
  - `production-network` 라는 Docker 네트워크를 사용합니다.

---

## 📄 코드 설명

**`main.go`**

```go
package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Golang with Docker Compose!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
```

- `/` 경로로 HTTP 요청이 오면 간단한 텍스트를 반환하는 서버입니다.
