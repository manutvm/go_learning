package main

import (
	"fmt"
)

func main() {
	a := 2
	switch a {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	}

	switch {
	case a > 10:
		fmt.Println("a is big")
	case a > 100:
		fmt.Println("a is very big")
	default:
		fmt.Println("a is small")
	}
}
