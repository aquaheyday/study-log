# 📊 Python Pandas 기초부터 활용까지

Pandas는 **데이터 분석과 처리**를 위한 Python 라이브러리로,  
**데이터프레임(DataFrame)** 을 이용하여 **엑셀과 유사한 방식으로 데이터를 다룰 수 있습니다.**

---

## 1. Pandas란?

- **구조화된 데이터(테이블 형태) 처리에 최적화된 라이브러리**
- **NumPy 기반으로 구축되어 고속 데이터 연산 지원**
- **CSV, Excel, SQL 등 다양한 데이터 포맷 지원**
- **정형 데이터 분석과 머신러닝 전처리에 필수적**

---

## 2. Pandas 설치 및 불러오기

### Pandas 설치
```sh
pip install pandas
```

### Pandas 불러오기
```python
import pandas as pd
```

✔ `pd`는 관례적으로 사용되는 Pandas의 별칭  

---

## 3. Pandas 데이터 구조

### Series (1차원 데이터)
```python
import pandas as pd

s = pd.Series([10, 20, 30, 40])
print(s)
```
#### 출력 결과
```
0    10
1    20
2    30
3    40
dtype: int64
```
✔ **인덱스(index)와 값(value)으로 구성된 1차원 데이터 구조**  

---

### DataFrame (2차원 데이터)
```python
data = {"Name": ["Alice", "Bob", "Charlie"],
        "Age": [25, 30, 35],
        "City": ["Seoul", "Busan", "Incheon"]}

df = pd.DataFrame(data)
print(df)
```
#### 출력 결과
```
     Name  Age     City
0   Alice   25   Seoul
1     Bob   30   Busan
2  Charlie   35  Incheon
```
✔ **엑셀과 유사한 형태의 2차원 데이터 구조**  

---

## 4. 데이터 불러오기 & 저장

### CSV 파일 불러오기
```python
df = pd.read_csv("data.csv")
```

### CSV 파일 저장
```python
df.to_csv("output.csv", index=False)
```

✔ `index=False` → **인덱스 없이 저장**  

---

## 5. 데이터프레임 기본 정보 확인

### 데이터프레임 개요 (`info()`, `describe()`)
```python
print(df.info())  # 데이터 타입, 결측값 확인
print(df.describe())  # 수치형 데이터 통계 정보
```

### 데이터 확인 (`head()`, `tail()`)
```python
print(df.head())   # 상위 5개 데이터
print(df.tail(3))  # 하위 3개 데이터
```

### 데이터 크기 (`shape`)
```python
print(df.shape)  # (행 개수, 열 개수)
```

---

## 6. 데이터 인덱싱 & 슬라이싱

### 특정 열 선택
```python
print(df["Name"])  # Series 반환
print(df[["Name", "Age"]])  # DataFrame 반환
```

---

### 특정 행 선택 (`loc`, `iloc`)
```python
print(df.loc[0])  # 첫 번째 행 (라벨 기반)
print(df.iloc[1])  # 두 번째 행 (숫자 인덱스 기반)
```

---

### 행과 열 동시 선택
```python
print(df.loc[0, "Name"])  # 특정 값 (0행, Name 열)
print(df.loc[:, ["Name", "City"]])  # 모든 행에서 Name, City 열 선택
print(df.iloc[1:3, 0:2])  # 1~2행, 0~1열 선택
```

---

## 7. 데이터 필터링

### 조건 필터링
```python
print(df[df["Age"] > 25])  # Age가 25 이상인 행 선택
```

### 여러 조건 필터링 (`&`, `|` 사용)
```python
print(df[(df["Age"] > 25) & (df["City"] == "Busan")])  # Age > 25 이면서 City가 Busan
```

---

## 8. 데이터 수정 및 추가

### 새로운 열 추가
```python
df["Salary"] = [50000, 60000, 70000]  # 새 열 추가
```

### 값 변경
```python
df.loc[1, "Age"] = 32  # 1번 행의 Age 값 변경
```

### 행 추가 (`append`)
```python
new_data = {"Name": "David", "Age": 28, "City": "Daegu"}
df = df.append(new_data, ignore_index=True)
```

✔ `ignore_index=True` → **새로운 인덱스 자동 부여**  

---

## 9. 결측값 처리 (Missing Data)

### 결측값 확인 (`isnull().sum()`)
```python
print(df.isnull().sum())  # 각 열의 결측값 개수 확인
```

### 결측값 제거 (`dropna`)
```python
df.dropna(inplace=True)  # 결측값이 포함된 행 제거
```

### 결측값 채우기 (`fillna`)
```python
df.fillna(value={"Age": df["Age"].mean()}, inplace=True)  # Age 결측값을 평균으로 대체
```

---

## 10. 데이터 정렬

### 특정 열 기준 정렬 (`sort_values`)
```python
df.sort_values(by="Age", ascending=False, inplace=True)  # Age 기준 내림차순 정렬
```

---

## 11. 데이터 그룹화 (`groupby`)

### 특정 열을 기준으로 그룹화
```python
grouped = df.groupby("City")["Age"].mean()
print(grouped)
```

✔ 각 도시별 `Age`의 평균 계산  

---

## 12. 데이터 병합

### `merge()` 를 사용한 데이터프레임 병합
```python
df1 = pd.DataFrame({"ID": [1, 2, 3], "Name": ["Alice", "Bob", "Charlie"]})
df2 = pd.DataFrame({"ID": [1, 2, 3], "Salary": [50000, 60000, 70000]})

merged_df = pd.merge(df1, df2, on="ID")  # ID를 기준으로 병합
print(merged_df)
```

---

## 13. 데이터 시각화 (`plot` 활용)

Pandas는 **Matplotlib** 와 함께 기본적인 그래프를 지원합니다.

### 라인 차트 (`plot()`)
```python
import matplotlib.pyplot as plt

df["Age"].plot(kind="line")
plt.show()
```

---

### 히스토그램 (`hist()`)
```python
df["Age"].hist()
plt.show()
```

---

## 14. Pandas vs Python 리스트 성능 비교

Pandas는 **벡터 연산을 지원하여 일반 리스트보다 훨씬 빠름**.

```python
import pandas as pd
import time

size = 1000000
data = list(range(size))
df = pd.Series(data)

# 리스트 연산
start = time.time()
sum_list = sum(data)
print("리스트 연산 시간:", time.time() - start)

# Pandas 연산
start = time.time()
sum_pandas = df.sum()
print("Pandas 연산 시간:", time.time() - start)
```

✔ **Pandas가 훨씬 빠름** (NumPy 기반 벡터 연산 최적화 덕분)  

---

## 🎯 정리

✔ **Pandas DataFrame** → 표 형식(엑셀과 유사) 데이터 처리  
✔ **데이터 불러오기 & 저장** → `read_csv()`, `to_csv()`  
✔ **데이터 탐색** → `info()`, `describe()`, `head()`  
✔ **데이터 필터링** → 조건식을 사용한 데이터 선택  
✔ **결측값 처리** → `dropna()`, `fillna()`  
✔ **데이터 병합** → `merge()`, `concat()`  
✔ **데이터 시각화** → `plot()`, `hist()` 지원  
