# 🔄 컨텍스트 스위칭(Context Switching)

**컨텍스트 스위칭은 CPU가 현재 실행 중인 작업의 상태를 저장하고, 다른 작업으로 전환하는 과정입니다.  
멀티태스킹을 구현하기 위한 필수 기능이며, 과도한 전환은 오버헤드를 유발할 수 있습니다.**

---

## 1️⃣ 컨텍스트 스위칭이란?

| 항목       | 설명 |
|------------|------|
| 정의       | **CPU가 현재 실행 중인 작업의 상태(Context)를 저장하고, 다른 작업으로 전환하는 과정** |
| 목적       | 멀티태스킹 구현, 사용자 응답성 향상 |
| 발생 시점  | CPU가 다른 프로세스 또는 스레드를 실행해야 할 때 |
| 전환 대상  | 프로세스 ↔ 프로세스, 스레드 ↔ 스레드 등 |

---

## 2️⃣ 왜 필요한가?

- 여러 프로세스/스레드를 **동시에 실행되는 것처럼 보이게** 하기 위해  
- I/O 등으로 **대기 중인 작업 대신 다른 작업을 실행**하기 위해  
- 사용자 응답성을 높이고 **시스템 자원을 효율적으로 활용**하기 위해

---

## 3️⃣ 컨텍스트 스위칭 과정

```
[실행 중인 작업 A]
      ↓ (스케줄러 개입)
→ 작업 A 상태 저장
→ 작업 B 상태 복원
      ↓
[작업 B 실행 시작]
```

> 이 과정에서 레지스터, 스택 포인터, 프로그램 카운터 등의 정보가 저장되고 복원됨

---

## 4️⃣ 프로세스 vs 스레드 컨텍스트 스위칭

| 항목           | 프로세스 전환             | 스레드 전환 (동일 프로세스 내) |
|----------------|---------------------------|---------------------------------|
| 전환 비용       | 높음                      | 낮음                            |
| 메모리 영역     | 완전히 다름                | 대부분 공유                     |
| 문맥 정보       | PCB 전체 저장              | 레지스터, 스택 등만 저장        |
| 캐시 영향       | 클 수 있음                 | 적음                            |

---

## 5️⃣ 컨텍스트 스위칭의 오버헤드

| 오버헤드 항목     | 설명 |
|-------------------|------|
| CPU 시간 낭비      | 저장/복원 작업 자체가 자원 소모 |
| 캐시 무효화        | CPU 캐시가 초기화될 수 있어 성능 저하 |
| 지연               | 스케줄링 + 문맥 처리로 작업 전환 시간 증가 |

---

## 6️⃣ 최소화 방법

| 방법                         | 설명 |
|------------------------------|------|
| 스레드 기반 처리              | 프로세스보다 전환 비용이 낮음 |
| 작업 묶기 (Batch 처리)        | 전환 횟수 자체를 줄임 |
| CPU Affinity 설정             | 같은 작업이 같은 코어에서 계속 실행되도록 설정 |
| 비동기 I/O 사용               | I/O 대기 줄이고 CPU 낭비 최소화 |

---

## 🎯 정리 요약

✔ 컨텍스트 스위칭은 **CPU가 작업을 바꾸는 필수 동작**  
✔ 프로세스 전환은 **스레드 전환보다 무겁다**  
✔ 너무 잦은 스위칭은 **성능 저하**를 유발할 수 있음  
✔ **멀티스레딩, affinity, 최적화된 스케줄링**을 통해 효율성 향상 가능
