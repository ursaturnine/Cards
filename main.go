package main

/*
	Create a variable and assign it a string that represents a playing card.

use fmt package to print out string
*/
func main() {
	//var | name | type
	cards := deck{"Ace of Diamonds", newCard()}
	//creates a new slice
	cards = append(cards, "Six of Spades")

	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
