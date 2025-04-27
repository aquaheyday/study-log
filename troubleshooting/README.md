# 🧯 Troubleshooting - 실전 장애 해결 기록

서비스 운영 중 겪었던 다양한 문제 상황과 그에 대한 **원인 분석 및 해결 방법**을 정리한 공간입니다.  

---

## 📂 디렉토리 구성

| 폴더/파일 | 설명 |
|-----------|------|
| [docker](./docker) | 해결한 Docker Troubleshooting 기록 |
| [nextjs](./nextjs) | 해결한 Next.js Troubleshooting 기록 |


### 📋 Troubleshooting 목록
| 주제 | 파일명 | 설명 |
|---|---|---|
| Docker `no space left on device` | [docker-image-disk-full.md](./docker/docker-image-disk-full.md) | Docker로 Blue/Green 배포시 `no space left on device` 오류 |
| Next.js + Apache Revers Proxy `CORS` | [nextjs-api-cors-error.md](./nextjs/nextjs-api-cors-error.md) | Next.js + Apache Reverse Proxy 환경에서 CORS 오류 |
| `Apache` 로그 파일 용량 폭주 문제 | [log-overgrowth-and-rotatelogs.md](./apache/log-overgrowth-and-rotatelogs.md) | Apache access_log, error_log 등의 로그 파일이 지속적으로 증가하는 이슈 |
