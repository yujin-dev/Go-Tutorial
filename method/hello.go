package main

import "fmt"

func main() {
	rect := Rect{10, 20}
	area := rect.area()
	println(area)
	fmt.Println(rect)
	area2 := rect.area2()
	println(area2)
	fmt.Println(rect)
}

type Rect struct {
	width, height int
}

// value receiver :struct 데이터를 copy( immutable )
func (r Rect) area() int { // (r Rect) : receiver
	r.width++
	return r.width * r.height
}

// 포인터 receiver :struct 포인터( mutable )
func (r *Rect) area2() int {
	r.width++
	return r.width * r.height
}
