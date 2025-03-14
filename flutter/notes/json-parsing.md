# 🌍 JSON 데이터 처리

Flutter에서 API 응답을 JSON으로 받아와 Dart 객체로 변환하거나, Dart 객체를 JSON으로 변환하는 방법을 학습합니다.

---

## 1. JSON 데이터란?

JSON(JavaScript Object Notation)은 **텍스트 기반 데이터 형식**으로 API와 데이터를 주고받을 때 많이 사용됩니다.

```json
{
  "id": 1,
  "name": "Flutter",
  "description": "JSON 데이터 처리 예제"
}
```

✔ **키-값 쌍** 형태의 데이터 구조  
✔ Flutter에서 `dart:convert` 패키지를 사용하여 변환 가능  

---

## 2. JSON 데이터를 Dart 객체로 변환 (디코딩)

### `dart:convert`를 사용한 JSON 디코딩

JSON 문자열을 **Dart 객체(Map/List)**로 변환할 수 있습니다.

```dart
import 'dart:convert';

void main() {
  String jsonString = '{"id": 1, "name": "Flutter", "description": "JSON 예제"}';
  
  Map<String, dynamic> jsonData = jsonDecode(jsonString);
  
  print(jsonData['name']); // Flutter
}
```

✔ `jsonDecode()` → JSON 문자열을 Dart 객체(Map)로 변환  

---

## 3. Dart 객체를 JSON 데이터로 변환 (인코딩)

Dart 객체를 **JSON 문자열**로 변환할 수 있습니다.

```dart
import 'dart:convert';

void main() {
  Map<String, dynamic> data = {
    'id': 1,
    'name': 'Flutter',
    'description': 'JSON 예제'
  };
  
  String jsonString = jsonEncode(data);
  
  print(jsonString); // {"id":1,"name":"Flutter","description":"JSON 예제"}
}
```

✔ `jsonEncode()` → Dart 객체를 JSON 문자열로 변환  

---

## 4. JSON 데이터를 모델 클래스로 변환

Flutter에서는 JSON 데이터를 쉽게 다루기 위해 **모델 클래스를 정의**하는 것이 일반적입니다.

### 모델 클래스 정의

```dart
class User {
  final int id;
  final String name;
  final String email;

  User({required this.id, required this.name, required this.email});

  // JSON -> Dart 객체 (디코딩)
  factory User.fromJson(Map<String, dynamic> json) {
    return User(
      id: json['id'],
      name: json['name'],
      email: json['email'],
    );
  }

  // Dart 객체 -> JSON (인코딩)
  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'name': name,
      'email': email,
    };
  }
}
```

✔ `factory User.fromJson(Map<String, dynamic> json)` → JSON 데이터를 Dart 객체로 변환  
✔ `toJson()` → Dart 객체를 JSON 데이터로 변환  

---

### API에서 JSON을 가져와 모델 클래스로 변환

```dart
import 'package:http/http.dart' as http;
import 'dart:convert';

Future<User> fetchUser() async {
  final response = await http.get(Uri.parse('https://jsonplaceholder.typicode.com/users/1'));

  if (response.statusCode == 200) {
    return User.fromJson(jsonDecode(response.body));
  } else {
    throw Exception('Failed to load user');
  }
}
```

✔ `jsonDecode(response.body)` → API 응답을 JSON으로 변환  
✔ `User.fromJson(json)` → JSON을 Dart 객체로 변환  

---

## 5. JSON 리스트 데이터 변환

API 응답이 **배열(List)** 형태라면 리스트 변환이 필요합니다.

### JSON 리스트를 Dart 객체 리스트로 변환

```dart
Future<List<User>> fetchUsers() async {
  final response = await http.get(Uri.parse('https://jsonplaceholder.typicode.com/users'));

  if (response.statusCode == 200) {
    List<dynamic> jsonList = jsonDecode(response.body);
    return jsonList.map((json) => User.fromJson(json)).toList();
  } else {
    throw Exception('Failed to load users');
  }
}
```

✔ `jsonList.map((json) => User.fromJson(json)).toList()` → 리스트 변환  

---

## 6. `json_serializable`을 이용한 자동 변환

`json_serializable` 패키지를 사용하면 **JSON 변환 코드를 자동 생성**할 수 있습니다.

### 패키지 설치

```sh
flutter pub add json_annotation
flutter pub add build_runner json_serializable --dev
```

### 모델 클래스 작성

```dart
import 'package:json_annotation/json_annotation.dart';

part 'user.g.dart';

@JsonSerializable()
class User {
  final int id;
  final String name;
  final String email;

  User({required this.id, required this.name, required this.email});

  factory User.fromJson(Map<String, dynamic> json) => _$UserFromJson(json);

  Map<String, dynamic> toJson() => _$UserToJson(this);
}
```

✔ `@JsonSerializable()` → JSON 변환을 자동화  
✔ `_$UserFromJson(json)` → 자동 생성된 JSON 변환 함수  

### 코드 생성

```sh
flutter pub run build_runner build
```

---

## 7. JSON과 UI 연결 예제

Flutter UI에서 **JSON 데이터를 표시**하는 예제입니다.

```dart
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

class User {
  final int id;
  final String name;
  final String email;

  User({required this.id, required this.name, required this.email});

  factory User.fromJson(Map<String, dynamic> json) {
    return User(
      id: json['id'],
      name: json['name'],
      email: json['email'],
    );
  }
}

class UserScreen extends StatefulWidget {
  @override
  _UserScreenState createState() => _UserScreenState();
}

class _UserScreenState extends State<UserScreen> {
  late Future<User> user;

  @override
  void initState() {
    super.initState();
    user = fetchUser();
  }

  Future<User> fetchUser() async {
    final response = await http.get(Uri.parse('https://jsonplaceholder.typicode.com/users/1'));

    if (response.statusCode == 200) {
      return User.fromJson(jsonDecode(response.body));
    } else {
      throw Exception('Failed to load user');
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text("User Info")),
      body: FutureBuilder<User>(
        future: user,
        builder: (context, snapshot) {
          if (snapshot.connectionState == ConnectionState.waiting) {
            return Center(child: CircularProgressIndicator());
          } else if (snapshot.hasError) {
            return Center(child: Text("Error: ${snapshot.error}"));
          } else if (snapshot.hasData) {
            return Center(
              child: Column(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Text("Name: ${snapshot.data!.name}", style: TextStyle(fontSize: 20)),
                  Text("Email: ${snapshot.data!.email}", style: TextStyle(fontSize: 18)),
                ],
              ),
            );
          } else {
            return Center(child: Text("No Data"));
          }
        },
      ),
    );
  }
}
```

✔ `FutureBuilder<User>` → API 데이터를 UI에 표시  
✔ `CircularProgressIndicator()` → 데이터 로딩 표시  

---

## 🎯 정리

✔ `jsonDecode()` → JSON 문자열을 Dart 객체(Map)로 변환        
✔ `jsonEncode()` → Dart 객체를 JSON 문자열로 변환  
✔ 모델 클래스를 활용하여 JSON 데이터를 쉽게 관리  
✔ `json_serializable` → JSON 자동 변환 가능  

