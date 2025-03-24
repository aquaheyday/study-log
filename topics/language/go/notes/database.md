# 🗄️ Go 언어 데이터베이스 연동

Go는 표준 패키지 `database/sql`을 통해 다양한 관계형 데이터베이스와 연동할 수 있습니다.  
실제 DB 연결은 별도의 **드라이버**를 사용하며, `sql.DB` 객체를 통해 쿼리를 수행합니다.

---

## 1️⃣ 드라이버 설치 예시
- 드라이버는 반드시 `import _` 형식으로 등록해야 함  
- 드라이버별 연결 문자열(DNS)은 다름

### 1) MySQL

```bash
go get -u github.com/go-sql-driver/mysql
```

```go
import _ "github.com/go-sql-driver/mysql"
```

---

### 2) PostgreSQL

```bash
go get -u github.com/lib/pq
```

```go
import _ "github.com/lib/pq"
```

---

### 3) SQLite

```bash
go get -u github.com/mattn/go-sqlite3
```

```go
import _ "github.com/mattn/go-sqlite3"
```

---

## 2️⃣ DB 연결하기

```go
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/mydb")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    err = db.Ping() // 연결 확인
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("DB 연결 성공")
}
```

✔ `sql.Open()`은 실제 연결을 열지 않음 → `db.Ping()`으로 확인  
✔ `defer db.Close()`로 연결 닫기  

---

## 3️⃣ 단일 행 조회 (`QueryRow`)

```go
var name string
err := db.QueryRow("SELECT name FROM users WHERE id = ?", 1).Scan(&name)
if err != nil {
    log.Fatal(err)
}
fmt.Println("Name:", name)
```

✔ `QueryRow()`는 결과가 하나일 때 사용  
✔ `.Scan()`에 각 컬럼 값을 받을 포인터 전달  

---

## 4️⃣ 여러 행 조회 (`Query`)

```go
rows, err := db.Query("SELECT id, name FROM users")
if err != nil {
    log.Fatal(err)
}
defer rows.Close()

for rows.Next() {
    var id int
    var name string
    rows.Scan(&id, &name)
    fmt.Println(id, name)
}
```

✔ `rows.Next()`로 반복  
✔ `defer rows.Close()`는 꼭 호출  

---

## 5️⃣ INSERT, UPDATE, DELETE

```go
res, err := db.Exec("INSERT INTO users (name) VALUES (?)", "Alice")
if err != nil {
    log.Fatal(err)
}

id, _ := res.LastInsertId()
fmt.Println("새 ID:", id)
```

✔ `Exec()`은 결과 없이 실행하는 쿼리에 사용  
✔ `LastInsertId()` 또는 `RowsAffected()`로 결과 확인  

---

## 6️⃣ Prepare 사용 (성능 + 보안)

```go
stmt, err := db.Prepare("SELECT name FROM users WHERE id = ?")
if err != nil {
    log.Fatal(err)
}
defer stmt.Close()

var name string
stmt.QueryRow(1).Scan(&name)
```

✔ SQL 문을 미리 준비 → 성능 향상  
✔ 반복되는 쿼리에 유리  
✔ SQL Injection 방지에도 도움  

---

## 7️⃣ 트랜잭션 처리

```go
tx, err := db.Begin()
if err != nil {
    log.Fatal(err)
}

_, err = tx.Exec("UPDATE accounts SET balance = balance - 100 WHERE id = ?", 1)
if err != nil {
    tx.Rollback()
    log.Fatal(err)
}

_, err = tx.Exec("UPDATE accounts SET balance = balance + 100 WHERE id = ?", 2)
if err != nil {
    tx.Rollback()
    log.Fatal(err)
}

err = tx.Commit()
if err != nil {
    log.Fatal(err)
}
```

✔ `db.Begin()` → 트랜잭션 시작  
✔ `tx.Exec()`로 실행, `Commit()` 또는 `Rollback()`으로 마무리  
✔ 예외 발생 시 꼭 `Rollback()` 호출  

---

## 8️⃣ 에러 처리

```go
if errors.Is(err, sql.ErrNoRows) {
    fmt.Println("결과 없음")
} else if err != nil {
    log.Fatal(err)
}
```

✔ `sql.ErrNoRows`는 조회 결과가 없을 때 발생하는 에러  
✔ 에러 타입 분기처리 시 `errors.Is()` 사용  

---

## 9️⃣ 추천 ORM 패키지

| 패키지 | 특징 |
|--------|------|
| [GORM](https://gorm.io) | 가장 널리 사용되는 ORM, 풍부한 기능 |
| [sqlx](https://github.com/jmoiron/sqlx) | `database/sql` 확장, 더 편리한 바인딩 |
| [ent](https://entgo.io) | Graph 기반 코드 생성형 ORM |

---

## 🎯 정리

✔ 연결 풀은 `sql.DB` 내부에서 자동 관리됨  
✔ 에러 처리는 꼭! `err != nil` 체크  
✔ SQL Injection을 막기 위해 `?` 플레이스홀더 + 인자 분리 사용  
✔ 대규모 프로젝트에선 ORM(GORM 등) 도입도 고려 가능
