# 📬 Data Structures - 큐 (Queue)

**큐(Queue)** 는 **먼저 들어온 데이터가 먼저 나오는** **선입선출(First-In, First-Out, FIFO)** 방식의 선형 자료구조입니다.

> 줄 서기, 은행 대기열처럼 먼저 온 순서대로 처리하는 구조

---

## 1️⃣ 큐의 주요 특징

- **FIFO 구조**: 먼저 들어온 데이터가 먼저 나감  
- 삽입은 뒤(rear), 삭제는 앞(front)에서 발생  
- 배열 또는 연결 리스트로 구현 가능

---

## 2️⃣ 주요 연산

| 연산          | 설명                               |
|---------------|------------------------------------|
| `enqueue(x)`  | 요소 x를 큐의 뒤쪽에 삽입            |
| `dequeue()`   | 큐의 앞쪽 요소 제거 및 반환          |
| `peek()`      | 가장 앞 요소를 제거하지 않고 확인       |
| `isEmpty()`   | 큐가 비어있는지 확인                 |
| `size()`      | 큐에 들어 있는 요소 개수 반환           |

---

## 3️⃣ 큐의 동작 예시

```
enqueue(10) enqueue(20) enqueue(30)

Queue 상태: front → [10, 20, 30] ← rear

dequeue() → 10 제거

Queue 상태: front → [20, 30] ← rear
```

---

## 4️⃣ 큐의 시간 복잡도

| 연산         | 시간 복잡도 |
|--------------|--------------|
| `enqueue`    | O(1)         |
| `dequeue`    | O(1)         |
| `peek`       | O(1)         |

> 배열 기반 큐에서는 **앞쪽 요소 제거 시 요소 이동이 필요할 수도 있음**  
> 이를 해결하기 위해 **원형 큐(Circular Queue)** 를 사용하기도 함

---

## 5️⃣ 큐의 구현 방식

- **배열 기반 큐**
  - 고정 크기 제한, 재할당 필요
  - front 이동 시 요소를 밀어야 하는 단점 있음
- **연결 리스트 기반 큐**
  - front와 rear 포인터로 삽입/삭제
  - 유연한 크기 조절 가능

---

## 6️⃣ 큐의 활용 사례

- **프린터 대기열**  
- **프로세스/스레드 스케줄링**  
- **네트워크 패킷 처리**  
- **BFS(너비 우선 탐색)**  
- **버퍼(Keyboard, Stream 등)**

---

## 7️⃣ 큐 vs 스택

| 항목      | 큐 (Queue)            | 스택 (Stack)            |
|-----------|------------------------|--------------------------|
| 구조      | 선입선출 (FIFO)          | 후입선출 (LIFO)            |
| 삽입 위치 | rear (뒤)               | top (위)                  |
| 삭제 위치 | front (앞)              | top (위)                  |
| 활용 예   | 작업 대기열, BFS 등      | 함수 호출, 되돌리기 등      |

---

## 🎯 정리 요약

✔ **큐(Queue)** 는 **선입선출(FIFO)** 구조의 선형 자료구조  
✔ `enqueue`, `dequeue` 모두 **O(1)** 로 빠름  
✔ 배열 또는 연결 리스트로 구현 가능  
✔ **작업 대기열, BFS, OS 스케줄링** 등 다양한 곳에 사용됨

