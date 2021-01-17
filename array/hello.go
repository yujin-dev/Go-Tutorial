package main

import "fmt"

func main() {
	defarray()
	initarray()
	mularray()
	slice()
	sliceMake()
	partSlice()
	appendSlice()
	changeSlice()
	appendSlices()
	copySlice()
}

func defarray() {
	// 배열: 연속적인 메모리 공간에 동일한 타입의 데이터를 순서대로 저장
	var a [3]int // 정수로 구성된 배열(길이:3)
	a[0] = 1
	a[1] = 2
	a[2] = 3
	println(a[1])
}
func initarray() {
	// 초기값 설정
	var a1 = [3]int{1, 2, 3}
	// 배열 크기 자동 설정
	var a3 = [...]int{1, 2, 3}
	println(a1[1], a3[1])
}

func mularray() {
	// 다배열
	var multiArray [3][4][5]int // 3*4*5 array
	multiArray[0][1][2] = 10
	println(multiArray[0][1][1]) // 값 할당안하면 초기값 = 0
	var multiArray2 = [2][3]int{ // 2행*3열 array
		{1, 2, 3},
		{4, 5, 6},
	}
	println(multiArray2[1][2]) // [1,2]
}

func slice() {
	// 동적 크기 설정
	var a []int // 크기 설정X
	a = []int{1, 2, 3}
	a[1] = 10
	fmt.Println(a)
}

func sliceMake() {
	// make : slice의 길이(len), 용량(cap)을 지정
	s := make([]int, 5, 10) // 타입, 길이, 용량
	println(len(s), cap(s))
	var nils []int
	if nils == nil {
		println(len(nils), cap(nils))
	}
}

func partSlice() {
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(s[2:5]) // 2,3,4
	fmt.Println(s[1:])  // 1, 2, 3, 4, 5
	fmt.Println(s[:])   // 0, 1, 2, 3, 4, 5
}

func appendSlice() {
	s := []int{0, 1}
	s = append(s, 2) // 0, 1, 2
	fmt.Println(s)
}

func changeSlice() {
	a := make([]int, 0, 3)
	for i := 1; i <= 15; i++ {
		a = append(a, i)
		fmt.Println("길이: ", len(a), "용량: ", cap(a)) // 용량 찰때마다 2배씩 증가
	}
	fmt.Println(a)
}

func appendSlices() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	a = append(a, b...)
	fmt.Println(a)
}

func copySlice() {
	source := []int{0, 1, 2}
	target := make([]int, len(source), cap(source)*2)
	copy(target, source) // [0, 1, 2]
	fmt.Println(target)
	println(len(target), cap(target))
}
