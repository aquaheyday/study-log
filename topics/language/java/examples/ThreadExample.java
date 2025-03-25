public class MultiThreadExample {

    // 스레드 정의 방법 1: Thread 클래스 상속
    static class MyThread extends Thread {
        private String name;

        public MyThread(String name) {
            this.name = name;
        }

        @Override
        public void run() {
            for (int i = 1; i <= 5; i++) {
                System.out.println(name + " 실행 중... (" + i + ")");
                try {
                    Thread.sleep(500); // 0.5초 대기
                } catch (InterruptedException e) {
                    System.out.println(name + " 인터럽트 발생!");
                }
            }
        }
    }

    // 스레드 정의 방법 2: Runnable 구현
    static class MyRunnable implements Runnable {
        private String name;

        public MyRunnable(String name) {
            this.name = name;
        }

        @Override
        public void run() {
            for (int i = 1; i <= 5; i++) {
                System.out.println(name + " 작업 중... (" + i + ")");
                try {
                    Thread.sleep(300); // 0.3초 대기
                } catch (InterruptedException e) {
                    System.out.println(name + " 인터럽트 발생!");
                }
            }
        }
    }

    // main 메서드
    public static void main(String[] args) {
        // Thread 클래스를 상속한 스레드 실행
        MyThread t1 = new MyThread("🧵 스레드 A");
        t1.start();

        // Runnable 인터페이스를 구현한 스레드 실행
        Thread t2 = new Thread(new MyRunnable("🧶 스레드 B"));
        t2.start();

        // 메인 스레드도 함께 작업
        for (int i = 1; i <= 5; i++) {
            System.out.println("💻 메인 스레드 실행... (" + i + ")");
            try {
                Thread.sleep(400);
            } catch (InterruptedException e) {
                System.out.println("메인 스레드 인터럽트 발생!");
            }
        }

        System.out.println("🏁 모든 스레드 작업 종료 (main 끝)");
    }
}
