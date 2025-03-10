# 🏗️ Flutter Examples

이 폴더는 **Flutter 학습 과정에서 작성한 기능별 예제 코드**를 모아둔 공간입니다.  
각 예제는 독립적으로 실행할 수 있는 형태로 작성되며, 학습 노트(`notes/`)와 연계하여 참고할 수 있습니다.  
Flutter 기본 위젯부터, 네트워크 통신, 상태관리 패턴까지 다양한 주제를 다룹니다.

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

## 📝 작성 가이드
- 파일명은 `스네이크 케이스`로 작성 (ex: `basic_layout.dart`)
- 각 예제는 반드시 `main.dart` 형태로 구성하여, 단독 실행 가능하게 작성
- 주석을 충분히 작성하여, 예제의 핵심 포인트를 명확하게 설명
- 필요 시 관련 노트(`notes/`)와 상호 링크 추가
- 공통적으로 필요한 위젯이나 유틸성 코드는 `common/` 폴더로 분리 가능

---

## 📚 참고 자료
- [Flutter 공식 문서](https://docs.flutter.dev/)
- [Flutter by Example](https://flutterbyexample.com/)
- [Awesome Flutter](https://github.com/Solido/awesome-flutter)
