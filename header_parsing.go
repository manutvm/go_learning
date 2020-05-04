package main

import (
	"fmt"
	"net/http"
)

func contentType(url string) (string, error) {
	content, err := http.Get(url)

	fmt.Println(content)
	if err != nil {
		return "", err
	}

	defer content.Body.Close()

	ctype := content.Header.Get("Content-Type")
	fmt.Println(ctype)

	if ctype == "" || ctype == nil {
		return "", fmt.Errorf("Cannot find the Content-Type header")
	}

	return ctype, nil
}

func main() {
	ctype, err := contentType("https://google.com")

	if err != nil {
		fmt.Println("Got content type")
		fmt.Println(ctype)
	} else {
		fmt.Println(err)
	}
}
