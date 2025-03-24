# 📄 Apache Cassandra 개요 정리

---

## 1️⃣ Cassandra란?

**Apache Cassandra**는 대용량 데이터를 처리하기 위한 **분산형 NoSQL 데이터베이스**입니다.  

- **페이스북**에서 개발되어, 현재는 **Apache 재단**에서 관리
- **수평 확장성**, **고가용성**, **무중단 아키텍처**에 강함
- 데이터는 여러 노드에 **분산 저장**, 장애 허용성과 확장성 우수
- **정형화된 스키마는 있지만**, RDB처럼 조인이나 트랜잭션은 제한적

---

## 2️⃣ 주요 특징

| 특징 | 설명 |
|------|------|
| **분산 구조 (Distributed)** | 데이터를 여러 노드에 자동 분산 저장 |
| **마스터리스 구조** | 모든 노드가 평등한 역할 (Single Point of Failure 없음) |
| **높은 쓰기 성능** | 빠른 쓰기 처리 (Log-structured Storage) |
| **Column-family 기반 구조** | RDB의 테이블과 유사하지만 유연한 구조 |
| **내결함성(Fault-tolerant)** | 노드 장애 시 자동 복구 및 리플리케이션 |
| **선형 확장성** | 노드를 추가할수록 처리량이 선형적으로 증가 |
| **튜닝 가능한 일관성** | Consistency Level을 쿼리마다 설정 가능 |

---

## 3️⃣ 데이터 구조

Cassandra는 **Keyspace → Table(Column Family) → Row → Column** 구조  

```
Keyspace
 └── Table (Column Family)
      └── Row (Primary Key 기반)
           └── Columns (Key-Value 쌍)
```

- 데이터는 내부적으로 **SSTable, MemTable, Commit Log**로 구성되어 있음

---

## 4️⃣ CAP 이론에서의 Cassandra

#### AP 기반 시스템
- ✅ **Availability (가용성)**  
- ✅ **Partition Tolerance (분할 허용성)**  
- ❌ **Consistency (일관성)** → 쿼리 시 설정 가능 (Tunability)

---

## 5️⃣ 기본 명령어 예시 (CQL - Cassandra Query Language)

```sql
-- 키 스페이스 생성
CREATE KEYSPACE mykeyspace WITH replication = {
  'class': 'SimpleStrategy',
  'replication_factor': 3
};

-- 테이블 생성
CREATE TABLE users (
  id UUID PRIMARY KEY,
  name TEXT,
  age INT
);

-- 데이터 삽입
INSERT INTO users (id, name, age) VALUES (uuid(), 'Alice', 25);

-- 데이터 조회
SELECT * FROM users WHERE id = ...;
```

---

## 6️⃣ Cassandra 주요 활용 사례

- SNS, 채팅, 피드 시스템
- IoT 센서 데이터 저장
- 로그 및 이벤트 데이터 저장
- 사용자 활동 기록, 시간 기반 데이터 처리
- 대규모 전자상거래 추천 및 개인화 시스템

---

## 🎯 정리

✔ 대규모 분산 시스템에 최적화된 NoSQL DB  
✔ 높은 쓰기 처리량 + 장애 허용성 + 유연한 확장  
✔ 마스터리스 구조로 무중단 서비스 가능  
✔ 단점: 조인 미지원, 복잡한 쿼리에는 부적합  

