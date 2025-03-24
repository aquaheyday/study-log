# 🚀 Flutter 애니메이션 기본

Flutter에서 애니메이션을 사용하면 UI를 더 생동감 있게 만들 수 있습니다.

---

## 1️⃣ 애니메이션의 종류

Flutter에서 제공하는 애니메이션 방식은 다음과 같습니다.

| 애니메이션 방식 | 설명 |
|---------------|------|
| `Implicit Animations` | 자동으로 애니메이션 효과 적용 (간단한 애니메이션) |
| `Explicit Animations` | 개발자가 직접 컨트롤하는 애니메이션 |
| `Tween Animations` | 값이 변하는 범위를 지정하는 애니메이션 |
| `Physics-based Animations` | 물리적인 힘(중력, 속도 등)을 적용하는 애니메이션 |

---

## 2️⃣ Implicit Animation (간단한 애니메이션)

### 1) `AnimatedContainer`
`AnimatedContainer`는 크기, 색상, 모양 등을 부드럽게 변화시킬 때 사용합니다.

```dart
import 'package:flutter/material.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: AnimatedContainerExample(),
    );
  }
}

class AnimatedContainerExample extends StatefulWidget {
  @override
  _AnimatedContainerExampleState createState() => _AnimatedContainerExampleState();
}

class _AnimatedContainerExampleState extends State<AnimatedContainerExample> {
  bool isBig = false;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text("AnimatedContainer Example")),
      body: Center(
        child: GestureDetector(
          onTap: () {
            setState(() {
              isBig = !isBig;
            });
          },
          child: AnimatedContainer(
            duration: Duration(seconds: 1),
            width: isBig ? 200 : 100,
            height: isBig ? 200 : 100,
            color: isBig ? Colors.blue : Colors.red,
            alignment: Alignment.center,
            child: Text("Tap Me!", style: TextStyle(color: Colors.white)),
          ),
        ),
      ),
    );
  }
}
```

✔ `AnimatedContainer` → 속성이 변경될 때 애니메이션 효과 적용  
✔ `duration` → 애니메이션 지속 시간 설정  
✔ `setState()` → 상태 변경 시 애니메이션 실행  

---

### 2) `AnimatedOpacity`
위젯의 투명도를 부드럽게 변경할 때 사용합니다.

```dart
AnimatedOpacity(
  duration: Duration(seconds: 1),
  opacity: isVisible ? 1.0 : 0.0,
  child: Text("Hello, Flutter!"),
);
```

✔ `opacity` → 1.0(보임), 0.0(숨김)  

---

## 3️⃣ Explicit Animation (컨트롤 가능한 애니메이션)

### 1) `AnimationController`와 `Tween`
`AnimationController`를 사용하면 애니메이션을 세밀하게 제어할 수 있습니다.

```dart
import 'package:flutter/material.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: ExplicitAnimationExample(),
    );
  }
}

class ExplicitAnimationExample extends StatefulWidget {
  @override
  _ExplicitAnimationExampleState createState() => _ExplicitAnimationExampleState();
}

class _ExplicitAnimationExampleState extends State<ExplicitAnimationExample> with SingleTickerProviderStateMixin {
  late AnimationController _controller;
  late Animation<double> _animation;

  @override
  void initState() {
    super.initState();
    _controller = AnimationController(
      duration: Duration(seconds: 2),
      vsync: this,
    )..repeat(reverse: true);

    _animation = Tween<double>(begin: 0, end: 300).animate(_controller);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text("Explicit Animation Example")),
      body: Center(
        child: AnimatedBuilder(
          animation: _animation,
          builder: (context, child) {
            return Container(
              width: _animation.value,
              height: _animation.value,
              color: Colors.blue,
            );
          },
        ),
      ),
    );
  }

  @override
  void dispose() {
    _controller.dispose();
    super.dispose();
  }
}
```

✔ `AnimationController` → 애니메이션 실행을 직접 제어  
✔ `Tween<double>(begin, end)` → 시작값과 종료값 설정  
✔ `AnimatedBuilder` → 애니메이션 값 변경 시 UI 업데이트  

---

## 4️⃣ Hero 애니메이션 (페이지 전환)

### 1) `Hero`
화면 전환 시 부드러운 이동 효과를 제공합니다.

```dart
// 첫 번째 화면
GestureDetector(
  onTap: () {
    Navigator.push(
      context,
      MaterialPageRoute(builder: (context) => SecondScreen()),
    );
  },
  child: Hero(
    tag: 'hero-tag',
    child: Container(
      width: 100,
      height: 100,
      color: Colors.blue,
    ),
  ),
);

// 두 번째 화면
Hero(
  tag: 'hero-tag',
  child: Container(
    width: 200,
    height: 200,
    color: Colors.blue,
  ),
);
```

✔ `Hero` → `tag` 속성이 동일한 위젯끼리 애니메이션 적용  

---

## 5️⃣ AnimatedList (리스트 애니메이션)

```dart
AnimatedList(
  initialItemCount: items.length,
  itemBuilder: (context, index, animation) {
    return SizeTransition(
      sizeFactor: animation,
      child: ListTile(title: Text(items[index])),
    );
  },
);
```

✔ `AnimatedList` → 리스트 아이템 추가/삭제 시 애니메이션 적용  

---

## 6️⃣ Lottie 애니메이션

#### 1. 패키지 설치

```sh
flutter pub add lottie
```

#### 2. JSON 애니메이션 적용

```dart
import 'package:flutter/material.dart';
import 'package:lottie/lottie.dart';

class LottieExample extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: Lottie.asset('assets/animation.json'),
      ),
    );
  }
}
```

✔ `Lottie.asset('path')` → JSON 애니메이션 적용  

---

## 🎯 정리

✔ `Implicit Animation` → `AnimatedContainer`, `AnimatedOpacity` 등 자동 애니메이션  
✔ `Explicit Animation` → `AnimationController`, `Tween` 등 세밀한 애니메이션  
✔ `Hero` → 페이지 전환 애니메이션  
✔ `AnimatedList` → 리스트 변경 애니메이션  
✔ `Lottie` → JSON 기반 애니메이션  
