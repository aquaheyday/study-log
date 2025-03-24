# Flutter 레이아웃 구성 정리

Flutter의 레이아웃 시스템은 위젯을 조합하여 다양한 UI를 만들 수 있도록 설계되었습니다. `Column`, `Row`, `Stack` 등의 기본 레이아웃 위젯과 `Expanded`, `Flexible`, `Spacer` 등을 적절히 활용하면 효과적인 레이아웃 구성이 가능합니다.

---

## 1. 기본 레이아웃 위젯

### `Container`
`Container`는 박스 형태의 위젯으로, 크기, 색상, 정렬 등을 설정할 수 있습니다.

```dart
Container(
  width: 200,
  height: 100,
  color: Colors.blue,
  child: Center(
    child: Text('Container', style: TextStyle(color: Colors.white)),
  ),
)
```

### `Column`
위젯을 세로 방향으로 정렬할 때 사용합니다.

```dart
Column(
  mainAxisAlignment: MainAxisAlignment.center,
  children: [
    Text('First Line'),
    Text('Second Line'),
  ],
)
```

### `Row`
위젯을 가로 방향으로 정렬할 때 사용합니다.

```dart
Row(
  mainAxisAlignment: MainAxisAlignment.spaceAround,
  children: [
    Icon(Icons.star),
    Icon(Icons.favorite),
    Icon(Icons.thumb_up),
  ],
)
```

---

## 2. 정렬 및 패딩

### `Alignment`
Flutter의 `Container` 또는 `Align` 위젯을 사용하여 위젯을 정렬할 수 있습니다.

```dart
Container(
  alignment: Alignment.center,
  child: Text('Centered Text'),
)
```

```dart
Align(
  alignment: Alignment.center,
  child: Text(
    'Centered Text',
    style: TextStyle(fontSize: 24),
  ),
)
```

### `Padding`
`Padding` 위젯을 사용하면 위젯의 내부 여백을 설정할 수 있습니다.

```dart
Padding(
  padding: EdgeInsets.all(16.0),
  child: Text('Padded Text'),
)
```

---

## 3. `Expanded` 와 `Flexible`

### `Expanded`
`Expanded`는 `Column`이나 `Row` 내에서 자식 위젯이 가용 가능한 모든 공간을 차지하도록 만듭니다.

```dart
Row(
  children: [
    Expanded(child: Container(color: Colors.red, height: 100)),
    Expanded(child: Container(color: Colors.green, height: 100)),
    Expanded(child: Container(color: Colors.blue, height: 100)),
  ],
)
```

### `Flexible`
`Flexible`은 `Expanded`와 유사하지만, 위젯의 크기를 유동적으로 조정할 수 있습니다.

```dart
Row(
  children: [
    Flexible(
      flex: 2,
      child: Container(color: Colors.orange, height: 100),
    ),
    Flexible(
      flex: 1,
      child: Container(color: Colors.purple, height: 100),
    ),
  ],
)
```

---

## 4. `Stack`과 `Positioned` 위젯
### `Stack`
`Stack`을 사용하면 여러 위젯을 겹쳐서 배치할 수 있습니다.

```dart
Stack(
  children: [
    Container(width: 200, height: 200, color: Colors.blue),
    Positioned(
      top: 50,
      left: 50,
      child: Container(width: 100, height: 100, color: Colors.red),
    ),
  ],
)
```

---

## 5. 리스트 및 그리드

### `ListView`
스크롤 가능한 리스트를 만들 때 사용합니다.

```dart
ListView(
  children: [
    ListTile(title: Text('Item 1')),
    ListTile(title: Text('Item 2')),
    ListTile(title: Text('Item 3')),
  ],
)
```

### `GridView`
격자 형태로 위젯을 배치할 때 사용합니다.

```dart
GridView.count(
  crossAxisCount: 2,
  children: [
    Container(color: Colors.red, height: 100),
    Container(color: Colors.green, height: 100),
    Container(color: Colors.blue, height: 100),
    Container(color: Colors.orange, height: 100),
  ],
)
```

---

## 6. `SizedBox`와 `Spacer`

### `SizedBox`
고정된 크기의 빈 공간을 만들 때 사용합니다.

```dart
SizedBox(height: 20),
```

### `Spacer`
`Column`이나 `Row`에서 남은 공간을 차지하도록 만듭니다.

```dart
Row(
  children: [
    Text('Start'),
    Spacer(),
    Text('End'),
  ],
)
```

---

## 7. 반응형 UI 만들기 (`MediaQuery` & `LayoutBuilder`)

### `MediaQuery`
화면 크기를 기반으로 레이아웃을 조정할 수 있습니다.

```dart
double screenWidth = MediaQuery.of(context).size.width;
```

### `LayoutBuilder`
위젯의 부모 크기에 따라 레이아웃을 조정할 수 있습니다.

```dart
LayoutBuilder(
  builder: (context, constraints) {
    if (constraints.maxWidth > 600) {
      return Text('Large Screen');
    } else {
      return Text('Small Screen');
    }
  },
)
```

