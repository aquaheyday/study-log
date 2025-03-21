# 🔄 Flutter Provider

**Provider**는 Flutter에서 공식적으로 권장하는 **전역 상태 관리 라이브러리**입니다.  
앱의 여러 위젯에서 데이터를 공유하고 UI를 자동으로 업데이트할 수 있습니다.

---

## 1️⃣ Provider란?

**Provider**는 **Flutter의 상태(State)를 관리하는 라이브러리**로,  
UI와 상태를 효율적으로 연결하여 업데이트할 수 있도록 도와줍니다.

#### Provider의 장점
- Flutter 공식 권장 방식
- 간단한 사용법 (`ChangeNotifier` 기반)
- UI 자동 업데이트 (`notifyListeners()` 활용)
- 전역 상태 관리 가능 (어디서든 데이터 접근)

---

## 2️⃣ Provider 설치

Flutter 프로젝트에서 `Provider`를 사용하려면 패키지를 추가해야 합니다.

```sh
flutter pub add provider
```

또는 `pubspec.yaml`에 직접 추가:

```yaml
dependencies:
  flutter:
    sdk: flutter
  provider: ^6.0.5  # 최신 버전 확인 후 변경 가능
```

그런 다음, 패키지를 설치합니다.

```sh
flutter pub get
```

---

## 3️⃣ Provider 기본 사용법

#### 1. `ChangeNotifier`를 사용한 상태 클래스 만들기

`ChangeNotifier`는 Provider에서 상태를 관리하는 기본 클래스입니다.

```dart
import 'package:flutter/material.dart';

class CounterProvider with ChangeNotifier {
  int _count = 0;

  int get count => _count; // 현재 상태 값 반환

  void increment() {
    _count++;
    notifyListeners(); // UI 업데이트
  }
}
```

✔ `notifyListeners()` → 상태가 변경되면 자동으로 UI를 업데이트합니다.


#### 2. `ChangeNotifierProvider`로 상태 제공하기

앱의 최상위 위젯에서 `ChangeNotifierProvider`로 상태를 제공해야 합니다.

```dart
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'counter_provider.dart'; // 상태 클래스 가져오기

void main() {
  runApp(
    ChangeNotifierProvider(
      create: (context) => CounterProvider(), // 상태 객체 생성
      child: MyApp(),
    ),
  );
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: CounterScreen(),
    );
  }
}
```

✔ `ChangeNotifierProvider` → `create:` 에 상태 클래스를 넣어 위젯 트리에 상태를 제공.

#### 3. UI 에서 Provider 상태 사용하기

이제 **UI 에서 Provider 상태를 가져와서 사용**할 수 있습니다.

```dart
class CounterScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    final counter = Provider.of<CounterProvider>(context); // 상태 가져오기

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

✔ `Provider.of<CounterProvider>(context)` → 현재 상태를 가져와 UI 에서 사용할 수 있음.

---

## 4️⃣ 상태 접근 방식 비교

Provider에서 상태를 가져오는 방법은 **3가지**가 있습니다.

### 1) `Provider.of<T>(context)`

`Provider.of`는 상태를 가져올 때 사용되며, 기본적으로 UI가 **자동 업데이트**됩니다.

```dart
final counter = Provider.of<CounterProvider>(context);
```

- `listen: false`를 설정하면 UI가 갱신되지 않습니다.
```dart
final counter = Provider.of<CounterProvider>(context, listen: false);
```

---

### 2) `Consumer<T>`

`Consumer`는 UI에서 **특정 부분만 상태 변경 시 업데이트**할 수 있습니다.  
UI의 일부만 다시 그려야 할 때 유용합니다.

```dart
Consumer<CounterProvider>(
  builder: (context, counter, child) {
    return Text("Count: ${counter.count}", style: TextStyle(fontSize: 24));
  },
);
```

---

### 3) `context.watch<T>()` / `context.read<T>()`

| 메서드 | 설명 | 사용 예 |
|--------|------|--------|
| `context.watch<T>()` | 상태를 읽고, 변경 시 UI를 다시 그림 | `context.watch<CounterProvider>().count` |
| `context.read<T>()` | 한 번만 상태를 읽고 UI를 다시 그리지 않음 | `context.read<CounterProvider>().increment()` |

```dart
// 상태 변경 시 UI 갱신
Text("Count: ${context.watch<CounterProvider>().count}");

// UI 갱신 없이 상태 변경
ElevatedButton(
  onPressed: () => context.read<CounterProvider>().increment(),
  child: Text("Increment"),
);
```

---

### 4) 방법 차이 
| 방법 | 설명 | 사용 예 |
|------|------|------|
| `Provider.of<T>(context)` | 즉시 상태를 가져오고 UI 갱신 | `Provider.of<CounterProvider>(context).count` |
| `Consumer<T>` | UI의 특정 부분만 업데이트 | `Consumer<CounterProvider>(builder: ...)` |
| `context.watch<T>()` / `context.read<T>()` | UI 갱신 여부에 따라 다름 | `context.watch<T>()` (UI 갱신), `context.read<T>()` (한 번만 읽음) |

---

## 5️⃣ 여러 상태 관리 (`MultiProvider`)

`MultiProvider`를 사용하면 **여러 개의 상태를 한 번에 제공**할 수 있습니다.

```dart
void main() {
  runApp(
    MultiProvider(
      providers: [
        ChangeNotifierProvider(create: (context) => CounterProvider()),
        ChangeNotifierProvider(create: (context) => ThemeProvider()), // 추가 가능
      ],
      child: MyApp(),
    ),
  );
}
```

✔ `MultiProvider` → 여러 상태를 한 번에 관리하고, UI 에서 쉽게 접근 가능.

---

## 6️⃣ Provider vs 다른 상태 관리 라이브러리

| 상태 관리 방식 | 장점 | 단점 |
|--------------|------|------|
| Provider | 간단하고 사용하기 쉬움, 공식 권장 | 큰 규모의 앱에서 상태가 많아지면 복잡해질 수 있음 |
| Riverpod | Provider보다 더 간결하고 구조적 | 새로운 문법 학습 필요 |
| Bloc (flutter_bloc) | 명확한 상태 흐름, 대규모 프로젝트에 적합 | 코드가 길어지고 복잡할 수 있음 |
| GetX | 코드가 간결하고 성능 최적화 | 패턴이 명확하지 않아 유지보수가 어려울 수 있음 |

---

## 🎯 정리

✔ `Provider` → Flutter 공식 권장 상태 관리 라이브러리  
✔ `ChangeNotifier` → 상태를 정의하고 `notifyListeners()`로 UI 갱신  
✔ `ChangeNotifierProvider` → 상태를 위젯 트리에 제공  
✔ `Provider.of`, `Consumer`, `context.watch/read`를 활용하여 상태 접근  
✔ `MultiProvider` → 여러 상태를 관리 가능  
