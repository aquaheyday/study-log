# 🛡️ Java - 예외 처리 (Exception Handling)

예외(Exception)는 프로그램 실행 중 발생할 수 있는 **예상치 못한 오류 상황**을 의미합니다.  
Java는 예외를 **객체로 다루며**, try-catch 문을 통해 안전하게 처리합니다.

---

## 1️⃣ 예외란?

- 실행 중 발생하는 **비정상적인 상황**
- 예: 0으로 나누기, 배열 범위 초과, null 접근 등

```java
int[] arr = {1, 2, 3};
System.out.println(arr[5]);  // ArrayIndexOutOfBoundsException
```

---

## 2️⃣ 예외 처리 기본 구조

```java
try {
    // 예외 발생 가능 코드
} catch (ExceptionType e) {
    // 예외 발생 시 처리 코드
} finally {
    // (선택) 항상 실행되는 코드
}
```

---

## 3️⃣ 예외 클래스 계층 구조

- 모든 예외 클래스는 `Throwable`을 상속합니다.

```
Throwable
 ├── Error (심각한 시스템 오류)
 └── Exception
      ├── Checked Exception
      └── Unchecked Exception (RuntimeException 하위)
```

---

## 4️⃣ Checked vs Unchecked 예외

| 구분 | 예시 | 처리 필요 여부 |
|------|------|----------------|
| Checked | IOException, SQLException | 반드시 처리 필요 |
| Unchecked | NullPointerException, IndexOutOfBoundsException | 처리 선택 가능 |

---

## 5️⃣ 다중 catch

- 여러 예외 유형을 구분해서 처리할 수 있음

```java
try {
    String s = null;
    s.length();
} catch (NullPointerException e) {
    System.out.println("Null 접근 예외");
} catch (Exception e) {
    System.out.println("기타 예외");
}
```

---

## 6️⃣ finally 블록

- 예외 발생 여부와 관계없이 **항상 실행**되는 블록

```java
try {
    // DB 연결
} catch (Exception e) {
    // 예외 처리
} finally {
    // 연결 해제
}
```

---

## 7️⃣ throws 키워드

- 메서드에서 예외를 **직접 처리하지 않고 위임**할 때 사용

```java
public void readFile(String path) throws IOException {
    FileReader fr = new FileReader(path);
}
```

---

## 8️⃣ throw 키워드

- 예외 객체를 **직접 발생**시킬 때 사용

```java
if (age < 0) {
    throw new IllegalArgumentException("나이는 0 이상이어야 합니다.");
}
```

---

## 9️⃣ 사용자 정의 예외

- 직접 예외 클래스를 만들어 사용할 수 있음

```java
public class CustomException extends Exception {
    public CustomException(String msg) {
        super(msg);
    }
}
```

```java
throw new CustomException("사용자 정의 예외 발생");
```

---

## 🔟 try-with-resources (자동 자원 해제)

- `AutoCloseable`을 구현한 객체는 try문에서 자동으로 close됩니다.

```java
try (BufferedReader br = new BufferedReader(new FileReader("file.txt"))) {
    System.out.println(br.readLine());
} catch (IOException e) {
    e.printStackTrace();
}
```

---

## 🎯 정리

✔ `try-catch` → 예외 발생 시 실행 흐름을 유지하며 처리  
✔ `finally` → 예외 유무와 관계없이 항상 실행  
✔ `throw` → 예외를 직접 발생시킴  
✔ `throws` → 예외를 메서드 호출자에게 전가  
✔ `Checked Exception` → 반드시 예외 처리해야 함 (컴파일 에러)  
✔ `Unchecked Exception` → 선택적으로 처리 가능  
✔ `Exception` → 대부분의 예외가 상속하는 클래스  
✔ `RuntimeException` → 실행 중에 발생하는 예외  
✔ `CustomException` → 사용자 정의 예외 클래스  
✔ `try-with-resources` → 파일/DB 등 자원 자동 해제

