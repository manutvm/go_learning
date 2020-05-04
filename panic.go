package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("app.lo")

	if err != nil {
		panic(err)
	}

	defer file.Close()
	fmt.Println("file is opened")
}
