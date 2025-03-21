# 🚀 Flutter 플랫폼별 기능 정리

Flutter는 **iOS, Android, 웹, 데스크톱(Windows, macOS, Linux)** 등 플랫폼별로 제공되는 기능과 차이점을 정리합니다.  

---

## 1️⃣ Flutter의 멀티 플랫폼 지원

| 플랫폼 | 지원 상태 | 주요 기능 |
|--------|---------|----------|
| Android | ✅ 완전 지원 | 모든 기능 사용 가능 |
| iOS | ✅ 완전 지원 | 일부 네이티브 기능 제한 |
| Web | ✅ 지원 | 일부 네이티브 API 제한 |
| Windows | ✅ 지원 | 파일 시스템, 네트워크 기능 사용 가능 |
| macOS | ✅ 지원 | iOS와 비슷한 UI/UX |
| Linux | ✅ 지원 | GUI 및 CLI 앱 개발 가능 |

✔ **코드 한 번 작성**으로 여러 플랫폼에서 실행 가능  
✔ `dart:io`, `dart:html` 등을 활용하여 플랫폼별 코드 작성 가능  

---

## 2️⃣ 플랫폼별 기능 차이

### 1) `path_provider` (파일 시스템 접근)

| 기능 | Android | iOS | Web | Windows | macOS | Linux |
|------|--------|----|----|----|----|----|
| 앱 데이터 저장 | ✅ | ✅ | ❌ | ✅ | ✅ | ✅ |
| 외부 저장소 접근 | ✅ | 제한적 | ❌ | ✅ | ✅ | ✅ |

#### `path_provider` 패키지 예시
```dart
import 'package:path_provider/path_provider.dart';

Future<void> getAppDirectory() async {
  final directory = await getApplicationDocumentsDirectory();
  print("App directory: ${directory.path}");
}
```

✔ Android, iOS, 데스크톱 지원  
✔ 웹에서는 `path_provider` 미지원  

---

### 2) `platform_channel` (네이티브 코드 실행)

| 기능 | Android | iOS | Web | Windows | macOS | Linux |
|------|--------|----|----|----|----|----|
| 네이티브 기능 호출 | ✅ | ✅ | ❌ | ✅ | ✅ | ✅ |

#### `MethodChannel` (플랫폼별 네이티브 코드 실행)

```dart
import 'package:flutter/services.dart';

const platform = MethodChannel('com.example/native');

Future<String> getNativeMessage() async {
  return await platform.invokeMethod('getMessage');
}
```

✔ `MethodChannel` → 네이티브 코드 실행 가능  
✔ 웹에서는 `MethodChannel` 미지원  

---

### 3) `device_info_plus` (기기 정보)

| 기능 | Android | iOS | Web | Windows | macOS | Linux |
|------|--------|----|----|----|----|----|
| 기기 모델 조회 | ✅ | ✅ | ❌ | ✅ | ✅ | ✅ |
| OS 버전 확인 | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |

#### `device_info_plus` 패키지 예시
```dart
import 'package:device_info_plus/device_info_plus.dart';

Future<void> getDeviceInfo() async {
  final deviceInfo = DeviceInfoPlugin();
  final info = await deviceInfo.androidInfo; // Android 정보 가져오기
  print("Device: ${info.model}");
}
```

✔ Android, iOS, 데스크톱 지원  
✔ 웹에서는 제한적 기능 제공  

---

### 4) `firebase_messaging` (푸시 알림)

| 기능 | Android | iOS | Web | Windows | macOS | Linux |
|------|--------|----|----|----|----|----|
| Firebase 푸시 | ✅ | ✅ | ✅ | ❌ | ❌ | ❌ |

#### `firebase_messaging` 패키지 사용 예시

```dart
import 'package:firebase_messaging/firebase_messaging.dart';

Future<void> setupFirebase() async {
  FirebaseMessaging messaging = FirebaseMessaging.instance;
  String? token = await messaging.getToken();
  print("Firebase Token: $token");
}
```

✔ 모바일 및 웹에서 사용 가능  
✔ Windows, macOS, Linux에서는 미지원  

---

### 5) `geolocator` (위치 정보)

| 기능 | Android | iOS | Web | Windows | macOS | Linux |
|------|--------|----|----|----|----|----|
| GPS 위치 조회 | ✅ | ✅ | ❌ | ✅ | ✅ | ✅ |

#### `geolocator` 패키지 사용 예시

```dart
import 'package:geolocator/geolocator.dart';

Future<void> getLocation() async {
  Position position = await Geolocator.getCurrentPosition();
  print("Latitude: ${position.latitude}, Longitude: ${position.longitude}");
}
```

✔ Android, iOS, 데스크톱 지원  
✔ 웹에서는 미지원  

---

### 6) `local_auth` (생체 인증)

| 기능 | Android | iOS | Web | Windows | macOS | Linux |
|------|--------|----|----|----|----|----|
| 지문 인증 | ✅ | ✅ | ❌ | ✅ | ✅ | ❌ |
| Face ID | ❌ | ✅ | ❌ | ❌ | ✅ | ❌ |

#### `local_auth` 패키지 사용 예시

```dart
import 'package:local_auth/local_auth.dart';

Future<void> authenticateUser() async {
  final auth = LocalAuthentication();
  bool isAuthenticated = await auth.authenticate(
    localizedReason: '지문으로 로그인하세요',
  );
  print("Authentication result: $isAuthenticated");
}
```

✔ 생체 인증 지원 (Windows, macOS 일부 지원)  
✔ 웹에서는 미지원  

---

### 7) `adaptive_theme` (다크 모드)

| 기능 | Android | iOS | Web | Windows | macOS | Linux |
|------|--------|----|----|----|----|----|
| 시스템 다크 모드 감지 | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |

#### `adaptive_theme` 패키지 사용 예시

```dart
import 'package:adaptive_theme/adaptive_theme.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  final savedThemeMode = await AdaptiveTheme.getThemeMode();
  runApp(MyApp(savedThemeMode: savedThemeMode));
}
```

✔ 모든 플랫폼에서 다크 모드 감지 가능  

---

## 🎯 정리

| 기능 | Android | iOS | Web | Windows | macOS | Linux |
|------|--------|----|----|----|----|----|
| 파일 저장 | ✅ | ✅ | ❌ | ✅ | ✅ | ✅ |
| 네이티브 코드 실행 | ✅ | ✅ | ❌ | ✅ | ✅ | ✅ |
| 기기 정보 | ✅ | ✅ | 제한적 | ✅ | ✅ | ✅ |
| 푸시 알림 | ✅ | ✅ | ✅ | ❌ | ❌ | ❌ |
| 위치 정보 | ✅ | ✅ | ❌ | ✅ | ✅ | ✅ |
| 생체 인증 | ✅ | ✅ | ❌ | ✅ | ✅ | ❌ |
| 다크 모드 | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
