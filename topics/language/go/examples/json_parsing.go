package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// 구조체 정의 (JSON 변환용)
type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	// JSON 인코딩 (구조체 → JSON)
	person := Person{Name: "Alice", Age: 30, Email: "alice@example.com"}
	jsonData, err := json.Marshal(person) // JSON 변환
	if err != nil {
		fmt.Println("JSON 인코딩 오류:", err)
	} else {
		fmt.Println("JSON 인코딩 결과:", string(jsonData))
	}

	// JSON 인코딩 (들여쓰기 포함)
	jsonPretty, err := json.MarshalIndent(person, "", "  ") // 들여쓰기 포함
	if err != nil {
		fmt.Println("JSON 인코딩 오류:", err)
	} else {
		fmt.Println("JSON 인코딩 (Pretty):")
		fmt.Println(string(jsonPretty))
	}

	// JSON 디코딩 (JSON → 구조체)
	jsonString := `{"name": "Bob", "age": 25, "email": "bob@example.com"}`
	var newPerson Person
	err = json.Unmarshal([]byte(jsonString), &newPerson) // JSON을 구조체로 변환
	if err != nil {
		fmt.Println("JSON 디코딩 오류:", err)
	} else {
		fmt.Println("JSON 디코딩 결과:", newPerson)
	}

	// JSON 파일 저장
	file, err := os.Create("person.json")
	if err != nil {
		fmt.Println("파일 생성 오류:", err)
	} else {
		defer file.Close()
		_, err = file.Write(jsonPretty)
		if err != nil {
			fmt.Println("파일 저장 오류:", err)
		} else {
			fmt.Println("JSON 데이터가 person.json 파일에 저장되었습니다.")
		}
	}

	// JSON 파일 읽기
	file, err = os.Open("person.json")
	if err != nil {
		fmt.Println("파일 열기 오류:", err)
	} else {
		defer file.Close()
		decoder := json.NewDecoder(file)
		var filePerson Person
		err = decoder.Decode(&filePerson) // JSON 파일을 구조체로 변환
		if err != nil {
			fmt.Println("파일 JSON 디코딩 오류:", err)
		} else {
			fmt.Println("파일에서 읽은 JSON 데이터:", filePerson)
		}
	}
}
