# 🔄 Flutter BLoC 패턴

**BLoC (Business Logic Component)** 패턴은 **Flutter에서 상태를 효율적으로 관리하는 아키텍처 패턴**입니다.  
BLoC 패턴은 **이벤트(Event)와 상태(State)를 기반으로 동작**하며, **대규모 애플리케이션에서 강력한 상태 관리**를 제공합니다.

---

## 1️⃣ BLoC 패턴이란?

- 이벤트(Event) → 비즈니스 로직(BLoC) → 새로운 상태(State)를 UI에 반영하는 구조.
- **Streams(스트림)** 을 기반으로 동작하여 비동기 처리에 강함.
- **단방향 데이터 흐름**을 유지하여 코드 유지보수성이 높아짐.

---

## 2️⃣ BLoC 패턴의 흐름
#### BLoC 흐름도
```
User Action (이벤트 발생) → Bloc (이벤트 처리) → State 변경 → UI 업데이트
```
#### 1. 사용자가 이벤트(Event)를 발생 (ex: 버튼 클릭)
#### 2. BLoC에서 이벤트를 수신하고 상태(State)를 변경
#### 3. UI는 새로운 상태(State)를 감지하고 업데이트

---

## 3️⃣ BLoC 패키지 설치

Flutter에서 BLoC 패턴을 쉽게 구현하기 위해 **flutter_bloc 패키지를 사용**합니다.

```sh
flutter pub add flutter_bloc
```

또는 `pubspec.yaml`에 추가:

```yaml
dependencies:
  flutter:
    sdk: flutter
  flutter_bloc: ^8.1.3  # 최신 버전 확인 후 업데이트
```

설치 후 패키지를 불러옵니다.

```dart
import 'package:flutter_bloc/flutter_bloc.dart';
```

---

## 4️⃣ BLoC 기본 구조

### 1) `Event` 정의 (사용자의 액션)

이벤트는 **사용자가 실행하는 동작**을 의미합니다.

```dart
abstract class CounterEvent {}

class IncrementEvent extends CounterEvent {} // 증가 버튼 클릭
class DecrementEvent extends CounterEvent {} // 감소 버튼 클릭
```

---

### 2) `State` 정의 (UI의 상태)

상태는 **UI의 현재 상태를 나타냅니다**.

```dart
abstract class CounterState {
  final int count;
  CounterState(this.count);
}

class CounterInitial extends CounterState {
  CounterInitial() : super(0);
}

class CounterUpdated extends CounterState {
  CounterUpdated(int count) : super(count);
}
```

---

### 3) `Bloc` 구현 (이벤트 처리 및 상태 관리)

`Bloc` 클래스는 **이벤트를 받아서 상태를 변경하는 로직을 포함**합니다.

```dart
class CounterBloc extends Bloc<CounterEvent, CounterState> {
  CounterBloc() : super(CounterInitial()) {
    on<IncrementEvent>((event, emit) {
      emit(CounterUpdated(state.count + 1)); // 상태 변경 (증가)
    });

    on<DecrementEvent>((event, emit) {
      emit(CounterUpdated(state.count - 1)); // 상태 변경 (감소)
    });
  }
}
```

✔ `on<EventType>()` → 특정 이벤트가 발생했을 때 실행할 로직 정의  
✔ `emit()` → 상태를 변경하여 UI를 업데이트  

---

## 5️⃣ UI 에서 BLoC 사용하기

### 1) `BlocProvider`로 BLoC 제공

**BLoC을 `BlocProvider`로 감싸서 위젯 트리에 제공**합니다.

```dart
void main() {
  runApp(
    BlocProvider(
      create: (context) => CounterBloc(),
      child: MyApp(),
    ),
  );
}
```

---

### 2) `BlocBuilder`로 상태 감지 및 UI 업데이트

**BlocBuilder**를 사용하여 UI를 상태(State)에 따라 변경합니다.

```dart
class CounterScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text("Flutter BLoC Example")),
      body: Center(
        child: BlocBuilder<CounterBloc, CounterState>(
          builder: (context, state) {
            return Text("Count: ${state.count}", style: TextStyle(fontSize: 24));
          },
        ),
      ),
      floatingActionButton: Column(
        mainAxisAlignment: MainAxisAlignment.end,
        children: [
          FloatingActionButton(
            onPressed: () => context.read<CounterBloc>().add(IncrementEvent()),
            child: Icon(Icons.add),
          ),
          SizedBox(height: 10),
          FloatingActionButton(
            onPressed: () => context.read<CounterBloc>().add(DecrementEvent()),
            child: Icon(Icons.remove),
          ),
        ],
      ),
    );
  }
}
```

✔ `BlocBuilder<Bloc, State>` → 상태(State)가 변경될 때마다 UI를 업데이트  
✔ `context.read<CounterBloc>().add(Event())` → 이벤트를 BLoC에 전달  

---

## 6️⃣ `BlocListener` 사용 (이벤트 기반 UI 변경)
특정 이벤트가 발생했을 때 **스낵바(Snackbar) 표시** 등의 효과를 줄 때 사용.

```dart
BlocListener<CounterBloc, CounterState>(
  listener: (context, state) {
    ScaffoldMessenger.of(context).showSnackBar(
      SnackBar(content: Text("State changed: ${state.count}")),
    );
  },
  child: CounterScreen(),
);
```

✔ `BlocListener` → **UI 업데이트 없이 이벤트를 감지하여 특정 동작 실행** 가능.

---

## 7️⃣ `MultiBlocProvider` (여러 BLoC 사용)
`MultiBlocProvider`를 사용하면 **여러 개의 BLoC을 한 번에 제공**할 수 있습니다.

```dart
void main() {
  runApp(
    MultiBlocProvider(
      providers: [
        BlocProvider(create: (context) => CounterBloc()),
        BlocProvider(create: (context) => AnotherBloc()),
      ],
      child: MyApp(),
    ),
  );
}
```

✔ `MultiBlocProvider` → **대규모 앱에서 여러 개의 상태를 관리할 때 유용**.

---

## 8️⃣ `BlocObserver` (디버깅 및 로깅)
**모든 BLoC의 상태 변화를 감지**하고 디버깅할 수 있습니다.

```dart
class MyBlocObserver extends BlocObserver {
  @override
  void onChange(BlocBase bloc, Change change) {
    super.onChange(bloc, change);
    print("${bloc.runtimeType} changed: $change");
  }
}

void main() {
  Bloc.observer = MyBlocObserver(); // BlocObserver 설정
  runApp(MyApp());
}
```

✔ `BlocObserver` → 모든 상태 변경을 추적하여 **디버깅이 쉬워짐**.

---

## 🎯 정리

| BLoC 구성 요소 | 설명 |
|--------------|------|
| **Event** | 사용자의 액션 (버튼 클릭 등) |
| **State** | 현재 UI의 상태 |
| **Bloc** | 이벤트를 받아 상태를 변경 |
| **BlocProvider** | BLoC을 앱 위젯 트리에 제공 |
| **BlocBuilder** | 상태 변화를 감지하여 UI 업데이트 |
| **BlocListener** | UI 업데이트 없이 특정 이벤트 감지 |
| **MultiBlocProvider** | 여러 개의 BLoC을 한 번에 제공 |
| **BlocObserver** | 상태 변화를 추적하는 디버깅 도구 |
