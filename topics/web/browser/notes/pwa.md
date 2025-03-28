# 📱 PWA (Progressive Web App)

**PWA (Progressive Web App)** 는 웹 기술을 사용해 만든 앱이지만, 네이티브 앱처럼 동작하도록 하는 웹 애플리케이션입니다.  
브라우저에서 실행되지만, 설치도 가능하고 오프라인에서도 동작할 수 있어 사용자 경험을 크게 향상시킵니다.

---

## 1️⃣ PWA란?

웹 앱이지만 네이티브 앱처럼 설치 가능하고, 빠르고, 오프라인에서도 동작하는 앱

- HTML, CSS, JS 기반
- 홈 화면에 앱처럼 설치 가능
- 오프라인 캐시 및 백그라운드 처리 지원
- 푸시 알림, 전체화면 등 다양한 기능 지원

---

## 2️⃣ 네이티브 앱 vs 웹 앱 vs PWA

| 항목 | 네이티브 앱 | 웹 앱 | PWA |
|------|------|------|------|
| 설치 | ✅ 필요 | ❌ 없음 | ✅ 선택적 |
| 오프라인 사용 |	✅ 가능 | ❌ 불가능 | ✅ 가능 |
| 성능 | 🔥 최고 |	😐 보통 |	💡 개선됨 |
| 접근성 | 기기 기능 완전 접근 |	거의 불가 |	일부 접근 가능 |
| 배포 방식 | 앱스토어 | URL | 앱스토어 또는 URL |


---

## 3️⃣ 구성 요소

### 1) Web App Manifest
- 앱 이름, 아이콘, 시작 URL 등 앱 메타데이터 설정
- JSON 파일 형식
- 브라우저가 앱으로 인식하는 핵심 요소

#### 예시
```
{
  "name": "My Cool App",
  "short_name": "CoolApp",
  "start_url": "/",
  "display": "standalone",
  "background_color": "#ffffff",
  "icons": [
    {
      "src": "/icons/icon-192.png",
      "sizes": "192x192",
      "type": "image/png"
    }
  ]
}
```

---

### 2) Service Worker
- 백그라운드에서 실행되는 JS 파일
- 네트워크 요청 가로채기, 캐싱, 푸시 알림 등 담당
- 비동기 이벤트 기반

#### `Service Worker` 등록 예시
```
 if ('serviceWorker' in navigator) {
  navigator.serviceWorker.register('/sw.js')
    .then(() => console.log('Service Worker registered!'));
}
```

---

## 4️⃣ 설치 흐름

1. 사용자가 PWA 지원 브라우저에서 사이트 접속
2. Manifest와 Service Worker가 정상 등록됨
3. "앱 설치" 배너 또는 + 버튼 노출
4. 사용자가 설치 선택 시 홈 화면에 앱 등록
5. 다음부터는 앱처럼 전체 화면에서 실행됨

---

## 5️⃣ 오프라인 동작 예시

- Service Worker가 페이지 요청을 캐싱
- 오프라인 시 캐시에서 페이지 제공
- 네트워크 연결 회복 시 백그라운드 동기화 가능

---

## 6️⃣ display 모드 종류

| 모드 | 설명 |
|------|------|
| standalone | 앱처럼 보여주고 브라우저 UI 제거 |
| fullscreen | 완전한 전체화면 (상태바도 제거) |
| minimal-ui | 주소창 최소화된 형태 |
| browser | 일반 브라우저로 동작 |

---

## 7️⃣ 브라우저 지원

| 브라우저 | 지원 여부 |
|----------|-----------|
| Chrome | ✅ 완전 지원 |
| Edge | ✅ 완전 지원 |
| Firefox | ✅ 대부분 지원 |
| Safari (iOS) | 🔸 일부 제한 (오프라인/알림 불완전) |
| Android WebView | ❌ 미지원 |

---

## 🎯 요약

✔ PWA는 웹 기술로 만든 앱이지만 **설치, 캐싱, 알림** 등의 **네이티브 기능**을 제공  
✔ 핵심 구성요소는 **Manifest** 와 **Service Worker**  
✔ 사용자 경험 향상, 성능 개선, 오프라인 지원이 핵심 장점  
✔ 모든 기능을 위해서는 **HTTPS 환경 필수**

