// ✅ JavaScript 배열 메서드 예제

const numbers = [1, 2, 3, 4, 5];

// ---------------------------
// 1️⃣ push(), pop()
// ---------------------------
numbers.push(6);      // 배열 끝에 추가
console.log(numbers); // [1, 2, 3, 4, 5, 6]

numbers.pop();        // 배열 끝에서 제거
console.log(numbers); // [1, 2, 3, 4, 5]

// ---------------------------
// 2️⃣ unshift(), shift()
// ---------------------------
numbers.unshift(0);   // 배열 앞에 추가
console.log(numbers); // [0, 1, 2, 3, 4, 5]

numbers.shift();      // 배열 앞에서 제거
console.log(numbers); // [1, 2, 3, 4, 5]

// ---------------------------
// 3️⃣ forEach()
// ---------------------------
numbers.forEach((num, index) => {
  console.log(`인덱스 ${index}: 값 ${num}`);
});

// ---------------------------
// 4️⃣ map()
// ---------------------------
const squared = numbers.map(n => n * n);
console.log(squared); // [1, 4, 9, 16, 25]

// ---------------------------
// 5️⃣ filter()
// ---------------------------
const evens = numbers.filter(n => n % 2 === 0);
console.log(evens); // [2, 4]

// ---------------------------
// 6️⃣ find(), findIndex()
// ---------------------------
const found = numbers.find(n => n > 3);      // 4
const index = numbers.findIndex(n => n > 3); // 3
console.log(found, index);

// ---------------------------
// 7️⃣ includes(), indexOf()
// ---------------------------
console.log(numbers.includes(3));   // true
console.log(numbers.indexOf(4));    // 3

// ---------------------------
// 8️⃣ some(), every()
// ---------------------------
console.log(numbers.some(n => n > 4));  // true (하나라도 조건 만족)
console.log(numbers.every(n => n > 0)); // true (모두 조건 만족)

// ---------------------------
// 9️⃣ reduce()
// ---------------------------
const sum = numbers.reduce((acc, curr) => acc + curr, 0);
console.log(`합계: ${sum}`); // 합계: 15

// ---------------------------
// 🔟 sort(), reverse()
// ---------------------------
const items = [3, 1, 4, 2];
items.sort(); // 유니코드 순서 정렬 → [1, 2, 3, 4]
console.log(items);

items.reverse(); // 역순 정렬
console.log(items); // [4, 3, 2, 1]

// 숫자를 정확히 정렬하고 싶다면:
items.sort((a, b) => a - b); // 오름차순
console.log(items); // [1, 2, 3, 4]

// ---------------------------
// 1️⃣1️⃣ concat(), join(), slice(), splice()
// ---------------------------
const arr1 = [1, 2];
const arr2 = [3, 4];
const combined = arr1.concat(arr2);
console.log(combined); // [1, 2, 3, 4]

console.log(combined.join(" - ")); // "1 - 2 - 3 - 4"

console.log(combined.slice(1, 3)); // [2, 3] (1~2 인덱스)

combined.splice(2, 1, 99); // 2번 인덱스 제거 후 99 삽입
console.log(combined); // [1, 2, 99, 4]

// ---------------------------
// 1️⃣2️⃣ flat(), flatMap()
// ---------------------------
const nested = [1, [2, [3, 4]]];
console.log(nested.flat());       // [1, 2, [3, 4]]
console.log(nested.flat(2));      // [1, 2, 3, 4]

const flatMapped = [1, 2, 3].flatMap(n => [n, n * 2]);
console.log(flatMapped); // [1, 2, 2, 4, 3, 6]
