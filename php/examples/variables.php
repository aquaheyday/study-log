<?php

//PHP에서 변수는 `$` 기호로 선언되며, 자료형은 자동으로 결정됩니다.  
//여러 종류의 자료형(문자열, 정수, 실수, 불리언, 배열, 객체 등)이 존재합니다.

// 1. 변수 선언과 사용
$name = "John";  // 문자열
$age = 25;       // 정수
$height = 5.9;   // 실수
$isStudent = true; // 불리언

echo "이름: $name, 나이: $age, 키: $height, 학생 여부: $isStudent";
// PHP는 동적 타입 언어로, 변수 타입을 명시하지 않아도 됨  
// 변수 이름은 문자로 시작하고 _를 포함할 수 있음  

// 2. 자료형 종류 및 확인 (gettype())
$var1 = "Hello";   // 문자열
$var2 = 10;        // 정수
$var3 = 10.5;      // 실
$var4 = true;      // 불리언
$var5 = array(1, 2, 3); // 배열

echo gettype($var1) . "<br>"; // string
echo gettype($var2) . "<br>"; // integer
echo gettype($var3) . "<br>"; // double
echo gettype($var4) . "<br>"; // boolean
echo gettype($var5) . "<br>"; // array

// 3. 형 변환 (Type Casting)
// PHP에서는 명시적 형 변환을 사용할 수 있습니다.
// 명시적 형 변환을 위해 (int), (float), (bool), (string), (array) 사용 가능
$number = "123"; // 문자열

$intNumber = (int) $number;  // 정수 변환
$floatNumber = (float) $number; // 실수 변환
$boolValue = (bool) $number;  // 불리언 변환

echo "정수: $intNumber, 실수: $floatNumber, 불리언: $boolValue";

// 4. 자동 형 변환 (Implicit Type Conversion)
// PHP는 자동으로 변수의 자료형을 변환할 수 있습니다.
$sum = "5" + 10; // 문자열 "5"가 정수 5로 변환됨
echo "결과: $sum"; // 15 출력

// 5. 불리언(Boolean) 값 변환
$var1 = (bool) 0; // false
$var2 = (bool) ""; // false
$var3 = (bool) "Hello"; // true
$var4 = (bool) 123; // true

var_dump($var1, $var2, $var3, $var4);
// 숫자 0과 빈 문자열 ""은 false로 변환됨
// 그 외의 값은 true로 간주됨
// 자동 형 변환을 통해 문자열과 숫자를 더할 수 있음

// 6. 문자열 처리 (strlen(), strtoupper())
// strlen()을 사용하여 문자열 길이 확인 가능 
// strtoupper()로 문자열을 대문자로 변환 가능  
$text = "Hello, PHP!";
echo strlen($text); // 문자열 길이 출력
echo strtoupper($text); // 대문자로 변환

// 7. 배열 (Indexed & Associative Arrays)
// 배열을 사용하여 여러 값을 저장 가능
// 인덱스 배열
$fruits = array("Apple", "Banana", "Cherry");
echo $fruits[0]; // Apple

// 연관 배열 (Associative Array)
$person = array("name" => "John", "age" => 25);
echo $person["name"]; // John

// 8. 객체(Object) 생성
// 클래스를 정의하고 객체를 생성할 수 있음
class Person {
    public $name;
    public $age;
    
    public function __construct($name, $age) {
        $this->name = $name;
        $this->age = $age;
    }

    public function introduce() {
        return "이름: $this->name, 나이: $this->age";
    }
}

$person1 = new Person("Alice", 30);
echo $person1->introduce();

// 9. 변수 존재 여부 확인 (isset(), empty())
// isset($var): 변수가 존재하는지 확인
// empty($var): 변수가 비어 있는지 확인
$name = "PHP";
$age = null;

echo isset($name) ? "변수 존재" : "변수 없음"; // "변수 존재"
echo empty($age) ? "비어 있음" : "비어 있지 않음"; // "비어 있음"

// 10. 상수(Constant) 선언
// define() 또는 const를 사용하여 상수를 선언 가능
// 한 번 정의된 상수는 변경할 수 없음
define("SITE_NAME", "My Website");
const PI = 3.14;

echo SITE_NAME; // My Website
echo PI; // 3.14
?>
