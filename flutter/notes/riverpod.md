# 🔄 Flutter Riverpod

**Riverpod**은 Flutter에서 강력한 **상태 관리(State Management) 라이브러리**입니다.  
**Provider의 단점을 보완한 업그레이드 버전**으로, 더 안전하고 간결한 코드 작성을 지원합니다.

---

## 1️⃣ Riverpod이란?

- Flutter 공식 상태 관리 라이브러리 (`Provider`의 개선 버전)
- 전역 상태 관리 가능
- build-context 없이 어디서든 사용 가능
- `notifyListeners()` 없이 자동으로 상태 업데이트
- 의존성 주입 가능 (Dependency Injection)

#### Provider vs Riverpod 차이점
| 특징 | Provider | Riverpod |
|------|----------|----------|
| `ChangeNotifier` 사용 | ✅ 사용 | ❌ 사용하지 않음 |
| `context` 필요 여부 | ✅ 필요 | ❌ 불필요 |
| `notifyListeners()` 필요 여부 | ✅ 필요 | ❌ 자동 상태 업데이트 |
| 의존성 주입 | ❌ 제한적 | ✅ 강력한 DI 지원 |

---

## 2️⃣ Riverpod 설치

```sh
flutter pub add flutter_riverpod
```

또는 `pubspec.yaml`에 추가:

```yaml
dependencies:
  flutter:
    sdk: flutter
  flutter_riverpod: ^2.0.0  # 최신 버전 확인 후 변경 가능
```

설치 후 패키지를 불러옵니다.

```dart
import 'package:flutter_riverpod/flutter_riverpod.dart';
```

---

## 3️⃣ Riverpod 기본 사용법

Riverpod에서 상태를 관리하려면 `Provider`를 생성해야 합니다.

### 1) `Provider` (읽기 전용 상태)

**변하지 않는 상태를 관리**할 때 사용합니다.

```dart
final helloProvider = Provider<String>((ref) {
  return "Hello, Riverpod!";
});
```

#### 사용법
```dart
Consumer(
  builder: (context, ref, child) {
    final message = ref.watch(helloProvider);
    return Text(message);
  },
);
```

---

### 2) `StateProvider` (기본적인 상태 관리)

**간단한 상태 값 변경**(예: 카운터)을 관리할 때 사용합니다.

```dart
final counterProvider = StateProvider<int>((ref) => 0);
```

#### UI에서 사용
```dart
Consumer(
  builder: (context, ref, child) {
    final count = ref.watch(counterProvider);
    return Column(
      children: [
        Text("카운트: $count"),
        ElevatedButton(
          onPressed: () => ref.read(counterProvider.notifier).state++,
          child: Text("증가"),
        ),
      ],
    );
  },
);
```

✔ `.notifier` → 상태 변경 가능  
✔ `Provider` → 자동으로 UI를 업데이트  

---

## 4️⃣ `StateNotifierProvider` (복잡한 상태 관리)

**클래스 기반 상태 관리** (`ChangeNotifier` 대체) 이며, 여러 개의 값을 가진 **복잡한 상태**를 관리할 때 사용합니다.

#### 1. `StateNotifier` 상태 클래스 만들기
```dart
import 'package:flutter_riverpod/flutter_riverpod.dart';

class CounterNotifier extends StateNotifier<int> {
  CounterNotifier() : super(0);

  void increment() => state++;
  void decrement() => state--;
}

final counterNotifierProvider =
    StateNotifierProvider<CounterNotifier, int>((ref) => CounterNotifier());
```

#### 2. UI 에서 사용
```dart
Consumer(
  builder: (context, ref, child) {
    final count = ref.watch(counterNotifierProvider);
    return Column(
      children: [
        Text("카운트: $count"),
        ElevatedButton(
          onPressed: () => ref.read(counterNotifierProvider.notifier).increment(),
          child: Text("증가"),
        ),
        ElevatedButton(
          onPressed: () => ref.read(counterNotifierProvider.notifier).decrement(),
          child: Text("감소"),
        ),
      ],
    );
  },
);
```

✔ **복잡한 상태를 `StateNotifier`로 관리 가능**  
✔ `.notifier` → 상태 변경  

---

## 5️⃣ `FutureProvider` (비동기 데이터 관리)

API 호출과 같은 **비동기 작업**을 처리할 때 사용합니다.

#### 1. `FutureProvider` 선언
```dart
final dataProvider = FutureProvider<String>((ref) async {
  await Future.delayed(Duration(seconds: 2)); // 2초 후 데이터 반환
  return "비동기 데이터 로드 완료!";
});
```

#### 2. UI 에서 사용
```dart
Consumer(
  builder: (context, ref, child) {
    final asyncValue = ref.watch(dataProvider);

    return asyncValue.when(
      data: (data) => Text(data),
      loading: () => CircularProgressIndicator(),
      error: (error, stack) => Text("에러 발생: $error"),
    );
  },
);
```

✔ `when()` 메서드를 사용하여 상태별 UI 처리 가능
- `data`: 성공 시 UI 표시
- `loading`: 로딩 중 UI 표시
- `error`: 에러 발생 시 UI 표시

---

## 6️⃣ `StreamProvider` (실시간 데이터)

Firebase 등 **실시간 데이터를 관리**할 때 사용합니다.

#### 1. `StreamProvider` 선언
```dart
final counterStreamProvider = StreamProvider<int>((ref) {
  return Stream.periodic(Duration(seconds: 1), (count) => count);
});
```

#### 2. UI 에서 사용
```dart
Consumer(
  builder: (context, ref, child) {
    final asyncValue = ref.watch(counterStreamProvider);

    return asyncValue.when(
      data: (data) => Text("카운트: $data"),
      loading: () => CircularProgressIndicator(),
      error: (error, stack) => Text("에러 발생: $error"),
    );
  },
);
```

✔ `StreamProvider` → 실시간 데이터 변경을 UI에 반영 가능.

---

## 7️⃣ `ProviderScope` (전역 Provider 관리)

Riverpod을 사용하려면 `ProviderScope`를 최상위 위젯에 추가해야 합니다.

```dart
void main() {
  runApp(ProviderScope(child: MyApp()));
}
```

✔ `ProviderScope`가 있어야 **상태가 전역적으로 관리됨.**

---

## 🎯 정리

| Riverpod Provider | 설명 |
|------------------|------|
| **Provider** | 읽기 전용 상태 (변경 불가) |
| **StateProvider** | 간단한 상태 관리 (`setState` 대체) |
| **StateNotifierProvider** | 복잡한 상태 관리 (`ChangeNotifier` 대체) |
| **FutureProvider** | 비동기 데이터 관리 (API 호출 등) |
| **StreamProvider** | 실시간 데이터 관리 (Firebase 등) |
