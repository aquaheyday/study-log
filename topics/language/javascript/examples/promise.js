// ✅ JavaScript Promise 사용 예제

// ---------------------------
// 1️⃣ 기본 Promise 생성
// ---------------------------
const promise1 = new Promise((resolve, reject) => {
  const success = true;

  if (success) {
    resolve("✅ 성공했습니다!");
  } else {
    reject("❌ 실패했습니다.");
  }
});

// ---------------------------
// 2️⃣ then, catch 사용
// ---------------------------
promise1
  .then((result) => {
    console.log("1. then:", result);
  })
  .catch((error) => {
    console.error("2. catch:", error);
  })
  .finally(() => {
    console.log("3. finally: 항상 실행됨");
  });

// ---------------------------
// 3️⃣ 비동기 작업 시뮬레이션 (setTimeout)
// ---------------------------
function asyncTask(message, delay) {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve(`✅ 완료: ${message}`);
    }, delay);
  });
}

asyncTask("첫 번째 작업", 1000).then(console.log); // 1초 후 완료 메시지

// ---------------------------
// 4️⃣ Promise 체이닝 (Chaining)
// ---------------------------
asyncTask("Step 1", 500)
  .then((res) => {
    console.log(res);
    return asyncTask("Step 2", 1000);
  })
  .then((res) => {
    console.log(res);
    return asyncTask("Step 3", 500);
  })
  .then((res) => {
    console.log(res);
  });

// ---------------------------
// 5️⃣ 에러 처리
// ---------------------------
function taskWithError() {
  return new Promise((resolve, reject) => {
    const error = true;
    if (error) reject("⚠️ 에러 발생!");
    else resolve("문제 없음");
  });
}

taskWithError()
  .then(console.log)
  .catch((err) => {
    console.error("에러 캐치:", err);
  });

// ---------------------------
// 6️⃣ Promise.all (병렬 실행)
// ---------------------------
const task1 = asyncTask("A", 1000);
const task2 = asyncTask("B", 500);
const task3 = asyncTask("C", 800);

Promise.all([task1, task2, task3]).then((results) => {
  console.log("✅ 모든 작업 완료:");
  console.log(results); // ["A 결과", "B 결과", "C 결과"]
});

// ---------------------------
// 7️⃣ async/await (Promise 사용의 편한 버전)
// ---------------------------
async function runTasks() {
  try {
    const result1 = await asyncTask("🌟 비동기 1", 700);
    console.log(result1);

    const result2 = await asyncTask("🌟 비동기 2", 700);
    console.log(result2);
  } catch (e) {
    console.error("에러:", e);
  }
}

runTasks();
