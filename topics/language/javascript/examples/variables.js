// 📦 변수 선언과 데이터 타입 예제

// ✅ 1. 변수 선언 방식
var a = 10;            // 함수 레벨 스코프 (지양)
let b = 20;            // 블록 레벨 스코프
const c = 30;          // 상수, 재할당 불가

// b = 25;            // 가능
// c = 35;            // ❌ 에러 발생 (const는 재할당 불가)

// ✅ 2. 기본 데이터 타입 (Primitive Types)
let str = "Hello";     // 문자열 (string)
let num = 42;          // 숫자 (number)
let bool = true;       // 불리언 (boolean)
let undef;             // undefined (초기화되지 않음)
let nul = null;        // null (값이 없음을 명시)
let sym = Symbol("id");// 심볼 (고유값)
let bigInt = 123456789012345678901234567890n; // 빅인트 (BigInt)

// typeof 연산자로 타입 확인
console.log(typeof str);    // string
console.log(typeof num);    // number
console.log(typeof bool);   // boolean
console.log(typeof undef);  // undefined
console.log(typeof nul);    // object (자바스크립트의 버그로 유명)
console.log(typeof sym);    // symbol
console.log(typeof bigInt); // bigint

// ✅ 3. 참조 타입 (Reference Types)
let arr = [1, 2, 3];        // 배열
let obj = { name: "Tom", age: 25 }; // 객체
let func = function () {
  console.log("나는 함수입니다");
};

// typeof 연산자로 확인
console.log(typeof arr);    // object
console.log(typeof obj);    // object
console.log(typeof func);   // function

// ✅ 4. 동적 타입 (Dynamic Typing)
let dynamic = "문자열";
console.log(typeof dynamic); // string

dynamic = 100;
console.log(typeof dynamic); // number

dynamic = { key: "value" };
console.log(typeof dynamic); // object

// ✅ 5. 변수 선언 없이 할당 (암묵적 전역변수, 지양!)
globalVar = "나는 전역변수야!";
console.log(globalVar);      // 브라우저에선 window.globalVar 로 접근됨

// ✅ 6. 템플릿 문자열 (Template Literals)
let name = "Alice";
let greeting = `안녕하세요, ${name}님!`;
console.log(greeting); // 안녕하세요, Alice님!
