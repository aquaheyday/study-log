# 🌍 Flutter HTTP 통신

Flutter에서 API와 통신하기 위해 `http` 및 `dio` 패키지를 활용하여 HTTP 요청을 수행할 수 있습니다. 두 패키지를 이용한 API 호출 방법을 정리합니다.

---

## 1️⃣ 패키지 설치
Flutter 프로젝트에서 `http` 및 `dio` 패키지를 추가해야 합니다.

```sh
flutter pub add http dio
```

또는 `pubspec.yaml`에 직접 추가:

```yaml
dependencies:
  flutter:
    sdk: flutter
  http: ^0.13.4  # 최신 버전 확인 후 변경 가능
  dio: ^5.3.2    # 최신 버전 확인 후 변경 가능
```

그런 다음 패키지를 설치합니다.

```sh
flutter pub get
```

---

## 2️⃣ `http` 패키지를 이용한 API 호출

### 1) GET 요청

```dart
import 'package:http/http.dart' as http;
import 'dart:convert';

Future<void> fetchData() async {
  final response = await http.get(Uri.parse('https://jsonplaceholder.typicode.com/posts'));

  if (response.statusCode == 200) {
    List data = jsonDecode(response.body);
    print(data);
  } else {
    throw Exception('Failed to load data');
  }
}
```
✔ `http.get()` → GET 요청 실행  
✔ `jsonDecode(response.body)` → JSON 응답을 Dart 객체로 변환  
✔ `statusCode == 200` → 정상 응답 여부 확인  

---

### 2) POST 요청

```dart
Future<void> sendData() async {
  final response = await http.post(
    Uri.parse('https://jsonplaceholder.typicode.com/posts'),
    headers: {'Content-Type': 'application/json'},
    body: jsonEncode({
      'title': 'Flutter HTTP',
      'body': 'HTTP 통신 예제',
      'userId': 1,
    }),
  );

  if (response.statusCode == 201) {
    print('Data sent successfully');
  } else {
    throw Exception('Failed to send data');
  }
}
```

✔ `http.post()` → POST 요청 실행  
✔ `headers` → 요청 헤더 설정 (JSON 데이터 전송)  
✔ `body: jsonEncode({...})` → JSON 데이터 변환 후 전송  

---

### 3) PUT 요청 (데이터 수정)

```dart
Future<void> updateData() async {
  final response = await http.put(
    Uri.parse('https://jsonplaceholder.typicode.com/posts/1'),
    headers: {'Content-Type': 'application/json'},
    body: jsonEncode({
      'id': 1,
      'title': 'Updated Title',
      'body': 'Updated Content',
      'userId': 1,
    }),
  );

  if (response.statusCode == 200) {
    print('Data updated successfully');
  } else {
    throw Exception('Failed to update data');
  }
}
```

✔ `http.put()` → 데이터를 업데이트하는 요청  

---

### 4) DELETE 요청 (데이터 삭제)

```dart
Future<void> deleteData() async {
  final response = await http.delete(Uri.parse('https://jsonplaceholder.typicode.com/posts/1'));

  if (response.statusCode == 200) {
    print('Data deleted successfully');
  } else {
    throw Exception('Failed to delete data');
  }
}
```

✔ `http.delete()` → 데이터 삭제 요청 실행  

---

## 3️⃣ `dio` 패키지를 이용한 API 호출

### 1) Dio 설정

```dart
import 'package:dio/dio.dart';

final Dio dio = Dio(BaseOptions(
  baseUrl: 'https://jsonplaceholder.typicode.com/',
  connectTimeout: Duration(seconds: 5),
  receiveTimeout: Duration(seconds: 3),
  headers: {'Content-Type': 'application/json'},
));
```

✔ `baseUrl` → API 기본 URL 설정  
✔ `connectTimeout`, `receiveTimeout` → 연결 및 응답 시간 설정  
✔ `headers` → 기본 요청 헤더 설정  

---

### 2) GET 요청

```dart
Future<void> fetchDataWithDio() async {
  try {
    Response response = await dio.get('/posts');
    print(response.data);
  } catch (e) {
    print('Error: $e');
  }
}
```

✔ `dio.get()` → GET 요청 실행  
✔ `response.data` → JSON 응답 데이터 확인  

---

### 3) POST 요청

```dart
Future<void> sendDataWithDio() async {
  try {
    Response response = await dio.post(
      '/posts',
      data: {
        'title': 'Flutter Dio',
        'body': 'Dio 패키지를 이용한 HTTP 통신 예제',
        'userId': 1,
      },
    );
    print(response.data);
  } catch (e) {
    print('Error: $e');
  }
}
```

✔ `dio.post()` → POST 요청 실행  
✔ `data` → 전송할 JSON 데이터  

---

### 4) PUT 요청 (데이터 수정)

```dart
Future<void> updateDataWithDio() async {
  try {
    Response response = await dio.put(
      '/posts/1',
      data: {
        'id': 1,
        'title': 'Updated with Dio',
        'body': 'Dio를 이용한 데이터 수정',
        'userId': 1,
      },
    );
    print(response.data);
  } catch (e) {
    print('Error: $e');
  }
}
```

✔ `dio.put()` → 데이터 수정 요청  

---

### 5) DELETE 요청 (데이터 삭제)

```dart
Future<void> deleteDataWithDio() async {
  try {
    Response response = await dio.delete('/posts/1');
    print(response.statusCode == 200 ? 'Deleted' : 'Failed');
  } catch (e) {
    print('Error: $e');
  }
}
```

✔ `dio.delete()` → 데이터 삭제 요청 실행  

---

## 4️⃣ `http` vs `dio` 비교

| 기능  | `http` 패키지 | `dio` 패키지 |
|-------|-------------|-------------|
| 요청 방법 | `http.get()` 등 | `dio.get()` 등 |
| JSON 변환 | `jsonDecode(response.body)` 필요 | 자동 변환 (`response.data`) |
| 에러 처리 | try-catch 사용 | 더 자세한 예외 처리 지원 |
| 헤더 설정 | `headers` 매개변수 필요 | `BaseOptions`에서 설정 가능 |
| 인터셉터 | X | O (요청 및 응답 가로채기 가능) |
| Multipart 지원 | O | O (더 간편한 파일 업로드) |

✔ `http` → 기본적인 API 호출이 필요할 때 사용  
✔ `dio` → 더 강력한 기능이 필요한 경우 사용 (인터셉터, 자동 JSON 변환 등)  

---

## 🎯 정리

✔ `http`와 `dio` 패키지를 사용하여 API와 통신 가능  
✔ `GET`, `POST`, `PUT`, `DELETE` 요청 수행 방법 이해  
✔ `dio`는 `http`보다 더 강력한 기능 제공 (인터셉터, 자동 JSON 변환)  
✔ 프로젝트 요구사항에 따라 적절한 패키지 선택 가능  
