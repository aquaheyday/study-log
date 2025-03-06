# 🍽️ Menu Order API

팀원들이 각자 메뉴를 접수하고, 랜덤으로 배달 담당을 선정하는 시스템입니다.  

---

## 📌 프로젝트 소개

### 주요 기능
| 기능 | 설명 |
|---|---|
| 로그인 | 팀별 계정으로 로그인 후 메뉴 접수방 생성 |
| 메뉴 접수방 생성 | 팀원들이 함께 참여할 접수방 생성 |
| 메뉴 접수 | 각자 원하는 메뉴 선택 및 접수 |
| 접수 마감 | 접수 종료 후 배달 담당 랜덤 선정 |
| 주문 내역 확인 | 모든 팀원이 접수 결과 확인 가능 |

---

## 🛠️ 기술 스택

| 구분 | 사용 기술 |
|---|---|
| Backend | Laravel 10 |
| Frontend | Flutter (iOS, Android, Web) |
| Database | MySQL |
| Infrastructure | Docker + AWS EC2 (Route 53 도메인 연결 및 SSL 인증) |
| 기타 | docker-compose로 컨테이너 관리 |

---

## 📲 기능 시연 이미지

|화면|설명|
|---|---|
| 로그인 | 팀 계정 로그인 화면 |
| 메뉴 접수방 생성 | 접수방 생성 및 팀원 초대 |
| 메뉴 접수 화면 | 팀원들이 각자 메뉴 접수 |
| 배달원 선정 화면 | 접수 종료 후 랜덤으로 배달 담당 선정 |

### 1. [POST] 회원가입
![회원가입 POST](docs/assets/images/회원가입.png)

### 2. [POST] 이메일 찾기
![이메일 찾기 POST](docs/assets/images/이메일찾기.png)

### 3. [POST] 비밀번호 찾기
![비밀번호 찾기 POST](docs/assets/images/비밀번호찾기.png)

### 4. [POST] 로그인
![로그인 POST](docs/assets/images/로그인.png)

### 5. [GET] 방목록
![방목록 GET](docs/assets/images/방목록.png)

---

## 📦 설치 및 실행 방법

### 1. 클론 및 환경 설정

```bash
git clone https://github.com/your-repo/team-menu-order.git
cd team-menu-order
cp .env.example .env
```
