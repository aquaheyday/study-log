# 🛢️ Database

이 폴더는 **데이터베이스 학습 과정에서 정리한 개념 노트**를 보관하는 공간입니다.  
각 노트는 **개념 정의, 기본 문법, 실습 예제, 최적화 팁, 관련 자료** 순으로 구성되어 체계적인 학습 자료로 활용할 수 있습니다.

---

## 📋 개념 정리 목록  

### 📌 기본 개념  
| 번호 | 주제 | 파일명 | 설명 |  
|---|---|---|---|  
| 01 | 데이터베이스 개요 | [database-intro.md](./database-intro.md) | 데이터베이스란 무엇인가? RDBMS vs NoSQL 비교 |  
| 02 | SQL 기본 문법 | [sql-basics.md](./sql-basics.md) | SELECT, INSERT, UPDATE, DELETE 문법 |  
| 03 | 관계형 데이터베이스 | [rdbms.md](./rdbms.md) | 관계형 데이터베이스의 기본 개념과 특징 |  
| 04 | NoSQL 개요 | [nosql-intro.md](./nosql-intro.md) | NoSQL의 개념 및 종류 (Key-Value, Document, Column, Graph) |  
| 05 | 데이터 모델링 | [data-modeling.md](./data-modeling.md) | ERD(Entity Relationship Diagram) 작성 방법 |  

### 🔲 SQL 문법 및 고급 쿼리  
| 번호 | 주제 | 파일명 | 설명 |  
|---|---|---|---|  
| 06 | 조인(Join) | [sql-join.md](./sql-join.md) | INNER JOIN, LEFT JOIN, RIGHT JOIN, FULL JOIN |  
| 07 | 서브쿼리 | [subquery.md](./subquery.md) | 서브쿼리 개념과 활용법 |  
| 08 | 윈도우 함수 | [window-functions.md](./window-functions.md) | RANK(), DENSE_RANK(), LEAD(), LAG() |  
| 09 | 트랜잭션 & ACID | [transactions.md](./transactions.md) | 트랜잭션 개념 및 COMMIT, ROLLBACK 사용법 |  
| 10 | 인덱스와 성능 최적화 | [index-optimization.md](./index-optimization.md) | 인덱스 개념 및 쿼리 최적화 방법 |  

### 🔄 NoSQL  
| 번호 | 주제 | 파일명 | 설명 |  
|---|---|---|---|  
| 11 | Redis 개요 | [redis.md](./redis.md) | Key-Value 저장소, 캐싱 전략 |  
| 12 | MongoDB 기초 | [mongodb-basics.md](./mongodb-basics.md) | 문서형 데이터베이스 기본 개념 및 CRUD |  
| 13 | Cassandra 개요 | [cassandra.md](./cassandra.md) | 분산형 컬럼 기반 데이터베이스 |  
| 14 | Graph DB 개념 | [graph-db.md](./graph-db.md) | Graph 데이터베이스 (Neo4j) 개요 및 활용 |  

### 🌍 아키텍처  
| 번호 | 주제 | 파일명 | 설명 |  
|---|---|---|---|  
| 15 | 데이터 정규화 | [normalization.md](./normalization.md) | 1NF, 2NF, 3NF, BCNF 개념 |  
| 16 | 샤딩과 레플리케이션 | [sharding-replication.md](./sharding-replication.md) | 데이터베이스 수평/수직 분할, Master-Slave 복제 |  
| 17 | 데이터베이스 백업 & 복구 | [backup-restore.md](./backup-restore.md) | MySQL, PostgreSQL, MongoDB 백업 및 복구 전략 |  
| 18 | 데이터베이스 보안 | [db-security.md](./db-security.md) | SQL Injection 방지, 권한 관리, 암호화 |  

### 🚀 고급 개념  
| 번호 | 주제 | 파일명 | 설명 |  
|---|---|---|---|  
| 19 | 분산 데이터베이스 | [distributed-db.md](./distributed-db.md) | CAP 이론, Eventual Consistency 개념 |  
| 20 | 데이터 웨어하우스 | [data-warehouse.md](./data-warehouse.md) | OLTP vs OLAP, 데이터 마트 개념 |  
| 21 | BigQuery & Snowflake | [bigquery-snowflake.md](./bigquery-snowflake.md) | 빅데이터 분석을 위한 데이터 웨어하우스 |  

### 🛠️ 성능 테스트 및 운영  
| 번호 | 주제 | 파일명 | 설명 |  
|---|---|---|---|  
| 22 | 데이터베이스 성능 테스트 | [db-performance-testing.md](./db-performance-testing.md) | `EXPLAIN ANALYZE`, 벤치마킹 도구 사용법 |  
| 23 | MySQL 성능 최적화 | [mysql-performance.md](./mysql-performance.md) | 쿼리 튜닝, 인덱스 최적화 |  
| 24 | PostgreSQL 성능 튜닝 | [postgres-performance.md](./postgres-performance.md) | `pg_stat_statements`, 쿼리 캐싱 |  
| 25 | 데이터베이스 모니터링 | [db-monitoring.md](./db-monitoring.md) | Prometheus, Grafana를 활용한 DB 모니터링 |  

---

## 📚 참고 자료 & 추천 도서  
- [SQL 첫걸음](https://www.yes24.com/Product/Goods/66325951) - 초보자를 위한 데이터베이스 입문
- [Designing Data](https://www.oreilly.com/library/view/designing-data-intensive-applications/9781491903063/) - Intensive Applications (데이터 중심 애플리케이션 설계)
- [PostgreSQL 공식 문서](https://www.postgresql.org/docs/)  
- [MongoDB 개발자 가이드](https://www.mongodb.com/docs/guides/)  
