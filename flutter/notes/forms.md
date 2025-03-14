# 📌 Flutter 폼과 입력 처리

Flutter에서 사용자의 입력을 처리하기 위해 **TextField**, **TextFormField**, **폼 유효성 검사**, **Focus**, **키보드 제어** 등을 사용할 수 있습니다.

---

## 1. `TextField` 기본 사용법
`TextField`는 기본적인 입력 필드입니다.

```dart
TextField(
  decoration: InputDecoration(
    labelText: '이름 입력',
    border: OutlineInputBorder(),
  ),
)
```

✔ `labelText` : 필드 라벨 설정  
✔ `border: OutlineInputBorder()` : 입력 필드에 테두리 추가  

---

## 2. `TextEditingController`를 활용한 입력값 가져오기
사용자가 입력한 값을 가져오려면 **TextEditingController**를 사용합니다.

```dart
import 'package:flutter/material.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: InputScreen(),
    );
  }
}

class InputScreen extends StatefulWidget {
  @override
  _InputScreenState createState() => _InputScreenState();
}

class _InputScreenState extends State<InputScreen> {
  final TextEditingController _controller = TextEditingController();

  void _printInput() {
    print("입력한 값: ${_controller.text}");
  }

  @override
  void dispose() {
    _controller.dispose(); // 메모리 누수 방지를 위해 dispose 필요
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text('TextField 입력 처리')),
      body: Padding(
        padding: EdgeInsets.all(16.0),
        child: Column(
          children: [
            TextField(
              controller: _controller,
              decoration: InputDecoration(labelText: '이름 입력'),
            ),
            SizedBox(height: 20),
            ElevatedButton(
              onPressed: _printInput,
              child: Text('출력하기'),
            ),
          ],
        ),
      ),
    );
  }
}
```

✔ `TextEditingController` 를 사용하여 입력값 가져오기  
✔ 입력값 사용 후 `dispose()`를 호출하여 메모리 누수 방지

---

## 3. `TextFormField`와 폼(Form) 활용

`TextFormField`는 `TextField`와 비슷하지만 **폼(Form)과 함께 사용하여 유효성 검사를 쉽게 처리**할 수 있습니다.

```dart
import 'package:flutter/material.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: FormScreen(),
    );
  }
}

class FormScreen extends StatefulWidget {
  @override
  _FormScreenState createState() => _FormScreenState();
}

class _FormScreenState extends State<FormScreen> {
  final _formKey = GlobalKey<FormState>();
  final TextEditingController _emailController = TextEditingController();

  void _submitForm() {
    if (_formKey.currentState!.validate()) {
      print("유효한 입력: ${_emailController.text}");
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text('폼 입력 및 유효성 검사')),
      body: Padding(
        padding: EdgeInsets.all(16.0),
        child: Form(
          key: _formKey,
          child: Column(
            children: [
              TextFormField(
                controller: _emailController,
                decoration: InputDecoration(labelText: '이메일 입력'),
                validator: (value) {
                  if (value == null || value.isEmpty) {
                    return '이메일을 입력하세요.';
                  } else if (!RegExp(r'^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$').hasMatch(value)) {
                    return '유효한 이메일 주소를 입력하세요.';
                  }
                  return null;
                },
              ),
              SizedBox(height: 20),
              ElevatedButton(
                onPressed: _submitForm,
                child: Text('제출'),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
```

✔ `Form` 과 `GlobalKey<FormState>` 를 사용하여 유효성 검사 가능  
✔ `validator` 속성을 사용하여 입력값 검증  
✔ 버튼 클릭 시 `validate()` 실행하여 입력값 확인

---

## 4. 포커스(Focus)와 키보드 숨기기

입력 필드 간 포커스를 이동하거나, 입력 후 키보드를 숨기려면 **FocusNode**를 사용합니다.

```dart
class FocusExample extends StatefulWidget {
  @override
  _FocusExampleState createState() => _FocusExampleState();
}

class _FocusExampleState extends State<FocusExample> {
  final FocusNode _focusNode1 = FocusNode();
  final FocusNode _focusNode2 = FocusNode();

  void _changeFocus() {
    FocusScope.of(context).requestFocus(_focusNode2);
  }

  void _hideKeyboard() {
    FocusScope.of(context).unfocus(); // 키보드 숨기기
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text('포커스 및 키보드 제어')),
      body: Padding(
        padding: EdgeInsets.all(16.0),
        child: Column(
          children: [
            TextField(
              focusNode: _focusNode1,
              decoration: InputDecoration(labelText: '첫 번째 필드'),
            ),
            SizedBox(height: 10),
            TextField(
              focusNode: _focusNode2,
              decoration: InputDecoration(labelText: '두 번째 필드'),
            ),
            SizedBox(height: 20),
            ElevatedButton(
              onPressed: _changeFocus,
              child: Text('다음 필드로 이동'),
            ),
            ElevatedButton(
              onPressed: _hideKeyboard,
              child: Text('키보드 숨기기'),
            ),
          ],
        ),
      ),
    );
  }
}
```

✔ `FocusNode` 를 사용하여 특정 입력 필드에 포커스 지정 가능  
✔ `FocusScope.of(context).requestFocus(focusNode)` 를 사용하여 포커스 이동  
✔ `FocusScope.of(context).unfocus()` 로 키보드 숨기기 가능  

---

## 5. 비밀번호 입력 필드 (`obscureText`)
비밀번호 입력을 위해 `obscureText: true`를 설정할 수 있습니다.

```dart
TextField(
  obscureText: true,
  decoration: InputDecoration(labelText: '비밀번호 입력'),
)
```

✔`obscureText: true` 를 사용하여 입력한 문자 숨김 처리  

---

## 6. 자동완성 및 키보드 타입 설정

### 6.1 키보드 타입 변경
```dart
TextField(
  keyboardType: TextInputType.emailAddress, // 이메일 입력에 적합한 키보드
  decoration: InputDecoration(labelText: '이메일 입력'),
)
```

### 6.2 자동완성 및 힌트 제공
```dart
TextField(
  autofillHints: [AutofillHints.email], // 자동완성 기능 제공
  decoration: InputDecoration(labelText: '이메일 입력'),
)
```

✔ `keyboardType` 을 설정하여 숫자, 이메일, 전화번호 입력에 최적화 가능  
✔ `autofillHints` 를 사용하여 자동완성 기능 활용 가능  

---

## 🎯 정리

| 기능 | 위젯 / 메서드 |
|------|-------------|
| 기본 입력 필드 | `TextField` |
| 입력값 가져오기 | `TextEditingController.text` |
| 폼 유효성 검사 | `TextFormField` + `Form` |
| 키보드 숨기기 | `FocusScope.of(context).unfocus()` |
| 비밀번호 입력 | `obscureText: true` |
| 자동완성 | `autofillHints: [AutofillHints.email]` |

