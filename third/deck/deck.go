package deck

import "fmt"

// Create a new type of 'deck' which is a slice of strings
type deck []string

// Fill the deck with cards
func New() deck {
	var Suits = []string{"Diamonds", "Hearts", "Clubs", "Spades"}
	var Ranks = []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	cards := deck{}
	for _, suit := range Suits {
		for _, rank := range Ranks {
			cards = append(cards, rank+" of "+suit)
		}
	}
	return cards
}

// Splits the deck into two hands, the first of which is the specified size and the second of which is the remainder
func Deal(d deck, leftHandSize int) (deck, deck) {
	return d[:leftHandSize], d[leftHandSize:]
}

// Print the contents of the deck to the console
func (d deck) Print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
