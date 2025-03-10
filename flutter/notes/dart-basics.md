# Flutter 기본 문법 정리

## 1. Flutter 개요
Flutter는 Google에서 개발한 오픈소스 UI 프레임워크로, 하나의 코드베이스로 Android, iOS, 웹, 데스크톱 애플리케이션을 개발할 수 있습니다.

---

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

---

## 3. Flutter 기본 위젯
Flutter에서 UI는 위젯(Widget)으로 구성됩니다. 위젯은 화면을 구성하는 기본 단위이며, 모든 요소(텍스트, 버튼, 레이아웃 등)가 위젯입니다.

위젯은 크게 두 가지로 나뉩니다.
- **StatelessWidget**: 상태를 가지지 않는 위젯
- **StatefulWidget**: 상태를 가질 수 있는 위젯

### StatelessWidget
`StatelessWidget`은 변경되지 않는 UI를 구성할 때 사용됩니다.

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
        appBar: AppBar(title: Text('StatelessWidget 예제')),
        body: Center(child: Text('Hello, Flutter!')),
      ),
    );
  }
}
```

### StatefulWidget
StatefulWidget은 변경 가능한 상태를 가지며, setState()를 호출하여 UI를 업데이트할 수 있습니다.

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
      appBar: AppBar(title: Text('StatefulWidget 예제')),
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
텍스트를 화면에 표시하는 위젯입니다.
```dart
Text(
  'Hello, Flutter!',
  style: TextStyle(fontSize: 24, fontWeight: FontWeight.bold),
)
```

#### Cotainer
컨테이너는 박스 형태의 위젯으로, 크기, 배경색, 패딩 등을 설정할 수 있습니다.
```dart
Container(
  width: 200,
  height: 100,
  color: Colors.blue,
  child: Center(child: Text('Container Example')),
)
```

#### Colunm & Row
위젯을 수직(Column)으로 배치하는 레이아웃 위젯입니다.
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
위젯을 수평(Row)으로 배치하는 레이아웃 위젯입니다.
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
클릭할 수 있는 버튼을 생성합니다.
```dart
ElevatedButton(
  onPressed: () {
    print('Button Pressed');
  },
  child: Text('Click Me'),
)
```

#### Image
이미지를 표시하는 위젯입니다.
```dart
Image.network(
  'https://flutter.dev/images/flutter-logo-sharing.png',
  width: 100,
  height: 100,
)
```
```dart
Image.asset(
  'assets/my_image.png',
  width: 100,
  height: 100,
)
```


### 사용자 입력 위젯
#### TextField
사용자가 텍스트를 입력할 수 있도록 하는 위젯입니다.
```dart
TextField(
  decoration: InputDecoration(
    labelText: 'Enter your name',
    border: OutlineInputBorder(),
  ),
)
```

#### Checkbox
체크박스를 생성합니다.
```dart
Checkbox(
  value: true,
  onChanged: (bool? newValue) {
    print(newValue);
  },
)
```

#### Switch
스위치를 생성합니다.
```dart
Switch(
  value: true,
  onChanged: (bool newValue) {
    print(newValue);
  },
)
```

#### Slider
슬라이더를 생성합니다.
```dart
Slider(
  value: 0.5,
  min: 0.0,
  max: 1.0,
  onChanged: (double newValue) {
    print(newValue);
  },
)
```

---

### 4. 상태 관리 (StatefulWidget)
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

---

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

---

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
