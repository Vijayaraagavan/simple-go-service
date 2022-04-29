package main

import (
	"fmt"
)
var y int =0
func test(x chan int) {
	for i:=0; i<1000000; i++ {
		y += i
	}
	data := <- x
	fmt.Println(data,y)
	fmt.Println(x)
}
func main(){
	data := make(chan int,3);
	// fmt.Printf("%T",data)
	go test(data)
	go test(data)
	data <- 10
	// data <- 15
	// x := <- data
	// go test(data)
	fmt.Println("main ends")
}
// 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