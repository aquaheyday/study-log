# 🖥️ Operating System - 운영체제 개념 정리

운영체제(OS)는 하드웨어와 사용자 프로그램 사이의 중재자 역할을 하며,  
프로세스 관리, 메모리 관리, 파일 시스템, 입출력, 동기화 등의 핵심 기능을 제공합니다.  
이 디렉토리는 운영체제의 주요 개념들을 분류별로 정리합니다.

---

### 🧠 운영체제 기초
| 주제 | 파일명 | 설명 |
|------|--------|------|
| 운영체제란? | [what-is-os.md](./notes/what-is-os.md) | OS의 역할, 종류, 커널 구조 |
| 시스템 콜 | [system-call.md](./notes/system-call.md) | 사용자와 커널 간 인터페이스 |
| 커널 구조 | [kernel.md](./notes/kernel.md) | 모놀리식 vs 마이크로커널 구조 |
| 사용자 모드 / 커널 모드 | [user-kernel-mode.md](./notes/user-kernel-mode.md) | CPU의 권한 모드 개념 |

---

### ⚙️ 프로세스 & 스레드
| 주제 | 파일명 | 설명 |
|------|--------|------|
| 프로세스 개념 | [process.md](./notes/process.md) | 프로세스의 상태, 구조, PCB |
| 스레드 개념 | [thread.md](./notes/thread.md) | 경량 프로세스, 멀티스레딩 |
| 컨텍스트 스위칭 | [context-switching.md](./notes/context-switching.md) | CPU가 다른 프로세스로 전환하는 과정 |
| 멀티프로세싱 & 멀티스레딩 | [multi-processing-threading.md](./notes/multi-processing-threading.md) | 병렬성과 자원 분리 차이점 |
| 프로세스 간 통신 (IPC) | [ipc.md](./notes/ipc.md) | 파이프, 메시지 큐, 공유 메모리 |

---

### ⏱️ CPU 스케줄링
| 주제 | 파일명 | 설명 |
|------|--------|------|
| 스케줄링 개요 | [scheduling.md](./notes/scheduling.md) | 프로세스를 실행 순서대로 정하는 전략 |
| 스케줄링 알고리즘 | [scheduling-algorithms.md](./notes/scheduling-algorithms.md) | FCFS, SJF, RR, Priority 등 비교 |

---

### 🧮 메모리 관리
| 주제 | 파일명 | 설명 |
|------|--------|------|
| 메모리 구조 개요 | [memory-structure.md](./notes/memory-structure.md) | 스택, 힙, 데이터, 코드 영역 |
| 가상 메모리 | [virtual-memory.md](./notes/virtual-memory.md) | 페이징, 세그멘테이션, 주소 변환 |
| 페이징 & 페이지 교체 | [paging.md](./notes/paging.md) | 페이지 테이블, 교체 알고리즘 |
| 캐시 메모리 | [memory-cache.md](./notes/memory-cache.md) | 메모리 접근 속도 향상 구조 |

---

### 💾 파일 시스템
| 주제 | 파일명 | 설명 |
|------|--------|------|
| 파일 시스템 개요 | [file-system.md](./notes/file-system.md) | 디스크에 데이터 저장하는 구조 |
| 디스크 스케줄링 | [disk-scheduling.md](./notes/disk-scheduling.md) | SSTF, SCAN 등 디스크 접근 방식 |
| inode & 디렉토리 구조 | [inode.md](./notes/inode.md) | 리눅스 계열 파일 시스템 구조 |

---

### 🔐 동기화 & 병행성
| 주제 | 파일명 | 설명 |
|------|--------|------|
| 임계구역과 상호 배제 | [critical-section.md](./notes/critical-section.md) | 동시성 문제를 해결하는 방법 |
| 뮤텍스 vs 세마포어 | [mutex-semaphore.md](./notes/mutex-semaphore.md) | 대표적인 동기화 기법 비교 |
| 데드락 | [deadlock.md](./notes/deadlock.md) | 교착 상태의 원인과 해결 방법 |
| 스핀락 & 모니터 | [spinlock-monitor.md](./notes/spinlock-monitor.md) | 커널 수준 동기화 방식 |
