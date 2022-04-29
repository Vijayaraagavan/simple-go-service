package main

import (
	"fmt"
	// "time"
	"log"
)
func count(a,b int, ch chan string) {
	fmt.Println("waiting: ",<-ch)
	fmt.Println("enter data send")
	for i:=a; i<b; i++ {
		log.Println(i)
	}
	// ch <- "hello"
	// time.Sleep(time.Second + 1)
	// fmt.Println("ending")
}

func cou(a,b int, ch chan string) {
	fmt.Println("enter data send")
	for i:=a; i<b; i++ {
		log.Println(i)
		if i==450 {
			ch <- "ok"
		}
	}
	// ch <- "hello"
	// time.Sleep(time.Second + 1)
	// fmt.Println("ending")
}

func main() {
	var data = make(chan string)
	go cou(0,500,data)
	count(51,100,data)
	// time.Sleep(time.Second +2)
	// fmt.Println(res)
	// fmt.Println("enter data request")
	// fmt.Println(<-data)
	// time.Sleep(time.Second)

}
// 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