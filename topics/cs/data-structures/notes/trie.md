# 🔡 Data Structures - 트라이 (Trie, Prefix Tree)

**트라이(Trie)** 는 문자열을 **공통 접두사(prefix)** 를 기준으로 분기하여 저장하는 **트리 기반 자료구조**입니다.  
주로 **문자열 검색**, **자동완성**, **사전(dictionary)** 등에 사용됩니다.

> 트리(Tree) 구조에서 파생된 구조로, 문자 단위로 노드를 분기하여 문자열 집합을 표현합니다.

---

## 1️⃣ 트라이의 주요 특징

- 각 노드는 **한 글자(문자)** 를 나타냄
- 루트 노드는 **비어 있는 상태**
- 문자열의 **접두사(prefix)를 공유**함으로써 **공간 절약**
- 문자열의 끝을 나타내는 **종료 표시(End of Word)** 가 존재

---

## 2️⃣ 동작 예시

문자열 집합: `"cat"`, `"can"`, `"car"`

```
        (root)
        /  |  \
      c    ...  
      |
      a
    / | \
   n  r  t
```

→ `"ca"` 접두사 공유, 이후 문자에서 분기됨  
→ 효율적인 **검색 / 저장** 가능

---

## 3️⃣ 주요 연산

| 연산           | 설명                                      | 시간 복잡도 |
|----------------|-------------------------------------------|--------------|
| 삽입 (Insert)  | 문자 하나씩 따라가며 노드 삽입              | O(L)         |
| 탐색 (Search)  | 문자를 따라가며 존재 여부 확인              | O(L)         |
| 삭제 (Delete)  | 존재 여부 확인 후 노드 삭제 및 정리         | O(L)         |
| 자동완성 (Prefix) | 주어진 접두사로 시작하는 단어 나열          | O(P + 결과 수) |

> L: 문자열 길이  
> P: 접두사 길이

---

## 4️⃣ 트라이 구현 예시 (Python)

```
class TrieNode:
    def __init__(self):
        self.children = dict()
        self.is_end = False

class Trie:
    def __init__(self):
        self.root = TrieNode()

    def insert(self, word):
        node = self.root
        for ch in word:
            if ch not in node.children:
                node.children[ch] = TrieNode()
            node = node.children[ch]
        node.is_end = True

    def search(self, word):
        node = self.root
        for ch in word:
            if ch not in node.children:
                return False
            node = node.children[ch]
        return node.is_end

    def starts_with(self, prefix):
        node = self.root
        for ch in prefix:
            if ch not in node.children:
                return False
            node = node.children[ch]
        return True
```

---

## 5️⃣ 트라이의 장점

✅ **빠른 문자열 검색**  
✅ **공통 접두사 활용 → 공간 절약**  
✅ **접두사 기반 기능에 특화 (자동완성 등)**  
✅ **정렬된 순서로 탐색 가능** (사전 순)

---

## 6️⃣ 트라이의 단점

⚠️ 문자가 다양할 경우 **자식 노드 수 증가**  
⚠️ 메모리 사용량이 높을 수 있음  
⚠️ 일반적인 해시 테이블보다 단순 키 탐색은 느릴 수도 있음

---

## 7️⃣ 트라이 활용 사례

- **자동완성 시스템 (Autocomplete)**
- **단어 사전(Dictionary)**
- **검색 추천 기능**
- **문자열 패턴 매칭 (예: 문자열 필터링)**
- **IP 라우팅 테이블 (이진 트라이 형태)**

---

## 8️⃣ 변형 구조

| 구조               | 설명                                      |
|--------------------|-------------------------------------------|
| 압축 트라이 (Radix Trie) | 노드 간 중복 문자열을 압축하여 공간 절약       |
| 이진 트라이 (Binary Trie) | 비트 단위 분기로 구성 (네트워크/라우팅에 사용) |
| DAWG (Directed Acyclic Word Graph) | 중복 문자열 트리 최소화 그래프 구조           |

---

## 9️⃣ 정리 요약 🎯

✔ 트라이는 **문자열 검색과 접두사 처리를 위한 트리형 자료구조**  
✔ 삽입/탐색 O(L), 접두사 자동완성 등 특화  
✔ **문자 단위 저장 + 종료 표시** 구조  
✔ **자동완성, 검색엔진, 필터링, 사전 구현** 등에 자주 활용
