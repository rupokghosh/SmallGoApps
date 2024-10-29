package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	testDeck := newDeck()

	if len(testDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(testDeck))
	}
	if testDeck[0] != "Diamonds of Ace" {
		t.Errorf("Expected Diamond of Ace, got %v", testDeck[0])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52{
		t.Errorf("Expected deck length of 52, but got %d", len(loadedDeck))
	}

	os.Remove("_decktesting")

}
