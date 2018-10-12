package main

import "fmt"

/**
 * Code to create and manipulate a deck.
**/

func main() {

	cards := newDeck()
	numCards := 5

	cards.saveToFile("my_cards")
	newDeckFromFile("my_cards")

	fmt.Printf("\nDECK BEFORE SHUFFLING: %s.\n", cards.printAsFormattedString())
	cards.shuffle()
	fmt.Printf("\nDECK AFTER SHUFFLING: %s.\n", cards.printAsFormattedString())

	dl, card := deal(cards, numCards)
	fmt.Printf("\nDECK CONTAINING %d CARD(S): %s.\nREMAINING DECK OF CARDS: %s.\n",
		numCards, dl.printAsFormattedString(), card.printAsFormattedString())

	removeFile("my_cards")
}
