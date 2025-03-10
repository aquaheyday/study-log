# Flutter 파일 입출력 (File I/O)

Flutter에서 **파일을 읽고 쓰는 방법**은 `dart:io`의 `File` 클래스를 사용합니다.  
주로 `path_provider` 패키지를 활용하여 앱의 저장 디렉터리를 찾습니다.

---

## 1. `path_provider` 패키지 설치
Flutter에서 파일을 저장하려면 **앱의 저장소 경로**를 가져와야 합니다.  
이를 위해 `path_provider` 패키지를 설치해야 합니다.

```yaml
dependencies:
  flutter:
    sdk: flutter
  path_provider: ^2.0.11  # 최신 버전 확인 후 적용
```

📌 `pub get`을 실행하여 패키지를 설치하세요.

---

## 2. 파일 저장 경로 가져오기
`path_provider`를 사용하여 **앱의 저장 디렉터리를 가져올 수 있습니다.**

```dart
import 'package:path_provider/path_provider.dart';
import 'dart:io';

Future<String> getFilePath() async {
  final directory = await getApplicationDocumentsDirectory();
  return '${directory.path}/my_file.txt';
}
```

✅ `getApplicationDocumentsDirectory()` : 앱의 문서 저장소 (iOS, Android에서 사용 가능)  
✅ `getTemporaryDirectory()` : 임시 파일 저장소  

---

## 3. 파일 쓰기 (저장)

파일을 생성하고 데이터를 저장하려면 `File.writeAsString()`을 사용합니다.

```dart
Future<void> writeToFile(String content) async {
  final path = await getFilePath();
  final file = File(path);
  await file.writeAsString(content);
  print("파일 저장 완료: $path");
}
```

✅ `writeAsString(content)` : 문자열을 파일에 저장  
✅ `writeAsBytes(bytes)` : 바이너리 데이터를 저장  

---

## 4. 파일 읽기

저장된 파일을 읽으려면 `File.readAsString()`을 사용합니다.

```dart
Future<String> readFromFile() async {
  try {
    final path = await getFilePath();
    final file = File(path);
    return await file.readAsString();
  } catch (e) {
    return "파일을 찾을 수 없습니다.";
  }
}
```

✅ `readAsString()` : 파일의 내용을 문자열로 읽음  
✅ `readAsBytes()` : 파일의 내용을 바이너리 데이터로 읽음  

---

## 5. 파일 삭제

파일을 삭제하려면 `File.delete()`를 사용합니다.

```dart
Future<void> deleteFile() async {
  final path = await getFilePath();
  final file = File(path);

  if (await file.exists()) {
    await file.delete();
    print("파일 삭제 완료");
  } else {
    print("파일이 존재하지 않습니다.");
  }
}
```

✅ `file.exists()` : 파일 존재 여부 확인  
✅ `file.delete()` : 파일 삭제  

---

## 6. 파일 입출력 예제 (UI 포함)

```dart
import 'package:flutter/material.dart';
import 'package:path_provider/path_provider.dart';
import 'dart:io';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: FileStorageScreen(),
    );
  }
}

class FileStorageScreen extends StatefulWidget {
  @override
  _FileStorageScreenState createState() => _FileStorageScreenState();
}

class _FileStorageScreenState extends State<FileStorageScreen> {
  final TextEditingController _controller = TextEditingController();
  String _fileContent = "파일을 읽어보세요!";

  Future<String> getFilePath() async {
    final directory = await getApplicationDocumentsDirectory();
    return '${directory.path}/my_file.txt';
  }

  Future<void> writeToFile(String content) async {
    final path = await getFilePath();
    final file = File(path);
    await file.writeAsString(content);
    print("파일 저장 완료: $path");
  }

  Future<void> readFromFile() async {
    try {
      final path = await getFilePath();
      final file = File(path);
      String content = await file.readAsString();
      setState(() {
        _fileContent = content;
      });
    } catch (e) {
      setState(() {
        _fileContent = "파일을 찾을 수 없습니다.";
      });
    }
  }

  Future<void> deleteFile() async {
    final path = await getFilePath();
    final file = File(path);

    if (await file.exists()) {
      await file.delete();
      setState(() {
        _fileContent = "파일 삭제됨";
      });
    } else {
      setState(() {
        _fileContent = "파일이 존재하지 않습니다.";
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text('파일 입출력 예제')),
      body: Padding(
        padding: EdgeInsets.all(16.0),
        child: Column(
          children: [
            TextField(
              controller: _controller,
              decoration: InputDecoration(labelText: '저장할 내용 입력'),
            ),
            SizedBox(height: 10),
            ElevatedButton(
              onPressed: () => writeToFile(_controller.text),
              child: Text('파일 저장'),
            ),
            ElevatedButton(
              onPressed: readFromFile,
              child: Text('파일 읽기'),
            ),
            ElevatedButton(
              onPressed: deleteFile,
              child: Text('파일 삭제'),
            ),
            SizedBox(height: 20),
            Text("파일 내용: $_fileContent"),
          ],
        ),
      ),
    );
  }
}
```

✅ **입력값을 저장, 읽기, 삭제할 수 있는 예제**  
✅ **파일 경로는 `getApplicationDocumentsDirectory()`에서 가져옴**  

---

## 7. 바이너리 데이터 (이미지 저장)

이미지와 같은 바이너리 데이터를 저장하려면 `writeAsBytes()`와 `readAsBytes()`를 사용합니다.

```dart
Future<void> saveImage(Uint8List imageBytes) async {
  final path = await getFilePath();
  final file = File('$path/image.png');
  await file.writeAsBytes(imageBytes);
}
```

```dart
Future<Uint8List?> loadImage() async {
  final path = await getFilePath();
  final file = File('$path/image.png');

  if (await file.exists()) {
    return await file.readAsBytes();
  } else {
    return null;
  }
}
```

✅ `writeAsBytes()` : 바이너리 데이터 저장  
✅ `readAsBytes()` : 바이너리 데이터 읽기  

---

## 8. JSON 파일 저장 및 읽기

JSON 데이터를 파일로 저장하고 읽을 수도 있습니다.

### 8.1 JSON 데이터 저장
```dart
import 'dart:convert';

Future<void> saveJson(Map<String, dynamic> data) async {
  final path = await getFilePath();
  final file = File('$path/data.json');
  await file.writeAsString(jsonEncode(data));
}
```

### 8.2 JSON 데이터 읽기
```dart
Future<Map<String, dynamic>?> readJson() async {
  final path = await getFilePath();
  final file = File('$path/data.json');

  if (await file.exists()) {
    return jsonDecode(await file.readAsString());
  } else {
    return null;
  }
}
```

✅ `jsonEncode()` : Map을 JSON 문자열로 변환  
✅ `jsonDecode()` : JSON 문자열을 Map으로 변환  

---

## 9. 결론

| 기능 | 메서드 |
|------|------------|
| 파일 저장 | `File.writeAsString()` |
| 파일 읽기 | `File.readAsString()` |
| 파일 삭제 | `File.delete()` |
| 바이너리 저장 | `File.writeAsBytes()` |
| JSON 저장 | `jsonEncode() + writeAsString()` |
| JSON 읽기 | `readAsString() + jsonDecode()` |
