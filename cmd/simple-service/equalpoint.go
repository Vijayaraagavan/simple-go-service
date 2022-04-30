package main

import "fmt"

func main() {
	in := "()()())((())"
	l := len(in)
	left, right := 0, 0
	for i := 0; i < l; i++ {
		if string(in[i]) == "(" {
			left++
		}
	}
	for i := l - 1; i >= 0; i-- {
		if left == right {
			i++
			fmt.Println(i)
			break
		}
		if string(in[i]) == ")" {
			right++
		} else {
			left--
		}
	}
	fmt.Println("no equal point")
}

/*
Given a string S of opening and closing brackets '(' and ')' only. The task is to find an equal point. An equal point is
an index such that the number of closing brackets on right from that point must be equal to number of opening brackets before that point.
*/
