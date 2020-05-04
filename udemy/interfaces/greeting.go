package main

import "fmt"

type Bot interface {
	getGreeting() string
}

type EnglishBot struct{}
type SpanishBot struct{}

func (EnglishBot) getGreeting() string {
	return "Hi there!"
}

func (SpanishBot) getGreeting() string {
	return "Hola!"
}

func (SpanishBot) getMessage() string {
	return "Halla!"
}

func printGreeting(b Bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := EnglishBot{}
	sb := SpanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}
