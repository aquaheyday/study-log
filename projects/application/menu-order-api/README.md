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
| Database | MySQL |
| 기타 | docker-compose로 컨테이너 관리 |

---

## 📲 기능 시연 이미지

### 1. [POST] 회원가입
![회원가입 POST](docs/assets/images/회원가입.png)

### 2. [POST] 이메일 찾기
![이메일 찾기 POST](docs/assets/images/이메일찾기.png)

### 3. [POST] 비밀번호 찾기
![비밀번호 찾기 POST](docs/assets/images/비밀번호찾기.png)

### 4. [POST] 로그인
![로그인 POST](docs/assets/images/로그인.png)

### 5. [GET] 방목록
![방목록 조회 GET](docs/assets/images/방목록조회.png)

### 6. [GET] 차트 정보
![방목록 차트 GET](docs/assets/images/방목록차트.png)

### 7. [GET] 방목록 탑 10
![방목록 탑 10 GET](docs/assets/images/방목록탑10.png)

### 8. [POST] 방생성
![방생성 POST](docs/assets/images/방생성.png)

### 9. [GET] 특정 방 조회
![특정 방 조회 GET](docs/assets/images/방조회.png)

### 10. [PUT] 방 상태 변경
![방 상태 변경 PUT](docs/assets/images/방상태변경.png)

### 11. [DELETE] 방 삭제
![방 삭제 DELETE](docs/assets/images/방삭제.png)

### 12. [GET] 메뉴 등록 목록
![메뉴 접수 목록 GET](docs/assets/images/메뉴접수목록.png)

### 13. [POST] 메뉴 등록
![메뉴 등록 POST](docs/assets/images/메뉴등록.png)

### 14. [PUT] 메뉴 수정
![메뉴 수정 PUT](docs/assets/images/메뉴수정.png)

### 15. [DELETE] 메뉴 삭제
![메뉴 삭제 DELETE](docs/assets/images/메뉴삭제.png)

### 16. [GET] 사용자 정보
![사용자 정보 GET](docs/assets/images/사용자정보조회.png)

### 17. [PUT] 사용자 정보 변경
![사용자 정보 변경 PUT](docs/assets/images/사용자정보변경.png)

---

## 📦 설치 및 실행 방법

### 1. 클론 및 환경 설정

```bash
git clone https://github.com/aquaheyday/study-log.git study-log/projects/application/menu-order-api
cd menu-order-api
cp .env.example .env
```
