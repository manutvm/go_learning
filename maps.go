package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`
	word_count := map[string]int{}
	words := strings.Fields(text)

	for _, word := range words {
		word_count[strings.ToLower(word)] = word_count[strings.ToLower(word)] + 1
	}

	fmt.Println(word_count)
}
