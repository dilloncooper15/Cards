package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

/**
 * Create a new type of 'deck', which is a slice of strings
**/

type deck []string

/**
* Generates a two separate slices of type string and appends the result to
* a deck. Returns a deck.
**/
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six",
		"Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

/**
* A Receiver Function that prints a value of type deck to the console.
**/
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

/**
* A Receiver Function that, when given a deck, returns a slice containing a single
* string of cards, where each card is seperated by a comma.
**/
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

/**
* A Receiver Function that converts a deck to a byte slice and writes the result
* to the provided file name passed as a parameter.
**/
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

/**
* Reads in the contents of the file [filename] and returns a deck. If the [filename]
* cannot be found, an error will be logged and the program will exit with code 1.
**/
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	ss := strings.Split(string(bs), ",")
	return deck(ss)
}

/**
* A Receiver Function that randomly reorganizes the order of a deck.
**/
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
