# 🐛 iOS Safari에서 MP4 영상 재생 불가

## ⚠️ 문제 상황
- 발생 날짜: 2025-07-04  
- 발생 환경:
  - 기기: iPhone 15 Pro
  - 브라우저: Safari, Chrome (모바일)
  - 파일: `mp4`
  - HTML 태그:
    ```html
    <video preload="metadata" muted autoplay loop playsinline webkit-playsinline x5-playsinline poster="img/img_01.jpg">
      <source src="video/video_main.mp4" type="video/mp4">
    </video>
    ```
- 재현 방법  
  1. iOS Safari에서 해당 페이지 접속  
  2. 영상이 로딩 상태에서 멈추고 재생되지 않음  

---

## 🔍 원인 분석
- 영상 파일은 `H.264 + AAC` 코덱으로 iOS 호환됨  
- **moov atom(영상 메타데이터)** 이 파일 끝에 위치 → iOS Safari는 스트리밍 재생 불가  
- Safari는 moov atom 이 파일 시작에 있어야 재생 가능  

---

## 🛠 해결 방법
- `ffmpeg`를 사용해 MP4 구조 수정:
    ```bash
    ffmpeg -i video_main.mp4 -c:v libx264 -profile:v baseline -level 3.0 -pix_fmt yuv420p -c:a aac -movflags +faststart output.mp4
    ```
- `-movflags +faststart`로 moov atom을 파일 앞쪽으로 이동 (MP4 파일을 스트리밍/웹 재생에 최적화)
- iOS, Safari, Android 브라우저에서 재생 정상 확인  

---

## 🚀 결과
✅ iOS Safari에서 영상 재생 정상화  
✅ 모든 주요 브라우저 호환성 확보  
✅ 향후 대응책 마련:
- 배포 전 MP4 파일 `faststart` 여부 확인
- Safari 테스트 항목 체크리스트에 영상 재생 추가
