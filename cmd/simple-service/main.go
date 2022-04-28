package main
//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 => github token
import (
	"fmt"
)
type workers interface {
	cal_salary() float64
	// printer()
}

type developer struct {
	id int
	name string
	salary int
	tax float64
}

type designer struct {
	id int
	name string
	salary int
	expense float64
}

type ali int

func (a ali) cal_salary() float64 {
	return 0.5 * float64(a) * 5000
}
var x = ali(5)
func cal_salary() workers {
	return x
}
func (a developer) cal_salary() float64 {
	taxes := a.tax
	roi := 0.5
	result := float64(a.salary) * taxes * roi
	return result
}

func (a designer) cal_salary() float64 {
	return float64(a.salary) - a.expense
}
func hello(a workers) {
	b := a

	fmt.Println(b.cal_salary())
}
func main() {
	// var obj workers
	a := developer{95,"vijay",45000,0.9}
	b := designer{100,"arun",64000,4512.6}
	c := ali(100)
	var obj1 = []workers{a,b,c}
	// obj1 = []workers{a,b}
	for _,v := range obj1 {
		res := v.cal_salary()
		fmt.Println(res)
	}
	hello(a)
	fmt.Println(cal_salary())
	}