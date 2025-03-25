# 🧱 Java - 클래스와 객체

Java는 객체지향 언어로, 클래스와 객체는 핵심 개념입니다.  
**클래스(Class)** 는 설계도, **객체(Object)** 는 그 설계도로 만들어진 실체입니다.

---

## 1️⃣ 클래스 정의

```java
public class Car {
    // 필드 (속성)
    String brand;
    int speed;

    // 메서드 (동작)
    void drive() {
        System.out.println(brand + " is driving at " + speed + " km/h");
    }
}
```

---

## 2️⃣ 객체 생성
- `new` 키워드로 객체(인스턴스) 생성
- 점(.) 연산자로 필드나 메서드에 접근

```java
Car myCar = new Car();
myCar.brand = "Hyundai";
myCar.speed = 100;
myCar.drive();  // 출력: Hyundai is driving at 100 km/h
```

---

## 3️⃣ 생성자(Constructor)

- 객체 생성 시 자동 호출되는 특별한 메서드  
- 클래스 이름과 같으며, 반환 타입이 없음

```java
public class Person {
    String name;

    // 생성자
    public Person(String name) {
        this.name = name;
    }
}
```

```java
Person p = new Person("Alice");
```

---

## 4️⃣ this 키워드

현재 객체 자신을 참조할 때 사용

```java
public class Dog {
    String name;

    public Dog(String name) {
        this.name = name;  // this로 인스턴스 필드와 구분
    }
}
```

---

## 5️⃣ 메서드 오버로딩 (Method Overloading)

같은 이름의 메서드를 **매개변수만 다르게** 여러 개 정의 가능

```java
public class MathUtil {
    int add(int a, int b) {
        return a + b;
    }

    double add(double a, double b) {
        return a + b;
    }
}
```

---

## 6️⃣ 접근 제어자 (Access Modifiers)

| 키워드 | 설명 |
|--------|------|
| `public` | 모든 클래스에서 접근 가능 |
| `private` | 현재 클래스 내에서만 접근 가능 |
| `protected` | 같은 패키지 또는 자식 클래스에서 접근 가능 |
| (default) | 같은 패키지 내에서만 접근 가능 (아무 키워드도 없을 때) |

---

## 7️⃣ getter / setter 메서드

`private` 필드를 외부에서 읽거나 수정할 수 있도록 제공하는 메서드

```java
public class User {
    private String name;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }
}
```

---

## 8️⃣ static 키워드

- 클래스 단위로 존재하는 필드/메서드에 사용
- 객체 없이 접근 가능

```java
public class MathUtil {
    static double PI = 3.14;

    static int square(int x) {
        return x * x;
    }
}

MathUtil.square(3);  // 객체 없이 사용 가능
```

---

## 9️⃣ 객체 배열

같은 타입의 여러 객체를 배열로 관리

```java
Student[] students = new Student[3];

students[0] = new Student("Tom");
students[1] = new Student("Jane");
students[2] = new Student("Alex");
```

---

## 🔟 객체 비교 (== vs equals)

```java
String a = new String("Hello");
String b = new String("Hello");

System.out.println(a == b);        // false (주소 비교)
System.out.println(a.equals(b));   // true (내용 비교)
```

---

## 🎯 정리

✔ `class` → 객체를 생성하기 위한 설계도  
✔ `object` → 클래스로부터 생성된 실체 (인스턴스)  
✔ `new` → 객체 생성 연산자  
✔ `constructor` → 객체 생성 시 호출되는 초기화 메서드  
✔ `this` → 현재 인스턴스를 참조하는 키워드  
✔ `method overloading` → 같은 이름의 메서드를 매개변수 다르게 여러 개 정의  
✔ `access modifier` → 접근 범위를 제한 (public, private 등)  
✔ `getter/setter` → 캡슐화된 필드 접근용 메서드  
✔ `static` → 클래스 단위로 공유되는 필드/메서드  
✔ `equals()` → 객체의 **내용** 비교 / `==` → 객체의 **참조 주소** 비교
