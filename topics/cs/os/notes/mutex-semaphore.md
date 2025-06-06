# 🔐 Mutex vs Semaphore

운영체제에서 **동기화(Synchronization)** 를 구현하기 위해 사용하는 대표적인 도구는 **뮤텍스(Mutex)** 와 **세마포어(Semaphore)** 입니다.  
이 두 개념은 비슷해 보이지만, **동작 방식, 용도, 구조**에서 차이가 있습니다.

---

## 1️⃣ 기본 정의

| 항목       | Mutex (뮤텍스) | Semaphore (세마포어) |
|------------|----------------|----------------------|
| **정의**   | 상호 배제를 위한 **잠금 객체** | 동기화 및 자원 제어용 **카운터 변수** |
| **기본 값** | 0 또는 1 (Binary) | 정수 (0 이상) |
| **주 용도** | 임계구역 보호 (1개 스레드만 접근) | 자원 수 제한 (N개까지 접근 가능) |

---

## 2️⃣ 동작 방식

| 항목           | Mutex            | Semaphore               |
|----------------|------------------|--------------------------|
| **소유권**     | **스레드가 소유** (Owner가 존재) | 소유권 없음 (누구나 signal 가능) |
| **lock/unlock**| 반드시 쌍으로 사용 (`lock`, `unlock`) | `wait(P)`, `signal(V)`로 조작 |
| **동작 예시**  | 하나의 스레드가 락을 걸고, 해제 | 여러 스레드가 동시에 접근 시, 자원 수 감소/복구 |

---

## 3️⃣ 종류

| 항목       | Mutex         | Semaphore                       |
|------------|---------------|----------------------------------|
| **종류**   | 일반적으로 1가지 | Binary Semaphore (0 or 1)<br>Counting Semaphore (0~N) |
| **Binary** | 기본적으로 Binary | Binary일 경우, Mutex와 유사 동작 |

---

## 4️⃣ 특징 비교

| 항목                | Mutex                         | Semaphore                          |
|---------------------|-------------------------------|-------------------------------------|
| **자원 보호 개수**   | 1개                           | N개 (복수 자원 관리 가능)          |
| **동기화 방식**      | 상호 배제만                  | 상호 배제 + 동기화 가능            |
| **Deadlock 위험**    | 있음 (lock 후 unlock 누락 시) | 있음 (wait/signal 불균형 시)       |
| **우선순위 반전**    | 발생 가능                    | 발생 가능                           |

---

## 5️⃣ 예시 코드 (C 스타일)

### Mutex 예시

```c
pthread_mutex_t lock;

void* thread_func(void* arg) {
    pthread_mutex_lock(&lock);
    // 임계구역
    pthread_mutex_unlock(&lock);
}
```

---

### Semaphore 예시

```c
sem_t sem;

void* thread_func(void* arg) {
    sem_wait(&sem);   // wait(P)
    // 임계구역
    sem_post(&sem);   // signal(V)
}
```

---

## 🎯 정리 요약

| 구분        | Mutex                          | Semaphore                      |
|-------------|---------------------------------|---------------------------------|
| **목적**     | 단일 임계구역 보호              | 다중 접근 제어 / 자원 동기화   |
| **동작 단위** | 1개의 스레드                   | 여러 스레드 또는 프로세스      |
| **소유권**   | 있음 (lock 주체만 unlock 가능)  | 없음 (누구나 signal 가능)      |
| **주 사용처** | 단일 스레드 간 상호 배제       | N개 제한 자원 접근 제어        |

✔ **Mutex**: 하나의 스레드만 접근 가능한 잠금  
✔ **Semaphore**: 여러 개 자원을 제어하는 카운터  
✔ 둘 다 임계구역 보호 가능, 하지만 목적과 스케일이 다름  
✔ 세마포어는 뮤텍스보다 더 일반화된 동기화 메커니즘

