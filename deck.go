package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// Create a new function for the type "deck"
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// _ is a placeholder for a value that we don't care about
	// in this case, we don't care about the index
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			// append to the slice
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// any variable of type "deck" now gets access to the "print" method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	// slice[startIndexIncluding : upToNotIncluding]
	return d[:handSize], d[handSize:]
}
