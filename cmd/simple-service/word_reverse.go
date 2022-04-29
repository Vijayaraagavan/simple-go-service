package main

import (
	"fmt"
	"strings"
)

func main() {
	// To get input from user use bufio package
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// in := scanner.Text()
	in := "i am the iking"
	fmt.Println(in)
	// sli := make([]string, 20)
	sli := strings.Split(in, " ")

	last := len(sli) - 1
	fmt.Println(last, sli[1])
	var out []string
	for i, _ := range sli {
		out = append(out, sli[last-i])
	}
	str := strings.Join(out, " ")
	fmt.Println(strings.Trim(str, " "))

}
