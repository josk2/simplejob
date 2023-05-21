package main

import (
	"log"
	"testing"
)

type TB struct {
	testing.B
}

// phương thức thuộc struct TB
//func (p *TB) Fatal(args ...interface{}) {
//	//fmt.Println("TB.Fatal disabled!")
//}

func main() {
	// khởi tạo một đối tượng thuộc interface testing.TB
	var tb testing.TB = new(TB)

	// lúc này nó có thể sử dụng phương thức Fatal mà TB đã hiện thực
	//tb.Fatal("Hello, playground")
	log.Fatalln(tb)
}
