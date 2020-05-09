package main

import (
	"fmt"
	"strings"
)

type deck []string

// print prints all cards in a deck (receiver function example)
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// createdeck creates new card deck
func createdeck() deck {
	num := []string{"1", "2", "3"}
	suit := []string{"spades", "hearts"}
	var cards deck
	for _, n := range num {
		for _, s := range suit {
			cards = append(cards, n+" of "+s)
		}
	}
	return cards
}

// deal returns dealt and remaining decks
func deal(n int, d deck) (deck, deck) {
	return d[:n], d[n:]
}

// tostring converts deck to string
func (d deck) tostring() string {
	return strings.Join(d, ", ")
}
