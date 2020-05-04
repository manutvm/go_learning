package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://stackoverflow.com",
		"http://amazon.com",
	}

	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(url string, c chan string) {
	_, err := http.Get(url)

	if err != nil {
		fmt.Println(url, "is NOT UP")
		c <- (url + " is NOT UP")
	}

	fmt.Println(url, "is UP")
	c <- (url + " is UP")
}
