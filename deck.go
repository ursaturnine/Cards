package main

import (
	"fmt"
	"strings"
)

// Create a new type of 'deck
// which is a slice of string

// deck extends a list of strings
type deck []string

// receiver function
// "deck" a reference to the type deck
// "d" a reference to the copy of array
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Create and return a list of playing cards
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// Each suit gets each number
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

// Create a 'hand' of cards
// takes arguments d, a deck, and handSize, an int
// two values of type deck returned
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Write File function to save a deck to our hard drive

// use "strings" package to encode a deck into a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
