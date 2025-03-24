# 📄 PostgreSQL 성능 최적화 정리

---

## 1️⃣ PostgreSQL 성능 최적화란?

PostgreSQL에서 **데이터 처리 속도와 효율성을 향상**시키기 위해 **쿼리, 인덱스, 설정, 하드웨어 자원 등 다양한 요소를 최적화**하는 작업

---

## 2️⃣ 주요 튜닝 영역

| 영역 | 설명 |
|------|------|
| **쿼리 최적화** | 느린 SQL 문 분석 및 개선 |
| **인덱스 관리** | 인덱스 구조 최적화 |
| **서버 설정 변경** | PostgreSQL 설정값 튜닝 (`postgresql.conf`) |
| **자원 활용** | CPU, 메모리, 디스크 효율적 사용 |
| **통계/캐시 활용** | 실행계획 정확도 향상, 캐시 적중률 증가 |

---

## 3️⃣ 쿼리 튜닝

#### 1. 실행 계획 확인

```sql
EXPLAIN ANALYZE SELECT * FROM users WHERE email = 'test@example.com';
```

✔ 실제 수행 시간, 행 수, 인덱스 사용 여부 등을 확인  
✔ `Seq Scan` → `Index Scan` 으로 전환할 수 있는지 분석  

#### 2. 쿼리 튜닝 팁

| 항목 | 설명 |
|------|------|
| SELECT * 지양 | 필요한 컬럼만 선택 |
| 조건절 인덱스 | WHERE절에 사용되는 컬럼에 인덱스 설정 |
| EXISTS 사용 | `IN`보다 `EXISTS`가 더 효율적인 경우 많음 |
| 서브쿼리 → JOIN | 복잡한 서브쿼리 대신 JOIN 구조로 변경 |

---

## 4️⃣ 인덱스 최적화

| 인덱스 종류 | 설명 |
|-------------|------|
| **B-Tree** | 기본 인덱스, 범위 조회에 적합 |
| **Hash** | 정확한 일치 검색에 빠름 |
| **GIN** | 배열, JSON, Full-Text Search에 사용 |
| **Partial Index** | 조건부 인덱스 (자주 조회되는 조건만 인덱싱) |
| **Covering Index** | 필요한 컬럼을 모두 포함하는 인덱스 (Index Only Scan 유도) |

```sql
-- 예: JSONB 필드 GIN 인덱스 생성
CREATE INDEX idx_info ON items USING GIN (data jsonb_path_ops);
```

---

## 5️⃣ 설정 파일 튜닝 (`postgresql.conf`)

| 파라미터 | 설명 |
|-----------|------|
| `shared_buffers` | DB 메모리 캐시 (전체 메모리의 25~40%) |
| `work_mem` | 정렬, 해시 연산 시 사용하는 메모리 (쿼리당 적용) |
| `maintenance_work_mem` | VACUUM, CREATE INDEX 등에 사용되는 메모리 |
| `effective_cache_size` | OS 파일시스템 캐시 추정치 (읽기 성능 참고용) |
| `wal_buffers` | Write-Ahead Log 버퍼 크기 |
| `max_connections` | 동시 접속자 수 제한 |
| `autovacuum` | 자동 VACUUM 활성화 여부 |

---

## 6️⃣ Autovacuum & Analyze

- PostgreSQL은 **MVCC 방식** 사용 → 쓰기 많으면 **데드 튜플** 쌓임  
- `VACUUM`, `ANALYZE` 주기적 수행 필수!

```bash
-- 수동으로 실행
VACUUM ANALYZE;
```

| 설정 항목 | 설명 |
|-----------|------|
| `autovacuum_vacuum_threshold` | 최소 튜플 수 변화 후 VACUUM 시작 |
| `autovacuum_vacuum_scale_factor` | 전체 튜플 대비 변화 비율 |

---

## 7️⃣ 모니터링 방법

| 도구 | 설명 |
|------|------|
| `pg_stat_activity` | 현재 실행 중인 쿼리 확인 |
| `pg_stat_user_tables` | 테이블별 튜플 통계 |
| `pg_stat_statements` | 쿼리별 수행 횟수, 시간 등 분석 |
| `auto_explain` | 자동 쿼리 실행 계획 로그 남기기 |
| 외부 도구 | pgAdmin, pganalyze, Prometheus + Grafana 등 |

---

## 8️⃣ 실전 튜닝 팁

- 느린 쿼리 → `EXPLAIN ANALYZE`로 구조 파악  
- 캐시 잘 쓰게끔 → Index Only Scan 유도  
- Autovacuum 주기 조정 → 데드 튜플 줄이기  
- 통계 정보 최신화 → `ANALYZE` 정기 수행  
- 쿼리 수 많을 땐 `work_mem` 점검 (낮으면 디스크 사용됨)

---

## 🎯 정리

✔ PostgreSQL 성능 튜닝은 **쿼리, 인덱스, 설정, 자원 활용 최적화**의 조화  
✔ `EXPLAIN`, `pg_stat_statements`, `autovacuum` 활용은 필수  
✔ 튜닝은 수치 기반으로 분석 → 변경 → 검증의 반복  
✔ 초과 튜닝보다 **적절한 튜닝과 모니터링**이 장기적인 성능 유지에 중요

