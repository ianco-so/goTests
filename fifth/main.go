package main

import (
	"fifth/deck"
	"fmt"
)

func main() {
	cards := deck.New()

	cards.SaveToFile("my_cards")
	_ = deck.LoadFromFile("my_cards")
	fmt.Println(cards.ToString())
	//os.Remove("my_cards")
}
