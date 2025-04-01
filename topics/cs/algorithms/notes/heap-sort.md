# 🧠 Algorithms - 힙 정렬 (Heap Sort)

**힙 정렬(Heap Sort)** 은 **힙(Heap)** 자료구조를 활용하여  
최댓값(또는 최솟값)을 반복적으로 꺼내어 정렬하는 **비교 기반 정렬 알고리즘**입니다.

> 힙 정렬은 **최대 힙(또는 최소 힙)** 을 구성한 후, 루트 노드(최댓값/최솟값)를 꺼내며 배열을 정렬합니다.

---

## 1️⃣ 핵심 아이디어

1. **배열을 힙 구조(Max Heap or Min Heap)** 로 변환
2. 루트(최댓값)를 꺼내서 배열 끝으로 이동
3. 힙을 다시 재정렬(Heapify)하고 반복
4. 모든 요소를 꺼낼 때까지 반복 → 정렬 완료

> 💡 Max Heap → 오름차순 정렬 / Min Heap → 내림차순 정렬

---

## 2️⃣ 힙(Heap) 구조란?

- **완전 이진 트리(Complete Binary Tree)** 기반
- 부모 노드 ≥ 자식 노드 ⇒ **Max Heap**
- 부모 노드 ≤ 자식 노드 ⇒ **Min Heap**

배열로 표현 시:
- 왼쪽 자식: `2i + 1`, 오른쪽 자식: `2i + 2`

---

## 3️⃣ JavaScript 구현 예시 (Max Heap 기준)

```js
function heapSort(arr) {
  const n = arr.length;

  // 1. Max Heap 구성
  for (let i = Math.floor(n / 2) - 1; i >= 0; i--) {
    heapify(arr, n, i);
  }

  // 2. 정렬 과정
  for (let i = n - 1; i > 0; i--) {
    [arr[0], arr[i]] = [arr[i], arr[0]]; // 최대값을 맨 뒤로
    heapify(arr, i, 0); // 남은 부분 힙 정렬
  }

  return arr;
}

function heapify(arr, heapSize, rootIdx) {
  let largest = rootIdx;
  const left = 2 * rootIdx + 1;
  const right = 2 * rootIdx + 2;

  if (left < heapSize && arr[left] > arr[largest]) largest = left;
  if (right < heapSize && arr[right] > arr[largest]) largest = right;

  if (largest !== rootIdx) {
    [arr[rootIdx], arr[largest]] = [arr[largest], arr[rootIdx]];
    heapify(arr, heapSize, largest);
  }
}
```

---

## 4️⃣ 시간/공간 복잡도

| 구분            | 복잡도     |
|------------------|------------|
| 최선 / 평균 / 최악 | O(n log n) |
| 공간 복잡도       | O(1)       |
| 정렬 안정성       | ❌ 불안정 정렬 (Stable X) |

> ✅ 시간 복잡도가 **언제나 일정하게 O(n log n)**  
> ❌ 같은 값의 상대 순서가 보장되지 않음 (불안정)

---

## 5️⃣ 장점 vs 단점

| 장점                          | 단점                         |
|-------------------------------|------------------------------|
| 항상 O(n log n)의 성능 유지     | 구현이 비교적 복잡함          |
| 추가 메모리 필요 없음 (in-place) | 정렬 안정성이 없음            |
| 삽입/삭제가 빠른 힙 구조 기반   | 실제 속도는 퀵 정렬보다 느림    |

---

## 🎯 정리

✔ 힙 정렬은 **힙 구조(Max Heap or Min Heap)** 를 이용해 정렬하는 알고리즘  
✔ **최댓값을 꺼내며 뒤에서부터 정렬** → 오름차순  
✔ 시간 복잡도는 항상 O(n log n)  
✔ 메모리 효율이 뛰어난 **in-place 정렬**  
✔ 다만 **정렬 안정성이 없다**는 단점이 있음  
✔ 실무보다는 알고리즘 구현/학습용으로 자주 사용됨
