package main

import "fmt"

func main() {

	// create new deck
	fmt.Println("step 1: add new deck function")
	cards := newDeck()
	fmt.Println("step 1 is done")

	// print deck values
	fmt.Println("step 2: print list of cards in deck")
	cards.print()
	fmt.Println("step 2 is done")

	// deal function
	fmt.Println("step 3: deal function")
	hand, leftHand := deal(cards, 3)
	fmt.Println(hand)
	fmt.Println(leftHand)
	fmt.Println("step 3 is done")

	// save to file
	fmt.Println("step 4: save deck to file func")
	cards.saveToFile("deck_file")
	fmt.Println("step 4 is done")

	// read deck from file
	fmt.Println("step 5: read deck from file")
	cards = newDeckFromFile("deck_file")
	cards.print()
	fmt.Println("step 5 is done")

	// shuffle cards
	fmt.Println("step 6: run shuffle operation")
	cards.shuffle()
	cards.print()
	fmt.Println("step 6 is done")

	// finish the game
	fmt.Println("Game Over")

}
