# 🧬 Java - 상속과 다형성

Java는 객체지향 프로그래밍 언어로,  
**상속(Inheritance)** 과 **다형성(Polymorphism)** 은 코드 재사용성과 유연성을 높여주는 핵심 개념입니다.

---

## 1️⃣ 상속이란?

- 기존 클래스를 **확장(extends)** 하여 **새로운 클래스를 정의**하는 것
- 부모 클래스의 필드와 메서드를 자식 클래스에서 재사용 가능

```java
class Animal {
    void sound() {
        System.out.println("동물이 소리를 냅니다.");
    }
}

class Dog extends Animal {
    void bark() {
        System.out.println("멍멍!");
    }
}
```

```java
Dog d = new Dog();
d.sound();  // 부모 클래스 메서드
d.bark();   // 자식 클래스 메서드
```

---

## 2️⃣ `extends` 키워드

- 클래스 상속 시 사용
- Java는 **단일 상속**만 지원

```java
class Parent { }
class Child extends Parent { }  // OK

class Another extends Parent, OtherClass { }  // ❌ 에러: 다중 상속 불가
```

---

## 3️⃣ 메서드 오버라이딩 (Overriding)

- 부모 클래스의 메서드를 **자식 클래스에서 재정의**
- `@Override` 애노테이션 사용 권장

```java
class Animal {
    void sound() {
        System.out.println("동물이 소리를 냅니다.");
    }
}

class Cat extends Animal {
    @Override
    void sound() {
        System.out.println("야옹~");
    }
}
```

---

## 4️⃣ super 키워드

- 부모 클래스의 **생성자, 메서드, 필드**에 접근할 때 사용

```java
class Parent {
    void greet() {
        System.out.println("안녕하세요!");
    }
}

class Child extends Parent {
    void greet() {
        super.greet();  // 부모 메서드 호출
        System.out.println("반갑습니다!");
    }
}
```

---

## 5️⃣ 다형성이란?

- **하나의 객체가 여러 타입으로 동작**할 수 있는 특성
- 자식 객체를 **부모 타입으로 참조** 가능

```java
Animal a = new Cat();  // Cat은 Animal이다
a.sound();             // "야옹~" (오버라이딩 된 메서드 실행)
```

---

## 6️⃣ 업캐스팅 (Upcasting)

- 자식 객체를 **부모 타입 변수**로 참조
- 자동 변환됨 (묵시적 캐스팅)

```java
Dog dog = new Dog();
Animal a = dog;  // 업캐스팅
```

---

## 7️⃣ 다운캐스팅 (Downcasting)

- 부모 타입의 객체를 **자식 타입으로 변환**
- 명시적 캐스팅 필요

```java
Animal a = new Dog();
Dog d = (Dog) a;  // 다운캐스팅
```

> ⚠️ 타입이 실제로 자식 클래스가 아닐 경우 `ClassCastException` 발생 가능

---

## 8️⃣ instanceof 연산자

- 객체가 특정 클래스의 인스턴스인지 확인할 때 사용

```java
if (a instanceof Dog) {
    Dog d = (Dog) a;
    d.bark();
}
```

---

## 9️⃣ final 클래스/메서드

- `final` 키워드가 붙은 클래스는 상속 불가
- `final` 메서드는 오버라이딩 불가

```java
final class Utility { }     // 이 클래스는 상속 불가
final void doSomething() {} // 이 메서드는 재정의 불가
```

---

## 🔟 추상 클래스 (미리 보기)

- 상속을 위한 **설계용 클래스**
- 객체 생성 불가, 반드시 상속 후 구현 필요
- 다음 주제에서 자세히 다룰 예정

```java
abstract class Animal {
    abstract void sound(); // 구현은 자식에게 맡김
}
```

---

## 🎯 정리

✔ `extends` → 클래스를 상속할 때 사용  
✔ `super` → 부모 클래스의 필드나 메서드에 접근  
✔ `@Override` → 부모 메서드 재정의 시 명시  
✔ `업캐스팅` → 자식 → 부모 (자동)  
✔ `다운캐스팅` → 부모 → 자식 (명시적 형변환 필요)  
✔ `다형성` → 부모 타입 변수로 자식 객체를 다룰 수 있음  
✔ `instanceof` → 객체의 실제 타입을 체크  
✔ `final` → 클래스 상속/메서드 오버라이딩 제한  
✔ `오버라이딩` → 부모의 메서드를 자식이 재정의  
✔ `추상 클래스` → 직접 객체 생성 불가, 반드시 상속 필요

