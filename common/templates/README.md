# 📝 Templates

이 폴더는 `language-archive`에서 각 언어별 학습 자료를 작성할 때 사용하는 **기본 템플릿**을 모아둔 공간입니다.  
새로운 노트, 예제, 프로젝트, README 등을 만들 때, 이곳에 있는 템플릿을 복사해서 사용하면 됩니다.

---

## 📂 템플릿 구성 목록

| 파일명 | 설명 | 용도 |
|---|---|---|
| `note-template.md` | 개념 정리 노트 템플릿 | `notes/` 작성용 |
| `example-template.md` | 예제 코드 설명 템플릿 | `examples/` 작성용 |
| `project-readme-template.md` | 프로젝트 설명 템플릿 | `projects/` 하위 프로젝트용 |
| `troubleshooting-template.md` | 트러블슈팅 기록 템플릿 | `troubleshooting/` 작성용 |
| `component-template.jsx` | React 컴포넌트 기본 템플릿 | React 예제/프로젝트 작성용 |
| `script-template.py` | Python 스크립트 기본 템플릿 | Python 예제 작성용 |

---

## 📑 사용 방법
1. 새로운 파일 작성 시, 해당 폴더에서 적합한 템플릿 복사
2. 파일명 변경 후 내용 작성
3. 필요 시 주석으로 작성 날짜, 작성자 정보 추가

### 예시
```bash
# 새 노트 작성 예시 (Python)
cp templates/note-template.md python/notes/new-topic.md

# 새 예제 작성 예시 (React 컴포넌트)
cp templates/component-template.jsx react/examples/MyComponent.jsx

