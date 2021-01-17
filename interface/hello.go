package main

import (
	"fmt"
	"math"
)

func main() {
	r := Rect{10., 20.}
	c := Circle{10}
	showArea(r, c)
	emptyInterface()
}

func emptyInterface() {
	var x interface{}
	// dynamice interface
	x = 1
	x = "Tom"
	printIt(x)
}

func printIt(v interface{}) {
	fmt.Println(v)
}

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.area() // shapes 인터페이스 메서드 호출
		println(a)
	}
}

type Shape interface {
	area() float64
	perimeter() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (r Rect) perimeter() float64 {
	return 2 * (r.width * r.height) // float32면 오류
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func typeAssert() {
	var a interface{} = 1
	i := a
	j := a.(int) // type assertion: x != nil && x가 int(type)임
	println(i)
	println(j)
}
