# 📜 오픈소스 기여 가이드 (Open Source Contribution Guide)

오픈소스 프로젝트에 기여할 때 **효율적인 협업을 위한 모범 사례(Best Practices)** 를 정리한 가이드입니다.  

---

## 1. 커밋 메시지 규칙 (Conventional Commits)
| 태그 | 의미 |
|------|------|
| `feat:` | 새로운 기능 추가 |
| `fix:` | 버그 수정 |
| `docs:` | 문서 수정 |
| `refactor:` | 코드 리팩토링 |
| `test:` | 테스트 코드 추가 |

---

## 2. 오픈소스 기여 프로세스
오픈소스 프로젝트에 기여하려면 **다음 순서를 따라 주세요.**  

**1. 저장소 Fork & Clone**

```bash
# 1. 프로젝트 Fork
git clone https://github.com/{your-username}/{repo-name}.git
cd {repo-name}

# 2. 원본 저장소(Remote Upstream) 연결
git remote add upstream https://github.com/{original-owner}/{repo-name}.git
git fetch upstream
```

---

**2. 작업을 위한 브랜치 생성**

```bash
git checkout -b feature/new-feature
```

✔ 브랜치명은 `feature/기능-이름`, `fix/버그-이름` 형식으로 생성

---

**3. 코드 변경 및 커밋 (커밋 메시지 규칙 준수)**
```bash
git add .
git commit -m "feat: 신규 기능 추가"
git push origin feature/new-feature
```

---

**4. GitHub 원본 저장소에 Pull Request(PR) 생성**  

- PR 제목 템플릿  
```md
# ✨ feat: 특정 사용자 조회 추가 (#123)
```

- PR 작성 템플릿  
```md
## ✨ 작업 내용
- 특정 사용자 조회 기능 추가
- 특정 사용자 조회 GET API 추가 (`/api/user/:id`)

## ✅ 테스트 내용
- `GET /api/user/:id` 요청 시 정상 조회 확인
- 유효하지 않은 id로 요청 시 404 반환 확인

## 🔗 관련 이슈
Closes #123
```

- PR 작성 체크리스트  
  ✔ PR 제목이 명확한가? (`feat: 특정 사용자 조회 추가`)  
  ✔ 관련 이슈 번호를 포함했는가? (`Closes #123`)  
  ✔ 변경된 코드에 대한 설명이 충분한가?  
  ✔ 테스트 내용이 포함되었는가?  

---

## 3. 코드 스타일 가이드 (Code Style Guide)
코드 스타일을 통일하면 코드 리뷰가 쉬워지고 유지보수가 편리합니다.  

### Python 스타일 가이드 (PEP 8)
```python
def fetch_user_data(user_id: int) -> dict:
    """사용자 정보를 가져오는 함수"""
    if user_id < 0:
        raise ValueError("User ID must be positive")
    return {"id": user_id, "name": "John Doe"}
```

📌 **PEP 8 규칙 요약**
- 4칸 들여쓰기(탭 대신 공백)
- 함수 및 변수명: `snake_case`
- 클래스명: `PascalCase`
- 한 줄 최대 길이: **79자**

---

### JavaScript 스타일 가이드 (ESLint)
```javascript
function fetchUserData(userId) {
  if (userId < 0) {
    throw new Error("User ID must be positive");
  }
  return { id: userId, name: "John Doe" };
}
```

**JavaScript 코드 스타일**
- `const`, `let` 사용 (`var` 지양)
- `camelCase` 변수명 사용
- `===` 연산자 사용 (느슨한 비교 `==` 지양)
- 함수 내 `return` 문 명확히 작성

---

## 4. 코드 리뷰 모범 사례 (Code Review Best Practices)
효율적인 코드 리뷰는 프로젝트 품질 향상과 팀워크 강화를 위한 필수 요소입니다.  

### 📝 리뷰할 때 체크할 사항
✔ 코드가 일관된 스타일을 따르는가?  
✔ 코드가 재사용 가능하도록 작성되었는가?  
✔ 성능을 고려하여 최적화가 되었는가?  
✔ 필요한 경우 적절한 주석이 추가되었는가?  
✔ 불필요한 코드(Dead Code)는 없는가?  

**리뷰 예시 (Good 👍)**
```md
✅ `fetchUserData` 함수에서 유효성 검사를 추가한 점이 좋습니다!  
💡 다만, `try-catch`를 추가하여 예외 처리를 명확하게 하면 더 좋을 것 같습니다.  
```

**리뷰 예시 (Bad 👎)**
```md
❌ "이 코드 왜 이렇게 짰나요?" (비판적인 어조 X)  
✅ "이 부분을 `fetch` 대신 `axios`로 변경하면 더 안정적일 것 같아요!" (건설적인 피드백 O)
```

---

## 🎯 정리
✔ 오픈소스 프로젝트에 기여할 때 **기본적인 Git 사용법을 숙지**하자  
✔ PR을 보낼 때는 **커밋 메시지와 PR 설명을 명확하게 작성**하자  
✔ 코드 리뷰는 **건설적인 피드백을 제공하는 방식**으로 진행하자  
