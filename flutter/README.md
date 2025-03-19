# 🐦 Flutter

이 폴더는 **Flutter 학습 과정**에서 정리한 자료, 예제 코드, 프로젝트를 보관하는 공간입니다.  
기본 개념부터 실습 예제, 프로젝트 구성까지 단계적으로 학습한 내용을 체계적으로 정리합니다.

---

## 📂 디렉토리 구성

| 폴더명 | 설명 |
|---|---|
| [notes](./notes) | 개념 정리 및 학습 노트 |
| [examples](./examples) | 주요 기능별 예제 코드 (컴포넌트, 상태관리 등) |
| [projects](./projects) | 미니 프로젝트 및 실습 프로젝트 |

---

## 📋 개념 정리 목록  

### 📌 기본 개념  
| 번호 | 주제 | 파일명 | 설명 |  
|---|---|---|---|  
| 01 | Flutter 개요 | [flutter-intro.md](./notes/flutter-intro.md) | Flutter란 무엇인가? 특징과 장점 |  
| 02 | 프로젝트 설정 | [setup-flutter.md](./notes/setup-flutter.md) | Flutter SDK 설치, 프로젝트 생성 |  
| 03 | Dart 기본 문법 | [dart-basics.md](./notes/dart-basics.md) | 변수, 함수, 클래스, 비동기 프로그래밍 |  
| 04 | 위젯 개요 | [widgets-intro.md](./notes/widgets-intro.md) | Stateless vs Stateful 위젯 개념 |  
| 05 | 상태 관리 | [state-management.md](./notes/state-management.md) | setState, InheritedWidget 기초 |  

### 🔲 UI 구성 및 위젯  
| 번호 | 주제 | 파일명 | 설명 |  
|---|---|---|---|  
| 06 | 레이아웃 기본 | [layout-basics.md](./notes/layout-basics.md) | Column, Row, Stack, Expanded 사용법 |  
| 07 | 리스트와 스크롤 | [listview.md](./notes/listview.md) | ListView, GridView, ScrollController |  
| 08 | 폼과 입력 필드 | [forms.md](./notes/forms.md) | TextField, Form, validation |  
| 09 | 버튼과 이벤트 처리 | [buttons.md](./notes/buttons.md) | ElevatedButton, GestureDetector |  
| 10 | 네비게이션 | [navigation.md](./notes/navigation.md) | Navigator, push/pop, named routes |  

### 🔄 상태 관리  
| 번호 | 주제 | 파일명 | 설명 |  
|---|---|---|---|  
| 11 | Provider 기본 | [provider-basics.md](./notes/provider-basics.md) | Provider를 활용한 상태 관리 |  
| 12 | Riverpod | [riverpod.md](./notes/riverpod.md) | Riverpod을 활용한 상태 관리 |  
| 13 | Bloc 패턴 | [bloc-pattern.md](./notes/bloc-pattern.md) | flutter_bloc 라이브러리 사용법 |  
| 14 | GetX 상태 관리 | [getx.md](./notes/getx.md) | GetX의 상태 관리와 라우팅 기능 |  

### 🌍 네트워크 및 데이터 관리  
| 번호 | 주제 | 파일명 | 설명 |  
|---|---|---|---|  
| 15 | HTTP 통신 | [http-request.md](./notes/http-request.md) | Dio, http 패키지를 이용한 API 호출 |  
| 16 | JSON 데이터 처리 | [json-parsing.md](./notes/json-parsing.md) | json_serializable을 활용한 파싱 |  
| 17 | 로컬 데이터 저장 | [local-storage.md](./notes/local-storage.md) | SharedPreferences, Hive, SQLite |  
| 18 | Firebase 연동 | [firebase.md](./notes/firebase.md) | Firebase Auth, Firestore, Storage |  

### 🚀 고급 개념  
| 번호 | 주제 | 파일명 | 설명 |  
|---|---|---|---|  
| 19 | 애니메이션 기본 | [animations.md](./notes/animations.md) | AnimatedBuilder, Hero, Lottie 사용 |  
| 20 | 커스텀 위젯 만들기 | [custom-widgets.md](./notes/custom-widgets.md) | 재사용 가능한 위젯 개발 |  
| 21 | 플랫폼별 기능 정리 | [platform-channels.md](./notes/platform-channels.md) | MethodChannel을 활용한 네이티브 연동 |  
| 22 | 멀티플랫폼 개발 가이드 | [flutter-web.md](./notes/flutter-web.md) | Flutter Web 및 데스크톱 지원 |  

### 🛠️ 테스트 및 배포  
| 번호 | 주제 | 파일명 | 설명 |  
|---|---|---|---|  
| 23 | 테스트 가이드 | [testing.md](./notes/testing.md) | Widget 테스트, Unit 테스트 |  
| 24 | 앱 최적화 가이드 | [performance.md](./notes/performance.md) | 렌더링 최적화, 코드 스플리팅 |  
| 25 | 앱 배포 가이드 | [deployment.md](./notes/deployment.md) | Android, iOS, 웹 배포 (Play Store, App Store) |  

---

## 📋 예제 목록

| 번호 | 주제 | 파일명 | 설명 |
|---|---|---|---|
| 01 | 기본 레이아웃 | [basic_layout.dart](./basic_layout.dart) | `Button`, `Column`, `Container` 기본 구성 |
| 02 | 카운터 앱 | [counter_app.dart](./counter_app.dart) | `StatefulWidget`을 활용한 카운터 앱 |
| 03 | HTTP 요청 | [http_request.dart](./http_request.dart) | REST API 요청(GET) 예제 |
| 04 | 폼 유효성 검사 | [form_validation.dart](./form_validation.dart) | `TextFormField`와 `Form`을 사용한 입력값 유효성 검사 |
| 05 | 갤러리/카메라 이미지 선택 | [image_picker.dart](./image_picker.dart) | 갤러리에서 이미지 선택하기 또는 카메라에서 사진 촬영 가능 |

---

## 📋 프로젝트 목록

| 번호 | 프로젝트명 | 설명 | 링크 |
|---|---|---|---|
| 01 | 메뉴 주문 웹 | API 와 연동하여 방을 생성하여 개인별 메뉴를 접수하는 웹 사이트 | [menu-order-web-front](./projects/menu-order-web-front) |

---

## 📢 업데이트 로그
- 2025-03-07: 초기 구성
- 2025-03-10: notes 초기 내용
- 2025-03-14: notes 내용 추가
- 2025-03-15: notes 내용 추가
- 2025-03-19: 전체 구성 수정 및 project, example 내용 추가
- 새로운 예제 추가 시 목록 업데이트 예정
