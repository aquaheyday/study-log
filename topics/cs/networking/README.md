# 🌐 Networking - 네트워크 개념 정리

컴퓨터 네트워크의 기본 개념부터 TCP/IP, HTTP, 소켓, DNS 등 실무와 연결된 기술들을 계층별로 정리합니다.  
OSI 7계층, 프로토콜, 보안, 성능 최적화까지 아우릅니다.

---

### 🧱 네트워크 기초
| 주제 | 파일명 | 설명 |
|------|--------|------|
| 네트워크란? | [networking-basics.md](./notes/what-is-networking.md) | 컴퓨터 간 통신의 원리와 개요 |
| OSI 7계층 모델 | [osi-model.md](./notes/osi-model.md) | 계층별 역할, 프로토콜 정리 |
| TCP/IP 모델 | [tcp-ip-model.md](./notes/tcp-ip-model.md) | 실무에서 사용하는 4계층 구조 |
| 프로토콜 개요 | [protocols.md](./notes/protocols.md) | TCP, UDP, ICMP 등 주요 프로토콜 개요 |

---

### 🔗 링크 계층 & 물리 계층
| 주제 | 파일명 | 설명 |
|------|--------|------|
| MAC 주소 | [mac-address.md](./notes/mac-address.md) | 하드웨어 식별 주소 |
| ARP 프로토콜 | [arp.md](./notes/arp.md) | IP → MAC 주소 매핑 방식 |
| 이더넷 | [ethernet.md](./notes/ethernet.md) | LAN 기반 데이터 전송 기술 |

---

### 📦 네트워크 계층 (IP)
| 주제 | 파일명 | 설명 |
|------|--------|------|
| IP 주소와 서브넷 | [ip-subnet.md](./notes/ip-subnet.md) | IPv4, IPv6, 서브넷 마스크 개념 |
| 라우팅 | [routing.md](./notes/routing.md) | 패킷의 목적지 결정 과정 |
| NAT / 포트포워딩 | [nat.md](./notes/nat.md) | 사설 IP와 공인 IP 연결 방식 |

---

### 🧭 전송 계층 (TCP/UDP)
| 주제 | 파일명 | 설명 |
|------|--------|------|
| TCP의 특징과 흐름제어 | [tcp.md](./notes/tcp.md) | 연결지향적 통신, 신뢰성 보장 |
| UDP의 특징과 사용처 | [udp.md](./notes/udp.md) | 비연결성 통신, 실시간 성능 |
| 3-way & 4-way 핸드셰이크 | [handshake.md](./notes/handshake.md) | 연결 및 종료 절차 |
| 포트 번호 | [ports.md](./notes/ports.md) | 서비스 식별 번호, well-known 포트 |

---

### 🌍 애플리케이션 계층
| 주제 | 파일명 | 설명 |
|------|--------|------|
| HTTP & HTTPS | [http.md](./notes/http.md) | 웹 통신 프로토콜 및 보안 |
| HTTP/1.1 vs HTTP/2 vs HTTP/3 | [http-versions.md](./notes/http-versions.md) | HTTP 프로토콜 진화 비교 |
| 쿠키, 세션, 캐시 | [cookie-session-cache.md](./notes/cookie-session-cache.md) | 상태 유지 및 성능 향상 전략 |
| DNS (Domain Name System) | [dns.md](./notes/dns.md) | 도메인 → IP 주소 매핑 시스템 |
| URL 구조 & 인코딩 | [url.md](./notes/url.md) | URI, URL, 인코딩 방식 |

---

### 🧪 실습 및 기타 개념
| 주제 | 파일명 | 설명 |
|------|--------|------|
| 소켓 프로그래밍 기초 | [socket.md](./notes/socket.md) | TCP/UDP 소켓 통신 구조 |
| 패킷 분석 도구 | [wireshark.md](./notes/wireshark.md) | 실시간 트래픽 분석 도구 |
| 네트워크 보안 개요 | [network-security.md](./notes/network-security.md) | MITM, 스니핑, 방화벽 등 위협과 대응 |
| CDN과 캐시 | [cdn.md](./notes/cdn.md) | 콘텐츠 분산 네트워크와 응답 속도 최적화 |

