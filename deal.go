package main

/**
* Given a deck, and number of cards, returns a deck containing the desired number
* of cards. A second deck is returned containing the remaining cards.
**/
func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]
}
