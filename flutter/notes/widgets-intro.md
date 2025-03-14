# 위젯 개요 (Flutter Widgets)

Flutter에서 **위젯(Widget)** 은 UI를 구성하는 기본 요소입니다.  
Flutter 앱의 모든 것은 **위젯으로 이루어져 있으며, 위젯을 조합하여 화면을 구성**합니다.

---

## 1. 위젯이란?

**위젯(Widget)**은 Flutter에서 화면을 구성하는 **UI 요소**입니다.  
모든 버튼, 텍스트, 이미지, 레이아웃 등이 **위젯**으로 표현됩니다.

✔ **Flutter UI = 위젯 트리 (Widget Tree)**
- 위젯을 중첩하여 트리 구조로 화면을 구성
- 부모-자식 관계를 형성하여 UI 계층 구조 생성

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
        appBar: AppBar(title: Text("Flutter Widgets")),
        body: Center(child: Text("Hello, Flutter!")),
      ),
    );
  }
}
```
📌 `MaterialApp`, `Scaffold`, `AppBar`, `Text` 등 모든 요소가 위젯!

---

## 2. 위젯의 종류

Flutter 위젯은 크게 두 가지로 나뉩니다.

| 종류 | 설명 | 예제 |
|------|------|------|
| **Stateless Widget** | 상태가 없는 정적 UI | `Text`, `Icon`, `Container` |
| **Stateful Widget** | 상태(State)를 가지며 동적 변경 가능 | `TextField`, `Checkbox`, `AnimatedContainer` |

---

## 3. Stateless 위젯 (StatelessWidget)

**StatelessWidget**은 상태가 변경되지 않는 **정적인 UI**를 구성할 때 사용됩니다.

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
        appBar: AppBar(title: Text("Stateless Widget")),
        body: Center(child: MyTextWidget()),
      ),
    );
  }
}

class MyTextWidget extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Text(
      "This is a Stateless Widget!",
      style: TextStyle(fontSize: 20, color: Colors.blue),
    );
  }
}
```

✔ `StatelessWidget`은 `build` 메서드에서 UI를 반환하며, 상태 변경이 필요하지 않음.

---

## 4. Stateful 위젯 (StatefulWidget)

**StatefulWidget**은 상태(State)를 가지며, 변경될 수 있는 **동적인 UI**를 구성할 때 사용됩니다.

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
        appBar: AppBar(title: Text("Stateful Widget")),
        body: CounterWidget(),
      ),
    );
  }
}

class CounterWidget extends StatefulWidget {
  @override
  _CounterWidgetState createState() => _CounterWidgetState();
}

class _CounterWidgetState extends State<CounterWidget> {
  int _count = 0;

  void _increment() {
    setState(() {
      _count++;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          Text("Count: $_count", style: TextStyle(fontSize: 24)),
          ElevatedButton(
            onPressed: _increment,
            child: Text("Increment"),
          ),
        ],
      ),
    );
  }
}
```

📌 `setState()`를 사용하여 상태를 변경하면 **UI가 다시 렌더링**됨.

---

## 5. 기본 위젯 정리

### 5-1. 레이아웃 관련 위젯
| 위젯 | 설명 |
|------|------|
| **Container** | 크기, 색상, 여백을 지정하는 박스 |
| **Row** | 가로 방향으로 위젯 배치 |
| **Column** | 세로 방향으로 위젯 배치 |
| **Stack** | 겹쳐서 배치 (Z축) |
| **Expanded** | 남은 공간을 차지하도록 확장 |
| **SizedBox** | 크기 지정 (공백 추가 가능) |

```dart
Container(
  padding: EdgeInsets.all(10),
  decoration: BoxDecoration(color: Colors.blue),
  child: Text("Hello, Container!"),
);
```

---

### 5-2. 입력 위젯
| 위젯 | 설명 |
|------|------|
| **TextField** | 텍스트 입력 필드 |
| **Checkbox** | 체크박스 |
| **Radio** | 라디오 버튼 |
| **Switch** | 스위치 버튼 |
| **Slider** | 슬라이더 (범위 조정) |

```dart
TextField(
  decoration: InputDecoration(labelText: "Enter text"),
);
```

---

### 5-3. 버튼 위젯
| 위젯 | 설명 |
|------|------|
| **ElevatedButton** | 기본 버튼 |
| **TextButton** | 텍스트 버튼 |
| **OutlinedButton** | 테두리 버튼 |
| **IconButton** | 아이콘 버튼 |

```dart
ElevatedButton(
  onPressed: () {
    print("Button Pressed");
  },
  child: Text("Click Me"),
);
```

---

### 5-4. 목록 & 스크롤 위젯
| 위젯 | 설명 |
|------|------|
| **ListView** | 스크롤 가능한 리스트 |
| **GridView** | 그리드 형태의 리스트 |
| **SingleChildScrollView** | 단일 스크롤 가능 뷰 |

```dart
ListView(
  children: [
    ListTile(title: Text("Item 1")),
    ListTile(title: Text("Item 2")),
    ListTile(title: Text("Item 3")),
  ],
);
```

---

## 6. 위젯 트리와 BuildContext

### 6-1. 위젯 트리란?
- Flutter는 **위젯 트리(WIDGET TREE)** 구조로 UI를 구성합니다.
- 모든 위젯은 **부모-자식 관계**를 가지며 계층적으로 정리됩니다.

```dart
MaterialApp
 ├── Scaffold
 │   ├── AppBar
 │   ├── Body (Center)
 │       ├── Column
 │           ├── Text
 │           ├── ElevatedButton
```

✔ 위젯이 변경되면 **Flutter는 최소한의 UI만 다시 그려서 성능을 최적화**합니다.

---

## 🎯 정리

- **Flutter의 UI는 위젯으로 구성되며, 모든 요소가 위젯**  
- **위젯은 Stateless(정적)과 Stateful(동적)로 나뉨**  
- **Flutter UI는 "위젯 트리" 형태로 구성됨**  
- **레이아웃, 버튼, 리스트 등의 다양한 위젯 제공**  
- **StatefulWidget에서는 `setState()`를 사용하여 UI 갱신**

