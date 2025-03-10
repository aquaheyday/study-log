# Python 객체지향 프로그래밍(OOP) 기초 정리

Python은 객체지향 프로그래밍(Object-Oriented Programming, OOP)을 지원하는 언어입니다.  
객체지향 프로그래밍은 코드의 재사용성을 높이고 유지보수를 쉽게 합니다.

---

## 1. 클래스(Class)와 객체(Object)

클래스는 객체를 생성하기 위한 청사진(설계도)이며, 객체는 클래스의 인스턴스(instance)입니다.

```python
# 클래스 정의
class Person:
    def __init__(self, name, age):
        self.name = name  # 인스턴스 변수
        self.age = age    # 인스턴스 변수
    
    def introduce(self):
        return f"My name is {self.name} and I am {self.age} years old."

# 객체 생성
person1 = Person("Alice", 25)
print(person1.introduce())  # My name is Alice and I am 25 years old.
```

---

## 2. 생성자(`__init__`)와 소멸자(`__del__`)

### 2.1 생성자 (`__init__`)
객체가 생성될 때 자동으로 호출됩니다.

```python
class Dog:
    def __init__(self, name):
        self.name = name
        print(f"{self.name}가 태어났습니다!")

dog = Dog("Buddy")  # Buddy가 태어났습니다!
```

### 2.2 소멸자 (`__del__`)
객체가 소멸될 때 자동으로 호출됩니다.

```python
class Animal:
    def __init__(self, name):
        self.name = name

    def __del__(self):
        print(f"{self.name} 객체가 삭제되었습니다.")

animal = Animal("Tiger")
del animal  # Tiger 객체가 삭제되었습니다.
```

---

## 3. 클래스 변수와 인스턴스 변수

- **클래스 변수**: 모든 객체가 공유하는 변수
- **인스턴스 변수**: 각 객체가 개별적으로 가지는 변수

```python
class Counter:
    count = 0  # 클래스 변수

    def __init__(self):
        Counter.count += 1  # 객체 생성 시마다 증가

print(Counter.count)  # 0
c1 = Counter()
c2 = Counter()
print(Counter.count)  # 2
```

---

## 4. 메서드(Method)의 종류

### 4.1 인스턴스 메서드
객체에서 호출하며, `self`를 통해 인스턴스 변수에 접근할 수 있습니다.

```python
class Car:
    def __init__(self, brand):
        self.brand = brand
    
    def show_info(self):
        return f"Car brand: {self.brand}"

car = Car("Toyota")
print(car.show_info())  # Car brand: Toyota
```

### 4.2 클래스 메서드 (`@classmethod`)
클래스 변수를 다룰 때 사용하며, `cls`를 통해 접근합니다.

```python
class Vehicle:
    count = 0  # 클래스 변수

    def __init__(self):
        Vehicle.count += 1

    @classmethod
    def get_count(cls):
        return f"Total Vehicles: {cls.count}"

v1 = Vehicle()
v2 = Vehicle()
print(Vehicle.get_count())  # Total Vehicles: 2
```

### 4.3 정적 메서드 (`@staticmethod`)
클래스나 인스턴스와 관계없이 독립적으로 동작합니다.

```python
class MathUtils:
    @staticmethod
    def add(a, b):
        return a + b

print(MathUtils.add(3, 5))  # 8
```

---

## 5. 상속(Inheritance)

기존 클래스(부모 클래스)를 재사용하여 새로운 클래스(자식 클래스)를 정의할 수 있습니다.

```python
class Animal:
    def __init__(self, name):
        self.name = name

    def speak(self):
        return "소리를 냅니다."

class Dog(Animal):  # Animal 클래스를 상속받음
    def speak(self):
        return "멍멍!"

dog = Dog("Buddy")
print(dog.speak())  # 멍멍!
```

### 5.1 `super()`를 이용한 부모 클래스 접근
```python
# 부모 클래스 (Animal)
class Animal:
    def __init__(self, name):
        self.name = name  # 이름 속성 정의

    def speak(self):
        return "소리를 냅니다."  # 기본적인 메서드

# 자식 클래스 (Bird)
class Bird(Animal):  # Animal 클래스를 상속받음
    def __init__(self, name, color):
        super().__init__(name)  # 부모 클래스의 생성자 호출
        self.color = color  # 새의 색깔 속성 추가

    def speak(self):  # 부모 클래스의 speak 메서드 오버라이딩
        return f"{self.color} {self.name}가 짹짹거립니다."

# 객체 생성 및 테스트
bird = Bird("참새", "갈색")  
print(bird.speak())  # 갈색 참새가 짹짹거립니다.
```

---

## 6. 다형성(Polymorphism)

같은 메서드 이름을 사용하더라도 객체에 따라 다르게 동작할 수 있습니다.

```python
class Animal:
    def speak(self):
        return "소리를 냅니다."

class Dog(Animal):
    def speak(self):
        return "멍멍!"

class Cat(Animal):
    def speak(self):
        return "야옹!"

animals = [Dog(), Cat(), Animal()]

for animal in animals:
    print(animal.speak())
```

출력:
```
멍멍!
야옹!
소리를 냅니다.
```

---

## 7. 추상 클래스(Abstract Class)

추상 클래스는 인스턴스를 만들 수 없으며, 반드시 **자식 클래스에서 메서드를 구현해야 합니다.**  
`abc` 모듈의 `ABC` 클래스를 상속받고, `@abstractmethod` 데코레이터를 사용합니다.

```python
from abc import ABC, abstractmethod

class Animal(ABC):
    @abstractmethod
    def speak(self):
        pass  # 반드시 자식 클래스에서 구현해야 함

class Dog(Animal):
    def speak(self):
        return "멍멍!"

dog = Dog()
print(dog.speak())  # 멍멍!
```

---

## 8. 캡슐화(Encapsulation)

캡슐화는 데이터 보호를 위해 변수의 접근 제한을 설정하는 기법입니다.

- **공개 멤버 (Public)**: `self.var`
- **보호 멤버 (Protected)**: `_var` (관례적으로 보호됨)
- **비공개 멤버 (Private)**: `__var` (외부 접근 불가)

```python
class Person:
    def __init__(self, name, age):
        self.name = name       # 공개 변수
        self._age = age        # 보호 변수
        self.__id = "1234ABCD" # 비공개 변수

    def get_id(self):
        return self.__id  # 비공개 변수는 내부에서만 접근 가능

person = Person("Alice", 25)
print(person.name)  # Alice
print(person._age)  # 25 (외부에서 접근 가능하지만, 사용을 권장하지 않음)
# print(person.__id)  # 오류 발생 (비공개 변수)
print(person.get_id())  # 1234ABCD (메서드를 통해 접근)
```
