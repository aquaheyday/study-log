# 🐛 Flutter Troubleshooting

이 폴더는 Flutter 개발 중 발생한 **에러 및 문제 상황**과 그에 대한 원인 분석, 해결 방법을 정리하는 공간입니다.  
각 문제는 독립적인 파일로 관리하며, 아래 흐름에 맞춰 작성합니다:

**문제 상황 → 원인 분석 → 해결 방법 → 추가 학습 포인트**

---

## 📋 문제 목록

| 번호 | 문제 요약 | 파일명 | 링크 |
|---|---|---|---|
| 01 | 빌드 시 Flutter SDK 경로 오류 | [flutter-sdk-path-error.md](./flutter-sdk-path-error.md) | [바로가기](./flutter-sdk-path-error.md) |
| 02 | Hot Reload가 동작하지 않는 문제 | [hot-reload-not-working.md](./hot-reload-not-working.md) | [바로가기](./hot-reload-not-working.md) |
| 03 | Android 에뮬레이터 네트워크 불가 | [emulator-network-issue.md](./emulator-network-issue.md) | [바로가기](./emulator-network-issue.md) |
| 04 | iOS 빌드 시 Pod 관련 에러 | [ios-pod-error.md](./ios-pod-error.md) | [바로가기](./ios-pod-error.md) |
| 05 | FlutterFire 초기화 에러 | [flutterfire-init-error.md](./flutterfire-init-error.md) | [바로가기](./flutterfire-init-error.md) |

---

## 📑 작성 가이드
- 파일명은 `케밥 케이스`로 작성 (ex: `flutter-sdk-path-error.md`)
- 각 문제는 템플릿(`troubleshooting-template.md`) 기반으로 작성
- 에러 로그/캡처/메시지는 원본 그대로 포함
- 해결 과정에서 참고한 공식 문서/블로그 링크를 하단에 정리
- **재발 방지를 위한 교훈 및 팁**도 추가

---

## 📂 디렉토리 구성 예시
```text
troubleshooting/
├── README.md                              # 현재 파일
├── flutter-sdk-path-error.md              # 문제별 정리 파일
├── hot-reload-not-working.md
├── emulator-network-issue.md
├── ios-pod-error.md
└── flutterfire-init-error.md
```

---

## 📚 참고 자료
- [Flutter 공식 문서](https://docs.flutter.dev/)
- [Flutter Issue Tracker](https://github.com/flutter/flutter/issues)
- [Flutter Awesome 트러블슈팅 모음](https://github.com/Solido/awesome-flutter#troubleshooting)
