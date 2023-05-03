package deck_test

import (
	"testing"
	"third/deck"
)

var deckForTesting = []string{
	"Ace of Diamonds",
	"Two of Diamonds",
	"Three of Diamonds",
	"Four of Diamonds",
	"Five of Diamonds",
	"Six of Diamonds",
	"Seven of Diamonds",
	"Eight of Diamonds",
	"Nine of Diamonds",
	"Ten of Diamonds",
	"Jack of Diamonds",
	"Queen of Diamonds",
	"King of Diamonds",
	"Ace of Hearts",
	"Two of Hearts",
	"Three of Hearts",
	"Four of Hearts",
	"Five of Hearts",
	"Six of Hearts",
	"Seven of Hearts",
	"Eight of Hearts",
	"Nine of Hearts",
	"Ten of Hearts",
	"Jack of Hearts",
	"Queen of Hearts",
	"King of Hearts",
	"Ace of Clubs",
	"Two of Clubs",
	"Three of Clubs",
	"Four of Clubs",
	"Five of Clubs",
	"Six of Clubs",
	"Seven of Clubs",
	"Eight of Clubs",
	"Nine of Clubs",
	"Ten of Clubs",
	"Jack of Clubs",
	"Queen of Clubs",
	"King of Clubs",
	"Ace of Spades",
	"Two of Spades",
	"Three of Spades",
	"Four of Spades",
	"Five of Spades",
	"Six of Spades",
	"Seven of Spades",
	"Eight of Spades",
	"Nine of Spades",
	"Ten of Spades",
	"Jack of Spades",
	"Queen of Spades",
	"King of Spades"}

func TestNew(t *testing.T) {
	cards := deck.New()
	if len(cards) != 52 {
		t.Error("Expected 52 cards, got", len(cards))
	}
	for i, card := range cards {
		if card != deckForTesting[i] {
			t.Error("Expected", deckForTesting[i], "but got", card)
		}
	}
}

func TestDeal(t *testing.T) {
	cards := deck.New()
	leftHand, rightHand := deck.Deal(cards, 5)
	if len(leftHand) != 5 {
		t.Error("Expected 5 cards in left hand, got", len(leftHand))
	}
	if len(rightHand) != 47 {
		t.Error("Expected 47 cards in right hand, got", len(rightHand))
	}
	for i, card := range leftHand {
		if card != deckForTesting[i] {
			t.Error("Expected", deckForTesting[i], "but got", card)
		}
	}
	for i, card := range rightHand {
		if card != deckForTesting[i+5] {
			t.Error("Expected", deckForTesting[i+5], "but got", card)
		}
	}
}
