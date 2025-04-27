# 📦 Next.js Project with Docker

Docker로 컨테이너화하고 Docker Compose로 관리하는 간단한 Next.js 프로젝트입니다.

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
git clone https://github.com/aquaheyday/study-log.git study-log/projects/docker/nextjs/
cd nextjs
```

### 3. 컨테이너 빌드 및 실행 (개발 모드)

```bash
docker-compose up -d
```

- `nextjs` 컨테이너가 생성됩니다.
- 브라우저에서 [http://localhost:3000](http://localhost:3000) 접속하면 Next.js 앱이 뜹니다.

---

## ⚙️ 주요 설정

- **Dockerfile**
  - `node:21` 이미지를 사용해 Next.js 앱을 빌드하고 실행합니다.
  - `npm run dev`로 개발 서버를 실행하거나 `npm run start`로 프로덕션 서버를 실행합니다.

- **docker-compose.yml**
  - `/app` 경로에 프로젝트 소스를 마운트합니다.
  - `node_modules`는 컨테이너 내부 전용으로 관리합니다.
  - 기본 포트는 `3000`입니다.

---

## 📄 코드 설명

**Dockerfile**

```Dockerfile
# 빌드 스테이지
FROM node:21 AS builder

WORKDIR /app

COPY ./app/package*.json ./
RUN npm ci

COPY ./app .

RUN npm run build

# 실행 스테이지
FROM node:21-slim

WORKDIR /app

COPY --from=builder /app/.next ./.next
COPY --from=builder /app/public ./public
COPY --from=builder /app/package*.json ./

RUN npm install --only=production

EXPOSE 3000

CMD ["npm", "run", "start"]
```

---

**docker-compose.yml**

```yaml
version: '3.8'

services:
  nextjs:
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "3000:3000"
    volumes:
      - ./app:/app
      - /app/node_modules
    environment:
      - NODE_ENV=development
```

---

## ✨ 추가 작업

- 개발 모드로 실행하고 싶다면 Dockerfile의 `CMD`를 `npm run dev`로 바꿔도 됩니다.
- 환경변수 `.env` 파일을 만들어서 `NEXT_PUBLIC_API_URL` 같은 설정도 추가할 수 있습니다.
