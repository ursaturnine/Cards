package main

/*
	Create a variable and assign it a string that represents a playing card.

use fmt package to print out string
*/
func main() {
	//var | name | type
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}
