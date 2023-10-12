package main

func main() {

	cards := newDeck()

	//hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	cards.saveToFile("my_cards")

	otherCards := newDeckFromFile("./my_cards")

	otherCards.shuffleDeck()

	otherCards.print()
}
