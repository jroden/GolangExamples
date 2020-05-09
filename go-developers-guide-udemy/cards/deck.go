package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

// savetofile saves deck in byte slice to file
func (d deck) savetofile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.tostring()), 0666)
}

// newdeckfromfile reads deck from file
func newdeckfromfile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil { // error handling
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ", "))
}

// shuffledeck returns a 'shuffled' deck of cards
func (d deck) shuffledeck() {
	// random function seed based on time
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newpos := r.Intn(len(d) - 1)
		d[i], d[newpos] = d[newpos], d[i]
	}
}
