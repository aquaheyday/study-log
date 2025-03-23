# 🏗️ Go 언어 구조체와 메서드

Go는 클래스는 없지만 **구조체(struct)** 를 통해 객체와 유사한 데이터 구조를 만들고, **메서드**를 통해 동작을 정의할 수 있습니다.  

---

## 1️⃣ 구조체(Struct) 정의

- 구조체는 **필드의 집합**으로 이루어진 사용자 정의 타입입니다.
- `type` 키워드로 새로운 타입 정의

#### 기본 문법 예제

```go
type User struct {
    Name string
    Age  int
}
```

---

## 2️⃣ 구조체 초기화

필드 순서를 지켜야 하는 위치 기반 초기화는 `u1`처럼 가능  

```go
u1 := User{"Alice", 30}
u2 := User{Name: "Bob"}
u3 := new(User)      // 포인터 반환
u3.Name = "Carol"
```

✔ 추천 방식은 **필드명을 명시한 초기화 (`u2`)**  

---

## 3️⃣ 구조체 필드 접근

```go
fmt.Println(u1.Name) // "Alice"
u1.Age = 31
```

✔ 구조체 필드에 **점(.)으로 접근**하고 수정 가능  

---

## 4️⃣ 메서드 정의

#### 메서드 기본 문법

```go
func (u User) Greet() string {
    return "Hello, " + u.Name
}
```

- `(u User)` 부분은 **리시버(receiver)**: 메서드를 소유하는 타입
- 일반 함수와 다르게 특정 타입에 **속한 함수**

```go
fmt.Println(u1.Greet()) // "Hello, Alice"
```

---

## 5️⃣ 값 리시버 vs 포인터 리시버

| 구분 | 설명 | 사용 시점 |
|------|------|-----------|
| `func (u User)` | 값 복사 | 읽기 전용, 값 변형 없음 |
| `func (u *User)` | 포인터 전달 | 구조체 필드 변경 필요할 때 |

#### 포인터 리시버 예시

```go
func (u *User) SetName(name string) {
    u.Name = name
}
```

```go
u := User{Name: "Old"}
u.SetName("New")
fmt.Println(u.Name) // "New"
```

✔ 포인터 리시버는 **원본 구조체 값을 변경**할 수 있습니다.

---

## 6️⃣ 메서드와 함수 차이

| 항목 | 함수 | 메서드 |
|------|------|--------|
| 정의 방식 | `func 함수명()` | `func (t 타입) 메서드명()` |
| 소속     | 없음  | 특정 타입에 소속됨 |
| 호출 방식 | `함수()` | `인스턴스.메서드()` |

---

## 7️⃣ 익명 구조체 (Anonymous Struct)

```go
person := struct {
    Name string
    Age  int
}{
    Name: "Eve",
    Age:  28,
}
```

✔ 구조체를 간단하게 즉시 선언하여 사용할 수 있음

---

## 8️⃣ 중첩 구조체 (Nested Struct)

```go
type Address struct {
    City  string
    Zip   string
}

type Customer struct {
    Name    string
    Address Address
}
```

```go
c := Customer{
    Name: "Tom",
    Address: Address{
        City: "Seoul",
        Zip:  "12345",
    },
}
```

---

## 9️⃣ 구조체 임베딩 (Embedding)

```go
type Animal struct {
    Name string
}

func (a Animal) Speak() {
    fmt.Println(a.Name + " makes a sound")
}

type Dog struct {
    Animal // 임베딩
    Breed  string
}
```

```go
d := Dog{Animal: Animal{Name: "Buddy"}, Breed: "Shiba"}
d.Speak() // Buddy makes a sound
```

✔ 임베딩하면 **상속 없이 메서드 포함 가능**  
✔ Go의 **구성(Composition) 기반 설계**

---

## 🎯 정리

✔ Go의 구조체는 다양한 필드를 묶어 새로운 타입을 정의하는 도구입니다.  
✔ 메서드는 `func (r Type)` 형식으로 선언되며, 포인터 리시버를 사용하면 **구조체 필드 변경 가능**  
✔ Go는 클래스는 없지만, **구조체 + 메서드** + **인터페이스**를 통해 객체 지향적 설계가 가능합니다.  
✔ 구조체 임베딩은 **상속 없이 재사용성 확보**에 유용한 패턴입니다.

