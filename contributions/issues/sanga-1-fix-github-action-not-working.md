# 🐛 GitHub Actions Not Working (#1)
>📌 이슈 링크: [GitHub Actions Not Working (#1)](https://github.com/daewoungkim/sanga/issues/1) (🔒 Private Repository)

---

## ⚠️ 문제 상황
1. git push 시 서버에서 git pull을 위한 github action 이 동작하지 않고 있는 상황

---

## 🔍 원인 분석
1. 생성된 디렉토리 권한
2. 서버 접속 정보 오류

#### `.github\workflows\autopull.yml`
```yml
name: Auto Pull

on:
  push:
    branches:
      - main

jobs:
  pull:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Setup SSH
        run: |
          mkdir -p $HOME/.ssh
          echo "${{ secrets.SSH_PRIVATE_KEY }}" > $HOME/.ssh/id_rsa
          chmod 600 $HOME/.ssh/id_rsa
          ssh-keyscan -H *.***.**.*** >> $HOME/.ssh/known_hosts

      - name: Git Pull
        run: |
          ssh -T git@*.***.**.*** || true
          git pull origin main
```


---

## 🛠 해결 방법
1. $HOME/.ssh 디렉토리 생성 후 700 권한 설정
2. $HOME/.ssh/known_hosts 파일 생성 후 644 권한 설정
3. 웹 서버 와 git repo 에 ssh key 설정
4. git pull을 위한 웹 서버 접속 ip, port, id, 프로젝트 경로 설정
5. ip, port, id 보안을 위해 secret 설정

#### `.github\workflows\autopull.yml`
```yml
name: Auto Pull

on:
  push:
    branches:
      - main

jobs:
  pull:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Setup SSH
        run: |
          mkdir -p $HOME/.ssh
          chmod 700 $HOME/.ssh

          if [ ! -f "$HOME/.ssh/id_rsa" ]; then
            echo -e "${{ secrets.SSH_PRIVATE_KEY }}" > "$HOME/.ssh/id_rsa"
            chmod 600 $HOME/.ssh/id_rsa
            echo "SSH key 생성"
          else
            echo "ssh key 이미 있음."
          fi
          
          touch $HOME/.ssh/known_hosts
          ssh-keyscan -p ${{ secrets.AUTOPULL_PORT }} -H ${{ secrets.AUTOPULL_IP }} >> $HOME/.ssh/known_hosts
          chmod 644 "$HOME/.ssh/known_hosts"
      - name: Test SSH Connection
        run: ssh -p ${{ secrets.AUTOPULL_PORT }} -T -o StrictHostKeyChecking=no -o UserKnownHostsFile=$HOME/.ssh/known_hosts ${{ secrets.AUTOPULL_ID }}@${{ secrets.AUTOPULL_IP }} || true

      - name: Git Pull
        run: |
          ssh -p ${{ secrets.AUTOPULL_PORT }} -T -o BatchMode=yes ${{ secrets.AUTOPULL_ID }}@${{ secrets.AUTOPULL_IP }} << 'EOF'
            cd /home/users/${{ secrets.AUTOPULL_ID }}/code/eoc/src/renew_sanga || exit 1
            git pull origin main
          EOF
```

---

## 🚀 결과
✅ 웹 서버에서 git pull origni main 을 위한 github action 이 정상 작동 (테스트 통과) 
