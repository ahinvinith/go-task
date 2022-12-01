package main

func main() {
	// card := newDeck()

	// hand, remainingCard := deal(card, 5)

	// hand.printDeck()
	// remainingCard.printDeck()

	//Writing a file
	// card := newDeck()

	// card.saveToFile("my_cards")

	// card := newDeckFromFile("my_cards")
	// card.printDeck()

	card := newDeck()
	card.shuffle()
	card.printDeck()
}
