# Flutter 성능 최적화

Flutter 앱의 성능을 개선하기 위한 방법을 정리합니다.  
주요 최적화 대상: **렌더링 성능, 빌드 최적화, 애니메이션 최적화, 네트워크 성능 개선** 등

---

## 1. `build()` 최적화 (위젯 렌더링 최소화)

### 1.1 `const` 키워드 사용하기
`const`를 사용하면 **불필요한 위젯 재빌드를 방지**할 수 있습니다.

```dart
const Text("이것은 정적인 텍스트");
```

✅ **불필요한 `build()` 호출 방지**  
✅ **앱 성능 개선에 도움**  

---

### 1.2 `StatelessWidget`을 최대한 활용하기
**상태가 변하지 않는 위젯**은 `StatefulWidget` 대신 `StatelessWidget`을 사용하세요.

```dart
class MyWidget extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Text("이 위젯은 상태가 필요 없습니다.");
  }
}
```

✅ `StatelessWidget`은 불필요한 상태 관리 비용을 줄임  

---

### 1.3 `ListView.builder()` 사용하기
대량의 리스트 데이터를 효율적으로 렌더링하려면 **`ListView.builder()`**를 사용하세요.

```dart
ListView.builder(
  itemCount: 1000,
  itemBuilder: (context, index) {
    return ListTile(title: Text("아이템 $index"));
  },
)
```

✅ **필요한 아이템만 빌드하여 메모리 절약**  
✅ `ListView.builder()` 사용 시 **스크롤 성능 개선**  

---

## 2. `setState()` 최소화

`setState()`를 호출하면 **해당 위젯과 자식 위젯이 모두 다시 빌드**됩니다.  
필요한 부분만 `setState()` 하도록 구조를 개선하세요.

### 2.1 `setState()`를 위젯 트리 상위에서 호출하지 않기 (비효율적인 예제)
```dart
setState(() {
  _counter++;
});
```

✔ **최적화된 구조**: `setState()`를 하위 위젯으로 분리  
✔ **상태를 효율적으로 관리**하기 위해 **Provider, Riverpod** 사용 고려  

---

## 3. 이미지 최적화

### 3.1 `cacheWidth`, `cacheHeight` 지정
큰 이미지를 직접 로드하는 대신, 적절한 크기로 캐시하면 **메모리 사용량 감소** 가능

```dart
Image.asset(
  'assets/image.png',
  cacheWidth: 200, // 200px로 크기 제한
  cacheHeight: 200,
)
```

### 3.2 `flutter_image_compress`를 사용한 이미지 압축
```yaml
dependencies:
  flutter_image_compress: ^1.1.0
```

```dart
import 'package:flutter_image_compress/flutter_image_compress.dart';

Future<void> compressImage(String path) async {
  final result = await FlutterImageCompress.compressAndGetFile(
    path, "$path_compressed.jpg",
    quality: 80,
  );
}
```

✅ **이미지 로딩 속도 향상**  
✅ **메모리 사용량 감소**  

---

## 4. 애니메이션 최적화

### 4.1 `AnimatedBuilder` 사용하기
애니메이션을 적용할 때 **위젯 전체가 다시 빌드되지 않도록** 최적화합니다.

```dart
class AnimatedBox extends StatefulWidget {
  @override
  _AnimatedBoxState createState() => _AnimatedBoxState();
}

class _AnimatedBoxState extends State<AnimatedBox> with SingleTickerProviderStateMixin {
  late AnimationController _controller;
  
  @override
  void initState() {
    super.initState();
    _controller = AnimationController(vsync: this, duration: Duration(seconds: 2))..repeat();
  }

  @override
  Widget build(BuildContext context) {
    return AnimatedBuilder(
      animation: _controller,
      builder: (context, child) {
        return Transform.rotate(
          angle: _controller.value * 6.28,
          child: child,
        );
      },
      child: Container(width: 100, height: 100, color: Colors.blue),
    );
  }
}
```

✅ **애니메이션이 변경될 때 `build()` 호출 최소화**  
✅ **애니메이션 성능 개선 가능**  

---

## 5. 네트워크 성능 개선

### 5.1 `http` 요청을 `compute()`로 최적화
백그라운드에서 JSON 파싱을 실행하여 **UI 스레드 성능 저하 방지**

```dart
import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:flutter/foundation.dart';

Future<List<dynamic>> fetchData() async {
  final response = await http.get(Uri.parse('https://jsonplaceholder.typicode.com/posts'));
  return compute(parseJson, response.body);
}

List<dynamic> parseJson(String responseBody) {
  return jsonDecode(responseBody);
}
```

✅ **`compute()`를 사용하여 백그라운드 스레드에서 JSON 처리**  
✅ **UI 프레임 드랍 방지**  

---

### 5.2 `dio`를 사용하여 효율적인 HTTP 요청 관리
```yaml
dependencies:
  dio: ^5.0.0
```

```dart
import 'package:dio/dio.dart';

final dio = Dio();

Future<void> fetchData() async {
  final response = await dio.get('https://jsonplaceholder.typicode.com/posts');
  print(response.data);
}
```

✅ `Dio`는 **자동 JSON 파싱, 네트워크 캐싱 기능 제공**  
✅ HTTP 요청 속도 최적화 가능  

---

## 6. 기타 성능 최적화 팁

### 6.1 `Tree Shake Icons` 활성화
사용하지 않는 아이콘을 제거하여 **앱 크기를 줄일 수 있음**

```yaml
flutter:
  fonts:
    - family: MaterialIcons
      fonts:
        - asset: fonts/MaterialIcons-Regular.otf
```

✅ **앱 크기 최적화 가능**  

---

### 6.2 `isolate`를 활용한 멀티스레딩
복잡한 연산을 별도의 `isolate`에서 처리하여 **UI 프레임 드랍 방지**

```dart
import 'dart:isolate';

Future<void> computeHeavyTask() async {
  final receivePort = ReceivePort();
  await Isolate.spawn(doHeavyWork, receivePort.sendPort);
}

void doHeavyWork(SendPort sendPort) {
  // 무거운 연산 수행
}
```

✅ **CPU 사용량이 높은 작업을 백그라운드에서 실행 가능**  

---

## 7. 성능 분석 도구 활용

### 7.1 Flutter DevTools
```sh
flutter pub global activate devtools
flutter run --debug
```

✔ **위젯 리빌드 검사**  
✔ **프레임 속도 및 메모리 사용량 분석 가능**  

---

### 7.2 `debugProfileBuildsEnabled`
`debugProfileBuildsEnabled`를 설정하면 **어떤 위젯이 자주 다시 빌드되는지 확인 가능**

```dart
void main() {
  debugProfileBuildsEnabled = true;
  runApp(MyApp());
}
```

✔ **불필요한 `build()` 호출 감지 가능**  

---

## 8. 결론

| 최적화 방법 | 설명 |
|-------------|------------------------------|
| `const` 키워드 | 불필요한 위젯 재빌드 방지 |
| `setState()` 최소화 | 불필요한 위젯 빌드 방지 |
| `ListView.builder()` | 효율적인 리스트 렌더링 |
| 이미지 최적화 | `cacheWidth`, `flutter_image_compress` |
| 애니메이션 최적화 | `AnimatedBuilder` 사용 |
| 네트워크 최적화 | `compute()`, `Dio` 사용 |
| 성능 분석 도구 | `Flutter DevTools`, `debugProfileBuildsEnabled` |
