package main

import (
	"fmt"
	"sort"
)

/* program to find longest common prefix in an array
   so the concept is simple
   sort() will arrange words in small to large both in len and dictioinay order
   if smallest lenth word is of size of 4, then the prefix can be withing this limit only
   alphabetically last word will have most difference ( eg: geeks vs geezer, z is so diff)
   if last word has no match then even if other words had match, it is useless

*/

func main() {

	var in = []string{"geeksforgeeks", "geeks", "geek", "geezer"}
	sort.Strings(in)
	for _, v := range in {
		fmt.Println(v)
	}
	start := in[0]
	end := in[len(in)-1]
	l := len(start)
	out := ""
	for i := 0; i < l; i++ {
		if start[i] == end[i] {
			out += string(start[i])
		}
	}
	fmt.Println(out)
}
