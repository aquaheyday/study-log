# Flutter 플러그인 활용

Flutter에서 플러그인은 **기본 기능을 확장하여 네이티브 기능(Android/iOS)을 사용**할 수 있도록 도와줍니다.  
예를 들어, **네트워크, 데이터 저장, 위치 서비스, 카메라, 푸시 알림** 등을 플러그인으로 쉽게 구현할 수 있습니다.

---

## 1. 플러그인 추가 방법

Flutter에서 플러그인을 사용하려면 **pub.dev**에서 원하는 플러그인을 찾아 `pubspec.yaml`에 추가합니다.

### 1.1 플러그인 설치
```yaml
dependencies:
  flutter:
    sdk: flutter
  http: ^0.13.0  # 네트워크 요청을 위한 http 패키지
  shared_preferences: ^2.0.0  # 로컬 저장을 위한 플러그인
```

📌 **설치 후 터미널에서 실행**
```sh
flutter pub get
```

---

## 2. 주요 Flutter 플러그인

### 2.1 네트워크 요청 (`http`)
HTTP 요청을 보낼 때 `http` 패키지를 사용합니다.

#### 설치
```yaml
dependencies:
  http: ^0.13.0
```

#### 사용 예제
```dart
import 'package:http/http.dart' as http;
import 'dart:convert';

Future<void> fetchData() async {
  final response = await http.get(Uri.parse('https://jsonplaceholder.typicode.com/posts/1'));

  if (response.statusCode == 200) {
    var data = jsonDecode(response.body);
    print('Title: ${data['title']}');
  } else {
    throw Exception('Failed to load data');
  }
}
```

✅ **REST API와 통신 가능**  
✅ **GET, POST, PUT, DELETE 요청 지원**  

---

### 2.2 로컬 저장 (`shared_preferences`)
**앱의 간단한 데이터를 저장**할 때 사용합니다.

#### 설치
```yaml
dependencies:
  shared_preferences: ^2.0.0
```

#### 사용 예제
```dart
import 'package:shared_preferences/shared_preferences.dart';

Future<void> saveData() async {
  final prefs = await SharedPreferences.getInstance();
  await prefs.setString('username', 'flutter_dev');
}

Future<void> loadData() async {
  final prefs = await SharedPreferences.getInstance();
  String? username = prefs.getString('username');
  print('저장된 값: $username');
}
```

✅ **앱 종료 후에도 데이터 유지 가능**  
✅ **간단한 설정값 저장에 적합 (ex. 토큰, 유저 설정값)**  

---

### 2.3 파일 저장 (`path_provider`)
**파일을 저장하거나 불러올 때 사용**합니다.

#### 설치
```yaml
dependencies:
  path_provider: ^2.0.11
```

#### 사용 예제
```dart
import 'package:path_provider/path_provider.dart';
import 'dart:io';

Future<String> getFilePath() async {
  final directory = await getApplicationDocumentsDirectory();
  return '${directory.path}/my_file.txt';
}

Future<void> writeToFile(String content) async {
  final path = await getFilePath();
  final file = File(path);
  await file.writeAsString(content);
  print("파일 저장 완료: $path");
}
```

✅ **로컬 파일 저장 및 읽기 가능**  
✅ **네이티브 파일 시스템 활용 가능**  

---

### 2.4 위치 정보 (`geolocator`)
**현재 GPS 위치를 가져오거나 거리 계산할 때 사용**합니다.

#### 설치
```yaml
dependencies:
  geolocator: ^9.0.2
```

#### 사용 예제
```dart
import 'package:geolocator/geolocator.dart';

Future<void> getCurrentLocation() async {
  Position position = await Geolocator.getCurrentPosition(
      desiredAccuracy: LocationAccuracy.high);
  print("위치: ${position.latitude}, ${position.longitude}");
}
```

✅ **GPS 및 위치 서비스 활용 가능**  
✅ **거리 계산 기능 제공**  

---

### 2.5 카메라 (`camera`)
**카메라를 사용하여 사진을 촬영하거나, 실시간 미리보기를 제공**할 수 있습니다.

#### 설치
```yaml
dependencies:
  camera: ^0.10.0
```

#### 사용 예제
```dart
import 'package:camera/camera.dart';

late List<CameraDescription> cameras;
late CameraController controller;

Future<void> initCamera() async {
  cameras = await availableCameras();
  controller = CameraController(cameras[0], ResolutionPreset.high);
  await controller.initialize();
}
```

✅ **실시간 카메라 미리보기 가능**  
✅ **사진 및 동영상 촬영 가능**  

---

### 2.6 푸시 알림 (`firebase_messaging`)
**Firebase Cloud Messaging(FCM)을 통해 푸시 알림을 받을 수 있습니다.**

#### 설치
```yaml
dependencies:
  firebase_messaging: ^14.0.0
  firebase_core: ^2.0.0
```

#### 사용 예제
```dart
import 'package:firebase_messaging/firebase_messaging.dart';

Future<void> setupPushNotifications() async {
  FirebaseMessaging messaging = FirebaseMessaging.instance;

  // 푸시 알림 권한 요청
  await messaging.requestPermission();

  // FCM 토큰 가져오기
  String? token = await messaging.getToken();
  print("FCM 토큰: $token");
}
```

✅ **앱이 백그라운드/포그라운드 상태에서도 알림 수신 가능**  
✅ **iOS 및 Android에서 푸시 알림 지원**  

---

### 2.7 바코드/QR 스캔 (`qr_code_scanner`)
**QR 코드 및 바코드 스캔 기능을 구현할 수 있습니다.**

#### 설치
```yaml
dependencies:
  qr_code_scanner: ^1.0.0
```

#### 사용 예제
```dart
import 'package:qr_code_scanner/qr_code_scanner.dart';

final GlobalKey qrKey = GlobalKey(debugLabel: 'QR');
late QRViewController controller;

Widget build(BuildContext context) {
  return QRView(
    key: qrKey,
    onQRViewCreated: (QRViewController controller) {
      this.controller = controller;
      controller.scannedDataStream.listen((scanData) {
        print('QR 코드 값: ${scanData.code}');
      });
    },
  );
}
```

✅ **QR 코드 및 바코드 스캔 가능**  
✅ **카메라를 활용하여 실시간 인식 가능**  

---

## 3. 결론

| 플러그인 | 기능 |
|----------|------------------------------|
| `http` | REST API 통신 |
| `shared_preferences` | 간단한 로컬 저장 |
| `path_provider` | 파일 저장 및 읽기 |
| `geolocator` | GPS 위치 정보 |
| `camera` | 카메라 사용 |
| `firebase_messaging` | 푸시 알림 기능 |
| `qr_code_scanner` | QR 코드 스캔 |
