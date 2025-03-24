# 🍽️ Menu Order Front

팀원들이 각자 메뉴를 접수하고, 랜덤으로 배달 담당을 선정하는 시스템입니다.  

---

## 📌 프로젝트 소개

### 주요 기능
| 기능 | 설명 |
|---|---|
| 로그인 | 개인 계정으로 로그인 후 메뉴 접수방 생성 |
| 메뉴 접수방 생성 | 팀원들이 함께 참여할 접수방 생성 |
| 메뉴 접수 | 각자 원하는 메뉴 선택 및 접수 |
| 접수 마감 | 접수 종료 후 배달 담당 랜덤 선정 |
| 주문 내역 확인 | 모든 팀원이 접수 결과 확인 가능 |

---

## 🛠️ 기술 스택

| 구분 | 사용 기술 |
|---|---|
| Frontend |  flutter |
| 기타 | docker-compose로 컨테이너 관리 |

---

## 📲 기능 시연 이미지

### 1. 로그인
- flutter에서 이메일, 비밀번호의 유효성검사를 합니다
- Api 서버에 로그인 요청을 하여 token을 발급받아 로그인 세션을 유지합니다.
<img src="https://github.com/user-attachments/assets/1d101960-90d6-4581-b913-91a34a080eb4"  width="600"/>

### 2. 회원가입
- fluuter에서 회원정보의 유효성검사를 합니다.
- Api 서버에서 계정 중복 확인 후 회원가입 처리 및 로그인 token을 발급합니다.
<img src="https://github.com/user-attachments/assets/12b4110f-775e-48df-837b-f669ede51c00"  width="600"/>

### 3. 비밀번호찾기
- 이메일을 이용하여 비밀번호를 재발급받습니다.
<img src="https://github.com/user-attachments/assets/f0aa437b-fea4-40f9-b634-31721f736586"  width="600"/>

### 4. 이용약관 & 개인정보처리방침
- showModalBottomSheet 에 작성된 내용을 노출합니다.
<img src="https://github.com/user-attachments/assets/7e035724-3e7e-4036-8a24-e10d1e743825"  width="600"/>

### 5. 방 생성
- 커피 구매처를 지정하여 제목, 비밀번호를 입력하여 생성합니다.
- 생성된 방의 url 정보를 공유하면 비밀번호 입력없이 바로 입장을 제공합니다.
<img src="https://github.com/user-attachments/assets/d8d3c192-f965-4d83-9733-4ebbbb5cc2e0"  width="600"/>

### 6. 방 목록 전환
- 탭 전환시 내용이 리빌드되며, 무한 스크롤의 현재 위치를 기억하고 있게됩니다.
<img src="https://github.com/user-attachments/assets/01c0c56b-8521-4ae9-8451-e6f6cfbe828d"  width="600"/>

### 7. 방 입장
- 방 생성자 또는 메뉴 주문자의 경우 바로 입장이 가능하며, 그 외 일 경우 비밀번호를 입력해야합니다.
<img src="https://github.com/user-attachments/assets/c43dc925-12f9-4ded-a537-6fdf57c8b1ed"  width="600"/>

### 8. 메뉴 주문
- 메인과 서브로 주문이 가능하며, 메인메뉴가 품절일때, 서브메뉴를 확인하여 주문할 수 있습니다. 
<img src="https://github.com/user-attachments/assets/83e13794-ae67-4c3b-9d8b-23355d2cbe76"  width="600"/>

### 9. 주문마감 & 주문오픈
- 주문 마감시 주문자 리스트중 랜덤으로 배달원이 배정됩니다.
- 마감된 주문방에서 재오픈시 추가 주문을 받을 수 있습니다.
<img src="https://github.com/user-attachments/assets/86d61a57-8102-460b-8e25-85340aece517"  width="600"/>

### 10. 내정보 수정
- 회원의 시퀀스 번호를 기반으로 데이터를 저장하고 있어, 이메일을 포함한 모든 정보를 자유롭게 변경가능합니다.
<img src="https://github.com/user-attachments/assets/3ad7d10e-4fa6-4a8d-bc44-97f4c27f4930"  width="600"/>
---

## 📦 설치 및 실행 방법

### 1. 클론 및 환경 설정

```bash
git clone https://github.com/aquaheyday/language-archive.git language-archive/flutter/projects/menu-order-front/
cd menu-order-front
```
