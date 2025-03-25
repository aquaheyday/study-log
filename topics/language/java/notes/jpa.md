# 🗃️ Java - JPA 개요

**JPA(Java Persistence API)** 는 자바 객체(Object)와 관계형 데이터베이스(RDB) 간의 **매핑(ORM)** 을 지원하는 **자바 표준 인터페이스**입니다.

---

## 1️⃣ JPA란?

- 자바 ORM(Object-Relational Mapping) 기술의 **표준 스펙**
- 자바 객체와 DB 테이블 간의 매핑을 선언적으로 처리
- SQL을 직접 작성하지 않고, 객체 지향적으로 DB를 조작 가능

---

## 2️⃣ ORM(Object-Relational Mapping)이란?

- **객체(클래스)** ↔ **테이블**  
- **필드** ↔ **컬럼**  
- **객체 간 연관관계** ↔ **외래키(FK)**

✔ 즉, 객체를 DB에 자동으로 저장/조회하는 기술

---

## 3️⃣ JPA vs Hibernate

| 항목 | JPA | Hibernate |
|------|-----|-----------|
| 개념 | 자바 ORM 표준 | JPA 구현체 (가장 대표적) |
| 역할 | 인터페이스 | 실제 동작을 구현하는 클래스 |
| 사용 | `javax.persistence.*` | 내부적으로 `org.hibernate.*` 사용 |

✔ 실무에서는 Hibernate를 통해 JPA를 사용하는 경우가 많습니다.

---

## 4️⃣ JPA의 핵심 구성요소

| 구성요소 | 설명 |
|----------|------|
| `Entity` | DB 테이블에 매핑되는 클래스 (`@Entity`) |
| `EntityManager` | 엔티티 저장/조회/수정/삭제 역할 |
| `Persistence Context` | 엔티티를 관리하는 1차 캐시 (영속성 컨텍스트) |
| `JPQL` | 객체 지향 쿼리 언어 (SQL 유사) |

---

## 5️⃣ 간단한 Entity 클래스 예시

```java
@Entity
public class User {
    
    @Id @GeneratedValue
    private Long id;

    private String name;

    @Column(name = "email_address")
    private String email;

    // getter/setter
}
```

---

## 6️⃣ JPA 기본 동작 흐름

#### 1. `EntityManager`를 통해 엔티티 저장 요청
#### 2. 영속성 컨텍스트(Persistence Context)에 저장
#### 3. 트랜잭션 커밋 시 실제 DB에 SQL 실행 (쓰기 지연)
#### 4. 조회 시 1차 캐시에서 먼저 찾음 → 없으면 DB 조회

---

## 7️⃣ Spring Data JPA 소개

Spring에서 JPA를 쉽게 사용할 수 있도록 도와주는 모듈

```java
public interface UserRepository extends JpaRepository<User, Long> {
    List<User> findByName(String name);
}
```

✔ 복잡한 쿼리 없이도 메서드 이름만으로 자동 SQL 생성 가능

---

## 8️⃣ 주요 어노테이션 요약

| 어노테이션 | 설명 |
|------------|------|
| `@Entity` | JPA가 관리할 클래스 지정 |
| `@Id` | 기본 키 지정 |
| `@GeneratedValue` | 자동 생성 전략 지정 (AUTO, IDENTITY 등) |
| `@Column` | 컬럼 설정 (이름, 길이 등) |
| `@Table` | 테이블명, 제약 조건 등 지정 |
| `@ManyToOne`, `@OneToMany` | 연관관계 매핑 |

---

## 9️⃣ JPA의 장점

- SQL 작성 최소화 (객체 중심 개발)
- 생산성 향상 & 유지보수 쉬움
- DB 독립성 확보 가능
- 트랜잭션 관리와 캐시 자동 처리

---

## 🔟 주의할 점

- 복잡한 쿼리나 성능 최적화는 여전히 SQL 또는 QueryDSL 필요
- 연관관계, 지연 로딩, N+1 문제 등 주의 필요
- 너무 자동화에 의존하면 추적/디버깅 어려움

---

## 🎯 정리

✔ JPA는 자바에서 ORM을 위한 **표준 API**  
✔ Hibernate는 JPA의 **구현체**  
✔ `@Entity`로 클래스 ↔ 테이블 매핑  
✔ `EntityManager`로 영속성 관리  
✔ `JPQL`을 사용해 객체 중심 쿼리 작성  
✔ Spring Data JPA로 **자동 CRUD 메서드 생성** 가능  
✔ SQL보다 객체 중심 개발로 생산성 향상  
✔ 단, 복잡한 쿼리나 성능은 QueryDSL 등 별도 고려 필요

