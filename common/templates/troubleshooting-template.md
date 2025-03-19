# 🐛 문제 제목 (예: NameError: 함수 호출 시 정의되지 않은 이름 에러 발생)

## 1️⃣ 문제 상황  
🔹 **발생 날짜:** YYYY-MM-DD  
🔹 **발생 환경:** (OS, 개발 언어, 프레임워크, 라이브러리 버전 등)  
🔹 **재현 방법:**  
1. (어떤 코드를 실행했는지)  
2. (어떤 조건에서 발생했는지)  
3. (특정 입력 값 또는 실행 흐름)  

📌 **오류 메시지 (있다면 포함)**  
🔹 **오류 로그 예제**
Traceback (most recent call last):
  File "main.py", line 10, in <module>
    print(hello_world())
NameError: name 'hello_world' is not defined

---

## 2️⃣ 원인 분석  
🔍 **원인:**  
- (문제가 발생한 원인을 설명)  
- (관련된 코드 또는 설정 문제 설명)  

📌 **관련 코드 스니펫 (있다면 포함)**
def hello():
    print("Hello, World!")

print(hello_world())  # 잘못된 함수명 호출

---

## 3️⃣ 해결 과정  
✅ **수정 방법:**  
1. (어떤 방법으로 해결했는지 설명)  
2. (수정된 코드 또는 설정 값)  

📌 **수정 후 코드**
def hello_world():
    print("Hello, World!")

print(hello_world())  # 올바른 함수명으로 수정

📌 **적용 후 정상 작동 여부:** ✅ 정상 동작 확인  
---

## 4️⃣ 배운 점 & 예방 방법  
📌 **핵심 정리:**  
- (이 오류를 통해 배운 점)  
- (유사한 오류를 방지하기 위한 방법)  

📌 **참고 자료:**  
- 🔗 [Stack Overflow - 관련 질문](https://stackoverflow.com/)  
- 📖 [공식 문서 링크](https://docs.python.org/3/)  

✅ **작성자:** (작성자명)  
✅ **업데이트 날짜:** YYYY-MM-DD  
