# 🌉 알고리즘 - 위상 정렬 (Topological Sort)

**위상 정렬**은 **방향성이 있는 그래프(DAG)** 에서 모든 노드의 **선행 관계를 지키면서 순서대로 나열**하는 알고리즘입니다.

---

## 1️⃣ 위상 정렬이란?

- **방향 그래프(DAG: Directed Acyclic Graph)** 에서 **“선행 순서가 있는 작업”들을 정렬**하는 방법
- 즉, 어떤 노드를 방문하기 전에 반드시 **먼저 처리해야 할 노드**가 있을 때

---

## 2️⃣ 예시 문제

### 예시: 선수 과목 조건

> 과목 1을 듣기 전에 과목 2를 들어야 하고,  
> 과목 3은 1과 2를 다 들어야 한다면?

```
선행 조건:
2 → 1  
1 → 3  
2 → 3
```

가능한 수강 순서 예시:

```
2 → 1 → 3
```

---

## 3️⃣ 해결 방법: 진입 차수(Indegree) + 큐(Queue)

1. 각 노드의 **진입 차수(선행 노드 수)** 를 계산
2. **진입 차수가 0인 노드**를 큐에 삽입
3. 큐에서 꺼내면서 **연결된 노드의 진입 차수 감소**
4. 진입 차수가 0이 되면 다시 큐에 삽입
5. 큐가 빌 때까지 반복 → 정렬 결과 완성

---

## 4️⃣ 파이썬 코드 예제

```python
from collections import deque

def topological_sort(n, edges):
    graph = [[] for _ in range(n + 1)]
    indegree = [0] * (n + 1)

    # 그래프와 진입 차수 초기화
    for a, b in edges:
        graph[a].append(b)
        indegree[b] += 1

    queue = deque()
    result = []

    # 처음에 진입 차수 0인 노드 삽입
    for i in range(1, n + 1):
        if indegree[i] == 0:
            queue.append(i)

    # 위상 정렬 수행
    while queue:
        now = queue.popleft()
        result.append(now)
        for neighbor in graph[now]:
            indegree[neighbor] -= 1
            if indegree[neighbor] == 0:
                queue.append(neighbor)

    return result
```

---

### 📦 예제 입력

```python
n = 3  # 노드 수
edges = [(2, 1), (1, 3), (2, 3)]  # 방향성 간선
print(topological_sort(n, edges))  # 출력: [2, 1, 3]
```

---

## 5️⃣ 시간 복잡도

- 노드: `V`, 간선: `E`
- `O(V + E)`

→ 노드당 1번, 간선당 1번씩만 처리됨 (BFS 기반)

---

## 6️⃣ 주요 포인트 요약

| 항목              | 설명 |
|-------------------|------|
| **DAG만 가능**     | 사이클이 있으면 위상 정렬 불가능 |
| **진입 차수**       | 선행 노드가 몇 개인지 저장 |
| **큐**             | 진입 차수 0인 노드부터 처리 |
| **결과**           | 모든 노드가 순서대로 정렬됨 |
| **여러 정렬 가능성** | 정답이 하나로 정해지지 않을 수도 있음 |

---

## 🎯 정리 요약

✔ **위상 정렬 = 선행 조건이 있는 작업의 순서 정하기**  
✔ **진입 차수 0부터 시작**해서 하나씩 정렬  
✔ **사이클이 없는 방향 그래프(DAG)** 만 가능  
✔ 시간복잡도는 `O(V + E)`, BFS 기반이라 효율적

