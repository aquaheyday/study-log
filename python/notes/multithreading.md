# 🚀 Python 멀티스레딩 (Multithreading)

Python에서 **멀티스레딩(Multithreading)** 을 사용하면 **여러 작업을 동시에 실행**할 수 있습니다.  
이 문서에서는 **`threading` 모듈을 활용한 멀티스레딩 구현 및 동기화 기법**을 설명합니다.

---

## 1. 멀티스레딩이란?

- **멀티스레딩(Multithreading)** 은 **하나의 프로세스에서 여러 개의 작업(스레드)을 동시에 실행**하는 기법입니다.
- CPU를 활용하는 작업보다는 **I/O 작업(파일 처리, 네트워크 요청 등)에 적합**합니다.
- Python의 **GIL(Global Interpreter Lock)로 인해 멀티코어 CPU 활용에는 제약**이 있습니다.

✅ **멀티스레딩 적합한 경우**: 파일 입출력, 웹 크롤링, 네트워크 요청, GUI 프로그램  
❌ CPU 집중 연산(이미지 처리, 데이터 분석)에는 적합하지 않음 → **멀티프로세싱 사용 권장**

---

## 2. `threading` 모듈로 멀티스레드 구현

Python에서 기본 제공하는 `threading` 모듈을 사용하여 멀티스레드를 실행할 수 있습니다.

### 기본적인 스레드 생성 (`threading.Thread`)
```python
import threading

def task():
    print("스레드 실행!")

# 스레드 생성
thread = threading.Thread(target=task)
thread.start()

# 메인 스레드에서 실행
print("메인 스레드 실행 완료!")
```
✔ `threading.Thread(target=함수)` → 새로운 스레드 생성  
✔ `thread.start()` → 스레드 실행  
✔ 스레드는 **비동기적으로 실행**되므로 메인 스레드가 먼저 종료될 수도 있음  

---

## 3. 여러 개의 스레드 실행

여러 개의 스레드를 실행하면 동시에 여러 작업을 수행할 수 있습니다.

```python
import threading
import time

def worker(n):
    print(f"스레드 {n} 시작")
    time.sleep(2)
    print(f"스레드 {n} 종료")

# 여러 개의 스레드 생성
threads = []
for i in range(3):
    thread = threading.Thread(target=worker, args=(i,))
    threads.append(thread)
    thread.start()

# 모든 스레드가 종료될 때까지 대기
for thread in threads:
    thread.join()

print("모든 스레드 종료!")
```
✔ `args=(i,)` → 스레드 함수에 인자 전달  
✔ `thread.join()` → 모든 스레드가 종료될 때까지 메인 스레드 대기  

---

## 4. 스레드 동기화 (`Lock` 사용)

멀티스레딩 환경에서는 **여러 스레드가 동시에 자원(변수, 파일 등)을 변경하면 충돌**이 발생할 수 있습니다.  
이를 방지하기 위해 **`Lock` (뮤텍스, Mutex)** 을 사용하여 **한 번에 하나의 스레드만 특정 코드 블록을 실행**하도록 할 수 있습니다.

### `Lock`을 사용한 동기화 예제
```python
import threading

lock = threading.Lock()
counter = 0

def increase():
    global counter
    for _ in range(100000):
        with lock:  # 한 번에 한 스레드만 접근 가능
            counter += 1

# 여러 개의 스레드 실행
threads = []
for _ in range(5):
    thread = threading.Thread(target=increase)
    threads.append(thread)
    thread.start()

# 모든 스레드 대기
for thread in threads:
    thread.join()

print("최종 카운터 값:", counter)
```
✔ `with lock:` → 해당 블록에서 **하나의 스레드만 실행 가능**  
✔ `counter`를 여러 스레드가 동시에 증가시키면 값이 꼬일 수 있음 → **Lock으로 해결**  

---

## 5. `ThreadPoolExecutor`를 사용한 간편한 멀티스레딩

Python의 **`concurrent.futures.ThreadPoolExecutor`** 를 사용하면 **더 간편하게 멀티스레딩을 구현**할 수 있습니다.

### `ThreadPoolExecutor` 사용 예제
```python
import concurrent.futures
import time

def task(n):
    time.sleep(2)
    return f"작업 {n} 완료"

with concurrent.futures.ThreadPoolExecutor(max_workers=3) as executor:
    results = executor.map(task, range(5))

for result in results:
    print(result)
```
✔ `max_workers=3` → **최대 3개의 스레드 동시 실행**  
✔ `executor.map(task, range(5))` → 여러 작업을 **비동기적으로 실행**  

---

## 6. GIL (Global Interpreter Lock) 이슈

- Python은 **GIL(Global Interpreter Lock)** 이 존재하여 **멀티코어 CPU를 활용한 병렬 처리에 제한**이 있음.
- GIL은 **한 번에 하나의 스레드만 Python 바이트코드를 실행하도록 강제**함.
- 따라서 **CPU 연산이 많은 작업은 멀티스레딩보다 멀티프로세싱이 더 적합**함.

✅ **멀티스레딩 적합한 경우**: **I/O 바운드 작업** (웹 크롤링, 파일 읽기/쓰기, 네트워크 요청)  
❌ **CPU 바운드 작업** (이미지 처리, 머신러닝 연산) → `multiprocessing` 사용 권장  

---

## 7. 멀티스레딩 vs 멀티프로세싱

멀티스레딩(`threading`)과 멀티프로세싱(`multiprocessing`)은 목적이 다릅니다.

| 구분 | 멀티스레딩 | 멀티프로세싱 |
|------|-----------|-------------|
| 실행 단위 | 스레드(Thread) | 프로세스(Process) |
| GIL 영향 | **받음** (한 번에 하나의 스레드 실행) | **받지 않음** (여러 프로세스가 개별 실행) |
| 적합한 작업 | **I/O 바운드 작업** (파일, 네트워크) | **CPU 바운드 작업** (데이터 연산, AI) |
| 메모리 공유 | **같은 메모리 공유** | **독립적인 메모리 사용** |

✔ **CPU 연산이 많은 작업** → `multiprocessing` 사용  
✔ **네트워크, 파일 처리 등 I/O 작업** → `threading` 사용  

---

## 🎯 정리

✔ **멀티스레딩(Multithreading)** → 여러 작업을 동시에 실행  
✔ **`threading.Thread`** → 새로운 스레드 생성  
✔ **`thread.start()`** → 스레드 실행  
✔ **`thread.join()`** → 모든 스레드 종료 대기  
✔ **`Lock`** → 공유 자원 접근 동기화  
✔ **`ThreadPoolExecutor`** → 더 간편한 멀티스레딩  
✔ **GIL 영향** → **CPU 작업은 멀티스레딩보다 멀티프로세싱 사용 권장**  

