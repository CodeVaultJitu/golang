package main

func main() {
	cards := newDeck()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// 	remainingCards.print()
	// fmt.Println((cards.joinString()))
	// cards.saveToFile("my_cards")
	// cards := newReadFile("my_cards")
	cards.shuffle()
	cards.print()
}
