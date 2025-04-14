# ☁️ Cloud - 클라우드 인프라 및 서비스

클라우드 환경에서의 인프라 구성, 서비스 모델, 주요 컴포넌트, 자동화, 비용 최적화 전략 등을 정리합니다.  
AWS, GCP, Azure 등의 주요 클라우드 플랫폼과 멀티클라우드 트렌드까지 포함합니다.

---

### 🌍 클라우드 개요 및 서비스 모델

| 주제 | 파일명 | 설명 |
|------|--------|------|
| 클라우드란? | [what-is-cloud.md](./notes/what-is-cloud.md) | 정의, 특징, 전통 인프라 대비 장점 |
| IaaS / PaaS / SaaS | [service-models.md](./notes/service-models.md) | 서비스 모델 3단계 비교 |
| 퍼블릭 vs 프라이빗 vs 하이브리드 | [cloud-types.md](./notes/cloud-types.md) | 환경별 장단점 비교 |

---

### 🧰 클라우드 핵심 컴포넌트

| 주제 | 파일명 | 설명 |
|------|--------|------|
| 가상 서버 (EC2 등) | [compute.md](./notes/compute.md) | 인스턴스, 오토스케일링, VM 관리 |
| 네트워크 구성 (VPC) | [vpc.md](./notes/vpc.md) | 서브넷, 게이트웨이, 보안 그룹 |
| 스토리지 서비스 | [storage.md](./notes/storage.md) | S3, EBS, 객체/블록/파일 스토리지 |
| 데이터베이스 서비스 | [managed-db.md](./notes/managed-db.md) | RDS, DynamoDB, Cloud SQL 등 |
| DNS와 로드밸런서 | [dns-lb.md](./notes/dns-lb.md) | Route 53, ALB/ELB 구조 |
| IAM (권한 관리) | [iam.md](./notes/iam.md) | 사용자, 역할, 정책 기반 접근 제어 |

---

### 🛠️ DevOps & 자동화

| 주제 | 파일명 | 설명 |
|------|--------|------|
| IaC 개념 | [iac.md](./notes/iac.md) | Terraform, CloudFormation 개요 |
| CI/CD 파이프라인 | [cicd.md](./notes/cicd.md) | 자동 빌드/테스트/배포 구성 |
| 모니터링 & 로깅 | [monitoring.md](./notes/monitoring.md) | CloudWatch, Stackdriver 등 |
| 비용 최적화 전략 | [cost-optimization.md](./notes/cost-optimization.md) | Reserved 인스턴스, 비용 분석 도구 등 |

---

### 🧱 서버리스 & 매니지드 서비스

| 주제 | 파일명 | 설명 |
|------|--------|------|
| 서버리스란? | [serverless.md](./notes/serverless.md) | Lambda, Cloud Functions 개요 |
| Event-driven 아키텍처 | [event-driven.md](./notes/event-driven.md) | SQS, SNS, Pub/Sub 연동 방식 |
| BaaS/FaaS 비교 | [baas-faas.md](./notes/baas-faas.md) | Backend-as-a-Service vs Function-as-a-Service |

---

### 🧭 클라우드 설계 & 보안 전략

| 주제 | 파일명 | 설명 |
|------|--------|------|
| 멀티 리전/가용영역 설계 | [multi-region.md](./notes/multi-region.md) | 고가용 아키텍처 설계 전략 |
| 백업 & 재해복구 (DR) | [backup-dr.md](./notes/backup-dr.md) | 스냅샷, 복제, RTO/RPO 설계 |
| 클라우드 보안 기본 | [cloud-security.md](./notes/cloud-security.md) | 데이터 암호화, 접근제어, 인증 정책 |
| CSPM / CWPP | [cspm-cwpp.md](./notes/cspm-cwpp.md) | 클라우드 보안 도구 및 자동 점검 전략 |
