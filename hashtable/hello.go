package main

import "fmt"

func main() {
	hashMap()
	appendMap()
	checkKey()
	forMap()
}

func hashMap() {
	var idMap map[int]string     // map[keyType]valueType
	idMap = make(map[int]string) // map 초기화
	fmt.Println(idMap)           // nil map
	tickers := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB":   "FaceBook", // , 붙여야함
	}
	fmt.Println(tickers)
}

func appendMap() {
	// 변수 선언
	var m map[int]string
	// map 초기화
	m = make(map[int]string)
	// 값 추가
	m[901] = "Apple"
	m[134] = "Grape"
	m[777] = "Tomato"
	// 값 일기
	str := m[134]
	println(str)
	noData := m[999]
	println(noData)
	delete(m, 777)
	fmt.Println(m)
}

func checkKey() {
	// 특정 key가 존재하는지 체크
	tickers := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB":   "FaceBook",
	}
	val, exists := tickers["MSFT"] // Microsoft, true
	if !exists {
		println("No MSFT ticker")
	}
	println(val, exists)
}

func forMap() {
	// range : key, value 뽑음
	newMap := map[string]string{
		"A": "Apple",
		"B": "Banana",
		"C": "Charlie",
	}
	for key, val := range newMap {
		fmt.Println(key, val)
	}
}
