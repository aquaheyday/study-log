# 🐦 Flutter

이 폴더는 **Flutter 학습 과정**에서 정리한 자료, 예제 코드, 프로젝트를 보관하는 공간입니다.  
기본 개념부터 실습 예제, 프로젝트 구성까지 단계적으로 학습한 내용을 체계적으로 정리합니다.

---

## 📂 디렉토리 구성

| 폴더명 | 설명 |
|---|---|
| [notes](./notes) | 개념 정리 및 학습 노트 |
| [examples](./examples) | 주요 기능별 예제 코드 |

---

## 📋 개념 정리 목록  

### 📌 기본 개념  
| 주제 | 파일명 | 설명 |  
|---|---|---|
| Flutter 개요 | [flutter-intro.md](./notes/flutter-intro.md) | Flutter란 무엇인가? 특징과 장점 |  
| 프로젝트 설정 | [setup-flutter.md](./notes/setup-flutter.md) | Flutter SDK 설치, 프로젝트 생성 |  
| Dart 기본 문법 | [dart-basics.md](./notes/dart-basics.md) | 변수, 함수, 클래스, 비동기 프로그래밍 |  
| 위젯 개요 | [widgets-intro.md](./notes/widgets-intro.md) | Stateless vs Stateful 위젯 개념 |  
| 상태 관리 | [state-management.md](./notes/state-management.md) | setState, InheritedWidget 기초 |  

### 🔲 UI 구성 및 위젯  
| 주제 | 파일명 | 설명 |  
|---|---|---|
| 레이아웃 기본 | [layout-basics.md](./notes/layout-basics.md) | Column, Row, Stack, Expanded 사용법 |  
| 리스트와 스크롤 | [listview.md](./notes/listview.md) | ListView, GridView, ScrollController |  
| 폼과 입력 필드 | [forms.md](./notes/forms.md) | TextField, Form, validation |  
| 버튼과 이벤트 처리 | [buttons.md](./notes/buttons.md) | ElevatedButton, GestureDetector |  
| 네비게이션 | [navigation.md](./notes/navigation.md) | Navigator, push/pop, named routes |  

### 🔄 상태 관리  
| 주제 | 파일명 | 설명 |  
|---|---|---|
| Provider 기본 | [provider-basics.md](./notes/provider-basics.md) | Provider를 활용한 상태 관리 |  
| Riverpod | [riverpod.md](./notes/riverpod.md) | Riverpod을 활용한 상태 관리 |  
| Bloc 패턴 | [bloc-pattern.md](./notes/bloc-pattern.md) | flutter_bloc 라이브러리 사용법 |  
| GetX 상태 관리 | [getx.md](./notes/getx.md) | GetX의 상태 관리와 라우팅 기능 |  

### 🌍 네트워크 및 데이터 관리  
| 주제 | 파일명 | 설명 |  
|---|---|---|
| HTTP 통신 | [http-request.md](./notes/http-request.md) | Dio, http 패키지를 이용한 API 호출 |  
| JSON 데이터 처리 | [json-parsing.md](./notes/json-parsing.md) | json_serializable을 활용한 파싱 |  
| 로컬 데이터 저장 | [local-storage.md](./notes/local-storage.md) | SharedPreferences, Hive, SQLite |  
| Firebase 연동 | [firebase.md](./notes/firebase.md) | Firebase Auth, Firestore, Storage |  

### 🚀 고급 개념  
| 주제 | 파일명 | 설명 |  
|---|---|---|
| 애니메이션 기본 | [animations.md](./notes/animations.md) | AnimatedBuilder, Hero, Lottie 사용 |  
| 커스텀 위젯 만들기 | [custom-widgets.md](./notes/custom-widgets.md) | 재사용 가능한 위젯 개발 |  
| 플랫폼별 기능 정리 | [platform-channels.md](./notes/platform-channels.md) | MethodChannel을 활용한 네이티브 연동 |  
| 멀티플랫폼 개발 가이드 | [flutter-web.md](./notes/flutter-web.md) | Flutter Web 및 데스크톱 지원 |  

### 🛠️ 테스트 및 배포  
| 주제 | 파일명 | 설명 |  
|---|---|---|
| 테스트 가이드 | [testing.md](./notes/testing.md) | Widget 테스트, Unit 테스트 |  
| 앱 최적화 가이드 | [performance.md](./notes/performance.md) | 렌더링 최적화, 코드 스플리팅 |  
| 앱 배포 가이드 | [deployment.md](./notes/deployment.md) | Android, iOS, 웹 배포 (Play Store, App Store) |  

---

## 📋 예제 목록

| 주제 | 파일명 | 설명 |
|---|---|---|
| 기본 레이아웃 | [basic_layout.dart](./examples/basic_layout.dart) | `Button`, `Column`, `Container` 기본 구성 |
| 카운터 앱 | [counter_app.dart](./examples/counter_app.dart) | `StatefulWidget`을 활용한 카운터 앱 |
| HTTP 요청 | [http_request.dart](./examples/http_request.dart) | REST API 요청(GET) 예제 |
| 폼 유효성 검사 | [form_validation.dart](./examples/form_validation.dart) | `TextFormField`와 `Form`을 사용한 입력값 유효성 검사 |
| 갤러리/카메라 이미지 선택 | [image_picker.dart](./examples/image_picker.dart) | 갤러리에서 이미지 선택하기 또는 카메라에서 사진 촬영 가능 |

---

## 📚 참고 자료
- [Flutter 공식 문서](https://docs.flutter.dev/) - Flutter의 기본 개념부터 심화된 내용까지 공식 가이드 제공  
- [Flutter API 문서](https://api.flutter.dev/) - Flutter 위젯 및 API 문서  
- [Dart 언어 문서](https://dart.dev/guides) - Flutter에서 사용하는 Dart 프로그래밍 언어 가이드  

---

## 📢 업데이트 로그
- 2025-03-07: 초기 구성
- 2025-03-10: notes 초기 내용
- 2025-03-14: notes 내용 추가
- 2025-03-15: notes 내용 추가
- 2025-03-19: 전체 구성 수정 및 project, example 내용 추가
- 2025-03-21: notes 템플릿 수정
