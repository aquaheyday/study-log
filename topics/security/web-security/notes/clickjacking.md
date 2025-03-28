# 🎯 Clickjacking (클릭재킹)

Clickjacking은 사용자가 클릭하고 있다고 생각하는 UI가 실제로는 **숨겨진 악성 요소**를 클릭하게 만드는 공격 기법입니다.  
보통 `iframe`을 이용해 **투명하게 덮어씌운 UI**를 보여주고, 사용자를 속여 특정 동작(좋아요, 결제, 설정 변경 등)을 유도합니다.

---

## 1️⃣ 정의

사용자가 실제로 클릭하는 대상이 보이는 것과 다른 UI 요소가 되도록 속이는 공격 기법

---

## 2️⃣ 동작 방식

1. 공격자가 자신의 웹사이트에 **희생자의 사이트를 iframe으로 삽입**
2. CSS로 해당 iframe을 **투명하게 만들고**, 정밀히 정렬
3. 버튼 등 유도할 영역 위에 클릭을 유도하는 요소 배치
4. 사용자가 보이는 UI를 클릭하면 **실제로는 iframe 안의 콘텐츠가 클릭됨**

---

## 3️⃣ 예시 시나리오

1. 사용자가 `evil.com` 접속
2. 그 안에 `bank.com/transfer` 페이지가 투명 iframe으로 로드됨
3. 사용자는 "게임 시작" 버튼을 클릭한다고 생각
4. 실제로는 "송금" 버튼을 클릭함 → 계좌 이체 실행

---

## 4️⃣ 코드 예시 (공격자 사이트)
```html
<style>
  iframe {
    position: absolute;
    top: 0;
    left: 0;
    width: 500px;
    height: 300px;
    opacity: 0; /* 투명 처리 */
    z-index: 10;
  }
  .fake-button {
    position: relative;
    z-index: 20;
  }
</style>

<iframe src="https://victim-site.com/transfer" ></iframe>
<button class="fake-button">🎮 게임 시작</button>
```

---

## 5️⃣ 공격 대상

- **SNS 버튼** (좋아요, 팔로우)
- **결제 / 계좌이체**
- **보안 설정 변경** (예: 2단계 인증 해제)
- **관리자 패널의 민감한 동작**

---

## 6️⃣ 방어 방법

| 방법 | 설명 |
|------|------|
| `X-Frame-Options` 헤더 설정 | iframe 삽입 자체를 차단 |
| `Content-Security-Policy: frame-ancestors` | iframe 삽입 가능한 출처 제한 |
| UI 확인용 토큰 사용 | 예: 클릭 전에 추가 인증 단계 삽입 |
| 클릭 이벤트에 사용자 확인 | 2단계 확인 또는 CAPTCHA 등 |

---

### 7️⃣ X-Frame-Options 예시
```
X-Frame-Options: DENY
X-Frame-Options: SAMEORIGIN
```

- `DENY` → 어떤 사이트에서도 iframe 삽입 차단  
- `SAMEORIGIN` → 같은 도메인만 iframe 허용

---

### 8️⃣ CSP 예시
```
Content-Security-Policy: frame-ancestors 'self'
```
- `'self'` → 현재 도메인만 iframe 허용  
- `'none'` → 어디서도 허용하지 않음

---

## 7️⃣ 사용자 방어법

- 신뢰되지 않은 사이트에서 무작정 클릭하지 않기
- 클릭을 유도하는 이벤트성 UI 주의
- 브라우저 보안 플러그인 활용

---

## 🎯 정리

✔ Clickjacking은 **iframe + CSS 조작**으로 사용자를 속이는 공격  
✔ 금융, SNS 등 **상태 변경 행위**를 속여 실행할 수 있음  
✔ `X-Frame-Options`, `CSP frame-ancestors` 로 효과적으로 방어 가능  
✔ **iframe 삽입 제한이 핵심 방어 전략**

