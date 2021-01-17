package main

import "fmt"

type person struct {
	// custom Data Type
	name string
	age  int
}
type dict struct {
	data map[int]string
}

func main() {
	makePerson()
	makeObj()
	dc := newDict()
	dc.data[1] = "A"
	fmt.Println(dc)
}

func makePerson() {
	p := person{}
	p.name = "Lee"
	p.age = 10
	fmt.Println(p)
}

func makeObj() {
	// struct 필드값을 순서대로 지정
	var p1 person
	p1 = person{"Bob", 20}
	p2 := person{name: "Sam", age: 50}
	fmt.Println(p1, p2)
	// 객체 생성(new)
	p := new(person)
	p.name = "Lee"
}

func newDict() *dict {
	// 생성자 함수: 초기화된 dict를 반환
	d := dict{}
	d.data = map[int]string{}
	return &d // 포인터 전달
}
