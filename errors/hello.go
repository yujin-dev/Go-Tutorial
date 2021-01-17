package main

import (
	"log"
	"os"
)

type error interface {
	Error() string
}

func main() {
	f, err := os.Open("C/temp/1.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	println(f.Name())
}

// 	_, err := otherFunc()
// 	switch err.(type) { // error type 체크해서 error case 처리
// 	default:
// 		println("OK")
// 	case MyError:
// 		log.Print("Log my error")
// 	case error:
// 		log.Fatal(err.Error())
// 	}
