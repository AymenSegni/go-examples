package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// declare new type called deck and assign to it a slice type

type deck []string

// func with receiver to iterate over a deck and print values

func (d deck) print() {
	for i, value := range d {
		fmt.Println(i, value)
	}
}

// Create and return newDeck

func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Tow", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards

}

// deal

func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]
}

// save deck to file and load a deck from file
//  to do: use the ioutil library
// to work with files in the ioutil library, will need to have a []byte instead
// will also need to import String package

// fun to convert a deck to string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
	// []string(d) // convert deck to slice of string

}

func (d deck) saveToFile(filename string) error {

	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)

}

// read deck from file
func newDeckFromFile(filename string) deck {

	bs, err := ioutil.ReadFile(filename)

	// error handling
	if err != nil {
		fmt.Println(err)
		//quite
		os.Exit(127)
	}

	// convert byte slice into string then to string slice to deck

	s := strings.Split(string(bs), ",")

	return deck(s)

}

// shuffle list of deck

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())

	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]

	}
}
