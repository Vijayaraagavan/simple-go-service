package main

import (
	"fmt"
	"Crypto/rand"
	"time"
	"math/big"
	"encoding/base32"
)
func gen() []byte {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err!=nil {
		panic(err)
	}
	return b
}

func main() {
	start := time.Now()
	fmt.Println("Crypto random")
	for i:=0; i<=10; i++ {
		data , _ := rand.Int(rand.Reader, big.NewInt(10000000))
		fmt.Println(data)
	}
	end := time.Now()
	fmt.Println(start.Sub(end))
	a := gen()
	fmt.Println(len(a), string(a))
	fmt.Println(base32.StdEncoding.EncodeToString(a))
	// i:=0
	// for {
	// 	if i>5 {
	// 		break;
	// 	}
	// 	fmt.Println(rand.Reader)

	// 	i++
	// }
}
