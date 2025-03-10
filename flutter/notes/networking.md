# Flutter 네트워크 통신 정리

Flutter에서 네트워크 통신을 수행하려면 **`http` 패키지** 또는 **Dio** 등의 라이브러리를 사용합니다.  
주로 **REST API 요청(GET, POST, PUT, DELETE)**을 통해 데이터를 주고받습니다.

---

## 1. `http` 패키지 설치

Flutter에서는 네트워크 요청을 위해 **`http` 패키지**를 사용합니다.  
우선, `pubspec.yaml` 파일에 **http 패키지**를 추가합니다.

```yaml
dependencies:
  flutter:
    sdk: flutter
  http: ^0.13.0  # 최신 버전 확인 후 적용
```

📌 **`pub get`을 실행하여 패키지를 설치**하세요.

---

## 2. GET 요청 (서버에서 데이터 가져오기)

```dart
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: NetworkScreen(),
    );
  }
}

class NetworkScreen extends StatefulWidget {
  @override
  _NetworkScreenState createState() => _NetworkScreenState();
}

class _NetworkScreenState extends State<NetworkScreen> {
  String _data = "데이터 없음";

  Future<void> fetchData() async {
    final response = await http.get(Uri.parse('https://jsonplaceholder.typicode.com/posts/1'));

    if (response.statusCode == 200) {
      var jsonData = jsonDecode(response.body);
      setState(() {
        _data = jsonData['title'];
      });
    } else {
      throw Exception('데이터를 불러오는 데 실패했습니다.');
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text('GET 요청 예제')),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text(_data, style: TextStyle(fontSize: 18)),
            ElevatedButton(
              onPressed: fetchData,
              child: Text('데이터 가져오기'),
            ),
          ],
        ),
      ),
    );
  }
}
```

✅ **`http.get()`을 사용하여 데이터를 가져옴**  
✅ **`jsonDecode()`로 JSON 데이터를 파싱하여 사용**  

---

## 3. POST 요청 (서버에 데이터 보내기)

```dart
Future<void> sendData() async {
  final response = await http.post(
    Uri.parse('https://jsonplaceholder.typicode.com/posts'),
    headers: {"Content-Type": "application/json"},
    body: jsonEncode({
      "title": "Flutter",
      "body": "Flutter 네트워크 요청",
      "userId": 1
    }),
  );

  if (response.statusCode == 201) {
    print("데이터 전송 성공: ${response.body}");
  } else {
    print("데이터 전송 실패");
  }
}
```

✅ **`http.post()`를 사용하여 데이터를 서버로 전송**  
✅ **`headers`에서 `Content-Type`을 `application/json`으로 설정**  
✅ **`jsonEncode()`를 사용하여 데이터를 JSON 형식으로 변환**  

---

## 4. PUT 요청 (데이터 수정)

```dart
Future<void> updateData() async {
  final response = await http.put(
    Uri.parse('https://jsonplaceholder.typicode.com/posts/1'),
    headers: {"Content-Type": "application/json"},
    body: jsonEncode({
      "id": 1,
      "title": "Updated Title",
      "body": "Updated Content",
      "userId": 1
    }),
  );

  if (response.statusCode == 200) {
    print("데이터 수정 성공: ${response.body}");
  } else {
    print("데이터 수정 실패");
  }
}
```

✅ **`http.put()`을 사용하여 데이터를 수정**  

---

## 5. DELETE 요청 (데이터 삭제)

```dart
Future<void> deleteData() async {
  final response = await http.delete(
    Uri.parse('https://jsonplaceholder.typicode.com/posts/1'),
  );

  if (response.statusCode == 200) {
    print("데이터 삭제 성공");
  } else {
    print("데이터 삭제 실패");
  }
}
```

✅ **`http.delete()`를 사용하여 데이터 삭제**  

---

## 6. Dio 패키지를 활용한 네트워크 요청

Dio는 더 강력한 기능을 제공하는 HTTP 클라이언트입니다.

### 6.1 `dio` 패키지 설치
```yaml
dependencies:
  dio: ^5.0.0
```

### 6.2 Dio를 사용한 GET 요청
```dart
import 'package:dio/dio.dart';

final dio = Dio();

Future<void> fetchData() async {
  try {
    final response = await dio.get('https://jsonplaceholder.typicode.com/posts/1');
    print("데이터 가져오기 성공: ${response.data}");
  } catch (e) {
    print("에러 발생: $e");
  }
}
```

✅ **더 간결한 코드로 네트워크 요청 가능**  
✅ **자동 JSON 파싱 지원**  
✅ **인터셉터(Interceptor) 기능 제공**  

---

## 7. JSON 데이터 모델로 변환하기

네트워크에서 받은 데이터를 **Dart 객체로 변환**하면 관리하기 쉽습니다.

### 7.1 JSON 데이터 모델 클래스 생성
```dart
class Post {
  final int userId;
  final int id;
  final String title;
  final String body;

  Post({required this.userId, required this.id, required this.title, required this.body});

  factory Post.fromJson(Map<String, dynamic> json) {
    return Post(
      userId: json['userId'],
      id: json['id'],
      title: json['title'],
      body: json['body'],
    );
  }
}
```

### 7.2 JSON을 모델로 변환하는 GET 요청
```dart
Future<Post> fetchPost() async {
  final response = await http.get(Uri.parse('https://jsonplaceholder.typicode.com/posts/1'));

  if (response.statusCode == 200) {
    return Post.fromJson(jsonDecode(response.body));
  } else {
    throw Exception('데이터를 불러오는 데 실패했습니다.');
  }
}
```

✅ **JSON을 Dart 모델로 변환하여 가독성을 높임**  
✅ **버그 발생 가능성을 줄이고, 유지보수 용이**  

---

## 8. 네트워크 상태 확인하기

```dart
import 'package:connectivity_plus/connectivity_plus.dart';

Future<void> checkInternetConnection() async {
  var connectivityResult = await Connectivity().checkConnectivity();
  if (connectivityResult == ConnectivityResult.mobile) {
    print("모바일 데이터 연결됨");
  } else if (connectivityResult == ConnectivityResult.wifi) {
    print("Wi-Fi 연결됨");
  } else {
    print("인터넷 연결 없음");
  }
}
```

✅ **오프라인 상태에서도 앱이 동작하도록 처리 가능**  

---

## 9. 결론

| 요청 방식 | 설명 | 메서드 |
|-----------|----------------|-----------|
| **GET** | 서버에서 데이터 가져오기 | `http.get()` |
| **POST** | 서버에 데이터 전송 | `http.post()` |
| **PUT** | 기존 데이터 수정 | `http.put()` |
| **DELETE** | 데이터 삭제 | `http.delete()` |

---

📌 **Flutter 네트워크 요청을 효율적으로 처리하려면**  
✔ `http` 패키지를 사용하여 기본 요청 처리  
✔ `Dio`를 사용하여 고급 네트워크 요청 수행  
✔ `connectivity_plus`를 이용해 인터넷 연결 상태 확인  
✔ JSON 데이터를 **Dart 모델 객체로 변환**하여 관리  
