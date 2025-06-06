# ⚡ inode & 디렉토리 구조

**inode와 디렉토리는 파일 시스템에서 파일 정보를 관리하고 계층 구조를 구성하는 핵심 요소입니다.**

---

## 1️⃣ inode란?

| 항목       | 설명 |
|------------|------|
| **정의**       | 파일의 메타데이터(속성, 위치 정보 등)를 저장하는 구조체 |
| **역할**       | 파일의 실제 내용 위치와 속성 정보 관리 |
| **식별 방식** | 각 파일은 고유한 inode 번호를 가짐 |
| **위치**       | 파일 시스템 내부, 일반적으로 고정된 영역에 존재 |
| **특징**       | 파일 이름은 저장하지 않음 (디렉토리에서 관리) |

---

## 2️⃣ inode 구조 예시

| 필드          | 설명 |
|---------------|------|
| **inode 번호**   | 고유 식별자 |
| **파일 유형**     | 일반 파일, 디렉토리, 링크 등 |
| **파일 크기**     | 바이트 단위 크기 |
| **접근 권한**     | rwx 정보 |
| **소유자 정보**   | UID, GID |
| **시간 정보**     | 생성, 수정, 접근 시간 |
| **데이터 블록 포인터** | 실제 데이터가 저장된 블록 위치 |

---

## 3️⃣ 디렉토리란?

| 항목       | 설명 |
|------------|------|
| **정의**       | 파일(또는 다른 디렉토리)을 포함하는 특수한 파일 |
| **역할**       | 파일 이름과 해당 inode 번호를 매핑 |
| **구조**       | `파일 이름 → inode 번호` 쌍으로 구성된 테이블 |
| **형태**       | 트리 구조로 계층화 가능 |
| **중요 디렉토리** | `/`, `/home`, `/etc`, `/bin` 등 (UNIX 계열 기준) |

---

## 4️⃣ inode와 디렉토리 관계

```
[ 디렉토리 ]
   └─ 파일 이름 → inode 번호 매핑
              ↓
        [ inode ]
           └─ 파일 속성 + 데이터 위치
                 ↓
             [ 데이터 블록 ]
```

- 디렉토리는 단순히 **이름과 inode를 연결**  
- 실제 데이터 정보는 **inode가 가지고 있음**  
- 따라서 **같은 inode를 참조하면 하드링크**가 됨

---

## 5️⃣ 관련 개념: 하드링크 vs 소프트링크

| 항목            | 하드링크 (Hard Link) | 소프트링크 (Symbolic Link) |
|------------------|----------------------|-----------------------------|
| **구조**           | 동일 inode 공유        | 별도 파일로 경로만 가리킴        |
| **파일 삭제 시**     | 링크 수 0이 되면 삭제     | 원본 삭제 시 링크 깨짐           |
| **범위**           | 같은 파일 시스템 내만 가능 | 다른 파일 시스템도 가능          |
| **명령어**         | `ln 원본 대상`           | `ln -s 원본 대상`               |

---

## 6️⃣ inode 번호 확인 예시

```bash
# 파일의 inode 번호 확인
ls -i sample.txt
# 출력 예시: 123456 sample.txt

# inode 기준으로 파일 찾기
find . -inum 123456
```

---

## 7️⃣ 디렉토리 구조 예시 (UNIX/Linux)

```
/
├── bin       # 기본 명령어 바이너리
├── boot      # 부팅 관련 파일
├── dev       # 장치 파일
├── etc       # 설정 파일
├── home      # 사용자 디렉토리
│   ├── user1
│   └── user2
├── tmp       # 임시 파일 저장소
└── var       # 로그 등 가변 데이터
```

---

## 🎯 정리 요약

✔ **inode는 파일의 정보를 저장하는 구조체**, 실제 데이터 블록 위치를 관리  
✔ **디렉토리는 이름과 inode를 연결하는 매핑 테이블** 역할  
✔ inode를 기반으로 **하드링크/소프트링크** 구현 가능  
✔ 파일 이름이 아닌 inode 번호가 실제 식별자 역할을 함  
✔ 디렉토리 구조는 트리 형태로 **파일 시스템을 계층화**

