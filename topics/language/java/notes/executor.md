# ⚙️ Java - Executor (스레드 풀)

`Executor`는 **자바의 멀티스레드 작업을 관리하기 위한 프레임워크**입니다.  
개별 스레드를 직접 생성하지 않고, **스레드 풀(Thread Pool)** 을 통해 작업을 효율적으로 실행하고 관리할 수 있게 도와줍니다.

---

## 1️⃣ 왜 Executor를 사용할까?

기존 방식:

```java
Thread t = new Thread(() -> doWork());
t.start();
```

✔ 스레드를 직접 만들면 **관리 힘듦 + 성능 저하 가능성↑**, 대량 작업에 적합하지 않음
✔ 해결책: Executor는 **스레드를 재사용하고 제한적으로 운영**함

---

## 2️⃣ 기본 구조

```java
Executor executor = Executors.newSingleThreadExecutor();
executor.execute(() -> {
    System.out.println("작업 실행 중");
});
```

✔ `execute(Runnable)` 메서드를 사용해 작업 제출

---

## 3️⃣ `ExecutorService` 사용

`ExecutorService`는 `Executor`의 하위 인터페이스로 **작업 종료, Future 반환, shutdown 등 제어 기능**을 포함합니다.

```java
ExecutorService service = Executors.newFixedThreadPool(3);

service.execute(() -> {
    System.out.println("작업 1");
});
```

---

## 4️⃣ 다양한 스레드 풀 생성 방법

```java
Executors.newSingleThreadExecutor();  // 스레드 1개 고정
Executors.newFixedThreadPool(n);      // 스레드 n개 고정
Executors.newCachedThreadPool();      // 필요한 만큼 생성 & 재사용
Executors.newScheduledThreadPool(n);  // 일정 시간 간격 작업
```

---

## 5️⃣ 작업 결과 받기 - `Callable` & `Future`

`Runnable`은 결과 반환 ❌  
`Callable`은 결과 반환 ✅ (`call()` 메서드)

```java
Callable<Integer> task = () -> {
    Thread.sleep(1000);
    return 42;
};
```

```java
Future<Integer> future = service.submit(task);
Integer result = future.get();  // 블로킹
```

---

## 6️⃣ `shutdown()` vs `shutdownNow()`

```java
service.shutdown();      // 기존 작업은 마치고 종료
// service.shutdownNow(); // 실행 중인 작업도 중단 요청 (강제)
```

- `shutdown()` → 안전하게 종료  
- `shutdownNow()` → 긴급 중단 요청 (정말 필요한 경우만!)

---

## 7️⃣ `invokeAll()` - 여러 Callable 처리

```java
List<Callable<Integer>> tasks = List.of(
    () -> 1, () -> 2, () -> 3
);

List<Future<Integer>> results = service.invokeAll(tasks);
for (Future<Integer> f : results) {
    System.out.println(f.get());
}
```

---

## 8️⃣ `ScheduledExecutorService` - 예약 실행

```java
ScheduledExecutorService scheduler = Executors.newScheduledThreadPool(1);

scheduler.schedule(() -> {
    System.out.println("3초 후 실행");
}, 3, TimeUnit.SECONDS);
```

---

## 9️⃣ 예외 처리

`Runnable`의 예외는 사라지지만  
`Callable` + `Future.get()`은 예외를 `ExecutionException`으로 감쌈

```java
try {
    future.get();
} catch (ExecutionException e) {
    Throwable cause = e.getCause();
    cause.printStackTrace();
}
```

---

## 🔟 스레드 풀 직접 설정 - `ThreadPoolExecutor`

```java
ExecutorService customPool = new ThreadPoolExecutor(
    2,               // corePoolSize
    4,               // maxPoolSize
    60,              // idle timeout
    TimeUnit.SECONDS,
    new LinkedBlockingQueue<>()
);
```

✔ 고급 설정이 필요할 때 `ThreadPoolExecutor` 직접 사용

---

## 🎯 정리

✔ `Executor` → 스레드 생성을 위임하는 인터페이스  
✔ `ExecutorService` → 스레드 풀 + 작업 제어 기능 포함  
✔ `submit(Callable)` → 결과 반환, `Future`로 받음  
✔ `shutdown()` → 작업 종료  
✔ `newFixedThreadPool(n)` → 스레드 n개 고정  
✔ `newCachedThreadPool()` → 가변적 스레드 재사용  
✔ `invokeAll()` → 여러 작업 일괄 처리  
✔ `ScheduledExecutorService` → 지연/반복 작업 예약  
✔ `ThreadPoolExecutor` → 세부 설정 가능  
✔ 실무에서는 `ExecutorService`로 멀티스레드를 안전하게 관리!

