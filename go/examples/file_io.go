package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 파일에 문자열 쓰기 (os.WriteFile 사용)
	err := ioutil.WriteFile("example1.txt", []byte("Hello, Go 파일 입출력!\n"), 0644)
	if err != nil {
		fmt.Println("파일 쓰기 오류:", err)
	} else {
		fmt.Println("example1.txt 파일이 성공적으로 저장되었습니다.")
	}

	// 파일에 문자열 추가하기 (os.OpenFile 사용)
	file, err := os.OpenFile("example1.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("파일 열기 오류:", err)
	} else {
		defer file.Close()
		_, err = file.WriteString("이 줄은 추가된 내용입니다.\n")
		if err != nil {
			fmt.Println("파일 쓰기 오류:", err)
		} else {
			fmt.Println("example1.txt 파일에 내용이 추가되었습니다.")
		}
	}

	// 파일에서 문자열 읽기 (ioutil.ReadFile 사용)
	content, err := ioutil.ReadFile("example1.txt")
	if err != nil {
		fmt.Println("파일 읽기 오류:", err)
	} else {
		fmt.Println("파일 내용 읽기 (ioutil.ReadFile):")
		fmt.Println(string(content))
	}

	// 파일에서 한 줄씩 읽기 (bufio.NewScanner 사용)
	file, err = os.Open("example1.txt")
	if err != nil {
		fmt.Println("파일 열기 오류:", err)
	} else {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		fmt.Println("파일 내용 읽기 (bufio.Scanner):")
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("파일 스캔 오류:", err)
		}
	}

	// 파일 삭제하기
	err = os.Remove("example1.txt")
	if err != nil {
		fmt.Println("파일 삭제 오류:", err)
	} else {
		fmt.Println("example1.txt 파일이 삭제되었습니다.")
	}
}
