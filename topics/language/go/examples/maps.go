package main

import "fmt"

func main() {
	// 맵 선언 및 초기화
	var person map[string]string // 빈 맵 선언 (nil 상태)
	person = make(map[string]string) // make()를 사용하여 초기화

	// 키-값 추가
	person["name"] = "Alice"
	person["age"] = "25"
	person["city"] = "Seoul"

	fmt.Println("맵 내용:", person)

	// 리터럴을 사용한 맵 초기화
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}

	fmt.Println("색상 맵:", colors)

	// 값 가져오기
	name := person["name"]
	fmt.Println("이름:", name)

	// 존재하지 않는 키 조회
	job, exists := person["job"]
	if exists {
		fmt.Println("직업:", job)
	} else {
		fmt.Println("직업 정보 없음")
	}

	// 키 삭제
	delete(person, "city")
	fmt.Println("삭제 후 맵:", person)

	// 맵의 길이 확인
	fmt.Println("맵의 길이:", len(person))

	// 맵 순회
	for key, value := range colors {
		fmt.Printf("키: %s, 값: %s\n", key, value)
	}

	// 맵의 값 변경
	person["age"] = "30"
	fmt.Println("나이 변경 후:", person)
}
