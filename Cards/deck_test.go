package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but got %v", len(d))}
	// } else{
	// 	t.Errorf("success. deck length is %d", len(d))
	// }

	if d[0] != "Ace of Spades"{
		t.Errorf("Expected first card Ace of Spades but got %v", d[0])
	} 
	// else{
	// 	t.Errorf("Successful. first card is %v", d[0])
	// }

	if d[len(d)-1] != "Four of hearts"{
		t.Errorf("Expected last card Ace of clubs but got %v", d[len(d)-1])
	} 
	// else{
	// 	t.Errorf("Successful. Expected last card is %v", d[len(d)-1])
	// }
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newReadFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

}