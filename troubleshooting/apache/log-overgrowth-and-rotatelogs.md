# 🐛 Apache 로그 파일 용량 폭주 문제

## ⚠️ 문제 상황
- 발생 날짜: 2025-04-25
- 발생 환경:
  - OS: Rocky Linux
  - Web Server: Apache
  - 기존 로그 방식: `CustomLog /경로/파일 combined` (직접 파일에 기록)
- 재현 방법 
1. 사이트에 HTTP(S) 요청이 계속 들어오면
2. `/access_log`, `/error_log` 등의 로그 파일이 지속적으로 증가  
3. `cat /dev/null`로 비워도 다시 로그가 즉시 쌓임  
4. `logrotate`로 회전해도 파일 핸들이 유지되어 용량이 줄지 않음

---

## 🔍 원인 분석
- Apache는 로그 파일을 "이름"이 아니라 "파일 핸들(file descriptor)"로 접근함  
- 기존 로그 설정(`CustomLog /경로/파일`)은 logrotate로 회전돼도 Apache가 이전 핸들을 계속 사용함  
- `copytruncate`를 사용해도 간혹 동작하지 않거나 reload가 제대로 반영되지 않음  

#### 기존 `ssl.conf` 에 설정된 내용
```apache
CustomLog /경로/파일 combined
ErrorLog  /경로/파일
```

---

## 🛠 해결 방법
- Apache 설정을 rotatelogs 방식으로 전환하여 로그를 날짜 단위로 자동 회전하도록 설정
- 로그 핸들이 자동 갱신되므로 logrotate나 cat /dev/null 불필요
- 무중단으로 반영 가능 (systemctl reload httpd)

#### 변경된 `ssl.conf`
```yml
CustomLog "|/usr/sbin/rotatelogs /경로/파일.%Y-%m-%d 86400" combined
ErrorLog  "|/usr/sbin/rotatelogs /경로/파일.%Y-%m-%d 86400"
```

---

## 🚀 결과
✅ 로그가 자동으로 날짜별 파일로 회전됨 `파일.2025-04-25`  
✅ 파일 핸들 문제가 해결되어 로그 누적 폭주 방지  
