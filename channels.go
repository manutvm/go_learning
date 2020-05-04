package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 35
		ch <- 36
	}()

	val := <-ch
	fmt.Printf("Value got from channel %v\n", val)
	val1 := <-ch
	fmt.Printf("Value got from channel %v\n", val1)
}
