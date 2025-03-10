# Flutter 기본 문법 정리

## 1. Flutter 개요
Flutter는 Google에서 개발한 오픈소스 UI 프레임워크로, 하나의 코드베이스로 Android, iOS, 웹, 데스크톱 애플리케이션을 개발할 수 있습니다.

## 2. Dart 기본 문법
Flutter는 Dart 언어를 사용합니다.

### 변수 선언
```dart
void main() {
  var name = 'Flutter';  // 타입 추론
  String language = 'Dart'; // 명시적 선언
  int version = 3;
  double pi = 3.14;
  bool isFlutterAwesome = true;

  print(name);
}
```

### 조건문
```dart
void main() {
  int number = 10;

  if (number > 5) {
    print('Number is greater than 5');
  } else {
    print('Number is 5 or less');
  }
}
```

### 반복문
```dart
void main() {
  for (int i = 0; i < 5; i++) {
    print('Iteration: \$i');
  }

  int count = 0;
  while (count < 5) {
    print('While loop: \$count');
    count++;
  }
}
```

### 함수 선언
```dart
int add(int a, int b) {
  return a + b;
}

void main() {
  int result = add(5, 10);
  print('Sum: \$result');
}
```

## 3. Flutter 기본 위젯
Flutter에서 UI는 위젯을 통해 구성됩니다.

### 기본 앱 구조
```dart
import 'package:flutter/material.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(title: Text('Flutter 기본 문법')),
        body: Center(child: Text('Hello, Flutter!')),
      ),
    );
  }
}
```

### 주요 위젯

#### Text
```dart
Text(
  'Hello, Flutter!',
  style: TextStyle(fontSize: 24, fontWeight: FontWeight.bold),
)
```

#### Container
```dart
Container(
  width: 200,
  height: 100,
  color: Colors.blue,
  child: Center(child: Text('Container Example')),
)
```

#### Column
```dart
Column(
  mainAxisAlignment: MainAxisAlignment.center,
  children: [
    Text('First Line'),
    Text('Second Line'),
  ],
)
```

#### Row
```dart
Row(
  mainAxisAlignment: MainAxisAlignment.spaceAround,
  children: [
    Icon(Icons.star),
    Icon(Icons.favorite),
    Icon(Icons.thumb_up),
  ],
)
```

#### ElevatedButton
```dart
ElevatedButton(
  onPressed: () {
    print('Button Pressed');
  },
  child: Text('Click Me'),
)
```

## 4. 상태 관리 (StatefulWidget)
Flutter는 StatelessWidget과 StatefulWidget을 사용하여 UI를 구성합니다.

```dart
import 'package:flutter/material.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: CounterScreen(),
    );
  }
}

class CounterScreen extends StatefulWidget {
  @override
  _CounterScreenState createState() => _CounterScreenState();
}

class _CounterScreenState extends State<CounterScreen> {
  int _counter = 0;

  void _incrementCounter() {
    setState(() {
      _counter++;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text('Counter Example')),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text('Counter: \$_counter', style: TextStyle(fontSize: 24)),
            ElevatedButton(
              onPressed: _incrementCounter,
              child: Text('Increase'),
            ),
          ],
        ),
      ),
    );
  }
}
```

## 5. 네비게이션 (페이지 이동)
```dart
Navigator.push(
  context,
  MaterialPageRoute(builder: (context) => SecondPage()),
);
```
```dart
Navigator.pop(context);
```

## 6. API 호출 (HTTP 요청)
```dart
import 'package:http/http.dart' as http;
import 'dart:convert';

void fetchData() async {
  final response = await http.get(Uri.parse('https://jsonplaceholder.typicode.com/posts/1'));
  
  if (response.statusCode == 200) {
    var data = jsonDecode(response.body);
    print('Title: ' + data['title']);
  } else {
    print('Failed to load data');
  }
}
```
