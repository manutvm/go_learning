package main

import (
	"fmt"
)

func main() {
	str := "The color of magic"

	fmt.Println(str)
	fmt.Println(len(str))
	fmt.Printf("book[0] = %v (type %T)\n", str[0], str[0])
	fmt.Printf("%v\n", str[:3])
	fmt.Printf("t%v\n", str[1:])
	fmt.Println("t" + str[1:])

	multiline := `
	This is a ssample test
	with multiline
	`

	fmt.Println(multiline)
}
