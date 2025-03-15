# 🛠️ Flutter 테스트 가이드

Flutter에서는 **단위 테스트(Unit Test), 위젯 테스트(Widget Test), 통합 테스트(Integration Test)** 를 지원합니다.  
테스트를 통해 코드의 안정성을 높이고, 앱의 동작을 검증할 수 있습니다.

---

## 1. Flutter 테스트의 종류

| 테스트 유형 | 설명 | 사용 패키지 |
|------------|------|------------|
| 단위 테스트 (Unit Test) | 개별 함수, 클래스 로직을 테스트 | `flutter_test` |
| 위젯 테스트 (Widget Test) | UI의 특정 위젯을 테스트 | `flutter_test` |
| 통합 테스트 (Integration Test) | 실제 앱 실행 환경에서 테스트 (네트워크, DB 포함) | `integration_test` |

✔ `flutter_test` 패키지는 기본적으로 포함됨  
✔ `integration_test`는 별도로 추가 필요 (`flutter pub add integration_test`)  

---

## 2. 단위 테스트 (Unit Test)

단위 테스트는 개별 함수, 비즈니스 로직 등을 검증하는 테스트입니다.

### 패키지 설치

```sh
flutter pub add test
```

### 단위 테스트 예제 (`test/calculator_test.dart`)

```dart
import 'package:flutter_test/flutter_test.dart';

int add(int a, int b) => a + b;

void main() {
  test('더하기 함수 테스트', () {
    expect(add(2, 3), 5);
  });
}
```

✔ `test('설명', () {...})` → 개별 테스트 실행  
✔ `expect(결과, 예상값)` → 예상값과 결과 비교  

### 단위 테스트 실행

```sh
flutter test test/calculator_test.dart
```

---

## 3. 위젯 테스트 (Widget Test)

위젯 테스트는 UI 위젯이 예상대로 동작하는지 확인하는 테스트입니다.

### 위젯 테스트 예제 (`test/counter_widget_test.dart`)

```dart
import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';

class CounterWidget extends StatefulWidget {
  @override
  _CounterWidgetState createState() => _CounterWidgetState();
}

class _CounterWidgetState extends State<CounterWidget> {
  int counter = 0;

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Text('Count: $counter', key: Key('counterText')),
        ElevatedButton(
          key: Key('incrementButton'),
          onPressed: () {
            setState(() {
              counter++;
            });
          },
          child: Text("증가"),
        ),
      ],
    );
  }
}

void main() {
  testWidgets('카운터 증가 테스트', (WidgetTester tester) async {
    await tester.pumpWidget(MaterialApp(home: CounterWidget()));

    expect(find.text('Count: 0'), findsOneWidget);

    await tester.tap(find.byKey(Key('incrementButton')));
    await tester.pump();

    expect(find.text('Count: 1'), findsOneWidget);
  });
}
```

✔ `tester.pumpWidget()` → 위젯을 렌더링  
✔ `find.text('Count: 0')` → 특정 텍스트 찾기  
✔ `tester.tap(find.byKey(Key('incrementButton')))` → 버튼 클릭 시뮬레이션  
✔ `tester.pump()` → UI 업데이트 반영  

### 위젯 테스트 실행

```sh
flutter test test/counter_widget_test.dart
```

---

## 4. 통합 테스트 (Integration Test)

통합 테스트는 실제 기기에서 앱을 실행하여 테스트하는 방식입니다.

### 패키지 설치

```sh
flutter pub add integration_test
```

### 통합 테스트 예제 (`integration_test/app_test.dart`)

```dart
import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:integration_test/integration_test.dart';

void main() {
  IntegrationTestWidgetsFlutterBinding.ensureInitialized();

  testWidgets('전체 앱 실행 테스트', (WidgetTester tester) async {
    await tester.pumpWidget(MyApp());

    expect(find.text("Hello, Flutter!"), findsOneWidget);
  });
}
```

✔ `IntegrationTestWidgetsFlutterBinding.ensureInitialized()` → 통합 테스트 초기화  
✔ `tester.pumpWidget()` → 앱 전체 렌더링  

### 통합 테스트 실행

```sh
flutter test integration_test/app_test.dart
```

---

## 5. Mock을 활용한 테스트

네트워크 요청과 같은 외부 의존성을 Mock으로 대체하여 테스트할 수 있습니다.

### `mockito` 패키지 설치

```sh
flutter pub add mockito
flutter pub add build_runner --dev
```

### Mock을 활용한 API 테스트 (`test/api_test.dart`)

```dart
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/annotations.dart';
import 'package:mockito/mockito.dart';

class ApiService {
  Future<String> fetchData() async {
    await Future.delayed(Duration(seconds: 1));
    return "데이터 로드 완료";
  }
}

@GenerateMocks([ApiService])
import 'api_test.mocks.dart';

void main() {
  late MockApiService mockApiService;

  setUp(() {
    mockApiService = MockApiService();
  });

  test('API 호출 테스트', () async {
    when(mockApiService.fetchData()).thenAnswer((_) async => "테스트 데이터");

    final result = await mockApiService.fetchData();
    expect(result, "테스트 데이터");
  });
}
```

✔ `when().thenAnswer()` → Mock 데이터 반환 설정  
✔ `expect(result, "테스트 데이터")` → API 응답 검증  

---

## 6. 테스트 자동화 (CI/CD)

테스트를 GitHub Actions, GitLab CI/CD 등에서 자동 실행할 수 있습니다.

### GitHub Actions 예제 (`.github/workflows/flutter-test.yml`)

```yaml
name: Flutter Test

on:
  push:
    branches:
      - main

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: subosito/flutter-action@v2
        with:
          flutter-version: "3.x"
      - run: flutter pub get
      - run: flutter test
```

✔ 코드가 `main` 브랜치에 푸시될 때 자동으로 테스트 실행  

---

## 🎯 정리

✔ 단위 테스트 → `flutter_test`를 활용하여 함수, 클래스 로직 검증  
✔ 위젯 테스트 → `tester.pumpWidget()`을 사용하여 UI 동작 검증  
✔ 통합 테스트 → `integration_test`를 사용하여 실제 앱 동작 검증  
✔ Mock 테스트 → `mockito`를 활용하여 외부 API를 Mock으로 대체  
✔ CI/CD 자동화 → GitHub Actions를 활용하여 테스트 자동 실행  
