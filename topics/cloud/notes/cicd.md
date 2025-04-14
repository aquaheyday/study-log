# 🧭 클라우드 운영 - CI/CD 파이프라인

**CI/CD**는 애플리케이션 개발부터 배포까지의 전체 흐름을 자동화하는 핵심 개발 전략입니다.  
CI(지속적 통합)는 코드 통합 및 테스트 자동화를, CD(지속적 배포 또는 제공)는 배포 자동화 및 무중단 배포를 목표로 합니다.

---

## 1️⃣ CI/CD란?

| 구분 | 설명 |
|------|------|
| **CI (Continuous Integration)** | 개발자들의 코드를 자동으로 병합, 빌드, 테스트하는 과정 |
| **CD (Continuous Delivery/Deployment)** | 테스트를 통과한 애플리케이션을 자동으로 배포하는 과정 (Delivery: 수동 승인 / Deployment: 완전 자동 배포) |

---

## 2️⃣ CI/CD 파이프라인 기본 흐름

```text
[개발자 코드 커밋]
        ↓
[CI] 코드 빌드 → 테스트 → 린트 검사 → 아티팩트 생성
        ↓
[CD] 스테이징 배포 → 승인(옵션) → 운영 배포
```

---

## 3️⃣ 주요 CI/CD 도구

| 도구            | 특징 |
|-----------------|------|
| **GitHub Actions** | GitHub과 연동 쉬움, 코드형 워크플로 |
| **GitLab CI/CD**   | Git 저장소와 통합된 올인원 도구 |
| **Jenkins**        | 플러그인 기반 유연한 오픈소스 CI 서버 |
| **CircleCI**       | 빠른 속도와 쉬운 설정 (YAML 기반) |
| **ArgoCD**         | 쿠버네티스 환경에서 GitOps 기반 CD 도구 |
| **CodePipeline (AWS)** | AWS 리소스와 쉽게 통합되는 CI/CD 관리형 서비스 |

---

## 4️⃣ 구성 요소

| 구성 요소         | 설명 |
|------------------|------|
| **빌드(Build)**     | 코드 → 실행 가능한 앱/컨테이너 이미지 생성 |
| **테스트(Test)**   | 단위/통합 테스트 자동 수행 |
| **배포(Deploy)**   | 스테이징/운영 환경에 앱 자동 배포 |
| **알림(Notification)** | 실패/성공 여부를 슬랙, 이메일 등으로 통지 |
| **승인(Approval)** | 운영 배포 전 수동 승인 가능 (Continuous Delivery의 특징) |

---

## 5️⃣ 실전 예시 (GitHub Actions 기반)

```yaml
name: CI/CD Pipeline

on:
  push:
    branches: [ main ]

jobs:
  build-and-deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - name: Install dependencies
        run: npm install
      - name: Run tests
        run: npm test
      - name: Build project
        run: npm run build
      - name: Deploy to S3
        run: aws s3 sync ./build s3://my-bucket --delete
```

---

## 6️⃣ CI/CD 도입 효과

| 효과               | 설명 |
|--------------------|------|
| ✅ 빠른 피드백       | 코드 푸시 직후 테스트/결과 확인 가능 |
| ✅ 배포 자동화       | 릴리즈 속도 향상 + 수작업 실수 방지 |
| ✅ 코드 품질 개선    | 테스트/정적 분석 자동화로 코드 안정성 ↑ |
| ✅ DevOps 문화 강화 | 개발팀과 운영팀의 협업 효율 ↑ |

---

## 🎯 정리 요약

✔ **CI/CD는 코드 통합 → 빌드 → 테스트 → 배포까지의 자동화 파이프라인**  
✔ CI는 테스트 자동화, CD는 배포 자동화를 의미  
✔ GitHub Actions, Jenkins, ArgoCD 등 다양한 도구 존재  
✔ 자동화로 **개발 속도 + 품질 + 운영 안정성**을 동시에 확보

