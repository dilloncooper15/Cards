package main

import "fmt"

/**
 * Code to create and manipulate a deck
**/

func main() {

	cards := newDeck()

	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()

	// fmt.Println(cards.toString())
	// fmt.Println([]byte(cards.toString()))
	cards.saveToFile("my_cards")
	fmt.Println(newDeckFromFile("my_cards"))

}
