package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 표준 출력 (fmt 패키지 사용)
	fmt.Print("fmt.Print 사용 - 개행 없음")  // 개행 없이 출력
	fmt.Println("fmt.Println 사용 - 개행 있음") // 개행 포함 출력
	fmt.Printf("fmt.Printf 사용 - 숫자: %d, 문자열: %s\n", 42, "Hello") // 형식 지정 출력

	// 표준 입력 (fmt.Scan 사용)
	var name string
	fmt.Print("이름을 입력하세요: ")
	fmt.Scan(&name) // 공백 전까지 입력받음
	fmt.Println("입력한 이름:", name)

	// 표준 입력 (fmt.Scanln 사용)
	var fullName string
	fmt.Print("풀네임을 입력하세요: ")
	fmt.Scanln(&fullName) // 한 줄 전체 입력받음
	fmt.Println("입력한 풀네임:", fullName)

	// 표준 입력 (bufio.NewReader 사용)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("메시지를 입력하세요: ")
	message, _ := reader.ReadString('\n') // 개행 문자까지 입력받음
	fmt.Println("입력한 메시지:", message)

	// 표준 입력 (파일에서 읽기)
	file, err := os.Open("example.txt") // 파일 열기
	if err != nil {
		fmt.Println("파일을 열 수 없습니다:", err)
	} else {
		defer file.Close() // 함수 종료 시 파일 닫기
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println("파일 내용:", scanner.Text()) // 파일에서 한 줄씩 읽어 출력
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("파일 읽기 오류:", err)
		}
	}

	// 표준 출력 (파일에 쓰기)
	outputFile, err := os.Create("output.txt") // 파일 생성
	if err != nil {
		fmt.Println("파일을 생성할 수 없습니다:", err)
	} else {
		defer outputFile.Close()
		_, err = outputFile.WriteString("Go 언어 파일 입출력 예제\n")
		if err != nil {
			fmt.Println("파일 쓰기 오류:", err)
		} else {
			fmt.Println("파일에 성공적으로 저장되었습니다: output.txt")
		}
	}
}
