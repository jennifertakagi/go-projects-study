package main

func main() {
	cards := newDeckFromFile("my_cards")
	//cards := newDeck()
	//cards.saveToFile("my_cards")

	cards.shuffle()
	cards.print()

	// hand, remainingCards := deal(cards, 5)
}
