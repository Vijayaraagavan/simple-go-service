package main

import "fmt"

func main() {
	x := "ACAB"
	y := "ACAA"
	res := true
	var out *bool = &res
	// *out = true
	data := map[byte]byte{}
	if len(x) != len(y) {
		fmt.Println("not a isomorphic string")
		// break
	}
	for i := 0; i < len(x); i++ {
		if data[x[i]] == 0 && data[y[i]] == 0 {
			data[x[i]] = y[i]
			data[y[i]] = x[i]
		}
		if data[x[i]] != 0 {
			if data[x[i]] != y[i] {
				*out = false
			}
		}
		if data[x[i]] == 0 && data[y[i]] != 0 {
			if data[y[i]] != x[i] {
				*out = false
			}
		}
	}
	if *out {
		fmt.Println("isomorphic string")
	} else {
		fmt.Println("Not a isomorphic")
	}
}

/*
Given two strings 'str1' and 'str2', check if these two strings are isomorphic to each other.
Two strings str1 and str2 are called isomorphic if there is a one to one mapping possible for every character of str1 to
every character of str2 while preserving the order.
Note: All occurrences of every character in ‘str1’ should map to the same character in ‘str2’
*/
