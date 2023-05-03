package cards_test

import (
	"second/cards"
	"testing"
)

var CardsForTesting := []string{"Jack of Diamonds", "Queen of Hearts", "King of Clubs", "Ace of Spades"}

func TestCard(t *testing.T) {
	for i, card := range CardsForTesting {
		if cards.Card((i+10)%13, i) != card {
			t.Error("Expected", card, "but got", cards.Card((i+10)%13, i))
		}
	}
}
