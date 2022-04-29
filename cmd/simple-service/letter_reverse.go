package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	in := scanner.Text()
	fmt.Println(in)
	var res []rune
	for _, val := range in {
		res = append(res, val)
		copy(res[1:], res)
		res[0] = val
	}
	var out string = ""
	for _, v := range res {
		out += string(v)
		fmt.Print(string(v))
	}
	fmt.Println(out)
}
