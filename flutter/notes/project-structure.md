# Flutter Notes 프로젝트 구조 정리

Flutter 프로젝트의 일반적인 구조와 각 폴더의 역할을 정리한 문서입니다.

## 프로젝트 구조

```text
flutter_notes/
│── android/               # 안드로이드 네이티브 코드 (Java/Kotlin)
│── ios/                   # iOS 네이티브 코드 (Swift/Objective-C)
│── lib/                   # 주요 애플리케이션 코드
│   ├── main.dart          # 앱 실행 진입점
│   ├── core/              # 핵심 로직 (예: API, 데이터 관리)
│   │   ├── models/        # 데이터 모델 클래스
│   │   ├── services/      # API 서비스 및 데이터 처리
│   │   ├── utils/         # 유틸리티 및 헬퍼 함수
│   ├── views/             # UI 화면 (페이지 및 위젯)
│   │   ├── home_screen.dart  # 홈 화면
│   │   ├── note_screen.dart  # 노트 상세 화면
│   ├── widgets/           # 재사용 가능한 위젯 모음
│── assets/                # 정적 파일 (이미지, 폰트 등)
│── test/                  # 테스트 코드
│── pubspec.yaml           # 프로젝트 설정 및 의존성 목록
│── README.md              # 프로젝트 설명
```
## 디렉토리 설명

- **`android/`** & **`ios/`**  
  - 네이티브 코드가 포함된 폴더입니다. 플랫폼별 설정 및 플러그인 관련 파일이 위치합니다.

- **`lib/`**  
  - Flutter 앱의 주요 코드가 포함됩니다.
  - `core/` 폴더는 데이터 및 비즈니스 로직을 담당합니다.
  - `views/` 폴더는 UI와 화면 관련 코드가 들어갑니다.
  - `widgets/` 폴더에는 재사용 가능한 UI 컴포넌트가 저장됩니다.

- **`assets/`**  
  - 앱에서 사용하는 이미지, 아이콘, 폰트 등의 정적 파일이 저장됩니다.
  - `pubspec.yaml`에 등록 후 사용해야 합니다.

- **`test/`**  
  - 단위 테스트 및 위젯 테스트가 포함됩니다.

- **`pubspec.yaml`**  
  - 프로젝트의 의존성 및 설정을 관리하는 파일입니다.

## 추가 고려 사항

- **상태 관리:**  
  - `Provider`, `Riverpod`, `Bloc` 등 원하는 상태 관리 패턴을 적용할 수 있습니다.
  
- **데이터 저장:**  
  - `SharedPreferences`, `Hive`, `SQLite` 등을 사용하여 노트 데이터를 저장할 수 있습니다.

- **네트워크 요청:**  
  - `Dio` 또는 `http` 패키지를 활용하여 API 요청을 처리할 수 있습니다.
