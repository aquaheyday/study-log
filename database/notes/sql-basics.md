# 🧾 SQL 기본 문법

SQL(Structured Query Language)은 **데이터베이스와 상호작용하기 위한 표준 언어**입니다.  
데이터를 **생성(Create)**, **조회(Read)**, **수정(Update)**, **삭제(Delete)**하는 데 사용됩니다.

---

## 1️⃣ SQL 명령 분류

| 유형 | 설명 | 예시 |
|------|------|------|
| DDL (정의) | 테이블/스키마 정의 | `CREATE`, `ALTER`, `DROP` |
| DML (조작) | 데이터 조작 | `SELECT`, `INSERT`, `UPDATE`, `DELETE` |
| DCL (제어) | 권한 부여 | `GRANT`, `REVOKE` |
| TCL (트랜잭션) | 트랜잭션 제어 | `COMMIT`, `ROLLBACK`, `SAVEPOINT` |

---

## 2️⃣ 테이블 생성 (CREATE TABLE)

```sql
CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE,
    age INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

✔ 주요 제약 조건  
- `PRIMARY KEY`: 기본 키  
- `NOT NULL`: NULL 불가  
- `UNIQUE`: 중복 불가  
- `DEFAULT`: 기본값  

---

## 3️⃣ 데이터 삽입 (INSERT)

```sql
INSERT INTO users (name, email, age)
VALUES ('Alice', 'alice@example.com', 30);
```

✔ 여러 건 한 번에 삽입:

```sql
INSERT INTO users (name, email) VALUES
('Bob', 'bob@example.com'),
('Charlie', 'charlie@example.com');
```

---

## 4️⃣ 데이터 조회 (SELECT)

```sql
SELECT name, age FROM users;
```

### 주요 옵션들:

| 기능 | 문법 예시 |
|------|-----------|
| 전체 컬럼 | `SELECT * FROM users` |
| 조건 | `WHERE age > 20` |
| 정렬 | `ORDER BY age DESC` |
| 제한 | `LIMIT 10` |
| 별칭 | `SELECT name AS username FROM users` |

```sql
SELECT name, age FROM users
WHERE age >= 18
ORDER BY age DESC
LIMIT 5;
```

---

## 5️⃣ 데이터 수정 (UPDATE)

```sql
UPDATE users
SET age = 31
WHERE name = 'Alice';
```

✔ `WHERE` 없이 쓰면 **전체 데이터 변경됨** → 주의

---

## 6️⃣ 데이터 삭제 (DELETE)

```sql
DELETE FROM users WHERE age < 18;
```

✔ 테이블 전체 삭제 시:

```sql
DELETE FROM users;
```

✔ **조심!** `WHERE` 없이 쓰면 전체 삭제됨

---

## 7️⃣ 조건절 연산자

| 연산자 | 의미 |
|--------|------|
| `=`    | 같다 |
| `!=` 또는 `<>` | 다르다 |
| `>` `<` `>=` `<=` | 비교 |
| `BETWEEN a AND b` | 범위 조건 |
| `IN (a, b, c)` | 여러 값 중 포함 |
| `LIKE '%abc%'` | 패턴 검색 (문자 포함) |
| `IS NULL`, `IS NOT NULL` | NULL 검사 |

---

## 8️⃣ 그룹화 및 집계 (GROUP BY)

```sql
SELECT age, COUNT(*) AS count
FROM users
GROUP BY age
HAVING count > 1;
```

✔ `GROUP BY`: 특정 컬럼으로 그룹화  
✔ `HAVING`: 그룹화된 결과에 조건  

### 주요 집계 함수

| 함수 | 설명 |
|------|------|
| `COUNT()` | 개수 |
| `SUM()` | 합계 |
| `AVG()` | 평균 |
| `MIN()` / `MAX()` | 최소/최대 |

---

## 9️⃣ 테이블 수정 (ALTER)

```sql
ALTER TABLE users ADD COLUMN phone VARCHAR(20);
ALTER TABLE users DROP COLUMN age;
ALTER TABLE users RENAME TO members;
```

✔ 테이블 구조 변경 시 사용

---

## 🔟 테이블 삭제 (DROP)

```sql
DROP TABLE users;
```

✔ 삭제된 테이블은 복구 불가 → 신중하게 사용

---

## 🔁 조인 (JOIN)

### 1) INNER JOIN (교집합)

```sql
SELECT users.name, orders.item
FROM users
JOIN orders ON users.id = orders.user_id;
```

---

### 2) LEFT JOIN (users는 모두, orders는 있으면)

```sql
SELECT users.name, orders.item
FROM users
LEFT JOIN orders ON users.id = orders.user_id;
```

✔ 실무에서 많이 쓰이는 JOIN 문법  
✔ `ON` 조건을 기준으로 두 테이블을 연결

---

## 🎯 정리

| 작업 | 키워드 |
|------|--------|
| 테이블 생성 | `CREATE TABLE` |
| 데이터 삽입 | `INSERT INTO` |
| 조회 | `SELECT`, `WHERE`, `ORDER BY`, `LIMIT` |
| 수정/삭제 | `UPDATE`, `DELETE` |
| 조건 비교 | `=`, `LIKE`, `IN`, `BETWEEN` |
| 집계 | `GROUP BY`, `COUNT`, `AVG` 등 |
| 테이블 변경/삭제 | `ALTER`, `DROP` |
| 테이블 연결 | `JOIN`, `ON` |
