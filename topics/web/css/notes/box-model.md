# 📦 CSS - 박스 모델(Box Model)

**박스 모델**은 모든 HTML 요소를 하나의 사각형 박스로 보고, 그 요소가 화면에서 **얼마만큼의 공간을 차지하는지 계산하는 기준**이 되는 모델입니다.

---

## 1️⃣ 박스 모델 구조

```
+----------------------------+ ← margin (바깥 여백)
|      ▲                   ▲ |
|      |    border         | |
|      ▼                   ▼ |
|  +---------------------+   ← padding (안쪽 여백)
|  |     content         |   ← 실제 내용
|  +---------------------+   
|                            |
+----------------------------+
```

✔ 안에서 바깥 순서: `content` → `padding` → `border` → `margin`  

---

## 2️⃣ 구성 요소 설명

| 구성 요소 | 설명 |
|-----------|------|
| `content` | 실제 콘텐츠 (텍스트, 이미지 등) |
| `padding` | 콘텐츠와 테두리 사이의 공간 (안쪽 여백) |
| `border` | 요소의 테두리 |
| `margin` | 요소와 다른 요소 사이의 공간 (바깥 여백) |

#### 예시

```css
.box {
  width: 200px;
  padding: 20px;
  border: 5px solid black;
  margin: 10px;
}
```

✔ 총 너비 계산: 200px (`content`) + 20px*2 (`padding`) + 5px*2 (`border`) = 250px  
✔ `margin`은 요소 외부 여백, 요소 간 거리에는 영향을 줌  

---

## 3️⃣ `box-sizing` 속성

### 1) 기본값: `content-box`

- `width`/`height`는 `content` 영역만 포함  
- padding, border는 **추가로 계산됨**

```css
.box {
  box-sizing: content-box; /* 기본값 */
}
```

---

### 2) 권장값: `border-box`

- `width`/`height`에 `padding`과 `border` 포함  
- **총 크기 계산이 쉬워짐 → 레이아웃 만들기 편함**

```css
.box {
  box-sizing: border-box;
}
```

✔ 실무에서는 `border-box`를 **초기화 스타일로 전역 설정**하는 경우가 많음  

#### 초기화 예시
```css
* {
  box-sizing: border-box;
}
```

---

## 🎯 정리

| 영역 | 설명 |
|------|------|
| content | 실제 내용 |
| padding | 콘텐츠와 테두리 사이 공간 |
| border | 테두리 |
| margin | 외부 여백 (다른 요소와의 거리) |

✔ `box-sizing: border-box` → 실무에서 더 많이 사용  
✔ 전체 요소 크기 계산할 땐 padding, border 포함 여부 확인  
✔ 모든 HTML 요소는 박스 모델의 구조를 따름  
