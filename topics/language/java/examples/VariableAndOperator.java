public class VariableAndOperator {
    public static void main(String[] args) {
        // 변수 선언 및 초기화
        int a = 10;
        int b = 3;

        // 산술 연산자
        System.out.println("a + b = " + (a + b));
        System.out.println("a - b = " + (a - b));
        System.out.println("a * b = " + (a * b));
        System.out.println("a / b = " + (a / b));  // 정수 나눗셈
        System.out.println("a % b = " + (a % b));  // 나머지

        // 증감 연산자
        int x = 5;
        System.out.println("x++ = " + (x++));  // 후위 증가
        System.out.println("++x = " + (++x));  // 전위 증가

        // 비교 연산자
        System.out.println("a == b: " + (a == b));
        System.out.println("a != b: " + (a != b));
        System.out.println("a > b: " + (a > b));
        System.out.println("a < b: " + (a < b));

        // 논리 연산자
        boolean result = (a > 5) && (b < 5);   // AND
        System.out.println("(a > 5) && (b < 5): " + result);

        result = (a > 5) || (b > 5);           // OR
        System.out.println("(a > 5) || (b > 5): " + result);

        result = !(a > b);                     // NOT
        System.out.println("!(a > b): " + result);
    }
}
