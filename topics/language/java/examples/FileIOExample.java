import java.io.*;

public class FileIOExample {

    public static void main(String[] args) {
        String filename = "example.txt";

        // 1. 파일 쓰기 (write)
        try (BufferedWriter writer = new BufferedWriter(new FileWriter(filename))) {
            writer.write("안녕하세요, Java 파일 입출력 예제입니다.\n");
            writer.write("두 번째 줄입니다.");
            System.out.println("✅ 파일 쓰기 완료");
        } catch (IOException e) {
            System.out.println("❌ 파일 쓰기 중 오류 발생: " + e.getMessage());
        }

        System.out.println("---------------");

        // 2. 파일 읽기 (read)
        try (BufferedReader reader = new BufferedReader(new FileReader(filename))) {
            String line;
            System.out.println("📄 파일 내용:");
            while ((line = reader.readLine()) != null) {
                System.out.println(line);
            }
        } catch (IOException e) {
            System.out.println("❌ 파일 읽기 중 오류 발생: " + e.getMessage());
        }
    }
}
