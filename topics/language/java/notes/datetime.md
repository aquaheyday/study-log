# 📅 Java - 날짜와 시간

Java에서 날짜와 시간 처리는 주로 `java.time` 패키지를 사용합니다.  
Java 8부터는 `LocalDate`, `LocalDateTime`, `ZonedDateTime` 등으로 **불변 객체 기반**의 안전한 날짜/시간 API가 제공됩니다.

---

## 1️⃣ 현재 날짜와 시간 얻기

```java
import java.time.*;

LocalDate date = LocalDate.now();                 // 현재 날짜 (연, 월, 일)
LocalTime time = LocalTime.now();                 // 현재 시간 (시, 분, 초)
LocalDateTime dateTime = LocalDateTime.now();     // 날짜 + 시간
ZonedDateTime zoned = ZonedDateTime.now();        // 시간대 포함
```

---

## 2️⃣ 날짜/시간 생성

```java
LocalDate date = LocalDate.of(2024, 3, 25);
LocalTime time = LocalTime.of(14, 30, 0);
LocalDateTime dt = LocalDateTime.of(date, time);
```

---

## 3️⃣ 날짜/시간 정보 추출

```java
LocalDate today = LocalDate.now();

int year = today.getYear();
int month = today.getMonthValue();     // 1 ~ 12
DayOfWeek dow = today.getDayOfWeek();  // MONDAY ~ SUNDAY
```

---

## 4️⃣ 날짜/시간 조작 (plus, minus)

```java
LocalDate date = LocalDate.now();

date.plusDays(3);       // 3일 뒤
date.minusMonths(1);    // 1달 전
date.plusYears(2);      // 2년 뒤
```

---

## 5️⃣ 날짜 비교

```java
LocalDate a = LocalDate.of(2024, 1, 1);
LocalDate b = LocalDate.now();

a.isBefore(b);  // true/false
a.isAfter(b);
a.equals(b);
```

---

## 6️⃣ 날짜 포맷 변환

```java
import java.time.format.DateTimeFormatter;

LocalDateTime now = LocalDateTime.now();
String formatted = now.format(DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss"));
```

```java
// 출력: 2025-03-25 14:30:00
```

---

## 7️⃣ 문자열 → 날짜 파싱

```java
String str = "2025-03-25 14:00:00";

LocalDateTime dt = LocalDateTime.parse(
    str,
    DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss")
);
```

---

## 8️⃣ 두 날짜 간 차이 계산

```java
import java.time.temporal.ChronoUnit;

LocalDate start = LocalDate.of(2024, 1, 1);
LocalDate end = LocalDate.now();

long days = ChronoUnit.DAYS.between(start, end);   // 며칠 차이
long months = ChronoUnit.MONTHS.between(start, end);
```

---

## 9️⃣ 시간대 (ZoneId)

```java
ZonedDateTime nowInKorea = ZonedDateTime.now(ZoneId.of("Asia/Seoul"));
ZonedDateTime nowInNY = ZonedDateTime.now(ZoneId.of("America/New_York"));
```

---

## 🔟 이전 방식 (java.util.Date / Calendar)

```java
Date date = new Date();           // 현재 날짜
Calendar cal = Calendar.getInstance();
cal.set(2024, Calendar.MARCH, 25);
Date d = cal.getTime();
```


✔ 현재는 `java.time` 사용을 **강력히 권장**합니다 (더 안전하고 간결)

---

## 🎯 정리

✔ `LocalDate`, `LocalTime`, `LocalDateTime` → 날짜, 시간, 날짜+시간 처리  
✔ `now()`, `of()` → 현재 또는 특정 날짜 생성  
✔ `plusX()`, `minusX()` → 날짜 연산  
✔ `isBefore()`, `isAfter()` → 날짜 비교  
✔ `format()` → 날짜 → 문자열  
✔ `parse()` → 문자열 → 날짜  
✔ `ChronoUnit` → 날짜 간 차이 계산  
✔ `ZoneId`, `ZonedDateTime` → 시간대 지원  
✔ 이전 방식인 `Date`, `Calendar`는 **호환용**으로만 사용  
✔ `java.time` 패키지는 **Immutable + Thread-safe**

