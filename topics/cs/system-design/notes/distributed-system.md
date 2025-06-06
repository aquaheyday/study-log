# 🛰️ 시스템 설계 - 분산 시스템 (Distributed Systems)

**분산 시스템(Distributed System)** 은 여러 대의 컴퓨터가 **네트워크로 연결되어 하나의 시스템처럼 동작하는 구조**입니다.  
확장성, 고가용성, 성능 향상 등을 위해 현대 서비스 대부분이 분산 아키텍처를 채택하고 있습니다.

---

## 1️⃣ 분산 시스템이란?

| 항목     | 설명 |
|----------|------|
| **정의** | 서로 다른 노드(Node)들이 네트워크를 통해 협력하여 하나의 시스템처럼 동작하는 구조 |
| **목적** | 단일 서버의 한계를 극복하고, **확장성, 장애 허용성, 고가용성** 확보 |
| **예시** | Google 검색, Netflix, 카카오톡, AWS/GCP 클라우드 등

---

## 2️⃣ 분산 시스템의 핵심 특성

| 특성         | 설명 |
|--------------|------|
| **확장성(Scalability)** | 노드를 수평적으로 확장하여 처리량 증가 가능 |
| **고가용성(Availability)** | 일부 노드가 죽어도 전체 서비스는 유지 |
| **장애 허용성(Fault Tolerance)** | 네트워크/노드 장애 발생 시 자동 복구 가능 |
| **일관성(Consistency)** | 모든 노드에서 동일한 데이터 상태 유지 |
| **지연 허용성(Latency Tolerance)** | 네트워크 지연이나 오류를 견딜 수 있는 유연성 |
| **분산성(Decentralization)** | 단일 장애 지점(SPOF) 없이 분산 처리

---

## 3️⃣ 분산 시스템의 도전 과제

| 문제                      | 설명 |
|---------------------------|------|
| **데이터 일관성 유지**         | 분산 환경에서 동일한 데이터 상태를 보장하기 어려움 (CAP 이론 관련) |
| **장애 탐지 및 복구**         | 노드/네트워크 장애를 정확히 감지하고 빠르게 복구해야 함 |
| **시간 동기화**              | 여러 시스템 간 시간 차이 존재 → 타임스탬프 처리 이슈 |
| **분산 트랜잭션 관리**        | 다수 노드 간 원자적 트랜잭션 처리 어려움 |
| **네트워크 분할 및 지연**      | 패킷 유실, 지연 등으로 인한 데이터 유실/중복 처리 가능성 |
| **복잡성 증가**              | 시스템 아키텍처, 테스트, 배포, 모니터링 복잡

---

## 4️⃣ CAP 이론 (CAP Theorem)

| 항목    | 설명 |
|---------|------|
| **Consistency** | 모든 노드에 동일한 데이터가 동시에 반영됨 |
| **Availability** | 모든 요청에 대해 응답을 받을 수 있음 |
| **Partition Tolerance** | 네트워크 단절(Partition) 상황에서도 시스템이 동작함 |

> **CAP 이론에 따르면, 3가지 중 동시에 2가지만 보장 가능**

| 조합       | 시스템 예시 |
|------------|-------------|
| CA         | 단일 서버 시스템 (분산 아님) |
| CP         | HDFS, MongoDB (Partition 시 가용성 희생) |
| AP         | Cassandra, DynamoDB (일관성보다 가용성 우선) |

---

## 5️⃣ 대표적인 분산 시스템 예시

| 시스템 유형       | 예시 |
|------------------|------|
| **분산 파일 시스템** | HDFS, Ceph, GlusterFS |
| **분산 데이터베이스** | Cassandra, MongoDB, DynamoDB |
| **분산 캐시**         | Redis Cluster, Memcached |
| **분산 메시징 시스템** | Kafka, RabbitMQ, Pulsar |
| **분산 컴퓨팅**        | Hadoop, Spark, Kubernetes |
| **분산 트랜잭션 관리** | Two-Phase Commit, SAGA 패턴 |

---

## 6️⃣ 분산 시스템 설계시 고려 요소

| 요소               | 설명 |
|--------------------|------|
| **데이터 복제 전략**   | 여러 노드에 데이터 복사 → 내구성/가용성 증가 |
| **샤딩 (Sharding)**    | 데이터를 키 or 범위 기준으로 분할 저장 |
| **리더-팔로워 구조**   | 리더 노드가 쓰기 처리, 팔로워는 읽기 (ex. Redis Sentinel) |
| **장애 복구 메커니즘** | Failover, Heartbeat, 헬스 체크 |
| **모니터링 및 알림**   | Prometheus, Grafana 등으로 실시간 상태 추적 |
| **보안**              | 노드 간 인증, 암호화, 접근 제어

---

## 🎯 정리 요약

✔ **분산 시스템은 여러 대의 노드가 협력하여 하나의 시스템처럼 동작**  
✔ 확장성, 장애 허용성, 고가용성 등의 이점을 제공하지만, 복잡성 증가와 일관성 문제가 존재  
✔ **CAP 이론**, **데이터 복제/샤딩**, **장애 복구** 등 핵심 개념 이해가 필수  
✔ 현대 대규모 서비스는 대부분 분산 시스템 기반으로 설계됨
