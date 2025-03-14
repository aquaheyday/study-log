# 📌 Flutter 개요

## 1. Flutter란 무엇인가?

**Flutter**는 Google에서 개발한 오픈소스 UI 프레임워크로, 하나의 코드베이스로 **iOS, Android, 웹, 데스크톱(Windows, macOS, Linux)** 애플리케이션을 개발할 수 있도록 지원합니다.

- **Dart 언어**를 기반으로 개발되며, 빠른 렌더링과 높은 성능을 제공합니다.
- **위젯 기반 프레임워크**로, UI를 구성하는 모든 요소를 위젯으로 표현합니다.
- **Hot Reload 기능**을 통해 코드 변경 사항을 즉시 확인할 수 있어 개발 속도가 빠릅니다.

---

## 2. Flutter의 특징과 장점

### ✅ 크로스 플랫폼 개발  
- 하나의 코드베이스로 여러 플랫폼(iOS, Android, Web, Desktop)에서 실행 가능  
- 유지보수 및 개발 비용 절감  

### ✅ 높은 성능  
- **Skia 그래픽 엔진**을 활용하여 빠른 렌더링 제공  
- 네이티브 ARM 코드로 컴파일되므로 네이티브 앱과 유사한 성능  

### ✅ 빠른 개발 속도  
- **Hot Reload**: 코드 수정 후 즉시 적용되어 UI 변화를 실시간으로 확인 가능  
- **풍부한 위젯 제공**: Material Design, Cupertino 위젯 포함  

### ✅ 강력한 UI 프레임워크  
- 커스텀 UI 제작이 용이하며, 애니메이션 및 테마 적용이 간편  
- 플랫폼에 구애받지 않고 일관된 디자인 구현 가능  

### ✅ 다양한 패키지와 플러그인 지원  
- 공식 패키지 및 커뮤니티 패키지 활용 가능 (pub.dev)  
- Firebase, REST API, SQLite, Google Maps 등의 다양한 플러그인 제공  

---

## 3. Flutter의 구조

Flutter는 3가지 주요 계층으로 구성됩니다.

### 1) **Flutter Framework**  
- UI를 구성하는 다양한 **위젯(Widgets)** 포함  
- `Material`, `Cupertino`, `Animation`, `Rendering`, `Gestures` 등의 모듈 제공  

### 2) **Flutter Engine**  
- **Dart VM**, Skia 그래픽 엔진, 네이티브 API와 상호작용  
- 빠른 렌더링 및 애니메이션 제공  

### 3) **Platform-specific Embedder**  
- iOS, Android, Web, Windows, macOS, Linux에서 실행되도록 지원  
- 각 플랫폼의 네이티브 기능과 통신  

---

## 4. Flutter와 다른 프레임워크 비교

| 특징 | Flutter | React Native | Swift/Kotlin |
|---|---|---|---|
| **언어** | Dart | JavaScript | Swift(iOS), Kotlin(Android) |
| **렌더링** | 자체 엔진 (Skia) | 네이티브 브리지 사용 | 네이티브 |
| **성능** | 네이티브 수준 | 네이티브보다 약간 낮음 | 최고 성능 |
| **개발 속도** | 빠름 (Hot Reload) | 빠름 (Fast Refresh) | 느림 |
| **UI 구성** | 커스텀 UI 위젯 | 네이티브 컴포넌트 | 네이티브 컴포넌트 |
| **코드 재사용성** | 90% 이상 | 80% 이상 | 0% (각 플랫폼 별 개발) |

---

## 5. Flutter로 개발할 수 있는 앱

Flutter는 다양한 유형의 애플리케이션을 개발할 수 있습니다.

- **모바일 앱**: Android, iOS 앱 개발 (예: Google Ads, eBay Motors)  
- **웹 애플리케이션**: Flutter Web을 활용한 PWA 개발  
- **데스크톱 애플리케이션**: Windows, macOS, Linux 지원  
- **임베디드 시스템**: 자동차, IoT 기기에도 적용 가능  

---

## 6. Flutter 주요 사용 사례

- **Google**: Google Pay, Google Ads, Google Assistant  
- **Alibaba**: 전자상거래 앱  
- **BMW**: 자동차 애플리케이션  
- **eBay Motors**: 중고차 거래 앱  
- **Reflectly**: 인기 저널링 앱  

---

## 7. Flutter 개발을 위한 준비 사항

### 필수 설치 도구  
1. **Flutter SDK** ([설치 가이드](https://docs.flutter.dev/get-started/install))  
2. **Dart SDK** (Flutter와 함께 설치됨)  
3. **개발 환경**  
   - VS Code + Flutter Extension  
   - Android Studio / IntelliJ IDEA  
4. **Android Emulator 또는 iOS Simulator**  

### 프로젝트 생성 명령어  
```sh
flutter create my_app
cd my_app
flutter run
```

---

## 8. Flutter 학습 자료

- 📚 공식 문서: [https://docs.flutter.dev](https://docs.flutter.dev)  
- 🎥 유튜브 강의: [Flutter 공식 채널](https://www.youtube.com/c/flutterdev)  
- 📦 패키지 검색: [pub.dev](https://pub.dev)  
- 🏛️ 커뮤니티: [Flutter Dev Forum](https://flutter.dev/community)  

---
🚀 **이제 Flutter를의 기본 개념을 이해했으니, 실습을 시작해 보세요!**  
다음 개념: [프로젝트 설정](./setup-flutter.md) →
