package main

func ex1() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	println("SUM: ", sum)
}

func ex2() {
	n := 1
	for n < 100 {
		n *= 2
		println(n)
	}
	println("FINAL: ", n)
}

func ex3() {
	names := []string{"홍길동", "이순신", "강감찬"}
	for index, name := range names { // 인덱스, 해당값
		println(index, name)
	}
}

func ex4() {
	a := 1
	for a < 15 {
		if a == 5 { // 5
			a += a // 5+5
			println("continue", a)
			continue // loop 시작으로 돌아감
		}
		a++ // 2,3,4, 5
		println(a)
		if a > 10 {
			print("breakout ", a)
			break
		}
	}
	if a == 11 {
		goto END
	}
	println(a)
END:
	println("END")
}

func ex5() {
	i := 0
L1:
	for { // 아무 조건없는 for : 무한루프
		if i == 0 {
			break L1 // break 레이블
		}
	}
	println(i, "OK")
}

func main() {
	// ex1()
	// ex2()
	// ex3()
	// ex4()
	ex5()
}
