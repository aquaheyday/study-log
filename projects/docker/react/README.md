# 📦 React App Deployment with Docker and Nginx

Docker로 컨테이너화하고 Nginx를 통해 서비스하는 간단한 React 애플리케이션입니다.

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
git clone https://github.com/aquaheyday/study-log.git study-log/projects/docker/react/
cd react
```

### 3. 컨테이너 빌드 및 실행

```bash
docker-compose up -d
```

- `react-app`이라는 이름의 컨테이너가 생성됩니다.
- 브라우저에서 [http://localhost](http://localhost) 로 접속하면 React 앱이 표시됩니다.

---

## ⚙️ 주요 설정

- **Dockerfile**
  - `node:18` 이미지를 사용하여 React 앱을 빌드합니다.
  - `nginx:alpine` 이미지를 사용하여 빌드된 정적 파일을 서빙합니다.
  - `/usr/share/nginx/html`에 React 빌드 파일(`dist`)을 복사합니다.

- **docker-compose.yml**
  - React 앱을 빌드하고 80번 포트에 노출합니다.
  - 컨테이너 자동 재시작(`restart: always`) 설정.

- **nginx.conf**
  - 모든 요청을 `index.html`로 리다이렉트하여 React 라우팅을 지원합니다.
  - 404 에러도 `index.html`로 처리합니다.

---

## 📄 코드 설명

**Dockerfile**

```Dockerfile
FROM node:18 AS build

WORKDIR /app

COPY ./app/package*.json ./

RUN npm install

COPY ./app ./

RUN npm run build

FROM nginx:alpine

COPY --from=build /app/dist /usr/share/nginx/html

COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
```

---

**docker-compose.yml**

```yaml
version: '3.8'

services:
  react-app:
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "80:80"
    restart: always
```

---

**nginx.conf**

```nginx
server {
    listen 80;

    server_name localhost;

    root /usr/share/nginx/html;

    index index.html;

    location / {
        try_files $uri /index.html;
    }

    error_page 404 /index.html;
}
```
