# 🧭 클라우드 운영 - IaC (Infrastructure as Code)

**IaC(Infrastructure as Code)** 는 서버, 네트워크, 스토리지 등 인프라 자원을 **코드로 정의하고 자동화된 방식으로 관리**하는 기술입니다.  
수작업 없이 코드로 인프라를 구성함으로써 일관성과 재현성을 보장합니다.

---

## 1️⃣ IaC란?

| 항목     | 설명 |
|----------|------|
| **정의** | 인프라 자원을 코드(스크립트, 선언형 설정 등)로 정의하고 자동으로 배포/관리하는 방식 |
| **목적** | 수작업 오류 방지, 환경 일관성 확보, 빠른 배포/롤백, DevOps 구현 |
| **핵심 철학** | "인프라도 애플리케이션처럼 코드로 관리하자" = 버전관리 + 테스트 + 자동화 가능 |

---

## 2️⃣ 대표 IaC 도구 비교

| 도구               | 형식     | 특징 |
|--------------------|----------|------|
| **Terraform**       | 선언형    | 멀티클라우드 지원, 상태파일 기반 인프라 추적 |
| **AWS CloudFormation** | 선언형 | AWS 전용, 서비스별 리소스 템플릿 지원 |
| **Pulumi**          | 프로그래밍형 | Python, TS 등 코드로 인프라 작성 가능 |
| **Ansible**         | 절차형    | 설정 관리에 강점, SSH 기반 |
| **Chef / Puppet**   | 절차형    | 서버 설정 중심, 에이전트 방식 |

---

## 3️⃣ 선언형 vs 절차형

| 항목       | 선언형 (Declarative)     | 절차형 (Imperative)       |
|------------|--------------------------|----------------------------|
| 예시 도구   | Terraform, CloudFormation | Ansible, Bash Script 등    |
| 방식       | "이렇게 되어야 한다" (결과 중심) | "이렇게 해라" (명령어 중심) |
| 장점       | 재현성, 상태 추적 용이           | 유연한 흐름 제어 가능         |
| 단점       | 디버깅 어려움, 복잡한 조건 분기 어려움 | 상태 추적 어려움             |

---

## 4️⃣ IaC 적용 흐름

```text
1. 인프라 정의
  - Terraform, YAML, JSON 등으로 서버, 네트워크, DB 등 작성

2. 코드 저장 및 버전 관리
  - Git에 저장 → 변경 이력 추적

3. 배포
  - CLI 또는 CI/CD 파이프라인 통해 자동 적용

4. 상태 확인 및 수정
  - 코드 수정 → 다시 적용 → 인프라 자동 업데이트
```

---

## 5️⃣ 실전 활용 예시

| 상황                                | IaC 활용 |
|-------------------------------------|----------|
| 신규 프로젝트 환경 1분만에 구성        | `terraform apply` 한 줄로 가능 |
| QA/스테이징/운영 환경 동일하게 유지     | 코드 기반으로 환경 재현 가능 |
| 배포 중 설정 누락/오류 자동 감지        | GitHub + IaC 도구로 변경 감시 |
| 재난복구/장애 복구 시 빠른 재배포        | 저장된 IaC 코드로 전체 인프라 복원 가능 |

---

## 🎯 정리 요약

✔ **IaC는 인프라를 코드로 다뤄서 자동화/버전관리/재현성을 확보하는 방식**  
✔ Terraform, CloudFormation 등 도구로 선언형 구성이 일반적  
✔ DevOps, CI/CD 파이프라인의 핵심 구성 요소  
✔ 수동 운영 리스크 줄이고, **신속하고 안정적인 인프라 운영을 가능하게 함**

