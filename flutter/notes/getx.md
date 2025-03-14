# 🔄 GetX

**GetX**는 Flutter에서 가장 간단하고 강력한 **상태 관리(State Management) 라이브러리**입니다.  
**별도의 코드 양이 적고 성능이 뛰어나며, 상태 관리뿐만 아니라 라우팅, 의존성 주입도 지원**합니다.

---

## 1. GetX란?

**GetX의 주요 기능**
- 간결한 상태 관리 (`obs`, `GetBuilder`, `GetX`)
- 라우팅 기능 내장 (`Get.to()`, `Get.off()`)
- 의존성 주입 가능 (`Get.put()`, `Get.find()`)
- 퍼포먼스 최적화 (`setState()` 없이 UI 업데이트)

✔ **GetX vs 다른 상태 관리 라이브러리**
| 기능 | GetX | Provider | BLoC |
|------|------|----------|------|
| 상태 관리 | ✅ 간단함 | ✅ 중간 난이도 | ❌ 복잡함 |
| 코드 양 | ✅ 적음 | ❌ 많음 | ❌ 많음 |
| 퍼포먼스 | ✅ 최적화 | ✅ 적절함 | ✅ 적절함 |
| 의존성 주입 | ✅ 지원 | ❌ 미지원 | ❌ 미지원 |
| 라우팅 지원 | ✅ 있음 | ❌ 없음 | ❌ 없음 |

---

## 2. GetX 설치

Flutter 프로젝트에서 `GetX`를 사용하려면 패키지를 추가해야 합니다.

```sh
flutter pub add get
```

또는 `pubspec.yaml`에 직접 추가:

```yaml
dependencies:
  flutter:
    sdk: flutter
  get: ^4.6.5  # 최신 버전 확인 후 변경 가능
```

설치 후 패키지를 불러옵니다.

```dart
import 'package:get/get.dart';
```

---

## 3. GetX 상태 관리 방식

| 상태 관리 방식 | 설명 |
|--------------|------|
| `.obs` (반응형 상태 관리) | 값이 변경될 때 UI가 자동으로 업데이트 |
| `GetBuilder` (단순 상태 관리) | 특정 위젯만 업데이트 |
| `GetX` (자동 상태 감지) | GetBuilder보다 더 강력한 기능 제공 |

---

## 4. `.obs` (반응형 상태 관리)

`obs`를 사용하면 변수 값이 변경될 때 UI가 자동 업데이트됩니다.

### `Controller` 만들기

```dart
import 'package:get/get.dart';

class CounterController extends GetxController {
  var count = 0.obs; // 반응형 상태

  void increment() {
    count++; // 상태 변경 시 자동 UI 업데이트
  }
}
```

✔ `.obs` → 상태를 감지할 수 있도록 선언.

---

### `Get.put()`을 사용하여 Controller 주입

```dart
final CounterController counterController = Get.put(CounterController());
```

✔ `Get.put(Controller())` → 컨트롤러를 메모리에 등록.

---

### `Obx`를 사용하여 UI 업데이트

```dart
class CounterScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    final CounterController counterController = Get.find();

    return Scaffold(
      appBar: AppBar(title: Text("GetX Counter Example")),
      body: Center(
        child: Obx(() => Text("Count: ${counterController.count}",
            style: TextStyle(fontSize: 24))),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: counterController.increment,
        child: Icon(Icons.add),
      ),
    );
  }
}
```

✔ `Obx(() => Widget())` → `obs` 상태가 변경될 때 자동으로 UI 업데이트  
✔ `Get.find<CounterController>()` → 이미 등록된 컨트롤러 가져오기  

---

## 5. `GetBuilder` (단순 상태 관리)
반응형(`obs`) 없이도 상태 변경 가능하고 `GetBuilder`는 **필요한 위젯만 업데이트**하는 방식으로, 퍼포먼스가 뛰어남.

### ✅ 1) `Controller` 만들기

```dart
class CounterController extends GetxController {
  int count = 0;

  void increment() {
    count++;
    update(); // UI 업데이트
  }
}
```

✔ `update()`를 호출해야 UI가 갱신됨.

---

### ✅ 2) `GetBuilder` 사용하기

```dart
class CounterScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text("GetBuilder Example")),
      body: Center(
        child: GetBuilder<CounterController>(
          init: CounterController(),
          builder: (controller) {
            return Text("Count: ${controller.count}", style: TextStyle(fontSize: 24));
          },
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () => Get.find<CounterController>().increment(),
        child: Icon(Icons.add),
      ),
    );
  }
}
```

✔ `update()`를 호출해야 **GetBuilder가 다시 렌더링됨**  
✔ `GetBuilder`는 특정 위젯만 업데이트하기 때문에 **퍼포먼스가 좋음**  

---

## 6. `GetX` (자동 상태 감지)

- **`GetX`는 `GetBuilder`보다 더 강력한 상태 관리 기능을 제공**.
- **자동으로 상태 변경을 감지하여 UI 업데이트**.

### ✅ 1) `Controller` 만들기

```dart
class CounterController extends GetxController {
  var count = 0.obs;

  void increment() {
    count++;
  }
}
```

---

### ✅ 2) `GetX` 사용하기

```dart
class CounterScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text("GetX Example")),
      body: Center(
        child: GetX<CounterController>(
          init: CounterController(),
          builder: (controller) {
            return Text("Count: ${controller.count}", style: TextStyle(fontSize: 24));
          },
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () => Get.find<CounterController>().increment(),
        child: Icon(Icons.add),
      ),
    );
  }
}
```

✔ **`Obx`와 유사하지만, `GetX`는 더 많은 기능을 제공**  

---

## 7. `Get.put()` vs `Get.lazyPut()` vs `Get.find()`

| 메서드 | 설명 |
|--------|------|
| **`Get.put()`** | 즉시 인스턴스 생성 (앱 실행과 동시에 메모리에 저장) |
| **`Get.lazyPut()`** | 처음 호출될 때 인스턴스 생성 (메모리 절약) |
| **`Get.find()`** | 기존에 생성된 인스턴스를 가져옴 |

```dart
final CounterController counterController = Get.put(CounterController()); // 즉시 생성
final CounterController counterLazy = Get.lazyPut(() => CounterController()); // 필요할 때 생성
final CounterController counterFind = Get.find<CounterController>(); // 기존 인스턴스 가져오기
```

---

## 8. `Get.delete()` (Controller 삭제)

- 사용하지 않는 컨트롤러를 삭제하여 **메모리 최적화 가능**.

```dart
Get.delete<CounterController>(); // 컨트롤러 삭제
```

---

## 🎯 정리

| GetX 상태 관리 방식 | 설명 |
|------------------|------|
| **`.obs` + `Obx`** | 반응형 상태 관리 (자동 UI 업데이트) |
| **`GetBuilder`** | 단순 상태 관리 (`update()` 필요) |
| **`GetX`** | 자동 상태 감지 및 UI 업데이트 |
| **`Get.put()`** | 즉시 인스턴스 생성 |
| **`Get.lazyPut()`** | 필요할 때 인스턴스 생성 (메모리 절약) |
| **`Get.find()`** | 기존 인스턴스 가져오기 |
| **`Get.delete()`** | 컨트롤러 삭제 (메모리 최적화) |

🚀 **이제 GetX를 활용하여 강력한 상태 관리를 구현하세요!**
