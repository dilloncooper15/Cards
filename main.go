package main

import "fmt"

/**
 * Code to create and manipulate a deck
**/

func main() {

	cards := newDeck()

	cards.saveToFile("my_cards")

	// fmt.Println(newDeckFromFile("my_cards"))

	cards.shuffle()

	fmt.Println(cards)
}
