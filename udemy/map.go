package main

import "fmt"

func main() {
	colors := make(map[string]string)
	colors["white"] = "#fffff"
	colors["black"] = "#00000"
	colors["red"] = "#f43000"

	addColor(colors)
	printHex(colors)
}

func addColor(c map[string]string) {
	c["blue"] = "#b34210"
}

func printHex(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex value of", color, "is", hex)
	}
}
