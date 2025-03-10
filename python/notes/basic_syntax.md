# Python 기본 문법 정리

## 1. 변수와 자료형
```python
# 변수 선언 및 할당
x = 10
y = 3.14
name = "Python"
is_valid = True

# 자료형 확인
print(type(x))        # <class 'int'>
print(type(y))        # <class 'float'>
print(type(name))     # <class 'str'>
print(type(is_valid)) # <class 'bool'>
```

---

## 2. 연산자

```python
# 산술 연산자
a = 10
b = 3
print(a + b)   # 13
print(a - b)   # 7
print(a * b)   # 30
print(a / b)   # 3.3333...
print(a // b)  # 3 (몫)
print(a % b)   # 1 (나머지)
print(a ** b)  # 1000 (제곱)

# 비교 연산자
print(a > b)   # True
print(a < b)   # False
print(a == b)  # False
print(a != b)  # True

# 논리 연산자
print(True and False) # False
print(True or False)  # True
print(not True)       # False
```

---

## 3. 조건문

```python
x = 10
if x > 5:
    print("x는 5보다 큽니다.")
elif x == 5:
    print("x는 5입니다.")
else:
    print("x는 5보다 작습니다.")
```

---

## 4. 반복문

```python
# for문
for i in range(5):  # 0부터 4까지 반복
    print(i)

# while문
count = 0
while count < 5:
    print(count)
    count += 1
```

---

## 5. 리스트 (List)

```python
numbers = [1, 2, 3, 4, 5]
print(numbers[0])    # 1 (인덱싱)
print(numbers[-1])   # 5 (음수 인덱싱)
print(numbers[1:4])  # [2, 3, 4] (슬라이싱)

# 리스트 메서드
numbers.append(6)    # 요소 추가
numbers.remove(3)    # 요소 삭제
numbers.sort()       # 정렬
numbers.reverse()    # 역순 정렬
print(numbers)
```

---

## 6. 튜플 (Tuple)

```python
t = (1, 2, 3, 4, 5)
print(t[0])    # 1
print(t[1:3])  # (2, 3)

# 튜플은 변경 불가능 (immutable)
# t[0] = 10  # 오류 발생!
```

---

## 7. 딕셔너리 (Dictionary)

```python
person = {"name": "Alice", "age": 25, "city": "Seoul"}
print(person["name"])  # Alice

# 딕셔너리 메서드
person["job"] = "Developer"  # 추가
del person["age"]            # 삭제
print(person.keys())         # 키 목록
print(person.values())       # 값 목록
print(person.items())        # (키, 값) 쌍 목록
```

---

## 8. 함수 (Function)

```python
def greet(name):
    return f"Hello, {name}!"

print(greet("Python"))  # Hello, Python!
```

---

## 9. 클래스 (Class)

```python
class Person:
    def __init__(self, name, age):
        self.name = name
        self.age = age
    
    def introduce(self):
        return f"My name is {self.name} and I am {self.age} years old."

p1 = Person("Alice", 25)
print(p1.introduce())  # My name is Alice and I am 25 years old.
```

---

## 10. 파일 입출

```python
# 파일 쓰기
with open("test.txt", "w") as f:
    f.write("Hello, Python!")

# 파일 읽기
with open("test.txt", "r") as f:
    content = f.read()
    print(content)
```

---

## 11. 예외 처리

```python
try:
    x = 1 / 0
except ZeroDivisionError as e:
    print("0으로 나눌 수 없습니다!", e)
finally:
    print("예외 처리 완료")
```

---

## 12. 리스트 내포 (List Comprehension)

```python
squares = [x**2 for x in range(5)]
print(squares)  # [0, 1, 4, 9, 16]
```

---

## 13. 람다 함수 (Lambda Function)

```python
square = lambda x: x ** 2
print(square(5))  # 25
```

---

## 14. 모듈과 패키지

```python
# 모듈 불러오기
import math
print(math.sqrt(16))  # 4.0
```
