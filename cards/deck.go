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

// Creates a new type of 'deck' which is a slice of strings
func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Jack", "Queen", "King"}

	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Prints the given deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Deals the hand according to the hand's size
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Transform the deck into string type
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Saves the deck on a file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Gets the deck from a file
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")

	return deck(s)
}

// Shuffles the deck
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
