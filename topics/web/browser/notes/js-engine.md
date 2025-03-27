# ⚙️ Browser - JavaScript 엔진 개요

## 1️⃣ JavaScript 엔진은
- 브라우저에서 **JS 코드를 해석하고 실행하는 핵심 컴포넌트**입니다.
- 브라우저마다 고유한 JS 엔진을 가지고 있으며, 모든 동적 웹 동작, DOM 조작, 이벤트 처리 등은 JS 엔진이 처리합니다.

---

## 2️⃣ 주요 JavaScript 엔진 종류

| 엔진 | 사용 브라우저 | 특징 |
|------|----------------|------|
| **V8** | Chrome, Edge, Opera, Brave 등 | 구글에서 개발, 빠른 성능, Node.js에서도 사용됨 |
| **SpiderMonkey** | Firefox | JS 최초의 엔진, Mozilla 개발 |
| **JavaScriptCore** | Safari | 애플의 JS 엔진, Nitro라고도 불림 |
| **Chakra (구형)** | 구버전 Edge | 현재는 V8 기반으로 변경됨 |

---

## 3️⃣ 동작 방식 (V8 기준)

JS 엔진은 단순히 코드를 해석하지 않고, **고속 실행을 위한 복잡한 처리**를 함

#### 1. Parsing (파싱)
   - JS 코드를 토큰화하고 AST(Abstract Syntax Tree) 생성

#### 2. Interpreter (Ignition)
   - AST를 바탕으로 **바이트코드**로 변환 후 실행

#### 3. JIT 컴파일 (TurboFan)
   - 반복 실행되는 코드 → **기계어(네이티브 코드)** 로 변환  
   - 최적화된 실행으로 속도 향상

✔️ 한 번만 쓰이는 코드는 빠르게 해석  
✔️ 자주 쓰이는 코드는 기계어로 바꿔 더 빠르게 실행  

---

### 1) AST(Abstract Syntax Tree)란?

JS 코드를 **문법적으로 쪼개어 계층 구조(Tree)** 로 나타낸 것

#### 예제 코드

```js
const x = 1 + 2;
```

#### 생성된 AST (간략 형태)

```
VariableDeclaration
├── VariableDeclarator
│   ├── Identifier (x)
│   └── BinaryExpression (+)
│       ├── Literal (1)
│       └── Literal (2)
```

#### 해석

- 변수 선언 → `VariableDeclaration`
- 변수 이름 → `Identifier: x`
- 값 → `BinaryExpression: 1 + 2`
  - 왼쪽 → `Literal: 1`
  - 오른쪽 → `Literal: 2`

---

### 2) 바이트코드란?

고급 언어(JS)를 저수준 명령어로 변환한 중간 코드, 아직 CPU가 바로 실행할 수는 없지만, 인터프리터가 해석할 수 있음

#### 예제 코드

```js
const x = 1 + 2;
```

#### 생성된 AST (간략 형태)

```
VariableDeclaration
├── VariableDeclarator
│   ├── Identifier (x)
│   └── BinaryExpression (+)
│       ├── Literal (1)
│       └── Literal (2)
```

#### 생성된 바이트코드
```js
LoadLiteral 2
LoadLiteral 3
Add
Store x
```

---

### 3) JIT 컴파일이란?

JIT(Just-In-Time) 컴파일은 자바스크립트 코드를 실행 도중에 반복적으로 실행되는 핫 코드(hot code)만 골라서 **기계어(네이티브 코드)** 로 변환해 성능을 극대화하는 기술

#### 예제 코드

```js
const x = 1 + 2;
```

#### 생성된 AST (간략 형태)

```
VariableDeclaration
├── VariableDeclarator
│   ├── Identifier (x)
│   └── BinaryExpression (+)
│       ├── Literal (1)
│       └── Literal (2)
```

#### 생성된 바이트코드

```js
LoadLiteral 2
LoadLiteral 3
Add
Store x
```

#### 바이트코드가 반복 실행시 JIT 적용된 코드

```
MOV R1, #1       ; R1 ← 1
ADD R1, R1, #2   ; R1 ← R1 + 2 (1 + 2)
MOV [x], R1      ; 전역 변수 x에 저장
```

✔️ 실제 기계어는 훨씬 복잡하고 플랫폼 의존적이지만, 핵심은 최대한 적은 명령어로 빠르게 처리하도록 바뀐다는 것입니다.

---

## 4️⃣ 엔진 내부 주요 구성 요소 (V8 예시)

| 구성 요소 | 설명 |
|-----------|------|
| **Parser** | JS 코드를 AST로 변환 |
| **Ignition (인터프리터)** | AST → 바이트코드 → 빠르게 실행 |
| **TurboFan (JIT 컴파일러)** | 바이트코드 → 최적화된 기계어 |
| **Garbage Collector** | 메모리 자동 회수 (Mark & Sweep 방식 등) |

---

## 5️⃣ 최적화 기법 (V8 기준)

- **Inline Caching**: 자주 쓰는 객체 구조는 미리 캐싱해서 빠르게 접근  
- **Hidden Class**: 객체 구조가 바뀌지 않도록 내부 클래스 생성  
- **Dead Code Elimination**: 사용되지 않는 코드는 컴파일에서 제거  
- **Escape Analysis**: 지역 객체는 힙이 아닌 스택에 저장해 성능 향상

---

## 6️⃣ 메모리 관리 (Garbage Collection)

JS는 수동 메모리 해제가 없고, JS 엔진이 자동으로 메모리를 관리함

- **Mark-and-Sweep 알고리즘**: 도달 가능한 객체만 남기고 나머지는 정리
- V8은 **Generational GC** 사용  
  - **New Space (젊은 세대)**: 짧게 살아있는 객체
  - **Old Space (오래된 객체)**: 살아남은 객체를 옮김

✔ 사용자가 따로 메모리 해제하지 않아도 됨  
✔ 단, 메모리 누수(예: 클로저, 이벤트 핸들러) 주의 필요

---

## 🎯 정리

✔ JavaScript 엔진 역할 → JS 코드를 해석하고 실행하는 핵심 엔진  
✔ 대표 엔진 → V8 (Chrome), SpiderMonkey (Firefox), JavaScriptCore (Safari)  
✔ 처리 과정 → 파싱 → 바이트코드 실행 → 반복 코드에 JIT 컴파일 적용  
✔ 최적화 → Inline caching, Dead code 제거, Hidden class, Escape analysis 등  
✔ 메모리 관리 → Mark & Sweep 기반의 자동 가비지 컬렉션 (GC)

