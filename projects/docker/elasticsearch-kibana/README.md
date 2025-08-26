# 📦 ElasticSearch + Kibana with Docker Compose

Docker로 컨테이너화하고 Docker Compose를 사용해 관리하는 간단한 ElasricSearch + Kibana 프로젝트입니다.

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
git clone https://github.com/aquaheyday/study-log.git study-log/projects/docker/elasticsearch-kibana/
cd elasticsearch-kibana
```

### 3. 컨테이너 빌드 및 실행

```bash
docker-compose up -d
```

- `localhost:5601` 에서 Kibana GUI를 확인할 수 있습니다.

---
