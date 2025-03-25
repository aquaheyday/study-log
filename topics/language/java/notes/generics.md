# 🧬 Java - 제네릭 (Generics)

**제네릭(Generic)** 은 클래스나 메서드에서 **타입을 파라미터화**해서 **타입 안정성**을 높이고 **코드 재사용성**을 높이기 위한 기능입니다.

---

## 1️⃣ 제네릭이 필요한 이유

```java
List list = new ArrayList();
list.add("hello");
list.add(123);  // ❗ 문제 발생 가능

String str = (String) list.get(0);  // 형변환 필요
```

✔ 타입이 명확하지 않아서 **형변환(casting)**이 필요하고, 런타임 에러 발생 가능성도 높아짐 ⚠️

---

## 2️⃣ 제네릭 사용 기본 예제

```java
List<String> list = new ArrayList<>();
list.add("hello");
// list.add(123);  // 컴파일 에러 (타입 안전성 보장)

String str = list.get(0);  // 형변환 필요 없음
```

---

## 3️⃣ 제네릭 클래스 정의

```java
public class Box<T> {
    private T item;

    public void setItem(T item) {
        this.item = item;
    }

    public T getItem() {
        return item;
    }
}
```

```java
Box<String> stringBox = new Box<>();
stringBox.setItem("Hello");

Box<Integer> intBox = new Box<>();
intBox.setItem(100);
```

---

## 4️⃣ 제네릭 메서드 정의

```java
public class Util {
    public static <T> T pickFirst(T[] array) {
        return array[0];
    }
}
```

```java
String[] names = {"Alice", "Bob"};
String name = Util.pickFirst(names);  // T는 String으로 추론됨
```

---

## 5️⃣ 제네릭 타입 제한 (Bounded Type)

```java
public class Calculator<T extends Number> {
    public double doubleValue(T t) {
        return t.doubleValue();
    }
}
```

✔ `T extends Number` → T는 Number 또는 그 하위 타입만 가능  
✔ `extends`는 클래스/인터페이스 상속, 구현 모두 포함  

---

## 6️⃣ 와일드카드 `<?>` (Unbounded)

- **모든 타입 허용**
- 읽기 전용 → `list.add(...)`는 불가

```java
public void printList(List<?> list) {
    for (Object obj : list) {
        System.out.println(obj);
    }
}
```

---

## 7️⃣ 와일드카드 상한 제한 `<? extends T>`

- T 또는 T의 하위 타입만 허용
- 읽기 가능, 쓰기 제한됨

```java
public void printNumbers(List<? extends Number> list) {
    for (Number n : list) {
        System.out.println(n);
    }
}
```

---

## 8️⃣ 와일드카드 하한 제한 `<? super T>`

- T 또는 T의 **상위 타입**만 허용
- 쓰기는 가능하지만, 읽을 때는 `Object`로 받아야 함

```java
public void addIntegers(List<? super Integer> list) {
    list.add(1);
    list.add(2);
}
```

---

## 9️⃣ 제네릭 배열은 불가

```java
// ❌ 불가능
// T[] arr = new T[10];

// ✅ 해결 방법: Object 배열 생성 후 캐스팅
@SuppressWarnings("unchecked")
T[] arr = (T[]) new Object[10];
```

---

## 🔟 타입 소거(Type Erasure)

- 컴파일 후에는 **제네릭 정보가 제거됨**
- JVM에서는 `List<String>`과 `List<Integer>`는 동일한 타입으로 취급

```java
List<String> list1 = new ArrayList<>();
List<Integer> list2 = new ArrayList<>();

System.out.println(list1.getClass() == list2.getClass());  // true
```

---

## 🎯 정리

✔ `<>` → 클래스나 메서드에서 타입을 파라미터화  
✔ `List<String>` → String만 저장 가능, 형변환 불필요  
✔ `<T>` → 제네릭 타입 선언  
✔ `<T extends Number>` → 상위 타입 제한  
✔ `<?>` → 모든 타입 허용 (읽기 전용)  
✔ `<? extends T>` → T 또는 하위 타입 (읽기 가능, 쓰기 제한)  
✔ `<? super T>` → T 또는 상위 타입 (쓰기 가능, 읽기 제한)  
✔ `Object[]` + `(T[])` 캐스팅 → 제네릭 배열 우회 생성  
✔ **타입 소거(Type Erasure)** → 런타임에 제네릭 정보 없음  
✔ 제네릭은 **타입 안전성 확보 + 재사용성 향상**

