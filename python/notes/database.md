# 🔄 Python 데이터베이스 (Database)

Python에서는 `sqlite3`, `MySQL`, `PostgreSQL` 등의 다양한 데이터베이스를 사용할 수 있습니다.  
이 문서에서는 **SQLite**를 기본 예제로 사용하여 Python에서 데이터베이스를 다루는 방법을 설명합니다.

---

## 1️⃣ 데이터베이스 개념

- 데이터베이스(Database)란 데이터를 저장하고 관리하는 시스템입니다.
- Python에서는 `sqlite3` 모듈을 사용하여 **SQLite** 데이터베이스를 쉽게 다룰 수 있습니다.
- `MySQL`, `PostgreSQL` 등 다른 데이터베이스도 `pymysql`, `psycopg2` 등의 라이브러리를 사용하여 연동할 수 있습니다.

---

## 2️⃣ SQLite란?

- **SQLite**는 파일 기반의 가벼운 데이터베이스 관리 시스템(DBMS)입니다.
- **서버가 필요하지 않으며**, 로컬 파일(`.db`)에 데이터를 저장합니다.
- Python의 **표준 라이브러리**에 포함되어 있어 별도의 설치 없이 사용할 수 있습니다.

#### SQLite 데이터베이스 파일 생성
```python
import sqlite3

# 데이터베이스 연결 (없으면 자동 생성)
conn = sqlite3.connect("example.db")

# 연결 종료
conn.close()
```

✔ 위 코드를 실행하면 `example.db` 파일이 생성됩니다.

---

## 3️⃣ 테이블 생성

데이터를 저장하기 위해 **테이블(Table)** 을 생성해야 합니다.

#### 테이블 생성 예제
```python
import sqlite3

# 데이터베이스 연결
conn = sqlite3.connect("example.db")
cursor = conn.cursor()  # SQL 실행을 위한 커서 생성

# 테이블 생성 (users)
cursor.execute('''
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        age INTEGER,
        email TEXT UNIQUE
    )
''')

conn.commit()  # 변경 사항 저장
conn.close()   # 연결 종료
```

✔ `CREATE TABLE` 문을 사용하여 테이블을 생성  
✔ `PRIMARY KEY`는 고유 ID를 자동 생성  
✔ `AUTOINCREMENT`를 사용하면 ID가 자동 증가  

---

## 4️⃣ 데이터 삽입 (INSERT)

데이터를 테이블에 추가하려면 `INSERT INTO` 문을 사용합니다.

### 1) 단일 데이터 삽입
```python
import sqlite3

conn = sqlite3.connect("example.db")
cursor = conn.cursor()

cursor.execute("INSERT INTO users (name, age, email) VALUES (?, ?, ?)", 
               ("Alice", 25, "alice@example.com"))

conn.commit()
conn.close()
```

✔ **`?` 플레이스홀더**를 사용하여 SQL Injection 방지  
✔ `commit()`을 호출해야 변경 사항이 저장됨  

---

### 2) 여러 데이터 삽입
```python
users = [
    ("Bob", 30, "bob@example.com"),
    ("Charlie", 28, "charlie@example.com"),
    ("David", 35, "david@example.com")
]

cursor.executemany("INSERT INTO users (name, age, email) VALUES (?, ?, ?)", users)

conn.commit()
conn.close()
```

✔ `executemany()`를 사용하면 여러 개의 데이터를 한 번에 삽입 가능  

---

## 5️⃣ 데이터 조회 (SELECT)

저장된 데이터를 가져오려면 `SELECT` 문을 사용합니다.

### 1) 모든 데이터 조회
```python
conn = sqlite3.connect("example.db")
cursor = conn.cursor()

cursor.execute("SELECT * FROM users")
rows = cursor.fetchall()  # 모든 행 가져오기

for row in rows:
    print(row)  # (1, 'Alice', 25, 'alice@example.com')

conn.close()
```

✔ `fetchall()` → 모든 결과를 리스트 형태로 반환  
✔ `fetchone()` → 첫 번째 행만 반환  

---

### 2) `WHERE` - 특정 데이터 조회
```python
cursor.execute("SELECT * FROM users WHERE age > ?", (28,))
rows = cursor.fetchall()

for row in rows:
    print(row)  # age가 28보다 큰 사용자 출력
```

✔ `WHERE` 조건을 사용하여 특정 데이터를 필터링 가능  

---

## 6️⃣ `UPDATE` - 데이터 수정

#### 특정 데이터 수정 예제
```python
cursor.execute("UPDATE users SET age = ? WHERE name = ?", (26, "Alice"))
conn.commit()
```

✔ `UPDATE` 문을 사용하여 특정 데이터를 변경  

---

## 7️⃣ `DELETE` - 데이터 삭제

### 특정 데이터 삭제 예제
```python
cursor.execute("DELETE FROM users WHERE name = ?", ("Charlie",))
conn.commit()
```

✔ `DELETE` 문을 사용하여 특정 데이터를 삭제  

---

## 8️⃣ `try-except` - 예외 처리

데이터베이스 작업 중 오류를 방지하기 위해 **예외 처리**를 추가할 수 있습니다.

#### 예외 처리 예제
```python
try:
    conn = sqlite3.connect("example.db")
    cursor = conn.cursor()

    cursor.execute("INSERT INTO users (name, age, email) VALUES (?, ?, ?)", 
                   ("Eve", 22, "eve@example.com"))

    conn.commit()
except sqlite3.Error as e:
    print(f"데이터베이스 오류 발생: {e}")
finally:
    conn.close()  # 오류가 발생하더라도 연결을 닫음
```

✔ `try-except-finally` 구문을 사용하여 **안전한 데이터베이스 작업 가능**  

---

## 9️⃣ MySQL과 PostgreSQL 사용법

SQLite 외에도 **MySQL** 및 **PostgreSQL** 데이터베이스를 사용할 수 있습니다.

### 1) MySQL 사용 (`pymysql` 패키지 필요)
```sh
pip install pymysql
```
```python
import pymysql

conn = pymysql.connect(host="localhost", user="root", password="1234", database="test_db")
cursor = conn.cursor()
cursor.execute("SELECT * FROM users")
rows = cursor.fetchall()

for row in rows:
    print(row)

conn.close()
```

---

### 2) PostgreSQL 사용 (`psycopg2` 패키지 필요)
```sh
pip install psycopg2
```
```python
import psycopg2

conn = psycopg2.connect(host="localhost", user="postgres", password="1234", dbname="test_db")
cursor = conn.cursor()
cursor.execute("SELECT * FROM users")
rows = cursor.fetchall()

for row in rows:
    print(row)

conn.close()
```

✔ **SQLite와 SQL 문법은 동일하지만, MySQL과 PostgreSQL에서는 드라이버가 필요함**  

---

## 🎯 정리

✔ **SQLite** → 파일 기반의 경량 데이터베이스 (`sqlite3` 모듈 사용)  
✔ **테이블 생성** → `CREATE TABLE` 문으로 데이터 구조 정의  
✔ **데이터 삽입** → `INSERT INTO` 문으로 데이터 추가  
✔ **데이터 조회** → `SELECT` 문을 사용하여 검색  
✔ **데이터 수정** → `UPDATE` 문을 사용하여 특정 데이터 변경  
✔ **데이터 삭제** → `DELETE` 문을 사용하여 특정 데이터 제거  
✔ **예외 처리** → `try-except`를 사용하여 안전한 데이터 처리  
✔ **MySQL 및 PostgreSQL** → `pymysql`, `psycopg2` 등의 라이브러리를 사용하여 연결 가능  
