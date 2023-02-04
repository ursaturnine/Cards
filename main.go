package main

/*
	Create a variable and assign it a string that represents a playing card.

use fmt package to print out string
*/
func main() {
	//var | name | type
	cards := newDeck()

	// the two return values from deal
	// will be assigned to vars hand and
	// remainingCards respectively
	hand, remainingCards := deal(cards, 5)

	// hand and remainingCards
	// are both of type deck
	hand.print()
	remainingCards.print()

}
