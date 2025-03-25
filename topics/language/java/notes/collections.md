# 📚 Java - List, Set, Map

Java의 컬렉션 프레임워크는 데이터를 **효율적으로 저장/관리**하기 위한 구조입니다.  
그 중 가장 많이 쓰이는 3종류는 `List`, `Set`, `Map`입니다.

---

## 1️⃣ List

- **순서 있음**, **중복 허용**
- 인덱스로 요소를 관리
- 대표 구현체: `ArrayList`, `LinkedList`

```java
import java.util.*;

List<String> list = new ArrayList<>();
list.add("apple");
list.add("banana");
list.add("apple");  // 중복 허용

System.out.println(list.get(0));       // apple
System.out.println(list.size());       // 3
System.out.println(list.contains("banana")); // true
```

---

## 2️⃣ Set

- **순서 없음**, **중복 허용 안 됨**
- 대표 구현체: `HashSet`, `LinkedHashSet`, `TreeSet`

```java
Set<String> set = new HashSet<>();
set.add("apple");
set.add("banana");
set.add("apple");  // 중복 무시됨

System.out.println(set.size());        // 2
System.out.println(set.contains("apple")); // true
```

> `TreeSet`은 정렬, `LinkedHashSet`은 입력 순서 유지

---

## 3️⃣ Map

- **key-value 쌍**으로 구성
- key는 **중복 불가**, value는 **중복 가능**
- 대표 구현체: `HashMap`, `LinkedHashMap`, `TreeMap`

```java
Map<String, Integer> map = new HashMap<>();
map.put("apple", 3);
map.put("banana", 5);
map.put("apple", 10);  // 기존 값 덮어씀

System.out.println(map.get("apple"));       // 10
System.out.println(map.containsKey("banana")); // true
```

---

## 4️⃣ 반복 (List/Set)

```java
for (String item : list) {
    System.out.println(item);
}

for (String item : set) {
    System.out.println(item);
}
```

---

## 5️⃣ 반복 (Map)

```java
for (Map.Entry<String, Integer> entry : map.entrySet()) {
    System.out.println(entry.getKey() + " = " + entry.getValue());
}
```

또는:

```java
map.forEach((k, v) -> System.out.println(k + ": " + v));
```

---

## 6️⃣ 주요 메서드 비교

| 기능 | List | Set | Map |
|------|------|-----|-----|
| 요소 추가 | `add()` | `add()` | `put(key, value)` |
| 요소 제거 | `remove(index)` | `remove(value)` | `remove(key)` |
| 포함 여부 | `contains(value)` | `contains(value)` | `containsKey(key)` |
| 길이 확인 | `size()` | `size()` | `size()` |
| 전체 삭제 | `clear()` | `clear()` | `clear()` |

---

## 7️⃣ 사용 예시 비교

| 구조 | 예시 상황 |
|------|------------|
| `List` | 순서가 있는 항목 목록 (장바구니, 댓글 목록 등) |
| `Set` | 중복 없는 데이터 (태그, 유저 ID 등) |
| `Map` | key로 빠르게 찾을 수 있는 데이터 (회원 정보, 설정 값 등) |

---

## 8️⃣ List vs Set vs Map 요약

| 특징 | List | Set | Map |
|------|------|-----|-----|
| 순서 | 있음 | 없음 (`LinkedHashSet`은 있음) | key 순서 없음 |
| 중복 | 허용 | 허용 안 됨 | key 중복 불가, value 가능 |
| 인덱스 접근 | 가능 | 불가 | key로 접근 |
| 주요 구현체 | `ArrayList` | `HashSet` | `HashMap` |

---

## 🎯 정리

✔ `List` → 순서 O, 중복 O (ex: ArrayList)  
✔ `Set` → 순서 X, 중복 X (ex: HashSet)  
✔ `Map` → key-value 구조, key는 중복 X  
✔ `ArrayList` → 가장 일반적인 순차형 리스트  
✔ `HashSet` → 빠른 검색 + 중복 제거  
✔ `HashMap` → key로 빠르게 값 조회  
✔ 반복 시 `for-each`, `entrySet()`, `forEach()` 등 활용  
✔ 사용 목적에 따라 적절한 컬렉션을 선택할 것!

