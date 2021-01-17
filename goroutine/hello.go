package main

import (
	"fmt"
	"sync"
	"time"
)

// goroutine : go 런타임을 관리하는 논리적 쓰레드
// os 쓰레드가 아닌 go 런타임이 자체 관리
// 기본적으로 os 쓰레드는 1MB 스택을 갖는반면, goroutine은 KB단위의 스택을 메모리에서 차지함
func main() {
	anonymous()
}

func say(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
}

func syncAsync() {
	// 1. 동기적 실행
	say("Sync")
	//2 .비동기적 실행: 별도의 goroutine이 동작함
	// 메인루틴은 계속 다음으로 진행함
	go say("Async1")
	go say("Async2")
	go say("Async3")
	time.Sleep(time.Second * 3)
}

func anonymous() {
	var wait sync.WaitGroup
	wait.Add(2) // 2개의 gorouinte을 대기
	go func() { // 익명함수
		defer wait.Done() // 끝나면 Done 호출
		fmt.Println("halo")
	}() // 실행
	go func(msg string) { // 익명함수: 파라미터 지정
		defer wait.Done()
		fmt.Println(msg)
	}("HI Param")
	wait.Wait()
}
