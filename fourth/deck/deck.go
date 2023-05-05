package deck

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck' which is a slice of strings
type deck []string

var suits = []string{"Diamonds", "Hearts", "Clubs", "Spades"}
var ranks = []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

// Set Suits
func SetSuits(suits_p []string) {
	suits = suits_p
}

// Set Ranks
func SetRanks(ranks_p []string) {
	ranks = ranks_p
}

// Return the Suits
func GetSuits() []string {
	return suits
}

// Return the Ranks
func GetRanks() []string {
	return ranks
}

// Fill the deck with cards (a deck must have more then one card)
func New() deck {
	cards := deck{}

	for _, suit := range suits {
		for _, rank := range ranks {
			cards = append(cards, rank+" of "+suit)
		}
	}
	return cards
}

// Splits the deck into two hands, the first of which is the specified size and the second of which is the remainder
func Deal(d deck, leftHandSize int) (deck, deck) {
	return d[:leftHandSize], d[leftHandSize:]
}

// Convert the deck to a string
func (d deck) ToString() string {
	return strings.Join([]string(d), ",")
}

// Save the deck to a file
func (d deck) SaveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.ToString()), 0666)
	//return nil
}

// Load the deck from a file
func LoadFromFile(filename string) deck {
	//bs is short for byte slice (i.e. []byte) and err is short for error
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Errorf("Error: %v", err)
		return nil
	}
	return deck(strings.Split(string(bs), ","))
}

// Shuffle the deck
func (d deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano()) //Create a new source for the random number generator using the current time
	r := rand.New(source)                           //Create a new random number generator using the source

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

// Return the number of cards in the deck
func (d deck) Size() int {
	return len(d)
}
