package main

func main() {
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()

	// remainingCards.print()

	// greeting := "Hi there!"
	// fmt.Println(greeting)
	// fmt.Println([]byte(greeting))

	// cards := newDeckFromFile("my_cards")
	// fmt.Println(cards)

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
