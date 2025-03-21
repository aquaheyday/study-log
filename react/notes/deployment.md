# 🚀 React 애플리케이션 배포

React 애플리케이션을 배포하는 방법을 단계별로 정리했습니다.  
**Vercel, Netlify, GitHub Pages, Firebase, AWS S3 등의 배포 플랫폼을 활용하는 방법**을 다룹니다.

---

## 1️⃣ React 애플리케이션 배포란?

- 배포(Deployment): 로컬 환경에서 개발한 애플리케이션을 웹에서 접근 가능하도록 서버에 업로드하는 과정
- 정적 사이트로 배포 가능 - React 앱은 빌드 후 HTML, CSS, JS 파일로 변환
- CDN(Content Delivery Network) 사용 가능 - 성능 최적화 및 글로벌 접근성 향상
- 자동화 배포 지원 - GitHub Actions, CI/CD 도구 활용 가능

---

## 2️⃣ React 프로젝트 빌드하기

React 애플리케이션을 배포하려면 먼저 **빌드(Build)** 해야 합니다.  

#### 1. React 빌드 명령어
```sh
npm run build
```
또는
```sh
yarn build
```

#### 2. `build` 폴더 생성
빌드 후 `build/` 디렉터리가 생성되며, 여기에는 **HTML, CSS, JS 등의 정적 파일이 포함**됩니다.

---

## 3️⃣ Vercel을 이용한 배포 (가장 쉬운 방법)

- Vercel은 React와 Next.js에 최적화된 무료 배포 서비스
- GitHub과 연동하여 자동 배포 가능

### 1) Vercel에 배포하는 방법
#### 1. [Vercel 공식 웹사이트](https://vercel.com/)에 가입
#### 2. `vercel` CLI 설치  
```sh
npm install -g vercel
```

#### 3. 프로젝트 디렉터리에서 실행  
```sh
vercel
```

#### 4. 도메인(URL) 자동 생성 및 배포 완료!

---

## 4️⃣ Netlify를 이용한 배포
- 무료 정적 사이트 호스팅 서비스
- GitHub와 연동하여 자동 배포 가능
- Drag & Drop 배포 지원

### 1) Netlify 배포 방법
#### 1. [Netlify 공식 웹사이트](https://www.netlify.com/) 가입  
#### 2. **GitHub 저장소 연결** 또는 **Drag & Drop**  
#### 3. `netlify-cli` 사용 (선택 사항)  
```sh
npm install -g netlify-cli
```
#### 4. 프로젝트에서 실행  
```sh
netlify deploy
```

---

## 5️⃣ GitHub Pages를 이용한 배포
- GitHub에서 제공하는 무료 정적 웹사이트 호스팅 서비스
- 개인 프로젝트, 포트폴리오 배포에 적합

### 1) GitHub Pages 배포 방법
#### 1. `gh-pages` 패키지 설치  
```sh
npm install --save-dev gh-pages
```
#### 2. `package.json` 수정
```json
{
  "homepage": "https://your-username.github.io/your-repo",
  "scripts": {
    "predeploy": "npm run build",
    "deploy": "gh-pages -d build"
  }
}
```
#### 3. 배포 실행  
```sh
npm run deploy
```
#### 4. GitHub Pages에서 배포된 사이트 확인

---

## 6️⃣ Firebase를 이용한 배포
- Google의 클라우드 플랫폼으로 빠르고 안정적인 배포 가능
- 무료 SSL, 서버리스 기능 제공

### 1) Firebase 배포 방법
#### 1. Firebase CLI 설치  
```sh
npm install -g firebase-tools
```
#### 2. Firebase 로그인  
```sh
firebase login
```
#### 3. Firebase 프로젝트 초기화  
```sh
firebase init
```
#### 4. 배포 실행  
```sh
firebase deploy
```
#### 5. Firebase Hosting에서 배포된 URL 확인

---

## 7️⃣ AWS S3 & CloudFront를 이용한 배포
- AWS S3(파일 저장) + CloudFront(CDN)로 성능 최적화된 배포 가능
- 기업용 배포에 적합

### 1) AWS S3 배포 방법
#### 1. `build/` 폴더를 S3에 업로드  
```sh
aws s3 sync build/ s3://your-bucket-name --acl public-read
```
#### 2. **CloudFront 배포 설정하여 CDN 적용**  

---

## 8️⃣ CI/CD를 활용한 자동 배포
- GitHub Actions, CircleCI, Travis CI 등을 사용하여 자동 배포 가능
- 코드 변경 시 자동으로 배포하여 개발 효율성 향상

### 1) GitHub Actions 예제 (`.github/workflows/deploy.yml`)
```yaml
name: Deploy React App

on:
  push:
    branches:
      - main

jobs:
  deploy:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout Repository
        uses: actions/checkout@v2

      - name: Install Dependencies
        run: npm install

      - name: Build Project
        run: npm run build

      - name: Deploy to GitHub Pages
        run: npm run deploy
```

---

## 9️⃣ 배포 시 고려할 사항
- 환경 변수 설정 (`.env` 파일 관리)
- SEO 최적화 (SSR 사용 고려)
- HTTPS 적용 (SSL 인증서 사용)
- 오류 추적 및 로깅 (Sentry, LogRocket 등 활용)
- 캐시 관리 (Cloudflare, CDN 설정 등 활용)

---

## 🎯 정리
✔ React 애플리케이션을 배포하려면 먼저 npm run build를 실행하여 정적 파일을 생성  
✔ Vercel, Netlify → 간편한 무료 배포, GitHub 연동 가능  
✔ GitHub Pages → 개인 프로젝트 및 포트폴리오 배포에 적합  
✔ Firebase Hosting → 빠르고 안정적인 Google 클라우드 기반 배포  
✔ AWS S3 + CloudFront → 기업용 대규모 서비스 배포에 적합  
✔ CI/CD 자동화 → GitHub Actions 등을 활용하여 코드 변경 시 자동 배포 가능  
