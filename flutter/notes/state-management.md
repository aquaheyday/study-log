# Flutter 상태관리 기본 개념

Flutter에서는 **StatelessWidget**과 **StatefulWidget**을 사용하여 UI를 구성합니다.  
상태(state)를 관리하는 방식에 따라 다양한 방법이 존재하며, 대표적으로 **setState, Provider, Riverpod, Bloc** 등을 사용할 수 있습니다.

---

## 1. 상태(State)란?
Flutter에서 **상태**란 UI가 변경될 수 있는 데이터나 속성을 의미합니다.  
예를 들어, **사용자의 입력, 버튼 클릭, API 응답** 등이 상태에 해당합니다.

- **StatelessWidget**: 상태가 없는 위젯 (변하지 않는 UI)
- **StatefulWidget**: 상태를 가질 수 있는 위젯 (변경 가능한 UI)

---

## 2. `setState`를 이용한 상태관리 (기본)
가장 기본적인 상태관리 방법으로, `StatefulWidget`에서 **setState()**를 사용하여 상태를 변경합니다.

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
      appBar: AppBar(title: Text('setState 예제')),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text('Counter: $_counter', style: TextStyle(fontSize: 24)),
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
✅ **setState()**를 호출하면 `build()` 함수가 다시 실행되어 UI가 업데이트됩니다.  
❌ 단점: **규모가 커지면 성능 저하 및 코드 복잡도 증가**

---

## 3. `InheritedWidget`을 이용한 상태관리
위젯 트리에서 데이터를 공유할 때 사용하는 방법입니다.

```dart
class MyData extends InheritedWidget {
  final int counter;

  MyData({required this.counter, required Widget child}) : super(child: child);

  static MyData? of(BuildContext context) {
    return context.dependOnInheritedWidgetOfExactType<MyData>();
  }

  @override
  bool updateShouldNotify(MyData oldWidget) {
    return oldWidget.counter != counter;
  }
}
```
✅ **setState보다 구조적인 데이터 공유 가능**  
❌ **코드가 길고 복잡해질 수 있음**  

---

## 4. `Provider`를 이용한 상태관리 (추천)
Flutter에서 공식적으로 권장하는 상태관리 패턴입니다.

### 4.1 `provider` 패키지 설치
```yaml
dependencies:
  flutter:
    sdk: flutter
  provider: ^6.0.0
```

### 4.2 `ChangeNotifier`와 `Consumer`를 사용한 상태관리
```dart
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

void main() {
  runApp(
    ChangeNotifierProvider(
      create: (context) => CounterProvider(),
      child: MyApp(),
    ),
  );
}

class CounterProvider extends ChangeNotifier {
  int _counter = 0;

  int get counter => _counter;

  void increment() {
    _counter++;
    notifyListeners();
  }
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: CounterScreen(),
    );
  }
}

class CounterScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text('Provider 예제')),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text(
              'Counter: ${context.watch<CounterProvider>().counter}',
              style: TextStyle(fontSize: 24),
            ),
            ElevatedButton(
              onPressed: () {
                context.read<CounterProvider>().increment();
              },
              child: Text('Increase'),
            ),
          ],
        ),
      ),
    );
  }
}
```
✅ **전역 상태 공유 가능**  
✅ **UI가 효율적으로 업데이트됨**  
❌ **초기 학습 필요**  

---

## 5. `Riverpod`를 이용한 상태관리 (더 간단한 Provider)
### 5.1 `flutter_riverpod` 패키지 설치
```yaml
dependencies:
  flutter:
    sdk: flutter
  flutter_riverpod: ^2.0.0
```

### 5.2 `Riverpod` 사용 예제
```dart
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

final counterProvider = StateProvider<int>((ref) => 0);

void main() {
  runApp(ProviderScope(child: MyApp()));
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: CounterScreen(),
    );
  }
}

class CounterScreen extends ConsumerWidget {
  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final counter = ref.watch(counterProvider);

    return Scaffold(
      appBar: AppBar(title: Text('Riverpod 예제')),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text('Counter: $counter', style: TextStyle(fontSize: 24)),
            ElevatedButton(
              onPressed: () {
                ref.read(counterProvider.notifier).state++;
              },
              child: Text('Increase'),
            ),
          ],
        ),
      ),
    );
  }
}
```
✅ **Provider보다 더 간단하고 강력한 상태관리 가능**  
✅ **전역 상태 관리 용이**  
❌ **초기 학습 필요**  

---

## 6. `Bloc`을 이용한 상태관리 (대규모 프로젝트)
**Bloc (Business Logic Component)** 패턴은 대규모 프로젝트에서 추천하는 상태관리 방식입니다.

🔹 `flutter_bloc` 패키지 사용  
🔹 이벤트(Event)와 상태(State)를 기반으로 동작  
🔹 구조적인 코드 작성 가능  

```yaml
dependencies:
  flutter_bloc: ^8.0.0
```

```dart
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

class CounterCubit extends Cubit<int> {
  CounterCubit() : super(0);

  void increment() => emit(state + 1);
}

void main() {
  runApp(
    BlocProvider(
      create: (context) => CounterCubit(),
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

class CounterScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text('Bloc 예제')),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            BlocBuilder<CounterCubit, int>(
              builder: (context, count) {
                return Text('Counter: $count', style: TextStyle(fontSize: 24));
              },
            ),
            ElevatedButton(
              onPressed: () {
                context.read<CounterCubit>().increment();
              },
              child: Text('Increase'),
            ),
          ],
        ),
      ),
    );
  }
}
```
✅ **대규모 프로젝트에 적합**  
✅ **명확한 데이터 흐름**  
❌ **구현이 복잡함**  

---

## 7. 결론
| 상태관리 방식 | 특징 |
|--------------|----------------------|
| `setState` | 작은 앱에 적합 |
| `InheritedWidget` | 데이터 공유 가능 |
| `Provider` | 가장 많이 사용됨 |
| `Riverpod` | Provider보다 간단 |
| `Bloc` | 대규모 프로젝트에 적합 |

