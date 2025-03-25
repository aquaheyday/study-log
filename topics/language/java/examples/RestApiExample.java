package com.example.demo;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.*;

import java.util.*;

@SpringBootApplication
@RestController
@RequestMapping("/users")
public class RestApiExample {

    // 사용자 목록 저장용 (임시 메모리)
    private Map<Long, String> userMap = new HashMap<>();
    private long idCounter = 1;

    public static void main(String[] args) {
        SpringApplication.run(RestApiExample.class, args);
    }

    // 사용자 생성 (POST /users)
    @PostMapping
    public Map<String, Object> createUser(@RequestBody Map<String, String> request) {
        String name = request.get("name");
        userMap.put(idCounter, name);

        return Map.of(
            "id", idCounter++,
            "name", name,
            "message", "✅ 사용자 생성 완료"
        );
    }

    // 사용자 전체 조회 (GET /users)
    @GetMapping
    public Map<Long, String> getAllUsers() {
        return userMap;
    }

    // 단일 사용자 조회 (GET /users/{id})
    @GetMapping("/{id}")
    public Map<String, Object> getUser(@PathVariable Long id) {
        String name = userMap.get(id);
        if (name == null) {
            return Map.of("error", "❌ 해당 ID의 사용자가 없습니다.");
        }
        return Map.of("id", id, "name", name);
    }

    // 사용자 수정 (PUT /users/{id})
    @PutMapping("/{id}")
    public Map<String, Object> updateUser(@PathVariable Long id, @RequestBody Map<String, String> request) {
        if (!userMap.containsKey(id)) {
            return Map.of("error", "❌ 사용자 없음");
        }
        userMap.put(id, request.get("name"));
        return Map.of("id", id, "updatedName", request.get("name"));
    }

    // 사용자 삭제 (DELETE /users/{id})
    @DeleteMapping("/{id}")
    public Map<String, String> deleteUser(@PathVariable Long id) {
        if (userMap.remove(id) != null) {
            return Map.of("message", "🗑️ 삭제 완료");
        }
        return Map.of("error", "❌ 해당 사용자가 없습니다.");
    }
}
