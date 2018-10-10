package main

func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]
}
