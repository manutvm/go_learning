package main

import (
	"fmt"
	"time"
)

func main() {
	out := make(chan float64)

	go func() {
		time.Sleep(100 * time.Millisecond)
		out <- 3.14
	}()

	select {
	case val := <-out:
		fmt.Println(val)
	case <-time.After(50 * time.Millisecond):
		fmt.Println("Timeout")
	}
}
