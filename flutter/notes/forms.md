# 🔲 폼과 입력 필드

Flutter에서 **폼(Form)과 입력 필드(Input Fields)** 는 **사용자의 입력을 처리하는 핵심 요소**입니다.  
사용자가 텍스트를 입력하거나, 버튼을 눌러 데이터를 제출하는 기능을 구현할 수 있습니다.

---

## 1. 폼(Form)이란?

- `Form` 위젯을 사용하여 여러 개의 입력 필드를 그룹화할 수 있습니다.
- `TextFormField`를 사용하면 **입력값 검증(Validation)** 이 가능합니다.
- `GlobalKey<FormState>`를 활용하여 **폼의 상태를 관리**할 수 있습니다.

---

## 2. 기본 입력 필드 (TextField)

Flutter에서 가장 기본적인 입력 필드는 `TextField`입니다.

```dart
TextField(
  decoration: InputDecoration(
    labelText: "이름",
    hintText: "이름을 입력하세요",
    border: OutlineInputBorder(), // 테두리 추가
  ),
)
```

✔ **TextField 주요 속성**
| 속성 | 설명 |
|------|------|
| `controller` | 입력된 값을 제어 |
| `decoration` | 입력 필드 디자인 변경 |
| `keyboardType` | 숫자 키패드 등 입력 방식 지정 |
| `obscureText` | 비밀번호 입력 (●●●) |
| `maxLength` | 최대 입력 글자 수 제한 |

---

## 3. `TextEditingController`를 활용한 입력값 가져오기

`TextEditingController`를 사용하면 입력된 값을 가져올 수 있습니다.

```dart
final TextEditingController _controller = TextEditingController();

void _printValue() {
  print("입력한 값: ${_controller.text}");
}

TextField(
  controller: _controller,
  decoration: InputDecoration(labelText: "이름"),
);

ElevatedButton(
  onPressed: _printValue,
  child: Text("출력"),
);
```

✔ `_controller.text` 를 사용하면 입력값을 가져올 수 있음.

---

## 4. 폼 (`Form`)과 입력값 검증 (`TextFormField`)

`TextFormField`는 `Form` 위젯과 함께 사용하여 **입력 검증(Validation)**이 가능합니다.

```dart
final _formKey = GlobalKey<FormState>();

Form(
  key: _formKey,
  child: Column(
    children: [
      TextFormField(
        decoration: InputDecoration(labelText: "이메일"),
        validator: (value) {
          if (value == null || value.isEmpty) {
            return "이메일을 입력하세요.";
          }
          return null;
        },
      ),
      ElevatedButton(
        onPressed: () {
          if (_formKey.currentState!.validate()) {
            print("폼이 유효합니다.");
          }
        },
        child: Text("제출"),
      ),
    ],
  ),
)
```

✔ **폼 검증(Validation)**
- `validator` → 입력값을 확인하고, 오류 메시지를 반환할 수 있음.
- `_formKey.currentState!.validate()` → 폼이 유효한지 검사.

---

## 5. 다양한 입력 필드 종류

### 5-1. 숫자 입력 (`keyboardType`)
```dart
TextField(
  keyboardType: TextInputType.number,
  decoration: InputDecoration(labelText: "나이"),
)
```

### 5-2. 비밀번호 입력 (`obscureText`)
```dart
TextField(
  obscureText: true,
  decoration: InputDecoration(labelText: "비밀번호"),
)
```

### 5-3. 다중 줄 입력 (`maxLines`)
```dart
TextField(
  maxLines: 3, // 여러 줄 입력 가능
  decoration: InputDecoration(labelText: "설명"),
)
```

---

## 6. 입력 필드 디자인 커스텀

입력 필드를 스타일링하려면 `InputDecoration`을 활용합니다.

```dart
TextField(
  decoration: InputDecoration(
    labelText: "이름",
    hintText: "이름을 입력하세요",
    prefixIcon: Icon(Icons.person),  // 왼쪽 아이콘 추가
    suffixIcon: Icon(Icons.clear),   // 오른쪽 아이콘 추가
    border: OutlineInputBorder(borderRadius: BorderRadius.circular(10)), // 둥근 테두리
    filled: true, 
    fillColor: Colors.grey[200], // 배경색
  ),
)
```

✔ **주요 속성**
| 속성 | 설명 |
|------|------|
| `labelText` | 필드 위에 표시되는 라벨 |
| `hintText` | 입력 전 힌트 텍스트 |
| `prefixIcon` | 왼쪽 아이콘 추가 |
| `suffixIcon` | 오른쪽 아이콘 추가 |
| `border` | 테두리 스타일 변경 |
| `filled` & `fillColor` | 배경색 지정 |

---

## 7. 체크박스, 라디오 버튼, 스위치

### 7-1. 체크박스 (`Checkbox`)
```dart
bool _isChecked = false;

Checkbox(
  value: _isChecked,
  onChanged: (bool? value) {
    setState(() {
      _isChecked = value!;
    });
  },
)
```

---

### 7-2. 라디오 버튼 (`Radio`)
```dart
String _selectedGender = "남성";

Column(
  children: [
    RadioListTile(
      title: Text("남성"),
      value: "남성",
      groupValue: _selectedGender,
      onChanged: (value) {
        setState(() {
          _selectedGender = value.toString();
        });
      },
    ),
    RadioListTile(
      title: Text("여성"),
      value: "여성",
      groupValue: _selectedGender,
      onChanged: (value) {
        setState(() {
          _selectedGender = value.toString();
        });
      },
    ),
  ],
)
```

✔ `groupValue` → 하나의 선택지만 가능하도록 설정.

---

### 7-3. 스위치 (`Switch`)
```dart
bool _isSwitched = false;

Switch(
  value: _isSwitched,
  onChanged: (value) {
    setState(() {
      _isSwitched = value;
    });
  },
)
```

---

## 8. 드롭다운 버튼 (DropdownButton)

사용자가 목록에서 하나를 선택할 수 있는 드롭다운 메뉴입니다.

```dart
String _selectedItem = "옵션 1";

DropdownButton<String>(
  value: _selectedItem,
  onChanged: (String? newValue) {
    setState(() {
      _selectedItem = newValue!;
    });
  },
  items: ["옵션 1", "옵션 2", "옵션 3"]
      .map<DropdownMenuItem<String>>((String value) {
    return DropdownMenuItem<String>(
      value: value,
      child: Text(value),
    );
  }).toList(),
)
```

---

## 🎯 정리

✔ `TextField` → 기본 입력 필드  
✔ `TextEditingController` → 입력값을 가져올 때 사용  
✔ `Form` & `TextFormField` → 입력값 검증 (Validation) 가능  
✔ 다양한 입력 필드 → 숫자 입력, 비밀번호 입력, 다중 줄 입력  
✔ 체크박스, 라디오 버튼, 스위치 → 다양한 사용자 입력 옵션  
✔ 드롭다운 버튼 → 사용자가 목록에서 선택 가능  
