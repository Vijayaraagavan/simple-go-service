package main

import "fmt"

func main() {
	s := "Bawds jog, flick quart, vex nymp"
	ch := map[rune]int{}
	count := 0
	for _, v := range s {
		if ch[v] == 0 && (v >= 65 && v <= 90) || (v >= 97 && v <= 122) {
			count++
			ch[v]++
		}
	}
	fmt.Println(count)
	if count > 25 {
		fmt.Println("It is a pangram string")
	} else {
		fmt.Print("Not a pangram string")
	}
}

/*
Given a string check if it is Pangram or not. A pangram is a sentence containing every letter in the English Alphabet.
*/
