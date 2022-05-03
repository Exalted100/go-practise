package main

import "fmt"

type deck []string

func newDeck() deck {
	values := deck{"Ace", "Two", "Three", "Four"}

	suits := deck{"Spades", "Diamonds", "Clubs", "Hearts"}

	var cards deck

	for _, value := range values {
		for _, suit := range suits {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}
