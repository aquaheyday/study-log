# 🚀 멀티플랫폼 개발 가이드

Flutter는 **모바일(Android, iOS), 웹, 데스크톱(Windows, macOS, Linux), 임베디드** 플랫폼을 하나의 코드베이스로 개발할 수 있는 프레임워크입니다.  
멀티플랫폼 개발 시 고려해야 할 사항과 설정 방법을 정리합니다.

---

## 1. Flutter 멀티플랫폼 지원

| 플랫폼 | 지원 상태 | 주요 특징 |
|--------|---------|----------|
| Android | ✅ 완전 지원 | 네이티브 성능 제공 |
| iOS | ✅ 완전 지원 | 일부 네이티브 기능 제한 |
| Web | ✅ 지원 | 일부 패키지 미지원 |
| Windows | ✅ 지원 | 독립 실행형 EXE 배포 가능 |
| macOS | ✅ 지원 | Mac 앱으로 배포 가능 |
| Linux | ✅ 지원 | 데스크톱 GUI 앱 가능 |
| 임베디드 | ⚠️ 일부 지원 | Raspberry Pi 등 특정 기기 지원 |

✔ **하나의 코드베이스**로 다양한 플랫폼에서 실행 가능  
✔ `flutter run -d <platform>`으로 실행 가능  

---

## 2. 프로젝트 설정

### 프로젝트 생성
Flutter에서 멀티플랫폼 프로젝트를 생성하려면 다음 명령어를 사용합니다.

```sh
flutter create my_app
```

### 멀티플랫폼 지원 활성화
Flutter는 기본적으로 Android와 iOS를 지원하지만, 웹과 데스크톱을 활성화해야 합니다.

```sh
flutter config --enable-web
flutter config --enable-macos-desktop
flutter config --enable-windows-desktop
flutter config --enable-linux-desktop
```

✔ `flutter doctor`를 실행하여 환경을 확인하세요.

---

## 3. 플랫폼별 실행 방법

| 플랫폼 | 실행 명령어 |
|--------|-----------|
| Android | `flutter run -d android` |
| iOS | `flutter run -d ios` |
| Web | `flutter run -d chrome` |
| Windows | `flutter run -d windows` |
| macOS | `flutter run -d macos` |
| Linux | `flutter run -d linux` |

✔ 실행 가능한 디바이스 목록 확인:  
```sh
flutter devices
```

---

## 4. 플랫폼별 UI 대응

### `Platform.is`를 사용한 분기 처리 (`dart:io`)
`dart:io` 패키지를 사용하여 플랫폼별 UI를 다르게 설정할 수 있습니다.

```dart
import 'dart:io';
import 'package:flutter/material.dart';

class PlatformText extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Text(
      Platform.isAndroid ? "안드로이드" :
      Platform.isIOS ? "iOS" :
      Platform.isWindows ? "Windows" :
      Platform.isMacOS ? "macOS" :
      Platform.isLinux ? "Linux" :
      "Unknown Platform",
    );
  }
}
```

✔ `Platform.isAndroid`, `Platform.isIOS` 등으로 플랫폼 감지 가능  
✔ `dart:io`는 웹에서 지원되지 않음  

---

### `Theme.of(context).platform`을 활용한 플랫폼 감지

```dart
import 'package:flutter/material.dart';

class PlatformAwareButton extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    final platform = Theme.of(context).platform;

    return ElevatedButton(
      onPressed: () {},
      child: Text(platform == TargetPlatform.android ? "Android 버튼" : "iOS 버튼"),
    );
  }
}
```

✔ `Theme.of(context).platform` → `TargetPlatform` 값을 반환  

---

## 5. 패키지 호환성 확인

### 패키지별 멀티플랫폼 지원 여부
Flutter에서 사용할 패키지가 특정 플랫폼에서 지원되는지 확인해야 합니다.

```sh
flutter pub outdated
```

### 웹과 호환되지 않는 패키지 예시

| 패키지 | Android/iOS | Web | 대체 가능 패키지 |
|--------|------------|-----|-----------------|
| `path_provider` | ✅ | ❌ | `universal_io` |
| `shared_preferences` | ✅ | ❌ | `flutter_secure_storage` |
| `firebase_messaging` | ✅ | ✅ | 웹은 일부 기능 제한 |
| `geolocator` | ✅ | ❌ | `location` |

✔ `pub.dev`에서 패키지 설명을 확인하고 호환성을 체크.   

---

## 6. 웹과 모바일의 차이점

| 기능 | 모바일(Android/iOS) | 웹 |
|------|----------------|----|
| 네이티브 코드 | ✅ | ❌ |
| 푸시 알림 | ✅ | 제한적 |
| 파일 저장 | ✅ | ❌ |
| 위치 정보 | ✅ | ❌ |

✔ 웹에서는 네이티브 기능 사용이 제한됨  

---

## 7. `platform_channel` (네이티브 코드 연동)

플랫폼별로 네이티브 기능을 추가해야 할 경우 `MethodChannel`을 사용할 수 있습니다.

### Android에서 네이티브 코드 실행 예제: 

#### Flutter 코드
```dart
import 'package:flutter/services.dart';

const platform = MethodChannel('com.example/native');

Future<String> getNativeMessage() async {
  return await platform.invokeMethod('getMessage');
}
```

#### Android 네이티브 코드 (Kotlin)
```kotlin
class MainActivity: FlutterActivity() {
    private val CHANNEL = "com.example/native"

    override fun configureFlutterEngine(flutterEngine: FlutterEngine) {
        super.configureFlutterEngine(flutterEngine)
        MethodChannel(flutterEngine.dartExecutor.binaryMessenger, CHANNEL).setMethodCallHandler { call, result ->
            if (call.method == "getMessage") {
                result.success("Hello from Android!")
            }
        }
    }
}
```

✔ `MethodChannel`을 통해 네이티브 코드와 통신 가능  

---

## 8. 플랫폼별 빌드 및 배포

### Android 빌드
```sh
flutter build apk
flutter build appbundle
```

### iOS 빌드
```sh
flutter build ios
```

### 웹 빌드
```sh
flutter build web
```

### Windows 빌드
```sh
flutter build windows
```

### macOS 빌드
```sh
flutter build macos
```

### Linux 빌드
```sh
flutter build linux
```

✔ 배포 시 플랫폼별 설정 파일 (`android/`, `ios/`, `web/`)을 확인하세요.  

---

## 🎯 정리

✔ Android, iOS, 웹, 데스크톱 지원 → 하나의 코드로 모든 플랫폼 실행  
✔ 플랫폼별 UI 대응 → `Platform.isAndroid`, `Theme.of(context).platform` 활용  
✔ 패키지 호환성 확인 → `flutter pub outdated` 사용  
✔ 네이티브 연동 가능 → `MethodChannel`로 네이티브 코드 실행  
✔ 멀티플랫폼 빌드 가능 → `flutter build <platform>` 사용  
