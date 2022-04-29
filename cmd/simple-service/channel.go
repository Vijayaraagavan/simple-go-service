package main

import (
	"fmt"
)

func prepare(data chan int) {
	fmt.Println("starting data input")
	for i:=0; i<=9; i++ {
		fmt.Printf("data no: %d\n",i)
		data <- i*i
	}
	close(data)
	fmt.Println("ending channel")
}

func main() {
	fmt.Println("main started")

	a := make(chan int)

	go prepare(a)

	for val := range a {
		fmt.Println("enter data read: ",val)
	}

	fmt.Println("main stopped")
}