# 📄 MySQL 성능 최적화 정리

---

## 1️⃣ MySQL 성능 최적화란?

MySQL에서 **데이터 처리 속도를 높이고**, **시스템 자원 사용을 효율화하며**, **안정적인 서비스 제공을 위한 설정 및 구조 개선 작업**입니다.

---

## 2️⃣ 최적화 대상 영역

| 영역 | 설명 |
|------|------|
| **쿼리(Query)** | 느린 SQL 개선, 실행 계획 분석 |
| **데이터 구조** | 인덱스 구성, 정규화/반정규화 |
| **서버 설정** | MySQL 설정값 튜닝 |
| **하드웨어 자원** | CPU, RAM, SSD 등 인프라 성능 |
| **애플리케이션 로직** | 반복 쿼리, 불필요한 요청 최소화 |

---

## 3️⃣ 쿼리 최적화

#### 1. 실행 계획 분석 (EXPLAIN)
```sql
EXPLAIN SELECT * FROM users WHERE email = 'test@example.com';
```
- `type`, `key`, `rows`, `Extra` 항목 분석  
- `ALL`(Full Scan)은 지양, `ref`, `const`, `eq_ref`가 이상적

#### 2. 슬로우 쿼리 로그 확인
```sql
SET GLOBAL slow_query_log = 'ON';
SET GLOBAL long_query_time = 1;
```

#### 3. 자주 쓰는 쿼리 패턴 개선
- SELECT * → 필요한 컬럼만  
- 서브쿼리 → JOIN 또는 EXISTS로 대체  
- LIKE '%word%' → Full Scan 주의

---

## 4️⃣ 인덱스 최적화

| 전략 | 설명 |
|------|------|
| 필요한 컬럼에 인덱스 추가 | WHERE, JOIN, ORDER BY에 사용되는 컬럼 |
| 복합 인덱스 활용 | 자주 함께 조회되는 컬럼 조합 |
| 과도한 인덱스 피하기 | INSERT/UPDATE 성능 저하 가능 |
| 인덱스 통계 갱신 | `ANALYZE TABLE` 주기적 수행 |

---

## 5️⃣ 테이블 구조 최적화

| 전략 | 설명 |
|------|------|
| 정규화 | 중복 최소화, 무결성 유지 |
| 반정규화 | 조인 최소화로 조회 성능 향상 |
| 데이터 타입 최소화 | VARCHAR vs CHAR, INT vs TINYINT 등 고려 |
| 파티셔닝 | 대용량 테이블 분할 (Range, List, Hash 등) |

---

## 6️⃣ 설정 파일 튜닝 (`my.cnf`)

| 파라미터 | 설명 |
|----------|------|
| `innodb_buffer_pool_size` | InnoDB 캐시 크기 (전체 메모리의 60~80%) |
| `query_cache_size` | SELECT 캐시 (MySQL 5.x 기준, 8.0에서는 제거됨) |
| `max_connections` | 동시 접속 허용 수 |
| `tmp_table_size`, `max_heap_table_size` | 임시 테이블 메모리 크기 설정 |
| `slow_query_log` | 느린 쿼리 로그 활성화 |

---

## 7️⃣ 기타 성능 개선 전략

- **Connection Pooling** → 매 요청마다 연결/해제 방지
- **배치 작업 스케줄링** → 트래픽 낮은 시간대에 처리
- **주기적 최적화** → `OPTIMIZE TABLE`로 조각화 방지
- **모니터링 도구 활용** → PMA, Percona Toolkit, NewRelic 등

---

## 8️⃣ 모니터링/분석 도구

| 도구 | 설명 |
|------|------|
| **MySQL Workbench** | 쿼리 시각화 및 성능 분석 |
| **Percona Toolkit** | 다양한 MySQL 진단 도구 |
| **slow_query_log** | 느린 쿼리 기록 분석 |
| **INFORMATION_SCHEMA** | 실시간 테이블/인덱스/트랜잭션 정보 조회 |

---

## 🎯 정리

✔ MySQL 성능 개선은 **쿼리 + 인덱스 + 설정 + 구조 + 자원**의 조화  
✔ 병목을 찾고, 데이터 접근 경로를 줄이는 것이 핵심  
✔ 튜닝 전후 비교를 통해 **수치 기반으로 개선 확인**  
✔ 무조건 빠르게가 아니라, **안정성과 확장성까지 고려한 최적화**가 중요

