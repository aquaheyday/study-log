public class InheritanceExample {

    // 부모 클래스
    static class Animal {
        protected String name;

        public Animal(String name) {
            this.name = name;
        }

        public void speak() {
            System.out.println(name + "이(가) 소리를 냅니다.");
        }
    }

    // 자식 클래스 - Dog
    static class Dog extends Animal {
        public Dog(String name) {
            super(name);
        }

        // 메서드 오버라이딩
        @Override
        public void speak() {
            System.out.println(name + "이(가) 멍멍 짖습니다.");
        }
    }

    // 자식 클래스 - Cat
    static class Cat extends Animal {
        public Cat(String name) {
            super(name);
        }

        @Override
        public void speak() {
            System.out.println(name + "이(가) 야옹 울어요.");
        }
    }

    // 실행 메서드
    public static void main(String[] args) {
        Animal generic = new Animal("동물");
        Dog dog = new Dog("바둑이");
        Cat cat = new Cat("나비");

        generic.speak();  // 동물이(가) 소리를 냅니다.
        dog.speak();      // 바둑이이(가) 멍멍 짖습니다.
        cat.speak();      // 나비이(가) 야옹 울어요.

        // 다형성 예제
        Animal pet = new Dog("초코");
        pet.speak();  // Dog의 speak 호출 (다형성), 출력: 초코이(가) 멍멍 짖습니다.
    }
}
