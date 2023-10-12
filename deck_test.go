package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) { //test is automatically ran with a test runner with the given parameter format

	d := newDeck()

	if len(d) != 52 {

		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Two of Spades" {
		t.Errorf("Expected Two of Spades, but got %v", d[0])
	}

}

func TestSaveToDeck(t *testing.T) {
	os.Remove("_decktesting") //remove dirty test files

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, but got %v", len(loadedDeck))
	}
}

func TestDeal(t *testing.T) {

	hand, remainingCards := deal(newDeck(), 7)

	if len(hand) != 7 {

		t.Errorf("Expected 7 cards in hand, but got %v", len(hand))
	}

	if len(remainingCards) != 45 {
		t.Errorf("Expected 45 cards in remaining deck, but got %v", len(remainingCards))
	}
}
