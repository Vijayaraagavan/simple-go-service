package main

import (
	"fmt"
	// "os/signal"
	"os"
	// "io"
	"github.com/gocarina/gocsv"
	// "strings"
	// "log"
	// "io/ioutil"
	// "encoding/csv"
	// "syscall"
	"time"
)
type Client struct {
	Year string `csv:"Year"`
	Aggr string `csv:"Aggr"`
	Code string `csv:"Code"`
	Industry_name string `csv:"Industry_name"`
	Units string `csv:"Units"`
	V_code string `csv:"V_code"`
	V_name string `csv:"V_name"`
	V_cat string `csv:"V_cat"`
	Value string `csv:"Value"`
	Industry_code string `csv:"Industry_code"`

}

func main() {
	// f,err := os.OpenFile("annual.csv",os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	// if err!=nil {
	// 	fmt.Println(err)
	// }
	// defer f.Close()
	// content,err := ioutil.ReadAll(f)
	// fmt.Println(content)
	// if err!=nil {
	// 	fmt.Println(err)
	// }
	// content := `firstName, lastName, age
	// Celina, Jones, 18 
	// Cailyn, Henderson, 13 
	// Cayden, Smith, 42
	// `
	// r := csv.NewReader(strings.NewReader(string(content)))
	// header := true
	// for i:=0; i<50; i++ {
	// 	record, err := r.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err !=nil {
	// 		log.Fatal(err)
	// 	}
	// 	if !header {
	// 		for i, v := range record {
	// 			switch i{
	// 			case 0:
	// 				fmt.Println("YEAR: ",v)
	// 			case 1:
	// 				fmt.Println("Industry_name: ",v)
	// 			case 2:
	// 				fmt.Println("Units: ",v)
	// 			}
	// 		}
	// 	}
	// 	fmt.Println("<--------------------->")
	// 	header = false
	// }
	s := time.Now()
	f, err := os.OpenFile("annual.csv",os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err!=nil {
		panic(err)
	}
	clients := []*Client{}
	if err:= gocsv.UnmarshalFile(f,&clients); err !=nil {
		panic(err)
	}
	count:=0
	for _, c := range clients {
		fmt.Println("Your name is: ", c.Year)
		fmt.Println("your age is: ", c.Industry_name)
		fmt.Println("<---------------------------->")
		count++
	}
	s2 := time.Now()
	res := s2.Sub(s)
	fmt.Println("Total rows : ", count)
	fmt.Println("Started at : ",s)
	fmt.Println("Ended at : ",s2)
	fmt.Println("Total time : ",res)

}