# 🧠 Computer Architecture - 컴퓨터 구조 개념 정리

컴퓨터 시스템의 하드웨어 구성과 내부 동작 원리에 대한 이론을 정리합니다.  
CPU, 메모리, 연산 장치, 버스 구조, 캐시, 파이프라인 등 시스템 하위 구조에 대한 개념을 포함합니다.

---

### 🧱 기본 개념
| 주제 | 파일명 | 설명 |
|------|--------|------|
| 컴퓨터 구조 개요 | [what-is-computer-architecture.md](./notes/what-is-computer-architecture.md) | 컴퓨터 시스템 구조 전반 요약 |
| 폰 노이만 구조 | [von-neumann.md](./notes/von-neumann.md) | 메모리와 연산 장치를 연결한 기본 구조 |
| 명령어 사이클 | [instruction-cycle.md](./notes/instruction-cycle.md) | fetch-decode-execute 사이클 |

---

### ⚙️ CPU 구조와 동작
| 주제 | 파일명 | 설명 |
|------|--------|------|
| 레지스터 | [registers.md](./notes/registers.md) | CPU 내 임시 저장소 |
| 산술 논리 연산 장치 (ALU) | [alu.md](./notes/alu.md) | 산술 및 논리 연산 수행 |
| 제어 장치 | [control-unit.md](./notes/control-unit.md) | 명령어 해석 및 신호 생성 |
| 파이프라인 | [pipeline.md](./notes/pipeline.md) | 명령어 병렬 처리 기술 |
| 명령어 집합 구조 (ISA) | [isa.md](./notes/isa.md) | CPU와 프로그램 간 인터페이스 정의 |

---

### 🧠 메모리와 저장 장치
| 주제 | 파일명 | 설명 |
|------|--------|------|
| 메모리 계층 구조 | [memory-hierarchy.md](./notes/memory-hierarchy.md) | 레지스터 ~ 디스크까지의 계층적 구조 |
| 캐시 메모리 | [cache.md](./notes/cache.md) | 고속 메모리, 지역성, 캐시 미스 |
| 주기억장치 (RAM/ROM) | [main-memory.md](./notes/main-memory.md) | RAM, DRAM, ROM 개념 정리 |
| 보조 기억장치 | [storage.md](./notes/storage.md) | HDD, SSD, 저장장치 인터페이스 |

---

### 🔌 입출력 및 버스
| 주제 | 파일명 | 설명 |
|------|--------|------|
| I/O 시스템 | [io-system.md](./notes/io-system.md) | 입출력 장치와 제어 방식 |
| 버스 구조 | [bus.md](./notes/bus.md) | 데이터, 주소, 제어 버스 설명 |
| 인터럽트 | [interrupt.md](./notes/interrupt.md) | CPU가 처리하는 비동기 이벤트 |

---

### 📊 성능 및 병렬성
| 주제 | 파일명 | 설명 |
|------|--------|------|
| CPI, MIPS | [performance.md](./notes/performance.md) | 명령어 수와 성능 지표 계산 |
| 병렬 처리 구조 | [parallelism.md](./notes/parallelism.md) | 멀티코어, SIMD, MIMD 등 |
| 슈퍼스칼라 & VLIW | [superscalar-vliw.md](./notes/superscalar-vliw.md) | 명령어 수준 병렬성 구현 방식 |
