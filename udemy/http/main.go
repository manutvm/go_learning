package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, error := http.Get("http://google.co.in")

	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)
	}

	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}
