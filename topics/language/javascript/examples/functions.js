// ✅ JavaScript 함수 예제 모음

// ---------------------------
// 1️⃣ 함수 선언식 (Function Declaration)
// ---------------------------
function greet(name) {
  console.log(`안녕하세요, ${name}님!`);
}
greet("Alice"); // 안녕하세요, Alice님!

// ---------------------------
// 2️⃣ 함수 표현식 (Function Expression)
// ---------------------------
const sayHi = function () {
  console.log("Hi there!");
};
sayHi(); // Hi there!

// ---------------------------
// 3️⃣ 화살표 함수 (Arrow Function)
// ---------------------------
const add = (a, b) => a + b;
console.log(add(3, 4)); // 7

const multiply = (a, b) => {
  const result = a * b;
  return result;
};
console.log(multiply(2, 5)); // 10

// ---------------------------
// 4️⃣ 매개변수 기본값
// ---------------------------
function greetUser(name = "게스트") {
  console.log(`반가워요, ${name}님!`);
}
greetUser();         // 반가워요, 게스트님!
greetUser("Tom");    // 반가워요, Tom님!

// ---------------------------
// 5️⃣ 반환값 (return)
// ---------------------------
function square(x) {
  return x * x;
}
const result = square(5);
console.log(`제곱 결과: ${result}`); // 제곱 결과: 25

// ---------------------------
// 6️⃣ 콜백 함수 (함수를 매개변수로 전달)
// ---------------------------
function processUserInput(callback) {
  const name = "Jane";
  callback(name);
}

processUserInput(function (name) {
  console.log(`콜백으로 받은 이름: ${name}`);
});

// ---------------------------
// 7️⃣ 함수 안에 함수 (중첩 함수)
// ---------------------------
function outer() {
  console.log("바깥 함수");

  function inner() {
    console.log("안쪽 함수");
  }

  inner();
}
outer();
// 바깥 함수
// 안쪽 함수

// ---------------------------
// 8️⃣ 재귀 함수 (자기 자신을 호출)
// ---------------------------
function countdown(n) {
  if (n <= 0) {
    console.log("출발!");
    return;
  }
  console.log(n);
  countdown(n - 1);
}
countdown(3);
// 3
// 2
// 1
// 출발!

// ---------------------------
// 9️⃣ 익명 함수 & 즉시 실행 함수 (IIFE)
// ---------------------------
(function () {
  console.log("즉시 실행 함수 실행됨!");
})();

// ---------------------------
// 🔟 함수 타입 확인
// ---------------------------
console.log(typeof greet);  // function
