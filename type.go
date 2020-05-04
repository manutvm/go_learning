package main

import (
	"fmt"
)

func main() {
	var a float64
	var b float64

	a = 1
	b = 2

	var mean float64

	mean = (a + b) / 2
	fmt.Printf("Variable a= %v Type= %T\n", a, a)
	fmt.Printf("Variable b= %v Type= %T\n", b, b)
	fmt.Printf("mean of %v and %v is %v\n", a, b, mean)
}
