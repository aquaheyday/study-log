# 📦 Redis with Docker Compose

Docker로 컨테이너화하고 Docker Compose로 관리하는 간단한 Redis 서비스입니다.

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
git clone https://github.com/aquaheyday/study-log.git study-log/projects/docker/redis/
cd redis
```

### 3. 컨테이너 빌드 및 실행

```bash
docker-compose up -d
```

- `event_redis`라는 이름의 Redis 컨테이너가 생성됩니다.
- `localhost:6379` 포트로 Redis 서버에 접근할 수 있습니다.

---

## ⚙️ 주요 설정

- **Redis 비밀번호 설정**
  - `REDIS_PASSWORD=pwd`
  - 컨테이너 실행 시 `--requirepass pwd` 옵션으로 비밀번호 인증을 강제합니다.

- **Persistent 데이터 저장**
  - Redis 데이터를 `redis_data`라는 Docker 볼륨에 영구 저장합니다.
  - Redis 설정 옵션 `--appendonly yes`를 통해 AOF(Append Only File) 모드 활성화.

- **컨테이너 기본 세팅**
  - `restart: always` 옵션으로 컨테이너 자동 재시작 설정.

---

## 📄 코드 설명

**docker-compose.yml**

```yaml
version: "3.8"

services:
  redis:
    image: redis:latest
    container_name: event_redis
    restart: always
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data
    environment:
      - REDIS_PASSWORD=pwd # Redis 비밀번호 (원하면 제거 가능)
    command: ["redis-server", "--appendonly", "yes", "--requirepass", "pwd"]

volumes:
  redis_data:
    driver: local
```
