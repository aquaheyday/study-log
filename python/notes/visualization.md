# 📊 Python 데이터 시각화 (Data Visualization)

데이터를 효과적으로 분석하고 전달하기 위해 Python에서는 **Matplotlib**과 **Seaborn** 같은 시각화 라이브러리를 사용합니다.  
이 문서에서는 **기본적인 그래프부터 고급 스타일 설정까지** 다룹니다.

---

## 1. 데이터 시각화란?

- 데이터를 시각적으로 표현하여 **트렌드와 패턴을 쉽게 이해**할 수 있음.
- 데이터 분석, 보고서 작성, 머신러닝 결과 해석 등에 필수적.
- **주요 라이브러리**:
  - **Matplotlib** → 기본적인 그래프 제공
  - **Seaborn** → 고급 통계 그래프 지원
  - **Pandas** → DataFrame에서 직접 시각화 지원

---

## 2. 라이브러리 설치 및 불러오기

### 설치
```sh
pip install matplotlib seaborn
```

### 라이브러리 불러오기
```python
import matplotlib.pyplot as plt
import seaborn as sns
import numpy as np
import pandas as pd
```

✔ `plt` → **Matplotlib의 pyplot 모듈**  
✔ `sns` → **Seaborn 라이브러리**  

---

## 3. Matplotlib 기본 사용법

### 선 그래프(Line Plot) 그리기
```python
x = [1, 2, 3, 4, 5]
y = [10, 20, 25, 30, 40]

plt.plot(x, y, marker="o", linestyle="--", color="b", label="데이터 라인")
plt.xlabel("X축 라벨")
plt.ylabel("Y축 라벨")
plt.title("Matplotlib 선 그래프")
plt.legend()
plt.grid(True)
plt.show()
```
✔ `marker="o"` → 데이터 포인트에 원(circle) 표시  
✔ `linestyle="--"` → 점선 스타일 적용  
✔ `color="b"` → 파란색(blue) 설정  
✔ `plt.legend()` → 범례 표시  
✔ `plt.grid(True)` → 격자(grid) 추가  

---

### 막대 그래프(Bar Chart)
```python
categories = ["A", "B", "C", "D"]
values = [3, 7, 5, 8]

plt.bar(categories, values, color=["red", "blue", "green", "orange"])
plt.xlabel("카테고리")
plt.ylabel("값")
plt.title("Matplotlib 막대 그래프")
plt.show()
```

✔ `plt.bar()` → 세로 막대 그래프  
✔ `color` 리스트로 막대 색상 설정  

---

### 히스토그램(Histogram)
```python
data = np.random.randn(1000)  # 정규분포 데이터 생성

plt.hist(data, bins=30, color="purple", edgecolor="black")
plt.xlabel("값")
plt.ylabel("빈도수")
plt.title("Matplotlib 히스토그램")
plt.show()
```
✔ `bins=30` → 구간 개수 설정  
✔ `edgecolor="black"` → 막대 테두리 색상 지정  

---

### 산점도(Scatter Plot)
```python
x = np.random.rand(50)
y = np.random.rand(50)

plt.scatter(x, y, color="red", alpha=0.7)
plt.xlabel("X값")
plt.ylabel("Y값")
plt.title("Matplotlib 산점도")
plt.show()
```
✔ `alpha=0.7` → 투명도 설정(0: 완전 투명, 1: 불투명)  

---

## 4. Seaborn을 활용한 고급 그래프

Seaborn은 Matplotlib을 기반으로 더욱 **미려한 스타일과 통계 분석 기능**을 제공하는 라이브러리입니다.

### Seaborn 스타일 설정
```python
sns.set_theme(style="darkgrid")  # 스타일 적용
```
✔ `darkgrid`, `whitegrid`, `dark`, `white`, `ticks` 등 다양한 스타일 제공  

---

### Seaborn 막대 그래프 (`barplot`)
```python
tips = sns.load_dataset("tips")  # 예제 데이터셋 불러오기
sns.barplot(x="day", y="total_bill", data=tips, palette="pastel")
plt.title("Seaborn 막대 그래프")
plt.show()
```
✔ `palette="pastel"` → 부드러운 색상 적용  
✔ `data=tips` → Pandas DataFrame 사용 가능  

---

### Seaborn 히스토그램 (`histplot`)
```python
sns.histplot(tips["total_bill"], bins=30, kde=True, color="blue")
plt.title("Seaborn 히스토그램")
plt.show()
```
✔ `kde=True` → 커널 밀도 함수(KDE) 추가  

---

### Seaborn 박스 플롯(Box Plot)
```python
sns.boxplot(x="day", y="total_bill", data=tips, palette="coolwarm")
plt.title("Seaborn 박스 플롯")
plt.show()
```
✔ 박스 플롯은 **데이터 분포와 이상치를 시각화**할 때 유용  

---

### Seaborn 산점도 + 회귀선 (`regplot`)
```python
sns.regplot(x="total_bill", y="tip", data=tips)
plt.title("Seaborn 회귀선 그래프")
plt.show()
```
✔ `regplot()` → **산점도 + 회귀선 표시**  

---

## 5. Pandas 데이터 시각화 (`plot` 사용)

Pandas의 DataFrame에서도 직접 시각화가 가능합니다.

### Pandas 선 그래프
```python
df = pd.DataFrame({
    "월": ["1월", "2월", "3월", "4월"],
    "매출": [100, 150, 130, 170]
})

df.plot(x="월", y="매출", kind="line", marker="o", title="Pandas 선 그래프")
plt.show()
```
✔ `kind="line"` → 선 그래프  
✔ `marker="o"` → 원형 마커 추가  

---

## 6. 그래프 스타일 & 꾸미기

### 그래프 크기 설정 (`figsize`)
```python
plt.figure(figsize=(10, 5))  # 가로 10, 세로 5 크기 설정
```

### 눈금(font 크기 조정)
```python
plt.xticks(fontsize=12)
plt.yticks(fontsize=12)
```

### 여러 개의 그래프 그리기 (`subplot`)
```python
fig, axes = plt.subplots(1, 2, figsize=(12, 5))

axes[0].bar(["A", "B", "C"], [5, 7, 3], color="blue")
axes[0].set_title("첫 번째 그래프")

axes[1].plot([1, 2, 3], [10, 20, 30], marker="o", color="red")
axes[1].set_title("두 번째 그래프")

plt.show()
```
✔ `plt.subplots(행, 열, figsize=(너비, 높이))` → 여러 개의 그래프 표시 가능  

---

## 🎯 정리

✔ **Matplotlib** → 기본적인 시각화 지원 (`plot`, `bar`, `hist`, `scatter`)  
✔ **Seaborn** → 통계 기반 고급 시각화 (`barplot`, `boxplot`, `regplot`)  
✔ **Pandas** → DataFrame에서 바로 그래프 생성 (`df.plot()`)  
✔ **그래프 꾸미기** → `figsize`, `grid`, `legend`, `title` 등 활용  
✔ **여러 개의 그래프** → `subplot()` 사용  
