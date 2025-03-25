import java.util.ArrayList;

public class ArrayListExample {
    public static void main(String[] args) {

        // ArrayList 선언 및 생성
        ArrayList<String> fruits = new ArrayList<>();

        // 요소 추가 (add)
        fruits.add("🍎 사과");
        fruits.add("🍌 바나나");
        fruits.add("🍇 포도");

        // 요소 출력 (get)
        System.out.println("첫 번째 과일: " + fruits.get(0));

        // 전체 반복 출력 (for-each)
        System.out.println("과일 목록:");
        for (String fruit : fruits) {
            System.out.println("- " + fruit);
        }

        // 요소 삽입 (특정 위치에 추가)
        fruits.add(1, "🍊 오렌지");

        // 요소 수정 (set)
        fruits.set(2, "🍍 파인애플");

        // 요소 삭제 (remove)
        fruits.remove("🍇 포도");

        // 크기 확인
        System.out.println("총 과일 개수: " + fruits.size());

        // 전체 출력 (인덱스 기반)
        System.out.println("최종 과일 목록:");
        for (int i = 0; i < fruits.size(); i++) {
            System.out.println(i + ": " + fruits.get(i));
        }

        // 포함 여부 확인 (contains)
        if (fruits.contains("🍌 바나나")) {
            System.out.println("🍌 바나나가 목록에 있습니다!");
        } else {
            System.out.println("🍌 바나나는 목록에 없습니다.");
        }

        // 전체 삭제 (clear)
        fruits.clear();
        System.out.println("모든 과일 제거됨. 현재 크기: " + fruits.size());
    }
}
