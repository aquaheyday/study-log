// ✅ JavaScript 클로저 & 스코프 예제

// ---------------------------
// 1️⃣ 스코프 (Scope) 기본
// ---------------------------

let globalVar = "나는 전역 변수";

function outerScope() {
  let outerVar = "나는 외부 함수 변수";

  function innerScope() {
    let innerVar = "나는 내부 함수 변수";

    console.log(globalVar); // ✅ 접근 가능
    console.log(outerVar);  // ✅ 접근 가능
    console.log(innerVar);  // ✅ 접근 가능
  }

  innerScope();
  // console.log(innerVar); // ❌ 오류: 내부 함수 변수는 외부에서 접근 불가
}

outerScope();

// ---------------------------
// 2️⃣ 함수 외부에서 선언된 변수 접근 (클로저)
// ---------------------------

function makeCounter() {
  let count = 0;

  return function () {
    count++;
    console.log(`현재 카운트: ${count}`);
  };
}

const counter = makeCounter();
counter(); // 현재 카운트: 1
counter(); // 현재 카운트: 2
counter(); // 현재 카운트: 3

// 위 함수는 `count` 변수에 계속 접근할 수 있음 → 클로저!

// ---------------------------
// 3️⃣ 클로저로 private 변수 흉내내기
// ---------------------------

function createUser(name) {
  let _secret = "비밀";

  return {
    getName: function () {
      return name;
    },
    getSecret: function () {
      return _secret;
    },
    setSecret: function (newSecret) {
      _secret = newSecret;
    }
  };
}

const user = createUser("Alice");

console.log(user.getName());   // Alice
console.log(user.getSecret()); // 비밀
user.setSecret("새로운 비밀");
console.log(user.getSecret()); // 새로운 비밀

// _secret은 외부에서 직접 접근 불가 (private처럼 작동)

// ---------------------------
// 4️⃣ 반복문 클로저 주의 (var vs let)
// ---------------------------

// ❌ 문제 발생: var는 함수 스코프라서 클로저가 공유됨
for (var i = 0; i < 3; i++) {
  setTimeout(() => {
    console.log(`var i: ${i}`); // 항상 3이 출력됨
  }, 100);
}

// ✅ 해결: let은 블록 스코프 → 클로저가 각기 다름
for (let j = 0; j < 3; j++) {
  setTimeout(() => {
    console.log(`let j: ${j}`); // 0, 1, 2 출력됨
  }, 100);
}
