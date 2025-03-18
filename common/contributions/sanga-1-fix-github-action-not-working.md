# 🐛 Issue 해결: GitHub Actions Not Working (#1)
📌 **이슈 링크:** [#1](https://github.com/daewoungkim/sanga/issues/1)

## 🔍 문제 상황
- git push 시 서버에서 git pull을 위한 github action 이 동작하지 않고 있는 상황

## 🛠 해결 방법
1. $HOME/.ssh 디렉토리 생성 후 700 권한 설정
2. $HOME/.ssh/known_hosts 파일 생성 후 644 권한 설정
3. 웹 서버 와 git repo 에 ssh key 설정
4. git pull을 위한 웹 서버 접속 ip, port, id, 프로젝트 경로 설정

## 🚀 결과
✅ 웹 서버에서 git pull origni main 을 위한 github action 이 정상 작동 (테스트 통과)  
