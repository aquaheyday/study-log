public class ControlStatement {
    public static void main(String[] args) {

        // 1. 조건문 if-else
        int score = 85;

        if (score >= 90) {
            System.out.println("A 학점");
        } else if (score >= 80) {
            System.out.println("B 학점");
        } else if (score >= 70) {
            System.out.println("C 학점");
        } else {
            System.out.println("D 학점");
        }

        // 2. switch-case 문
        int day = 3;

        switch (day) {
            case 1:
                System.out.println("월요일");
                break;
            case 2:
                System.out.println("화요일");
                break;
            case 3:
                System.out.println("수요일");
                break;
            default:
                System.out.println("그 외 요일");
                break;
        }

        // 3. for 반복문
        for (int i = 1; i <= 5; i++) {
            System.out.println("for 반복: " + i);
        }

        // 4. while 반복문
        int count = 1;
        while (count <= 3) {
            System.out.println("while 반복: " + count);
            count++;
        }

        // 5. do-while 반복문
        int n = 1;
        do {
            System.out.println("do-while 반복: " + n);
            n++;
        } while (n <= 2);

        // 6. break & continue
        for (int i = 1; i <= 5; i++) {
            if (i == 3) {
                continue; // 3은 건너뜀
            }
            if (i == 4) {
                break;    // 4에서 반복 종료
            }
            System.out.println("break/continue 예제: " + i);
        }
    }
}
