# 📄 MongoDB 개요

---

## 1️⃣ MongoDB란?

**MongoDB**는 문서(Document) 기반의 **NoSQL 데이터베이스**입니다.  
JSON과 유사한 **BSON(Binary JSON)** 형식으로 데이터를 저장합니다.

- **스키마가 유연**하여 다양한 형태의 데이터를 저장 가능  
- 수직 확장보다 **수평 확장(샤딩)**에 적합  
- 복잡한 조인보다는 **빠른 읽기/쓰기, 확장성, 유연성**에 초점  
- MongoDB Inc.에서 개발, **무료 + 유료 클라우드 서비스(MongoDB Atlas)** 제공

---

## 2️⃣ 기본 용어

| 용어 | 설명 |
|------|------|
| **Database** | 데이터베이스 단위 (SQL의 DB와 유사) |
| **Collection** | 문서들의 집합 (SQL의 Table과 유사) |
| **Document** | 실제 데이터 단위, JSON-like 구조 (SQL의 Row와 유사) |
| **Field** | 키-값 쌍으로 구성된 데이터 항목 (SQL의 Column과 유사) |
| **BSON** | Binary JSON, MongoDB 내부 데이터 저장 형식 |

---

## 3️⃣ 특징

| 항목 | 설명 |
|------|------|
| 비정형 데이터 저장 | 스키마가 고정되지 않아 다양한 형식의 데이터 저장 가능 |
| 고성능 | 인메모리 캐시와 비동기식 쓰기 처리로 빠른 성능 |
| 복제(Replica Set) | 장애 대비를 위한 자동 복제 기능 |
| 샤딩(Sharding) | 수평 확장을 위한 데이터 분산 저장 |
| 강력한 쿼리 | JSON 기반 쿼리, Aggregation Framework, 인덱싱 등 |
| 클라우드 | MongoDB Atlas로 클라우드 기반 서비스 제공 |

---

## 4️⃣ 주요 명령어

### 1) 데이터베이스 관련
```js
show dbs               // 모든 DB 목록 확인
use mydb               // mydb 데이터베이스로 전환 (없으면 생성)
db.dropDatabase()      // 현재 DB 삭제
```

### 2) 컬렉션 관련
```js
db.createCollection("users")   // 컬렉션 생성
show collections               // 컬렉션 목록 보기
db.users.drop()                // 컬렉션 삭제
```

### 3) 데이터 조작
```js
// 문서 삽입
db.users.insertOne({ name: "Alice", age: 25 })

// 문서 조회
db.users.find({ age: { $gt: 20 } })

// 문서 수정
db.users.updateOne({ name: "Alice" }, { $set: { age: 26 } })

// 문서 삭제
db.users.deleteOne({ name: "Alice" })
```

---

## 5️⃣ 인덱스
인덱스는 검색 성능 향상을 위해 사용됨  

```js
db.users.createIndex({ name: 1 })   // name 필드에 오름차순 인덱스 생성
```

---

## 6️⃣ Aggregation (집계)

MongoDB의 강력한 데이터 처리 파이프라인  

```js
db.orders.aggregate([
  { $match: { status: "delivered" } },
  { $group: { _id: "$customer", total: { $sum: "$amount" } } }
])
```

---

## 🎯 정리

✔ 문서 기반 NoSQL 데이터베이스  
✔ 스키마가 유연하고, 수평 확장에 강함  
✔ JSON 스타일 쿼리, 집계, 인덱스 등 풍부한 기능  
✔ 캐시, 세션 저장소, 로그, 대용량 데이터 처리에 적합  
✔ MongoDB Atlas로 클라우드 환경도 쉽게 구축 가능

