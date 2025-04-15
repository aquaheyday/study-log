# 🛠️ Projects - 실습 및 사이드 프로젝트 모음

이 디렉토리는 학습 목적 또는 실무 실습용으로 진행한 프로젝트들을 정리한 공간입니다.  
각 프로젝트는 독립적인 기능 구현 또는 아키텍처 실험 등을 목표로 하며, 개별 디렉토리 내에 상세한 README와 코드가 포함되어 있습니다.

---

## 📂 디렉토리 구성

| 폴더명 | 설명 |
|---|---|
| [application](./application) | 어플리케이션 프로젝트 기록 |
| [docker](./docker) | 도커 프로젝트 기록 |
| [machine-learning](./machine-learning) | 머신러닝 프로젝트 기록 |

---

## 📋 App Project 목록

| 번호 | 주제 | 기술 스택 | 파일명 |
|---|---|---|---|
| 1 | 배달 주문 서비스의 백엔드 API 서버 구현 | Node.js, Express, MongoDB | [delivery-api-server](./application/delivery-api-server) |
| 2 | 커피 주문 시스템 백엔드 | Nginx, Laravel, MySQL | [menu-order-api](./application/menu-order-api) |
| 3 | 커피 주문 시스템 프론트엔드 | Nginx, Flutter | [menu-order-web-front](./application/menu-order-web-front) |
| 4 | 대기열 기반 티켓 발급 시스템 | Laravel | [queue-ticket](./application/queue-ticket) | 
| 5 | 관리자 CRUD RESTful API | Go | [restful-admin-crud](./application/restful-admin-crud) |
| 6 | 관리자 대시보드 프론트엔드 | Vite, TypeScript, SCSS | [vite-ts-scss-admin-dashboard](./application/vite-ts-scss-admin-dashboard) | 

---

## 📋 Docker Project 목록
| 번호 | 주제 | 기술 스택 | 파일명 |
|---|---|---|---|
| 1    | Go API 서버 구성            | Go, Dockerfile, Docker Compose      | `golang-api` |
| 2    | Laravel + MySQL 통합 환경   | Laravel, MySQL, PHP, Docker Compose | `laravel-mysql` |
| 3    | Next.js Green-Blue 배포     | Next.js, Nginx, Docker Compose      | `nextjs-green-blue-deploy` |
| 4    | Next.js 개발 컨테이너       | Next.js, Node.js, Docker Compose    | `nextjs` |
| 5    | React 개발 컨테이너         | React, Node.js, Docker Compose      | `react` |
| 6    | Redis 단독 실행 테스트 환경 | Redis, Docker Compose               | `redis` |

---

## 📋 Machine Learning Project 목록
| 번호 | 주제 | 기술 스택 | 파일명 |
|---|---|---|---|
| 1 | MNIST 손글씨 숫자 분류 CNN | TensorFlow, Keras, CNN, Data Augmentation, Matplotlib | [mnist_cnn_model](./machine-learning/mnist_cnn_model) |
