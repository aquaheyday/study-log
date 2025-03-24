# 🔲 Python 클래스(Class)와 객체(Object)

Python에서 **클래스(Class)** 는 객체를 생성하는 틀이며,  
**객체(Object)** 는 클래스에서 생성된 인스턴스를 의미합니다.

---

## 1️⃣ 클래스와 객체 개념

### 1) 클래스란?
- 클래스는 **속성(변수)과 동작(메서드)** 을 정의하는 **설계도(틀)** 입니다.
- 클래스를 사용하면 코드를 **재사용**하고 **구조화**할 수 있습니다.

---

### 2) 객체란?
- 객체는 **클래스에서 생성된 인스턴스(Instance)** 입니다.
- 동일한 클래스를 기반으로 여러 개의 객체를 생성할 수 있습니다.

---

## 2️⃣ 클래스 정의 및 객체 생성

### 기본적인 클래스 정의
```python
class Dog:
    def __init__(self, name, breed):  # 생성자
        self.name = name  # 인스턴스 변수
        self.breed = breed  # 인스턴스 변수

    def bark(self):  # 메서드
        print(f"{self.name}가 멍멍! 하고 짖습니다.")

# 객체 생성
dog1 = Dog("Buddy", "Golden Retriever")
dog2 = Dog("Charlie", "Beagle")

# 메서드 호출
dog1.bark()  # "Buddy가 멍멍! 하고 짖습니다."
dog2.bark()  # "Charlie가 멍멍! 하고 짖습니다."
```

✔ `__init__` → 생성자(Constructor)로 객체가 생성될 때 실행됨  
✔ `self` → 객체 자기 자신을 가리키는 키워드  
✔ 클래스 내부의 변수(속성)는 `self.변수명` 형태로 정의  

---

## 3️⃣ 인스턴스 변수와 클래스 변수

### 1) 클래스 변수 (Class Variable)
- 모든 객체가 공유하는 변수 (`클래스명.변수명` 형태로 정의)

---

### 2) 인스턴스 변수 (Instance Variable)
- `self.변수명` 형태로 객체마다 다른 값을 저장하는 변수

```python
class Car:
    wheels = 4  # 클래스 변수 (모든 객체가 공유)

    def __init__(self, brand):
        self.brand = brand  # 인스턴스 변수

car1 = Car("Tesla")
car2 = Car("BMW")

print(car1.brand, car1.wheels)  # Tesla 4
print(car2.brand, car2.wheels)  # BMW 4
```

✔ `wheels`는 클래스 변수로 모든 객체가 공유  
✔ `brand`는 인스턴스 변수로 각 객체마다 다름  

---

## 4️⃣ 메서드의 종류

### 1) 인스턴스 메서드 (Instance Method)
- `self`를 사용하여 객체의 속성을 다루는 일반적인 메서드

```python
class Person:
    def __init__(self, name):
        self.name = name

    def greet(self):
        print(f"안녕하세요, 저는 {self.name}입니다.")

person1 = Person("Alice")
person1.greet()  # "안녕하세요, 저는 Alice입니다."
```

---

### 2) 클래스 메서드 (Class Method)
- `@classmethod` 데코레이터 사용
- `cls`를 매개변수로 받아 클래스 변수에 접근 가능

```python
class Animal:
    species = "동물"

    @classmethod
    def show_species(cls):
        print(f"이 클래스의 종(species)은 {cls.species}입니다.")

Animal.show_species()  # "이 클래스의 종(species)은 동물입니다."
```

✔ `cls`를 통해 클래스 변수에 접근 가능  

---

### 3) 정적 메서드 (Static Method)
- `@staticmethod` 데코레이터 사용
- `self`나 `cls` 없이 독립적인 기능 수행

```python
class Math:
    @staticmethod
    def add(a, b):
        return a + b

print(Math.add(3, 5))  # 8
```

✔ 정적 메서드는 클래스와 관계없이 독립적인 기능을 수행  

---

## 5️⃣ 클래스 상속 (Inheritance)

- **기존 클래스(부모 클래스, Superclass)의 기능을 그대로 물려받아 새로운 클래스를 생성**
- **코드 재사용** 및 **유지보수성** 향상

### 기본 상속 예제
```python
class Animal:
    def __init__(self, name):
        self.name = name

    def speak(self):
        print("소리를 냅니다.")

class Dog(Animal):  # Animal 클래스를 상속
    def speak(self):
        print(f"{self.name}가 멍멍!")

class Cat(Animal):  # Animal 클래스를 상속
    def speak(self):
        print(f"{self.name}가 야옹!")

dog = Dog("Buddy")
cat = Cat("Kitty")

dog.speak()  # "Buddy가 멍멍!"
cat.speak()  # "Kitty가 야옹!"
```

✔ `Dog`와 `Cat` 클래스는 `Animal` 클래스를 상속받아 사용  
✔ 부모 클래스의 기능을 그대로 사용하거나 **재정의(오버라이딩)** 가능  

---

## 6️⃣ 메서드 오버라이딩 (Method Overriding)

부모 클래스의 메서드를 자식 클래스에서 재정의(Override)하여 동작을 변경할 수 있음

```python
class Parent:
    def show(self):
        print("부모 클래스 메서드")

class Child(Parent):
    def show(self):
        print("자식 클래스에서 재정의된 메서드")

child = Child()
child.show()  # "자식 클래스에서 재정의된 메서드"
```

✔ 자식 클래스에서 동일한 이름의 메서드를 정의하면 부모 메서드를 덮어씀  

---

## 7️⃣ 캡슐화 (Encapsulation)

- 클래스 내부의 속성을 **외부에서 직접 접근하지 못하도록 보호**하는 개념
- 속성 앞에 **`_`(protected)** 또는 **`__`(private)** 를 붙여서 접근 제한 가능

### 캡슐화 예제
```python
class BankAccount:
    def __init__(self, balance):
        self.__balance = balance  # private 변수

    def deposit(self, amount):
        self.__balance += amount

    def get_balance(self):
        return self.__balance

account = BankAccount(1000)
account.deposit(500)
print(account.get_balance())  # 1500
```

✔ `__balance`는 클래스 내부에서만 접근 가능 (외부에서 `account.__balance` 직접 접근 불가)  

---

## 🎯 정리

✔ **클래스(Class)**: 객체를 생성하는 설계도  
✔ **객체(Object)**: 클래스로부터 만들어진 실체  
✔ **상속(Inheritance)**: 기존 클래스를 확장하여 새로운 클래스 생성  
✔ **오버라이딩(Overriding)**: 부모 클래스의 메서드를 자식 클래스에서 재정의  
✔ **캡슐화(Encapsulation)**: `__private` 속성으로 정보 보호  
✔ **메서드 종류**:
  - 인스턴스 메서드 (`self`)
  - 클래스 메서드 (`@classmethod`)
  - 정적 메서드 (`@staticmethod`)  
