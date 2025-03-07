# 🏗️ Flutter Examples

이 폴더는 **Flutter 학습 과정에서 작성한 기능별 예제 코드**를 모아둔 공간입니다.  
각 예제는 독립적으로 실행할 수 있는 형태로 작성되며, 학습 노트(`notes/`)와 연계하여 참고할 수 있습니다.  
Flutter 기본 위젯부터, 네트워크 통신, 상태관리 패턴까지 다양한 주제를 다룹니다.

---

## 📋 예제 목록

| 번호 | 주제 | 파일명 | 설명 |
|---|---|---|---|
| 01 | 기본 레이아웃 | [basic_layout.dart](./basic_layout.dart) | Row, Column, Container 기본 구성 |
| 02 | 버튼 및 이벤트 핸들링 | [button_event.dart](./button_event.dart) | 버튼 클릭 이벤트 처리 |
| 03 | 리스트뷰 구성 | [listview_example.dart](./listview_example.dart) | 리스트뷰 데이터 표시 |
| 04 | 상태관리 기본 | [simple_state.dart](./simple_state.dart) | setState를 활용한 상태관리 |
| 05 | Provider 상태관리 | [provider_example.dart](./provider_example.dart) | Provider 패턴 적용 예제 |
| 06 | 폼 입력 처리 | [form_example.dart](./form_example.dart) | TextField, Form 위젯 사용법 |
| 07 | 네트워크 요청 | [http_example.dart](./http_example.dart) | Dio 활용 REST API 요청 |
| 08 | 애니메이션 기본 | [animation_example.dart](./animation_example.dart) | 애니메이션 컨트롤러 기초 |
| 09 | 네비게이션 관리 | [navigation_example.dart](./navigation_example.dart) | Navigator 활용 화면 이동 |
| 10 | 테마 설정 | [theme_example.dart](./theme_example.dart) | 라이트/다크 모드 적용 |

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
