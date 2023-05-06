package deck_test

import (
	"fifth/deck"
	"os"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	//SETUP DECK
	deck.SetSuits([]string{"Hearts", "Spades"})
	deck.SetRanks([]string{"Ace", "Jack", "Queen", "King"})
	cards := deck.New()

	sizeSuits := len(deck.GetSuits())
	sizeRanks := len(deck.GetRanks())

	if len(cards) != sizeSuits*sizeRanks {
		t.Error("Expected ", sizeSuits*sizeRanks, " cards, got ", len(cards))
	}
	//Test if the cards are correct.
	for i, suit := range deck.GetSuits() {
		for j, rank := range deck.GetRanks() {
			//card_i := deck.card{rank: rank, suit: suit}
			if cards[i*sizeRanks+j].ToString() != rank+" of "+suit {
				t.Error("Expected ", rank, " of ", suit, " but got ", cards[i*sizeRanks+j].ToString())
			}
		}
	}
}

// Test if the deck size is correct.
func TestSize(t *testing.T) {
	//SETUP DECK
	deck.SetSuits([]string{"Hearts", "Spades"})
	deck.SetRanks([]string{"Ace", "Jack", "Queen", "King"})
	cards := deck.New()

	if cards.Size() != 8 {
		t.Error("Expected a deck of size 8, got ", cards.Size())
	}
}

func TestDeal(t *testing.T) {
	//SETUP DECK
	deck.SetSuits([]string{"Hearts", "Spades"})
	deck.SetRanks([]string{"Ace", "Jack", "Queen", "King"})
	cards := deck.New()

	leftHandSize := 3

	leftHand, rightHand := deck.Deal(cards, leftHandSize)
	if len(leftHand) != leftHandSize {
		t.Error("Expected ", leftHandSize, " cards in left hand, got ", len(leftHand))
	}
	if len(rightHand) != len(cards)-leftHandSize {
		t.Error("Expected ", len(cards)-leftHandSize, " cards in right hand, got ", len(rightHand))
	}

	//Test if the cards are the same.
	for i, lh := range leftHand {
		if !lh.Equals(cards[i]) {
			t.Error("Expected ", cards[i].ToString(), " in left hand, got ", lh.ToString())
		}
	}
	for i, rh := range rightHand {
		if !rh.Equals(cards[i+leftHandSize]) {
			t.Error("Expected ", cards[i+leftHandSize].ToString(), " in right hand, got ", rh.ToString())
		}
	}
}

func TestToString(t *testing.T) {
	//SETUP DECK
	deck.SetSuits([]string{"Hearts", "Spades"})
	deck.SetRanks([]string{"Ace", "Jack", "Queen", "King"})
	cards := deck.New()

	var strSliceForTest []string
	for _, suit := range deck.GetSuits() {
		for _, rank := range deck.GetRanks() {
			strSliceForTest = append(strSliceForTest, rank+" of "+suit)
		}
	}
	strForTest := strings.Join(strSliceForTest, "\n")
	if strForTest != cards.ToString() {
		t.Error("Expected\n", strForTest, "\n AND GOT:\n\n", cards.ToString())
	}
}
func TestSaveToFile(t *testing.T) {
	//SETUP DECK
	deck.SetSuits([]string{"Hearts", "Spades"})
	deck.SetRanks([]string{"Ace", "Jack", "Queen", "King"})
	cards := deck.New()
	//Delete the file if it exists.
	os.Remove("_deck_test.txt")

	e := cards.SaveToFile("_deck_test.txt")
	if e != nil {
		t.Error("Expected no error, got", e)
	}
	//Test if the file was created and it has contents.
	if _, err := os.Stat("_deck_test.txt"); os.IsNotExist(err) {
		t.Error("Expected the file to exist, got", err)
	}
	//Test if the contents of the file are correct.
	fileContents, err2 := os.ReadFile("_deck_test.txt")
	if err2 != nil {
		t.Error("Expected no error, got", err2)
	}
	if string(fileContents) != cards.ToString() {
		t.Error("Expected", cards.ToString()[:10], "... but got", string(fileContents), "...")
	}
	//Delete the file after the test is done.
	err3 := os.Remove("_deck_test.txt")
	if err3 != nil {
		t.Error("Expected no error, got", err3)
	}
}

func TestLoadFromFile(t *testing.T) {
	//SETUP DECK
	deck.SetSuits([]string{"Hearts", "Spades"})
	deck.SetRanks([]string{"Ace", "Jack", "Queen", "King"})
	cards := deck.New()
	//Delete the file if it exists.
	os.Remove("__deck_test.txt")

	//Save the deck to a file
	e := cards.SaveToFile("__deck_test.txt")
	if e != nil {
		t.Error("Expected no error, got", e)
	}

	//Load the deck from the file
	cardsFromFile := deck.LoadFromFile("__deck_test.txt")

	//Test if the deck loaded from the file is correct.
	if len(cardsFromFile) != len(cards) {
		t.Error("Expected ", len(cards), " cards, got", len(cardsFromFile))
	}
	for i, card := range cardsFromFile {
		if card != cards[i] {
			t.Error("Expected", cards[i], "but got", card)
		}
	}

	//Delete the file after the test is done.
	e = os.Remove("__deck_test.txt")
	if e != nil {
		t.Error("Expected no error, got", e)
	}

	//Test if has no file to load from, it'll return nil.
	cardsFromFile = deck.LoadFromFile("__deck_test.txt")
	if cardsFromFile != nil {
		t.Error("Expected nil, got", cardsFromFile.ToString()[:10], "...")
	}
}

func TestShuffle(t *testing.T) {
	//SETUP DECK
	deck.SetSuits([]string{"Hearts", "Spades"})
	deck.SetRanks([]string{"Ace", "Jack", "Queen", "King"})
	cards := deck.New()

	//Create a copy of the deck
	cardsCopy := deck.New()

	//Shuffle the deck
	cards.Shuffle()
	//Test if the deck is shuffled
	if cards.ToString() == cardsCopy.ToString() {
		t.Error("Expected the deck to be shuffled, got", cards.ToString())
	}
}
