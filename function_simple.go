package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func divmod(a int, b int) (int, int) {
	return a / b, a % b
}

func main() {
	a, b := 7, 2
	sum := add(a, b)

	fmt.Printf("sum of %v and %v is %v\n", a, b, sum)

	div, mod := divmod(a, b)

	fmt.Printf("Divide= %v Mod= %v\n", div, mod)
}
