# 🛡️ Browser - `sandbox` & `iframe` 보안

## 1️⃣ `iframe`이란?

`iframe`(inline frame)은 **웹 페이지 안에 다른 HTML 문서를 삽입**할 수 있는 HTML 요소입니다.

#### 사용 예시

```html
<iframe src="https://example.com"></iframe>
```

✔ 하지만 외부 콘텐츠를 로드할 경우 **보안 취약점(예: XSS, 클릭재킹)** 을 유발할 수 있기 때문에 **제어 수단이 필요**합니다.  

---

## 2️⃣ `sandbox` 속성이란?

`iframe`에 **제한적인 권한만 부여**해서 **더 안전하게 외부 콘텐츠를 포함**할 수 있도록 해주는 HTML 속성입니다.

```html
<iframe src="https://example.com" sandbox></iframe>
```

✔ `sandbox`를 사용하면 iframe 내부 페이지는 **기본적으로 모든 권한이 차단**됩니다.  

---

## 3️⃣ `sandbox` 기본 제약사항

#### `sandbox` 속성만 명시하면, iframe 내부 페이지는 다음과 같은 기능이 **모두 차단**됩니다

| 차단되는 기능 | 설명 |
|---------------|------|
| 스크립트 실행 | `<script>` 작동 안함 |
| 폼 제출 | `form` 전송 불가 |
| 플러그인 사용 | Flash, Java 등 불가능 |
| 팝업 창 열기 | `window.open` 차단 |
| 부모 접근 | `window.parent` 접근 차단 |
| 자동 실행 미디어 | 오디오/비디오 자동재생 불가 |

---

## 4️⃣ `sandbox` 허용 옵션

`sandbox` 속성에 **세부 권한을 명시**해서 일부 기능을 허용할 수 있습니다.  

```html
<iframe src="..." sandbox="allow-scripts allow-forms"></iframe>
```

| 속성 | 설명 |
|------|------|
| `allow-scripts` | 스크립트 실행 허용 (`eval`, `new Function` 등은 여전히 제한) |
| `allow-forms` | 폼 제출 허용 |
| `allow-same-origin` | 동일 출처처럼 동작 (주의! XSS 위험) |
| `allow-popups` | 팝업 창 열기 허용 |
| `allow-modals` | `alert`, `prompt`, `confirm` 허용 |
| `allow-presentation` | 프레젠테이션 모드 사용 (예: 전체화면) |
| `allow-downloads` | 사용자 상호작용을 통한 다운로드 허용 |
| `allow-top-navigation` | 최상위 창으로 리디렉션 허용 |

---

## 5️⃣ 보안 팁 🧩

| 상황 | 권장 설정 |
|------|-----------|
| 외부 콘텐츠 신뢰 불가 | `sandbox` 필수 사용 |
| 신뢰된 내부 콘텐츠 | `sandbox` 생략 or `allow-same-origin` 사용 가능 |
| 로그인/세션 포함 콘텐츠 | `allow-same-origin` + CORS 정책 철저 관리 |
| 광고/제3자 위젯 | `sandbox="allow-scripts"` 정도로 제한 권한만 부여 |

---

## 6️⃣ 안전한 광고 `iframe` 예시

```html
<iframe src="https://ads.partner.com/ad.html" sandbox="allow-scripts"></iframe>
```

✔ 광고는 보통 `allow-scripts`만 허용  
❌ `allow-same-origin`은 절대 사용 X → 광고가 부모 페이지에 접근할 수 있으므로 **심각한 보안 위협**  

---

## 🎯 정리

✔ `iframe`은 외부 콘텐츠를 삽입할 수 있지만 **보안상 주의가 필요**  
✔ `sandbox` 속성은 iframe에 **권한 제한을 부여**  
✔ 기본은 **모든 기능 차단**, 옵션으로 필요한 권한만 **하나씩 허용**  
✔ 외부 콘텐츠일수록 **더 강력하게 제한**해야 안전
