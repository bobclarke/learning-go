package main

import "testing"

func TestNewDeck(t *testing.T) {

	// Create a new deck
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected length of 16, got %v", len(d))
	}

}
