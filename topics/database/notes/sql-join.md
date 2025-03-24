# 🔗 데이터베이스 조인(Join) 정리

**조인(Join)** 은 여러 테이블의 데이터를 **공통된 컬럼을 기준으로 연결하여 조회**할 때 사용됩니다.  
관계형 데이터베이스의 핵심 기능으로, 테이블 간 관계를 통해 **복합적인 데이터 조회**를 가능하게 합니다.

---

## 1️⃣ 조인이란?

두 개 이상의 테이블을 **공통된 키(주로 PK ↔ FK)** 를 기준으로 연결하여 결과를 반환하는 SQL 기능

#### 예:
- `users` 테이블과 `orders` 테이블을 `user_id` 로 연결하여 사용자 정보 + 주문 정보를 한 번에 조회

---

## 2️⃣ 조인 종류 요약

| 종류 | 설명 |
|------|------|
| **INNER JOIN** | 양쪽 모두 일치하는 데이터만 |
| **LEFT JOIN** | 왼쪽 테이블은 모두, 오른쪽은 일치하는 것만 |
| **RIGHT JOIN** | 오른쪽 테이블은 모두, 왼쪽은 일치하는 것만 |
| **FULL OUTER JOIN** | 양쪽 모두 포함, 없으면 NULL |
| **SELF JOIN** | 같은 테이블을 자기 자신과 조인 |
| **CROSS JOIN** | 모든 조합 반환 (카테시안 곱) |

---

## 3️⃣ INNER JOIN

```sql
SELECT u.name, o.item
FROM users u
INNER JOIN orders o ON u.id = o.user_id;
```

✔ 양쪽 모두 매칭되는 행만 조회  
✔ 실무에서 가장 많이 사용됨

---

## 4️⃣ LEFT JOIN (LEFT OUTER JOIN)

```sql
SELECT u.name, o.item
FROM users u
LEFT JOIN orders o ON u.id = o.user_id;
```

✔ `users`에 있는 모든 사용자 + `orders`가 있는 경우만 매칭  
✔ 주문 내역 없는 사용자도 포함 (NULL 반환됨)

---

## 5️⃣ RIGHT JOIN (RIGHT OUTER JOIN)

```sql
SELECT u.name, o.item
FROM users u
RIGHT JOIN orders o ON u.id = o.user_id;
```

✔ `orders` 테이블에 있는 모든 주문 + 해당 사용자 정보  
✔ 사용자 정보 없으면 `users` 컬럼은 NULL

---

## 6️⃣ FULL OUTER JOIN (일부 DB 지원)

```sql
SELECT u.name, o.item
FROM users u
FULL OUTER JOIN orders o ON u.id = o.user_id;
```

✔ `users`, `orders` 모두 포함  
✔ 어느 한 쪽에만 있어도 결과에 포함됨  
⚠️ MySQL은 직접 지원하지 않음 → `UNION`으로 구현  

---

## 7️⃣ SELF JOIN

```sql
SELECT a.name AS employee, b.name AS manager
FROM employees a
JOIN employees b ON a.manager_id = b.id;
```

✔ 같은 테이블에서 부모-자식, 상하 관계 등을 표현  
✔ 테이블에 별칭으로로 구분

---

## 8️⃣ CROSS JOIN 

```sql
SELECT color, size
FROM colors
CROSS JOIN sizes;
```

✔ 각 색상마다 모든 사이즈 조합 반환 (카테시안 곱)  
✔ `n × m` 개 결과 생성  

---

## 9️⃣ 예시 데이터로 보는 JOIN

#### 1. `users` 테이블

| id | name  |
|----|-------|
| 1  | Alice |
| 2  | Bob   |
| 3  | Carol |

#### 2. `orders` 테이블

| id | user_id | item   |
|----|---------|--------|
| 1  | 1       | Book   |
| 2  | 2       | Pen    |

#### 3. INNER JOIN 결과
```sql
SELECT u.name, o.item
FROM users u
INNER JOIN orders o ON u.id = o.user_id;
```

```text
Alice | Book  
Bob   | Pen
```

#### 4. LEFT JOIN 결과
```sql
SELECT u.name, o.item
FROM users u
LEFT JOIN orders o ON u.id = o.user_id;
```

```text
Alice | Book  
Bob   | Pen  
Carol | NULL
```

---

### 🎯 정리

| JOIN 종류 | 설명 |
|-----------|------|
| INNER JOIN | 양쪽 테이블에 모두 존재하는 경우만 |
| LEFT JOIN | 왼쪽 테이블 모두 + 오른쪽 매칭된 데이터 |
| RIGHT JOIN | 오른쪽 테이블 모두 + 왼쪽 매칭된 데이터 |
| FULL OUTER JOIN | 양쪽 모두 포함 (없으면 NULL) |
| SELF JOIN | 자기 자신과 조인 |
| CROSS JOIN | 가능한 모든 조합 반환 |
