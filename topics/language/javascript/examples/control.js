// ✅ JavaScript 조건문 & 반복문 예제

// ---------------------------
// 1️⃣ 조건문 (if, else if, else)
// ---------------------------
let score = 85;

if (score >= 90) {
  console.log("A 학점");
} else if (score >= 80) {
  console.log("B 학점");
} else if (score >= 70) {
  console.log("C 학점");
} else {
  console.log("F 학점");
}

// ---------------------------
// 2️⃣ 삼항 연산자 (조건 ? 참 : 거짓)
// ---------------------------
let age = 20;
let isAdult = (age >= 18) ? "성인" : "미성년자";
console.log(isAdult); // 성인

// ---------------------------
// 3️⃣ switch 문
// ---------------------------
let day = 3;

switch (day) {
  case 1:
    console.log("월요일");
    break;
  case 2:
    console.log("화요일");
    break;
  case 3:
    console.log("수요일");
    break;
  default:
    console.log("알 수 없는 요일");
}

// ---------------------------
// 4️⃣ for 반복문
// ---------------------------
for (let i = 1; i <= 5; i++) {
  console.log(`for 반복문: ${i}`);
}

// ---------------------------
// 5️⃣ while 반복문
// ---------------------------
let count = 1;
while (count <= 3) {
  console.log(`while 반복문: ${count}`);
  count++;
}

// ---------------------------
// 6️⃣ do...while 반복문
// ---------------------------
let num = 1;
do {
  console.log(`do...while 반복문: ${num}`);
  num++;
} while (num <= 2);

// ---------------------------
// 7️⃣ for...of 반복문 (배열 순회)
// ---------------------------
const fruits = ["🍎", "🍌", "🍇"];

for (const fruit of fruits) {
  console.log(`for...of: ${fruit}`);
}

// ---------------------------
// 8️⃣ for...in 반복문 (객체 속성 순회)
// ---------------------------
const user = {
  name: "Tom",
  age: 28,
  job: "Developer"
};

for (const key in user) {
  console.log(`${key}: ${user[key]}`);
}

// ---------------------------
// 9️⃣ break & continue
// ---------------------------
for (let i = 1; i <= 5; i++) {
  if (i === 3) continue; // 3은 건너뜀
  if (i === 5) break;    // 5에서 멈춤
  console.log(`break/continue 예제: ${i}`);
}
