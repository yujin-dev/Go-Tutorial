package main

import "fmt"

// switch
func grade(score int) {
	switch {
	case score >= 90:
		println("A")
	case score >= 80:
		println("B")
	case score >= 70:
		println("C")
	default:
		println("No Hope")
	}
}

func check(val int) {
	switch val {
	case 1:
		fmt.Println("1 이하")
		fallthrough
	case 2:
		fmt.Println("2 이하")
		fallthrough // 아래 조건 무조건 실행
	case 3:
		fmt.Println("3 이하")
		fallthrough
	default:
		fmt.Println("default")
	}
}

func main() {
	// grade(100)
	check(2)
}
