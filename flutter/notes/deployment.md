# 🚀 Flutter 앱 배포 가이드

Flutter 앱을 Android 및 iOS 스토어에 배포하고 최적화하는 과정을 설명합니다.  

---

## 1. Flutter 배포 개요

| 플랫폼 | 배포 파일 형식 | 배포 방법 |
|--------|------------|----------|
| Android | `.apk`, `.aab` | Google Play Store 또는 수동 배포 |
| iOS | `.ipa` | Apple App Store |

✔ Android → `.apk` (테스트), `.aab` (Play Store 배포)  
✔ iOS → `.ipa` (App Store 배포)  

---

## 2. Android 앱 배포 (Google Play Store)

### 키스토어(KeyStore) 생성

```sh
keytool -genkey -v -keystore my-release-key.jks -keyalg RSA -keysize 2048 -validity 10000 -alias my-key-alias
```

✔ KeyStore 파일 생성 (`android/app/` 폴더에 저장)      
✔ 암호 및 키 정보는 안전하게 보관  

---

### `key.properties` 설정 (`android/` 폴더)

`android/key.properties` 파일 생성 후 아래 내용 추가

```properties
storePassword=<your-password>
keyPassword=<your-password>
keyAlias=my-key-alias
storeFile=my-release-key.jks
```

✔ **비밀번호와 키 정보는 Git에 저장하지 않도록 주의**  

---

### `gradle`에 서명 설정 (`android/app/build.gradle`)

```gradle
android {
    ...
    signingConfigs {
        release {
            storeFile file("../my-release-key.jks")
            storePassword project.property("storePassword")
            keyAlias project.property("keyAlias")
            keyPassword project.property("keyPassword")
        }
    }
    buildTypes {
        release {
            signingConfig signingConfigs.release
        }
    }
}
```

✔ `storeFile`, `storePassword`, `keyAlias`, `keyPassword`를 `key.properties`에서 불러오기  

---

### `.aab` 빌드 및 Play Store 업로드

```sh
flutter build appbundle
```

✔ `app-release.aab` 파일 생성 (`build/app/outputs/bundle/release/`)  
✔ Google Play Console에서 앱 업로드  

---

### `.apk` 빌드 (테스트용)

```sh
flutter build apk --release
```

✔ `app-release.apk` 파일 생성 (`build/app/outputs/flutter-apk/`)  

---

## 3. iOS 앱 배포 (Apple App Store)

### Apple Developer 계정 생성

- [Apple Developer Program](https://developer.apple.com/) 가입  
- Xcode 및 Apple ID 설정  

---

### iOS 빌드 설정 (`ios/Runner.xcodeproj`)

```sh
open ios/Runner.xcworkspace
```

✔ Xcode에서 `Runner` 프로젝트 열기  
✔ `General` → `Bundle Identifier` 설정  
✔ `Signing & Capabilities` → `Automatically manage signing` 활성화  

---

### `Podfile` 업데이트 및 빌드

```sh
cd ios
pod install
flutter build ios --release
```

✔ 실제 기기에서 실행하려면 개발자 계정 필요  

---

### Xcode에서 앱 아카이브 후 App Store Connect 업로드

1. Xcode에서 `Product` → `Archive` 클릭  
2. `Distribute App` 선택  
3. `App Store Connect`로 업로드  

✔ App Store에 업로드 후 TestFlight 또는 App Store 배포 가능  

---

## 4. Firebase App Distribution (테스트 버전 배포)

### Firebase CLI 설치

```sh
npm install -g firebase-tools
firebase login
```

### Firebase 프로젝트 설정

```sh
firebase init
```

✔ Firebase 프로젝트에 연결  

---

### APK/AAB 업로드 (Android)

```sh
firebase appdistribution:distribute build/app/outputs/flutter-apk/app-release.apk \
  --app <FIREBASE_APP_ID> \
  --release-notes "새로운 기능 업데이트" \
  --groups testers
```

✔ 테스트 버전 배포 후 테스터 이메일로 초대 가능  

---

## 5. 웹 앱 배포 (Flutter Web)

### 웹 빌드 실행

```sh
flutter build web
```

✔ `build/web/` 폴더 생성  

---

### Firebase Hosting에 배포

```sh
firebase init hosting
firebase deploy
```

✔ Firebase Hosting을 사용하여 빠르게 웹 배포 가능  

---

## 6. 배포 후 앱 업데이트 및 유지보수

### Android & iOS 버전 관리 (`pubspec.yaml`)

```yaml
version: 1.0.0+1
```

✔ `1.0.0` → 앱 버전  
✔ `+1` → 빌드 넘버  

---

### OTA 업데이트 (CodePush 대안)

- **Android:** `In-App Update API` 사용 가능  
- **iOS:** 앱스토어 자동 업데이트  

✔ Flutter 웹 앱은 즉시 배포 가능하지만, 모바일 앱은 앱스토어 심사 필요  

---

## 🎯 정리

✔ Android 배포** → `.aab` 생성 후 Play Store 업로드 (KeyStore 생성 및 서명 설정 필수)
✔ iOS 배포 → Xcode에서 `Archive` 후 App Store 업로드 (Apple Developer 계정 필요)  
✔ 테스트 배포 → Firebase App Distribution 활용  
✔ Flutter Web 배포 → Firebase Hosting 또는 GitHub Pages 사용  
✔ 앱 업데이트 관리 →`pubspec.yaml` 버전 변경 후 배포  
