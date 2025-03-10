# Flutter 기타 유용한 팁

Flutter 개발을 하면서 유용하게 사용할 수 있는 다양한 팁을 정리했습니다.  
**핫 리로드, 다크 모드, 반응형 UI, 키보드 숨기기, 로깅, 성능 분석** 등을 포함합니다.

---

## 1. 개발 속도를 높이는 팁

### 1.1 핫 리로드 & 핫 리스타트 차이점
| 기능 | 설명 |
|------|---------------------------|
| **Hot Reload** | 상태 유지한 채 UI만 새로고침 (빠름) |
| **Hot Restart** | 앱을 다시 시작 (상태 초기화됨) |

📌 **Hot Reload 사용 방법**  
터미널에서 **`r`** 키를 누르면 적용됩니다.

---

## 2. 다크 모드 지원

Flutter에서 다크 모드를 지원하려면 `themeMode`를 설정합니다.

```dart
MaterialApp(
  themeMode: ThemeMode.system, // 시스템 설정에 따라 변경
  theme: ThemeData.light(),    // 밝은 테마
  darkTheme: ThemeData.dark(), // 어두운 테마
);
```

✅ **다크 모드를 자동으로 적용 가능**  
✅ `ThemeMode.light` 또는 `ThemeMode.dark`를 설정하여 강제 적용 가능  

---

## 3. 반응형 UI 구현

Flutter에서 **반응형 UI**를 만들려면 `MediaQuery` 또는 `LayoutBuilder`를 사용합니다.

### 3.1 `MediaQuery`를 사용한 반응형 UI
```dart
double screenWidth = MediaQuery.of(context).size.width;

Widget build(BuildContext context) {
  return screenWidth > 600 ? TabletLayout() : MobileLayout();
}
```

### 3.2 `LayoutBuilder` 사용 예제
```dart
LayoutBuilder(
  builder: (context, constraints) {
    if (constraints.maxWidth > 600) {
      return TabletLayout();
    } else {
      return MobileLayout();
    }
  },
)
```

✅ **다양한 화면 크기에 대응 가능**  
✅ **웹 및 태블릿 환경에서도 최적화 가능**  

---

## 4. 키보드 숨기기

사용자가 입력을 완료하면 **키보드를 숨기는 기능**을 추가할 수 있습니다.

```dart
FocusScope.of(context).unfocus(); // 키보드 숨기기
```

또는 `GestureDetector`를 사용하여 **화면을 터치하면 키보드가 사라지도록 설정**할 수 있습니다.

```dart
GestureDetector(
  onTap: () {
    FocusScope.of(context).unfocus();
  },
  child: Scaffold(
    body: Center(child: Text("화면을 터치하면 키보드가 사라집니다.")),
  ),
)
```

✅ **텍스트 필드에서 다른 곳을 터치하면 키보드가 닫힘**  
✅ **사용자 경험(UX) 개선 가능**  

---

## 5. 앱 로딩 화면 (Splash Screen)

Flutter 기본 스플래시 화면을 설정하려면 **flutter_native_splash** 패키지를 사용합니다.

### 5.1 `flutter_native_splash` 패키지 설치
```yaml
dependencies:
  flutter_native_splash: ^2.2.0
```

### 5.2 설정 파일 (`pubspec.yaml`)
```yaml
flutter_native_splash:
  color: "#ffffff"
  image: assets/splash.png
```

📌 설정 적용  
```sh
flutter pub run flutter_native_splash:create
```

✅ **앱이 실행될 때 로딩 화면을 보여줄 수 있음**  

---

## 6. 디버깅 및 로깅

### 6.1 `print()` 대신 `debugPrint()` 사용
```dart
debugPrint("긴 문자열도 잘 출력됩니다.", wrapWidth: 1024);
```

✅ **긴 로그를 자동으로 줄여줌**  

---

### 6.2 `dart:developer`의 `log()`
```dart
import 'dart:developer';

log("디버깅 메시지", name: "MyApp");
```

✅ **필터링이 가능한 디버깅 로그 제공**  

---

## 7. 애니메이션 쉽게 적용하기

Flutter에서 기본 애니메이션을 사용할 때 `AnimatedContainer`, `AnimatedOpacity` 등을 활용하면 편리합니다.

### 7.1 `AnimatedContainer` 예제
```dart
AnimatedContainer(
  duration: Duration(seconds: 1),
  width: isExpanded ? 200 : 100,
  height: 100,
  color: Colors.blue,
)
```

### 7.2 `AnimatedOpacity` 예제
```dart
AnimatedOpacity(
  opacity: isVisible ? 1.0 : 0.0,
  duration: Duration(milliseconds: 500),
  child: Text("페이드 인/아웃 효과"),
)
```

✅ **애니메이션을 쉽게 적용 가능**  

---

## 8. 상태 관리 팁

Flutter에서 상태 관리를 할 때 **Provider** 또는 **Riverpod**을 사용하면 성능을 최적화할 수 있습니다.

### 8.1 `Provider` 예제
```dart
ChangeNotifierProvider(
  create: (context) => CounterProvider(),
  child: MyApp(),
)
```

✅ **전역 상태 공유 가능**  
✅ **`setState()` 최소화 가능**  

---

## 9. 앱 크기 줄이기

### 9.1 사용하지 않는 아이콘 제거
```yaml
flutter:
  fonts:
    - family: MaterialIcons
      fonts:
        - asset: fonts/MaterialIcons-Regular.otf
```

### 9.2 `flutter clean` 실행
```sh
flutter clean
```

✅ **불필요한 파일 제거 후 빌드 가능**  

---

## 10. Flutter 앱 배포

### 10.1 Android APK 빌드
```sh
flutter build apk --release
```

### 10.2 iOS 앱 빌드
```sh
flutter build ios --release
```

✅ **배포를 위한 최적화된 빌드 생성 가능**  

---

## 11. 결론

| 기능 | 방법 |
|------|------------------------------|
| 핫 리로드 | `r` 키 입력 |
| 다크 모드 지원 | `ThemeMode.system` 사용 |
| 반응형 UI | `MediaQuery`, `LayoutBuilder` 활용 |
| 키보드 숨기기 | `FocusScope.of(context).unfocus()` |
| 로깅 | `debugPrint()`, `log()` 사용 |
| 애니메이션 | `AnimatedContainer`, `AnimatedOpacity` 사용 |
| 상태 관리 | `Provider`, `Riverpod` 사용 |
| 앱 크기 줄이기 | `flutter clean` 실행 |
| 앱 배포 | `flutter build apk --release` |
