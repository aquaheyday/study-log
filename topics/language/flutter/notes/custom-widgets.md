# 🚀 Flutter 커스텀 위젯 만들기

Flutter에서 UI를 효율적으로 관리하기 위해 **재사용 가능한 커스텀 위젯**을 만드는 방법을 정리합니다.

---

## 1️⃣ 커스텀 위젯이 필요한 이유
- 코드 중복 제거 → 같은 UI 요소를 여러 번 사용 가능  
- 유지보수 용이 → UI 수정이 필요할 때 한 곳만 변경하면 됨  
- 구조화된 코드 → 코드가 더 깔끔하고 읽기 쉬움  

---

## 2️⃣ StatelessWidget을 이용한 커스텀 위젯

#### 기본적인 커스텀 버튼 만들기
```dart
import 'package:flutter/material.dart';

class CustomButton extends StatelessWidget {
  final String text;
  final VoidCallback onPressed;

  const CustomButton({required this.text, required this.onPressed, Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ElevatedButton(
      onPressed: onPressed,
      style: ElevatedButton.styleFrom(
        padding: EdgeInsets.symmetric(horizontal: 20, vertical: 10),
        textStyle: TextStyle(fontSize: 18),
      ),
      child: Text(text),
    );
  }
}

// 사용 예시
CustomButton(
  text: "클릭하세요",
  onPressed: () {
    print("버튼 클릭됨");
  },
);
```

✔ `StatelessWidget` → 상태가 없는 커스텀 위젯  
✔ `required` 키워드 → 필수 매개변수 설정  
✔ `VoidCallback` → 버튼 클릭 시 실행할 함수 전달  

---

## 3️⃣ StatefulWidget을 이용한 커스텀 위젯

#### 상태를 가지는 커스텀 토글 버튼 만들기

```dart
import 'package:flutter/material.dart';

class ToggleButton extends StatefulWidget {
  @override
  _ToggleButtonState createState() => _ToggleButtonState();
}

class _ToggleButtonState extends State<ToggleButton> {
  bool isOn = false;

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () {
        setState(() {
          isOn = !isOn;
        });
      },
      child: Container(
        padding: EdgeInsets.all(10),
        decoration: BoxDecoration(
          color: isOn ? Colors.green : Colors.grey,
          borderRadius: BorderRadius.circular(10),
        ),
        child: Text(
          isOn ? "ON" : "OFF",
          style: TextStyle(color: Colors.white, fontSize: 20),
        ),
      ),
    );
  }
}

// 사용 예시
ToggleButton();
```

✔ `StatefulWidget` → 상태를 가지는 커스텀 위젯  
✔ `setState()` → UI 업데이트  

---

## 4️⃣ 매개변수가 있는 커스텀 위젯

#### 이미지 카드 위젯 만들기

```dart
import 'package:flutter/material.dart';

class ImageCard extends StatelessWidget {
  final String imageUrl;
  final String title;

  const ImageCard({required this.imageUrl, required this.title, Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Card(
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(15)),
      elevation: 5,
      child: Column(
        children: [
          ClipRRect(
            borderRadius: BorderRadius.vertical(top: Radius.circular(15)),
            child: Image.network(imageUrl, height: 150, width: double.infinity, fit: BoxFit.cover),
          ),
          Padding(
            padding: EdgeInsets.all(10),
            child: Text(title, style: TextStyle(fontSize: 18, fontWeight: FontWeight.bold)),
          ),
        ],
      ),
    );
  }
}

// 사용 예시
ImageCard(
  imageUrl: "https://via.placeholder.com/150",
  title: "샘플 이미지",
);
```

✔ `ClipRRect` → 이미지 모서리 둥글게 만들기  
✔ `Card` → 기본적인 카드 형태 제공  

---

## 5️⃣ ListView에서 재사용 가능한 아이템 위젯

#### 리스트 아이템 만들기

```dart
class ListItem extends StatelessWidget {
  final String title;
  final String subtitle;

  const ListItem({required this.title, required this.subtitle, Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ListTile(
      title: Text(title, style: TextStyle(fontWeight: FontWeight.bold)),
      subtitle: Text(subtitle),
      trailing: Icon(Icons.arrow_forward),
    );
  }
}

// ListView에서 사용
ListView(
  children: [
    ListItem(title: "첫 번째 아이템", subtitle: "설명 1"),
    ListItem(title: "두 번째 아이템", subtitle: "설명 2"),
  ],
);
```

✔ `ListTile` → 리스트 형식의 UI 쉽게 구성 가능  
✔ `trailing` → 리스트 아이템의 우측에 아이콘 추가  

---

## 6️⃣ 테마를 활용한 커스텀 위젯

#### 공통 스타일을 적용한 커스텀 버튼

```dart
class ThemedButton extends StatelessWidget {
  final String text;
  final VoidCallback onPressed;

  const ThemedButton({required this.text, required this.onPressed, Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ElevatedButton(
      onPressed: onPressed,
      style: ElevatedButton.styleFrom(
        primary: Theme.of(context).primaryColor,
        padding: EdgeInsets.symmetric(horizontal: 20, vertical: 10),
        textStyle: TextStyle(fontSize: 18),
      ),
      child: Text(text),
    );
  }
}
```

✔ `Theme.of(context).primaryColor` → 앱의 테마 색상을 가져와 적용  

---

## 7️⃣ 애니메이션이 포함된 커스텀 위젯

#### 크기가 변화하는 애니메이션 위젯

```dart
class AnimatedBox extends StatefulWidget {
  @override
  _AnimatedBoxState createState() => _AnimatedBoxState();
}

class _AnimatedBoxState extends State<AnimatedBox> {
  bool isExpanded = false;

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () {
        setState(() {
          isExpanded = !isExpanded;
        });
      },
      child: AnimatedContainer(
        duration: Duration(seconds: 1),
        width: isExpanded ? 200 : 100,
        height: isExpanded ? 200 : 100,
        color: isExpanded ? Colors.blue : Colors.red,
      ),
    );
  }
}
```

✔ `AnimatedContainer` → 상태가 변경될 때 부드러운 애니메이션 적용  

---

## 🎯 정리

✔ `StatelessWidget` → `CustomButton`, `ImageCard`처럼 상태 없는 위젯 제작  
✔ `StatefulWidget` → `ToggleButton`, `AnimatedBox`처럼 상태가 있는 위젯 제작  
✔ `재사용 가능한 UI` → 버튼, 카드, 리스트 아이템 등 공통적으로 사용되는 UI 제작  
✔ `테마 활용` → `Theme.of(context).primaryColor`를 사용하여 일관된 스타일 유지  
