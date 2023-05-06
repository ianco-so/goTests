package deck

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var suits = []string{"Diamonds", "Hearts", "Clubs", "Spades"}
var ranks = []string{"Ace", "Jack", "Queen", "King"}

// Create a new type of 'deck' which is a slice of strings
type card struct {
	rank string
	suit string
}

type deck []card //A deck is a slice of cards

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
	deck_i := deck{}

	for _, suit_i := range suits {
		for _, rank_i := range ranks {
			deck_i = append(deck_i, card{rank: rank_i, suit: suit_i})
		}
	}
	return deck_i
}

// Splits the deck into two hands, the first of which is the specified size and the second of which is the remainder
func Deal(d deck, leftHandSize int) (deck, deck) {
	return d[:leftHandSize], d[leftHandSize:]
}

// Convert the deck to a string
func (d deck) ToString() string {
	var deckString []string
	for _, card_i := range d {
		deckString = append(deckString, card_i.ToString())
	}
	return strings.Join(deckString, "\n")
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
	//return deck(strings.Split(string(bs), ","))
	deck := deck{}
	stringDeck := strings.Split(string(bs), "\n")
	for _, card_s := range stringDeck {
		card_temp := strings.Split(card_s, " of ")
		deck = append(deck, card{rank: card_temp[0], suit: card_temp[1]})
	}
	return deck
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

func (c card) Equals(c2 card) bool {
	return c.rank == c2.rank && c.suit == c2.suit
}

func (c card) ToString() string {
	return c.rank + " of " + c.suit
}
