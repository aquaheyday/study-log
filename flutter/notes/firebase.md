# 🌍 Firebase 연동

Flutter에서 Firebase를 연동하는 방법을 정리합니다.  
Firebase는 백엔드 서비스로, 인증, 데이터베이스, 스토리지 등을 제공합니다.

---

## 1. Firebase 프로젝트 설정

### Firebase 콘솔에서 프로젝트 생성
1. [Firebase Console](https://console.firebase.google.com/)에 접속  
2. **새 프로젝트 만들기** 클릭  
3. 프로젝트 이름 입력 후 **계속**  
4. Google 애널리틱스 활성화 여부 선택 후 프로젝트 생성  

---

## 2. Flutter 프로젝트에 Firebase 추가

### Firebase CLI 설치

```sh
npm install -g firebase-tools
```

### Firebase 프로젝트에 Flutter 앱 추가
1. Firebase 콘솔에서 **iOS/Android 앱 추가**  
2. 패키지 이름 입력 (`com.example.myapp`)  
3. **Google 서비스 파일 다운로드**
   - `google-services.json` (Android → `android/app/` 폴더에 추가)
   - `GoogleService-Info.plist` (iOS → `ios/Runner/` 폴더에 추가)
4. 앱 등록 완료 후 Firebase 초기화  

---

## 3. Firebase 패키지 설치

```sh
flutter pub add firebase_core
```

그리고 필요한 Firebase 서비스 패키지를 추가합니다.

| 서비스 | 패키지 |
|--------|---------|
| 인증(Authentication) | `firebase_auth` |
| Firestore DB | `cloud_firestore` |
| 실시간 DB | `firebase_database` |
| 스토리지 | `firebase_storage` |
| 메시징 | `firebase_messaging` |

예제:

```sh
flutter pub add firebase_auth cloud_firestore
```

---

## 4. Firebase 초기화

### `main.dart`에서 Firebase 초기화

```dart
import 'package:firebase_core/firebase_core.dart';
import 'package:flutter/material.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await Firebase.initializeApp();
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(title: Text("Firebase Example")),
        body: Center(child: Text("Firebase Initialized")),
      ),
    );
  }
}
```

✔ `WidgetsFlutterBinding.ensureInitialized();` → Firebase 초기화 전에 Flutter 바인딩  
✔ `await Firebase.initializeApp();` → Firebase 앱 초기화  

---

## 5. Firebase Authentication (로그인)

### 패키지 설치

```sh
flutter pub add firebase_auth
```

### 이메일 로그인 구현

```dart
import 'package:firebase_auth/firebase_auth.dart';

Future<void> signIn(String email, String password) async {
  try {
    UserCredential userCredential = await FirebaseAuth.instance.signInWithEmailAndPassword(
      email: email,
      password: password,
    );
    print("User signed in: ${userCredential.user?.email}");
  } catch (e) {
    print("Sign in failed: $e");
  }
}
```

✔ `signInWithEmailAndPassword()` → 이메일/비밀번호 로그인  

---

### 회원가입 구현

```dart
Future<void> signUp(String email, String password) async {
  try {
    UserCredential userCredential = await FirebaseAuth.instance.createUserWithEmailAndPassword(
      email: email,
      password: password,
    );
    print("User registered: ${userCredential.user?.email}");
  } catch (e) {
    print("Sign up failed: $e");
  }
}
```

✔ `createUserWithEmailAndPassword()` → 이메일/비밀번호 회원가입  

---

### 로그아웃

```dart
Future<void> signOut() async {
  await FirebaseAuth.instance.signOut();
  print("User signed out");
}
```

✔ `signOut()` → 현재 로그인한 사용자 로그아웃  

---

## 6. Firestore (NoSQL 데이터베이스)

### 패키지 설치

```sh
flutter pub add cloud_firestore
```

### 데이터 추가

```dart
import 'package:cloud_firestore/cloud_firestore.dart';

Future<void> addUser(String name, int age) async {
  await FirebaseFirestore.instance.collection('users').add({
    'name': name,
    'age': age,
    'createdAt': FieldValue.serverTimestamp(),
  });
}
```

✔ `.collection('users').add({...})` → Firestore에 새 문서 추가  

---

### 데이터 읽기

```dart
Future<void> fetchUsers() async {
  QuerySnapshot snapshot = await FirebaseFirestore.instance.collection('users').get();
  for (var doc in snapshot.docs) {
    print("${doc.id} => ${doc.data()}");
  }
}
```

✔ `.collection('users').get()` → Firestore에서 모든 문서 가져오기  

---

### 데이터 업데이트

```dart
Future<void> updateUser(String userId, String newName) async {
  await FirebaseFirestore.instance.collection('users').doc(userId).update({
    'name': newName,
  });
}
```

✔ `.doc(userId).update({...})` → 특정 문서 업데이트  

---

### 데이터 삭제

```dart
Future<void> deleteUser(String userId) async {
  await FirebaseFirestore.instance.collection('users').doc(userId).delete();
}
```

✔ `.doc(userId).delete()` → 특정 문서 삭제  

---

## 7. Firebase Storage (파일 저장)

### 패키지 설치

```sh
flutter pub add firebase_storage
```

### 파일 업로드

```dart
import 'dart:io';
import 'package:firebase_storage/firebase_storage.dart';

Future<void> uploadFile(File file) async {
  try {
    await FirebaseStorage.instance.ref('uploads/myfile.jpg').putFile(file);
    print("File uploaded");
  } catch (e) {
    print("Upload failed: $e");
  }
}
```

✔ `.ref('path/to/file').putFile(file)` → Firebase Storage에 파일 업로드  

---

### 파일 다운로드 URL 가져오기

```dart
Future<String> getDownloadUrl(String path) async {
  return await FirebaseStorage.instance.ref(path).getDownloadURL();
}
```

✔ `.getDownloadURL()` → 업로드된 파일의 다운로드 URL 가져오기  

---

## 🎯 정리
 
✔ `firebase_auth` → 사용자 인증 (이메일 로그인)  
✔ `cloud_firestore` → 실시간 데이터 저장/조회  
✔ `firebase_storage` → 파일 업로드 및 다운로드  
