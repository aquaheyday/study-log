// ✅ JavaScript async/await 예제

// ---------------------------
// 1️⃣ 기본 구조
// ---------------------------
function delay(ms) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

async function sayHello() {
  console.log("1. 시작");
  await delay(1000); // 1초 기다림
  console.log("2. 안녕하세요!");
  await delay(500);
  console.log("3. 반가워요~");
}

sayHello();

// ---------------------------
// 2️⃣ 비동기 함수에서 값 반환
// ---------------------------
async function getData() {
  return "🎉 데이터 도착!";
}

getData().then((result) => {
  console.log("데이터:", result); // 🎉 데이터 도착!
});

// ---------------------------
// 3️⃣ await와 Promise 함께 사용
// ---------------------------
function fetchData(name, delayMs) {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve(`✅ ${name} 완료`);
    }, delayMs);
  });
}

async function runTasks() {
  const result1 = await fetchData("작업 1", 1000);
  console.log(result1); // ✅ 작업 1 완료

  const result2 = await fetchData("작업 2", 500);
  console.log(result2); // ✅ 작업 2 완료
}

runTasks();

// ---------------------------
// 4️⃣ try...catch 에러 처리
// ---------------------------
function fetchWithError() {
  return new Promise((_, reject) => {
    setTimeout(() => {
      reject("❌ 에러 발생!");
    }, 700);
  });
}

async function errorHandlingExample() {
  try {
    const result = await fetchWithError();
    console.log(result);
  } catch (error) {
    console.error("에러 캐치:", error); // ❌ 에러 발생!
  } finally {
    console.log("마무리 작업 (finally)");
  }
}

errorHandlingExample();

// ---------------------------
// 5️⃣ 여러 Promise 병렬 처리
// ---------------------------
async function runInParallel() {
  const [a, b, c] = await Promise.all([
    fetchData("병렬 작업 A", 800),
    fetchData("병렬 작업 B", 300),
    fetchData("병렬 작업 C", 500)
  ]);

  console.log(a); // ✅ 병렬 작업 A 완료
  console.log(b); // ✅ 병렬 작업 B 완료
  console.log(c); // ✅ 병렬 작업 C 완료
}

runInParallel();
