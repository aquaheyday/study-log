# 📦 Data Structures - 배열 (Array)

**배열(Array)** 은 **동일한 자료형의 데이터를 연속된 메모리 공간에 저장하는 자료구조**입니다.  
인덱스를 이용해 빠르게 접근 가능하며, 가장 기본이 되는 선형 구조입니다.

> 데이터를 "순서대로 줄 세워" 저장하는 구조  
> 메모리상 **연속적 공간**에 저장되어 인덱스로 즉시 접근 가능

---

## 1️⃣ 배열의 특징

- **같은 자료형**의 데이터만 저장 가능  
- **인덱스(index)** 를 이용해 요소에 접근 (0번부터 시작)  
- **연속된 메모리 공간** 사용 → 빠른 접근 속도 (O(1))  
- **고정 크기**: 생성 시 크기 지정 → 크기 변경 불가 (일반 배열 기준)

---

## 2️⃣ 배열의 장점과 단점

| 항목   | 장점                                    | 단점                                      |
|--------|-----------------------------------------|-------------------------------------------|
| 접근   | 인덱스로 **빠르게 접근 가능 (O(1))**      | 없음                                      |
| 삽입/삭제 | 중간 삽입/삭제 시 **모든 요소 이동 필요** | 비효율적 (O(n))                           |
| 메모리 | **연속된 공간 사용 → 캐시 적중률 높음**    | 크기 변경 불가, 메모리 낭비 가능           |

---

## 3️⃣ 배열의 시간 복잡도

| 연산           | 시간 복잡도 |
|----------------|--------------|
| **접근 (Access)**   | O(1)         |
| **탐색 (Search)**   | O(n)         |
| **삽입 (Insert)**   | O(n)         |
| **삭제 (Delete)**   | O(n)         |

> 중간에 값을 추가하거나 제거하면 전체 데이터를 이동해야 하므로 느림

---

## 4️⃣ 배열과 리스트 비교

| 항목       | 배열 (Array)           | 연결 리스트 (Linked List)     |
|------------|-------------------------|--------------------------------|
| 메모리 구조 | 연속된 공간               | 비연속 (노드 + 포인터)         |
| 접근 속도   | 빠름 (O(1))              | 느림 (O(n))                    |
| 삽입/삭제   | 느림 (O(n))              | 빠름 (O(1) ~ O(n))             |
| 메모리 낭비 | 가능                     | 포인터 공간 추가로 메모리 더 필요 |

---

## 5️⃣ 배열의 종류 (언어에 따라 차이 있음)

- **1차원 배열**: 가장 기본적인 형태 (ex. `int arr[5]`)  
- **다차원 배열**: 2차원 이상의 배열 (ex. 행렬, `arr[3][3]`)  
- **동적 배열 (Dynamic Array)**: 크기를 동적으로 늘릴 수 있는 배열  
  - 예: C++의 `vector`, Python의 `list`, Java의 `ArrayList`

---

## 6️⃣ 배열의 활용 예

- 테이블, 행렬, 이미지 데이터 저장  
- 해시테이블 구현의 내부 버킷 구조  
- 스택, 큐 등의 기반 구조  
- 정렬, 탐색 알고리즘의 기본 입력 구조

---

## 🎯 정리

✔ **배열**은 같은 타입 데이터를 **연속된 메모리 공간**에 저장하는 자료구조  
✔ 인덱스를 통해 **빠른 접근 가능 (O(1))**  
✔ **삽입/삭제는 느리며**, 고정 크기라는 제약이 있음  
✔ 언어마다 동적 배열 등 다양한 변형 제공  
✔ 많은 기본 알고리즘과 자료구조의 **기초가 되는 구조**

