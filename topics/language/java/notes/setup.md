# 🛠️ Java 개발 환경 설정

Java 개발을 시작하기 위한 환경 구성 방법을 정리합니다.  
JDK 설치부터 CLI 실행, IDE 선택, 기본 프로젝트 구성까지 다룹니다.

---

## 1️⃣ JDK란?

- **Java Development Kit**의 약자
- 자바를 **개발하고 실행하기 위한 도구 모음**
- 컴파일러(`javac`), 실행기(`java`), 문서 생성기(`javadoc`), 패키징 도구(`jar`) 포함

## 2️⃣ JDK 설치
### 1) 설치 경로
- Oracle JDK: https://www.oracle.com/java/technologies/javase-downloads.html  
- OpenJDK: https://jdk.java.net/  
- SDKMAN (macOS/Linux): https://sdkman.io/

---

### 2) 설치 확인

설치가 완료되면 터미널(또는 CMD)에서 아래 명령어로 확인합니다:

```sh
java -version
javac -version
```

#### 출력 예시

```sh
java version "17.0.8" 2023-07-18 LTS
javac 17.0.8
```

---

### 3) 환경 변수 설정 (Windows 기준)

JDK 설치 후 환경 변수 등록이 필요할 수 있습니다.

#### 1. `JAVA_HOME` 설정  
   → 예: `C:\Program Files\Java\jdk-17`

#### 2. `Path`에 `%JAVA_HOME%\bin` 추가

```sh
# 환경 변수 적용 확인
echo %JAVA_HOME%
java -version
```

---

## 3️⃣ 첫 번째 Java 파일 실행

```java
// HelloWorld.java
public class HelloWorld {
    public static void main(String[] args) {
        System.out.println("Hello, Java!");
    }
}
```

#### 컴파일 및 실행

```sh
javac HelloWorld.java
java HelloWorld
```

#### 결과

```
Hello, Java!
```

---

## 4️⃣ IDE 추천

| IDE | 특징 |
|-----|------|
| IntelliJ IDEA | JetBrains에서 만든 강력한 Java 전용 IDE (추천) |
| Eclipse | 전통적인 오픈소스 IDE |
| VS Code | 경량화된 에디터 + 확장팩으로 Java 개발 가능 |

---

## 5️⃣ 기본 프로젝트 구조 예시

```plaintext
MyJavaProject/
├── src/
│   └── Main.java
├── out/           ← 컴파일된 클래스 파일 (IDE 생성)
├── lib/           ← 외부 라이브러리 위치
└── README.md
```

---

## 📌 팁

- 최신 프로젝트는 Java 17 또는 Java 21 (LTS) 권장
- Gradle 또는 Maven을 통한 빌드 자동화 환경도 익혀두면 좋음
- VS Code 사용 시 `Extension Pack for Java` 설치 필요
