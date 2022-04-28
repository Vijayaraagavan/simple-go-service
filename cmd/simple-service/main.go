package main

import (
	"fmt"
)
func main(){
	// x := 5
	var slice = []string{"golang","javascript", "bootstrap"}
	for i, val := range slice {
		fmt.Printf("The %dst value is : %s\n",i+1,val)
	}
}