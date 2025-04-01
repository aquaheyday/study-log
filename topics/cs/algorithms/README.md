# 🧠 Algorithms - 개념 정리 목록

자료구조 기반의 주요 알고리즘을 유형별로 정리한 목록입니다.  
각 주제는 개념 요약 + 시간복잡도 + 예제 + 구현 코드 등을 포함합니다.

---

### 🔢 정렬 (Sorting)
| 주제 | 파일명 | 설명 |
|------|--------|------|
| 버블 정렬 (Bubble Sort) | [bubble-sort.md](./notes/bubble-sort.md) | 인접한 값을 비교하며 정렬하는 기본적인 정렬 |
| 선택 정렬 (Selection Sort) | [selection-sort.md](./notes/selection-sort.md) | 가장 작은 값을 선택해 앞에서부터 채워가는 방식 |
| 삽입 정렬 (Insertion Sort) | [insertion-sort.md](./notes/insertion-sort.md) | 앞에서부터 정렬된 부분에 값을 끼워 넣는 방식 |
| 병합 정렬 (Merge Sort) | [merge-sort.md](./notes/merge-sort.md) | 분할정복을 이용한 안정 정렬 알고리즘 |
| 퀵 정렬 (Quick Sort) | [quick-sort.md](./notes/quick-sort.md) | 피벗을 기준으로 분할하는 비안정 정렬 |
| 힙 정렬 (Heap Sort) | [heap-sort.md](./notes/heap-sort.md) | 힙 자료구조를 이용한 정렬 알고리즘 |

---

### 🔍 탐색 (Searching)
| 주제 | 파일명 | 설명 |
|------|--------|------|
| 선형 탐색 (Linear Search) | [linear-search.md](./notes/linear-search.md) | 순차적으로 비교하며 값을 찾는 방식 |
| 이진 탐색 (Binary Search) | [binary-search.md](./notes/binary-search.md) | 정렬된 배열에서 중간값을 기준으로 탐색 |
| BFS (너비 우선 탐색) | [bfs.md](./notes/bfs.md) | 큐를 사용한 그래프 탐색 |
| DFS (깊이 우선 탐색) | [dfs.md](./notes/dfs.md) | 스택/재귀를 이용한 깊이 우선 그래프 탐색 |

---

### 📐 분할 정복 (Divide & Conquer)
| 주제 | 파일명 | 설명 |
|------|--------|------|
| 병합 정렬 | [merge-sort.md](./notes/merge-sort.md) | 분할 후 정복 단계에서 병합 |
| 퀵 정렬 | [quick-sort.md](./notes/quick-sort.md) | 분할 후 각 파티션을 정렬 |
| 거듭제곱 계산 | [power.md](./notes-conquer/power.md) | O(log n)으로 xⁿ 계산 |

---

### 📊 그리디 알고리즘 (Greedy)
| 주제 | 파일명 | 설명 |
|------|--------|------|
| 거스름돈 문제 | [change.md](./notes/change.md) | 가장 큰 동전부터 선택해 문제 해결 |
| 회의실 배정 문제 | [meeting-room.md](./notes/meeting-room.md) | 가장 먼저 끝나는 회의 우선 선택 |
| 최소 신장 트리 (Kruskal) | [kruskal.md](./notes/kruskal.md) | 그리디 기반의 MST 알고리즘 |

---

### 📈 동적 계획법 (DP)
| 주제 | 파일명 | 설명 |
|------|--------|------|
| 피보나치 수열 | [fibonacci.md](./notes/fibonacci.md) | DP vs 재귀 성능 비교 |
| 0/1 배낭 문제 | [knapsack.md](./notes/knapsack.md) | DP로 최적 부분합 구하기 |
| LIS (최장 증가 부분 수열) | [lis.md](./notes/lis.md) | 증가하는 수열 중 가장 긴 것 찾기 |

---

### 🔁 백트래킹 (Backtracking)
| 주제 | 파일명 | 설명 |
|------|--------|------|
| N-Queen 문제 | [nqueen.md](./notes/nqueen.md) | 체스판에 퀸 배치하며 충돌 검사 |
| 순열/조합 생성 | [permutation-combination.md](./notes/permutation-combination.md) | 재귀로 가능한 모든 조합 생성 |

---

### 📉 최단 경로 (Shortest Path)
| 주제 | 파일명 | 설명 |
|------|--------|------|
| 다익스트라 알고리즘 | [dijkstra.md](./notes/dijkstra.md) | 우선순위 큐를 활용한 최단 경로 탐색 |
| 플로이드 워셜 알고리즘 | [floyd-warshall.md](./notes/floyd-warshall.md) | 모든 노드 간 최단 거리 계산 |
| 벨만 포드 알고리즘 | [bellman-ford.md](./notes/bellman-ford.md) | 음수 가중치 허용, 음의 사이클 탐지 |

---

### 🌉 그래프 알고리즘
| 주제 | 파일명 | 설명 |
|------|--------|------|
| 유니온 파인드 | [union-find.md](./notes/union-find.md) | 서로소 집합 자료구조 |
| 위상 정렬 | [topological-sort.md](./notes/topological-sort.md) | 방향 그래프의 정렬 순서 구하기 |
| 최소 신장 트리 (Prim) | [prim.md](./notes/prim.md) | MST 알고리즘 (그리디 기반) |

---

### 🎲 기타
| 주제 | 파일명 | 설명 |
|------|--------|------|
| 비트 마스크 | [bitmask.md](./notes/bitmask.md) | 비트를 이용한 상태 저장 |
| 누적합 (Prefix Sum) | [prefix-sum.md](./notes/prefix-sum.md) | 배열의 부분합을 빠르게 계산 |
| 슬라이딩 윈도우 | [sliding-window.md](./notes/sliding-window.md) | 고정 크기 또는 가변 크기의 윈도우 처리 |
