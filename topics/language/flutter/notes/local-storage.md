# 🌍 Flutter 로컬 데이터 저장

Flutter에서 앱 내에서 데이터를 저장하고 불러오는 다양한 방식들을 알아봅니다.

---

## 1️⃣ 로컬 데이터 저장 방식
Flutter에서 로컬 데이터를 저장하는 방법은 여러 가지가 있습니다.

| 저장 방식 | 사용 목적 | 장점 | 단점 |
|----------|----------|------|------|
| `SharedPreferences` | 간단한 키-값 저장 (설정, 토큰 등) | 사용이 간단함 | 구조화된 데이터 저장 어려움 |
| `Hive` | 경량 NoSQL 데이터베이스 | 빠른 속도, 직렬화 필요 없음 | 관계형 데이터 부족 |
| `Sqflite` | SQLite 기반 데이터베이스 | SQL 지원, 관계형 데이터 저장 가능 | ORM 지원 부족, 코드가 길어짐 |
| `Isar` | 빠른 NoSQL DB | 성능이 뛰어남 | 생태계가 작음 |
| `File` | 파일 입출력 | JSON, 텍스트 저장 가능 | 직접 관리 필요 |

---

## 2️⃣ `SharedPreferences` (간단한 키-값 저장)

#### 1. 패키지 설치

```sh
flutter pub add shared_preferences
```

#### 2. 데이터 저장

```dart
import 'package:shared_preferences/shared_preferences.dart';

Future<void> saveData() async {
  final prefs = await SharedPreferences.getInstance();
  await prefs.setString('username', 'flutter_user');
}
```

✔ `SharedPreferences.getInstance()` → 인스턴스 생성  
✔ `prefs.setString('key', value)` → 값 저장  

#### 3. 데이터 불러오기

```dart
Future<void> loadData() async {
  final prefs = await SharedPreferences.getInstance();
  String? username = prefs.getString('username');
  print(username);
}
```

✔ `prefs.getString('key')` → 저장된 값 가져오기  

#### 4. 데이터 삭제

```dart
Future<void> removeData() async {
  final prefs = await SharedPreferences.getInstance();
  await prefs.remove('username');
}
```

✔ `prefs.remove('key')` → 특정 키 데이터 삭제  

---

## 3️⃣ `Hive` (빠른 NoSQL DB)

#### 1. 패키지 설치

```sh
flutter pub add hive
flutter pub add hive_flutter
```

#### 2. 초기화 및 데이터 저장

```dart
import 'package:hive_flutter/hive_flutter.dart';

Future<void> initHive() async {
  await Hive.initFlutter();
  var box = await Hive.openBox('myBox');
  box.put('username', 'flutter_user');
}
```

✔ `Hive.initFlutter()` → Hive 초기화  
✔ `Hive.openBox('boxName')` → 데이터 저장 박스 열기  
✔ `box.put('key', value)` → 데이터 저장  

#### 3. 데이터 불러오기

```dart
Future<void> loadHiveData() async {
  var box = await Hive.openBox('myBox');
  String? username = box.get('username');
  print(username);
}
```
✔ `box.get('key')` → 값 가져오기  

#### 4. 데이터 삭제

```dart
Future<void> deleteHiveData() async {
  var box = await Hive.openBox('myBox');
  await box.delete('username');
}
```

✔ `box.delete('key')` → 특정 데이터 삭제  

---

## 4️⃣ `Sqflite` (SQLite 데이터베이스)

#### 1. 패키지 설치

```sh
flutter pub add sqflite path
```

#### 2. 데이터베이스 생성 및 테이블 만들기

```dart
import 'package:sqflite/sqflite.dart';
import 'package:path/path.dart';

Future<Database> openDatabaseConnection() async {
  return openDatabase(
    join(await getDatabasesPath(), 'my_database.db'),
    onCreate: (db, version) {
      return db.execute(
        "CREATE TABLE users(id INTEGER PRIMARY KEY, name TEXT)",
      );
    },
    version: 1,
  );
}
```

✔ `getDatabasesPath()` → 데이터베이스 경로 설정  
✔ `db.execute("CREATE TABLE ...")` → 테이블 생성  

#### 3. 데이터 삽입

```dart
Future<void> insertUser(Database db) async {
  await db.insert(
    'users',
    {'id': 1, 'name': 'Flutter User'},
    conflictAlgorithm: ConflictAlgorithm.replace,
  );
}
```

✔ `db.insert('tableName', data)` → 데이터 삽입  

#### 4. 데이터 조회

```dart
Future<List<Map<String, dynamic>>> getUsers(Database db) async {
  return await db.query('users');
}
```

✔ `db.query('tableName')` → 데이터 조회  

#### 5. 데이터 삭제

```dart
Future<void> deleteUser(Database db) async {
  await db.delete('users', where: "id = ?", whereArgs: [1]);
}
```

✔ `db.delete('tableName', where: "...")` → 특정 데이터 삭제  

---

## 5️⃣ 파일을 이용한 JSON 데이터 저장

#### 1. 패키지 설치

```sh
flutter pub add path_provider
```

#### 2. JSON 데이터 저장

```dart
import 'dart:io';
import 'dart:convert';
import 'package:path_provider/path_provider.dart';

Future<File> getFile() async {
  final directory = await getApplicationDocumentsDirectory();
  return File('${directory.path}/data.json');
}

Future<void> saveJson(Map<String, dynamic> data) async {
  final file = await getFile();
  file.writeAsString(jsonEncode(data));
}
```

✔ `getApplicationDocumentsDirectory()` → 저장 경로 가져오기  
✔ `file.writeAsString(jsonEncode(data))` → JSON 데이터 저장  

#### 3. JSON 데이터 불러오기

```dart
Future<Map<String, dynamic>> loadJson() async {
  try {
    final file = await getFile();
    String contents = await file.readAsString();
    return jsonDecode(contents);
  } catch (e) {
    return {};
  }
}
```

✔ `file.readAsString()` → JSON 파일 읽기  
✔ `jsonDecode(contents)` → JSON을 Dart 객체로 변환  

---

## 6️⃣ `Isar` (고속 NoSQL 데이터베이스)

#### 1. 패키지 설치

```sh
flutter pub add isar isar_flutter_libs
flutter pub add build_runner isar_generator --dev
```

#### 2. 모델 클래스 생성

```dart
import 'package:isar/isar.dart';

@Collection()
class User {
  Id id = Isar.autoIncrement;
  late String name;
}
```

✔ `@Collection()` → Isar 데이터 모델 생성  
✔ `Isar.autoIncrement` → 자동 증가 ID  

#### 3. 데이터 저장 및 조회

```dart
Future<void> saveUser(Isar isar) async {
  final user = User()..name = "Flutter User";
  await isar.writeTxn(() async {
    await isar.users.put(user);
  });
}

Future<List<User>> getUsers(Isar isar) async {
  return isar.users.where().findAll();
}
```

✔ `isar.writeTxn(() async { ... })` → 데이터 변경 트랜잭션  
✔ `isar.users.put(user)` → 데이터 저장  

---

## 🎯 정리

✔ `SharedPreferences` → 키-값 저장 (설정, 토큰)  
✔ `Hive` → 빠른 NoSQL DB (직렬화 불필요)  
✔ `Sqflite` → 관계형 데이터 저장 (SQL 지원)  
✔ `Isar` → 고속 NoSQL (ORM 지원)  
✔ `File` → JSON 파일 저장  
