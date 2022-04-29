package main

import (
	"fmt"
)

func main() {
	var in = map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000}
	data := "MCMIV"
	res := 0
	l := len(data)
	for i := 0; i < l; i++ {
		start := in[string(data[i])]
		next := in[string(data[i+1])]
		if start >= next {
			res += start
		} else {
			res += next - start
			i++
		}
	}
	fmt.Println(res)
}

/*
SYMBOL       VALUE
  I            1
  IV           4
  V            5
  IX           9
  X            10
  XL           40
  L            50
  XC           90
  C            100
  CD           400
  D            500
  CM           900
  M            1000

  it is sum of values
  like XI is 5 +1 = 6, but with one condition
  it is when  next digit's value is greater than current digit's value
  eg: IX is 1 > 5 , so 5-1 = 4

  I placed before V or X indicates one less, so four is IV (one less than 5) and 9 is IX (one less than 10).
X placed before L or C indicates ten less, so forty is XL (10 less than 50) and 90 is XC (ten less than a hundred).
C placed before D or M indicates a hundred less, so four hundred is CD (a hundred less than five hundred) and nine hundred is CM (a hundred less than a thousand).

just subtract the values and then add to sum if ith value is < i + 1 th value or else just sum all
if you subtract make sure you ignore the next value since you already utilised it, so increment i++

*/
