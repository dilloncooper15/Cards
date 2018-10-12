package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Clubs" {
		t.Errorf("Expected first card to be 'Ace of Clubs', but got %v", d[0])
	}

	if d[len(d)-1] != "King of Spades" {
		t.Errorf("Expected last card to be 'King of Spades', but got %v", d[0])
	}
}

func TestDeal(t *testing.T) {
	d := newDeck()
	hand, remainingDeck := deal(d, 5)

	if len(hand) != 5 || len(remainingDeck) != 47 {
		t.Errorf("Expected handsize to be 5, but got %v. Expected remainingDeck size to be 47, but got %v",
			len(hand), len(remainingDeck))
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	removeFile("_decktesting")
	d := newDeck()

	d.saveToFile("_decktesting")
	ld := newDeckFromFile("_decktesting")

	if len(ld) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(ld))
	}

	removeFile("_decktesting")
}
