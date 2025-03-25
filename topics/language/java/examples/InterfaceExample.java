public class InterfaceExample {

    // 인터페이스 정의
    interface Animal {
        void speak(); // 추상 메서드
        void move();
    }

    // 인터페이스 구현 클래스 - Dog
    static class Dog implements Animal {
        @Override
        public void speak() {
            System.out.println("멍멍!");
        }

        @Override
        public void move() {
            System.out.println("강아지가 네 발로 뛰어요.");
        }
    }

    // 인터페이스 구현 클래스 - Bird
    static class Bird implements Animal {
        @Override
        public void speak() {
            System.out.println("짹짹!");
        }

        @Override
        public void move() {
            System.out.println("새가 날아갑니다.");
        }
    }

    // 실행 메서드
    public static void main(String[] args) {
        Animal dog = new Dog();
        Animal bird = new Bird();

        dog.speak();  // 멍멍!
        dog.move();   // 강아지가 네 발로 뛰어요.

        bird.speak(); // 짹짹!
        bird.move();  // 새가 날아갑니다.
    }
}
