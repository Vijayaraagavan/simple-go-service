package main

import "fmt"

func main() {
	x := "fodrf"
	y := "gorkl"
	k := 2
	h1 := map[byte]int{}
	h2 := map[byte]int{}
	count := 0
	if len(x) == len(y) {
		fmt.Println("not a k-anagrams")
		return
	}
	for i := 0; i < len(x); i++ {
		h1[x[i]]++
		h2[y[i]]++
	}
	for i := 0; i < len(x); i++ {
		if h1[x[i]]-h2[y[i]] > 0 {
			count += h1[x[i]] - h2[y[i]]
		}
	}
	if count <= k {
		fmt.Println("both are k-anagrams")
	} else {
		fmt.Println("not possible")
	}
}

/*
Given two strings of lowercase alphabets and a value K, your task is to complete the given function which tells if  two strings are K-anagrams of each other or not.

Two strings are called K-anagrams if both of the below conditions are true.
1. Both have same number of characters.
2. Two strings can become anagram by changing at most K characters in a string.
*/
