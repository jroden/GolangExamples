package main

import "fmt"

func main() {

	cards := createdeck()

	// receiver function example
	// cards.print()

	// deal example
	// hand, remaining := deal(3, cards)
	// fmt.Println("cards in hand: ")
	// hand.print()
	// fmt.Println("cards remaining: ")
	// remaining.print()

	// deck to string example
	// fmt.Println(cards.tostring())

	// save deck to file
	// cards.savetofile("outputfile")

	// new deck from file
	// newdeck := newdeckfromfile("outputfile")
	// newdeck.print()

	// shuffledeck example
	cards.print()
	cards.shuffledeck()
	fmt.Println("printing shuffled deck")
	cards.print()

}
