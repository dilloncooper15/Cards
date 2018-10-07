package main

/**
 * Code to create and manipulate a deck
**/

func main() {

	cards := deck{newCard(), newCard()}
	cards = append(cards, "Six of Spades")

	// cards.print()
	newDeck().print()
}

func newCard() string {
	return "Five of Diamonds"
}
