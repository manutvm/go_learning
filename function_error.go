package main

import (
	"fmt"
	"math"
)

func sqrt(val float64) (float64, error) {
	if val < 0 {
		return 0.0, fmt.Errorf("sqrt of %f cannot be calculated", val)
	}

	return math.Sqrt(val), nil
}

func main() {
	val1 := 2.0

	s, err := sqrt(val1)

	if err == nil {
		fmt.Println(s)
	} else {
		fmt.Println(err)
	}

	val2 := -2.0

	s1, err1 := sqrt(val2)

	if err1 == nil {
		fmt.Println(s1)
	} else {
		fmt.Println(err1)
	}
}
