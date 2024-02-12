package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

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

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func deal(d deck, handSize int) (deck, deck) {
	// slice[startIndexIncluding : upToNotIncluding]
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(filename string) error {
	// 0666 is everyone can read and write to the file
	return os.WriteFile(filename, []byte(d.toString()), 0666)

}

func newDeckFromFile(fileName string) deck {
	// bs is a slice of bytes
	// errs is an error
	bs, err := os.ReadFile(fileName)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error:", err)
		// exit the program
		// 0 means everything went well
		// 1 means there was an error
		os.Exit(1)
		//return newDeck()
	}

	// convert the byte slice to a string
	s := strings.Split(string(bs), ",")
	return deck(s)

}

func (d deck) shuffle() {
	for i := range d {
		// generate a random number
		newPosition := rand.Intn(len(d) - 1)

		// swap the elements
		d[i], d[newPosition] = d[newPosition], d[i]
	}

}
