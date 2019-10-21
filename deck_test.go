package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected 'Ace of Spades', instead got %s", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected 'Four of Clubs', instead got %s", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, for %d", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
