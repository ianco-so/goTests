package main

import (
	"fourth/deck"
	"os"
)

func main() {
	cards := deck.New()

	cards.SaveToFile("my_cards")
	_ = deck.LoadFromFile("my_cards")
	os.Remove("my_cards")
}
