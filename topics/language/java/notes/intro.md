# ☕ Java란?

Java는 플랫폼에 독립적인 객체지향 프로그래밍 언어로, **Write Once, Run Anywhere (WORA)** 철학에 기반하여 다양한 환경에서 실행될 수 있습니다.  

---

## 1️⃣ 특징 요약

| 항목 | 설명 |
|------|------|
| 플랫폼 독립성 | JVM 위에서 실행되므로 OS에 종속되지 않음 |
| 객체지향 프로그래밍 | 캡슐화, 상속, 다형성 등 OOP 개념 지원 |
| 자동 메모리 관리 | 가비지 컬렉터를 통해 메모리 해제 자동화 |
| 풍부한 표준 라이브러리 | 다양한 네트워크, 파일, 유틸리티 API 제공 |
| 멀티스레딩 지원 | 내장된 Thread 클래스로 병렬 처리 가능 |

---

## 2️⃣ Java 실행 구조

#### 1. `.java` 파일 → 자바 소스 코드
HelloWorld.java 파일 생성

#### 2. `javac`로 컴파일 → `.class` 바이트코드 생성
```sh
javac HelloWorld.java   # 컴파일
```

#### 3. `java` 명령으로 JVM에서 실행
```sh
java HelloWorld         # 실행
```

---

## 3️⃣ Java 개발 도구 (JDK 구성요소)

| 구성 요소 | 설명 |
|-----------|------|
| `javac` | 자바 컴파일러 |
| `java` | 자바 애플리케이션 실행기 |
| `javadoc` | 문서 생성 도구 |
| `jar` | Java Archive, 패키징 도구 |

---

## 4️⃣ 예제 코드

```java
public class HelloWorld {
    public static void main(String[] args) {
        System.out.println("Hello, Java!");
    }
}
```

---

## 5️⃣ 주요 개념 용어

| 용어 | 설명 |
|------|------|
| JDK | Java Development Kit - 개발에 필요한 도구 모음 |
| JRE | Java Runtime Environment - 실행에 필요한 환경 |
| JVM | Java Virtual Machine - 바이트코드 실행 엔진 |
