package main

import (
	"fmt"
	"strconv"
)

func rev(r string) string {
	l := len(r)
	out := ""
	for i := l - 1; i >= 0; i-- {
		out += string(r[i])
	}
	return out
}

func main() {
	str := "aaaaaaaaaaa"
	out := ""
	l := len(str)
	var count int64 = 0
	for i := 0; i < l; i++ {
		if i == l-1 {
			if str[i] == str[i-1] {
				count++
			}
			if count > 0 {
				out += string(str[i])
				out += strconv.FormatInt(count, 16)
				break
			} else {
				out += string(str[i])
				out += "1"
				break
			}
		}
		if str[i] == str[i+1] {
			count++
			continue
		}
		if count > 0 {
			out += string(str[i])
			out += strconv.FormatInt(count, 16)
		} else {
			out += string(str[i])
			out += "1"
			count = 0
		}

	}
	fmt.Println(out)
	fmt.Println(rev(out))
}

/*
You are given a string S. Every sub-string of identical letters is replaced by a single instance of that 
letter followed by the hexadecimal representation of the number of occurrences of that letter. 
Then, the string thus obtained is further encrypted by reversing it 
 note : All Hexadecimal letters should be converted to Lowercase letters
 
