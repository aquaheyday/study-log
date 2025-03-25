# 🔤 Java - 문자열 처리

Java에서 문자열은 **문자(char)의 집합**이며, `String` 클래스를 통해 다양한 기능을 제공합니다.  
문자열은 **불변(immutable)** 하며, 문자열 조작을 위한 메서드가 매우 풍부합니다.

---

## 1️⃣ 문자열 선언과 생성

```java
String str1 = "Hello";
String str2 = new String("Hello");  // 명시적 생성
```

- 문자열 리터럴은 **String constant pool**에 저장됨
- `new` 키워드는 **heap 메모리**에 별도로 할당됨

---

## 2️⃣ 문자열 비교

```java
String a = "apple";
String b = new String("apple");

System.out.println(a == b);            // false (주소 비교)
System.out.println(a.equals(b));       // true  (값 비교)
```

- `==` → 주소(참조값) 비교  
- `.equals()` → 내용(값) 비교  
- `.equalsIgnoreCase()` → 대소문자 무시 비교

---

## 3️⃣ 문자열 길이

```java
String text = "Hello World";
System.out.println(text.length());  // 11
```

---

## 4️⃣ 문자 추출

```java
String str = "Java";
char c = str.charAt(1);  // 'a'
```

---

## 5️⃣ 부분 문자열 추출

```java
String s = "Hello Java";
String sub = s.substring(6);     // "Java"
String sub2 = s.substring(0, 5); // "Hello"
```

---

## 6️⃣ 문자열 포함 여부

```java
String msg = "Spring Boot";
msg.contains("Boot");    // true
msg.startsWith("Spr");   // true
msg.endsWith("oot");     // true
```

---

## 7️⃣ 문자열 치환

```java
String msg = "Hello World";
String replaced = msg.replace("World", "Java"); // "Hello Java"
```

---

## 8️⃣ 문자열 분리 (split)

```java
String data = "apple,banana,grape";
String[] arr = data.split(",");

for (String item : arr) {
    System.out.println(item);
}
```

---

## 9️⃣ 문자열 합치기

```java
String first = "Hello";
String second = "Java";

String result = first + " " + second;               // "Hello Java"
String result2 = String.join("-", "a", "b", "c");   // "a-b-c"
```

✔ 문자열 연산이 많은 경우 `StringBuilder` 사용 권장 (성능↑)

---

## 🔟 StringBuilder / StringBuffer

- 문자열을 **변경 가능한 객체**로 다룰 수 있음  
- `StringBuilder` → 단일 스레드용  
- `StringBuffer` → 멀티 스레드 안전

```java
StringBuilder sb = new StringBuilder("Hello");
sb.append(" Java");
System.out.println(sb.toString());  // "Hello Java"
```

---

## 🎯 정리

✔ `String` → 불변 객체 (값 변경 시 새로운 객체 생성)  
✔ `==` → 주소 비교 / `.equals()` → 값 비교  
✔ `.length()` → 문자열 길이 반환  
✔ `.charAt()` → 특정 인덱스 문자 추출  
✔ `.substring()` → 부분 문자열 추출  
✔ `.contains()`, `.startsWith()`, `.endsWith()` → 포함 여부 확인  
✔ `.replace()` → 문자열 치환  
✔ `.split()` → 구분자로 문자열 나누기  
✔ `+`, `String.join()` → 문자열 결합  
✔ `StringBuilder` → 가변 문자열 처리, 반복 결합에 적합

