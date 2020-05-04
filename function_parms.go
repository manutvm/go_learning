package main

import (
	"fmt"
)

func double(values []int) {
	for i, value := range values {
		values[i] = value * 2
	}
}

func doubleValue(value int) {
	value = value * 2
}

func doubleValuePtr(value *int) {
	*value = *value * 2
}

func main() {
	slice := []int{1, 2, 3, 4}
	fmt.Println(slice)

	double(slice)

	fmt.Println(slice)

	n := 3
	doubleValue(n)

	fmt.Println(n)

	doubleValuePtr(&n)

	fmt.Println(n)
}
