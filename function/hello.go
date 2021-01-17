package main

func main() {
	msg := "hello"
	byValue(msg)
	byRef(&msg)
	println(msg)
	varFunc("This", "is", "a", "book")
	println(sum(1, 3, 5, 7, 9))
	println(sums(1, 3, 5, 7, 9))
	println(namedSums(1, 3, 5, 7, 9))

	add := func(i int, j int) int { // 익명 함수
		return i + j
	}
	println(calc(add, 10, 20))
	println(calc(func(x int, y int) int { return x - y }, 10, 20))

	callClosure()
}
func byValue(msg string) {
	// pass by value : msg 문자열이 복사되어 전달됨
	println(msg)
}

func byRef(msg *string) {
	// pass by reference : &로 변수의 주소(포인터)를 받음
	println(*msg)
	*msg = "changed" // dereferencing : 주소에 해당되는 데이터를 가져옴
}

func varFunc(msg ...string) {
	// 가변인자 : ...string
	total := ""
	for _, s := range msg {
		total += s + " "
	}
	println(total)
}

func sum(nums ...int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}

func sums(nums ...int) (int, int) {
	s := 0
	count := 0
	for _, n := range nums {
		s += n
		count++
	}
	return s, count
}

func namedSums(nums ...int) (count int, total int) {
	// Named Return Parameter
	for _, n := range nums {
		total += n
	}
	count = len(nums)
	return // empty
}

func calc(f func(int, int) int, a int, b int) int { // 함수를 인자로 받음
	result := f(a, b)
	return result
}

type calculator func(int, int) int // Custom Type

func calc2(f calculator, a int, b int) int {
	result := f(a, b)
	return result
}

func nextValue() func() int { // closure
	i := 0
	return func() int {
		i++
		return i
	}
}

func callClosure() {
	next := nextValue()
	println(next())
	println(next())
	println(next())
	anotherNext := nextValue()
	println(anotherNext())
	println(anotherNext())
}
