# 🔁 Java - 제어문 (조건문 & 반복문)

제어문은 프로그램의 흐름을 제어하는 구문으로, **조건 분기(if, switch)** 와 **반복 처리(for, while)** 를 포함합니다.

---

## 1️⃣ if 조건문

조건에 따라 실행 흐름을 제어하는 가장 기본적인 조건문입니다.

```java
int age = 20;

if (age >= 18) {
    System.out.println("성인입니다.");
} else {
    System.out.println("미성년자입니다.");
}
```

---

## 2️⃣ if - else if - else

여러 조건을 순차적으로 검사할 수 있습니다.

```java
int score = 85;

if (score >= 90) {
    System.out.println("A학점");
} else if (score >= 80) {
    System.out.println("B학점");
} else {
    System.out.println("C학점 이하");
}
```

---

## 3️⃣ 중첩 if

조건문 안에 또 다른 조건문을 포함할 수 있습니다.

```java
boolean isMember = true;
int age = 25;

if (isMember) {
    if (age >= 20) {
        System.out.println("성인 회원");
    } else {
        System.out.println("청소년 회원");
    }
}
```

---

## 4️⃣ switch문

값에 따라 여러 분기를 처리할 때 사용합니다.

```java
int day = 3;

switch (day) {
    case 1:
        System.out.println("월요일");
        break;
    case 2:
        System.out.println("화요일");
        break;
    case 3:
        System.out.println("수요일");
        break;
    default:
        System.out.println("기타 요일");
}
```

> `break` 를 사용하지 않으면 **fall-through** 현상 발생 주의

---

## 5️⃣ while 반복문

조건이 **참인 동안** 계속 반복합니다.

```java
int i = 1;

while (i <= 5) {
    System.out.println("i = " + i);
    i++;
}
```

---

## 6️⃣ do-while 반복문

**최소 1회는 실행**되는 반복문입니다.

```java
int i = 1;

do {
    System.out.println("i = " + i);
    i++;
} while (i <= 5);
```

---

## 7️⃣ for 반복문

초기값, 조건, 증감식이 모두 포함된 반복문입니다.

```java
for (int i = 0; i < 5; i++) {
    System.out.println("i = " + i);
}
```

---

## 8️⃣ 향상된 for문 (foreach)

배열이나 컬렉션을 순회할 때 간결하게 사용합니다.

```java
String[] fruits = {"사과", "바나나", "포도"};

for (String fruit : fruits) {
    System.out.println(fruit);
}
```

---

## 9️⃣ break 문

반복문을 **즉시 종료**합니다.

```java
for (int i = 1; i <= 10; i++) {
    if (i == 5) break;
    System.out.println(i);
}
```

---

## 🔟 continue 문

이번 반복만 **건너뛰고**, 다음 반복으로 진행합니다.

```java
for (int i = 1; i <= 5; i++) {
    if (i == 3) continue;
    System.out.println(i); // 3은 출력되지 않음
}
```

---

## 🎯 정리

✔ `if` → 조건이 참일 때 실행  
✔ `else if` → 여러 조건 중 하나 선택  
✔ `switch` → 값에 따라 분기 처리  
✔ `while` → 조건이 참이면 반복 (0회 가능)  
✔ `do-while` → 조건과 관계없이 최소 1회 실행  
✔ `for` → 초기화 + 조건 + 증감 포함된 반복  
✔ `향상된 for` → 배열/컬렉션 순회에 적합  
✔ `break` → 반복문 즉시 종료  
✔ `continue` → 현재 반복만 건너뜀  
✔ `중첩 if` → 조건문 안에 또 다른 조건문

