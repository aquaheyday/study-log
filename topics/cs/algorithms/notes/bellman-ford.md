# 🔁 벨만 포드 알고리즘 (Bellman-Ford)

**벨만-포드 알고리즘**은 **단일 시작점에서 모든 노드까지의 최단 경로**를 구하는 알고리즘입니다.  
특히 **음수 간선이 있는 그래프에서도 사용할 수 있으며**, **음수 사이클 존재 여부까지 판별**할 수 있는 것이 핵심 특징입니다.

---

## 1️⃣ 문제 정의

### 목표:
- 하나의 시작 노드로부터 모든 노드까지의 최단 거리 계산

### 조건:
- **방향 그래프**  
- **음수 간선 허용**  
- **음수 사이클이 있으면 → 경로 없음으로 판단**

---

## 2️⃣ 예시 그래프

### 간선 정보:
- (1 → 2, 비용 4)  
- (1 → 3, 비용 5)  
- (2 → 3, 비용 -2)  
- (3 → 4, 비용 3)

### 시작 노드: 1

---

## 3️⃣ 핵심 아이디어

- 모든 간선을 **N-1번 반복(Relatex)** 하여 거리 갱신  
- `dist[u] + cost < dist[v]` → `dist[v] = dist[u] + cost`  
- **N번째 루프에서 값이 갱신되면 → 음수 사이클 존재**

---

## 4️⃣ 파이썬 코드 예제

```python
def bellman_ford(n, edges, start):
    INF = float('inf')  # 무한대를 표현
    dist = [INF] * (n + 1)  # 시작 노드 기준 최단 거리 배열 초기화 (1번 인덱스부터 사용)
    dist[start] = 0  # 시작 노드까지의 거리는 0으로 설정

    # 1. 간선을 N-1번 반복하여 Relax(거리 갱신) 수행
    for i in range(n - 1):  # 모든 정점에 대해 최대 N-1번 반복
        for u, v, cost in edges:  # 모든 간선을 하나씩 확인
            # u에서 v로 가는 거리가 현재보다 짧으면 갱신
            if dist[u] != INF and dist[u] + cost < dist[v]:
                dist[v] = dist[u] + cost

    # 2. N번째 반복 시에도 값이 갱신된다면 → 음수 사이클 존재
    for u, v, cost in edges:
        if dist[u] != INF and dist[u] + cost < dist[v]:
            return None  # 음수 사이클 감지 시 None 리턴

    # 3. 최종적으로 계산된 최단 거리 배열 반환
    return dist

```

---

## 5️⃣ 예제 입력

```python
n = 4
edges = [
    (1, 2, 4),
    (1, 3, 5),
    (2, 3, -2),
    (3, 4, 3)
]
start = 1

result = bellman_ford(n, edges, start)

if result is None:
    print("음수 사이클 존재")
else:
    for i in range(1, n + 1):
        print(f"{i}번 노드까지 거리: {result[i]}")
```

---

## 6️⃣ 시간 복잡도

- **O(N × E)**  
  - N: 노드 개수  
  - E: 간선 개수  
- 플로이드-워셜보다 효율적 (`O(N³)` → `O(N×E)`)

---

## 7️⃣ 주요 포인트 요약

| 항목               | 설명                                  |
|--------------------|---------------------------------------|
| 출발점 기준 최단거리 | 단일 시작점                          |
| 음수 간선           | 허용됨                            |
| 음수 사이클 탐지     | 가능 (`N번째 루프`로 확인)       |
| 그래프 형태         | 방향 그래프                          |
| 알고리즘 구조       | 반복 완화(Relax) 기반                |

---

## 🎯 정리 요약

✔ 음수 간선 O  
✔ 음수 사이클 검출 O  
✔ 다익스트라 대체 가능  
✔ O(N × E)의 효율적 알고리즘  
✔ 플로이드 워셜보다 빠르고, DP 기반은 아님
