package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got: %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got: %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got: %v", d[len(d)-1])
	}
}

func TestNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")
	d := newDeck()
	d.saveToFile("_deckTesting")
	loadedDeck := newDeckFromFile("_deckTesting")
	if loadedDeck == nil || len(loadedDeck) != 16 {
		t.Errorf("Expected 16 card in deck, but got: %v", len(loadedDeck))
	}

	os.Remove("_deckTesting")
}
