# 🌍 Python 비동기 프로그래밍 (Asynchronous Programming)

Python에서 **비동기 프로그래밍(Asynchronous Programming)** 은 **I/O 작업을 병렬로 실행하여 성능을 향상**시키는 기법입니다.  
이를 위해 `asyncio`, `async/await`, `aiohttp` 등의 라이브러리를 사용할 수 있습니다.

---

## 1. 비동기 프로그래밍이란?

- 일반적인 **동기(Synchronous) 프로그래밍**은 **하나의 작업이 끝나야 다음 작업이 실행**됨.
- **비동기(Asynchronous) 프로그래밍**은 **여러 작업을 동시에 실행**할 수 있음.
- 네트워크 요청, 파일 입출력, 데이터베이스 조회 등의 **I/O 작업을 최적화**하는 데 유용함.

### 동기 vs. 비동기 예제

#### 동기(Synchronous) 코드
```python
import time

def task(name, delay):
    time.sleep(delay)
    print(f"{name} 완료!")

task("작업 1", 2)
task("작업 2", 2)
print("모든 작업 완료!")
```
#### 출력 결과
```
작업 1 완료! (2초 후)
작업 2 완료! (4초 후)
모든 작업 완료! (총 4초)
```

---

#### 비동기(Asynchronous) 코드
```python
import asyncio

async def task(name, delay):
    await asyncio.sleep(delay)  # 비동기 대기
    print(f"{name} 완료!")

async def main():
    await asyncio.gather(
        task("작업 1", 2),
        task("작업 2", 2)
    )
    print("모든 작업 완료!")

asyncio.run(main())  # 비동기 실행
```
#### 출력 결과
```
작업 1 완료! (2초 후)
작업 2 완료! (2초 후)
모든 작업 완료! (총 2초)
```
✔ `async/await`을 사용하면 작업을 동시에 실행할 수 있어 실행 시간이 줄어듭니다.

---

## 2. `async/await` 기본 개념

### `async` 함수 정의
```python
async def my_function():
    print("비동기 함수 실행")
```
✔ `async` 키워드를 사용하면 **비동기 함수(코루틴)** 가 됩니다.

---

### `await` 키워드 사용
```python
import asyncio

async def task():
    print("작업 시작")
    await asyncio.sleep(2)  # 비동기 대기 (2초)
    print("작업 완료!")

asyncio.run(task())
```
✔ `await` 키워드를 사용하면 **비동기적으로 대기**하면서 다른 작업을 실행할 수 있습니다.

---

## 3. 여러 개의 작업 동시 실행 (`asyncio.gather()`)

여러 개의 비동기 함수를 동시에 실행할 때 `asyncio.gather()`를 사용하면 성능을 향상할 수 있습니다.

### 여러 작업 동시 실행
```python
import asyncio

async def task(name, delay):
    await asyncio.sleep(delay)
    print(f"{name} 완료!")

async def main():
    await asyncio.gather(
        task("작업 1", 2),
        task("작업 2", 3),
        task("작업 3", 1)
    )

asyncio.run(main())
```
#### 출력 결과
```
작업 3 완료! (1초 후)
작업 1 완료! (2초 후)
작업 2 완료! (3초 후)
```
✔ `asyncio.gather()`를 사용하면 **모든 작업을 동시에 실행**할 수 있음.

---

## 4. `aiohttp`를 사용한 비동기 HTTP 요청

기본적으로 `requests` 라이브러리는 **동기(Synchronous)** 방식이므로  
**비동기(Asynchronous) HTTP 요청**을 위해 `aiohttp` 라이브러리를 사용합니다.

### `aiohttp` 설치
```sh
pip install aiohttp
```

---

### 비동기 HTTP 요청 (`aiohttp`)
```python
import aiohttp
import asyncio

async def fetch(url):
    async with aiohttp.ClientSession() as session:
        async with session.get(url) as response:
            return await response.text()  # 응답 데이터 반환

async def main():
    urls = [
        "https://jsonplaceholder.typicode.com/posts/1",
        "https://jsonplaceholder.typicode.com/posts/2",
        "https://jsonplaceholder.typicode.com/posts/3"
    ]
    
    responses = await asyncio.gather(*(fetch(url) for url in urls))
    for i, data in enumerate(responses, 1):
        print(f"응답 {i}: {data[:50]}...")  # 첫 50글자만 출력

asyncio.run(main())
```
### 결과 (세 개의 URL을 동시에 요청하여 빠르게 응답받음)
```
응답 1: { "userId": 1, "id": 1, "title": "... 
응답 2: { "userId": 1, "id": 2, "title": "... 
응답 3: { "userId": 1, "id": 3, "title": "... 
```
✔ `aiohttp` + `asyncio.gather()`를 활용하면 **여러 개의 HTTP 요청을 병렬로 실행 가능!**  

---

## 5. 비동기 태스크 관리 (`asyncio.create_task()`)

비동기 태스크를 백그라운드에서 실행하려면 `asyncio.create_task()`를 사용할 수 있음.

### `asyncio.create_task()` 사용 예제
```python
import asyncio

async def background_task():
    await asyncio.sleep(2)
    print("백그라운드 작업 완료!")

async def main():
    task = asyncio.create_task(background_task())  # 백그라운드 실행
    print("메인 작업 실행 중...")
    await asyncio.sleep(1)  # 메인 작업 대기
    print("메인 작업 완료!")
    await task  # 백그라운드 작업 대기

asyncio.run(main())
```
#### 출력 결과
```
메인 작업 실행 중...
메인 작업 완료!
백그라운드 작업 완료!
```
✔ `create_task()`를 사용하면 특정 작업을 **백그라운드에서 실행** 가능.  

---

## 6. 비동기 큐 (`asyncio.Queue`)

비동기 작업 간 데이터 공유를 위해 **큐(Queue)** 를 사용할 수 있음.

### `asyncio.Queue` 사용 예제
```python
import asyncio

async def producer(queue):
    for i in range(5):
        await asyncio.sleep(1)
        await queue.put(i)
        print(f"데이터 {i} 생성")

async def consumer(queue):
    while True:
        item = await queue.get()
        print(f"데이터 {item} 처리 완료!")
        queue.task_done()

async def main():
    queue = asyncio.Queue()
    
    # 프로듀서 & 컨슈머 동시 실행
    await asyncio.gather(producer(queue), consumer(queue))

asyncio.run(main())
```
✔ `queue.put(item)` → 데이터를 큐에 추가  
✔ `queue.get()` → 큐에서 데이터 가져오기  

---

## 🎯 정리

✔ **`async/await` 사용** → 비동기 함수를 정의하고 실행  
✔ **`asyncio.gather()`** → 여러 작업을 동시에 실행  
✔ **`aiohttp`** → 비동기 HTTP 요청 처리  
✔ **`asyncio.create_task()`** → 백그라운드에서 태스크 실행  
✔ **`asyncio.Queue`** → 비동기 데이터 공유  
