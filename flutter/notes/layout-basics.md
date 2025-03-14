# 🔲 레이아웃 기본

Flutter에서 UI를 구성하는 기본 요소는 **위젯(Widget)** 이며, **레이아웃(Layout) 위젯**을 사용하여 다양한 화면 구성을 할 수 있습니다.

---

## 1. Flutter의 레이아웃 개념

Flutter에서 **모든 UI는 위젯을 조합하여 구성**됩니다.  
화면을 만들기 위해 **Container, Row, Column, Stack** 등의 **레이아웃 위젯**을 활용합니다.

✔ Flutter 레이아웃 위젯 특징
- 위젯 기반 트리 구조: 부모-자식 관계로 UI를 구성.
- 반응형 디자인 가능: `Flexible`, `Expanded`, `MediaQuery`를 사용하여 유동적인 레이아웃을 구현.
- 다양한 배치 방식 지원: 가로(Row), 세로(Column), 겹치기(Stack) 등.

---

## 2. 기본 레이아웃 위젯

### `Container` (박스 레이아웃)
크기, 색상, 패딩, 마진 등을 적용하는 박스 위젯.

```dart
Container(
  width: 200,
  height: 100,
  padding: EdgeInsets.all(10),
  margin: EdgeInsets.all(20),
  decoration: BoxDecoration(
    color: Colors.blue,
    borderRadius: BorderRadius.circular(10),
  ),
  child: Center(child: Text("Hello, Container!")),
)
```

---

### `Row` (가로 방향 배치)
위젯을 **가로(수평) 방향**으로 배치.

```dart
Row(
  mainAxisAlignment: MainAxisAlignment.spaceEvenly,
  children: [
    Icon(Icons.star, size: 40, color: Colors.yellow),
    Icon(Icons.star, size: 40, color: Colors.yellow),
    Icon(Icons.star, size: 40, color: Colors.yellow),
  ],
)
```

✔ 속성 설명
- `mainAxisAlignment`: **Row의 주축(가로) 정렬 방식** 지정  
  - `MainAxisAlignment.start` (왼쪽 정렬)
  - `MainAxisAlignment.center` (가운데 정렬)
  - `MainAxisAlignment.spaceBetween` (양 끝 정렬, 아이템 간격 없음)
  - `MainAxisAlignment.spaceAround` (양 끝 간격 반)
  - `MainAxisAlignment.spaceEvenly` (균등 간격 배치)
- `crossAxisAlignment`: **세로축 정렬 방식** (`start`, `center`, `end`)

---

### `Column` (세로 방향 배치)
위젯을 **세로(수직) 방향**으로 배치.

```dart
Column(
  mainAxisAlignment: MainAxisAlignment.center,
  crossAxisAlignment: CrossAxisAlignment.start,
  children: [
    Text("Item 1", style: TextStyle(fontSize: 24)),
    Text("Item 2", style: TextStyle(fontSize: 24)),
    Text("Item 3", style: TextStyle(fontSize: 24)),
  ],
)
```

✔ `Column`의 정렬 방식은 `Row`와 동일하게 설정 가능.

---

### `Stack` (겹쳐서 배치)
위젯을 **Z축(깊이) 방향으로 정렬** 하거나 여러 위젯을 **겹쳐서 배치**할 때 사용.

```dart
Stack(
  alignment: Alignment.center,
  children: [
    Container(width: 200, height: 200, color: Colors.blue),
    Container(width: 150, height: 150, color: Colors.red),
    Container(width: 100, height: 100, color: Colors.green),
  ],
)
```

✔ **위젯 순서대로 쌓이며, 가장 마지막에 선언한 위젯이 최상단에 위치**.

---

### `Expanded` & `Flexible` (공간 분배)
`Row` 또는 `Column` 내에서 **위젯 크기를 자동으로 조정**.

```dart
Row(
  children: [
    Expanded(
      flex: 2, // 2배 크기
      child: Container(color: Colors.blue, height: 100),
    ),
    Expanded(
      flex: 1, // 1배 크기
      child: Container(color: Colors.red, height: 100),
    ),
  ],
)
```

✔ `Expanded`와 `Flexible`의 차이:
- `Expanded` → 가능한 최대 공간을 차지.
- `Flexible` → 필요한 만큼만 차지하고 남은 공간 유지.

---

### `SizedBox` (크기 지정 및 공백 추가)
특정 크기의 공간을 만들거나, 위젯 크기 제한 가능.

```dart
SizedBox(
  width: 200,
  height: 50,
  child: ElevatedButton(
    onPressed: () {},
    child: Text("Click Me"),
  ),
)
```

✔ **공백 추가 용도로도 사용 가능**
```dart
SizedBox(height: 20) // 세로 여백 추가
```

---

## 3. 반응형 레이아웃 만들기

### `MediaQuery` (화면 크기 감지)
기기의 화면 크기에 맞게 UI를 동적으로 조정.

```dart
double screenWidth = MediaQuery.of(context).size.width;
double screenHeight = MediaQuery.of(context).size.height;
```

```dart
Container(
  width: MediaQuery.of(context).size.width * 0.8, // 화면 너비의 80%
  height: 200,
  color: Colors.green,
)
```

---

### `LayoutBuilder` (반응형 위젯)
부모 위젯의 크기에 따라 UI를 동적으로 변경.

```dart
LayoutBuilder(
  builder: (context, constraints) {
    if (constraints.maxWidth > 600) {
      return Row(
        children: [Expanded(child: Text("Large Screen"))],
      );
    } else {
      return Column(
        children: [Text("Small Screen")],
      );
    }
  },
)
```

---

## 4. 리스트와 스크롤 가능한 레이아웃

### `ListView` (스크롤 가능한 리스트)
많은 데이터를 표시할 때 사용.

```dart
ListView(
  children: [
    ListTile(title: Text("Item 1")),
    ListTile(title: Text("Item 2")),
    ListTile(title: Text("Item 3")),
  ],
)
```

---

### `GridView` (그리드 형태 레이아웃)
카드형 UI를 만들 때 유용.

```dart
GridView.count(
  crossAxisCount: 2, // 2열로 배치
  children: [
    Container(color: Colors.blue, height: 100),
    Container(color: Colors.red, height: 100),
  ],
)
```

---

## 🎯 정리

✔ Container → 크기, 색상, 여백을 설정하는 기본 박스 위젯  
✔ Row / Column → 가로 및 세로 정렬을 위한 핵심 레이아웃 위젯  
✔ Stack → 위젯을 겹쳐 배치 (Z축 활용)  
✔ Expanded / Flexible → 공간을 자동으로 분배  
✔ MediaQuery / LayoutBuilder → 반응형 UI 구현  
✔ ListView / GridView → 리스트 및 그리드 레이아웃 구성  
