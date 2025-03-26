# 📦 JavaScript - 모듈 시스템(Module System)

**모듈**은 JavaScript 코드의 **재사용성과 유지보수성을 높이기 위해 파일 단위로 기능을 분리하는 방법**입니다.  
여러 파일로 코드를 나누고 필요한 부분만 **가져오거나 내보낼 수 있는 시스템**을 말합니다.

---

## 1️⃣ 모듈이 필요한 이유

- 전역 스코프 오염 방지 (`변수 충돌 문제`)
- 코드 재사용성 증가
- 유지보수 용이
- 의존성 명확화

---

## 2️⃣ ES Modules (ESM)

- 브라우저와 최신 Node.js에서 지원
- `.js` 또는 `.mjs` 확장자 사용
- `import` / `export` 문법 사용

#### 1. 내보내기 (`export`)

```js
// math.js
export const add = (a, b) => a + b;
export const sub = (a, b) => a - b;
```

```js
// 또는 기본 내보내기
export default function multiply(a, b) {
  return a * b;
}
```

#### 2. 가져오기 (`import`)

```js
// main.js
import { add, sub } from './math.js';
import multiply from './math.js';

console.log(add(2, 3));     // 5
console.log(multiply(4, 5)); // 20
```

#### 3. `type="module"` 속성을 `<script>` 태그에 붙여야 브라우저에서 작동:

```html
<script type="module" src="main.js"></script>
```

---

## 3️⃣ CommonJS

- **Node.js에서 기본으로 사용되는 모듈 시스템**
- `require()` / `module.exports` 문법 사용
- 브라우저에서는 기본적으로 작동하지 않음 (Webpack, Browserify 등 필요)

#### 1. 내보내기

```js
// calc.js
function divide(a, b) {
  return a / b;
}

module.exports = {
  divide
};
```

#### 2. 가져오기

```js
// app.js
const calc = require('./calc');
console.log(calc.divide(10, 2)); // 5
```

---

## 4️⃣ 차이점 비교 (ESM vs CommonJS)

| 항목               | ES Modules (ESM)                 | CommonJS                         |
|--------------------|----------------------------------|----------------------------------|
| 문법               | `import` / `export`             | `require()` / `module.exports`  |
| 실행 시점          | 정적(compile 시 분석)            | 동적(실행 중 로딩됨)             |
| 지원 환경         | 브라우저, Node.js (ESM 지원 시)  | Node.js (기본 방식)              |
| 비동기/동기        | 비동기 지원                      | 동기 방식                         |
| 파일 확장자        | `.js`, `.mjs`                    | `.js`                            |

---

## 5️⃣ 모듈 종류 정리

| 모듈 유형      | 사용 환경             | 대표 문법                 |
|----------------|------------------------|---------------------------|
| **ES Modules** | 최신 브라우저, Node.js | `import`, `export`        |
| **CommonJS**   | Node.js                | `require`, `module.exports` |
| **AMD**        | 브라우저 (RequireJS)   | `define()`, `require()`   |
| **UMD**        | Universal (ES + CJS)   | 다양한 환경 대응 가능     |
| **IIFE**       | 과거 모듈화 방식       | 즉시 실행 함수 사용       |

---

## 🎯 정리

✔ **모듈 시스템**은 코드 구조화와 재사용을 위한 필수 도구  
✔ 브라우저에서는 `ES Modules`, Node.js에서는 `CommonJS`를 주로 사용  
✔ `import/export`는 정적, `require()`는 동적으로 동작  
✔ 모듈 번들러 (Webpack, Rollup 등)를 활용하면 다양한 방식 혼용도 가능
