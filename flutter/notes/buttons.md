# 🔲 Flutter 버튼과 이벤트 처리

Flutter에서 버튼은 **사용자의 액션(클릭, 탭 등)을 처리하는 주요 UI 요소**입니다.  
버튼을 클릭하면 특정 동작을 실행하는 **이벤트(Event Handling)**을 설정할 수 있습니다.

---

## 1️⃣ 기본 버튼 종류

Flutter에서는 다양한 버튼 위젯을 제공합니다.

| 버튼 위젯 | 설명 |
|----------|------|
| `ElevatedButton` | 기본 버튼 (입체 효과) |
| `TextButton` | 텍스트만 있는 버튼 (평면) |
| `OutlinedButton` | 테두리가 있는 버튼 |
| `IconButton` | 아이콘 버튼 |
| `FloatingActionButton` | 둥근 플로팅 버튼 |

---

## 2️⃣ `ElevatedButton` - 입체 효과 버튼

가장 일반적인 버튼으로, **입체 효과가 있는 기본 버튼**입니다.

```dart
ElevatedButton(
  onPressed: () {
    print("버튼 클릭!");
  },
  child: Text("클릭"),
)
```

✔ **`onPressed`** → 버튼 클릭 시 실행할 동작을 정의  
✔ **`child`** → 버튼 안에 표시할 위젯 (`Text`, `Icon` 등)  

---

## 3️⃣ `TextButton` - 평면 버튼

**배경이 없는** 간단한 텍스트 버튼입니다.

```dart
TextButton(
  onPressed: () {
    print("텍스트 버튼 클릭");
  },
  child: Text("텍스트 버튼"),
)
```

---

## 4️⃣ `OutlinedButton` - 테두리 버튼

**테두리만 있는 버튼**으로, 배경 없이 강조할 때 사용합니다.

```dart
OutlinedButton(
  onPressed: () {
    print("Outlined 버튼 클릭");
  },
  child: Text("Outlined 버튼"),
)
```

---

## 5️⃣ `IconButton` - 아이콘 버튼

**아이콘만 표시되는 버튼**.

```dart
IconButton(
  icon: Icon(Icons.thumb_up),
  onPressed: () {
    print("아이콘 버튼 클릭");
  },
)
```

✔ `icon` 속성을 사용하여 **아이콘 추가** 가능.

---

## 6️⃣ `FloatingActionButton` - 화면 위에 떠 있는 버튼

**스크롤 시에도 유지되는 버튼**으로, `Scaffold`에서 사용.

```dart
FloatingActionButton(
  onPressed: () {
    print("FAB 클릭");
  },
  child: Icon(Icons.add),
)
```

✔ `Scaffold`에서 **floatingActionButton** 속성으로 설정 가능.

```dart
Scaffold(
  floatingActionButton: FloatingActionButton(
    onPressed: () {
      print("FAB 클릭");
    },
    child: Icon(Icons.add),
  ),
)
```

---

## 7️⃣ 버튼 스타일 적용 (`style` 속성)

버튼의 **색상, 모양, 패딩, 크기** 등을 커스텀할 수 있습니다.

```dart
ElevatedButton(
  style: ElevatedButton.styleFrom(
    primary: Colors.blue, // 배경색
    onPrimary: Colors.white, // 글자색
    padding: EdgeInsets.symmetric(horizontal: 30, vertical: 15), // 패딩
    shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(10)), // 모서리 둥글게
  ),
  onPressed: () {
    print("스타일 적용 버튼 클릭");
  },
  child: Text("Styled Button"),
)
```

#### 속성
- `primary` → 배경색  
- `onPrimary` → 글자색  
- `padding` → 버튼 내부 여백  
- `shape` → 모양 변경 (`RoundedRectangleBorder`로 둥글게)  

---

## 8️⃣ `onPressed: null` - 버튼 비활성화

버튼을 비활성화하려면 `onPressed`를 `null`로 설정합니다.

```dart
ElevatedButton(
  onPressed: null, // 버튼 비활성화
  child: Text("비활성화된 버튼"),
)
```

---

## 9️⃣ 이벤트 처리 (onPressed)

버튼 클릭 시 특정 동작을 실행하는 이벤트 처리.

```dart
ElevatedButton(
  onPressed: () {
    print("버튼 클릭 이벤트 발생!");
  },
  child: Text("이벤트 버튼"),
)
```

✔ `onPressed` → 버튼을 클릭하면 실행할 **콜백 함수(익명 함수)**를 정의.

---

## 🔟 버튼으로 변수 값 변경

버튼을 클릭할 때 변수 값을 변경하고, UI를 업데이트하는 예제.

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
  int _count = 0; // 상태 변수

  void _increment() {
    setState(() {
      _count++; // 변수 값 증가
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text("버튼 이벤트 처리")),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text("카운트: $_count", style: TextStyle(fontSize: 24)),
            ElevatedButton(
              onPressed: _increment,
              child: Text("카운트 증가"),
            ),
          ],
        ),
      ),
    );
  }
}
```

✔ 버튼 클릭 시 `_increment()` 함수가 실행되고, `_count` 값이 증가 
✔ `setState()`를 사용하여 UI를 업데이트

---

## 1️⃣1️⃣ 여러 버튼 동작 처리

```dart
void _handlePress(String action) {
  print("버튼 클릭: $action");
}

ElevatedButton(
  onPressed: () => _handlePress("확인"),
  child: Text("확인"),
);

TextButton(
  onPressed: () => _handlePress("취소"),
  child: Text("취소"),
);
```

✔ `onPressed`에 **함수를 직접 호출**할 수도 있음.

---

## 1️⃣2️⃣ `GestureDetector` - 커스텀 제스처 이벤트

Flutter에서는 버튼 외에도 다양한 **제스처 이벤트**(탭, 길게 누르기, 스와이프 등)를 감지할 수 있습니다.

```dart
GestureDetector(
  onTap: () {
    print("화면을 터치했습니다!");
  },
  onLongPress: () {
    print("길게 눌렀습니다!");
  },
  child: Container(
    width: 200,
    height: 100,
    color: Colors.blue,
    child: Center(child: Text("눌러보세요!", style: TextStyle(color: Colors.white))),
  ),
)
```

✔ `onTap` → 터치(클릭) 감지  
✔ `onLongPress` → 길게 누르기 감지  

---

## 🎯 정리

✔ `ElevatedButton` → 기본 버튼  
✔ `TextButton` → 텍스트 버튼  
✔ `OutlinedButton` → 테두리 버튼  
✔ `IconButton` → 아이콘 버튼  
✔ `FloatingActionButton` → 둥근 플로팅 버튼  
✔ 버튼 스타일 (`style` 속성) → 색상, 모양, 크기 변경 가능  
✔ 이벤트 처리 (`onPressed`) → 버튼 클릭 시 특정 동작 실행  
✔ GestureDetector → 버튼 외의 터치 이벤트 감지 가능  
