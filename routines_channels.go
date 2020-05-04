package main

import (
	"fmt"
	"net/http"
)

func returnType(url string, ch chan string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	ch <- fmt.Sprintf("%s -> %s\n", url, ctype)
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	ch := make(chan string)
	for _, url := range urls {
		go func(url string) {
			returnType(url, ch)
		}(url)
	}

	for range urls {
		out := <-ch
		fmt.Println(out)
	}
}
