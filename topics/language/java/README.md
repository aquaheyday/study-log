# ☕ Java

이 폴더는 **Java 학습 과정**에서 정리한 자료, 예제 코드, 프로젝트를 보관하는 공간입니다.  
기본 문법부터 객체지향 개념, 예외 처리, 컬렉션, 스레드, 그리고 스프링 등 실무와 연결되는 내용까지 단계별로 정리합니다.

---

## 📂 디렉토리 구성

| 폴더명 | 설명 |
|---|---|
| [notes](./notes) | 개념 정리 및 학습 노트 |
| [examples](./examples) | 주요 기능별 예제 코드 |

---

## 📋 Java 언어 개념 정리 목록

### 📌 기본 개념
| 주제 | 파일명 | 설명 |
|---|---|---|
| Java 소개 | [intro.md](./notes/intro.md) | 특징, 플랫폼 독립성, JVM 개요 |
| 개발 환경 구성 | [setup.md](./notes/setup.md) | JDK 설치, 컴파일/실행 구조 |
| 변수와 타입 | [variables.md](./notes/variables.md) | 원시 타입, 참조 타입 |
| 제어문 | [control.md](./notes/control.md) | 조건문, 반복문, switch 사용 |

### 🧱 클래스와 객체지향
| 주제 | 파일명 | 설명 |
|---|---|---|
| 클래스 & 객체 | [class-object.md](./notes/class-object.md) | 인스턴스 생성, 필드, 메서드 |
| 상속과 다형성 | [inheritance.md](./notes/inheritance.md) | `extends`, 오버라이딩 개념 |
| 추상 클래스 & 인터페이스 | [interface.md](./notes/interface.md) | 설계 기반 추상화 구현 |
| 접근 제어자 | [access-modifiers.md](./notes/access-modifiers.md) | public, private, protected 개념 |

### 🧪 예외 처리 & API
| 주제 | 파일명 | 설명 |
|---|---|---|
| 예외 처리 | [exception.md](./notes/exception.md) | try-catch, throw/throws, 사용자 정의 예외 |
| 문자열 처리 | [string.md](./notes/string.md) | `String`, `StringBuilder`, 메서드 활용 |
| 날짜와 시간 | [datetime.md](./notes/datetime.md) | LocalDate, DateTimeFormatter 사용법 |

### 📚 컬렉션 & 제네릭
| 주제 | 파일명 | 설명 |
|---|---|---|
| 배열 | [array.md](./notes/array.md) | 1차원, 2차원 배열 선언과 활용 |
| List, Set, Map | [collections.md](./notes/collections.md) | ArrayList, HashSet, HashMap 등 |
| 제네릭 | [generics.md](./notes/generics.md) | 타입 안정성 확보, 제네릭 클래스/메서드 |

### 🧵 쓰레드 & 동시성
| 주제 | 파일명 | 설명 |
|---|---|---|
| 스레드 기초 | [thread.md](./notes/thread.md) | `Thread`, `Runnable` 구현 방식 |
| 동기화 & Lock | [synchronization.md](./notes/synchronization.md) | synchronized 키워드, 경쟁 조건 방지 |
| Executor | [executor.md](./notes/executor.md) | ThreadPoolExecutor 등 병렬 처리 도구 |

### 🌍 웹 & 스프링 입문
| 주제 | 파일명 | 설명 |
|---|---|---|
| Java로 웹 서버 만들기 | [http-server.md](./notes/http-server.md) | 기본 Servlet 사용법 |
| Spring Boot 시작하기 | [spring-boot.md](./notes/spring-boot.md) | 초기 설정, 기본 컨트롤러 구현 |
| REST API 설계 | [rest-api.md](./notes/rest-api.md) | 요청 매핑, 응답 객체, JSON 처리 |
| JPA 개요 | [jpa.md](./notes/jpa.md) | 엔티티, Repository, ORM 기초 개념 |

---

## 📋 예제 목록

| 주제 | 파일명 | 설명 |
|---|---|---|
| 변수와 연산자 | [Variables.java](./examples/Variables.java) | 변수 선언, 기본 연산 |
| 제어문 | [ControlFlow.java](./examples/ControlFlow.java) | 조건문, 반복문 |
| 클래스 정의 | [ClassExample.java](./examples/ClassExample.java) | 필드, 생성자, 메서드 |
| 상속 & 오버라이딩 | [Inheritance.java](./examples/Inheritance.java) | extends, super 사용 |
| 인터페이스 구현 | [InterfaceExample.java](./examples/InterfaceExample.java) | 다형성과 추상화 예시 |
| 예외 처리 | [TryCatch.java](./examples/TryCatch.java) | checked / unchecked 예외 |
| ArrayList 사용 | [ArrayListExample.java](./examples/ArrayListExample.java) | 자료구조 활용 예제 |
| 파일 읽기/쓰기 | [FileIO.java](./examples/FileIO.java) | BufferedReader, FileWriter 사용 |
| 멀티스레드 | [ThreadExample.java](./examples/ThreadExample.java) | Thread, Runnable |
| Spring REST API | [RestApiExample.java](./examples/RestApiExample.java) | Controller, @GetMapping 예제 |

---

## 📢 업데이트 로그

- 2025-03-25: 초기 구성
- 예제 코드 및 개념 정리는 계속 추가 예정

