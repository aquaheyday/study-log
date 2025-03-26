// ✅ JavaScript 객체 다루기 예제

// ---------------------------
// 1️⃣ 객체 생성
// ---------------------------
const user = {
  name: "Alice",
  age: 25,
  isAdmin: true
};

console.log(user); // 전체 객체 출력

// ---------------------------
// 2️⃣ 속성 접근 (점 표기법 & 대괄호 표기법)
// ---------------------------
console.log(user.name);      // Alice
console.log(user["age"]);    // 25

// 변수로 키 접근할 때는 대괄호 사용
const key = "isAdmin";
console.log(user[key]);      // true

// ---------------------------
// 3️⃣ 속성 추가 / 수정 / 삭제
// ---------------------------
user.email = "alice@example.com";  // 추가
user.age = 26;                     // 수정
delete user.isAdmin;              // 삭제

console.log(user);

// ---------------------------
// 4️⃣ 객체 메서드 정의
// ---------------------------
const person = {
  name: "Tom",
  sayHello: function () {
    console.log(`안녕하세요, 저는 ${this.name}입니다.`);
  },
  sayAge() {
    console.log("나이는 비밀입니다!");
  }
};

person.sayHello(); // 안녕하세요, 저는 Tom입니다.
person.sayAge();   // 나이는 비밀입니다!

// ---------------------------
// 5️⃣ 객체 안의 객체 (중첩 객체)
// ---------------------------
const company = {
  name: "TechCorp",
  ceo: {
    name: "Jane",
    age: 40
  }
};

console.log(company.ceo.name); // Jane

// ---------------------------
// 6️⃣ 객체 순회 (for...in)
// ---------------------------
for (const key in user) {
  console.log(`${key}: ${user[key]}`);
}

// ---------------------------
// 7️⃣ Object 메서드들
// ---------------------------
const product = {
  id: 101,
  name: "노트북",
  price: 1500000
};

console.log(Object.keys(product));   // ["id", "name", "price"]
console.log(Object.values(product)); // [101, "노트북", 1500000]
console.log(Object.entries(product)); // [["id", 101], ["name", "노트북"], ["price", 1500000]]

// ---------------------------
// 8️⃣ 객체 복사 (얕은 복사)
// ---------------------------
const original = { a: 1, b: 2 };
const copied = Object.assign({}, original);
copied.a = 100;

console.log(original.a); // 1 (원본은 안 바뀜)
console.log(copied.a);   // 100

// ---------------------------
// 9️⃣ 구조 분해 할당 (Destructuring)
// ---------------------------
const member = {
  username: "rainbow",
  level: 5
};

const { username, level } = member;
console.log(username); // rainbow
console.log(level);    // 5

// ---------------------------
// 🔟 객체 배열 활용 예시
// ---------------------------
const users = [
  { id: 1, name: "Anna" },
  { id: 2, name: "Brian" },
  { id: 3, name: "Chris" }
];

users.forEach((u) => {
  console.log(`${u.id}: ${u.name}`);
});
