package main

import "fmt"

/**
 * Create a new type of 'deck', which is a slice of strings
**/

type deck []string

func newDeck() deck {
	cardSuits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six",
		"Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	cards := deck{}

	for i := range cardSuits {
		for m := range cardValues {
			cards = append(cards, fmt.Sprintf("%v of %v", cardValues[m], cardSuits[i]))
		}
		cards = append(cards)
	}

	return cards
}

func (d deck) print() { //Receiver
	for i, card := range d {
		fmt.Println(i, card)
	}
}
