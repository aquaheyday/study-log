# 🗃️ JavaScript - IndexedDB

**IndexedDB**는 브라우저 내에 **비동기적으로 구조화된 데이터를 저장**할 수 있는 **로컬 데이터베이스 API**입니다.  
파일, 객체, 배열 등 **복잡한 데이터 구조**도 저장 가능하며, **오프라인 웹 앱**이나 대용량 클라이언트 데이터 처리에 매우 유용합니다.

---

## 1️⃣ 특징

- **비동기 기반** API (콜백 또는 Promise)
- **객체 기반 저장** (`key-value store`)
- **인덱스 지원** → 빠른 검색 가능
- 수십~수백 MB까지 저장 가능
- 브라우저 탭마다 독립적이며, 같은 origin 내에서 공유됨

---

## 2️⃣ 기본 구조

- **Database**: 하나의 DB (ex: "MyAppDB")
- **Object Store**: 테이블처럼 데이터를 저장하는 단위
- **Transaction**: 데이터 작업 단위 (읽기, 쓰기 등)
- **Index**: 검색을 빠르게 하기 위한 보조 구조

---

## 3️⃣ 기본 사용 예시

```js
const request = indexedDB.open("MyAppDB", 1);

// DB 버전 변경 시 or 처음 생성될 때 실행됨
request.onupgradeneeded = function (event) {
  const db = event.target.result;
  const store = db.createObjectStore("users", { keyPath: "id" }); // id를 기본 키로
  store.createIndex("name", "name", { unique: false });
};

// 연결 성공
request.onsuccess = function (event) {
  const db = event.target.result;

  // 트랜잭션 생성 (읽기/쓰기)
  const tx = db.transaction("users", "readwrite");
  const store = tx.objectStore("users");

  // 데이터 추가
  store.add({ id: 1, name: "Alice", age: 25 });

  tx.oncomplete = () => {
    console.log("저장 완료");
  };
};
```

---

## 4️⃣ 데이터 조회

```js
const tx = db.transaction("users", "readonly");
const store = tx.objectStore("users");

const getRequest = store.get(1);

getRequest.onsuccess = function () {
  console.log(getRequest.result); // { id: 1, name: "Alice", age: 25 }
};
```

---

## 5️⃣ 전체 데이터 순회 (커서)

```js
const tx = db.transaction("users", "readonly");
const store = tx.objectStore("users");

const cursorRequest = store.openCursor();

cursorRequest.onsuccess = function (event) {
  const cursor = event.target.result;
  if (cursor) {
    console.log(cursor.key, cursor.value);
    cursor.continue();
  } else {
    console.log("끝!");
  }
};
```

---

## 6️⃣ 주요 메서드 요약

| 메서드            | 설명                         |
|-------------------|------------------------------|
| `open(name, version)` | DB 열기 or 생성               |
| `createObjectStore()` | 객체 저장소 생성              |
| `add()`, `put()`       | 데이터 추가 / 갱신             |
| `get()`, `getAll()`    | 데이터 조회                    |
| `delete()`            | 항목 삭제                     |
| `clear()`             | 저장소 전체 삭제              |
| `openCursor()`        | 데이터 순회용 커서             |
| `createIndex()`       | 인덱스 생성 (검색 최적화용)    |

---

## ✅ IndexedDB vs 다른 저장소

| 항목              | `localStorage` / `sessionStorage` | `Cookies`           | `IndexedDB`                     |
|-------------------|-----------------------------------|----------------------|----------------------------------|
| 용량              | 약 5~10MB                         | 약 4KB               | 수십~수백 MB                    |
| 구조              | 문자열만 저장                    | 문자열               | 객체, 구조화된 데이터           |
| 비동기 여부       | ❌ 동기                           | ❌ 동기              | ✅ 비동기                       |
| 서버 전송         | ❌ 없음                           | ✅ 자동 전송         | ❌ 없음                         |
| 사용 용도         | 간단한 설정 저장                 | 인증, 세션 유지       | 대용량 오프라인 데이터 저장     |

---

## 🎯 정리

✔ IndexedDB는 **브라우저 내 로컬 데이터베이스**  
✔ 객체 단위로 저장되며, **인덱스/검색/커서 기능까지 제공**  
✔ 비동기 방식이므로 콜백 or `on*` 이벤트 기반으로 처리  
✔ 대용량 데이터를 클라이언트에 저장할 때 매우 유용  
✔ 로컬 앱, 오프라인 저장소, 캐싱 등에 적합

---

📌 IndexedDB는 API 구조가 복잡할 수 있으니,  
보다 쉽게 사용하려면 `idb` 라이브러리(Promise 기반 래퍼)를 활용하는 것도 추천
