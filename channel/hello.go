package main

import (
	"fmt"
	"time"
)

// channel : 채널을 통해 데이터를 주고받음
// make로 미리 생성
// 별도의 lock없이 데이터를 동기화하는데 사용됨

func main() {
	// makeChan()
	// waitChannel()
	buffered()

	ch := make(chan string, 1)
	sendChan(ch)
	receiveChan(ch)

	closeChan()
	rangeCloseChan()

	selectChan()
}

// unbuffered channel
func makeChan() {
	ch := make(chan int)
	go func() {
		ch <- 123 // 채털에 123을 전송: 채널 <- 데이터
	}()
	var i int
	i = <-ch // 채널로부터 123을 받음: <- 채널
	println(i)
}

func waitChannel() {
	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		done <- true
	}()
	<-done // goroutine에서 작업 실행 중이면 메인루틴에서 수신 대기 : 없으면 기다리지 않고 종료됨
}

// buffered channel
func buffered() {
	ch := make(chan int, 1)
	ch <- 101
	fmt.Println(<-ch)
}

func sendChan(ch chan<- string) { // 송신
	ch <- "Data"
}
func receiveChan(ch <-chan string) { // 수신
	data := <-ch
	fmt.Println(data)
}

func closeChan() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	println(<-ch)
	println(<-ch)
	if _, success := <-ch; !success {
		println("더이상 데이터 없음")
	}
}

func rangeCloseChan() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	for i := range ch {
		println(i)
	}
}

func selectChan() {
	done1 := make(chan bool)
	done2 := make(chan bool)
	go run1(done1)
	go run2(done2)
EXIT:
	for { // 무한루프
		select {
		case <-done1:
			println("RUN1 완료")
		case <-done2:
			println("RUN2 완료")
			break EXIT
		}
	}
}

func run1(done chan bool) {
	time.Sleep(1 * time.Second)
	done <- true
}

func run2(done chan bool) {
	time.Sleep(2 * time.Second)
	done <- true
}
