# 🔲 네비게이션

Flutter에서 **네비게이션(Navigation)** 은 화면(페이지) 간 이동을 관리하는 중요한 개념입니다.  
Flutter의 네비게이션 시스템은 `Navigator`를 기반으로 하며, **페이지 이동(push/pop), 네임드 라우트(named routes), 탭 네비게이션(Tab Navigation), Drawer(사이드바 메뉴)** 등을 지원합니다.

---

## 1. Flutter 네비게이션 개념

- Flutter는 **스택(Stack) 기반 네비게이션**을 사용합니다.
- `Navigator.push()` → 새 화면을 스택에 추가 (페이지 이동)
- `Navigator.pop()` → 현재 화면을 스택에서 제거 (이전 페이지로 이동)

---

## 2. 기본 네비게이션 (`Navigator` 사용)

### `Navigator.push()` (새 페이지로 이동)

```dart
Navigator.push(
  context,
  MaterialPageRoute(builder: (context) => SecondPage()),
);
```

✔ `MaterialPageRoute` → 새로운 화면을 표시하는 기본적인 라우트 방식  
✔ `builder` → 이동할 페이지를 생성하는 함수  

---

### `Navigator.pop()` (이전 페이지로 이동)

```dart
Navigator.pop(context);
```

✔ 현재 페이지를 닫고 이전 페이지로 돌아감.

---

## 3. 네임드 라우트 (`Named Routes`)

Flutter는 **경로(URL) 기반의 네비게이션**을 지원하며, **네임드 라우트(named routes)**를 사용할 수 있습니다.

### 네임드 라우트 설정 (`MaterialApp`에 등록)

```dart
void main() {
  runApp(MaterialApp(
    initialRoute: '/',
    routes: {
      '/': (context) => HomePage(),
      '/second': (context) => SecondPage(),
    },
  ));
}
```

### `Navigator.pushNamed()`로 페이지 이동

```dart
Navigator.pushNamed(context, '/second');
```

### `Navigator.pop()`으로 뒤로가기

```dart
Navigator.pop(context);
```

✔ **네임드 라우트 사용 시의 장점**
- 여러 화면을 이동할 때 경로(URL) 기반으로 관리 가능
- 유지보수 및 코드 가독성이 좋아짐

---

## 4. 데이터 전달하기

### `push()`로 데이터 전달

```dart
Navigator.push(
  context,
  MaterialPageRoute(
    builder: (context) => SecondPage(data: "Hello from Home!"),
  ),
);
```

**SecondPage.dart**
```dart
class SecondPage extends StatelessWidget {
  final String data;

  SecondPage({required this.data});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text("Second Page")),
      body: Center(child: Text("받은 데이터: $data")),
    );
  }
}
```

---

### `pushNamed()`로 데이터 전달 (arguments 사용)
```dart
Navigator.pushNamed(
  context,
  '/second',
  arguments: "Hello from Home!",
);
```

**SecondPage.dart**
```dart
class SecondPage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    final String data = ModalRoute.of(context)!.settings.arguments as String;

    return Scaffold(
      appBar: AppBar(title: Text("Second Page")),
      body: Center(child: Text("받은 데이터: $data")),
    );
  }
}
```

✔ `arguments` → 네임드 라우트에서도 데이터를 전달할 수 있음.

---

## 5. `Navigator.pushReplacement()` (현재 페이지 대체)

현재 페이지를 새로운 페이지로 대체하고, 이전 페이지를 **스택에서 제거**.

```dart
Navigator.pushReplacement(
  context,
  MaterialPageRoute(builder: (context) => SecondPage()),
);
```

✔ 로그인 후 홈 화면으로 이동할 때 유용 (`로그인 페이지를 뒤로가기 불가능하게 만들기`)

---

## 6. `Navigator.pushAndRemoveUntil()` (이전 페이지 삭제)

- 특정 조건이 충족될 때까지 모든 이전 페이지를 제거하고 새 페이지로 이동.

```dart
Navigator.pushAndRemoveUntil(
  context,
  MaterialPageRoute(builder: (context) => HomePage()),
  (Route<dynamic> route) => false, // 모든 이전 페이지 제거
);
```

✔ 예제: 로그인 후, 이전의 로그인 페이지를 삭제하고 홈 화면으로 이동.

---

## 7. `BottomNavigationBar` (하단 탭 네비게이션)

앱 하단에 **탭바(Tab Bar)** 를 추가하여 여러 화면을 쉽게 전환할 수 있음.

```dart
class BottomNavExample extends StatefulWidget {
  @override
  _BottomNavExampleState createState() => _BottomNavExampleState();
}

class _BottomNavExampleState extends State<BottomNavExample> {
  int _selectedIndex = 0;

  final List<Widget> _pages = [
    HomePage(),
    SettingsPage(),
    ProfilePage(),
  ];

  void _onItemTapped(int index) {
    setState(() {
      _selectedIndex = index;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: _pages[_selectedIndex],
      bottomNavigationBar: BottomNavigationBar(
        items: [
          BottomNavigationBarItem(icon: Icon(Icons.home), label: "Home"),
          BottomNavigationBarItem(icon: Icon(Icons.settings), label: "Settings"),
          BottomNavigationBarItem(icon: Icon(Icons.person), label: "Profile"),
        ],
        currentIndex: _selectedIndex,
        onTap: _onItemTapped,
      ),
    );
  }
}
```

✔ `BottomNavigationBar` → 여러 페이지를 쉽게 전환 가능.

---

## 8. `Drawer` (사이드 네비게이션 메뉴)

왼쪽에서 열리는 **사이드바 메뉴 (햄버거 메뉴)** 를 추가할 수 있음.

```dart
class DrawerExample extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text("Drawer Example")),
      drawer: Drawer(
        child: ListView(
          children: [
            DrawerHeader(
              decoration: BoxDecoration(color: Colors.blue),
              child: Text("메뉴", style: TextStyle(color: Colors.white, fontSize: 24)),
            ),
            ListTile(
              title: Text("홈"),
              onTap: () {
                Navigator.pushNamed(context, "/");
              },
            ),
            ListTile(
              title: Text("설정"),
              onTap: () {
                Navigator.pushNamed(context, "/settings");
              },
            ),
          ],
        ),
      ),
      body: Center(child: Text("메인 화면")),
    );
  }
}
```

✔ **Drawer** → 사이드바 메뉴를 쉽게 추가할 수 있음.

---

## 🎯 정리

✔ `Navigator.push()` → 새 페이지로 이동  
✔ `Navigator.pop()` → 이전 페이지로 돌아가기  
✔ `pushReplacement()` → 현재 페이지를 새로운 페이지로 대체  
✔ `pushAndRemoveUntil()` → 특정 페이지까지 스택을 정리  
✔ `pushNamed()` & `arguments` → 네임드 라우트와 데이터 전달  
✔ `BottomNavigationBar` → 하단 탭 네비게이션  
✔ `Drawer` → 사이드바 네비게이션 메뉴  
