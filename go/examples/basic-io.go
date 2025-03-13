package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 1. 표준 출력 (Print, Println, Printf)
	fmt.Print("🚀 Print: 개행 없음")       // 개행 없음
	fmt.Println("✅ Println: 개행 있음")  // 개행 있음
	fmt.Printf("🔢 Printf: 숫자 출력 %d\n", 100) // 형식 지정 출력

	// 2. 표준 입력 (fmt.Scan, fmt.Scanf)
	var name string
	fmt.Print("📝 이름을 입력하세요: ")
	fmt.Scan(&name) // 공백 이전까지만 입력 받음
	fmt.Println("👋 반갑습니다,", name)

	// 3. 표준 입력 (fmt.Scanln) - 한 줄 전체 입력 받기
	var fullName string
	fmt.Print("📝 풀네임을 입력하세요: ")
	fmt.Scanln(&fullName) // 한 줄 전체 입력 받음
	fmt.Println("👤 입력한 풀네임:", fullName)

	// 4. 표준 입력 (bufio.NewReader) - 띄어쓰기 포함 입력 받기
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("💬 메시지를 입력하세요: ")
	message, _ := reader.ReadString('\n') // 개행 문자까지 입력 받음
	fmt.Println("📩 입력한 메시지:", message)
}
