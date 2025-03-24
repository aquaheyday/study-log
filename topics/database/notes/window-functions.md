# 🪟 SQL 윈도우 함수(Window Function)

**윈도우 함수**는 **행(Row) 간 비교, 집계, 순위 계산** 등 복잡한 분석 작업을 테이블의 **행 단위로 계산**할 수 있게 해주는 SQL 기능입니다.

---

## 1️⃣ 윈도우 함수란?

`OVER()` 절을 사용해 **행 그룹(윈도우)** 내에서 각 행에 대한 **누적, 순위, 비교, 통계** 등을 수행하는 함수  

✔ 일반 집계 함수(SUM, AVG 등)와 달리  
✔ **모든 행을 유지하면서 추가 컬럼처럼 값 계산**

---

## 2️⃣ 기본 문법

```sql
SELECT col1,
       RANK() OVER (PARTITION BY dept ORDER BY salary DESC) AS rank_in_dept
FROM employees;
```

| 구성 요소 | 설명 |
|-----------|------|
| `OVER()`        | 윈도우 함수임을 명시 |
| `PARTITION BY` | 윈도우(그룹) 구분 기준 |
| `ORDER BY`     | 윈도우 내부 정렬 기준 |

---

## 3️⃣ 주요 윈도우 함수

| 함수 | 설명 |
|------|------|
| `ROW_NUMBER()` | 정렬 기준에 따라 행 번호 부여 |
| `RANK()` | 동일 순위 부여, 건너뛰기 있음 |
| `DENSE_RANK()` | 동일 순위 부여, 건너뛰기 없음 |
| `SUM()`, `AVG()` | 누적 합계/평균 |
| `LAG()`, `LEAD()` | 이전/다음 행 값 참조 |
| `FIRST_VALUE()`, `LAST_VALUE()` | 그룹 내 첫/마지막 값 |

---

## 4️⃣ 순위 함수 예제

```sql
SELECT name, dept, salary,
  RANK() OVER (PARTITION BY dept ORDER BY salary DESC) AS dept_rank
FROM employees;
```

✔ 부서(dept)별 급여 순위  
✔ 동점자 존재 시 순위 건너뜀 (`RANK()` vs `DENSE_RANK()` 차이)  

---

## 5️⃣ 누적 합계 (Running Total)

```sql
SELECT name, salary,
  SUM(salary) OVER (ORDER BY name) AS running_total
FROM employees;
```

✔ 이름순 정렬 기준으로 **누적 합계** 계산  
✔ 실무에서 매출 누적, 시간 순 누적값 등 자주 사용  

---

## 6️⃣ 이전/다음 행 참조 (LAG / LEAD)

```sql
SELECT name, salary,
  LAG(salary) OVER (ORDER BY id) AS prev_salary,
  LEAD(salary) OVER (ORDER BY id) AS next_salary
FROM employees;
```

✔ `LAG()` = 이전 행 값  
✔ `LEAD()` = 다음 행 값  
✔ 전월 대비 매출, 변화량 분석 등에 사용

---

## 7️⃣ PARTITION BY 없이 사용

```sql
SELECT name,
  ROW_NUMBER() OVER (ORDER BY created_at DESC) AS recent_rank
FROM users;
```

✔ 전체 행을 하나의 그룹으로 간주  
✔ 최근 가입자 순위 등 전체 기준 순위 매기기

---

## 8️⃣ 실무 예제: 최근 주문 순번 + 누적 금액

```sql
SELECT user_id, order_date, amount,
  ROW_NUMBER() OVER (PARTITION BY user_id ORDER BY order_date DESC) AS recent_order,
  SUM(amount) OVER (PARTITION BY user_id ORDER BY order_date) AS total_spent
FROM orders;
```

✔ 사용자별 최근 주문 순번  
✔ 사용자별 누적 주문 금액 계산

---

## 9️⃣ 윈도우 함수 vs 집계 함수

| 항목 | 집계 함수 | 윈도우 함수 |
|------|-----------|-------------|
| 예: `SUM()` | 전체/그룹 결과 1행만 반환 | 각 행마다 누적/그룹 값 추가 |
| 결과 형태 | 행 수 줄어듦 | 행 수 유지 |
| 용도 | 요약 | 분석/비교 |

---

## 🎯 정리

| 기능 | 함수 예시 |
|------|-----------|
| 순위/번호 | `ROW_NUMBER()`, `RANK()`, `DENSE_RANK()` |
| 누적 | `SUM() OVER`, `AVG() OVER` |
| 이전/다음 값 | `LAG()`, `LEAD()` |
| 처음/끝 값 | `FIRST_VALUE()`, `LAST_VALUE()` |
| 그룹 나누기 | `PARTITION BY` |
| 정렬 기준 | `ORDER BY` |
