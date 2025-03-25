// Person.java - 클래스 정의 예제
public class Person {

    // 필드 (멤버 변수)
    private String name;
    private int age;

    // 생성자
    public Person(String name, int age) {
        this.name = name;
        this.age = age;
    }

    // 메서드
    public void introduce() {
        System.out.println("안녕하세요, 제 이름은 " + name + "이고 나이는 " + age + "살입니다.");
    }

    public void birthday() {
        age++;
        System.out.println(name + "의 생일입니다! 이제 " + age + "살이에요.");
    }

    // Getter
    public String getName() {
        return name;
    }

    public int getAge() {
        return age;
    }
}


// Main.java - Person 클래스를 사용하는 예제
public class Main {
    public static void main(String[] args) {
        // 객체 생성
        Person p1 = new Person("홍길동", 25);
        Person p2 = new Person("이순신", 40);

        // 메서드 호출
        p1.introduce();
        p2.introduce();

        p1.birthday();
        p2.birthday();

        System.out.println(p1.getName() + "의 현재 나이: " + p1.getAge());
    }
}

