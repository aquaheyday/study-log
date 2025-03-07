# 🚀 FastAPI Delivery Service

이 프로젝트는 **FastAPI + pymysql** 기반으로 구성된 간단한 택배사/요금제/문의 관리 API 서비스입니다.  
Docker 기반 실행 환경을 제공하며, FastAPI의 기본 구성과 MariaDB 연동을 학습하는 데 목적을 둡니다.

---

## 📦 기술 스택

| 기술 | 버전 | 설명 |
|---|---|---|
| Python | 3.10 | 런타임 환경 |
| FastAPI | >=0.68.0, <0.69.0 | 웹 프레임워크 |
| Pydantic | >=1.8.0, <2.0.0 | 데이터 검증 및 모델링 |
| Uvicorn |>=0.15.0, <0.16.0 | ASGI 서버 |
| pymysql | 최신 버전 (1.x) | MySQL 연동 라이브러리 |
| Docker | latest | 컨테이너 실행 환경 |

---

## 📂 주요 파일 설명

| 파일명 | 설명 |
|---|---|
| Dockerfile | 컨테이너 환경 정의 (Python 3.10 + 의존성 설치) |
| requirements.txt | FastAPI, Pydantic, Uvicorn, pymysql 버전 고정 |
| .gitignore | 캐시, 환경파일 등 버전 관리 제외 항목 정의 |
| main.py | FastAPI 엔드포인트 정의 및 서비스 로직 구현 |

---

# 📋 주요 기능

| 기능 | 엔드포인트 | 메서드 |
|---|---|---|
| 택배사 목록 조회 | `/api/companyList` | GET |
| 요금제 조회 | `/api/ratePlan` | GET |
| 문의 목록 조회 | `/api/description` | GET |
| 문의 등록 | `/api/description` | POST |

---

## 🔧 환경 구성 및 실행 방법

### 1. Docker 이미지 빌드
```bash
docker build -t fastapi-delivery-service .
```

### 2. 도커 컨테이너 실행
```bash
docker run -p 8000:8000 fastapi-delivery-service uvicorn main:app --host 0.0.0.0 --port 8000
```

---

# 📚 참고 자료
- [FastAPI 공식 문서](https://fastapi.tiangolo.com/)
- [pymysql 공식 문서](https://pymysql.readthedocs.io/en/latest/)
- [Docker 공식 문서](https://docs.docker.com/)



