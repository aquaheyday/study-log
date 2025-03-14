# 📌 Dart 기본 문법

## 1. Dart란 무엇인가?

**Dart**는 Google에서 개발한 프로그래밍 언어로, **Flutter의 공식 언어**입니다.  
Dart는 **객체 지향 언어**이며, **정적 및 동적 타입 지정이 가능**하고, **JIT(Just-In-Time) 및 AOT(Ahead-Of-Time) 컴파일**을 지원하여 빠른 실행 속도를 제공합니다.

✔ Dart의 특징
- **C 계열 문법**을 기반으로 한 쉽고 직관적인 문법
- **JIT & AOT 컴파일**을 활용한 빠른 실행 및 성능 최적화
- **Garbage Collection 지원**으로 메모리 관리 자동화
- **비동기 프로그래밍 지원** (`async/await`, `Future`, `Stream`)
- **모바일, 웹, 서버 및 데스크톱 애플리케이션 개발 가능**

---

## 2. 기본 문법

### 변수 선언

Dart에서는 `var`, `final`, `const`, `dynamic` 키워드를 사용하여 변수를 선언할 수 있습니다.

```dart
void main() {
  var name = "Flutter";  // 타입 추론 (String)
  String language = "Dart";  // 명시적 선언
  final String version = "3.0";  // 런타임 상수
  const double pi = 3.1415;  // 컴파일 타임 상수

  dynamic variable = 42;  // 타입 변경 가능
  variable = "Hello";  // 문제 없음

  print(name);
  print(language);
  print(version);
  print(pi);
  print(variable);
}
```

✔ **차이점**  
- `final`: 한 번 할당하면 변경할 수 없는 **런타임 상수**  
- `const`: **컴파일 타임 상수**, `final`보다 더 엄격  
- `dynamic`: 런타임에 타입이 결정되며 변경 가능  

---

### 데이터 타입

| 타입 | 설명 | 예제 |
|---|---|---|
| **int** | 정수형 | `int num = 10;` |
| **double** | 실수형 | `double pi = 3.14;` |
| **String** | 문자열 | `String name = "Flutter";` |
| **bool** | 참/거짓 | `bool isFlutter = true;` |
| **List** | 리스트(배열) | `List<int> numbers = [1, 2, 3];` |
| **Set** | 중복 없는 집합 | `Set<String> fruits = {"apple", "banana"};` |
| **Map** | 키-값 쌍 | `Map<String, int> ages = {"Alice": 25, "Bob": 30};` |
| **dynamic** | 동적 타입 | `dynamic value = "Hello"; value = 42;` |

---

### 예제: 데이터 타입

```dart
void main() {
  int age = 30;
  double height = 1.75;
  String message = "Hello, Dart!";
  bool isLearning = true;

  List<String> languages = ["Dart", "Flutter", "JavaScript"];
  Set<int> uniqueNumbers = {1, 2, 3, 4, 4};  // {1, 2, 3, 4}
  Map<String, int> scores = {"Alice": 95, "Bob": 85};

  print("$message I'm $age years old.");
  print("Height: $height m");
  print("Languages: $languages");
  print("Unique Numbers: $uniqueNumbers");
  print("Scores: $scores");
}
```

---

### 연산자

```dart
void main() {
  int a = 10;
  int b = 3;

  print(a + b);  // 덧셈
  print(a - b);  // 뺄셈
  print(a * b);  // 곱셈
  print(a / b);  // 나눗셈 (소수점 포함)
  print(a ~/ b); // 나눗셈 (정수형 결과)
  print(a % b);  // 나머지

  // 비교 연산자
  print(a > b);  // true
  print(a < b);  // false
  print(a == b); // false
  print(a != b); // true

  // 논리 연산자
  bool isFlutter = true;
  bool isDart = false;

  print(isFlutter && isDart); // false
  print(isFlutter || isDart); // true
  print(!isFlutter);          // false
}
```

---

## 3. 제어문

### 조건문 (if-else, switch-case)
```dart
void main() {
  int score = 85;

  if (score >= 90) {
    print("A 학점");
  } else if (score >= 80) {
    print("B 학점");
  } else {
    print("C 학점");
  }

  String grade = "B";

  switch (grade) {
    case "A":
      print("Excellent!");
      break;
    case "B":
      print("Good job!");
      break;
    default:
      print("Keep trying!");
  }
}
```

---

### 반복문 (for, while, do-while)
```dart
void main() {
  // for 문
  for (int i = 1; i <= 5; i++) {
    print("Number: $i");
  }

  // while 문
  int j = 1;
  while (j <= 5) {
    print("While: $j");
    j++;
  }

  // do-while 문
  int k = 1;
  do {
    print("Do-While: $k");
    k++;
  } while (k <= 5);
}
```

---

## 4. 함수

### 기본 함수 선언
```dart
void greet(String name) {
  print("Hello, $name!");
}

int add(int a, int b) {
  return a + b;
}

void main() {
  greet("Dart");
  int result = add(3, 5);
  print("Sum: $result");
}
```

### `Lambda` (화살표 함수)
```dart
int multiply(int x, int y) => x * y;

void main() {
  print(multiply(4, 5)); // 20
}
```

### 선택적 매개변수
```dart
void printInfo(String name, {int? age}) {
  print("Name: $name, Age: ${age ?? 'Unknown'}");
}

void main() {
  printInfo("Alice");
  printInfo("Bob", age: 25);
}
```

---

## 5. 비동기 프로그래밍 (`Async/Await`, `Future`)

### `Future` (비동기 작업)
```dart
Future<String> fetchData() async {
  await Future.delayed(Duration(seconds: 2));  // 2초 대기
  return "Data Loaded";
}

void main() async {
  print("Fetching...");
  String data = await fetchData();
  print(data);
}
```

---

## 🎯 정리
✔ Dart는 객체 지향 및 비동기 프로그래밍을 지원하는 언어  
✔ 변수 선언 시 `var`, `final`, `const`, `dynamic` 사용  
✔ 리스트, 맵, 세트 등의 컬렉션 타입 제공  
✔ 제어문 (if, switch, for, while) 활용 가능  
✔ 비동기 프로그래밍을 위해 `Future`와 `async/await` 지원  
