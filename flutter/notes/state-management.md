# 📌 Flutter 상태 관리

Flutter에서 **상태(State)** 는 **UI에 영향을 주는 데이터**를 의미합니다.  
상태 관리는 **어떻게 데이터를 변경하고, 변경된 데이터를 UI에 반영할 것인지**를 결정하는 중요한 개념입니다.

---

## 1️⃣ 상태(State)란?

Flutter에서 **상태(State)** 는 **앱이 실행되는 동안 변경될 수 있는 데이터**입니다.

#### 1. 상태가 필요한 경우
- 버튼을 클릭할 때 카운트 증가  
- 사용자 입력을 저장하여 화면에 표시  
- API 요청 후 데이터를 업데이트  

#### 2. 상태(State) 변경 방식
- 로컬 상태(Local State) → 개별 위젯 내부에서 관리 (`setState()`)
- 전역 상태(Global State) → 여러 위젯에서 공유 (Provider, Riverpod, Bloc 등 사용)

---

## 2️⃣ StatefulWidget을 사용한 상태 관리 (setState)

Flutter에서 가장 기본적인 상태 관리는 **StatefulWidget**과 `setState()`를 활용하는 방식입니다.

### 1) `StatefulWidget` 구조
- `StatefulWidget` 클래스 → 상태를 관리하는 `State`를 생성
- `State` 클래스 → UI를 빌드하고, `setState()`를 사용하여 UI 갱신  

#### `setState()`란?
로컬 상태 관리(Local State Management) 방식이며 `setState()`를 호출하면 `build()`가 다시 실행되어 UI가 갱신되며 StatefulWidget 내부에서만 사용 가능합니다.

#### ⚠️ 주의점
- `setState()`를 너무 많이 호출하면 성능이 저하될 수 있음.  
- 위젯의 범위를 벗어난 상태(State)는 `setState()`로 관리할 수 없음.

---

### 2) 카운터 증가 버튼 예제

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

  void _incrementCounter() {
    setState(() {
      _count++; // 상태 변경 -> UI 갱신
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text("Counter App")),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text("Count: $_count", style: TextStyle(fontSize: 24)),
            ElevatedButton(
              onPressed: _incrementCounter,
              child: Text("Increment"),
            ),
          ],
        ),
      ),
    );
  }
}
```

---

## 3️⃣ 여러 위젯에서 상태를 공유하려면? (전역 상태 관리)

위의 `setState()` 방식은 위젯 하나의 상태만 관리할 수 있습니다.  
하지만, 앱 전체에서 상태를 공유하려면 전역 상태 관리 방법이 필요합니다.

#### 대표적인 전역 상태 관리 라이브러리
| 방법 | 특징 | 사용 예 |
|------|------|------|
| InheritedWidget | 기본 내장 방식, 적은 데이터 공유에 적합 | 앱 테마, 언어 설정 |
| Provider | Flutter 공식 권장 방식, 간단한 상태 관리 | 로그인 상태, UI 업데이트 |
| Riverpod | Provider의 개선 버전, 간단한 사용법 | API 데이터 관리 |
| Bloc (flutter_bloc) | 이벤트 기반 상태 관리 (Redux와 유사) | 대규모 프로젝트 |
| GetX | 간단한 코드, 성능 최적화 | 상태 및 네비게이션 관리 |

---

## 4️⃣ `InheritedWidget` (Flutter 기본 상태 관리)

Flutter에서 기본적으로 제공하는 전역 상태 관리 방식입니다.  
하지만 코드가 복잡하고 사용이 어렵기 때문에 `Provider` 사용을 추천합니다.

```dart
class MyInheritedWidget extends InheritedWidget {
  final int count;

  MyInheritedWidget({Key? key, required Widget child, required this.count}) 
      : super(key: key, child: child);

  static MyInheritedWidget? of(BuildContext context) {
    return context.dependOnInheritedWidgetOfExactType<MyInheritedWidget>();
  }

  @override
  bool updateShouldNotify(MyInheritedWidget oldWidget) {
    return count != oldWidget.count;
  }
}
```
✔ **복잡한 구조로 인해 실무에서는 거의 사용되지 않음** → `Provider` 또는 `Riverpod`을 사용 추천.

---

## 5️⃣ `Provider` (Flutter 공식 권장 상태 관리)

`Provider`는 Flutter에서 공식적으로 권장하는 상태 관리 라이브러리입니다.

#### 장점
- 코드가 간결하고 사용하기 쉬움.
- `ChangeNotifier`를 사용하여 상태를 자동으로 감지하고 UI 업데이트.

### 1. `Provider` 패키지 설치

```sh
flutter pub add provider
```

#### 2. `Provider` 상태 모델 생성 (`ChangeNotifier` 사용)

```dart
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

// 상태 관리 클래스
class CounterProvider with ChangeNotifier {
  int _count = 0;

  int get count => _count;

  void increment() {
    _count++;
    notifyListeners(); // 상태 변경 알림
  }
}
```

#### 3. `Provider` 적용하기 (`MyApp` 수정)

```dart
void main() {
  runApp(
    ChangeNotifierProvider(
      create: (context) => CounterProvider(),
      child: MyApp(),
    ),
  );
}
```

#### 4. UI에서 상태 사용

```dart
class CounterScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    final counter = Provider.of<CounterProvider>(context);

    return Scaffold(
      appBar: AppBar(title: Text("Provider Example")),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text("Count: ${counter.count}", style: TextStyle(fontSize: 24)),
            ElevatedButton(
              onPressed: () => counter.increment(),
              child: Text("Increment"),
            ),
          ],
        ),
      ),
    );
  }
}
```

✔ `Provider.of<CounterProvider>(context)` → 상태를 가져오고 UI를 업데이트함.

---

## 6️⃣ 기타 상태 관리 방법

| 방식 | 특징 |
|------|------|
| Riverpod | Provider의 개선 버전, 더 간결한 코드 |
| Bloc (flutter_bloc) | Redux 패턴과 유사한 이벤트 기반 상태 관리 |
| GetX | 코드가 간단하고 높은 성능 |
| Redux | 대규모 프로젝트에 적합한 패턴 |

---

## 🎯 정리

✔ Flutter의 상태(State)는 UI에 영향을 주는 데이터  
✔ setState() → 간단한 로컬 상태 관리 (StatefulWidget 사용)  
✔ Provider → Flutter 공식 권장 전역 상태 관리 방식  
✔ Bloc, Riverpod, GetX → 프로젝트 규모에 따라 선택 가능  
