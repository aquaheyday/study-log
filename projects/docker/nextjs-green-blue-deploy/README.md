# Next.js Docker Green/Blue Deployment

> 이 레포는 **Next.js 앱을 Docker 환경에서 Green/Blue 배포 방식으로 무중단 배포**하는 방법을 담고 있습니다.  
> Apache 리버스 프록시를 통해 트래픽 전환을 제어하며, `deploy.sh` 스크립트를 통해 자동화된 전환을 수행합니다.

---

## 🚀 구성 개요

### 💡 Green/Blue 배포란?
Green/Blue 배포는 **두 개의 독립적인 서비스 버전**(green, blue)을 번갈아 실행하면서  
트래픽을 안정적으로 전환해 **무중단 배포**를 실현하는 전략입니다.

- `nextjs-app-blue`와 `nextjs-app-green` 두 컨테이너를 운영
- Apache가 현재 활성 컨테이너로 트래픽을 전달
- `deploy.sh` 실행 시 다음 버전을 build & up → Apache config 수정 → 이전 버전 down

---

## 🧱 구조

```text
.
├── Dockerfile                  # Next.js 앱 Docker 이미지 빌드 설정
├── docker-compose.yml          # Blue/Green 컨테이너 정의 및 포트 설정
├── deploy.sh                   # 자동 배포 및 전환 스크립트
├── app/                        # 실제 Next.js 앱 소스 코드
│   ├── package.json
│   ├── next.config.js
│   └── ...
└── README.md
```

---

## ⚙️ 사용법

### 1. 사전 준비
Apache 설정에서 `httpd-vhost.conf`, `ssl.conf` 내 리버스 프록시 구문이 아래와 같이 있어야 합니다.

```apache
# httpd-vhost.conf 예시
ProxyPass / http://127.0.0.1:6960/
ProxyPassReverse / http://127.0.0.1:6960/
```

- 포트 `6960`, `6961`은 각각 Blue, Green 컨테이너로 매핑됨
- `sudo` 권한으로 Apache 설정을 수정하고 재시작해야 하므로 `deploy.sh` 실행 시 권한이 필요함

---

### 2. 배포

```bash
chmod +x deploy.sh
./deploy.sh
```

- 현재 활성 컨테이너를 판별
- 비활성 컨테이너를 빌드 및 실행
- Apache 설정 파일을 수정하여 트래픽을 새 컨테이너로 전환
- 이전 컨테이너는 자동 중단 및 제거됨

---

## 🐳 Docker 구성

### Dockerfile (멀티 스테이지)

- 빌드 전용 단계에서 `npm run build`
- 운영 이미지에서는 `node:18-alpine`으로 경량화

### docker-compose.yml

- `nextjs-app-blue` → `6960:3000`
- `nextjs-app-green` → `6961:3000`
- 공통 네트워크 `nextjs-blue-green-network` 사용

---

## 📂 환경 변수

- Dockerfile 내부에 `NODE_ENV=production` 설정 포함
