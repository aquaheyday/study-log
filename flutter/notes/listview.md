# 🔲 리스트와 스크롤

Flutter에서 **리스트(ListView)와 스크롤(ScrollView)** 을 활용하면 긴 목록을 표시하고, 화면을 넘겨서 콘텐츠를 확인할 수 있습니다.

---

## 1. 리스트(ListView)란?

Flutter에서 `ListView`는 **스크롤 가능한 목록을 생성하는 가장 기본적인 위젯**입니다.

✔ 리스트를 사용하는 경우
- 채팅 메시지 목록
- 쇼핑몰 상품 목록
- 뉴스 피드, 게시판

✔ ListView의 주요 기능
- 자동 스크롤 지원 (수직, 수평 방향)
- 리스트 최적화 (`ListView.builder`)
- 고정된 리스트 (`ListView`) & 동적 리스트 (`ListView.builder`) 지원
- Grid 형태 지원 (`GridView`)

---

## 2. `ListView` 기본 사용법

### 2-1. 기본 `ListView`
```dart
ListView(
  children: [
    ListTile(title: Text("Item 1")),
    ListTile(title: Text("Item 2")),
    ListTile(title: Text("Item 3")),
  ],
)
```

✔ **ListTile** → 리스트의 각 항목을 쉽게 만들 수 있는 기본 제공 위젯

---

### 2-2. `ListView.builder` (동적 리스트)
- 데이터가 많거나 **반복되는 UI**를 효율적으로 만들 때 사용.
- **성능이 좋고, 필요할 때만 항목을 렌더링**함.

```dart
ListView.builder(
  itemCount: 10, // 리스트 아이템 개수
  itemBuilder: (context, index) {
    return ListTile(
      title: Text("Item ${index + 1}"),
      leading: Icon(Icons.star), // 왼쪽 아이콘 추가
      trailing: Icon(Icons.arrow_forward), // 오른쪽 아이콘 추가
    );
  },
)
```

✔ **itemCount** → 리스트의 개수 지정  
✔ **itemBuilder** → 반복적으로 아이템을 생성하는 함수  

---

### 2-3. `ListView.separated` (구분선 있는 리스트)
- 각 아이템 사이에 **구분선(divider)** 추가.

```dart
ListView.separated(
  itemCount: 10,
  separatorBuilder: (context, index) => Divider(), // 구분선 추가
  itemBuilder: (context, index) {
    return ListTile(
      title: Text("Item ${index + 1}"),
    );
  },
)
```

✔ `separatorBuilder` → 항목 사이의 **구분선**을 정의  

---

## 3. 리스트 아이템 꾸미기

### 3-1. `ListTile`을 사용한 리스트 아이템
```dart
ListTile(
  leading: CircleAvatar(
    backgroundColor: Colors.blue,
    child: Text("A"),
  ),
  title: Text("Alice"),
  subtitle: Text("Flutter Developer"),
  trailing: Icon(Icons.message),
)
```

✔ `leading` → 왼쪽 아이콘 (예: 프로필 사진)  
✔ `title` → 주요 텍스트  
✔ `subtitle` → 보조 텍스트  
✔ `trailing` → 오른쪽 아이콘 (예: 메시지 아이콘)

---

### 3-2. `Card`를 활용한 리스트 아이템
```dart
Card(
  elevation: 4, // 그림자 효과
  shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(10)),
  child: ListTile(
    leading: Icon(Icons.person, size: 40),
    title: Text("John Doe"),
    subtitle: Text("Mobile Developer"),
    trailing: Icon(Icons.call),
  ),
)
```

✔ `Card` → 리스트 아이템을 **카드 형태로 꾸밀 때 사용**  
✔ `elevation` → 그림자 효과 추가  

---

## 4. 스크롤 가능한 레이아웃 만들기

### 4-1. `SingleChildScrollView` (스크롤 가능 영역)
- 일반적인 위젯을 **스크롤 가능하게 만들 때** 사용.

```dart
SingleChildScrollView(
  child: Column(
    children: List.generate(
      20,
      (index) => Padding(
        padding: EdgeInsets.all(8.0),
        child: Text("Item ${index + 1}", style: TextStyle(fontSize: 20)),
      ),
    ),
  ),
)
```

✔ `SingleChildScrollView`는 한 번에 **모든 데이터를 로드하므로 성능이 떨어질 수 있음.**  
✔ 데이터가 많을 경우 `ListView`를 사용하는 것이 더 적절함.  

---

### 4-2. `GridView` (그리드 형태 리스트)
- `ListView`와 유사하지만, **격자(그리드) 형태로 배치**할 때 사용.

```dart
GridView.count(
  crossAxisCount: 2, // 열 개수
  children: List.generate(
    10,
    (index) => Card(
      margin: EdgeInsets.all(8),
      child: Center(child: Text("Item ${index + 1}")),
    ),
  ),
)
```

✔ `crossAxisCount` → 열 개수 지정  
✔ `List.generate()` → 반복문처럼 위젯을 생성  

---

### 4-3. `GridView.builder` (동적 생성)
- 많은 데이터를 효율적으로 관리.

```dart
GridView.builder(
  gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
    crossAxisCount: 2, // 2열 배치
    crossAxisSpacing: 10,
    mainAxisSpacing: 10,
    childAspectRatio: 1.0, // 가로세로 비율
  ),
  itemCount: 10,
  itemBuilder: (context, index) {
    return Container(
      color: Colors.blueAccent,
      child: Center(child: Text("Item ${index + 1}")),
    );
  },
)
```

✔ `SliverGridDelegateWithFixedCrossAxisCount` → **고정된 열 개수**  
✔ `SliverGridDelegateWithMaxCrossAxisExtent` → **최대 너비 지정**  

---

## 5. 스크롤 컨트롤 (Custom ScrollView)

### 5-1. `CustomScrollView + SliverList`
- 더 **세밀한 스크롤 컨트롤** 가능.

```dart
CustomScrollView(
  slivers: [
    SliverAppBar(
      floating: true,
      expandedHeight: 200,
      flexibleSpace: FlexibleSpaceBar(title: Text("Sliver App Bar")),
    ),
    SliverList(
      delegate: SliverChildBuilderDelegate(
        (context, index) => ListTile(title: Text("Item ${index + 1}")),
        childCount: 20,
      ),
    ),
  ],
)
```

✔ `SliverAppBar` → **스크롤 시 애니메이션 효과** 지원  
✔ `SliverList` → **최적화된 리스트** 렌더링  

---

## 🎯 정리

✔ `ListView` → 기본적인 스크롤 가능한 리스트  
✔ `ListView.builder` → 동적으로 리스트 생성 (성능 최적화)  
✔ `ListView.separated` → 리스트 항목 사이 구분선 추가  
✔ `SingleChildScrollView` → 단순한 스크롤 가능 UI  
✔ `GridView` → 격자 형태 레이아웃  
✔ `CustomScrollView` → 고급 스크롤 컨트롤  
