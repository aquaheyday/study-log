# 📌 프로젝트 설정

Flutter를 시작하기 위해 필요한 환경 설정 및 프로젝트 생성 방법을 정리합니다.

---

## 1. Flutter 설치

Flutter 개발을 위해 **Flutter SDK**를 설치해야 합니다.

### Flutter SDK 다운로드
Flutter 공식 사이트에서 최신 버전을 다운로드합니다.

🔗 [Flutter 설치 가이드](https://docs.flutter.dev/get-started/install)

### OS별 설치 방법

#### Windows
```sh
# Chocolatey 패키지 매니저를 사용한 설치
choco install flutter
```
✔ 설치 후 환경 변수를 설정해야 합니다.

#### macOS
```sh
# Homebrew를 사용한 설치
brew install flutter
```
✔ iOS 개발을 위해 **Xcode**와 **CocoaPods**도 설치해야 합니다.
```sh
sudo gem install cocoapods
```

#### Linux
```sh
# Snap 패키지 매니저를 사용한 설치
sudo snap install flutter --classic
```

### Flutter 환경 변수 설정 (Windows)
Flutter를 설치한 후, 환경 변수를 설정해야 합니다.
1. **Flutter SDK 경로**를 `PATH`에 추가 (`C:\flutter\bin`)
2. `flutter doctor` 명령어 실행하여 설정 확인

---

## 2. Flutter 개발 환경 설정

### 필수 개발 도구
- **Flutter SDK**
- **Dart SDK** (Flutter에 포함)
- **Android Studio 또는 Visual Studio Code**
- **Xcode (macOS에서 iOS 개발 시 필수)**

### Flutter 상태 점검
설치가 완료되었는지 확인하려면 아래 명령어를 실행합니다.
```sh
flutter doctor
```
출력 예시:
```
Doctor summary (to see all details, run flutter doctor -v):
[✓] Flutter (Channel stable, 3.10.0, on macOS 12.5 64-bit)
[✓] Android toolchain - develop for Android devices
[✓] Xcode - develop for iOS and macOS
[✓] Chrome - develop for the web
[✓] Visual Studio Code (version 1.78.0)
[✓] Connected device (1 available)
```
⚠️ 에러가 있는 경우, Flutter가 제시하는 해결 방법을 따르면 됩니다.

---

## 3. Flutter 프로젝트 생성

Flutter 프로젝트를 생성하려면 아래 명령어를 실행합니다.

```sh
flutter create my_app
cd my_app
flutter run
```

✔ `flutter create my_app` → 새로운 Flutter 프로젝트를 `my_app`이라는 이름으로 생성합니다.  
✔ `cd my_app` → 프로젝트 폴더로 이동합니다.  
✔ `flutter run` → 프로젝트를 실행합니다. 실행 후  `http://localhost:12345` 또는 에뮬레이터에서 기본 앱을 확인할 수 있습니다.  

---

## 4. 프로젝트 폴더 구조

Flutter 프로젝트가 생성되면 아래와 같은 구조를 가집니다.

```
my_app/
 ├── android/         # Android 네이티브 코드
 ├── ios/             # iOS 네이티브 코드
 ├── lib/             # 앱의 주요 코드 (Dart 파일)
 │   ├── main.dart    # 앱의 진입점
 ├── test/            # 테스트 코드
 ├── pubspec.yaml     # 프로젝트 설정 및 의존성 관리
 ├── README.md        # 프로젝트 설명 파일
```

✔ 주요 폴더 및 파일 설명
- `lib/main.dart` → 앱의 진입점 (앱 실행 코드)
- `pubspec.yaml` → 패키지 및 의존성 관리
- `android/`, `ios/` → 네이티브 코드 포함

---

## 5. 개발 환경 설정

### 실행 가능한 디바이스 확인
Flutter에서 실행 가능한 기기를 확인하려면 다음 명령어를 입력합니다.

```sh
flutter devices
```

출력 예시:
```
1 connected device:

• iPhone 13 (mobile)  • ios      • com.apple.CoreSimulator.SimRuntime.iOS-15-5
```

### 실행 명령어
Flutter 프로젝트를 실행하려면:

```sh
flutter run
```
✔ 연결된 에뮬레이터 또는 실제 기기에서 앱이 실행됩니다.

특정 플랫폼에서 실행하려면:
```sh
flutter run -d chrome      # 웹 실행
flutter run -d ios         # iOS 실행
flutter run -d android     # Android 실행
```

---

## 6. IDE 설정 및 플러그인 설치

### VS Code 설정
1. VS Code에서 **Flutter & Dart 플러그인** 설치  
2. `Ctrl + Shift + P` → **Flutter: New Project** 실행  
3. `main.dart`를 열고 실행 (`F5`)

### Android Studio 설정
1. **Flutter 플러그인** 및 **Dart 플러그인** 설치  
2. **AVD Manager**에서 Android Emulator 설정  
3. 프로젝트를 실행 (`Shift + F10`)

---

## 7. 디버깅 및 Hot Reload

### Hot Reload

코드 변경 사항을 즉시 반영하려면:

```sh
r
```

✔ `flutter run` 실행 중 `r`을 입력

### Hot Restart

전체 애플리케이션을 다시 실행하려면:

```sh
R
```

✔ `flutter run` 실행 중 `R`을 입력

---

## 8. 패키지 추가 및 관리

### 패키지 추가 (`pubspec.yaml` 수정)
예를 들어, HTTP 요청을 위한 `http` 패키지를 추가하려면:

```yaml
dependencies:
  flutter:
    sdk: flutter
  http: ^0.13.5
```

### 패키지 설치
```sh
flutter pub get
```

### 패키지 업데이트
```sh
flutter pub upgrade
```

---

## 9. 앱 빌드 및 배포

### Android APK 빌드
```sh
flutter build apk
```
✔ `build/app/outputs/flutter-apk/app-release.apk` 에서 APK 확인 가능

### iOS 빌드
```sh
flutter build ios
```
✔ Mac 및 Xcode 필요

### 웹 빌드
```sh
flutter build web
```
✔ `build/web` 폴더에 생성됨

---

## 🎯 정리

✔ Flutter 설치 후 `flutter doctor`로 상태 확인  
✔ 새 Flutter 프로젝트 생성: `flutter create my_app`  
✔ 프로젝트 실행: `flutter run`  
✔ 핫 리로드(`r`), 핫 리스타트(`R`)로 빠른 개발 가능  
✔ 패키지 추가는 `pubspec.yaml`에서 설정  
✔ 앱 빌드는 `flutter build` 명령어 사용  
