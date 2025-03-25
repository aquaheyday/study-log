# 📦 Java - 배열 (Array)

배열은 **같은 타입의 데이터를 순차적으로 저장**할 수 있는 자료구조입니다.  
Java에서는 배열도 **객체**이며, 크기가 고정되고 인덱스(0부터 시작)로 접근합니다.

---

## 1️⃣ 배열 선언과 생성

```java
int[] scores = new int[5];     // 길이 5짜리 int 배열
String[] names = new String[3];
```

또는

```java
int scores[] = new int[5];     // 가능은 하지만 권장 X (가독성 ↓)
```

---

## 2️⃣ 배열 초기화 (리터럴 방식)

```java
int[] nums = {10, 20, 30, 40};
String[] colors = {"red", "green", "blue"};
```

✔ 선언과 동시에 값 지정 가능

---

## 3️⃣ 배열 요소 접근 및 변경

```java
int[] arr = {1, 2, 3};

System.out.println(arr[0]);  // 1
arr[1] = 100;                // 두 번째 값을 100으로 변경
```

---

## 4️⃣ 배열 길이 확인

```java
int[] arr = {10, 20, 30};

System.out.println(arr.length);  // 3
```

✔ `.length`는 필드이며, 괄호 없이 사용

---

## 5️⃣ 배열 순회 (for문)

```java
int[] nums = {1, 2, 3, 4, 5};

for (int i = 0; i < nums.length; i++) {
    System.out.println(nums[i]);
}
```

---

## 6️⃣ 향상된 for문 (for-each)

```java
String[] animals = {"cat", "dog", "fox"};

for (String animal : animals) {
    System.out.println(animal);
}
```

✔ 단, 인덱스가 필요한 작업에는 일반 for문 사용

---

## 7️⃣ 다차원 배열

```java
int[][] matrix = {
    {1, 2},
    {3, 4},
    {5, 6}
};

System.out.println(matrix[1][1]);  // 4
```

✔ 2차원 배열 이상도 가능 (표처럼 사용)

---

## 8️⃣ 배열 복사

```java
int[] original = {1, 2, 3};
int[] copy = Arrays.copyOf(original, original.length);
```

또는:

```java
System.arraycopy(original, 0, copy, 0, original.length);
```

✔ `import java.util.Arrays;` 필요

---

## 9️⃣ 배열 정렬

```java
int[] nums = {5, 1, 3, 2};
Arrays.sort(nums);  // 오름차순 정렬

System.out.println(Arrays.toString(nums));  // [1, 2, 3, 5]
```

---

## 🔟 배열 출력

```java
int[] data = {1, 2, 3};

System.out.println(Arrays.toString(data));  // [1, 2, 3]
```

✔ `System.out.println(data);`는 메모리 주소 출력이므로 주의!

---

## 🎯 정리

✔ `int[] arr = new int[5];` → 배열 선언 및 생성  
✔ `{...}` → 배열 리터럴 초기화  
✔ `arr.length` → 배열의 길이  
✔ `arr[i]` → 배열 요소 접근  
✔ `for`, `for-each` → 배열 반복  
✔ `int[][] arr` → 다차원 배열 선언  
✔ `Arrays.copyOf()` / `System.arraycopy()` → 배열 복사  
✔ `Arrays.sort()` → 배열 정렬  
✔ `Arrays.toString()` → 배열 내용 출력  
✔ 배열은 **크기 고정 + 같은 타입**만 저장 가능

