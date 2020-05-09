package main

import "fmt"

func main() {
	// range example
	// cards := deck{"one", "two"}
	// for index, card := range cards {
	// 	fmt.Println(index, card)
	// }

	// receiver function example
	cards := createdeck()
	cards.print()

	// deal example
	hand, remaining := deal(3, cards)
	fmt.Println("cards in hand: ")
	hand.print()
	fmt.Println("cards remaining: ")
	remaining.print()

	// deck to string example
	fmt.Println(cards.tostring())
}
