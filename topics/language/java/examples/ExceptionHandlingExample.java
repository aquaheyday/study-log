import java.io.*;

public class ExceptionHandlingExample {

    public static void main(String[] args) {

        System.out.println("📌 예외 처리 예제 시작");

        // 1. Checked 예외 처리 예제
        try {
            readFile("nonexistent.txt");
        } catch (IOException e) {
            System.out.println("📂 [Checked 예외] 파일 오류: " + e.getMessage());
        }

        // 2. Unchecked 예외 처리 예제
        try {
            int result = divide(10, 0);
            System.out.println("결과: " + result);
        } catch (ArithmeticException e) {
            System.out.println("🧮 [Unchecked 예외] 산술 오류: " + e.getMessage());
        }

        System.out.println("✅ 예외 처리 예제 종료");
    }

    // Checked 예외 - 반드시 예외 처리 필요
    public static void readFile(String filename) throws IOException {
        FileReader reader = new FileReader(filename); // 파일이 없으면 FileNotFoundException
        BufferedReader br = new BufferedReader(reader);
        System.out.println("파일 내용: " + br.readLine());
        br.close();
    }

    // Unchecked 예외 - 실행 시 발생, 처리 선택 가능
    public static int divide(int a, int b) {
        return a / b; // b가 0이면 ArithmeticException 발생
    }
}
