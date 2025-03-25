# 🧰 Java - 동기화(Synchronization) & Lock

멀티스레드 환경에서는 **여러 스레드가 동시에 공유 자원에 접근**할 수 있기 때문에,  
데이터 충돌과 오류를 막기 위해 **동기화(Synchronization)**가 필요합니다.

---

## 1️⃣ 동기화(Synchronization)란?

- 여러 스레드가 **공유 자원에 동시에 접근하지 못하도록 제한**하는 것  
- **경쟁 조건(Race Condition)** 을 방지

---

## 2️⃣ `synchronized` 키워드

- 공유 자원 접근을 **하나의 스레드만 허용**하는 키워드

#### 1. 메서드 전체 동기화

```java
public synchronized void print() {
    // 임계 영역
}
```

#### 2. 블록 단위 동기화

```java
public void print() {
    synchronized (this) {
        // 임계 영역
    }
}
```

✔ `this` → 현재 인스턴스를 lock 대상으로 사용

---

## 3️⃣ 정적 메서드 동기화 (`static synchronized`)

```java
public static synchronized void log() {
    // 클래스 단위 lock
}
```

✔`static` 메서드는 클래스 레벨에서 lock이 걸립니다 → `synchronized(ClassName.class)`와 동일

---

## 4️⃣ 임계 영역(Critical Section)

- **동시에 하나의 스레드만 실행 가능한 코드 영역**
- 동기화를 통해 보호해야 할 부분

```java
synchronized (sharedObject) {
    // 임계 구역: 공유 자원 접근
}
```

---

## 5️⃣ 경쟁 조건(Race Condition)

- 여러 스레드가 **동시에 값을 읽고, 쓰는 과정에서 충돌** 발생
- 동기화가 없을 경우 데이터가 일관성 없이 변경될 수 있음

---

## 6️⃣ 동기화 미적용 카운터 예제

```java
class Counter {
    int count = 0;

    void increment() {
        count++;
    }
}
```

✔ 여러 스레드에서 `increment()`를 동시에 호출하면 `count` 값이 잘못될 수 있음 (Race Condition 발생)

---

## 7️⃣ 동기화 적용 카운터터 예제

```java
class Counter {
    int count = 0;

    synchronized void increment() {
        count++;
    }
}
```

✔ 이제 하나의 스레드만 `increment()`를 실행할 수 있어 안정적

---

## 8️⃣ `java.util.concurrent.locks.Lock` 인터페이스

- `synchronized`보다 **더 유연한 제어** 가능
- 명시적으로 lock을 걸고 해제해야 함

```java
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

class Counter {
    private int count = 0;
    private final Lock lock = new ReentrantLock();

    public void increment() {
        lock.lock();  // 🔒 Lock 획득
        try {
            count++;
        } finally {
            lock.unlock();  // 🔓 반드시 해제
        }
    }
}
```

---

## 9️⃣ `ReentrantLock` 특징

| 기능 | 설명 |
|------|------|
| 재진입 가능 | 한 스레드가 여러 번 lock 획득 가능 |
| `tryLock()` | 대기하지 않고 즉시 lock 시도 |
| `lockInterruptibly()` | 대기 중 인터럽트 가능 |
| `fair` 옵션 | 먼저 기다린 스레드에게 lock 제공 (공정성)

---

## 🔟 `volatile` 키워드

- **변수 값이 캐시되지 않도록 보장**
- 모든 스레드가 **메인 메모리의 값을 직접 읽음**
- 하지만 `volatile`만으로는 복합 연산(count++) 보호 불가 → `synchronized` 필요

```java
private volatile boolean running = true;
```

---

## 🎯 정리

✔ `synchronized` → 공유 자원에 대한 동시 접근을 방지  
✔ 메서드 또는 블록 단위로 사용 가능  
✔ `Race Condition` → 동기화하지 않으면 발생하는 충돌 문제  
✔ `Lock` (`ReentrantLock`) → 보다 세밀한 락 제어  
✔ `lock()` → 락 획득 / `unlock()` → 해제 (항상 `finally`에서)  
✔ `tryLock()` → 락 점유 실패 시 대기하지 않음  
✔ `volatile` → 변수 값의 메모리 일관성 유지 (동기화는 아님)  
✔ 실무에서는 `synchronized` 보다 `Lock`, `ExecutorService` 등 고급 API 사용 권장

