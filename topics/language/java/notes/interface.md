# 🧩 Java - 추상 클래스 & 인터페이스

Java는 객체지향 언어로서 **공통된 동작의 설계**를 위해  
**추상 클래스(Abstract Class)** 와 **인터페이스(Interface)** 를 제공합니다.

---

## 1️⃣ 추상 클래스란?

- `abstract` 키워드로 선언  
- **일부만 구현**되어 있고, **나머지는 자식이 구현**  
- **객체 생성 불가** → 반드시 상속해서 사용

```java
abstract class Animal {
    abstract void sound();  // 추상 메서드
    void breathe() {
        System.out.println("숨을 쉰다");
    }
}
```

---

## 2️⃣ 추상 클래스 상속 & 구현

```java
class Dog extends Animal {
    @Override
    void sound() {
        System.out.println("멍멍!");
    }
}
```

```java
Animal dog = new Dog();  // 다형성
dog.sound();              // 멍멍!
dog.breathe();            // 숨을 쉰다
```

---

## 3️⃣ 추상 메서드

- 메서드 본문 없이 **선언만 존재**
- 자식 클래스에서 반드시 `@Override` 해야 함

```java
abstract void run();
```

---

## 4️⃣ 인터페이스란?

- **모든 메서드가 추상 메서드**인 설계도  
- `implements` 키워드로 구현  
- **다중 구현 가능** (Java는 다중 상속은 안 되지만 인터페이스는 가능)

```java
interface Flyable {
    void fly();
}
```

---

## 5️⃣ 인터페이스 구현

```java
class Bird implements Flyable {
    @Override
    public void fly() {
        System.out.println("날아간다!");
    }
}
```

✔ 인터페이스는 **메서드는 자동으로 public abstract**, 구현 시 `public` 생략하면 컴파일 에러 발생

---

## 6️⃣ 다중 인터페이스 구현

```java
interface Walkable {
    void walk();
}

interface Swimmable {
    void swim();
}

class Penguin implements Walkable, Swimmable {
    public void walk() { System.out.println("걷는다"); }
    public void swim() { System.out.println("수영한다"); }
}
```

---

## 7️⃣ 인터페이스 vs 추상 클래스

| 항목 | 추상 클래스 | 인터페이스 |
|------|-------------|-------------|
| 키워드 | `abstract class` | `interface` |
| 상속/구현 | `extends` | `implements` |
| 다중 구현 | ❌ 불가능 | ✅ 가능 |
| 필드 | 가질 수 있음 | 상수(public static final)만 |
| 메서드 구현 | 가능 | Java 8 이후 default/static만 가능 |
| 용도 | 기본 동작 + 공통 틀 제공 | 기능 계약 (기능 보장) |

---

## 8️⃣ default 메서드 (Java 8+)

인터페이스에서도 **기본 구현을 제공**할 수 있음

```java
interface Printer {
    default void print() {
        System.out.println("기본 출력");
    }
}
```

---

## 9️⃣ static 메서드 (Java 8+)

```java
interface Utils {
    static void hello() {
        System.out.println("Hello from static!");
    }
}
```

> 호출: `Utils.hello();`

---

## 🔟 Functional Interface (Java 8+)

- **메서드가 딱 1개인 인터페이스**
- 람다 표현식으로 사용 가능
- `@FunctionalInterface` 애노테이션 사용

```java
@FunctionalInterface
interface Calculator {
    int calculate(int a, int b);
}

Calculator add = (a, b) -> a + b;
System.out.println(add.calculate(3, 5));  // 결과: 8
```

---

## 🎯 정리

✔ `abstract class` → 일부만 구현된 클래스, 상속 전용  
✔ `abstract method` → 자식 클래스가 반드시 구현해야 하는 메서드  
✔ `interface` → 모든 메서드가 추상, 다중 구현 가능  
✔ `implements` → 인터페이스 구현할 때 사용하는 키워드  
✔ `extends` → 클래스나 추상 클래스를 상속할 때 사용  
✔ `default method` → 인터페이스에 기본 구현 제공 (Java 8+)  
✔ `static method` → 클래스처럼 호출 가능한 유틸성 메서드  
✔ `@FunctionalInterface` → 메서드 하나만 갖는 인터페이스, 람다용  
✔ `다중 인터페이스` → 여러 개의 interface를 동시에 구현 가능  
✔ `객체 생성 불가` → 추상 클래스와 인터페이스는 직접 인스턴스 생성 불가

