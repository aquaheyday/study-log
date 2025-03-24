# 📄 데이터베이스 모니터링 (Database Monitoring) 정리

---

## 1️⃣ 데이터베이스 모니터링이란?

데이터베이스의 **상태, 성능, 자원 사용량 등을 실시간으로 관찰하고 분석**하여, 장애 예방과 성능 최적화를 지원하는 활동입니다.

---

## 2️⃣ 모니터링 목적

| 목적 | 설명 |
|------|------|
| 성능 유지 | 응답 속도 저하, 병목 현상 조기 발견 |
| 장애 예방 | 과부하, 디스크 부족, 연결 폭주 등 사전 감지 |
| 자원 최적화 | CPU, 메모리, I/O 사용량 분석 |
| 용량 계획 | 데이터 증가 추세 분석 → 스케일링 시점 예측 |
| 원인 분석 | 장애 발생 시 신속한 트러블슈팅 가능 |

---

## 3️⃣ 주요 모니터링 지표

| 지표 | 설명 |
|------|------|
| **쿼리 응답 시간** | SELECT, INSERT 등의 평균/최대 응답 시간 |
| **TPS/QPS** | 초당 트랜잭션/쿼리 수 |
| **CPU/Memory 사용량** | DB 프로세스의 시스템 자원 사용 현황 |
| **디스크 I/O** | 읽기/쓰기 속도, IOPS |
| **접속자 수 / 연결 수** | 동시에 접속 중인 세션 수 |
| **슬로우 쿼리 수** | 느리게 수행된 쿼리 수 및 목록 |
| **Lock 대기/경합** | 테이블/행 잠금 충돌 현황 |
| **Buffer Cache Hit Ratio** | 디스크 대신 캐시에서 읽은 비율 (높을수록 좋음) |
| **Autovacuum 상태** (PostgreSQL) | 정리 작업 수행 현황 |

---

## 4️⃣ 모니터링 도구

| 도구 | 대상 | 특징 |
|------|------|------|
| **MySQL Workbench** | MySQL | 실시간 성능, 세션 모니터링 |
| **pgAdmin / pganalyze** | PostgreSQL | 실행 쿼리, Autovacuum, 인덱스 등 분석 |
| **Percona Monitoring** | MySQL | 오픈소스 모니터링 툴 |
| **Prometheus + Grafana** | 모든 DB (커스텀 가능) | 시각화 대시보드, 알림 설정 |
| **Datadog / NewRelic** | 클라우드 기반 통합 모니터링 | 다양한 DB 연동 지원 |
| **CloudWatch / Stackdriver** | AWS / GCP 전용 | 클라우드 DB (RDS, BigQuery 등) 모니터링 지원 |

---

## 5️⃣ 실시간 모니터링 SQL 예시 (MySQL 기준)

```sql
-- 현재 접속된 세션 보기
SHOW PROCESSLIST;

-- 슬로우 쿼리 개수 확인
SHOW GLOBAL STATUS LIKE 'Slow_queries';

-- 캐시 히트율 확인
SHOW STATUS LIKE 'Qcache%';

-- 현재 락 대기 쿼리
SELECT * FROM information_schema.innodb_locks;
```

---

## 6️⃣ PostgreSQL 전용 뷰 예시

```sql
-- 현재 실행 중인 쿼리
SELECT * FROM pg_stat_activity;

-- 쿼리별 통계 정보
SELECT * FROM pg_stat_statements ORDER BY total_time DESC;

-- 자동 vacuum 상태
SELECT * FROM pg_stat_user_tables;
```

---

## 7️⃣ 모니터링 잘하는 팁 💡

- [ ] **슬로우 쿼리 로그 활성화**
- [ ] **리소스 임계값 설정 → 알림 시스템 연동**
- [ ] **지표 수집 + 시각화 → 추세 분석**
- [ ] **장애 발생 전후 로그 자동 저장**
- [ ] **모니터링 도구 + SQL 병행 사용**

---

## 🎯 정리

✔ DB 모니터링은 **문제 예방 + 성능 유지**의 핵심  
✔ 쿼리/리소스/세션/IO 등 다양한 지표 관찰 필요  
✔ SQL + 전용 도구 + 시각화 툴 조합이 이상적  
✔ **모니터링은 튜닝보다 먼저**, 평소부터 실시간 지표 확인이 중요

