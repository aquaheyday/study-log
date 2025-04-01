# 🛡️ Secure Coding - SQL Injection 대응

**SQL Injection(SQLi)** 은 공격자가 사용자 입력값을 악용하여 **원래 의도하지 않은 SQL 쿼리를 실행**시키는 공격입니다.  
데이터 유출, 조작, 삭제는 물론이고 **전체 시스템 탈취**로도 이어질 수 있는 치명적인 취약점입니다.

---

## 1️⃣ SQL Injection의 위험성

| 영향               | 설명 |
|--------------------|------|
| 인증 우회           | 로그인 우회 (`' OR '1'='1`) |
| 데이터 유출         | 민감 정보 추출 (`SELECT * FROM users`) |
| 데이터 변경/삭제    | `UPDATE`, `DELETE` 조작 |
| 시스템 명령 실행     | DB 계정 권한에 따라 시스템 공격 가능 |

---

## 2️⃣ 공격 예시

#### ❌ 취약한 코드 예시 (Node.js)

```js
const query = `SELECT * FROM users WHERE id = '${req.query.id}'`;
db.query(query, (err, result) => { ... });
```

> 입력값: `1' OR '1'='1`  
> 실제 실행: `SELECT * FROM users WHERE id = '1' OR '1'='1'` → 모든 사용자 조회됨

---

## 3️⃣ SQL Injection 대응 전략

| 대응 방법             | 설명 |
|-----------------------|------|
| **Prepared Statement** | 입력값을 파라미터로 안전하게 분리 |
| **ORM 사용**          | Sequelize, TypeORM 등 안전한 쿼리 생성 |
| **입력 검증**          | 예상되는 입력값 형식 검증 (예: 숫자만 허용) |
| **최소 권한 DB 계정**  | 읽기/쓰기 제한, DROP 권한 차단 등 |
| **에러 메시지 숨김**   | SQL 오류로 내부 구조 노출 방지 |

---

## 4️⃣ 안전한 코드 예시

#### 1. Prepared Statement (Node.js - mysql2)

```js
const id = req.query.id;
const [rows] = await db.execute("SELECT * FROM users WHERE id = ?", [id]);
```

#### 2. ORM 사용 예시 (Sequelize)

```js
const user = await User.findOne({ where: { id: req.query.id } });
```

---

## 5️⃣ 입력 검증 예시 (숫자만 허용)

```js
const id = req.query.id;
if (!/^\d+$/.test(id)) {
  return res.status(400).send("Invalid ID format");
}
```

---

## 6️⃣ SQLi 대응 체크리스트

- [ ] 모든 SQL 쿼리에 Prepared Statement 또는 ORM을 사용하고 있는가?
- [ ] 사용자 입력값에 대해 엄격한 검증이 적용되어 있는가?
- [ ] 데이터베이스 연결 계정은 최소 권한으로 설정되어 있는가?
- [ ] SQL 에러 메시지는 사용자에게 노출되지 않는가?
- [ ] 로그/모니터링 시스템으로 이상 쿼리 패턴을 감지하고 있는가?

---

## 🎯 정리

✔ SQL Injection은 **사용자 입력이 쿼리에 그대로 삽입되어 발생**  
✔ Prepared Statement나 ORM 사용으로 **입력값을 안전하게 분리**  
✔ 입력 검증, DB 권한 최소화, 에러 메시지 숨김 등 **다중 방어** 적용  
