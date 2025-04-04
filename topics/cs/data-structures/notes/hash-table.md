# 🗃️ Data Structures - 해시 테이블 (Hash Table)

**해시 테이블**은 `Key-Value` 쌍 데이터를 저장하는 자료구조입니다.  
**해시 함수(Hash Function)** 를 통해 키를 배열의 인덱스로 매핑하여 빠르게 데이터를 **저장/탐색/삭제**할 수 있습니다.

> 해시 테이블은 **빠른 검색과 삽입**이 핵심인 구조이며, **딕셔너리(Dictionary)** 형태로 자주 사용됩니다.

---

## 1️⃣ 해시 테이블의 주요 특징

- 데이터를 **Key → Value** 형태로 저장  
- **해시 함수**를 사용해 키를 인덱스로 변환  
- **평균적으로 매우 빠른 접근 속도 (`O(1)`)**  
- **충돌(Collision)** 이 발생할 수 있으며, 이를 해결해야 함

---

## 2️⃣ 동작 구조 예시

1. 저장: `"apple" → 100` 저장  
2. 해시 함수로 `"apple"`을 인덱스 계산 → 예: `hash("apple") % N = 3`  
3. 배열의 인덱스 3에 (key, value) 저장 → `(apple, 100)`  
4. 검색 시도: `"apple"` → 해시 함수 → 3번 인덱스 확인

---

## 3️⃣ 주요 연산 및 시간 복잡도

| 연산       | 설명                           | 평균 시간 | 최악 시간 |
|------------|--------------------------------|-----------|-----------|
| 삽입       | Key → Value 저장               | O(1)      | O(N)      |
| 탐색       | Key 기반으로 Value 검색        | O(1)      | O(N)      |
| 삭제       | Key를 찾아 삭제                | O(1)      | O(N)      |

> 최악의 경우는 **해시 충돌**이 심하게 일어난 경우

---

## 4️⃣ 해시 충돌(Collision)과 해결법

해시 충돌: **서로 다른 키**가 **같은 인덱스**로 해싱되는 현상

#### 해결 방법

| 방법            | 설명                                      |
|-----------------|-------------------------------------------|
| 체이닝 (Chaining) | 동일 인덱스에 리스트로 여러 데이터 저장  |
| 개방 주소법 (Open Addressing) | 빈 슬롯을 찾아 다음 인덱스로 이동 저장 |

---

## 5️⃣ 해시 테이블 구현 예시 (Python)

### Python 기본 딕셔너리 사용

```python
# 기본 dict 사용
table = dict()
table["apple"] = 100
print(table["apple"])  # 출력: 100
```

### 해시 테이블 직접 구현 (체이닝 방식)

```python
class HashTable:
    def __init__(self, size=10):
        self.size = size
        self.table = [[] for _ in range(size)]

    def _hash(self, key):
        return hash(key) % self.size

    def insert(self, key, value):
        idx = self._hash(key)
        for i, (k, v) in enumerate(self.table[idx]):
            if k == key:
                self.table[idx][i] = (key, value)
                return
        self.table[idx].append((key, value))

    def get(self, key):
        idx = self._hash(key)
        for k, v in self.table[idx]:
            if k == key:
                return v
        return None
```

---

## 6️⃣ 해시 테이블의 장점

✅ **빠른 데이터 접근 (`O(1)` 평균)**  
✅ **Key 기반 정밀 검색**에 적합  
✅ **데이터 삽입/삭제** 효율적  
✅ 실용적이고 구현 간단 (Python dict, Java HashMap 등)

---

## 7️⃣ 해시 테이블의 단점

⚠️ **해시 충돌 처리** 필요  
⚠️ **해시 함수 품질**에 따라 성능 편차 발생  
⚠️ **정렬된 탐색 불가능** (순서 없음)  
⚠️ **메모리 사용량**이 많을 수 있음

---

## 8️⃣ 활용 사례

- **딕셔너리 자료형 (Python dict, Java HashMap 등)**  
- **중복 체크 (Seen Set)**  
- **캐싱 / 메모이제이션 (Memoization)**  
- **데이터베이스 인덱스 구조**  
- **해시 기반 집합(Set), 매핑(Map) 자료구조**

---

## 🎯 정리 요약 

✔ 해시 테이블은 **Key → Value 매핑 구조**  
✔ 빠른 탐색/삽입 가능 (`O(1)` 평균)  
✔ **충돌(Collision)** 발생 시 해결 기법 필요  
✔ 딕셔너리처럼 **검색 속도가 중요한 상황**에서 강력한 성능  
✔ 언어 내장 자료구조로 자주 활용됨 (`dict`, `set`, `map`)
