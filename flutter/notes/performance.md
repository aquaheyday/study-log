# 🛠️ 앱 최적화 가이드

Flutter 앱의 성능을 최적화하는 방법을 정리합니다.  
앱이 빠르고 원활하게 실행되도록 **렌더링, 메모리 관리, 네트워크 성능** 등을 개선하는 방법을 알아봅니다.

---

## 1. Flutter 앱 최적화의 주요 영역

| 최적화 영역 | 설명 |
|------------|------|
| 렌더링 성능 | 화면을 빠르게 그리도록 최적화 |
| 빌드 성능 | 위젯 트리의 불필요한 빌드 방지 |
| 메모리 관리 | 불필요한 객체 생성 줄이기 |
| 네트워크 최적화 | API 호출 속도 및 데이터 캐싱 |
| 패키지 최적화 | 불필요한 패키지 제거 및 코드 크기 감소 |

---

## 2. 렌더링 성능 최적화

### `const` 키워드 사용
`const`를 사용하면 불필요한 위젯 리빌드를 방지할 수 있습니다.

```dart
class MyWidget extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return const Text("Hello, Flutter!");
  }
}
```

✔ `const` 키워드를 사용하면 변경되지 않는 위젯을 캐싱하여 **렌더링 성능**을 향상  

---

### `RepaintBoundary`를 사용하여 불필요한 리렌더링 방지

```dart
RepaintBoundary(
  child: Image.network("https://example.com/image.jpg"),
)
```

✔ `RepaintBoundary` → UI 업데이트 시 특정 부분만 다시 그림  

---

### `Opacity` 대신 `Visibility` 사용

```dart
Visibility(
  visible: true, // false로 설정하면 화면에서 숨김
  child: Text("보이거나 숨길 수 있는 텍스트"),
)
```

✔ `Opacity`는 보이지 않아도 계속 렌더링되지만, `Visibility`는 **렌더링 자체를 방지**  

---

## 3. 빌드 성능 최적화

### `const` 생성자를 활용한 위젯 재사용

```dart
class MyButton extends StatelessWidget {
  const MyButton({Key? key}) : super(key: key);
}
```

✔ `const` 생성자를 사용하면 **불필요한 위젯 재생성 방지**  

---

### `ListView.builder` 사용
화면에 보이는 항목만 렌더링하여 **메모리 사용을 줄임**.

```dart
ListView.builder(
  itemCount: 1000,
  itemBuilder: (context, index) {
    return ListTile(title: Text("아이템 $index"));
  },
)
```

✔ `ListView.builder` → **동적 리스트 최적화**  

---

### `AutomaticKeepAliveClientMixin`으로 리스트 상태 유지

```dart
class MyListView extends StatefulWidget {
  @override
  _MyListViewState createState() => _MyListViewState();
}

class _MyListViewState extends State<MyListView> with AutomaticKeepAliveClientMixin {
  @override
  bool get wantKeepAlive => true;

  @override
  Widget build(BuildContext context) {
    super.build(context);
    return ListView.builder(
      itemCount: 100,
      itemBuilder: (context, index) => ListTile(title: Text("Item $index")),
    );
  }
}
```

✔ **페이지 전환 시 상태 유지** → 리스트 스크롤 위치가 초기화되지 않음  

---

## 4. 메모리 관리 최적화

### `dispose()`를 활용하여 리소스 해제

```dart
class MyWidget extends StatefulWidget {
  @override
  _MyWidgetState createState() => _MyWidgetState();
}

class _MyWidgetState extends State<MyWidget> {
  late TextEditingController _controller;

  @override
  void initState() {
    super.initState();
    _controller = TextEditingController();
  }

  @override
  void dispose() {
    _controller.dispose(); // 메모리 해제
    super.dispose();
  }
}
```

✔ `dispose()` → 사용이 끝난 객체를 해제하여 **메모리 누수 방지**  

---

### `image_cache.clear()`로 불필요한 이미지 캐시 삭제

```dart
void clearCache() {
  imageCache.clear();
  imageCache.clearLiveImages();
}
```

✔ `imageCache.clear()` → 불필요한 **이미지 캐시 제거**  

---

## 5. 네트워크 성능 최적화

### `http` 패키지 대신 `dio` 사용 (더 빠른 API 호출)

```dart
import 'package:dio/dio.dart';

final dio = Dio();

Future<void> fetchData() async {
  final response = await dio.get('https://jsonplaceholder.typicode.com/posts/1');
  print(response.data);
}
```

✔ `dio`는 `http`보다 **빠르고 효율적인 네트워크 요청** 지원  

---

### API 응답 캐싱 (`dio_cache_interceptor` 활용)

```sh
flutter pub add dio_cache_interceptor
```

```dart
import 'package:dio/dio.dart';
import 'package:dio_cache_interceptor/dio_cache_interceptor.dart';

final dio = Dio();
final cacheOptions = CacheOptions(store: MemCacheStore());

void setupDio() {
  dio.interceptors.add(DioCacheInterceptor(options: cacheOptions));
}
```

✔ **API 응답을 캐싱하여 불필요한 요청 방지**  

---

## 6. 패키지 및 코드 최적화

### 사용하지 않는 패키지 제거

```sh
flutter pub outdated
flutter pub remove <패키지명>
```

✔ 불필요한 패키지를 삭제하여 **앱 크기 및 빌드 시간 단축**  

---

### 코드 난독화 및 앱 크기 줄이기

```sh
flutter build apk --release --split-per-abi
flutter build ios --release
```

✔ `--split-per-abi` → **앱 크기 최적화** (안드로이드)  
✔ `flutter build ios --release` → **iOS 최적화 빌드**  

---

### `flutter analyze`를 활용한 코드 정리

```sh
flutter analyze
```

✔ **불필요한 코드 및 오류 탐색**  

---

## 7. 앱 실행 성능 최적화

### `Flutter DevTools` 활용

```sh
flutter pub global activate devtools
flutter run --profile
```

✔ `devtools` → **실시간 성능 모니터링**  
✔ `profile mode` → **퍼포먼스 분석 가능**  

---

### `Isolate`를 사용하여 백그라운드 작업 수행

```dart
import 'dart:isolate';

void heavyTask() {
  print("백그라운드 작업 실행");
}

void main() {
  Isolate.spawn(heavyTask, null);
}
```

✔ 백그라운드에서 무거운 작업 실행하여 메인 스레드 최적화  

---

## 🎯 최적화 체크리스트

✔ 렌더링 성능 최적화 → `const` 사용, `RepaintBoundary` 활용  
✔ 위젯 빌드 최적화 → `ListView.builder` 사용, 불필요한 리빌드 방지  
✔ 메모리 관리 최적화 → `dispose()` 사용, `imageCache.clear()` 활용  
✔ 네트워크 최적화 → `dio` 사용, API 응답 캐싱  
✔ 코드 최적화 → 불필요한 패키지 제거, 난독화 적용  
✔ 실행 성능 최적화 → `Flutter DevTools`로 성능 모니터링  
