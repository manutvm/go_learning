package main

import "fmt"

func main() {
	//deck := newDeck()
	//deck.print()
	//deck.saveToFile("deck.txt")
	deck := newDeckFromFile("deck.txt")
	deck.shuffle()
	deck.print()

	deckHand, deckRemaining := deal(deck, 4)
	fmt.Println("Deck Hand")
	deckHand.print()
	fmt.Println("Deck Remaining")
	deckRemaining.print()
}
