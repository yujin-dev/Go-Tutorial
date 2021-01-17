package main

import (
	"fmt"
	"os"
)

func main() {
	// withDefer()
	// openFile("Invalid.txt")
	withRecover("Invalid.txt")
	println("Even errors, done without exception")
}

func withDefer() {
	f, err := os.Open("1.txt")
	if err != nil {
		panic(err) // 현재 함수를 즉히 멈추고 defer 실행 후 리턴함
	}

	defer f.Close() // main마지막에 실행됨: 파일 닫음
	bytes := make([]byte, 1024)
	f.Read(bytes)
	println(len(bytes))
}

func openFile(fn string) {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()
}

func withRecover(fn string) {
	// recover : panic함수를 정상으로 되돌림.(exit 하지 않음)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR: ", r)
		}
	}()
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()
}
