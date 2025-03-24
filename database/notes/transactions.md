# 🔐 데이터베이스 트랜잭션 & ACID

**트랜잭션(Transaction)** 은 데이터베이스에서 하나 이상의 작업을 **하나의 논리적인 작업 단위로 묶는 것**입니다.  
⚠️ 반드시 **모두 성공하거나, 전부 실패해야 함**.

---

## 1️⃣ 트랜잭션(Transaction)이란?

데이터베이스의 **일련의 작업을 하나의 단위로 처리하는 것**, 중간에 오류가 발생하면 전체 작업을 취소(rollback)하여 **일관성 유지**

### 은행 계좌 이체 예시
#### 1. A 계좌에서 1000원 출금  
#### 2. B 계좌에 1000원 입금  
⚠️ 둘 중 하나라도 실패하면 전체 취소되어야 안전함

---

## 2️⃣ 트랜잭션의 4가지 속성 (ACID)

| 속성 | 설명 |
|------|------|
| **A - Atomicity (원자성)** | 모두 성공하거나 모두 실패해야 함 |
| **C - Consistency (일관성)** | 트랜잭션 전후 데이터의 무결성 보장 |
| **I - Isolation (격리성)** | 동시에 실행되는 트랜잭션 간 간섭 없음 |
| **D - Durability (지속성)** | 트랜잭션 성공 시 데이터는 영구 저장됨 |

#### 💡 ACID 쉽게 이해하기
- **Atomicity**: 다 아니면 0 (All or Nothing)  
- **Consistency**: 무결성 규칙 깨지면 안 됨  
- **Isolation**: 동시에 실행되어도 결과는 순차 실행처럼  
- **Durability**: 커밋된 데이터는 시스템 꺼져도 보존됨

---

## 3️⃣ 트랜잭션 제어문 (SQL 문법)

```sql
BEGIN;         -- 트랜잭션 시작
UPDATE accounts SET balance = balance - 1000 WHERE id = 1;
UPDATE accounts SET balance = balance + 1000 WHERE id = 2;
COMMIT;        -- 모두 성공 시 반영
-- ROLLBACK;   -- 실패 시 전체 취소
```

✔ 일반적으로 **BEGIN ~ COMMIT** 또는 **ROLLBACK** 형태로 사용  
✔ 일부 DBMS에서는 `START TRANSACTION`도 사용 가능

---

## 4️⃣ Isolation Level (격리 수준)

#### 동시성 제어를 위한 트랜잭션의 격리 레벨

| 격리 수준 | 설명 | 현상 예시 |
|-----------|------|-----------|
| **READ UNCOMMITTED** | 커밋 전 데이터도 읽음 | 더러운 읽기(Dirty Read) 발생 가능 |
| **READ COMMITTED** | 커밋된 데이터만 읽음 | 반복 불가능한 읽기(PNR) 발생 가능 |
| **REPEATABLE READ** | 읽은 데이터는 트랜잭션 중 항상 동일 | 팬텀 리드 발생 가능 |
| **SERIALIZABLE** | 가장 엄격, 완전한 순차 실행처럼 처리 | 모든 현상 방지 (성능 낮음) |

---

## 5️⃣ 트랜잭션이 필요한 대표 사례

- 계좌 이체, 포인트 충전/사용
- 장바구니 → 주문 전환
- 재고 차감 처리
- 복수 테이블/로우 변경이 필요한 경우

---

## 6️⃣ 실무 팁 & 주의사항

- 트랜잭션 범위는 **짧고 명확하게** 유지  
- 꼭 필요한 경우에만 사용 (락 걸리면 성능 저하 우려)  
- 예외 발생 시 `ROLLBACK` 처리 필수  
- ORM 사용 시 자동 트랜잭션 지원 여부 확인  
- **커밋 or 롤백 누락** 없도록 예외 처리 잘 구성

---

## 7️⃣ Go 예시 (SQL 트랜잭션 처리)

```go
tx, err := db.Begin()
if err != nil {
    log.Fatal(err)
}

_, err1 := tx.Exec("UPDATE accounts SET balance = balance - 1000 WHERE id = 1")
_, err2 := tx.Exec("UPDATE accounts SET balance = balance + 1000 WHERE id = 2")

if err1 != nil || err2 != nil {
    tx.Rollback()
    log.Fatal("Transaction failed")
} else {
    tx.Commit()
}
```

---

## 🎯 정리

| 항목 | 내용 |
|------|------|
| 트랜잭션 | 여러 SQL 작업을 하나의 단위로 묶는 기능 |
| ACID | 원자성, 일관성, 격리성, 지속성 |
| 주요문법 | `BEGIN`, `COMMIT`, `ROLLBACK` |
| 격리 수준 | READ UNCOMMITTED → SERIALIZABLE |
| 사용예시 | 계좌이체, 재고관리, 주문처리 등 |
